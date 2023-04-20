// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.17;

import "./libraries/UintUtilsLib.sol";
import "./DataEntities.sol";
import "./libraries/EdgeChallengeManagerLib.sol";
import "../libraries/Constants.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

interface IEdgeChallengeManager {
     // Creates a layer zero edge in a challenge.
    function createLayerZeroEdge(CreateEdgeArgs memory args, bytes calldata, bytes calldata)
        external
        payable
        returns (bytes32);

    // Bisects an edge. Emits both children's edge IDs in an event.
    function bisectEdge(bytes32 eId, bytes32 prefixHistoryRoot, bytes memory prefixProof)
        external
        returns (bytes32, bytes32);

    // Checks if both children of an edge are already confirmed in order to confirm the edge.
    function confirmEdgeByChildren(bytes32 eId) external;

    // Confirms an edge by edge ID and an array of ancestor edges based on total time unrivaled
    function confirmEdgeByTime(bytes32 eId, bytes32[] memory ancestorIds) external;

    // If we have created a subchallenge, confirmed a layer 0 edge already, we can use a claim id to confirm edge ids.
    // All edges have two children, unless they only have a link to a claim id.
    function confirmEdgeByClaim(bytes32 eId, bytes32 claimId) external;

    // when we reach a one step fork in a small step challenge we can confirm
    // the edge by executing a one step proof to show the edge is valid
    function confirmEdgeByOneStepProof(
        bytes32 edgeId,
        OneStepData calldata oneStepData,
        bytes32[] calldata beforeHistoryInclusionProof,
        bytes32[] calldata afterHistoryInclusionProof
    ) external;

    function calculateEdgeId(
        EdgeType edgeType,
        bytes32 originId,
        uint256 startHeight,
        bytes32 startHistoryRoot,
        uint256 endHeight,
        bytes32 endHistoryRoot
    ) external pure returns (bytes32);

    function calculateMutualId(
        EdgeType edgeType,
        bytes32 originId,
        uint256 startHeight,
        bytes32 startHistoryRoot,
        uint256 endHeight
    ) external pure returns (bytes32);

    function edgeExists(bytes32 edgeId) external view returns (bool);

    function getEdge(bytes32 edgeId) external view returns (ChallengeEdge memory);

    function edgeLength(bytes32 edgeId) external view returns (uint256);

    function hasRival(bytes32 edgeId) external view returns (bool);

    function hasLengthOneRival(bytes32 edgeId) external view returns (bool);

    function timeUnrivaled(bytes32 edgeId) external view returns (uint256);

    function getPrevAssertionId(bytes32 edgeId) external view returns (bytes32);

    function firstRival(bytes32 edgeId) external view returns (bytes32);
}

// // CHRIS: TODO: check the ministake was provided

// // CHRIS: TODO: invariants
// // 1. edges are only created, never destroyed
// // 2. all edges have at least one parent, or a claim id - other property invariants exist
// // 3. all edges have a mutual id, and that mutual id must have an entry in firstRivals
// // 4. all values of firstRivals are existing edges (must be in the edge mapping), or are the NO_RIVAL magic hash
// // 5. where to check edge prefix proofs? in bisection, or in add?

