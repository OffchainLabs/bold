// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.17;

import "./libraries/UintUtilsLib.sol";
import "./DataEntities.sol";
import "./libraries/EdgeChallengeManagerLib.sol";
import "../libraries/Constants.sol";

interface IEdgeChallengeManager {
    // Checks if an edge by ID exists.
    function edgeExists(bytes32 eId) external view returns (bool);

    function initialize(
        IAssertionChain _assertionChain,
        uint256 _challengePeriodSec,
        IOneStepProofEntry _oneStepProofEntry
    ) external;

    // // Checks if an edge by ID exists.
    // function edgeExists(bytes32 eId) external view returns (bool);
    // Gets an edge by ID.
    function getEdge(bytes32 eId) external view returns (ChallengeEdge memory);

    // Gets the current time unrivaled by edge ID. TODO: Needs more thinking.
    function timeUnrivaled(bytes32 eId) external view returns (uint256);

    // We define a mutual ID as hash(EdgeType  ++ originId ++ hash(startCommit ++ startHeight)) as a way
    // of checking if an edge has rivals. Rivals edges share the same mutual ID.
    function calculateMutualId(
        EdgeType edgeType,
        bytes32 originId,
        uint256 startHeight,
        bytes32 startHistoryRoot,
        uint256 endHeight
    ) external returns (bytes32);

    function calculateEdgeId(
        EdgeType edgeType,
        bytes32 originId,
        uint256 startHeight,
        bytes32 startHistoryRoot,
        uint256 endHeight,
        bytes32 endHistoryRoot
    ) external returns (bytes32);

    // Checks if an edge's mutual ID corresponds to multiple rivals and checks if a one step fork exists.
    function hasRival(bytes32 eId) external view returns (bool);

    // Checks if an edge's mutual ID corresponds to multiple rivals and checks if a one step fork exists.
    function hasLengthOneRival(bytes32 eId) external view returns (bool);

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
}

// /// @notice Data for creating a layer zero edge
// struct CreateEdgeArgs {
//     /// @notice The type of edge to be created
//     EdgeType edgeType;
//     /// @notice The end history root of the edge to be created
//     bytes32 endHistoryRoot;
//     /// @notice The end height of the edge to be created.
//     /// @dev    End height is deterministic for different edge types but supplying it here gives the
//     ///         caller a bit of extra security that they are supplying data for the correct type of edge
//     uint256 endHeight;
//     /// @notice The edge, or assertion, that is being claimed correct by the newly created edge.
//     bytes32 claimId;
// }

// /// @notice Data parsed raw proof data
// struct ProofData {
//     /// @notice The first state being committed to by an edge
//     bytes32 startState;
//     /// @notice The last state being committed to by an edge
//     bytes32 endState;
//     /// @notice A proof that the end state is included in the egde
//     bytes32[] inclusionProof;
// }

// // CHRIS: TODO: check the ministake was provided

// // CHRIS: TODO: invariants
// // 1. edges are only created, never destroyed
// // 2. all edges have at least one parent, or a claim id - other property invariants exist
// // 3. all edges have a mutual id, and that mutual id must have an entry in firstRivals
// // 4. all values of firstRivals are existing edges (must be in the edge mapping), or are the NO_RIVAL magic hash
// // 5. where to check edge prefix proofs? in bisection, or in add?

// /// @notice Adding layer zero edges requires specific checks. This library contains logic for ensuring that only
// ///         valid layer zero edges can be created
// library LayerZeroEdgeLib {
//     using EdgeChallengeManagerLib for EdgeStore;
//     using ChallengeEdgeLib for ChallengeEdge;

//     /// @notice Conduct checks that are specific to the edge type.
//     /// @dev    Since different edge types also require different proofs, we also include the specific
//     ///         proof parsing logic and return the common parts for later use.
//     /// @param store            The store containing existing edges
//     /// @param assertionChain   The assertion chain containing assertions
//     /// @param args             The edge creation args
//     /// @param proof            Additional proof data to be pars4ed and used
//     /// @return                 Data parsed from the proof, or fetched from elsewhere. Also the origin id for the to be created.
//     function typeSpecificChecks(
//         EdgeStore storage store,
//         IAssertionChain assertionChain,
//         CreateEdgeArgs memory args,
//         bytes memory proof
//     ) internal view returns (ProofData memory, bytes32) {
//         if (args.edgeType == EdgeType.Block) {
//             // origin id is the assertion which is the root of challenge
//             // all rivals and their children share the same origin id - it is a link to the information
//             // they agree on
//             bytes32 originId = assertionChain.getPredecessorId(args.claimId);
//             // if the assertion is already confirmed or rejected then it cant be referenced as a claim
//             require(assertionChain.isPending(args.claimId), "Claim assertion is not pending");
//             // CHRIS: TODO: rename this to "getSibling"? Is it even important?
//             require(assertionChain.getSuccessionChallenge(originId) != 0, "Assertion is not in a fork");

