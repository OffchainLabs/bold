// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/OffchainLabs/nitro-contracts/blob/main/LICENSE
// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.0;

import "./AbsBoldStakingPool.sol";
import "../challengeV2/EdgeChallengeManager.sol";

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/// @notice Staking pool contract for creating layer zero edges.
/// Allows users to deposit stake, create an edge once required stake amount is reached,
/// and reclaim their stake when and if the edge is confirmed.
/// @dev Unlike the assertion staking pool, there is no need for a function to claim the stake back into the pool.
contract EdgeStakingPool is AbsBoldStakingPool {
    using SafeERC20 for IERC20;

    /// @notice The targeted challenge manager contract
    EdgeChallengeManager public immutable challengeManager;
    /// @notice todo
    bytes32 public immutable createEdgeArgsHash;
    /// @notice The required stake amount to create the edge
    uint256 public immutable requiredStake;

    /// @notice The provided arguments to not match createEdgeArgsHash
    error IncorrectCreateEdgeArgs();

    /// @param _challengeManager EdgeChallengeManager contract
    /// @param createEdgeArgs Arguments to be passed into EdgeChallengeManager.createLayerZeroEdge
    constructor(
        address _challengeManager,
        CreateEdgeArgs memory createEdgeArgs
    ) AbsBoldStakingPool(EdgeChallengeManager(_challengeManager).stakeToken()) {
        challengeManager = EdgeChallengeManager(_challengeManager);
        createEdgeArgsHash = keccak256(abi.encode(createEdgeArgs));
        requiredStake = challengeManager.stakeAmounts(createEdgeArgs.level);
    }

    /// @notice Create the edge. Callable only if required stake has been reached and edge has not been created yet.
    function createEdge(CreateEdgeArgs calldata args) external {
        if (!isCorrectCreateEdgeArgs(args)) {
            revert IncorrectCreateEdgeArgs();
        }
        stakeToken.safeIncreaseAllowance(address(challengeManager), requiredStake);
        challengeManager.createLayerZeroEdge(args);
    }

    /// @notice Check that the provided arguments match the expected arguments
    function isCorrectCreateEdgeArgs(CreateEdgeArgs calldata args) public view returns (bool) {
        return keccak256(abi.encode(args)) == createEdgeArgsHash;
    }
}
