// Package challengetree includes logic for keeping track of royal edges within a challenge
// with utilities for computing cumulative path timers for said edges. This is helpful during
// the confirmation process needed by edge trackers.
//
// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE
package challengetree

import (
	"context"
	"fmt"
	"math"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/OffchainLabs/bold/containers/threadsafe"
	l2stateprovider "github.com/OffchainLabs/bold/layer2-state-provider"
	"github.com/pkg/errors"
)

// MetadataReader can read certain information about edges from the backend.
type MetadataReader interface {
	AssertionUnrivaledBlocks(ctx context.Context, assertionHash protocol.AssertionHash) (uint64, error)
	TopLevelAssertion(ctx context.Context, edgeId protocol.EdgeId) (protocol.AssertionHash, error)
	TopLevelClaimHeights(ctx context.Context, edgeId protocol.EdgeId) (protocol.OriginHeights, error)
	SpecChallengeManager(ctx context.Context) (protocol.SpecChallengeManager, error)
	ReadAssertionCreationInfo(
		ctx context.Context, id protocol.AssertionHash,
	) (*protocol.AssertionCreatedInfo, error)
}

type creationTime uint64

// OriginPlusMutualId combines a mutual id and origin id as a key for a mapping.
// This is used for computing the rivals of an edge, as all rivals share a mutual id.
// However, we also add the origin id as that allows us to namespace to lookup
// to a specific challenge.
type OriginPlusMutualId [64]byte

func buildEdgeCreationTimeKey(originId protocol.OriginId, mutualId protocol.MutualId) OriginPlusMutualId {
	var key OriginPlusMutualId
	copy(key[0:32], originId[:])
	copy(key[32:64], mutualId[:])
	return key
}

// RoyalChallengeTree keeps track of royal edges the honest node agrees with in a particular challenge.
// All edges tracked in this data structure are part of the same, top-level assertion challenge.
type RoyalChallengeTree struct {
	edges                         *threadsafe.Map[protocol.EdgeId, protocol.SpecEdge]
	edgeInheritedTimers           *threadsafe.Map[protocol.EdgeId, uint64]
	claimedEdgesWithUpdatedTimers *threadsafe.Set[protocol.EdgeId]
	edgeCreationTimes             *threadsafe.Map[OriginPlusMutualId, *threadsafe.Map[protocol.EdgeId, creationTime]]
	topLevelAssertionHash         protocol.AssertionHash
	metadataReader                MetadataReader
	histChecker                   l2stateprovider.HistoryChecker
	validatorName                 string
	totalChallengeLevels          uint8
	royalRootEdgesByLevel         *threadsafe.Map[protocol.ChallengeLevel, *threadsafe.Slice[protocol.ReadOnlyEdge]]
}

func New(
	assertionHash protocol.AssertionHash,
	metadataReader MetadataReader,
	histChecker l2stateprovider.HistoryChecker,
	numBigStepLevels uint8,
	validatorName string,
) *RoyalChallengeTree {
	return &RoyalChallengeTree{
		edges:                         threadsafe.NewMap[protocol.EdgeId, protocol.SpecEdge](threadsafe.MapWithMetric[protocol.EdgeId, protocol.SpecEdge]("edges")),
		edgeInheritedTimers:           threadsafe.NewMap[protocol.EdgeId, uint64](),
		claimedEdgesWithUpdatedTimers: threadsafe.NewSet[protocol.EdgeId](),
		edgeCreationTimes:             threadsafe.NewMap[OriginPlusMutualId, *threadsafe.Map[protocol.EdgeId, creationTime]](threadsafe.MapWithMetric[OriginPlusMutualId, *threadsafe.Map[protocol.EdgeId, creationTime]]("edgeCreationTimes")),
		topLevelAssertionHash:         assertionHash,
		metadataReader:                metadataReader,
		histChecker:                   histChecker,
		validatorName:                 validatorName,
		// The total number of challenge levels include block challenges, small step challenges, and N big step challenges.
		totalChallengeLevels:  numBigStepLevels + 2,
		royalRootEdgesByLevel: threadsafe.NewMap[protocol.ChallengeLevel, *threadsafe.Slice[protocol.ReadOnlyEdge]](threadsafe.MapWithMetric[protocol.ChallengeLevel, *threadsafe.Slice[protocol.ReadOnlyEdge]]("royalRootEdgesByLevel")),
	}
}

// RoyalBlockChallengeRootEdge gets the royal, root challenge block edge for the top level assertion
// being challenged.
func (ht *RoyalChallengeTree) RoyalBlockChallengeRootEdge() (protocol.ReadOnlyEdge, error) {
	// In our locally tracked challenge tree implementation, the
	// block challenge level is equal to the total challenge levels - 1.
	blockChallengeLevel := protocol.ChallengeLevel(ht.totalChallengeLevels) - 1
	if rootEdges, ok := ht.royalRootEdgesByLevel.TryGet(blockChallengeLevel); ok {
		if rootEdges.Len() != 1 {
			return nil, fmt.Errorf(
				"expected one royal root block challenge edge for challenged assertion %#x",
				ht.topLevelAssertionHash,
			)
		}
		return rootEdges.Get(0).Unwrap(), nil
	}
	return nil, fmt.Errorf("no royal root edges for block challenge level for assertion %#x", ht.topLevelAssertionHash)
}