//             // parse the inclusion proof for later use
//             require(proof.length > 0, "Block edge specific proof is empty");
//             bytes32[] memory inclusionProof = abi.decode(proof, (bytes32[]));

//             bytes32 startState = assertionChain.getStateHash(originId);
//             bytes32 endState = assertionChain.getStateHash(args.claimId);
//             return (ProofData(startState, endState, inclusionProof), originId);
//         } else {
//             ChallengeEdge storage claimEdge = store.get(args.claimId);

//             // origin id is the mutual id of the claim
//             // all rivals and their children share the same origin id - it is a link to the information
//             // they agree on
//             bytes32 originId = claimEdge.mutualId();

//             // once a claim is confirmed it's status can never become pending again, so there is no point
//             // opening a challenge that references it
//             require(claimEdge.status == EdgeStatus.Pending, "Claim is not pending");

//             // Claim must be length one. If it is unrivaled then its unrivaled time is ticking up, so there's
//             // no need to create claims against it
//             require(store.hasLengthOneRival(args.claimId), "Claim does not have length 1 rival");

//             // the edge must be a level down from the claim
//             require(args.edgeType == EdgeChallengeManagerLib.nextEdgeType(claimEdge.eType), "Invalid claim edge type");

//             // parse the proofs
//             require(proof.length > 0, "Edge type specific proof is empty");
//             (
//                 bytes32 startState,
//                 bytes32 endState,
//                 bytes32[] memory claimStartInclusionProof,
//                 bytes32[] memory claimEndInclusionProof,
//                 bytes32[] memory edgeInclusionProof
//             ) = abi.decode(proof, (bytes32, bytes32, bytes32[], bytes32[], bytes32[]));

//             // if the start and end states are consistent with the claim edge
//             // this guarantees that the edge we're creating is a 'continuation' of the claim edge, it is
//             // a commitment to the states that between start and end states of the claim
//             MerkleTreeLib.verifyInclusionProof(
//                 claimEdge.startHistoryRoot, startState, claimEdge.startHeight, claimStartInclusionProof
//             );
//             // it's doubly important to check the end state since if the end state since the claim id is
//             // not part of the edge id, so we need to ensure that it's not possible to create two edges of the
//             // same id, but with different claim id. Ensuring that the end state is linked to the claim,
//             // and later ensuring that the end state is part of the history commitment of the new edge ensures
//             // that the end history root of the new edge will be different for different claim ids, and therefore
//             // the edge ids will be different
//             MerkleTreeLib.verifyInclusionProof(
//                 claimEdge.endHistoryRoot, endState, claimEdge.endHeight, claimEndInclusionProof
//             );

//             return (ProofData(startState, endState, edgeInclusionProof), originId);
//         }
//     }

//     /// @notice Zero layer edges have to be a fixed height.
//     ///         This function returns the end height for a given edge height
//     function getLayerZeroEndHeight(EdgeType eType) internal pure returns (uint256) {
//         if (eType == EdgeType.Block) {
//             return LAYERZERO_BLOCKEDGE_HEIGHT;
//         } else if (eType == EdgeType.BigStep) {
//             return LAYERZERO_BIGSTEPEDGE_HEIGHT;
//         } else if (eType == EdgeType.SmallStep) {
//             return LAYERZERO_SMALLSTEPEDGE_HEIGHT;
//         } else {
//             revert("Unrecognised edge type");
//         }
//     }

//     function checkLayerZeroEdge(
//         EdgeStore storage store,
//         IAssertionChain assertionChain,
//         CreateEdgeArgs memory args,
//         bytes calldata prefixProof,
//         bytes calldata proof
//     ) internal view returns (bytes32, bytes32) {
//         // each edge type requires some specific checks
//         (ProofData memory proofData, bytes32 originId) =
//             LayerZeroEdgeLib.typeSpecificChecks(store, assertionChain, args, proof);

