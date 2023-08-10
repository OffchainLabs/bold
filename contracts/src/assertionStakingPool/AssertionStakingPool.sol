// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE
// SPDX-License-Identifier: BUSL-1.1
//

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../rollup/IRollupLogic.sol";
import "../rollup/IRollupCore.sol";
import "./StakingPoolErrors.sol";

contract AssertionStakingPoolCreator {
    function createPoolForAssertion(address _rollup, bytes32 _expectedAssertionHash) external {
        new AssertionStakingPool(_rollup, _expectedAssertionHash);
    }
}

contract AssertionStakingPool {
    using SafeERC20 for IERC20;

    address public immutable rollup;
    bytes32 public immutable assertionHash;
    IERC20 public immutable stakeToken;
    uint256 public immutable baseStake;

    mapping(address => uint256) public depositedTokenBalances;

    PoolState public poolState = PoolState.PENDING;

    constructor(address _rollup, bytes32 _assertionHash) {
        rollup = _rollup;
        assertionHash = _assertionHash;
        stakeToken = IERC20(IRollupCore(rollup).stakeToken());
        baseStake = IRollupCore(rollup).baseStake();

        stakeToken.approve(rollup, baseStake);
    }

    function depositIntoPool(uint256 _amount) external {
        if (poolState != PoolState.PENDING) {
            revert PoolNotInPendingState(poolState);
        }

        uint256 currentPoolBalance = stakeToken.balanceOf(address(this));
        if (currentPoolBalance >= baseStake) {
            revert PoolStakeAlreadyReached(baseStake);
        }

        uint256 amountToTransfer;

        if (currentPoolBalance + _amount > baseStake) {
            amountToTransfer = baseStake - currentPoolBalance;
        } else {
            amountToTransfer = _amount;
        }

        stakeToken.safeTransferFrom(msg.sender, address(this), amountToTransfer);
        depositedTokenBalances[msg.sender] += amountToTransfer;
    }

    function createAssertion(AssertionInputs calldata assertionInputs) external {
        if (poolState != PoolState.PENDING) {
            revert PoolNotInPendingState(poolState);
        }

        uint256 balance = stakeToken.balanceOf(address(this));

        if (balance < baseStake) {
            revert NotEnoughStake(balance, baseStake);
        }

        IRollupUser(rollup).newStakeOnNewAssertion(baseStake, assertionInputs, assertionHash);
        poolState = PoolState.ASSERTED;
    }

    function returnOldStakeBackToPool() external {
        if (poolState != PoolState.ASSERTED) {
            revert PoolNotInAssertedState(poolState);
        }

        IRollupUser(rollup).returnOldDeposit();
        poolState = PoolState.CONFIRMED;
    }

    function withdrawFromPool() external {
        if (poolState == PoolState.ASSERTED) {
            revert PoolNotInPendingOrConfirmedState(poolState);
        }
        uint256 balance = depositedTokenBalances[msg.sender];
        if (balance == 0) {
            revert NoBalanceToWithdraw(msg.sender);
        }
        depositedTokenBalances[msg.sender] = 0;
        stakeToken.safeTransfer(msg.sender, balance);
    }
}
