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
/// @dev    Unlike the assertion staking pool, there is no need for a function to claim the stake back into the pool.
contract EdgeStakingPool is AbsBoldStakingPool {
    using SafeERC20 for IERC20;

    /// @notice The targeted challenge manager contract
    EdgeChallengeManager public immutable challengeManager;
    /// @notice The ID of the edge to be created (see ChallengeEdgeLib.id)
    bytes32 public immutable edgeId;

    /// @notice The provided arguments to not match createEdgeArgsHash
    error IncorrectEdgeId(bytes32 actual, bytes32 expected);

    /// @param _challengeManager EdgeChallengeManager contract
    /// @param _edgeId The ID of the edge to be created (see ChallengeEdgeLib.id)
    constructor(
        address _challengeManager,
        bytes32 _edgeId
    ) AbsBoldStakingPool(EdgeChallengeManager(_challengeManager).stakeToken()) {
        challengeManager = EdgeChallengeManager(_challengeManager);
        edgeId = _edgeId;
    }

    /// @notice Create the edge. Callable only if required stake has been reached and edge has not been created yet.
    function createEdge(CreateEdgeArgs calldata args) external {
        uint256 requiredStake = challengeManager.stakeAmounts(args.level);
        stakeToken.safeIncreaseAllowance(address(challengeManager), requiredStake);
        bytes32 newEdgeId = challengeManager.createLayerZeroEdge(args);
        if (newEdgeId != edgeId) {
            revert IncorrectEdgeId(newEdgeId, edgeId);
        }
    }
}
