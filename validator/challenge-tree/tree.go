package challengetree

import (
	"fmt"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/OffchainLabs/challenge-protocol-v2/util/threadsafe"
)

type edgeId string
type claimId string
type originId string
type mutualId string
type commit string

type edge struct {
	id           edgeId
	edgeType     protocol.EdgeType
	startHeight  uint64
	startCommit  commit
	endHeight    uint64
	endCommit    commit
	originId     originId
	claimId      claimId
	allHeights   *edgeHeights
	lowerChildId edgeId
	upperChildId edgeId
	creationTime uint64
}

type commitHeights struct {
	start uint64
	end   uint64
}

type edgeHeights struct {
	blockChal     commitHeights
	bigStepChal   util.Option[commitHeights]
	smallStepChal util.Option[commitHeights]
}

func (e *edge) claimHeights() *edgeHeights {
	return e.allHeights
}

func (e *edge) computeMutualId() mutualId {
	return mutualId(e.id[:len(e.id)-1]) // Strip off the last char.
	// return mutualId(fmt.Sprintf(
	// 	"%d-%#x-%d-%#x-%d",
	// 	e.edgeType,
	// 	e.originId,
	// 	e.startHeight,
	// 	e.startCommit,
	// 	e.endHeight,
	// ))
}

// Can check if the local challenge manager agrees with an edge's start
// commitment. If the edge is a block challenge edge, we check if we
// agree with the commitment at the block challenge height.
type HistoryChecker interface {
	AgreesWithStartHistoryCommitment(
		heights *edgeHeights,
		commitMerkle commit,
	) bool
}

type EdgeMetadataReader interface {
	TopLevelAssertion(edgeId) protocol.AssertionId
}

// A challenge tree keeps track of the honest branch for a challenged
// assertion in the protocol. All edges tracked in this data structure
// are part of the same challenge.
type challengeTree struct {
	// TODO: Needs to be thread-safe.
	timeRef                          util.TimeReference
	topLevelAssertionId              protocol.AssertionId
	metadataReader                   EdgeMetadataReader
	histChecker                      HistoryChecker
	edges                            *threadsafe.Map[edgeId, *edge]
	mutualIds                        *threadsafe.Map[mutualId, *threadsafe.Set[edgeId]]
	rivaledEdges                     *threadsafe.Set[edgeId]
	computedPathTimers               *threadsafe.Map[edgeId, uint64]
	honestBlockChalLevelZeroEdge     util.Option[*edge]
	honestBigStepChalLevelZeroEdge   util.Option[*edge]
	honestSmallStepChalLevelZeroEdge util.Option[*edge]
}

func (ct *challengeTree) addEdge(eg *edge) {
	prevAssertionId := ct.metadataReader.TopLevelAssertion(eg.id)
	if ct.topLevelAssertionId != prevAssertionId {
		// Do nothing - this edge should not be part of this challenge tree.
		return
	}

	// Check if the edge id should be added to the rivaled edges set.
	mutualId := eg.computeMutualId()
	if mutuals, ok := ct.mutualIds.Get(mutualId); ok {
		if mutuals.Has(eg.id) {
			ct.rivaledEdges.Insert(eg.id)
		} else {
			mutuals.Insert(eg.id)
		}
	}

	// We only need to check that we agree with the edge's start commitment,
	// and then we will necessarily track all edges we care about for the sake
	// of honest edge confirmations.
	if ct.histChecker.AgreesWithStartHistoryCommitment(
		eg.claimHeights(),
		eg.startCommit,
	) {
		ct.edges.Insert(eg.id, eg)
		if eg.claimId != "" {
			switch eg.edgeType {
			case protocol.BlockChallengeEdge:
				ct.honestBlockChalLevelZeroEdge = util.Some(eg)
			case protocol.BigStepChallengeEdge:
				ct.honestBigStepChalLevelZeroEdge = util.Some(eg)
			case protocol.SmallStepChallengeEdge:
				ct.honestSmallStepChalLevelZeroEdge = util.Some(eg)
			default:
			}
		}
	}
}

// Gets the computed path timer for an edge id being tracked by the challenge tree.
func (ct *challengeTree) GetPathTimer(edgeId edgeId) (uint64, error) {
	total, ok := ct.computedPathTimers.Get(edgeId)
	if !ok {
		return 0, fmt.Errorf("edge id %#x not found in path timers map", edgeId)
	}
	return total, nil
}

// Updates the path timers at the current timestamp for all edges being
// tracked by the challenge tree.
func (ct *challengeTree) updateCumulativeTimers() {
	now := uint64(ct.timeRef.Get().Unix())
	for _, k := range ct.edges.Keys() {
		edge, _ := ct.edges.Get(k)
		timer := ct.pathTimer(edge, now)
		ct.computedPathTimers.Insert(k, timer)
	}
}

// Computes the set of rivals for an edge being tracked by the challenge tree.
// We do this by computing the mutual id of the edge and fetching all edge ids
// that share the same one from a set the challenge tree keeps track of.
// We exclude the specified edge from the returned list of rivals.
func (ct *challengeTree) rivals(edge *edge) []edgeId {
	rivalIds := make([]edgeId, 0)
	if !ct.rivaledEdges.Has(edge.id) {
		return rivalIds
	}
	mutualId := edge.computeMutualId()
	mutuals, ok := ct.mutualIds.Get(mutualId)
	if !ok {
		return rivalIds
	}
	for item := range mutuals.CopyItems() {
		if item == edge.id {
			continue
		}
		rivalIds = append(rivalIds, item)
	}
	return rivalIds
}

