package challengetree

import (
	"fmt"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/OffchainLabs/challenge-protocol-v2/util/threadsafe"
	"github.com/ethereum/go-ethereum/common"
)

type edge struct {
	id           protocol.EdgeId
	edgeType     protocol.EdgeType
	startHeight  uint64
	startCommit  common.Hash
	endHeight    uint64
	endCommit    common.Hash
	originId     common.Hash
	claimId      common.Hash
	lowerChildId common.Hash
	upperChildId common.Hash
	allHeights   *edgeHeights
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

func (e *edge) computeMutualId() protocol.MutualId {
	return protocol.MutualId(common.BytesToHash([]byte(
		fmt.Sprintf(
			"%d-%#x-%d-%#x-%d",
			e.edgeType,
			e.originId,
			e.startHeight,
			e.startCommit,
			e.endHeight,
		),
	)))
}

type chain interface {
	timeUnrivaled(edgeId protocol.EdgeId) uint64
}

// Can check if the local challenge manager agrees with an edge's start
// commitment. If the edge is a block challenge edge, we check if we
// agree with the commitment at the block challenge height.
type HistoryChecker interface {
	AgreesWithStartHistoryCommitment(
		heights *edgeHeights,
		commitMerkle common.Hash,
	) bool
}

type edgeMetadataReader interface {
	topLevelAssertion(protocol.EdgeId) protocol.AssertionId
}

// A challenge tree keeps track of the honest branch for a challenged
// assertion in the protocol. All edges tracked in this data structure
// are part of the same challenge.
type challengeTree struct {
	// TODO: Needs to be thread-safe.
	topLevelAssertionId              protocol.AssertionId
	chain                            chain
	metadataReader                   edgeMetadataReader
	histChecker                      HistoryChecker
	edges                            *threadsafe.Map[protocol.EdgeId, *edge]
	mutualIds                        *threadsafe.Map[protocol.MutualId, *threadsafe.Set[protocol.EdgeId]]
	rivaledEdges                     *threadsafe.Set[protocol.EdgeId]
	honestUnrivaledCumulativeTimers  *threadsafe.Map[protocol.EdgeId, uint64]
	honestBlockChalLevelZeroEdge     util.Option[*edge]
	honestBigStepChalLevelZeroEdge   util.Option[*edge]
	honestSmallStepChalLevelZeroEdge util.Option[*edge]
}

func (ct *challengeTree) addEdge(eg *edge) {
	prevAssertionId := ct.metadataReader.topLevelAssertion(eg.id)
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

	// TODO: We only need to check that we agree with the edge's start commitment,
	// and then we will necessarily track all edges we care about.
	// If this is an honest edge from our perspective, we keep track
	// of it as such in our implementation.
	if ct.histChecker.AgreesWithStartHistoryCommitment(
		eg.claimHeights(),
		eg.startCommit,
	) {
		ct.edges.Insert(eg.id, eg)
		if eg.claimId != (common.Hash{}) {
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

func (ct *challengeTree) CumulativeTimeUnrivaled(edgeId protocol.EdgeId) (uint64, error) {
	total, ok := ct.honestUnrivaledCumulativeTimers.Get(edgeId)
	if !ok {
		return 0, fmt.Errorf("edge id %#x not found in cumulative timers map", edgeId)
	}
	return total, nil
}

// Get the honest level zero edge from our list of honest
// edges (keep track of them per challenge level).
// Recursively go down its children and then update their
// cumulative timers accordingly.
func (ct *challengeTree) updateCumulativeTimers() {
	if ct.honestBlockChalLevelZeroEdge.IsNone() {
		return
	}
	blockEdge := ct.honestBlockChalLevelZeroEdge.Unwrap()
	ct.innerCumulativeUpdate(0, blockEdge.id)
}

func (ct *challengeTree) innerCumulativeUpdate(
	cumulativeUnrivaledTime uint64,
	edgeId protocol.EdgeId,
) {
	edge, _ := ct.edges.Get(edgeId)
	blocksUnrivaled := ct.chain.timeUnrivaled(edgeId)
	total := blocksUnrivaled + cumulativeUnrivaledTime
	ct.honestUnrivaledCumulativeTimers.Insert(edgeId, total)
	if !hasChildren(edge) {
		// If the edge has length 1, we then perform a few special checks.
		if edgeLength(edge) == 1 {

			// In case the edge is a small step challenge of length 1, we simply return.
			if edge.edgeType == protocol.SmallStepChallengeEdge {
				return
			}

			// If the edge is unrivaled, we return.
			hasRival := ct.rivaledEdges.Has(edgeId)
			if !hasRival {
				return
			}

			// If the edge is a block challenge, we continue the recursion starting from the honest
			// big step level zero edge, if it exists.
			if edge.edgeType == protocol.BlockChallengeEdge {
				if ct.honestBigStepChalLevelZeroEdge.IsNone() {
					return
				}
				honestLowerLevelEdge := ct.honestBigStepChalLevelZeroEdge.Unwrap()

				// Defensive check ensuring the honest level zero edge one challenge level below
				// claims the current edge id as its claim id.
				if honestLowerLevelEdge.claimId != common.Hash(edgeId) {
					return
				}

				ct.innerCumulativeUpdate(total, honestLowerLevelEdge.id)
				return
			}
			// If the edge is a big step challenge, we continue the recursion starting from the honest
			// big step level zero edge, if it exists.
			if edge.edgeType == protocol.BigStepChallengeEdge {
				if ct.honestSmallStepChalLevelZeroEdge.IsNone() {
					return
				}
				honestLowerLevelEdge := ct.honestSmallStepChalLevelZeroEdge.Unwrap()

				// Defensive check ensuring the honest level zero edge one challenge level below
				// claims the current edge id as its claim id.
				if honestLowerLevelEdge.claimId != common.Hash(edgeId) {
					return
				}

				ct.innerCumulativeUpdate(total, honestLowerLevelEdge.id)
			}
		}
		return
	}
	// We recursively update the children's timers.
	ct.innerCumulativeUpdate(total, protocol.EdgeId(edge.lowerChildId))
	ct.innerCumulativeUpdate(total, protocol.EdgeId(edge.upperChildId))
}

func (ct *challengeTree) rivalIds(edge *edge) []protocol.EdgeId {
	mutualId := edge.computeMutualId()
	mutuals, ok := ct.mutualIds.Get(mutualId)
	if !ok {
		return make([]protocol.EdgeId, 0)
	}
	rivalIds := make([]protocol.EdgeId, 0)
	for item := range mutuals.CopyItems() {
		if item == edge.id {
			continue
		}
		rivalIds = append(rivalIds, item)
	}
	return rivalIds
}

func (ct *challengeTree) ancestorsForHonestEdge(id protocol.EdgeId) []protocol.EdgeId {
	// Consider the following set of edges in a challenge.
	//
	// Honest edges
	// 0-----------------------8
	// 0-------4, 4------------8
	//            4----6, 6----8
	//
	// Evil edges
	// 0-----------------------8'
	//            4------------8'
	//            4----6',6'---8'
	//
	// The honest branch is the one that goes from 0-8. The evil edge is 0-8'.
	// The evil edge 0-8' bisects, but agrees with the honest one from 0-4.
	// Therefore, there is only a single 0-4 edge in the set.
	//
	// In this case, the challenge tree's list of honest edge ids will be:
	//
	//   [id(0,4), id(4,6), id(6,8), id(0,8)]
	//
	// In order to retrieve ancestors for an edge with id=I, we start from the honest,
	// block challenge level zero edge and recursively traverse its children,
	// reducing the ids along the way into a slice until we hit a child that
	// matches id=I and return the slice.
	if ct.honestBlockChalLevelZeroEdge.IsNone() {
		return make([]protocol.EdgeId, 0)
	}
	blockEdge := ct.honestBlockChalLevelZeroEdge.Unwrap()
	ancestors, ok := ct.ancestorQuery(
		make([]protocol.EdgeId, 0),
		blockEdge,
		id,
	)
	if !ok {
		return nil
	}
	// The confirm by time function in Solidity requires ancestors to be specified
	// from earliest to oldest, which is the reverse result of our recursion.
	reverse(ancestors)
	return ancestors
}

func (ct *challengeTree) ancestorQuery(
	accum []protocol.EdgeId,
	curr *edge,
	queryingFor protocol.EdgeId,
) ([]protocol.EdgeId, bool) {
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

			rivalIds := ct.rivalIds(curr)

			// If the edge is a block challenge edge, we continue the recursion starting from the honest
			// big step level zero edge, if it exists.
			if curr.edgeType == protocol.BlockChallengeEdge {
				if ct.honestBigStepChalLevelZeroEdge.IsNone() {
					return accum, false
				}
				honestLowerLevelEdge := ct.honestBigStepChalLevelZeroEdge.Unwrap()

				// Defensive check ensuring the honest level zero edge one challenge level below
				// claims the current edge id as its claim id.
				if honestLowerLevelEdge.claimId != common.Hash(curr.id) {
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
				if honestLowerLevelEdge.claimId != common.Hash(curr.id) {
					return accum, false
				}

				accum = append(accum, rivalIds...)
				accum = append(accum, curr.id)
				return ct.ancestorQuery(accum, honestLowerLevelEdge, queryingFor)
			}
		}
		return accum, false
	}
	rivalIds := ct.rivalIds(curr)
	accum = append(accum, rivalIds...)
	accum = append(accum, curr.id)
	// If the edge id we are querying for is a child of the current edge, we append
	// the current edge to the ancestors list and return true.
	if isChild(curr, queryingFor) {
		return accum, true
	}
	lowerChild, lowerOk := ct.edges.Get(protocol.EdgeId(curr.lowerChildId))
	if !lowerOk {
		panic("not lower")
	}
	upperChild, upperOk := ct.edges.Get(protocol.EdgeId(curr.upperChildId))
	if !upperOk {
		panic(fmt.Sprintf("not upper curr %s, upper=%s", curr.id, curr.upperChildId.Bytes()))
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

func edgeLength(eg *edge) uint64 {
	return eg.endHeight - eg.startHeight
}

func isChild(eg *edge, childId protocol.EdgeId) bool {
	return eg.lowerChildId == common.Hash(childId) || eg.upperChildId == common.Hash(childId)
}

func hasChildren(eg *edge) bool {
	return eg.lowerChildId != (common.Hash{}) || eg.upperChildId != (common.Hash{})
}

func reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
