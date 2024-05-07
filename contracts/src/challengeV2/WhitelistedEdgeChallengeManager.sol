// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE
// SPDX-License-Identifier: BUSL-1.1
//
pragma solidity ^0.8.17;

import "./EdgeChallengeManager.sol";
import "../rollup/IRollupCore.sol";

interface IWhitelistedEdgeChallengeManager is IEdgeChallengeManager {
    /// @notice Thrown when an account has already created a rivalling edge
    error AccountHasMadeRival(address account);
    /// @notice Thrown when an account is not a validator
    error NotValidator(address account);

    /// @notice Checks if an account has created an edge with the given mutualId
    function accountHasMadeRival(bytes32 mutualId, address account) external view returns (bool);
}

/// @notice Extension of EdgeChallengeManager that restricts the creation of layer zero edges to the rollup's whitelisted validators.
///         If the rollup's whitelist is enabled, a validator cannot create a layer zero edge that rivals an edge it previously created.
contract WhitelistedEdgeChallengeManager is EdgeChallengeManager, IWhitelistedEdgeChallengeManager {
    using ChallengeEdgeLib for ChallengeEdge;

    /// @inheritdoc IWhitelistedEdgeChallengeManager
    mapping(bytes32 => mapping(address => bool)) public accountHasMadeRival;

    /// @inheritdoc IEdgeChallengeManager
    function createLayerZeroEdge(CreateEdgeArgs calldata args) public override(EdgeChallengeManager, IEdgeChallengeManager) returns (bytes32) {
        IRollupCore rollup = IRollupCore(address(assertionChain));

        if (rollup.validatorWhitelistDisabled()) {
            return EdgeChallengeManager.createLayerZeroEdge(args);
        }

        if (!rollup.isValidator(msg.sender)) {
            revert NotValidator(msg.sender);
        }

        bytes32 edgeId = EdgeChallengeManager.createLayerZeroEdge(args);
        bytes32 mutualId = store.edges[edgeId].mutualId();

        if (accountHasMadeRival[mutualId][msg.sender]) {
            revert AccountHasMadeRival(msg.sender);
        }

        accountHasMadeRival[mutualId][msg.sender] = true;

        return edgeId;
    }
}