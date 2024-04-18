// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/OffchainLabs/nitro-contracts/blob/main/LICENSE
// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.0;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/// @notice Abstract contract for handling deposits and withdrawals of trustless edge/assertion staking pools.
/// @dev    The total deposited amount can exceed the required stake amount. 
///         If the total deposited amount exceeds the required amount, any depositor can withdraw some stake early even after the protocol move has been made.
///         This is okay because the protocol move will still be created once the required stake amount is reached, 
///         and all depositors will still be eventually refunded.
abstract contract AbsBoldStakingPool {
    using SafeERC20 for IERC20;

    IERC20 public immutable stakeToken;
    
    mapping(address => uint256) public depositedTokenBalances;

    event StakeDeposited(address indexed sender, uint256 amount);
    event StakeWithdrawn(address indexed sender, uint256 amount);

    /// @notice Cannot deposit or withdraw zero amount
    error ZeroAmount();
    /// @notice Amount exceeds balance
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

        depositedTokenBalances[msg.sender] += amount;
        stakeToken.safeTransferFrom(msg.sender, address(this), amount);

        emit StakeDeposited(msg.sender, amount);
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
}
