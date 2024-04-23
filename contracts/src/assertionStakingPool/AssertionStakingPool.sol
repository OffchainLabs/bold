// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/OffchainLabs/nitro-contracts/blob/main/LICENSE
// SPDX-License-Identifier: BUSL-1.1
//

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../rollup/IRollupLogic.sol";
import "./AbsBoldStakingPool.sol";

/// @notice Staking pool contract for target assertion.
/// Allows users to deposit stake, create assertion once required stake amount is reached,
/// and reclaim their stake when and if the assertion is confirmed.
contract AssertionStakingPool is AbsBoldStakingPool {
    using SafeERC20 for IERC20;
    address public immutable rollup;
    bytes32 public immutable assertionHash;

    /// @param _rollup Rollup contract of target chain
    /// @param _assertionHash Assertion hash to be passed into Rollup.stakeOnNewAssertion
    constructor(
        address _rollup,
        bytes32 _assertionHash
    ) AbsBoldStakingPool(IERC20(IRollupCore(_rollup).stakeToken())) {
        rollup = _rollup;
        assertionHash = _assertionHash;
    }

    /// @notice Create assertion. Callable only if required stake has been reached and assertion has not been asserted yet.
    function createAssertion(AssertionInputs calldata assertionInputs) external {
        uint256 requiredStake = assertionInputs.beforeStateData.configData.requiredStake;
        // approve spending from rollup for newStakeOnNewAssertion call
        stakeToken.safeIncreaseAllowance(rollup, requiredStake);
        // reverts if pool doesn't have enough stake and if assertion has already been asserted
        IRollupUser(rollup).newStakeOnNewAssertion(requiredStake, assertionInputs, assertionHash);
    }

    /// @notice Make stake withdrawable.
    /// @dev Separate call from withdrawStakeBackIntoPool since returnOldDeposit reverts with 0 balance (in e.g., case of admin forceRefundStaker)
    function makeStakeWithdrawable() public {
        // this checks for active staker
        IRollupUser(rollup).returnOldDeposit();
    }

    /// @notice Move stake back from rollup contract to this contract.
    /// Callable only if this contract has already created an assertion and it's now inactive.
    /// @dev Separate call from makeStakeWithdrawable since returnOldDeposit reverts with 0 balance (in e.g., case of admin forceRefundStaker)
    function withdrawStakeBackIntoPool() public {
        IRollupUser(rollup).withdrawStakerFunds();
    }

    /// @notice Combines makeStakeWithdrawable and withdrawStakeBackIntoPool into single call
    function makeStakeWithdrawableAndWithdrawBackIntoPool() external {
        makeStakeWithdrawable();
        withdrawStakeBackIntoPool();
    }
}
