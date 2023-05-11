package challengetree

import (
	"context"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/OffchainLabs/challenge-protocol-v2/util/threadsafe"
	"github.com/pkg/errors"
)

type MetadataReader interface {
	TopLevelAssertion(ctx context.Context, edgeId protocol.EdgeId) (protocol.AssertionId, error)
	ClaimHeights(ctx context.Context, edgeId protocol.EdgeId) (*ClaimHeights, error)
}

type ClaimHeights struct {
	AssertionClaimHeight      uint64
	BlockChallengeClaimHeight uint64
	BigStepClaimHeight        uint64
}

type Agreement struct {
	IsHonestEdge          bool
	AgreesWithStartCommit bool
}

type HistoryChecker interface {
	AgreesWithHistoryCommitment(
		ctx context.Context,
		heights *ClaimHeights,
		startCommit,
		endCommit util.HistoryCommitment,
	) (Agreement, error)
}

// An honestChallengeTree keeps track of edges the honest node agrees with in a particular challenge.
// All edges tracked in this data structure are part of the same, top-level assertion challenge.
type honestChallengeTree struct {
	edges                            *threadsafe.Map[protocol.EdgeId, protocol.EdgeSnapshot]
	mutualIds                        *threadsafe.Map[protocol.MutualId, *threadsafe.Set[protocol.EdgeId]]
	rivaledEdges                     *threadsafe.Set[protocol.EdgeId]
	topLevelAssertionId              protocol.AssertionId
	honestBlockChalLevelZeroEdge     util.Option[protocol.EdgeSnapshot]
	honestBigStepChalLevelZeroEdge   util.Option[protocol.EdgeSnapshot]
	honestSmallStepChalLevelZeroEdge util.Option[protocol.EdgeSnapshot]
	metadataReader                   MetadataReader
	histChecker                      HistoryChecker
}

func (ct *honestChallengeTree) addEdge(ctx context.Context, eg protocol.EdgeSnapshot) error {
	prevAssertionId, err := ct.metadataReader.TopLevelAssertion(ctx, eg.Id())
	if err != nil {
		return errors.Wrapf(err, "could not get top level assertion for edge %#x", eg.Id())
	}
	if ct.topLevelAssertionId != prevAssertionId {
		// Do nothing - this edge should not be part of this challenge tree.
		return nil
	}

	// We only track edges we fully agree with (honest edges).
	startHeight, startCommit := eg.StartCommitment()
	endHeight, endCommit := eg.EndCommitment()
	heights, err := ct.metadataReader.ClaimHeights(ctx, eg.Id())
	if err != nil {
		return errors.Wrapf(err, "could not get claim heights for edge %#x", eg.Id())
	}
	agreement, err := ct.histChecker.AgreesWithHistoryCommitment(
		ctx,
		heights,
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
		return errors.Wrapf(err, "could not check if agrees with history commit for edge %#x", eg.Id())
	}

	// If we agree with the edge, we add it to our edges mapping and if it is level zero,
	// we keep track of it specifically in our struct.
	if agreement.IsHonestEdge {
		ct.edges.Put(eg.Id(), eg)
		if !eg.ClaimId().IsNone() {
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

	// Check if the edge id should be added to the rivaled edges set.
	// Here we only care about edges here th
	if agreement.AgreesWithStartCommit || agreement.IsHonestEdge {
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
	}
	return nil
}
