// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE
// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol";

import {IRollupUser} from "./IRollupLogic.sol";
import "../libraries/UUPSNotUpgradeable.sol";
import "./RollupCore.sol";
import "./IRollupLogic.sol";
import {ETH_POS_BLOCK_TIME} from "../libraries/Constants.sol";

abstract contract AbsRollupUserLogic is RollupCore, UUPSNotUpgradeable, IRollupUserAbs {
    using AssertionNodeLib for AssertionNode;
    using GlobalStateLib for GlobalState;

    modifier onlyValidator() {
        require(isValidator[msg.sender] || validatorWhitelistDisabled, "NOT_VALIDATOR");
        _;
    }

    uint256 internal immutable deployTimeChainId = block.chainid;

    function _chainIdChanged() internal view returns (bool) {
        return deployTimeChainId != block.chainid;
    }

    /**
     * @notice Extra number of blocks the validator can remain inactive before considered inactive
     *         This is 7 days assuming a 13.2 seconds block time
     */
    uint256 public constant VALIDATOR_AFK_BLOCKS = 45818;

    function _validatorIsAfk() internal view returns (bool) {
        AssertionNode memory latestAssertion = getAssertionStorage(latestAssertionCreated());
        if (latestAssertion.createdAtBlock == 0) return false;
        if (latestAssertion.createdAtBlock + confirmPeriodBlocks + VALIDATOR_AFK_BLOCKS < block.number) {
            return true;
        }
        return false;
    }

    function removeWhitelistAfterFork() external {
        require(!validatorWhitelistDisabled, "WHITELIST_DISABLED");
        require(_chainIdChanged(), "CHAIN_ID_NOT_CHANGED");
        validatorWhitelistDisabled = true;
    }

    function removeWhitelistAfterValidatorAfk() external {
        require(!validatorWhitelistDisabled, "WHITELIST_DISABLED");
        require(_validatorIsAfk(), "VALIDATOR_NOT_AFK");
        validatorWhitelistDisabled = true;
    }

    function isERC20Enabled() public view override returns (bool) {
        return stakeToken != address(0);
    }

    /**
     * @notice Reject the next unresolved assertion
     * @param winningEdgeId The winning challenge edge of the prev's succession challenge
     */
    function rejectNextAssertion(bytes32 winningEdgeId) external onlyValidator whenNotPaused {
        requireUnresolvedExists();
        uint64 latestConfirmedAssertionNum = latestConfirmed();
        uint64 firstUnresolvedAssertionNum = firstUnresolvedAssertion();
        AssertionNode storage firstUnresolvedAssertion_ = getAssertionStorage(firstUnresolvedAssertionNum);

        if (firstUnresolvedAssertion_.prevNum == latestConfirmedAssertionNum) {
            /**
             * If the first unresolved assertion is a child of the latest confirmed assertion, to prove it can be rejected, we show:
             * a) Its prev's child confirmation deadline has expired
             * b) There is more than 1 child on the prev
             * c) The assertion is not the winner of the prev's succession challenge
             */

            getAssertionStorage(latestConfirmedAssertionNum).requirePastChildConfirmDeadline();
            getAssertionStorage(latestConfirmedAssertionNum).requireMoreThanOneChild();

            ChallengeEdge memory winningEdge = challengeManager.getEdge(winningEdgeId);
            require(winningEdge.status == EdgeStatus.Confirmed, "EDGE_NOT_CONFIRMED");
            require(winningEdge.eType == EdgeType.Block, "EDGE_NOT_BLOCK_TYPE");
            require(
                winningEdge.originId == getAssertionStorage(latestConfirmedAssertionNum).assertionHash,
                "EDGE_NOT_FROM_PREV"
            );
            require(winningEdge.claimId != firstUnresolvedAssertion_.assertionHash, "IS_WINNER");
        }
        // Simpler case: if the first unreseolved assertion doesn't point to the last confirmed assertion, another branch was confirmed and can simply reject it outright
        _rejectNextAssertion();

        emit AssertionRejected(firstUnresolvedAssertionNum);
    }

    /**
     * @notice Confirm the next unresolved assertion
     * @param blockHash The block hash at the end of the assertion
     * @param sendRoot The send root at the end of the assertion
     * @param winningEdgeId The winning edge id if a challenge is started
     */
    function confirmNextAssertion(bytes32 blockHash, bytes32 sendRoot, bytes32 winningEdgeId)
        external
        onlyValidator
        whenNotPaused
    {
        requireUnresolvedExists();

        uint64 assertionNum = firstUnresolvedAssertion();
        AssertionNode storage assertion = getAssertionStorage(assertionNum);

        // Check that prev is latest confirmed
        assert(assertion.prevNum == latestConfirmed());

        AssertionNode storage prevAssertion = getAssertionStorage(assertion.prevNum);
        prevAssertion.requirePastChildConfirmDeadline();

        if (prevAssertion.secondChildBlock > 0) {
            // check if assertion is the challenge winner
            ChallengeEdge memory winningEdge = challengeManager.getEdge(winningEdgeId);
            require(getAssertionNum(winningEdge.claimId) == assertionNum, "NOT_WINNER");
            require(winningEdge.status == EdgeStatus.Confirmed, "EDGE_NOT_CONFIRMED");
        }

        confirmAssertion(assertionNum, blockHash, sendRoot);
    }

    /**
     * @notice Create a new stake
     * @param depositAmount The amount of either eth or tokens staked
     */
    function _newStake(uint256 depositAmount) internal onlyValidator whenNotPaused {
        // Verify that sender is not already a staker
        require(!isStaked(msg.sender), "ALREADY_STAKED");
        // TODO: HN: review this logic
        // require(!isZombie(msg.sender), "STAKER_IS_ZOMBIE");
        require(depositAmount >= currentRequiredStake(), "NOT_ENOUGH_STAKE");

        createNewStake(msg.sender, depositAmount);
    }

    /**
     * @notice Create a new assertion and move stake onto it
     * @param assertion The assertion data
     * @param expectedAssertionHash The hash of the assertion being created (protects against reorgs)
     */
    function stakeOnNewAssertion(AssertionInputs calldata assertion, bytes32 expectedAssertionHash)
        public
        onlyValidator
        whenNotPaused
    {
        // Early revert on duplicated assertion if expectedAssertionHash is set
        require(
            expectedAssertionHash == bytes32(0) || !isAssertionExists(expectedAssertionHash), "EXPECTED_ASSERTION_SEEN"
        );

        require(isStaked(msg.sender), "NOT_STAKED");
        require(amountStaked(msg.sender) >= currentRequiredStake(), "INSUFFICIENT_STAKE");
        // Staker can create new assertion only if
        // a) its last staked assertion is the prev; or
        // b) its last staked assertion have a child
        uint64 prevAssertion = getAssertionNum(
            RollupLib.assertionHash(
                assertion.beforeStateData.prevAssertionHash,
                assertion.beforeState,
                assertion.beforeStateData.sequencerBatchAcc,
                assertion.beforeStateData.wasmRoot
            )
        ); // TODO: HN: we calculated this hash again in createNewAssertion
        {
            uint64 lastAssertion = latestStakedAssertion(msg.sender);
            require(
                lastAssertion == prevAssertion || getAssertionStorage(lastAssertion).firstChildBlock > 0,
                "STAKED_ON_ANOTHER_BRANCH"
            );
        }

        {
            uint256 timeSincePrev = block.number - getAssertion(prevAssertion).createdAtBlock;
            // Verify that assertion meets the minimum Delta time requirement
            require(timeSincePrev >= minimumAssertionPeriod, "TIME_DELTA");

            // CHRIS: TODO: this is an extra storage call
            // CHRIS: TODO: we should be doing this inside the createNewAssertion call
            //              since otherwise an admin created assertion would be challengeable if created with the wrong count
            uint64 prevAssertionNextInboxPosition = getAssertionStorage(prevAssertion).nextInboxPosition;

            // Minimum size requirement: any assertion must consume exactly all inbox messages
            // put into L1 inbox before the prev node’s L1 blocknum.
            // We make an exception if the machine enters the errored state,
            // as it can't consume future batches.
            require(
                assertion.afterState.machineStatus == MachineStatus.ERRORED
                    || assertion.afterState.globalState.getInboxPosition() == prevAssertionNextInboxPosition,
                "WRONG_INBOX_POS"
            );

            // The rollup cannot advance normally from an errored state
            // CHRIS: TODO: this is interesting? How do we recover from errored state?
            require(assertion.beforeState.machineStatus == MachineStatus.FINISHED, "BAD_PREV_STATUS");
        }
        createNewAssertion(assertion, prevAssertion, expectedAssertionHash);

        stakeOnAssertion(msg.sender, latestAssertionCreated());

        if (getAssertionStorage(prevAssertion).secondChildBlock > 0) {
            // only 1 of the children can be confirmed and get their stake refunded
            // so we send the other child's stake to the excess stake receiver
            // TODO: HN: if the losing staker have staked more than currentRequiredStake, the excess stake will be stuck
            // TODO: HN: use prev stake amount instead of currentRequiredStake
            increaseWithdrawableFunds(excessStakeReceiver, currentRequiredStake());
        }
    }

    /**
     * @notice Refund a staker that is currently staked on or before the latest confirmed assertion
     * @param stakerAddress Address of the staker whose stake is refunded
     */
    function returnOldDeposit(address stakerAddress) external override onlyValidator whenNotPaused {
        requireInactiveStaker(msg.sender);
        withdrawStaker(stakerAddress);
    }

    /**
     * @notice Increase the amount staked for the given staker
     * @param stakerAddress Address of the staker whose stake is increased
     * @param depositAmount The amount of either eth or tokens deposited
     */
    function _addToDeposit(address stakerAddress, uint256 depositAmount) internal onlyValidator whenNotPaused {
        increaseStakeBy(stakerAddress, depositAmount);
    }

    /**
     * @notice Reduce the amount staked for the sender (difference between initial amount staked and target is creditted back to the sender).
     * @param target Target amount of stake for the staker. If this is below the current minimum, it will be set to minimum instead
     */
    function reduceDeposit(uint256 target) external onlyValidator whenNotPaused {
        requireInactiveStaker(msg.sender);
        uint256 currentRequired = currentRequiredStake();
        if (target < currentRequired) {
            target = currentRequired;
        }
        reduceStakeTo(msg.sender, target);
    }

    /**
     * @notice Calculate the current amount of funds required to place a new stake in the rollup
     * @return The current minimum stake requirement
     */
    function requiredStake(uint256 blockNumber, uint64 firstUnresolvedAssertionNum, uint64 latestCreatedAssertion)
        external
        view
        returns (uint256)
    {
        return baseStake;
    }

    function owner() external view returns (address) {
        return _getAdmin();
    }

    function currentRequiredStake() public view returns (uint256) {
        return baseStake;
    }

    /**
     * @notice Verify that there are some number of assertions still unresolved
     */
    function requireUnresolvedExists() public view override {
        uint256 firstUnresolved = firstUnresolvedAssertion();
        require(firstUnresolved > latestConfirmed() && firstUnresolved <= latestAssertionCreated(), "NO_UNRESOLVED");
    }

    function requireUnresolved(uint256 assertionNum) public view override {
        require(assertionNum >= firstUnresolvedAssertion(), "ALREADY_DECIDED");
        require(assertionNum <= latestAssertionCreated(), "DOESNT_EXIST");
    }

    /**
     * @notice Verify that the given staker is not active
     * @param stakerAddress Address to check
     */
    function requireInactiveStaker(address stakerAddress) private view {
        require(isStaked(stakerAddress), "NOT_STAKED");
        // A staker is inactive if
        // a) their last staked assertion is the latest confirmed assertion
        // b) their last staked assertion have a child
        uint64 lastestAssertion = latestStakedAssertion(stakerAddress);
        bool isLatestConfirmed = lastestAssertion == latestConfirmed();
        bool haveChild = getAssertionStorage(lastestAssertion).firstChildBlock > 0;
        require(isLatestConfirmed || haveChild, "STAKE_ACTIVE");
    }
}

