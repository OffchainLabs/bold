// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE
// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.0;

import "../libraries/DelegateCallAware.sol";
import "../osp/IOneStepProofEntry.sol";
import "../state/GlobalState.sol";
import "./IChallengeResultReceiver.sol";
import "./ChallengeLib.sol";
import "./NewChallengeLib.sol";
import "./IChallengeManager.sol";

import {NO_CHAL_INDEX} from "../libraries/Constants.sol";

contract NewChallengeManager is DelegateCallAware, IChallengeManager {
    using GlobalStateLib for GlobalState;
    using MachineLib for Machine;
    using NewChallengeLib for NewChallengeLib.Challenge;

    enum ChallengeModeRequirement {
        ANY,
        BLOCK,
        EXECUTION
    }

    string private constant NO_CHAL = "NO_CHAL";
    uint256 private constant MAX_CHALLENGE_DEGREE = 40;

    uint64 public totalChallengesCreated;
    mapping(uint256 => NewChallengeLib.Challenge) public challenges;
    mapping(uint256 => uint64) public totalVertexCreated;
    mapping(uint256 => mapping(uint256 => NewChallengeLib.Vertex)) public vertices;
    mapping(bytes32 => bool) public seenHistroyHash;

    IChallengeResultReceiver public resultReceiver;

    ISequencerInbox public sequencerInbox;
    IBridge public bridge;
    IOneStepProofEntry public osp;

    function challengeInfo(uint64 challengeIndex)
        external
        view
        override
        returns (NewChallengeLib.Challenge memory)
    {
        return challenges[challengeIndex];
    }

    function initialize(
        IChallengeResultReceiver resultReceiver_,
        ISequencerInbox sequencerInbox_,
        IBridge bridge_,
        IOneStepProofEntry osp_
    ) external override onlyDelegated {
        require(address(resultReceiver) == address(0), "ALREADY_INIT");
        require(address(resultReceiver_) != address(0), "NO_RESULT_RECEIVER");
        resultReceiver = resultReceiver_;
        sequencerInbox = sequencerInbox_;
        bridge = bridge_;
        osp = osp_;
    }

    function createChallenge(
        MachineStatus startMachineStatuses_,
        GlobalState calldata startGlobalStates_,
        bytes32 wasmModuleRoot_,
        uint256 confirmPeriodBlocks
    ) external override returns (uint64) {
        require(msg.sender == address(resultReceiver), "ONLY_ROLLUP_CHAL");
        uint64 challengeIndex = ++totalChallengesCreated;
        // The following is an assertion since it should never be possible, but it's an important invariant
        require(challengeIndex != NO_CHAL_INDEX, "challengeIndex overflow");
        NewChallengeLib.Challenge storage challenge = challenges[challengeIndex];
        challenge.wasmModuleRoot = wasmModuleRoot_;
        challenge.confirmPeriodBlocks = confirmPeriodBlocks;

        ++totalVertexCreated[challengeIndex]; // skip 0 index as nil

        // emit InitiatedChallenge( // TODO: Fix event

        return challengeIndex;
    }

    function addChallengeVertex(
        uint64 challengeIndex,
        uint64 assertionNum,
        NewChallengeLib.HistoryCommitment calldata history
    ) external payable returns (uint64) {
        // TODO: mini stake
        // TODO: verify commitment
        // TODO: validator whitelist?
        require(challenges[challengeIndex].wasmModuleRoot != bytes32(0), "CHAL_NOT_FOUND");

        bytes32 historyHash = NewChallengeLib.historyHash(history);
        require(!seenHistroyHash[historyHash], "HISTORY_SEEN");
        seenHistroyHash[historyHash] = true;

        uint64 vertexIndex = ++totalVertexCreated[challengeIndex];
        require(vertexIndex != NO_CHAL_INDEX, "vertexIndex overflow");

        NewChallengeLib.Vertex storage vertex = vertices[challengeIndex][vertexIndex];
        vertex.validator = msg.sender;
        vertex.isLeaf = true;
        vertex.status = NewChallengeLib.VertexStatus.Pending;
        vertex.history = history.merkleRoot;
        vertex.height = history.height;
        vertex.prev = 1; // root
        vertex.presumptivSuccessor = 0; // none
        vertex.winnerIfConfirmed = assertionNum;
        updatePresumptivSuccessor(challengeIndex, vertex.prev, vertexIndex);
        return vertexIndex;
    }

    function confirmForPSTimer(uint64 challengeIndex, uint64 vertexIndex) external {
        // TODO: other confirm rules
        NewChallengeLib.Vertex storage vertex = vertices[challengeIndex][vertexIndex];
        require(vertex.validator != address(0), "VERTEX_NOT_FOUND");
        require(vertex.status == NewChallengeLib.VertexStatus.Pending, "NOT_PENDING_VERTEX");
        updatePresumptivSuccessor(challengeIndex, vertex.prev, 0);
        require(vertex.psTimer >= challenges[challengeIndex].confirmPeriodBlocks, "PSTIMER_LOW");
        vertex.status = NewChallengeLib.VertexStatus.Confirmed;
        if(vertex.winnerIfConfirmed > 0){
            resultReceiver.completeChallenge(challengeIndex, vertex.winnerIfConfirmed);
        }
    }

    function updatePresumptivSuccessor(
        uint64 challengeIndex,
        uint64 vertexIndex,
        uint64 potentialVertex
    ) internal {
        NewChallengeLib.Vertex storage v0 = vertices[challengeIndex][vertexIndex];
        if (v0.presumptivSuccessor == 0) {
            v0.presumptivSuccessor = potentialVertex;
        } else {
            NewChallengeLib.Vertex storage currentPs = vertices[challengeIndex][
                v0.presumptivSuccessor
            ];
            currentPs.psTimer += uint64(block.number - v0.lastPsUpdate);
            NewChallengeLib.Vertex storage v1 = vertices[challengeIndex][potentialVertex];
            // do not change ps if it can be confirmed already
            if (
                currentPs.psTimer < challenges[challengeIndex].confirmPeriodBlocks &&
                v1.height < currentPs.height &&
                potentialVertex != 0
            ) {
                v0.presumptivSuccessor = potentialVertex;
            }
        }
        v0.lastPsUpdate = uint64(block.number);
    }

    function bisect(
        uint64 challengeIndex, 
        uint64 vertexIndex,
        NewChallengeLib.HistoryCommitment calldata history,
        bytes32[] calldata proof
    ) external returns (uint256) {
        bytes32 historyHash = NewChallengeLib.historyHash(history);
        require(!seenHistroyHash[historyHash], "HISTORY_SEEN");
        seenHistroyHash[historyHash] = true;

        NewChallengeLib.Vertex storage vertex = vertices[challengeIndex][vertexIndex];
        require(vertex.validator != address(0), "VERTEX_NOT_FOUND");
        require(vertex.status == NewChallengeLib.VertexStatus.Pending, "NOT_PENDING_VERTEX");
        require(vertex.prev > 0, "PREV_IS_NIL");
        NewChallengeLib.Vertex storage prev = vertices[challengeIndex][vertex.prev];
        require(prev.presumptivSuccessor != vertexIndex, "ALREADY_IS_PS");
        NewChallengeLib.Vertex storage prevps = vertices[challengeIndex][prev.presumptivSuccessor];
        require(prevps.psTimer < challenges[challengeIndex].confirmPeriodBlocks, "PREV_PS_CONFIRMED");
        require(prev.height <= vertex.height - 2, "SHOULD_OSP");
        require(NewChallengeLib.bisectHeight(vertex.height) == history.height, "BAD_BISECT_HIGHT");
        require(NewChallengeLib.verifyPrefixProof({
            prefix: history.merkleRoot,
            root: vertex.history,
            proof: proof
        }), "BAD_PREFIX_PROOF");
        
        // update psTimer
        updatePresumptivSuccessor(challengeIndex, vertex.prev, 0);

        uint64 bisectIndex = ++totalVertexCreated[challengeIndex];
        NewChallengeLib.Vertex storage bisected = vertices[challengeIndex][bisectIndex]; // aka N
        bisected.validator = msg.sender;
        bisected.isLeaf = false;
        bisected.status = NewChallengeLib.VertexStatus.Pending;
        bisected.history = history.merkleRoot;
        bisected.height = history.height;
        bisected.prev = vertex.prev; // root
        bisected.presumptivSuccessor = vertexIndex; // none
        bisected.lastPsUpdate = vertex.psTimer;
        bisected.winnerIfConfirmed = 0; // TODO: this is only for leaf right?
        vertex.prev = bisectIndex;
        updatePresumptivSuccessor(challengeIndex, bisected.prev, bisectIndex);
        // TODO: spec said V's chess clock is stopped, but isn't it the PS of N?
        return bisectIndex;
    }

    function merge(uint64 challengeIndex, uint64 vertexFromIndex, uint64 vertexToIndex, bytes32[] calldata proof) external returns (uint256) {

        NewChallengeLib.Vertex storage vertexFrom = vertices[challengeIndex][vertexFromIndex];
        require(vertexFrom.validator != address(0), "VERTEX_NOT_FOUND");
        require(vertexFrom.status == NewChallengeLib.VertexStatus.Pending, "NOT_PENDING_VERTEX");
        NewChallengeLib.Vertex storage vertexTo = vertices[challengeIndex][vertexToIndex];
        require(vertexTo.validator != address(0), "VERTEX_NOT_FOUND");


        require(vertexFrom.prev > 0, "PREV_IS_NIL");
        NewChallengeLib.Vertex storage prev = vertices[challengeIndex][vertexFrom.prev];
        require(prev.height <= vertexFrom.height - 2, "SHOULD_OSP");
        require(NewChallengeLib.bisectHeight(vertexFrom.height) == vertexTo.height, "BAD_MERGE_HEIGHT");
        require(NewChallengeLib.verifyPrefixProof({
            prefix: vertexTo.history,
            root: vertexFrom.history,
            proof: proof
        }), "BAD_PREFIX_PROOF");

        // require(prev.presumptivSuccessor != vertexIndex, "ALREADY_IS_PS");
        // NewChallengeLib.Vertex storage prevps = vertices[challengeIndex][prev.presumptivSuccessor];
        // require(prevps.psTimer < challenges[challengeIndex].confirmPeriodBlocks, "PREV_PS_CONFIRMED");
        
        // update psTimer
        updatePresumptivSuccessor(challengeIndex, vertexFrom.prev, 0);

        vertexFrom.prev = vertexToIndex;
        vertexTo.psTimer += vertexFrom.psTimer;
        updatePresumptivSuccessor(challengeIndex, vertexFrom.prev, vertexFromIndex);
        // TODO: spec said V's chess clock is stopped, but isn't it the PS of N?
        return vertexToIndex;
    }

    function getChallengeVertex(uint64 challengeIndex, uint64 vertexIndex)
        external
        view
        returns (NewChallengeLib.Vertex memory)
    {
        return vertices[challengeIndex][vertexIndex];
    }

    function timeout(uint64 challengeIndex) external override {
        revert("DEPRECATED");
    }

    function clearChallenge(uint64 challengeIndex) external override {
        revert("DEPRECATED");
    }

    function currentResponder(uint64 challengeIndex) public view override returns (address) {
        revert("DEPRECATED");
    }

    function isTimedOut(uint64 challengeIndex) public view override returns (bool) {
        revert("DEPRECATED");
    }
}
