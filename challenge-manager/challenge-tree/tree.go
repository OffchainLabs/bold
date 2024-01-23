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

// RoyalChallengeTree keeps track of royal edges the honest node agrees with in a particular challenge.
// All edges tracked in this data structure are part of the same, top-level assertion challenge.
type RoyalChallengeTree struct {
	edges                 *threadsafe.Map[protocol.EdgeId, protocol.SpecEdge]
	mutualIds             *threadsafe.Map[protocol.MutualId, *threadsafe.Map[protocol.EdgeId, creationTime]]
	topLevelAssertionHash protocol.AssertionHash
	metadataReader        MetadataReader
	histChecker           l2stateprovider.HistoryChecker
	validatorName         string
	totalChallengeLevels  uint8
	royalRootEdgesByLevel *threadsafe.Map[protocol.ChallengeLevel, *threadsafe.Slice[protocol.ReadOnlyEdge]]
}

func New(
	assertionHash protocol.AssertionHash,
	metadataReader MetadataReader,
	histChecker l2stateprovider.HistoryChecker,
	numBigStepLevels uint8,
	validatorName string,
) *RoyalChallengeTree {
	return &RoyalChallengeTree{
		edges:                 threadsafe.NewMap[protocol.EdgeId, protocol.SpecEdge](),
		mutualIds:             threadsafe.NewMap[protocol.MutualId, *threadsafe.Map[protocol.EdgeId, creationTime]](),
		topLevelAssertionHash: assertionHash,
		metadataReader:        metadataReader,
		histChecker:           histChecker,
		validatorName:         validatorName,
		// The total number of challenge levels include block challenges, small step challenges, and N big step challenges.
		totalChallengeLevels:  numBigStepLevels + 2,
		royalRootEdgesByLevel: threadsafe.NewMap[protocol.ChallengeLevel, *threadsafe.Slice[protocol.ReadOnlyEdge]](),
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
