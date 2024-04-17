// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/OffchainLabs/nitro-contracts/blob/main/LICENSE
// SPDX-License-Identifier: BUSL-1.1
//

pragma solidity ^0.8.0;

import "./AssertionStakingPool.sol";
import "@openzeppelin/contracts/utils/Create2.sol";
import "@openzeppelin/contracts/utils/Address.sol";

/// @notice Creates staking pool contract for a target assertion. Can be used for any child Arbitrum chain running on top of the deployed AssertionStakingPoolCreator's chain.
contract AssertionStakingPoolCreator {
    event NewAssertionPoolCreated(
        address indexed rollup,
        bytes32 indexed _assertionHash,
        address assertionPool
    );

    error PoolDoesntExist(address rollup, AssertionInputs assertionInputs, bytes32 assertionHash);

    /// @notice Create a staking pool contract
    /// @param _rollup Rollup contract of target chain
    /// @param _assertionInputs Inputs to be passed into Rollup.stakeOnNewAssertion
    /// @param _assertionHash Assertion hash to be passed into Rollup.stakeOnNewAssertion
    function createPoolForAssertion(
        address _rollup,
        AssertionInputs memory _assertionInputs,
        bytes32 _assertionHash
    ) external returns (address) {
        address assertionPoolAddress = address(
            new AssertionStakingPool{salt: 0}(
                _rollup,
                _assertionInputs,
                _assertionHash
            )
        );

        emit NewAssertionPoolCreated(_rollup, _assertionHash, assertionPoolAddress);
        return assertionPoolAddress;
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
        bytes32 bytecodeHash = _getPoolByteCodeHash(_rollup, _assertionInputs, _assertionHash);

        address pool = Create2.computeAddress(0, bytecodeHash, address(this));
        if (Address.isContract(pool)) {
            return AssertionStakingPool(pool);
        } else {
            revert PoolDoesntExist(_rollup, _assertionInputs, _assertionHash);
        }
    }

    /// @notice get bytecodehash for create2 staking pool deployment
    function _getPoolByteCodeHash(
        address _rollup,
        AssertionInputs memory _assertionInputs,
        bytes32 _assertionHash
    ) internal pure returns (bytes32) {
        bytes memory bytecode = type(AssertionStakingPool).creationCode;
        return
            keccak256(
                abi.encodePacked(bytecode, abi.encode(_rollup, _assertionInputs, _assertionHash))
            );
    }
}
