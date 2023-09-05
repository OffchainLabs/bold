// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE

package challengetree

import (
	"context"
	"fmt"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
)

type findAncestorsRequest struct {
	ctx                   context.Context
	rootEdge              protocol.ReadOnlyEdge
	queryingFor           protocol.ReadOnlyEdge
	totalChallengeLevels  uint8
	currentChallengeLevel uint8
	blockNumber           uint64
}

type ancestorsQueryResponse struct {
	ancestorCumulativePathTimers []PathTimer
	ancestorEdgeIds              HonestAncestors
}

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

		// Expand the total ancestry list.
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

	return &ancestorsQueryResponse{
		ancestorCumulativePathTimers: make([]PathTimer, 0),
		ancestorEdgeIds:              ancestry,
	}, nil
}

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