//         // since zero layer edges have a start height of zero, we know that they are a size
//         // one tree containing only the start state. We can then compute the history root directly
//         bytes32 startHistoryRoot = MerkleTreeLib.root(MerkleTreeLib.appendLeaf(new bytes32[](0), proofData.startState));

//         // edge have a deterministic end height dependent on their type
//         uint256 endHeight = getLayerZeroEndHeight(args.edgeType);

//         // It isnt strictly necessary to pass in the end height, we know what it
//         // should be so we could just use the end height that we get from getLayerZeroEndHeight
//         // However it's a nice sanity check for the calling code to check that their local edge
//         // will have the same height as the one created here
//         require(args.endHeight == endHeight, "Invalid edge size");

//         // the end state is checked/detemined as part of the specific edge type
//         // We then ensure that that same end state is part of the end history root we're creating
//         // This ensures continuity of states between levels - the state is present in both this
//         // level and the one above
//         MerkleTreeLib.verifyInclusionProof(args.endHistoryRoot, proofData.endState, endHeight, proofData.inclusionProof);

//         // start root must always be a prefix of end root, we ensure that
//         // this new edge adheres to this. Future bisections will ensure that this
//         // property is conserved
//         require(prefixProof.length > 0, "Prefix proof is empty");
//         (bytes32[] memory preExpansion, bytes32[] memory preProof) = abi.decode(prefixProof, (bytes32[], bytes32[]));
//         MerkleTreeLib.verifyPrefixProof(startHistoryRoot, 1, args.endHistoryRoot, endHeight + 1, preExpansion, preProof);

//         return (originId, startHistoryRoot);
//     }
// }

contract EdgeChallengeManager is IEdgeChallengeManager {
    using EdgeChallengeManagerLib for EdgeStore;
    using ChallengeEdgeLib for ChallengeEdge;

    EdgeStore internal store;

    uint256 public challengePeriodSec;
    IAssertionChain internal assertionChain;
    IOneStepProofEntry oneStepProofEntry;

    constructor(IAssertionChain _assertionChain, uint256 _challengePeriodSec, IOneStepProofEntry _oneStepProofEntry) {
        // HN: TODO: remove constructor?
        initialize(_assertionChain, _challengePeriodSec, _oneStepProofEntry);
    }

    function initialize(
        IAssertionChain _assertionChain,
        uint256 _challengePeriodSec,
        IOneStepProofEntry _oneStepProofEntry
    ) public {
        require(address(assertionChain) == address(0), "ALREADY_INIT");
        assertionChain = _assertionChain;
        challengePeriodSec = _challengePeriodSec;
        oneStepProofEntry = _oneStepProofEntry;
    }

    function bisectEdge(bytes32 edgeId, bytes32 bisectionHistoryRoot, bytes memory prefixProof)
        external
        returns (bytes32, bytes32)
    {
        return store.bisectEdge(edgeId, bisectionHistoryRoot, prefixProof);
    }

    function createLayerZeroEdge(CreateEdgeArgs memory args, bytes calldata prefixProof, bytes calldata proof)
        external
        payable
        returns (bytes32)
    {
        return store.createLayerZeroEdge(assertionChain, args, prefixProof, proof);
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

    // CHRIS: TODO: remove these?
    ///////////////////////////////////////////////
    ///////////// VIEW FUNCS ///////////////

    function getPrevAssertionId(bytes32 edgeId) public view returns (bytes32) {
        return store.getPrevAssertionId(edgeId);
    }

    function hasRival(bytes32 edgeId) public view returns (bool) {
        return store.hasRival(edgeId);
    }

    function timeUnrivaled(bytes32 edgeId) public view returns (uint256) {
        return store.timeUnrivaled(edgeId);
    }

    function hasLengthOneRival(bytes32 edgeId) public view returns (bool) {
        return store.hasLengthOneRival(edgeId);
    }

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

    function firstRival(bytes32 edgeId) public view returns (bytes32) {
        return store.firstRivals[edgeId];
    }

    function edgeLength(bytes32 edgeId) public view returns (uint256) {
        return store.get(edgeId).length();
    }
}
