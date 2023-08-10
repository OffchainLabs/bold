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
    function createPoolForAssertion(
        address _rollup,
        AssertionInputs memory _assertionInputs,
        bytes32 _expectedAssertionHash
    ) external {
        new AssertionStakingPool(_rollup, _assertionInputs, _expectedAssertionHash);
    }
}

contract AssertionStakingPool {
    using SafeERC20 for IERC20;

    address public immutable rollup;
    bytes32 public immutable assertionHash;
    AssertionInputs public assertionInputs;
    IERC20 public immutable stakeToken;
    mapping(address => uint256) public depositedTokenBalances;

    PoolState public poolState = PoolState.PENDING;

    constructor(
        address _rollup,
        AssertionInputs memory _assertionInputs,
        bytes32 _assertionHash
    ) {
        rollup = _rollup;
        assertionHash = _assertionHash;
        assertionInputs = _assertionInputs;
        stakeToken = IERC20(IRollupCore(rollup).stakeToken());

        stakeToken.approve(rollup, getRequiredStake());
    }

    function depositIntoPool(uint256 _amount) external {
        if (poolState != PoolState.PENDING) {
            revert PoolNotInPendingState(poolState);
        }
        uint256 requiredStake = getRequiredStake();

        uint256 currentPoolBalance = stakeToken.balanceOf(address(this));
        if (currentPoolBalance >= requiredStake) {
            revert PoolStakeAlreadyReached(requiredStake);
        }

        uint256 amountToTransfer;

        if (currentPoolBalance + _amount > requiredStake) {
            amountToTransfer = requiredStake - currentPoolBalance;
        } else {
            amountToTransfer = _amount;
        }

        depositedTokenBalances[msg.sender] += amountToTransfer;
        stakeToken.safeTransferFrom(msg.sender, address(this), amountToTransfer);
    }

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
        IRollupUser(rollup).newStakeOnNewAssertion(requiredStake, assertionInputs, assertionHash);
    }

    function returnOldStakeBackToPool() external {
        if (poolState != PoolState.ASSERTED) {
            revert PoolNotInAssertedState(poolState);
        }

        poolState = PoolState.CONFIRMED;
        IRollupUser(rollup).returnOldDeposit();
        IRollupUser(rollup).withdrawStakerFunds();
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

    function getRequiredStake() public view returns (uint256) {
        return assertionInputs.beforeStateData.configData.requiredStake;
    }
}
