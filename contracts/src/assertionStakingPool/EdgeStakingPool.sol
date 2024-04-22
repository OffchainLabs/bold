// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/OffchainLabs/nitro-contracts/blob/main/LICENSE
// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.0;

import "./AbsBoldStakingPool.sol";
import "../challengeV2/EdgeChallengeManager.sol";

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/// @notice Trustless staking pool contract for creating layer zero edges.
///
///         Allows users to deposit stake, create an edge once required stake amount is reached,
///         and reclaim their stake when and if the edge is confirmed.
///
///         Honest participants should check that both edgeId and edgeLevel are correct before staking. 
///         If either is not expected, the pool should be ignored.
///
/// @dev    Unlike the assertion staking pool, there is no need for a function to claim the stake back into the pool.
contract EdgeStakingPool is AbsBoldStakingPool {
    using SafeERC20 for IERC20;

    EdgeChallengeManager public immutable challengeManager;
    bytes32 public immutable edgeId;
    uint8 public immutable edgeLevel;
    uint256 public immutable requiredStake;

    /// @notice The provided arguments to not match createEdgeArgsHash
    error IncorrectEdgeId(bytes32 actual, bytes32 expected);

    /// @param _challengeManager EdgeChallengeManager contract
    /// @param _edgeId The ID of the edge to be created (see ChallengeEdgeLib.id)
    /// @param _edgeLevel The level of the edge to be created
    constructor(
        address _challengeManager,
        bytes32 _edgeId,
        uint8 _edgeLevel
    ) AbsBoldStakingPool(EdgeChallengeManager(_challengeManager).stakeToken()) {
        challengeManager = EdgeChallengeManager(_challengeManager);
        edgeId = _edgeId;
        edgeLevel = _edgeLevel;
        requiredStake = challengeManager.stakeAmounts(_edgeLevel);
    }

    /// @notice Create the edge. Callable only if required stake has been reached and edge has not been created yet.
    function createEdge(CreateEdgeArgs calldata args) external {
        stakeToken.safeIncreaseAllowance(address(challengeManager), requiredStake);
        bytes32 newEdgeId = challengeManager.createLayerZeroEdge(args);
        if (newEdgeId != edgeId) {
            revert IncorrectEdgeId(newEdgeId, edgeId);
        }
    }
}
