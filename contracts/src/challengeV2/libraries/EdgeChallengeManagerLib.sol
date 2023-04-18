// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.17;

import "./UintUtilsLib.sol";
import "./MerkleTreeLib.sol";
import "./ChallengeEdgeLib.sol";
import "../../osp/IOneStepProofEntry.sol";

/// @notice Stores all edges and their rival status
struct EdgeStore {
    /// @dev A mapping of edge id to edges. Edges are never deleted, only created, and potentially confirmed.
    mapping(bytes32 => ChallengeEdge) edges;
    /// @dev A mapping of mutualId to edge id. Rivals share the same mutual id, and here we
    ///      store the edge id of the second edge that was created with the same mutual id - the first rival
    ///      When only one edge exists for a specific mutual id then a special magic string hash is stored instead
    ///      of the first rival id, to signify that a single edge does exist with this mutual id
    mapping(bytes32 => bytes32) firstRivals;
}

/// @notice Input data to a one step proof
struct OneStepData {
    /// @notice The one step proof execution context
    ExecutionContext execCtx;
    /// @notice The machine counter of the state that's being executed from
    uint256 machineStep;
    /// @notice The hash of the state that's being executed from
    bytes32 beforeHash;
    /// @notice Proof data to accompany the execution context
    bytes proof;
}

