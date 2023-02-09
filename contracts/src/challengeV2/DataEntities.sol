// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.17;

import "../osp/IOneStepProofEntry.sol";
import "./libraries/ChallengeVertexLib.sol";

enum Status {
    Pending,
    Confirmed,
    Rejected
}

interface IAssertionChain {
    function getPredecessorId(bytes32 assertionId) external view returns (bytes32);

    function getHeight(bytes32 assertionId) external view returns (uint256);

    function getInboxMsgCountSeen(bytes32 assertionId) external view returns (uint256);

    function getStateHash(bytes32 assertionId) external view returns (bytes32);

    function getSuccessionChallenge(bytes32 assertionId) external view returns (bytes32);

    function getFirstChildCreationTime(bytes32 assertionId) external view returns (uint256);

    function isFirstChild(bytes32 assertionId) external view returns (bool);
}

// CHRIS: TODO: move this to data entities?
struct OneStepData {
    ExecutionContext execCtx;
    uint256 machineStep;
    bytes32 beforeHash;
    bytes proof;
}

struct AddLeafArgs {
    bytes32 challengeId;
    bytes32 claimId;
    uint256 height;
    bytes32 historyRoot;
    bytes32 firstState;
    bytes firstStatehistoryProof;
    bytes32 lastState;
    bytes lastStatehistoryProof;
}

struct AddLeafLibArgs {
    uint256 miniStake;
    uint256 challengePeriod;
    AddLeafArgs leafData;
    bytes proof1;
    bytes proof2;
}

interface IChallengeManagerExternalView {
    function challengeExists(bytes32 challengeId) external view returns (bool);

    function getChallenge(bytes32 challengeId) external view returns (Challenge memory);

    function winningClaim(bytes32 challengeId) external view returns (bytes32);

    function vertexExists(bytes32 vId) external view returns (bool);

    function getVertex(bytes32 vId) external view returns (ChallengeVertex memory);

    function getCurrentPsTimer(bytes32 vId) external view returns (uint256);

    function hasConfirmedSibling(bytes32 vId) external view returns (bool);

    function isAtOneStepFork(bytes32 vId) external view returns (bool);
}

interface IChallengeManagerCore {
    function confirmForPsTimer(bytes32 vId) external;

    function confirmForSucessionChallengeWin(bytes32 vId) external;

    function createChallenge(bytes32 assertionId) external returns (bytes32);

    function createSubChallenge(bytes32 vId) external returns (bytes32);

    function bisect(bytes32 vId, bytes32 prefixHistoryRoot, bytes memory prefixProof) external returns (bytes32);

    function merge(bytes32 vId, bytes32 prefixHistoryRoot, bytes memory prefixProof) external returns (bytes32);

    function addLeaf(AddLeafArgs calldata leafData, bytes calldata proof1, bytes calldata proof2)
        external
        payable
        returns (bytes32);
}

interface IChallengeManager is IChallengeManagerCore, IChallengeManagerExternalView {}

enum ChallengeType {
    Block,
    BigStep,
    SmallStep,
    OneStep
}

struct Challenge {
    bytes32 rootId;
    // CHRIS: TODO: we could the leaf id here instead and just lookup the claim from the leaf
    bytes32 winningClaim;
    ChallengeType challengeType; // CHRIS: TODO: can use the keyword 'type' here?
}

// CHRIS: TODO: one step proof test just here for structure test
contract OneStepProofManager {
    mapping(bytes32 => bytes32) public winningClaims;

    function winningClaim(bytes32 challengeId) public view returns (bytes32) {
        return winningClaims[challengeId];
    }

    function createOneStepProof(bytes32 startState) public returns (bytes32) {
        revert("NOT_IMPLEMENTED");
    }

    function setWinningClaim(bytes32 startState, bytes32 _winner) public {
        winningClaims[startState] = _winner;
    }
}
