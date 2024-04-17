// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/OffchainLabs/nitro-contracts/blob/main/LICENSE
// SPDX-License-Identifier: BUSL-1.1
//

pragma solidity ^0.8.0;

import "./EdgeStakingPool.sol";
import "@openzeppelin/contracts/utils/Create2.sol";
import "@openzeppelin/contracts/utils/Address.sol";

/// @notice Creates staking pool contract for a target assertion. Can be used for any child Arbitrum chain running on top of the deployed AssertionStakingPoolCreator's chain.
contract EdgeStakingPoolCreator {
    event NewEdgeStakingPoolCreated(address indexed challengeManager, bytes32 indexed createEdgeArgsHash);

    error PoolDoesntExist();

    function createPool(
        address challengeManager,
        CreateEdgeArgs memory createEdgeArgs
    ) external returns (address) {
        address pool = address(
            new EdgeStakingPool{salt: 0}(challengeManager, createEdgeArgs)
        );

        emit NewEdgeStakingPoolCreated(challengeManager, keccak256(abi.encode(createEdgeArgs)));
        return pool;
    }

    function getPool(
        address challengeManager,
        CreateEdgeArgs memory createEdgeArgs
    ) public view returns (EdgeStakingPool) {
        bytes32 bytecodeHash = _getPoolByteCodeHash(challengeManager, createEdgeArgs);

        address pool = Create2.computeAddress(0, bytecodeHash, address(this));
        if (Address.isContract(pool)) {
            return EdgeStakingPool(pool);
        } else {
            revert PoolDoesntExist();
        }
    }

    /// @notice get bytecodehash for create2 staking pool deployment
    function _getPoolByteCodeHash(
        address challengeManager,
        CreateEdgeArgs memory createEdgeArgs
    ) internal pure returns (bytes32) {
        bytes memory bytecode = type(EdgeStakingPool).creationCode;
        return
            keccak256(
                abi.encodePacked(bytecode, abi.encode(challengeManager, createEdgeArgs))
            );
    }
}