// Consider the following set of edges in a challenge where evil
// edges are marked with a ' and a *:
//
//	/---6---8
//
// 0-----4
//
//	\---6'--8'
//
// 0*-------------8*
//
// The honest branch is the one that goes from 0-8. The evil edge is 0-8'.
// The evil edge 0-8' bisects, but agrees with the honest one from 0-4.
// Therefore, there is only a single 0-4 edge in the set.
//
// In this case, the set of ancestors for 4-6 is the following:
//
//	{4-8, 4-8', 0-8, 0-8'}
//
// All of these ancestors will contribute towards the path timer of 4-6 when
// confirmations by time are being attempted for the edge. Note there is another
// evil party that starts at 0*. In this case, it does not agree with the honest
// branch at all, so that party's edges can be ignored for this computation.
//
// In order to retrieve ancestors for an edge with id=I, we start from the honest,
// block challenge level zero edge and recursively traverse its children,
// reducing the ids and their rivals along the way into a slice until we hit a child that
// matches id=I and return the slice.
func (ct *challengeTree) ancestorsForHonestEdge(id edgeId) []edgeId {
	if ct.honestBlockChalLevelZeroEdge.IsNone() {
		return make([]edgeId, 0)
	}
	blockEdge := ct.honestBlockChalLevelZeroEdge.Unwrap()
	ancestors, ok := ct.ancestorQuery(
		make([]edgeId, 0),
		blockEdge,
		id,
	)
	if !ok {
		return nil
	}
	// The confirm by time function in Solidity requires ancestors to be specified
	// from earliest to oldest, which is the reverse result of our recursion.
	util.Reverse(ancestors)
	return ancestors
}

func (ct *challengeTree) ancestorQuery(
	accum []edgeId,
	curr *edge,
	queryingFor edgeId,
) ([]edgeId, bool) {
	if curr.id == queryingFor {
		return accum, true
	}
	if !hasChildren(curr) {
		// If the edge has length 1, we then perform a few special checks.
		if edgeLength(curr) == 1 {
			// In case the edge is a small step challenge of length 1, we simply return.
			if curr.edgeType == protocol.SmallStepChallengeEdge {
				return accum, false
			}

			// If the edge is unrivaled, we return.
			hasRival := ct.rivaledEdges.Has(curr.id)
			if !hasRival {
				return accum, false
			}

			rivalIds := ct.rivals(curr)

			// If the edge is a block challenge edge, we continue the recursion starting from the honest
			// big step level zero edge, if it exists.
			if curr.edgeType == protocol.BlockChallengeEdge {
				if ct.honestBigStepChalLevelZeroEdge.IsNone() {
					return accum, false
				}
				honestLowerLevelEdge := ct.honestBigStepChalLevelZeroEdge.Unwrap()

				// Defensive check ensuring the honest level zero edge one challenge level below
				// claims the current edge id as its claim id.
				if honestLowerLevelEdge.claimId != claimId(curr.id) {
					return accum, false
				}
				accum = append(accum, rivalIds...)
				accum = append(accum, curr.id)
				return ct.ancestorQuery(accum, honestLowerLevelEdge, queryingFor)
			}

			// If the edge is a big step challenge edge, we continue the recursion starting from the honest
			// small step level zero edge, if it exists.
			if curr.edgeType == protocol.BigStepChallengeEdge {
				if ct.honestSmallStepChalLevelZeroEdge.IsNone() {
					return accum, false
				}
				honestLowerLevelEdge := ct.honestSmallStepChalLevelZeroEdge.Unwrap()

				// Defensive check ensuring the honest level zero edge one challenge level below
				// claims the current edge id as its claim id.
				if honestLowerLevelEdge.claimId != claimId(curr.id) {
					return accum, false
				}

				accum = append(accum, rivalIds...)
				accum = append(accum, curr.id)
				return ct.ancestorQuery(accum, honestLowerLevelEdge, queryingFor)
			}
		}
		return accum, false
	}
	rivalIds := ct.rivals(curr)
	accum = append(accum, rivalIds...)
	accum = append(accum, curr.id)

	// If the edge id we are querying for is a direct child of the current edge, we append
	// the current edge to the ancestors list and return true.
	if isDirectChild(curr, queryingFor) {
		return accum, true
	}
	lowerChild, lowerOk := ct.edges.Get(curr.lowerChildId)
	if !lowerOk {
		panic("not lower")
	}
	upperChild, upperOk := ct.edges.Get(curr.upperChildId)
	if !upperOk {
		panic(fmt.Sprintf("not upper curr %s, upper=%s", curr.id, curr.upperChildId))
	}
	lowerAncestors, foundInLowerChildren := ct.ancestorQuery(
		accum, lowerChild, queryingFor,
	)
	upperAncestors, foundInUpperChildren := ct.ancestorQuery(
		accum,
		upperChild,
		queryingFor,
	)
	// If the edge we are querying for is found in the lower children,
	// we return the ancestry along such path.
	if foundInLowerChildren {
		return lowerAncestors, true
	}
	// If the edge we are querying for is found in the upper children,
	// we return the ancestry along such path.
	if foundInUpperChildren {
		return upperAncestors, true
	}
	return accum, false
}

// Computes the length of an edge by taking the difference between
// its end and start heights.
// SAFETY: We will never receive a malformed edge, as the challenge tree is
// created from events emitted by successful challenge addition events
// in the protocol smart contracts.
func edgeLength(eg *edge) uint64 {
	return eg.endHeight - eg.startHeight
}

// Checks if an edge id is a direct child of a specified parent edge.
func isDirectChild(parent *edge, childId edgeId) bool {
	return parent.lowerChildId == childId || parent.upperChildId == childId
}

// Checks if an edge has any children.
func hasChildren(eg *edge) bool {
	return eg.lowerChildId != "" || eg.upperChildId != ""
}
