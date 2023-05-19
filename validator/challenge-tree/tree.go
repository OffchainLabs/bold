package challengetree

import (
	"context"

	"fmt"
	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/state-manager"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/OffchainLabs/challenge-protocol-v2/util/threadsafe"
	"github.com/pkg/errors"
)

// MetadataReader can read certain information about edges from the backend.
type MetadataReader interface {
	AssertionUnrivaledTime(ctx context.Context, edgeId protocol.EdgeId) (uint64, error)
	TopLevelAssertion(ctx context.Context, edgeId protocol.EdgeId) (protocol.AssertionId, error)
	TopLevelClaimHeights(ctx context.Context, edgeId protocol.EdgeId) (*protocol.OriginHeights, error)
	SpecChallengeManager(ctx context.Context) (protocol.SpecChallengeManager, error)
	GetAssertionNum(ctx context.Context, assertionHash protocol.AssertionId) (protocol.AssertionSequenceNumber, error)
	ReadAssertionCreationInfo(
		ctx context.Context, seqNum protocol.AssertionSequenceNumber,
	) (*protocol.AssertionCreatedInfo, error)
}

type creationTime uint64

// An honestChallengeTree keeps track of edges the honest node agrees with in a particular challenge.
// All edges tracked in this data structure are part of the same, top-level assertion challenge.
type HonestChallengeTree struct {
	edges                         *threadsafe.Map[protocol.EdgeId, protocol.EdgeSnapshot]
	mutualIds                     *threadsafe.Map[protocol.MutualId, *threadsafe.Map[protocol.EdgeId, creationTime]]
	topLevelAssertionId           protocol.AssertionId
	honestBlockChalLevelZeroEdge  util.Option[protocol.EdgeSnapshot]
	honestBigStepLevelZeroEdges   *threadsafe.Slice[protocol.EdgeSnapshot]
	honestSmallStepLevelZeroEdges *threadsafe.Slice[protocol.EdgeSnapshot]
	metadataReader                MetadataReader
	histChecker                   statemanager.HistoryChecker
}

func New(
	prevAssertionId protocol.AssertionId,
	metadataReader MetadataReader,
	histChecker statemanager.HistoryChecker,
) *HonestChallengeTree {
	return &HonestChallengeTree{
		edges:                         threadsafe.NewMap[protocol.EdgeId, protocol.EdgeSnapshot](),
		mutualIds:                     threadsafe.NewMap[protocol.MutualId, *threadsafe.Map[protocol.EdgeId, creationTime]](),
		topLevelAssertionId:           prevAssertionId,
		honestBlockChalLevelZeroEdge:  util.None[protocol.EdgeSnapshot](),
		honestBigStepLevelZeroEdges:   threadsafe.NewSlice[protocol.EdgeSnapshot](),
		honestSmallStepLevelZeroEdges: threadsafe.NewSlice[protocol.EdgeSnapshot](),
		metadataReader:                metadataReader,
		histChecker:                   histChecker,
	}
}

// RefreshEdgesFromChain refreshes all edge snapshots from the chain.
func (ht *HonestChallengeTree) RefreshEdgesFromChain(ctx context.Context) error {
	edgeIds := make([]protocol.EdgeId, 0, ht.edges.NumItems())
	if err := ht.edges.ForEach(func(id protocol.EdgeId, _ protocol.EdgeSnapshot) error {
		edgeIds = append(edgeIds, id)
		return nil
	}); err != nil {
		return err
	}
	edgeReader, err := ht.metadataReader.SpecChallengeManager(ctx)
	if err != nil {
		return err
	}
	snapshots := make([]protocol.EdgeSnapshot, len(edgeIds))
	for i, edgeId := range edgeIds {
		edgeOpt, err := edgeReader.GetEdge(ctx, edgeId)
		if err != nil {
			return err
		}
		if edgeOpt.IsNone() {
			return fmt.Errorf("edge with id %#x not found", edgeId)
		}
		snapshots[i] = edgeOpt.Unwrap()
	}
	for i, edgeId := range edgeIds {
		ht.edges.Put(edgeId, snapshots[i])
	}
	return nil
}

// AddEdge to the honest challenge tree. Only honest edges are tracked, but we also keep track
// of rival ids in a mutual ids mapping internally for extra book-keeping.
func (ht *HonestChallengeTree) AddEdge(ctx context.Context, eg protocol.EdgeSnapshot) error {
	prevAssertionId, err := ht.metadataReader.TopLevelAssertion(ctx, eg.Id())
	if err != nil {
		return errors.Wrapf(err, "could not get top level assertion for edge %#x", eg.Id())
	}
	if ht.topLevelAssertionId != prevAssertionId {
		// Do nothing - this edge should not be part of this challenge tree.
		return nil
	}
	prevAssertionSeqNum, err := ht.metadataReader.GetAssertionNum(ctx, prevAssertionId)
	if err != nil {
		return err
	}
	prevCreationInfo, err := ht.metadataReader.ReadAssertionCreationInfo(ctx, prevAssertionSeqNum)
	if err != nil {
		return err
	}

	// We only track edges we fully agree with (honest edges).
	startHeight, startCommit := eg.StartCommitment()
	endHeight, endCommit := eg.EndCommitment()
	heights, err := ht.metadataReader.TopLevelClaimHeights(ctx, eg.Id())
	if err != nil {
		return errors.Wrapf(err, "could not get claim heights for edge %#x", eg.Id())
	}
	agreement, err := ht.histChecker.AgreesWithHistoryCommitment(
		ctx,
		eg.GetType(),
		prevCreationInfo.InboxMaxCount.Uint64(),
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
		ht.edges.Put(eg.Id(), eg)
		if !eg.ClaimId().IsNone() {
			switch eg.GetType() {
			case protocol.BlockChallengeEdge:
				ht.honestBlockChalLevelZeroEdge = util.Some(eg)
			case protocol.BigStepChallengeEdge:
				ht.honestBigStepLevelZeroEdges.Push(eg)
			case protocol.SmallStepChallengeEdge:
				ht.honestSmallStepLevelZeroEdges.Push(eg)
			default:
			}
		}
	}

	// Check if the edge id should be added to the rivaled edges set.
	// Here we only care about edges here that are either honest or those whose start
	// history commitments we agree with.
	if agreement.AgreesWithStartCommit || agreement.IsHonestEdge {
		mutualId := eg.MutualId()
		mutuals := ht.mutualIds.Get(mutualId)
		if mutuals == nil {
			ht.mutualIds.Put(mutualId, threadsafe.NewMap[protocol.EdgeId, creationTime]())
			mutuals = ht.mutualIds.Get(mutualId)
		}
		mutuals.Put(eg.Id(), creationTime(eg.CreatedAtBlock()))
	}
	return nil
}
