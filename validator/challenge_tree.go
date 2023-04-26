package validator

import (
	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/ethereum/go-ethereum/common"
)

type set[T comparable] struct {
	items map[T]bool
}

func newSet[T comparable]() *set[T] {
	return &set[T]{
		items: make(map[T]bool),
	}
}

func (s *set[T]) insert(t T) {
	s.items[t] = true
}

func (s *set[T]) has(t T) bool {
	return s.items[t]
}

func (s *set[T]) allItems() map[T]bool {
	return s.items
}

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
}

type chain interface {
	timeUnrivaled(edgeId protocol.EdgeId) uint64
}

type historyChecker interface {
	hasHistoryCommitment(
		startHeight uint64,
		startCommit common.Hash,
		endHeight uint64,
		endCommit common.Hash,
	) bool
}

type edgeMetadataReader interface {
	topLevelAssertion(protocol.EdgeId) protocol.AssertionId
	mutualId(protocol.EdgeId) protocol.MutualId
}

// A challenge tree keeps track of the honest branch for a challenged
// assertion in the protocol. All edges tracked in this data structure
// are part of the same challenge.
type challengeTree struct {
	// TODO: Needs to be thread-safe.
	topLevelAssertionId              protocol.AssertionId
	chain                            chain
	metadataReader                   edgeMetadataReader
	histChecker                      historyChecker
	edges                            map[protocol.EdgeId]*edge
	mutualIds                        map[protocol.MutualId]set[protocol.EdgeId]
	rivaledEdges                     set[protocol.EdgeId]
	honestEdges                      set[protocol.EdgeId]
	honestBranch                     []protocol.EdgeId
	honestUnrivaledCumulativeTimers  map[protocol.EdgeId]uint64
	honestBlockChalLevelZeroEdge     *edge
	honestBigStepChalLevelZeroEdge   *edge
	honestSmallStepChalLevelZeroEdge *edge
}

func (ct *challengeTree) addEdge(eg *edge) {
	prevAssertionId := ct.metadataReader.topLevelAssertion(eg.id)
	if ct.topLevelAssertionId != prevAssertionId {
		// Do nothing - this edge should not be part of this challenge tree.
		return
	}

	// Check if the edge id should be added to the rivaled edges set.
	mutualId := ct.metadataReader.mutualId(eg.id)
	if mutuals, ok := ct.mutualIds[mutualId]; ok {
		if mutuals.has(eg.id) {
			ct.rivaledEdges.insert(eg.id)
		} else {
			mutuals.insert(eg.id)
		}
	}

	// If this is an honest edge from our perspective, we keep track
	// of it as such in a set. TODO: Handle subchallenges.
	if ct.histChecker.hasHistoryCommitment(
		eg.startHeight,
		eg.startCommit,
		eg.endHeight,
		eg.endCommit,
	) {
		ct.honestEdges.insert(eg.id)
		if eg.claimId != (common.Hash{}) {
			switch eg.edgeType {
			case protocol.BlockChallengeEdge:
				ct.honestBlockChalLevelZeroEdge = eg
			case protocol.BigStepChallengeEdge:
				ct.honestBigStepChalLevelZeroEdge = eg
			case protocol.SmallStepChallengeEdge:
				ct.honestSmallStepChalLevelZeroEdge = eg
			default:
			}
		}
	}

	// Add the edge to the map of edge ids for the challenge.
	ct.edges[eg.id] = eg
}

// Get the honest level zero edge from our list of honest
// edges (keep track of them per challenge level).
// Recursively go down its children and then update their
// cumulative timers accordingly.
func (ct *challengeTree) updateCumulativeTimers() {
	blockEdge := ct.honestBlockChalLevelZeroEdge
	ct.innerCumulativeUpdate(0, blockEdge.id)
	// TODO: Figure out how to do for the lower challenge levels.
}

func (ct *challengeTree) innerCumulativeUpdate(
	cumulativeUnrivaledTime uint64,
	edgeId protocol.EdgeId,
) {
	edge := ct.edges[edgeId]
	blocksUnrivaled := ct.chain.timeUnrivaled(edgeId)
	total := blocksUnrivaled + cumulativeUnrivaledTime
	ct.honestUnrivaledCumulativeTimers[edgeId] = total
	if edge.lowerChildId != (common.Hash{}) {
		ct.innerCumulativeUpdate(total, protocol.EdgeId(edge.lowerChildId))
	}
	if edge.upperChildId != (common.Hash{}) {
		ct.innerCumulativeUpdate(total, protocol.EdgeId(edge.upperChildId))
	}
}

// TODO: How to get ancestors list based on our construct?
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
	// The deviant edge 0-8' bisects, but agrees with the honest one from 0-4.
	// Therefore, there is only a single 0-4 edge in the set.
	//
	// In this case, the challenge tree's list of honest edge ids will be:
	//
	//   [id(0,4), id(4,6), id(6,8), id(0,8)]
	//
	// from here, the ancestor list can be determined as we only care about direct ancestors
	// of an honest edge, and the evil edges would not be considered. For example, if
	// we want the ancestors for id(6,8), they would be id(4,8), and id(0,8).
	//
	// In order to retrieve ancestors for an edge with id=I, we start from the honest,
	// block challenge level zero edge and recursively traverse its children,
	// reducing the ids along the way into a slice until we hit a child that
	// matches id=I and return the slice.
	return ct.ancestorQuery(
		make([]protocol.EdgeId, 0),
		ct.honestBlockChalLevelZeroEdge,
		id,
	)
}

func (ct *challengeTree) ancestorQuery(
	accum []protocol.EdgeId,
	curr *edge,
	queryingFor protocol.EdgeId,
) []protocol.EdgeId {
	if curr.lowerChildId == (common.Hash{}) && curr.upperChildId == (common.Hash{}) {
		return accum
	}
	if curr.lowerChildId == common.Hash(queryingFor) || curr.upperChildId == common.Hash(queryingFor) {
		return append(accum, curr.id)
	}
	lowerChild := ct.edges[protocol.EdgeId(curr.lowerChildId)]
	upperChild := ct.edges[protocol.EdgeId(curr.upperChildId)]
	lowerAncestors := ct.ancestorQuery(append(accum, curr.id), lowerChild, queryingFor)
	upperAncestors := ct.ancestorQuery(append(accum, curr.id), upperChild, queryingFor)
	allAncestors := accum
	allAncestors = append(allAncestors, lowerAncestors...)
	allAncestors = append(allAncestors, upperAncestors...)
	return allAncestors
}
