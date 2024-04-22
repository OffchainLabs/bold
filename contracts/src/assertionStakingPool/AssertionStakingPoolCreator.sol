// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/OffchainLabs/nitro-contracts/blob/main/LICENSE
// SPDX-License-Identifier: BUSL-1.1
//

pragma solidity ^0.8.0;

import "./AssertionStakingPool.sol";
import "./StakingPoolCreatorUtils.sol";

/// @notice Creates staking pool contract for a target assertion. Can be used for any child Arbitrum chain running on top of the deployed AssertionStakingPoolCreator's chain.
contract AssertionStakingPoolCreator {
    event NewAssertionPoolCreated(
        address indexed rollup,
        bytes32 indexed _assertionHash,
        address assertionPool
    );

    /// @notice Create a staking pool contract
    /// @param _rollup Rollup contract of target chain
    /// @param _assertionInputs Inputs to be passed into Rollup.stakeOnNewAssertion
    /// @param _assertionHash Assertion hash to be passed into Rollup.stakeOnNewAssertion
    function createPool(
        address _rollup,
        AssertionInputs memory _assertionInputs,
        bytes32 _assertionHash
    ) external returns (AssertionStakingPool) {
        AssertionStakingPool assertionPool = new AssertionStakingPool{salt: 0}(_rollup, _assertionInputs, _assertionHash);
        emit NewAssertionPoolCreated(_rollup, _assertionHash, address(assertionPool));
        return assertionPool;
    }

    /// @notice get staking pool deployed with provided inputs; reverts if pool contract doesn't exist.
    /// @param _rollup Rollup contract of target chain
    /// @param _assertionInputs Inputs to be passed into Rollup.stakeOnNewAssertion
    /// @param _assertionHash Assertion hash to be passed into Rollup.stakeOnNewAssertion
    function getPool(
        address _rollup,
        AssertionInputs memory _assertionInputs,
        bytes32 _assertionHash
    ) public view returns (AssertionStakingPool) {
        return AssertionStakingPool(
            StakingPoolCreatorUtils.getPool(
                type(AssertionStakingPool).creationCode, 
                abi.encode(_rollup, _assertionInputs, _assertionHash)
            )
        );
    }
}