contract RollupUserLogic is AbsRollupUserLogic, IRollupUser {
    /// @dev the user logic just validated configuration and shouldn't write to state during init
    /// this allows the admin logic to ensure consistency on parameters.
    function initialize(address _stakeToken) external view override onlyProxy {
        require(_stakeToken == address(0), "NO_TOKEN_ALLOWED");
        require(!isERC20Enabled(), "FACET_NOT_ERC20");
    }

    /**
     * @notice Create a new stake on a new assertion
     * @param assertion Assertion describing the state change between the old assertion and the new one
     * @param expectedAssertionHash Assertion hash of the assertion that will be created
     */
    function newStakeOnNewAssertion(AssertionInputs calldata assertion, bytes32 expectedAssertionHash)
        external
        payable
        override
    {
        _newStake(msg.value);
        stakeOnNewAssertion(assertion, expectedAssertionHash);
    }

    /**
     * @notice Increase the amount staked eth for the given staker
     * @param stakerAddress Address of the staker whose stake is increased
     */
    function addToDeposit(address stakerAddress) external payable override onlyValidator whenNotPaused {
        _addToDeposit(stakerAddress, msg.value);
    }

    /**
     * @notice Withdraw uncommitted funds owned by sender from the rollup chain
     */
    function withdrawStakerFunds() external override onlyValidator whenNotPaused returns (uint256) {
        uint256 amount = withdrawFunds(msg.sender);
        // This is safe because it occurs after all checks and effects
        // solhint-disable-next-line avoid-low-level-calls
        (bool success,) = msg.sender.call{value: amount}("");
        require(success, "TRANSFER_FAILED");
        return amount;
    }
}

