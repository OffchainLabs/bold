// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE
// SPDX-License-Identifier: BUSL-1.1
//

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../rollup/IRollupLogic.sol";
import "../rollup/IRollupCore.sol";
import "./StakingPoolErrors.sol";

/// @notice Staking pool contract for target assertion. Allows users to deposit stake, create assertion once required stake amount is reached, and reclaim their stake when and if the assertion is confirmed.
contract AssertionStakingPool {
    using SafeERC20 for IERC20;
    address public immutable rollup;
    bytes32 public immutable assertionHash;
    AssertionInputs public assertionInputs;
    IERC20 public immutable stakeToken;
    mapping(address => uint256) public depositedTokenBalances;

    PoolState public poolState = PoolState.PENDING;

    event StakeDeposited(address indexed sender, uint256 amount);
    event AssertionCreated();
    event PoolStateInactive();
    event StakeReturned();
    event StakeWithrawn(address indexed sender, uint256 amount);

    /// @param _rollup Rollup contract of target chain
    /// @param _assertionInputs Inputs to be passed into Rollup.stakeOnNewAssertion
    /// @param _assertionHash Assertion hash to be passed into Rollup.stakeOnNewAssertion
    constructor(
        address _rollup,
        AssertionInputs memory _assertionInputs,
        bytes32 _assertionHash
    ) {
        rollup = _rollup;
        assertionHash = _assertionHash;
        assertionInputs = _assertionInputs;
        stakeToken = IERC20(IRollupCore(rollup).stakeToken());
    }

    /// @notice Deposit stake into pool contract. Callable only if assertion has not been asserted yet.
    /// @param _amount amount of stake token to deposit
    function depositIntoPool(uint256 _amount) external {
        if (poolState != PoolState.PENDING) {
            revert PoolNotInPendingState(poolState);
        }

        depositedTokenBalances[msg.sender] += _amount;
        stakeToken.safeTransferFrom(msg.sender, address(this), _amount);
        emit StakeDeposited(msg.sender, _amount);
    }

    /// @notice Create assertion. Callable only if required stake has been reached and assertion has not been asserted yet.
    function createAssertion() external {
        if (poolState != PoolState.PENDING) {
            revert PoolNotInPendingState(poolState);
        }

        uint256 balance = stakeToken.balanceOf(address(this));
        uint256 requiredStake = getRequiredStake();
        if (balance < requiredStake) {
            revert NotEnoughStake(balance, requiredStake);
        }

        poolState = PoolState.ASSERTED;
        // approve spending from rollup for newStakeOnNewAssertion call
        stakeToken.safeIncreaseAllowance(rollup, requiredStake);
        IRollupUser(rollup).newStakeOnNewAssertion(requiredStake, assertionInputs, assertionHash);
        emit AssertionCreated();
    }

    /// @notice update pool state if assertion is no longer active (confirmed or has a child) , and make deposit withdrawable.
    function setPoolStateInactive() external {
        if (poolState != PoolState.ASSERTED) {
            revert PoolNotInAssertedState(poolState);
        }   

        poolState = PoolState.INACTIVE;

        if (!IRollupCore(rollup).stakerIsInactive(address(this))) {
            revert AssertionNotInactive(assertionHash);
        }
        IRollupUser(rollup).returnOldDeposit();
        emit PoolStateInactive();
    }

    /// @notice Move stake back from rollup contract to this contract. Calalble only if this contract has already created an assertion and it's now inactive.
    /// @dev Separate call from setPoolStateInactive since withdrawStakerFunds reverts with 0 balance (in e.g., case of admin forceRefundStaker)
    function returnOldStakeBackToPool() external {
        if (poolState != PoolState.INACTIVE) {
            revert PoolNotInInactiveState(poolState);
        }
        IRollupUser(rollup).withdrawStakerFunds();
        emit StakeReturned();
    }

    /// @notice Send stake from this contract back to its depositor.
    function withdrawFromPool() external {
        uint256 balance = depositedTokenBalances[msg.sender];
        if (balance == 0) {
            revert NoBalanceToWithdraw(msg.sender);
        }
        depositedTokenBalances[msg.sender] = 0;
        stakeToken.safeTransfer(msg.sender, balance);
        emit StakeWithrawn(msg.sender, balance);
    }

    /// @notice Get required stake for pool's assertion. Requried stake for a given assertion is set in the previous assertion's config data
    function getRequiredStake() public view returns (uint256) {
        return assertionInputs.beforeStateData.configData.requiredStake;
    }
}