contract EdgeChallengeManager is IEdgeChallengeManager, Initializable {
    using EdgeChallengeManagerLib for EdgeStore;
    using ChallengeEdgeLib for ChallengeEdge;

    EdgeStore internal store;

    uint256 public challengePeriodSec;
    IAssertionChain public assertionChain;
    IOneStepProofEntry public oneStepProofEntry;

    constructor() {
        _disableInitializers();
    }

    function initialize(
        IAssertionChain _assertionChain,
        uint256 _challengePeriodSec,
        IOneStepProofEntry _oneStepProofEntry
    ) public initializer {
        require(address(assertionChain) == address(0), "ALREADY_INIT");
        assertionChain = _assertionChain;
        challengePeriodSec = _challengePeriodSec;
        oneStepProofEntry = _oneStepProofEntry;
    }

    /////////////////////////////
    // STATE MUTATING SECTIION //
    /////////////////////////////

    function createLayerZeroEdge(CreateEdgeArgs memory args, bytes calldata prefixProof, bytes calldata proof)
        external
        payable
        returns (bytes32)
    {
        return store.createLayerZeroEdge(assertionChain, args, prefixProof, proof);
    }

    function bisectEdge(bytes32 edgeId, bytes32 bisectionHistoryRoot, bytes memory prefixProof)
        external
        returns (bytes32, bytes32)
    {
        return store.bisectEdge(edgeId, bisectionHistoryRoot, prefixProof);
    }

    function confirmEdgeByChildren(bytes32 edgeId) public {
        store.confirmEdgeByChildren(edgeId);
    }

    function confirmEdgeByClaim(bytes32 edgeId, bytes32 claimingEdgeId) public {
        store.confirmEdgeByClaim(edgeId, claimingEdgeId);
    }

    function confirmEdgeByTime(bytes32 edgeId, bytes32[] memory ancestorEdges) public {
        store.confirmEdgeByTime(edgeId, ancestorEdges, challengePeriodSec);
    }

    function confirmEdgeByOneStepProof(
        bytes32 edgeId,
        OneStepData calldata oneStepData,
        bytes32[] calldata beforeHistoryInclusionProof,
        bytes32[] calldata afterHistoryInclusionProof
    ) public {
        bytes32 prevAssertionId = store.getPrevAssertionId(edgeId);
        ExecutionContext memory execCtx = ExecutionContext({
            maxInboxMessagesRead: assertionChain.getInboxMsgCountSeen(prevAssertionId),
            bridge: assertionChain.bridge(),
            initialWasmModuleRoot: assertionChain.getWasmModuleRoot(prevAssertionId)
        });

        store.confirmEdgeByOneStepProof(
            edgeId, oneStepProofEntry, oneStepData, execCtx, beforeHistoryInclusionProof, afterHistoryInclusionProof
        );
    }

    ///////////////////////
    // VIEW ONLY SECTION //
    ///////////////////////

    function calculateEdgeId(
        EdgeType edgeType,
        bytes32 originId,
        uint256 startHeight,
        bytes32 startHistoryRoot,
        uint256 endHeight,
        bytes32 endHistoryRoot
    ) public pure returns (bytes32) {
        return
            ChallengeEdgeLib.idComponent(edgeType, originId, startHeight, startHistoryRoot, endHeight, endHistoryRoot);
    }

    function calculateMutualId(
        EdgeType edgeType,
        bytes32 originId,
        uint256 startHeight,
        bytes32 startHistoryRoot,
        uint256 endHeight
    ) public pure returns (bytes32) {
        return ChallengeEdgeLib.mutualIdComponent(edgeType, originId, startHeight, startHistoryRoot, endHeight);
    }

    function edgeExists(bytes32 edgeId) public view returns (bool) {
        return store.edges[edgeId].exists();
    }

    function getEdge(bytes32 edgeId) public view returns (ChallengeEdge memory) {
        return store.get(edgeId);
    }

    function edgeLength(bytes32 edgeId) public view returns (uint256) {
        return store.get(edgeId).length();
    }

    function hasRival(bytes32 edgeId) public view returns (bool) {
        return store.hasRival(edgeId);
    }

    function hasLengthOneRival(bytes32 edgeId) public view returns (bool) {
        return store.hasLengthOneRival(edgeId);
    }

    function timeUnrivaled(bytes32 edgeId) public view returns (uint256) {
        return store.timeUnrivaled(edgeId);
    }

    function getPrevAssertionId(bytes32 edgeId) public view returns (bytes32) {
        return store.getPrevAssertionId(edgeId);
    }

    function firstRival(bytes32 edgeId) public view returns (bytes32) {
        return store.firstRivals[edgeId];
    }
}
