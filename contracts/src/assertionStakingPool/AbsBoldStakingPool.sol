// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/OffchainLabs/nitro-contracts/blob/main/LICENSE
// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.0;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

abstract contract AbsBoldStakingPool {
    using SafeERC20 for IERC20;

    IERC20 public immutable stakeToken;
    mapping(address => uint256) public depositedTokenBalances;
    uint256 public totalDepositedTokens;

    event StakeDeposited(address indexed sender, uint256 amount);
    event StakeWithdrawn(address indexed sender, uint256 amount);

    error RequiredStakeAmountMet();
    error ZeroAmount();
    error AmountExceedsBalance(address sender, uint256 amount, uint256 balance);

    constructor(IERC20 _stakeToken) {
        stakeToken = _stakeToken;
    }

    /// @notice Deposit stake into pool contract.
    /// @param amount amount of stake token to deposit
    function depositIntoPool(uint256 amount) external {
        if (amount == 0) {
            revert ZeroAmount();
        }
        uint256 _totalDepositedTokens = totalDepositedTokens;
        uint256 _requiredStake = getRequiredStake();
        if (_totalDepositedTokens >= _requiredStake) {
            revert RequiredStakeAmountMet();
        }

        uint256 cappedAmount = _totalDepositedTokens + amount > _requiredStake ? _requiredStake - totalDepositedTokens : amount;

        depositedTokenBalances[msg.sender] += cappedAmount;
        totalDepositedTokens = _totalDepositedTokens + cappedAmount;

        stakeToken.safeTransferFrom(msg.sender, address(this), cappedAmount);

        emit StakeDeposited(msg.sender, cappedAmount);
    }

    /// @notice Send supplied amount of stake from this contract back to its depositor.
    /// @param amount stake amount to withdraw
    function withdrawFromPool(uint256 amount) public {
        if (amount == 0) {
            revert ZeroAmount();
        }
        uint256 balance = depositedTokenBalances[msg.sender];
        if (amount > balance) {
            revert AmountExceedsBalance(msg.sender, amount, balance);
        }

        depositedTokenBalances[msg.sender] = balance - amount;
        stakeToken.safeTransfer(msg.sender, amount);
        
        emit StakeWithdrawn(msg.sender, amount);
    }

    /// @notice Send full balance of stake from this contract back to its depositor.
    function withdrawFromPool() external {
        withdrawFromPool(depositedTokenBalances[msg.sender]);
    }

    /// @notice Make the move. This could be creating an assertion or challenge edge.
    function createMove() external virtual;

    /// @notice Reclaim the pool's stake from the rollup or challenge manager to this contract.
    function withdrawStakeBackIntoPool() external virtual;

    /// @notice Get required stake to create a move.
    function getRequiredStake() public virtual view returns (uint256);
}
