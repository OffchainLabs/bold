// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE

package challengetree

import (
	"context"
	"fmt"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/OffchainLabs/bold/containers"
	bisection "github.com/OffchainLabs/bold/math"
	"github.com/pkg/errors"
)

// EdgeLocalTimer is the local unrivaled timer of a specific edge (not a cumulative path timer).
type EdgeLocalTimer uint64

type ancestorsQueryResponse struct {
	ancestorLocalTimers []EdgeLocalTimer
	ancestorEdgeIds     HonestAncestors
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
	localTimers := make([]EdgeLocalTimer, 0)

	for currentChallengeLevel < protocol.ChallengeLevel(ht.totalChallengeLevels) {
		// Compute the root edge for the current challenge level.
		rootEdge, err := ht.honestRootAncestorAtChallengeLevel(currentEdge, currentChallengeLevel)
		if err != nil {
			return nil, err
		}

		// Compute the ancestors for the current edge in the current challenge level.
		ancestorLocalTimers, ancestorsAtLevel, err := ht.findHonestAncestorsWithinChallengeLevel(ctx, rootEdge, currentEdge, blockNumber)
		if err != nil {
			return nil, err
		}

		// Advance the challenge level.
		currentChallengeLevel += 1

		// Expand the total ancestry list. We want ancestors from
		// the bottom-up, so we must reverse the output slice from the find function.
		containers.Reverse(ancestorLocalTimers)
		containers.Reverse(ancestorsAtLevel)
		ancestry = append(ancestry, ancestorsAtLevel...)
		localTimers = append(localTimers, ancestorLocalTimers...)

		if currentChallengeLevel == protocol.ChallengeLevel(ht.totalChallengeLevels) {
			break
		}

		// Update the current edge to the one the root edge at this challenge claims
		// at the next challenge level to link between levels.
		nextLevelClaimedEdge, err := ht.getClaimedEdge(rootEdge)
		if err != nil {
			return nil, err
		}
		claimEdgeLocalTimer, err := ht.localTimer(nextLevelClaimedEdge, blockNumber)
		if err != nil {
			return nil, err
		}
		currentEdge = nextLevelClaimedEdge

		// Include the next level claimed edge in the ancestry list.
		ancestry = append(ancestry, nextLevelClaimedEdge.Id())
		localTimers = append(localTimers, EdgeLocalTimer(claimEdgeLocalTimer))
	}

	// If the ancestry is empty, we just return an empty response.
	if len(ancestry) == 0 {
		return &ancestorsQueryResponse{
			ancestorLocalTimers: make([]EdgeLocalTimer, 0),
			ancestorEdgeIds:     ancestry,
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
	return &ancestorsQueryResponse{
		ancestorLocalTimers: localTimers,
		ancestorEdgeIds:     ancestry,
	}, nil
}

// Computes the list of ancestors in a challenge type from a root edge down
// to a specified child edge within a single challenge level. The edge we are querying must be
// a child of this start edge for this function to succeed without error.
func (ht *HonestChallengeTree) findHonestAncestorsWithinChallengeLevel(
	ctx context.Context,
	start protocol.ReadOnlyEdge,
	queryingFor protocol.ReadOnlyEdge,
	blockNumber uint64,
) ([]EdgeLocalTimer, []protocol.EdgeId, error) {
	found := false
	curr := start
	ancestry := make([]protocol.EdgeId, 0)
	localTimers := make([]EdgeLocalTimer, 0)
	wantedEdgeStart, _ := queryingFor.StartCommitment()

	for {
		if ctx.Err() != nil {
			return nil, nil, ctx.Err()
		}
		if curr.Id() == queryingFor.Id() {
			found = true
			break
		}
		ancestry = append(ancestry, curr.Id())

		currStart, _ := curr.StartCommitment()
		currEnd, _ := curr.EndCommitment()
		bisectTo, err := bisection.Bisect(uint64(currStart), uint64(currEnd))
		if err != nil {
			return nil, nil, errors.Wrapf(err, "could not bisect start=%d, end=%d", currStart, currEnd)
		}
		// If the wanted edge's start commitment is < the bisection height of the current
		// edge in the loop, it means it is part of its lower children.
		if uint64(wantedEdgeStart) < bisectTo {
			lowerSnapshot, lowerErr := curr.LowerChild(ctx)
			if lowerErr != nil {
				return nil, nil, errors.Wrapf(lowerErr, "could not get lower child for edge %#x", curr.Id())
			}
			if lowerSnapshot.IsNone() {
				return nil, nil, fmt.Errorf("edge %#x had no lower child", curr.Id())
			}
			curr = ht.edges.Get(lowerSnapshot.Unwrap())
		} else {
			// Else, it is part of the upper children.
			upperSnapshot, upperErr := curr.UpperChild(ctx)
			if upperErr != nil {
				return nil, nil, errors.Wrapf(upperErr, "could not get upper child for edge %#x", curr.Id())
			}
			if upperSnapshot.IsNone() {
				return nil, nil, fmt.Errorf("edge %#x had no upper child", curr.Id())
			}
			curr = ht.edges.Get(upperSnapshot.Unwrap())
		}
		timer, err := ht.localTimer(curr, blockNumber)
		if err != nil {
			return nil, nil, err
		}
		localTimers = append(localTimers, EdgeLocalTimer(timer))
	}
	if !found {
		return nil, nil, errNotFound(queryingFor.Id())
	}
	return localTimers, ancestry, nil
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
