// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE

package challengetree

import (
	"context"
	"fmt"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/OffchainLabs/bold/containers"
	"github.com/pkg/errors"
)

type ancestorsQueryResponse struct {
	ancestorCumulativePathTimers []PathTimer
	ancestorEdgeIds              HonestAncestors
}

// computeAncestorsWithTimers computes the ancestors of the given edge and their respective path timers, even
// across challenge levels. Ancestor lists are linked through challenge levels via claimed edges. It is generalized
// to any number of challenge levels in the protocol.
func (ht *HonestChallengeTree) computeAncestorsWithTimers(
	ctx context.Context,
	edgeId protocol.EdgeId,
	blockNumber uint64,
) (*ancestorsQueryResponse, error) {
	startEdge, ok := ht.edges.TryGet(edgeId)
	if !ok {
		return nil, errNotFound(edgeId)
	}
	if ht.honestBlockChalLevelZeroEdge.IsNone() {
		return nil, ErrNoHonestTopLevelEdge
	}
	currentChallengeLevel, err := startEdge.GetChallengeLevel()
	if err != nil {
		return nil, err
	}
	currentEdge := startEdge

	ancestry := make([]protocol.EdgeId, 0)

	for currentChallengeLevel < protocol.ChallengeLevel(ht.totalChallengeLevels) {
		// Compute the root edge for the current challenge level.
		rootEdge, err := ht.honestRootAncestorAtChallengeLevel(currentEdge, currentChallengeLevel)
		if err != nil {
			return nil, err
		}

		// Compute the ancestors for the current edge in the current challenge level.
		_, ancestorsAtLevel, err := ht.findAncestorsInChallenge(ctx, rootEdge, currentEdge, blockNumber)
		if err != nil {
			return nil, err
		}

		// Advance the challenge level.
		currentChallengeLevel += 1

		// Expand the total ancestry list. We want ancestors from
		// the bottom-up, so we must reverse the output slice
		// from the find function.
		containers.Reverse(ancestorsAtLevel)
		ancestry = append(ancestry, ancestorsAtLevel...)

		if currentChallengeLevel == protocol.ChallengeLevel(ht.totalChallengeLevels) {
			break
		}

		// Update the current edge to the one the root edge at this challenge claims
		// at the next challenge level to link between levels.
		nextLevelClaimedEdge, err := ht.getClaimedEdge(rootEdge)
		if err != nil {
			return nil, err
		}
		currentEdge = nextLevelClaimedEdge

		// Include the next level claimed edge in the ancestry list.
		ancestry = append(ancestry, nextLevelClaimedEdge.Id())
	}

	// If the ancestry is empty, we just return an empty response.
	if len(ancestry) == 0 {
		return &ancestorsQueryResponse{
			ancestorCumulativePathTimers: make([]PathTimer, 0),
			ancestorEdgeIds:              ancestry,
		}, nil
	}

	// If the ancestry list is non-empty, the last edge in the ancestry should
	// be the honest block challenge level root edge we agree with. We perform this
	// safety check at the end of this function to ensure we are returning
	// a proper ancestry list.
	if ht.honestBlockChalLevelZeroEdge.IsNone() {
		// Should never happen, just an extra check against panics.
		return nil, errors.New("no honest block challenge root edge found")
	}
	rootChallengeEdgeId := ht.honestBlockChalLevelZeroEdge.Unwrap().Id()
	lastAncestryEdgeId := ancestry[len(ancestry)-1]
	if rootChallengeEdgeId != lastAncestryEdgeId {
		return nil, fmt.Errorf(
			"last edge in ancestry %#x is not the top-level, root honest edge %#x",
			lastAncestryEdgeId,
			rootChallengeEdgeId,
		)
	}

	// In addition, each path timer returned should also take into account
	// the assertion unrivaled timer from the assertion chain. To do this, we compute it here and add it to
	// each entry path timer before we return it.
	assertionUnrivaledTimer, err := ht.metadataReader.AssertionUnrivaledBlocks(
		ctx, ht.topLevelAssertionHash,
	)
	if err != nil {
		return nil, err
	}
	ancestorCumulativeTimers := make([]PathTimer, 0)
	for i := range ancestorCumulativeTimers {
		ancestorCumulativeTimers[i] += PathTimer(assertionUnrivaledTimer)
	}

	return &ancestorsQueryResponse{
		ancestorCumulativePathTimers: ancestorCumulativeTimers,
		ancestorEdgeIds:              ancestry,
	}, nil
}

// Computes the root edge for a given child edge at a challenge level.
func (ht *HonestChallengeTree) honestRootAncestorAtChallengeLevel(
	childEdge protocol.ReadOnlyEdge,
	challengeLevel protocol.ChallengeLevel,
) (protocol.ReadOnlyEdge, error) {
	originId := childEdge.OriginId()
	if challengeLevel == protocol.ChallengeLevel(ht.totalChallengeLevels)-1 {
		if ht.honestBlockChalLevelZeroEdge.IsNone() {
			return nil, errNoLevelZero(originId)
		}
		return ht.honestBlockChalLevelZeroEdge.Unwrap(), nil
	}
	rootEdgesAtLevel, ok := ht.honestRootEdgesByLevel.TryGet(challengeLevel)
	if !ok {
		return nil, fmt.Errorf("no honest edges found at challenge level %d", challengeLevel)
	}
	if rootEdgesAtLevel == nil {
		return nil, fmt.Errorf("no honest edges found at challenge level %d", challengeLevel)
	}
	rootAncestor, found := findOriginEdge(originId, rootEdgesAtLevel)
	if !found {
		return nil, fmt.Errorf("no honest root edge with origin id %#x found at challenge level %d", originId, challengeLevel)
	}
	return rootAncestor, nil
}
