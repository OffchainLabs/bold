// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/OffchainLabs/nitro-contracts/blob/main/LICENSE
// SPDX-License-Identifier: BUSL-1.1
//

pragma solidity ^0.8.0;

import {AbsBoldStakingPool} from "./AbsBoldStakingPool.sol";

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "./StakingPoolErrors.sol";

/// @notice Staking pool contract for target assertion.
/// Allows users to deposit stake, create assertion once required stake amount is reached,
/// and reclaim their stake when and if the assertion is confirmed.
contract AssertionStakingPool is AbsBoldStakingPool {
    using SafeERC20 for IERC20;
    address public immutable rollup;
    bytes32 public immutable assertionHash;
    AssertionInputs public assertionInputs;

    /// @param _rollup Rollup contract of target chain
    /// @param _assertionInputs Inputs to be passed into Rollup.stakeOnNewAssertion
    /// @param _assertionHash Assertion hash to be passed into Rollup.stakeOnNewAssertion
    constructor(
        address _rollup,
        AssertionInputs memory _assertionInputs,
        bytes32 _assertionHash
    ) AbsBoldStakingPool(IERC20(IRollupCore(_rollup).stakeToken())) {
        rollup = _rollup;
        assertionHash = _assertionHash;
        assertionInputs = _assertionInputs;
    }

    /// @notice Create assertion. Callable only if required stake has been reached and assertion has not been asserted yet.
    function createMove() external override {
        uint256 requiredStake = getRequiredStake();
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
    function withdrawStakeBackIntoPool() external override {
        IRollupUser(rollup).withdrawStakerFunds();
    }

    /// @notice Get required stake for pool's assertion.
    /// Requried stake for a given assertion is set in the previous assertion's config data
    function getRequiredStake() public view override returns (uint256) {
        return assertionInputs.beforeStateData.configData.requiredStake;
    }
}
