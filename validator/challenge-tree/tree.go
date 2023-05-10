package challengetree

import (
	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/OffchainLabs/challenge-protocol-v2/util/threadsafe"
)

type metadataReader interface {
	TopLevelAssertion(edgeId protocol.EdgeId) (protocol.AssertionId, error)
	ClaimHeights(edgeId protocol.EdgeId) *ClaimHeights
}

type ClaimHeights struct {
	AssertionClaimHeight      uint64
	BlockChallengeClaimHeight uint64
	BigStepClaimHeight        uint64
}

type HistoryChecker interface {
	AgreesWithHistoryCommitment(
		heights *ClaimHeights,
		startCommit,
		endCommit util.HistoryCommitment,
	) bool
}

// A challenge tree keeps track of edges the honest node agrees with in a particular challenge.
// All edges tracked in this data structure are part of the same, top-level assertion challenge.
type challengeTree struct {
	edges                        *threadsafe.Map[protocol.EdgeId, protocol.EdgeSnapshot]
	mutualIds                    *threadsafe.Map[protocol.MutualId, *threadsafe.Set[protocol.EdgeId]]
	rivaledEdges                 *threadsafe.Set[protocol.EdgeId]
	topLevelAssertionId          protocol.AssertionId
	honestBlockChalLevelZeroEdge util.Option[protocol.SpecEdge]
	metadataReader               metadataReader
	histChecker                  HistoryChecker
}

func (ct *challengeTree) addEdge(eg protocol.EdgeSnapshot) {
	prevAssertionId, err := ct.metadataReader.TopLevelAssertion(eg.Id())
	if err != nil {
		panic(err)
	}
	if ct.topLevelAssertionId != prevAssertionId {
		// Do nothing - this edge should not be part of this challenge tree.
		return
	}

	// Check if the edge id should be added to the rivaled edges set.
	mutualId := eg.MutualId()
	mutuals := ct.mutualIds.Get(mutualId)
	if mutuals == nil {
		ct.mutualIds.Put(mutualId, threadsafe.NewSet[protocol.EdgeId]())
		mutuals = ct.mutualIds.Get(mutualId)
	}
	if mutuals.Has(eg.Id()) {
		ct.rivaledEdges.Insert(eg.Id())
	} else {
		mutuals.Insert(eg.Id())
	}

	// We only track edges we fully agree with (honest edges).
	startHeight, startCommit := eg.StartCommitment()
	endHeight, endCommit := eg.EndCommitment()
	agrees, err := ct.histChecker.AgreesWithHistoryCommitment(
		ct.metadataReader.ClaimHeights(eg.Id()),
		util.HistoryCommitment{
			Height: uint64(startHeight),
			Merkle: startCommit,
		},
		util.HistoryCommitment{
			Height: uint64(endHeight),
			Merkle: endCommit,
		},
	)
	if err != nil {
		panic(err)
	}
	if agrees {
		ct.edges.Put(eg.Id(), eg)
		if eg.claimId != "" {
			switch eg.GetType() {
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