/// @title  Core functionality for the Edge Challenge Manager
/// @notice The edge manager library allows edges to be added and bisected, and keeps track of the amount
///         of time an edge remained unrivaled.
library EdgeChallengeManagerLib {
    using ChallengeEdgeLib for ChallengeEdge;

    /// @dev Magic string hash to represent that a edges with a given mutual id have no rivals
    bytes32 constant UNRIVALED = keccak256(abi.encodePacked("UNRIVALED"));

    /// @notice A new edge has been added to the challenge manager
    /// @param edgeId       The id of the newly added edge
    /// @param mutualId     The mutual id of the added edge - all rivals share the same mutual id
    /// @param originId     The origin id of the added edge - origin ids link an edge to the level above
    /// @param hasRival     Does the newly added edge have a rival upon creation
    /// @param length       The length of the new edge
    /// @param eType        The type of the new edge
    /// @param isLayerZero  Whether the new edge was added at layer zero - has a claim and a staker
    event EdgeAdded(
        bytes32 indexed edgeId,
        bytes32 indexed mutualId,
        bytes32 indexed originId,
        bool hasRival,
        uint256 length,
        EdgeType eType,
        bool isLayerZero
    );

    /// @notice An edge has been bisected
    /// @param edgeId                   The id of the edge that was bisected
    /// @param lowerChildId             The id of the lower child created during bisection
    /// @param upperChildId             The id of the upper child created during bisection
    /// @param lowerChildAlreadyExists  When an edge is bisected the lower child may already exist - created by a rival.
    event EdgeBisected(
        bytes32 indexed edgeId, bytes32 indexed lowerChildId, bytes32 indexed upperChildId, bool lowerChildAlreadyExists
    );

    /// @notice An edge can be confirmed if both of its children were already confirmed.
    /// @param edgeId   The edge that was confirmed
    /// @param mutualId The mutual id of the confirmed edge
    event EdgeConfirmedByChildren(bytes32 indexed edgeId, bytes32 indexed mutualId);

    /// @notice An edge can be confirmed if the cumulative time unrivaled of it and a direct chain of ancestors is greater than a threshold
    /// @param edgeId               The edge that was confirmed
    /// @param mutualId             The mutual id of the confirmed edge
    /// @param totalTimeUnrivaled   The cumulative amount of time this edge spent unrivaled
    event EdgeConfirmedByTime(bytes32 indexed edgeId, bytes32 indexed mutualId, uint256 totalTimeUnrivaled);

    /// @notice An edge can be confirmed if a zero layer edge in the level below claims this edge
    /// @param edgeId           The edge that was confirmed
    /// @param mutualId         The mutual id of the confirmed edge
    /// @param claimingEdgeId   The id of the zero layer edge that claimed this edge
    event EdgeConfirmedByClaim(bytes32 indexed edgeId, bytes32 indexed mutualId, bytes32 claimingEdgeId);

    /// @notice A SmallStep edge of length 1 can be confirmed via a one step proof
    /// @param edgeId   The edge that was confirmed
    /// @param mutualId The mutual id of the confirmed edge
    event EdgeConfirmedByOneStepProof(bytes32 indexed edgeId, bytes32 indexed mutualId);

    /// @notice Get an edge from the store
    /// @dev    Throws if the edge does not exist in the store
    /// @param store    The edge store to fetch an id from
    /// @param edgeId   The id of the edge to fetch
    function get(EdgeStore storage store, bytes32 edgeId) internal view returns (ChallengeEdge storage) {
        require(store.edges[edgeId].exists(), "Edge does not exist");
        return store.edges[edgeId];
    }

    /// @notice Adds a new edge to the store
    /// @dev    Updates first rival info for later use in calculating time unrivaled
    /// @param store    The store to add the edge to
    /// @param edge     The edge to add
    function add(EdgeStore storage store, ChallengeEdge memory edge) internal {
        bytes32 eId = edge.idMem();
        // add the edge if it doesnt exist already
        require(!store.edges[eId].exists(), "Edge already exists");
        store.edges[eId] = edge;

        // edges that are rivals share the same mutual id
        // we use records of whether a mutual id has ever been added to decide if
        // the new edge is a rival. This will later allow us to calculate time an edge
        // stayed unrivaled
        bytes32 mutualId = ChallengeEdgeLib.mutualIdComponent(
            edge.eType, edge.originId, edge.startHeight, edge.startHistoryRoot, edge.endHeight
        );
        bytes32 firstRival = store.firstRivals[mutualId];

        // the first time we add a mutual id we store a magic string hash against it
        // We do this to distinguish from there being no edges
        // with this mutual. And to distinguish it from the first rival, where we
        // will use an actual edge id so that we can look up the created when time
        // of the first rival, and use it for calculating time unrivaled
        if (firstRival == 0) {
            store.firstRivals[mutualId] = UNRIVALED;
        } else if (firstRival == UNRIVALED) {
            store.firstRivals[mutualId] = eId;
        } else {
            // after we've stored the first rival we dont need to keep a record of any
            // other rival edges - they will all have a zero time unrivaled
        }

        emit EdgeAdded(
            eId,
            mutualId,
            edge.originId,
            firstRival != 0,
            store.edges[eId].length(),
            edge.eType,
            edge.staker != address(0)
        );
    }

    /// @notice Does this edge currently have one or more rivals
    ///         Rival edges share the same startHeight, startHistoryCommitment and the same endHeight,
    ///         but they have a different endHistoryRoot. Rival edges have the same mutualId
    /// @param store    The edge store containing the edge
    /// @param edgeId   The edge if to test if it is unrivaled
    function hasRival(EdgeStore storage store, bytes32 edgeId) internal view returns (bool) {
        require(store.edges[edgeId].exists(), "Edge does not exist");

        // rivals have the same mutual id
        bytes32 mutualId = store.edges[edgeId].mutualId();
        bytes32 firstRival = store.firstRivals[mutualId];
        // Sanity check: it should never be possible to create an edge without having an entry in firstRivals
        require(firstRival != 0, "Empty first rival");

        // can only have no rival if the firstRival is the UNRIVALED magic hash
        return firstRival != UNRIVALED;
    }

    /// @notice Is the edge a single step in length, and does it have at least one rival.
    /// @param store    The edge store containing the edge
    /// @param edgeId   The edge id to test for single step and rivaled
    function hasLengthOneRival(EdgeStore storage store, bytes32 edgeId) internal view returns (bool) {
        // must be length 1 and have rivals - all rivals have the same length
        return (hasRival(store, edgeId) && store.edges[edgeId].length() == 1);
    }

    /// @notice The amount of time this edge has spent without rivals
    ///         This value is increasing whilst an edge is unrivaled, once a rival is created
    ///         it is fixed. If an edge has rivals from the moment it is created then it will have
    ///         a zero time unrivaled
    function timeUnrivaled(EdgeStore storage store, bytes32 edgeId) internal view returns (uint256) {
        require(store.edges[edgeId].exists(), "Edge does not exist");

        bytes32 mutualId = store.edges[edgeId].mutualId();
        bytes32 firstRival = store.firstRivals[mutualId];
        // Sanity check: it's not possible to have a 0 first rival for an edge that exists
        require(firstRival != 0, "Empty rival record");

        // this edge has no rivals, the time is still going up
        // we give the current amount of time unrivaled
        if (firstRival == UNRIVALED) {
            return block.timestamp - store.edges[edgeId].createdWhen;
        } else {
            // Sanity check: it's not possible an edge does not exist for a first rival record
            require(store.edges[firstRival].exists(), "Rival edge does not exist");

            // rivals exist for this edge
            uint256 firstRivalCreatedWhen = store.edges[firstRival].createdWhen;
            uint256 edgeCreatedWhen = store.edges[edgeId].createdWhen;
            if (firstRivalCreatedWhen > edgeCreatedWhen) {
                // if this edge was created before the first rival then we return the difference
                // in createdWhen times
                return firstRivalCreatedWhen - edgeCreatedWhen;
            } else {
                // if this was created at the same time as, or after the the first rival
                // then we return 0
                return 0;
            }
        }
    }

    /// @notice Given a start and an endpoint determine the bisection height
    /// @dev    Returns the highest power of 2 in the differing lower bits of start and end
    function mandatoryBisectionHeight(uint256 start, uint256 end) internal pure returns (uint256) {
        require(end - start >= 2, "Height difference not two or more");
        if (end - start == 2) {
            return start + 1;
        }

        uint256 diff = (end - 1) ^ start;
        uint256 mostSignificantSharedBit = UintUtilsLib.mostSignificantBit(diff);
        uint256 mask = type(uint256).max << mostSignificantSharedBit;
        return ((end - 1) & mask);
    }

    /// @notice Bisect and edge. This creates two child edges:
    ///         lowerChild: has the same start root and height as this edge, but a different end root and height
    ///         upperChild: has the same end root and height as this edge, but a different start root and height
    ///         The lower child end root and height are equal to the upper child start root and height. This height
    ///         is the mandatoryBisectionHeight
    /// @param store                The edge store containing the edge to bisect
    /// @param edgeId               Edge to bisect
    /// @param bisectionHistoryRoot The new history root to be used in the lower and upper children
    /// @param prefixProof          A proof to show that the bisectionHistoryRoot commits to a prefix of the current endHistoryRoot
    /// @return lowerChildId        The id of the newly created lower child edge
    /// @return upperChildId        The id of the newly created upper child edge
    function bisectEdge(EdgeStore storage store, bytes32 edgeId, bytes32 bisectionHistoryRoot, bytes memory prefixProof)
        internal
        returns (bytes32, bytes32)
    {
        require(store.edges[edgeId].status == EdgeStatus.Pending, "Edge not pending");
        require(hasRival(store, edgeId), "Cannot bisect an unrivaled edge");

        // cannot bisect an edge twice
        ChallengeEdge memory ce = get(store, edgeId);
        require(
            store.edges[edgeId].lowerChildId == 0 && store.edges[edgeId].upperChildId == 0, "Edge already has children"
        );

        // bisections occur at deterministic heights, this ensures that
        // rival edges bisect at the same height, and create the same child if they agree
        uint256 middleHeight = mandatoryBisectionHeight(ce.startHeight, ce.endHeight);
        {
            (bytes32[] memory preExpansion, bytes32[] memory proof) = abi.decode(prefixProof, (bytes32[], bytes32[]));
            MerkleTreeLib.verifyPrefixProof(
                bisectionHistoryRoot, middleHeight + 1, ce.endHistoryRoot, ce.endHeight + 1, preExpansion, proof
            );
        }

        bytes32 lowerChildId;
        bool lowerChildExists;
        {
            // midpoint proof it valid, create and store the children
            ChallengeEdge memory lowerChild = ChallengeEdgeLib.newChildEdge(
                ce.originId, ce.startHistoryRoot, ce.startHeight, bisectionHistoryRoot, middleHeight, ce.eType
            );
            lowerChildId = lowerChild.idMem();
            // it's possible that the store already has the lower child if it was created by a rival
            // (aka a merge move)
            if (store.edges[lowerChildId].exists()) {
                lowerChildExists = true;
            } else {
                add(store, lowerChild);
                lowerChildExists = false;
            }
        }

        bytes32 upperChildId;
        {
            ChallengeEdge memory upperChild = ChallengeEdgeLib.newChildEdge(
                ce.originId, bisectionHistoryRoot, middleHeight, ce.endHistoryRoot, ce.endHeight, ce.eType
            );
            upperChildId = upperChild.idMem();

            // Sanity check: it's not possible that the upper child already exists, for this to be the case
            // the edge would have to have been bisected already.
            require(!store.edges[upperChildId].exists(), "Store contains upper child");
            add(store, upperChild);
        }

        store.edges[edgeId].setChildren(lowerChildId, upperChildId);

        emit EdgeBisected(edgeId, lowerChildId, upperChildId, lowerChildExists);

        return (lowerChildId, upperChildId);
    }

    /// @notice Confirm an edge if both its children are already confirmed
    function confirmEdgeByChildren(EdgeStore storage store, bytes32 edgeId) internal {
        require(store.edges[edgeId].exists(), "Edge does not exist");
        require(store.edges[edgeId].status == EdgeStatus.Pending, "Edge not pending");

        bytes32 lowerChildId = store.edges[edgeId].lowerChildId;
        // Sanity check: it bisect should already enforce that this child exists
        require(store.edges[lowerChildId].exists(), "Lower child does not exist");
        require(store.edges[lowerChildId].status == EdgeStatus.Confirmed, "Lower child not confirmed");

        bytes32 upperChildId = store.edges[edgeId].upperChildId;
        // Sanity check: it bisect should already enforce that this child exists
        require(store.edges[upperChildId].exists(), "Upper child does not exist");
        require(store.edges[upperChildId].status == EdgeStatus.Confirmed, "Upper child not confirmed");

        store.edges[edgeId].setConfirmed();

        emit EdgeConfirmedByChildren(edgeId, store.edges[edgeId].mutualId());
    }

    /// @notice Returns the sub edge type of the provided edge type
    function nextEdgeType(EdgeType eType) internal pure returns (EdgeType) {
        if (eType == EdgeType.Block) {
            return EdgeType.BigStep;
        } else if (eType == EdgeType.BigStep) {
            return EdgeType.SmallStep;
        } else if (eType == EdgeType.SmallStep) {
            revert("No next type after SmallStep");
        } else {
            revert("Unexpected edge type");
        }
    }

    /// @notice Check that the originId of a claiming edge matched the mutualId() of a supplied edge
    /// @dev    Does some additional sanity checks to ensure that the claim id link is valid
    /// @param store            The store containing all edges and rivals
    /// @param edgeId           The edge being claimed
    /// @param claimingEdgeId   The edge with a claim id equal to edge id
    function checkClaimIdLink(EdgeStore storage store, bytes32 edgeId, bytes32 claimingEdgeId) private view {
        // we do some extra checks that edge being claimed is eligible to be claimed by the claiming edge
        // these shouldn't be necessary since it should be impossible to add layer zero edges that do not
        // satisfy the checks below, but we conduct these checks anyway for double safety

        // the origin id of an edge should be the mutual id of the edge in the level above
        require(store.edges[edgeId].mutualId() == store.edges[claimingEdgeId].originId, "Origin id-mutual id mismatch");
        // the claiming edge must be exactly one level below
        require(
            nextEdgeType(store.edges[edgeId].eType) == store.edges[claimingEdgeId].eType,
            "Edge type does not match claiming edge type"
        );
    }

    /// @notice If a confirmed edge exists whose claim id is equal to this edge, then this edge can be confirmed
    /// @dev    When zero layer edges are created they reference an edge, or assertion, in the level above. If a zero layer
    ///         edge is confirmed, it becomes possible to also confirm the edge that it claims
    /// @param store            The store containing all edges and rivals data
    /// @param edgeId           The id of the edge to confirm
    /// @param claimingEdgeId   The id of the edge which has a claimId equal to edgeId
    function confirmEdgeByClaim(EdgeStore storage store, bytes32 edgeId, bytes32 claimingEdgeId) internal {
        // this edge is pending
        require(store.edges[edgeId].exists(), "Edge does not exist");
        require(store.edges[edgeId].status == EdgeStatus.Pending, "Edge not pending");
        // the claiming edge is confirmed
        require(store.edges[claimingEdgeId].exists(), "Claiming edge does not exist");
        require(store.edges[claimingEdgeId].status == EdgeStatus.Confirmed, "Claiming edge not confirmed");

        checkClaimIdLink(store, edgeId, claimingEdgeId);
        require(edgeId == store.edges[claimingEdgeId].claimId, "Claim does not match edge");

        store.edges[edgeId].setConfirmed();

        emit EdgeConfirmedByClaim(edgeId, store.edges[edgeId].mutualId(), claimingEdgeId);
    }

    /// @notice An edge can be confirmed if the total amount of time it and a single chain of its direct ancestors
    ///         has spent unrivaled is greater than the challenge period.
    /// @dev    Edges inherit time from their parents, so the sum of unrivaled timers is compared against the threshold.
    ///         Given that an edge cannot become unrivaled after becoming rivaled, once the threshold is passed
    ///         it will always remain passed. The direct ancestors of an edge are linked by parent-child links for edges
    ///         of the same edgeType, and claimId-edgeid links for zero layer edges that claim an edge in the level above.
    /// @param store                    The edge store containing all edges and rival data
    /// @param edgeId                   The id of the edge to confirm
    /// @param ancestorEdgeIds          The ids of the direct ancestors of an edge. These are ordered from the parent first, then going to grand-parent,
    ///                                 great-grandparent etc. The chain can extend only as far as the zero layer edge of type Block.
    /// @param confirmationThresholdSec The amount of time in seconds that the total unrivaled time of an ancestor chain needs to exceed in
    ///                                 order to be confirmed
    function confirmEdgeByTime(
        EdgeStore storage store,
        bytes32 edgeId,
        bytes32[] memory ancestorEdgeIds,
        uint256 confirmationThresholdSec
    ) internal {
        require(store.edges[edgeId].exists(), "Edge does not exist");
        require(store.edges[edgeId].status == EdgeStatus.Pending, "Edge not pending");

        bytes32 currentEdgeId = edgeId;
        uint256 totalTimeUnrivaled = timeUnrivaled(store, edgeId);

        // ancestors start from parent, then extend upwards
        for (uint256 i = 0; i < ancestorEdgeIds.length; i++) {
            ChallengeEdge storage e = get(store, ancestorEdgeIds[i]);
            // the ancestor must either have a parent-child link
            // or have a claim id-edge link when the ancestor is of a different edge type to its child
            if (e.lowerChildId == currentEdgeId || e.upperChildId == currentEdgeId) {
                totalTimeUnrivaled += timeUnrivaled(store, e.id());
                currentEdgeId = ancestorEdgeIds[i];
            } else if (ancestorEdgeIds[i] == store.edges[currentEdgeId].claimId) {
                checkClaimIdLink(store, ancestorEdgeIds[i], currentEdgeId);
                totalTimeUnrivaled += timeUnrivaled(store, e.id());
                currentEdgeId = ancestorEdgeIds[i];
            } else {
                revert("Current is not a child of ancestor");
            }
        }

        require(
            totalTimeUnrivaled > confirmationThresholdSec,
            "Total time unrivaled not greater than confirmation threshold"
        );

        store.edges[edgeId].setConfirmed();

        emit EdgeConfirmedByTime(edgeId, store.edges[edgeId].mutualId(), totalTimeUnrivaled);
    }

    /// @notice Confirm an edge by executing a one step proof
    /// @dev    One step proofs can only be executed against edges that have length one and of type SmallStep
    /// @param store                        The edge store containing all edges and rival data
    /// @param edgeId                       The id of the edge to confirm
    /// @param oneStepProofEntry            The one step proof contract
    /// @param oneStepData                  Input data to the one step proof
    /// @param beforeHistoryInclusionProof  Proof that the state which is the start of the edge is committed to by the startHistoryRoot
    /// @param afterHistoryInclusionProof   Proof that the state which is the end of the edge is committed to by the endHistoryRoot
    function confirmEdgeByOneStepProof(
        EdgeStore storage store,
        bytes32 edgeId,
        IOneStepProofEntry oneStepProofEntry,
        OneStepData memory oneStepData,
        bytes32[] memory beforeHistoryInclusionProof,
        bytes32[] memory afterHistoryInclusionProof
    ) internal {
        require(store.edges[edgeId].exists(), "Edge does not exist");
        require(store.edges[edgeId].status == EdgeStatus.Pending, "Edge not pending");

        // edge must be length one and be of type SmallStep
        require(store.edges[edgeId].eType == EdgeType.SmallStep, "Edge is not a small step");
        require(store.edges[edgeId].length() == 1, "Edge does not have single step");

        // the state in the onestep data must be committed to by the startHistoryRoot
        require(
            MerkleTreeLib.verifyInclusionProof(
                store.edges[edgeId].startHistoryRoot,
                oneStepData.beforeHash,
                oneStepData.machineStep,
                beforeHistoryInclusionProof
            ),
            "Before state not in history"
        );

        // execute the single step to produce the after state
        bytes32 afterHash = oneStepProofEntry.proveOneStep(
            oneStepData.execCtx, oneStepData.machineStep, oneStepData.beforeHash, oneStepData.proof
        );

        // check that the after state was indeed committed to by the endHistoryRoot
        require(
            MerkleTreeLib.verifyInclusionProof(
                store.edges[edgeId].endHistoryRoot, afterHash, oneStepData.machineStep + 1, afterHistoryInclusionProof
            ),
            "After state not in history"
        );

        store.edges[edgeId].setConfirmed();

        emit EdgeConfirmedByOneStepProof(edgeId, store.edges[edgeId].mutualId());
    }
}
