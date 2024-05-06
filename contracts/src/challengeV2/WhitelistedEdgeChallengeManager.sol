// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE
// SPDX-License-Identifier: BUSL-1.1
//
pragma solidity ^0.8.17;

import "./EdgeChallengeManager.sol";
import "../rollup/IRollupCore.sol";

interface IWhitelistedEdgeChallengeManager is IEdgeChallengeManager {
    function accountHasMadeRival(bytes32 mutualId, address account) external view returns (bool);
}

contract WhitelistedEdgeChallengeManager is EdgeChallengeManager, IWhitelistedEdgeChallengeManager {
    using ChallengeEdgeLib for ChallengeEdge;

    mapping(bytes32 => mapping(address => bool)) public accountHasMadeRival;

    error AccountHasMadeRival(address account);
    error NotValidator(address account);

    function createLayerZeroEdge(CreateEdgeArgs calldata args) public override(EdgeChallengeManager, IEdgeChallengeManager) returns (bytes32) {
        if (!isValidator(msg.sender)) {
            revert NotValidator(msg.sender);
        }

        bytes32 edgeId = super.createLayerZeroEdge(args);
        bytes32 mutualId = store.edges[edgeId].mutualId();

        if (accountHasMadeRival[mutualId][msg.sender]) {
            revert AccountHasMadeRival(msg.sender);
        }

        accountHasMadeRival[mutualId][msg.sender] = true;

        return edgeId;
    }

    function isValidator(address account) public view returns (bool) {
        IRollupCore rollup = IRollupCore(address(assertionChain));
        return rollup.validatorWhitelistDisabled() || rollup.isValidator(account);
    }
}