contract ERC20RollupUserLogic is AbsRollupUserLogic, IRollupUserERC20 {
    /// @dev the user logic just validated configuration and shouldn't write to state during init
    /// this allows the admin logic to ensure consistency on parameters.
    function initialize(address _stakeToken) external view override onlyProxy {
        require(_stakeToken != address(0), "NEED_STAKE_TOKEN");
        require(isERC20Enabled(), "FACET_NOT_ERC20");
    }

    /**
     * @notice Create a new stake on a new assertion
     * @param tokenAmount Amount of the rollups staking token to stake
     * @param assertion Assertion describing the state change between the old assertion and the new one
     * @param expectedAssertionHash Assertion hash of the assertion that will be created
     */
    function newStakeOnNewAssertion(
        uint256 tokenAmount,
        AssertionInputs calldata assertion,
        bytes32 expectedAssertionHash
    ) external override {
        _newStake(tokenAmount);
        stakeOnNewAssertion(assertion, expectedAssertionHash);
        /// @dev This is an external call, safe because it's at the end of the function
        receiveTokens(tokenAmount);
    }

    /**
     * @notice Increase the amount staked tokens for the given staker
     * @param stakerAddress Address of the staker whose stake is increased
     * @param tokenAmount the amount of tokens staked
     */
    function addToDeposit(address stakerAddress, uint256 tokenAmount) external onlyValidator whenNotPaused {
        _addToDeposit(stakerAddress, tokenAmount);
        /// @dev This is an external call, safe because it's at the end of the function
        receiveTokens(tokenAmount);
    }

    /**
     * @notice Withdraw uncommitted funds owned by sender from the rollup chain
     */
    function withdrawStakerFunds() external override onlyValidator whenNotPaused returns (uint256) {
        uint256 amount = withdrawFunds(msg.sender);
        // This is safe because it occurs after all checks and effects
        require(IERC20Upgradeable(stakeToken).transfer(msg.sender, amount), "TRANSFER_FAILED");
        return amount;
    }

    function receiveTokens(uint256 tokenAmount) private {
        require(IERC20Upgradeable(stakeToken).transferFrom(msg.sender, address(this), tokenAmount), "TRANSFER_FAIL");
    }
}
