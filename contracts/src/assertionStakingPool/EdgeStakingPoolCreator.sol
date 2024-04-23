// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/OffchainLabs/nitro-contracts/blob/main/LICENSE
// SPDX-License-Identifier: BUSL-1.1
//

pragma solidity ^0.8.0;

import "./EdgeStakingPool.sol";
import "./StakingPoolCreatorUtils.sol";

/// @notice Creates EdgeStakingPool contracts.
contract EdgeStakingPoolCreator {
    event NewEdgeStakingPoolCreated(address indexed challengeManager, bytes32 indexed edgeId);

    /// @notice Create an edge staking pool contract
    /// @param challengeManager EdgeChallengeManager contract
    /// @param edgeId The ID of the edge to be created (see ChallengeEdgeLib.id)
    function createPool(
        address challengeManager,
        bytes32 edgeId
    ) external returns (EdgeStakingPool) {
        EdgeStakingPool pool = new EdgeStakingPool{salt: 0}(challengeManager, edgeId);
        emit NewEdgeStakingPoolCreated(challengeManager, edgeId);
        return pool;
    }

    /// @notice get staking pool deployed with provided inputs; reverts if pool contract doesn't exist.
    /// @param challengeManager EdgeChallengeManager contract
    /// @param edgeId The ID of the edge to be created (see ChallengeEdgeLib.id)
    function getPool(
        address challengeManager,
        bytes32 edgeId
    ) public view returns (EdgeStakingPool) {
        return EdgeStakingPool(
            StakingPoolCreatorUtils.getPool(
                type(EdgeStakingPool).creationCode,
                abi.encode(challengeManager, edgeId)
            )
        );
    }
}
