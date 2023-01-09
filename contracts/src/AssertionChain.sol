// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

interface IAssertionChain {
    struct Assertion {
        uint256 seqNum;
        StateCommitment stateCommitment;
        uint status;
        bool isFirstChild;
        uint256 firstChildCreationTimestamp;
        uint256 secondChildCreationTimestamp;
        address validator;
    }
    struct Challenge {
        uint256 seqNum;
        uint256 nextSeqNum;
        ChallengeVertex root;
        ChallengeVertex latestConfirmed;
        uint256 creationTimestamp;
        address actor;
    }
    struct ChallengeVertex {
        uint256 seqNum;
        bytes32 challengeParentStateCommitHash;
        bool actor;
        bool isLeaf;
        uint256 psTimer;
    }
    struct StateCommitment {
        uint256 height;
        bytes32 stateRoot;
    }
    struct HistoryCommitment {
        uint256 height;
        bytes32 merkleRoot;
    }

    // Read-only calls.
    function numAssertions() external view returns (uint256);
    function challengePeriodSeconds() external view returns (uint256);
    function latestConfirmedAssertion() external view returns (Assertion memory assertion);
    function getAssertion(uint256 seqNum) external view returns (Assertion memory assertion);
    function getChallenge(bytes32 parentStateCommitHash) external view returns (Challenge memory challenge);
    function getChallengeVertex(bytes32 challengeCommitHash, bytes32 vertexHash) external view returns (ChallengeVertex memory vertex);

    // Mutating calls.
    function createAssertion(
        Assertion calldata prev,
        StateCommitment calldata commit
    ) external payable returns (Assertion memory assertion);
    function confirmForWin(Assertion calldata assertion) external payable;
    function confirmNoRival(Assertion calldata assertion) external payable;
    function rejectForLoss(Assertion calldata assertion) external payable;
    function rejectForPrev(Assertion calldata assertion) external payable;
    function confirmForPSTimer(ChallengeVertex calldata vertex) external payable;
    function confirmForChallengeDeadline(ChallengeVertex calldata vertex) external payable;
    function confirmForSubchallengeWin(ChallengeVertex calldata vertex) external payable;
    function createChallenge(Assertion calldata prev) external payable returns (Challenge memory challenge);
}

contract AssertionChain is IAssertionChain {
    // Storage variables.
    Assertion[] assertions;
    mapping (bytes32 => bytes32) vertexHashByChallengeHash;
    mapping (bytes32 => ChallengeVertex) verticesByHash;
    mapping (bytes32 => Challenge) challengesByCommitHash;
    uint256 public challengePeriodSeconds;

    constructor(uint256 _challengePeriodSeconds) {
        challengePeriodSeconds = _challengePeriodSeconds;
        Assertion memory genesis = Assertion({
            seqNum: 0,
            stateCommitment: StateCommitment({
                height: 0,
                stateRoot: bytes32(0)
            }),
            status: 0,
            isFirstChild: false,
            firstChildCreationTimestamp: 0,
            secondChildCreationTimestamp: 0,
            validator: address(0)
        });
        assertions.push(genesis);
    }

    // Read-only calls.
    function numAssertions() external view returns (uint256) {
        return assertions.length;
    }

    function latestConfirmedAssertion() external view returns (Assertion memory assertion) {
        return assertions[assertions.length-1];
    }

    function getAssertion(uint256 seqNum) external view returns (Assertion memory assertion) {
        require(seqNum < assertions.length);
        return assertions[seqNum];
    }

    function getChallenge(bytes32 challengeCommitHash) external view returns (Challenge memory challenge) {
        return challengesByCommitHash[challengeCommitHash];
    }

    function getChallengeVertex(bytes32 challengeCommitHash, bytes32 vertexHash) external view returns (ChallengeVertex memory vertex) {
        require(vertexHashByChallengeHash[challengeCommitHash] == vertexHash);
        return verticesByHash[vertexHash];
    }

    // Mutating calls.
    function createAssertion(
        Assertion calldata prev,
        StateCommitment calldata commit
    ) external payable returns (Assertion memory assertion) {
        revert("unimplemented");
    }
    function confirmForWin(Assertion calldata assertion) external payable {
        revert("unimplemented");
    }
    function confirmNoRival(Assertion calldata assertion) external payable {
        revert("unimplemented");
    }
    function rejectForLoss(Assertion calldata assertion) external payable {
        revert("unimplemented");
    }
    function rejectForPrev(Assertion calldata assertion) external payable {
        revert("unimplemented");
    }
    function confirmForPSTimer(ChallengeVertex calldata vertex) external payable {
        revert("unimplemented");
    }
    function confirmForChallengeDeadline(ChallengeVertex calldata vertex) external payable {
        revert("unimplemented");
    }
    function confirmForSubchallengeWin(ChallengeVertex calldata vertex) external payable {
        revert("unimplemented");
    }
    function createChallenge(Assertion calldata prev) external payable returns (Challenge memory challenge) {
        revert("unimplemented");
    }
}
