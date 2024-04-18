// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/OffchainLabs/nitro-contracts/blob/main/LICENSE
// SPDX-License-Identifier: BUSL-1.1
//

pragma solidity ^0.8.0;

import "./EdgeStakingPool.sol";
import "./StakingPoolCreatorUtils.sol";

/// @notice Creates EdgeStakingPool contracts.
contract EdgeStakingPoolCreator {
    event NewEdgeStakingPoolCreated(address indexed challengeManager, bytes32 indexed createEdgeArgsHash);

    /// @notice Create an edge staking pool contract
    /// @param challengeManager ChallengeManager contract
    /// @param createEdgeArgs Arguments to be passed into EdgeChallengeManager.createLayerZeroEdge
    function createPool(
        address challengeManager,
        CreateEdgeArgs memory createEdgeArgs
    ) external returns (EdgeStakingPool) {
        EdgeStakingPool pool = new EdgeStakingPool{salt: 0}(challengeManager, createEdgeArgs);
        emit NewEdgeStakingPoolCreated(challengeManager, keccak256(abi.encode(createEdgeArgs)));
        return pool;
    }

    function getPool(
        address challengeManager,
        CreateEdgeArgs memory createEdgeArgs
    ) public view returns (EdgeStakingPool) {
        return EdgeStakingPool(
            StakingPoolCreatorUtils.getPool(
                type(EdgeStakingPool).creationCode,
                abi.encode(challengeManager, createEdgeArgs)
            )
        );
    }
}
