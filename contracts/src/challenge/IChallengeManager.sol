// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE
// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import "./IChallengeResultReceiver.sol";

struct ChallengeVertex {
    uint256 stub;
}

interface IChallengeManager {
    function createChallenge(bytes32 startId) external returns (bytes32);

    function winningClaim(bytes32 challengeId) external view returns (bytes32);

    function vertexExists(bytes32 challengeId, bytes32 vId) external view returns (bool);

    function getVertex(bytes32 challengeId, bytes32 vId)
        external
        view
        returns (ChallengeVertex memory);

    function getCurrentPsTimer(bytes32 challengeId, bytes32 vId) external view returns (uint256);

    function confirmForPsTimer(bytes32 challengeId, bytes32 vId) external;

    function confirmForSucessionChallengeWin(bytes32 challengeId, bytes32 vId) external;

    function createSubChallenge(
        bytes32 challengeId,
        bytes32 child1Id,
        bytes32 child2Id
    ) external;

    function bisect(
        bytes32 challengeId,
        bytes32 vId,
        bytes32 prefixHistoryCommitment,
        bytes memory prefixProof
    ) external;

    function merge(
        bytes32 challengeId,
        bytes32 vId,
        bytes32 prefixHistoryCommitment,
        bytes memory prefixProof
    ) external;

    function addLeaf(
        bytes32 challengeId,
        bytes32 claimId,
        uint256 height,
        bytes32 historyCommitment,
        bytes32 lastState,
        bytes memory lastStatehistoryProof,
        bytes memory proof1,
        bytes memory proof2
    ) external;
}