var (
	ErrAlreadyBeingTracked              = errors.New("edge already being tracked")
	ErrMismatchedChallengeAssertionHash = errors.New("edge challenged assertion hash is not the expected one for the challenge")
)

func (ht *RoyalChallengeTree) GetEdges() *threadsafe.Map[protocol.EdgeId, protocol.SpecEdge] {
	return ht.edges
}

func (ht *RoyalChallengeTree) HasRoyalEdge(edgeId protocol.EdgeId) bool {
	return ht.edges.Has(edgeId)
}

func (ht *RoyalChallengeTree) IsUnrivaledAtBlockNum(edge protocol.ReadOnlyEdge, blockNum uint64) (bool, error) {
	return ht.UnrivaledAtBlockNum(edge, blockNum)
}

func (ht *RoyalChallengeTree) TimeUnrivaled(edge protocol.ReadOnlyEdge, blockNum uint64) (uint64, error) {
	return ht.LocalTimer(edge, blockNum)
}

func (ht *RoyalChallengeTree) UpdateInheritedTimer(
	ctx context.Context,
	edgeId protocol.EdgeId,
	blockNum uint64,
) (uint64, error) {
	edge, ok := ht.edges.TryGet(edgeId)
	if !ok {
		return 0, fmt.Errorf("edge with id %#x not found", edgeId.Hash)
	}
	timeUnrivaled, err := ht.TimeUnrivaled(edge, blockNum)
	if err != nil {
		return 0, err
	}
	inheritedTimer := timeUnrivaled

	// Subchallenged edges always have their inherited timer fetched from onchain.
	// TODO: When to perform the tx? As long as we have called it once, we don't need to do it again.

	// If an edge has children, we use the min of its children if it
	// is not a root edge. Otherwise, use the max.
	hasChildren, err := edge.HasChildren(ctx)
	if err != nil {
		return 0, err
	}
	chalManager, err := ht.metadataReader.SpecChallengeManager(ctx)
	if err != nil {
		return 0, err
	}
	if hasChildren {
		lowerChildOpt, err := edge.LowerChild(ctx)
		if err != nil {
			return 0, err
		}
		upperChildOpt, err := edge.UpperChild(ctx)
		if err != nil {
			return 0, err
		}
		lowerChildId := lowerChildOpt.Unwrap()
		upperChildId := upperChildOpt.Unwrap()

		// We get the inherited timers of the children. If we don't have
		// them locally, we fetch them from onchain.
		lowerTimer, ok := ht.edgeInheritedTimers.TryGet(lowerChildId)
		if !ok {
			lowerInheritedTimer, err := chalManager.InheritedTimer(ctx, lowerChildId)
			if err != nil {
				return 0, err
			}
			lowerTimer = lowerInheritedTimer
			ht.edgeInheritedTimers.Put(lowerChildId, lowerInheritedTimer)
		}
		upperTimer, ok := ht.edgeInheritedTimers.TryGet(upperChildId)
		if !ok {
			upperInheritedTimer, err := chalManager.InheritedTimer(ctx, upperChildId)
			if err != nil {
				return 0, err
			}
			upperTimer = upperInheritedTimer
			ht.edgeInheritedTimers.Put(upperChildId, upperInheritedTimer)
		}

		if edge.ClaimId().IsSome() {
			val := lowerTimer
			if upperTimer > lowerTimer {
				val = upperTimer
			}
			inheritedTimer += val
		} else {
			val := lowerTimer
			if upperTimer < lowerTimer {
				val = upperTimer
			}
			inheritedTimer += val
		}
		ht.edgeInheritedTimers.Put(edgeId, inheritedTimer)
		return inheritedTimer, nil
	}

	// If the edge has been confirmed, is one-step, and is small step,
	// then its timer will be set to max uint64.
	status, err := edge.Status(ctx)
	if err != nil {
		return 0, err
	}
	startHeight, _ := edge.StartCommitment()
	endHeight, _ := edge.StartCommitment()
	diff := endHeight - startHeight
	if edge.GetChallengeLevel() == protocol.ChallengeLevel(edge.GetTotalChallengeLevels(ctx)-1) && status == protocol.EdgeConfirmed && diff == 1 {
		inheritedTimer = math.MaxUint64
		ht.edgeInheritedTimers.Put(edgeId, inheritedTimer)
		return inheritedTimer, nil
	}

	// Otherwise, the edge does not yet have children.
	ht.edgeInheritedTimers.Put(edgeId, inheritedTimer)
	return inheritedTimer, nil
}
