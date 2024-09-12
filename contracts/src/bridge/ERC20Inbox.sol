// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/OffchainLabs/nitro-contracts/blob/main/LICENSE
// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.4;

import "./AbsInbox.sol";
import "./IERC20Inbox.sol";
import "./IERC20Bridge.sol";
import "../libraries/AddressAliasHelper.sol";
import {L1MessageType_ethDeposit} from "../libraries/MessageTypes.sol";
import {AmountTooLarge} from "../libraries/Error.sol";
import {MAX_UPSCALE_AMOUNT} from "../libraries/Constants.sol";

import {DecimalsConverterHelper} from "../libraries/DecimalsConverterHelper.sol";

import {AddressUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/AddressUpgradeable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/**
 * @title Inbox for user and contract originated messages
 * @notice Messages created via this inbox are enqueued in the delayed accumulator
 * to await inclusion in the SequencerInbox
 */
contract ERC20Inbox is AbsInbox, IERC20Inbox {
    using SafeERC20 for IERC20;

    constructor(
        uint256 _maxDataSize
    ) AbsInbox(_maxDataSize) {}

    /// @inheritdoc IInboxBase
    function initialize(
        IBridge _bridge,
        ISequencerInbox _sequencerInbox
    ) external initializer onlyDelegated {
        __AbsInbox_init(_bridge, _sequencerInbox);

        // inbox holds native token in transit used to pay for retryable tickets, approve bridge to use it
        address nativeToken = IERC20Bridge(address(bridge)).nativeToken();
        IERC20(nativeToken).safeApprove(address(bridge), type(uint256).max);
    }

    /// @inheritdoc IERC20Inbox
    function depositERC20(
        uint256 amount
    ) public whenNotPaused onlyAllowed returns (uint256) {
        address dest = msg.sender;

        // solhint-disable-next-line avoid-tx-origin
        if (AddressUpgradeable.isContract(msg.sender) || tx.origin != msg.sender) {
            // isContract check fails if this function is called during a contract's constructor.
            dest = AddressAliasHelper.applyL1ToL2Alias(msg.sender);
        }

        uint256 amountToMintOnL2 = _fromNativeTo18Decimals(amount);
        return _deliverMessage(
            L1MessageType_ethDeposit, msg.sender, abi.encodePacked(dest, amountToMintOnL2), amount
        );
    }

    /// @inheritdoc IERC20Inbox
    function createRetryableTicket(
        address to,
        uint256 l2CallValue,
        uint256 maxSubmissionCost,
        address excessFeeRefundAddress,
        address callValueRefundAddress,
        uint256 gasLimit,
        uint256 maxFeePerGas,
        uint256 tokenTotalFeeAmount,
        bytes calldata data
    ) external whenNotPaused onlyAllowed returns (uint256) {
        return _createRetryableTicket(
            to,
            l2CallValue,
            maxSubmissionCost,
            excessFeeRefundAddress,
            callValueRefundAddress,
            gasLimit,
            maxFeePerGas,
            tokenTotalFeeAmount,
            data
        );
    }

    /// @inheritdoc IERC20Inbox
    function unsafeCreateRetryableTicket(
        address to,
        uint256 l2CallValue,
        uint256 maxSubmissionCost,
        address excessFeeRefundAddress,
        address callValueRefundAddress,
        uint256 gasLimit,
        uint256 maxFeePerGas,
        uint256 tokenTotalFeeAmount,
        bytes calldata data
    ) public whenNotPaused onlyAllowed returns (uint256) {
        return _unsafeCreateRetryableTicket(
            to,
            l2CallValue,
            maxSubmissionCost,
            excessFeeRefundAddress,
            callValueRefundAddress,
            gasLimit,
            maxFeePerGas,
            tokenTotalFeeAmount,
            data
        );
    }

    /// @inheritdoc IInboxBase
    function calculateRetryableSubmissionFee(
        uint256,
        uint256
    ) public pure override(AbsInbox, IInboxBase) returns (uint256) {
        // retryable ticket's submission fee is not charged when ERC20 token is used to pay for fees
        return 0;
    }

    function _deliverToBridge(
        uint8 kind,
        address sender,
        bytes32 messageDataHash,
        uint256 tokenAmount
    ) internal override returns (uint256) {
        // Fetch native token from sender if inbox doesn't already hold enough tokens to pay for fees.
        // Inbox might have been pre-funded in prior call, ie. as part of token bridging flow.
        address nativeToken = IERC20Bridge(address(bridge)).nativeToken();
        uint256 inboxNativeTokenBalance = IERC20(nativeToken).balanceOf(address(this));
        if (inboxNativeTokenBalance < tokenAmount) {
            uint256 diff = tokenAmount - inboxNativeTokenBalance;
            IERC20(nativeToken).safeTransferFrom(msg.sender, address(this), diff);
        }

        return IERC20Bridge(address(bridge)).enqueueDelayedMessage(
            kind, AddressAliasHelper.applyL1ToL2Alias(sender), messageDataHash, tokenAmount
        );
    }

    /// @inheritdoc AbsInbox
    function _fromNativeTo18Decimals(
        uint256 value
    ) internal view override returns (uint256) {
        // In order to keep compatibility of child chain's native currency with external 3rd party tooling we
        // expect 18 decimals to be always used for native currency. If native token uses different number of
        // decimals then here it will be normalized to 18. Keep in mind, when withdrawing from child chain back
        // to parent chain then the amount has to match native token's granularity, otherwise it will be rounded
        // down.
        uint8 nativeTokenDecimals = IERC20Bridge(address(bridge)).nativeTokenDecimals();

        // Also make sure that inflated amount does not overflow uint256
        if (nativeTokenDecimals < 18) {
            if (value > MAX_UPSCALE_AMOUNT) {
                revert AmountTooLarge(value);
            }
        }
        return DecimalsConverterHelper.adjustDecimals(value, nativeTokenDecimals, 18);
    }
}
