package challengetree

import (
	"fmt"
	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/pkg/errors"
)

// Consider the following set of edges in a challenge where evil
// edges are marked with a ' and a *:
//
//		     /---6---8
//	  0-----4
//		     \---6'--8'
//
// The honest branch is the one that goes from 0-8. The evil edge is 0-8'.
// The evil edge 0-8' bisects, but agrees with the honest one from 0-4.
// Therefore, there is only a single 0-4 edge in the set.
//
// In this case, the set of honest ancestors for 4-6 is the following:
//
//	{4-8, 0-8}
//
// In order to retrieve ancestors for an edge with id=I, we start from the honest,
// block challenge level zero edge and recursively traverse its children,
// reducing the ids and along the way into a slice until we hit a child that
// matches id=I and return the slice.
func (ht *HonestChallengeTree) AncestorsForHonestEdge(queryingFor protocol.EdgeId) ([]protocol.EdgeId, error) {
	wantedEdge, ok := ht.edges.TryGet(queryingFor)
	if !ok {
		return nil, errors.New("not found in honest challenge tree")
	}
	honestLevelZero := ht.honestBlockChalLevelZeroEdge.Unwrap()

	// Figure out what kind of edge this is, and apply different logic based on it.
	var curr protocol.EdgeSnapshot
	switch wantedEdge.GetType() {
	case protocol.BlockChallengeEdge:
		curr = honestLevelZero
		ancestry, err := ht.findAncestorsInChallenge(curr, wantedEdge)
		if err != nil {
			return nil, err
		}
		util.Reverse(ancestry)
		return ancestry, nil
	case protocol.BigStepChallengeEdge:
		ancestry := make([]protocol.EdgeId, 0)
		var bigStepLevelZeroEdge protocol.EdgeSnapshot
		foundLevelZeroEdge := false
		for _, e := range ht.honestBigStepLevelZeroEdges {
			if e.OriginId() == wantedEdge.OriginId() {
				bigStepLevelZeroEdge = e
				foundLevelZeroEdge = true
				break
			}
		}
		if !foundLevelZeroEdge {
			return nil, errors.New("no level zero edge with origin id found")
		}

		start := bigStepLevelZeroEdge
		searchFor := wantedEdge
		bigStepAncestry, err := ht.findAncestorsInChallenge(start, searchFor)
		if err != nil {
			return nil, err
		}
		if bigStepLevelZeroEdge.ClaimId().IsNone() {
			return nil, errors.New("does not claim any edge")
		}
		claimId := bigStepLevelZeroEdge.ClaimId().Unwrap()
		claimedEdge, ok := ht.edges.TryGet(protocol.EdgeId(claimId))
		if !ok {
			return nil, errors.New("claimed edge not found")
		}

		start = honestLevelZero
		searchFor = claimedEdge
		blockChalAncestry, err := ht.findAncestorsInChallenge(start, searchFor)
		if err != nil {
			return nil, err
		}
		ancestry = append(ancestry, blockChalAncestry...)
		ancestry = append(ancestry, claimedEdge.Id())
		ancestry = append(ancestry, bigStepAncestry...)

		util.Reverse(ancestry)
		return ancestry, nil
	case protocol.SmallStepChallengeEdge:
		ancestry := make([]protocol.EdgeId, 0)
		var smallStepLevelZeroEdge protocol.EdgeSnapshot
		foundLevelZeroEdge := false
		for _, e := range ht.honestSmallStepLevelZeroEdges {
			if e.OriginId() == wantedEdge.OriginId() {
				smallStepLevelZeroEdge = e
				foundLevelZeroEdge = true
				break
			}
		}
		if !foundLevelZeroEdge {
			return nil, errors.New("no level zero edge with origin id found")
		}
		start := smallStepLevelZeroEdge
		searchFor := wantedEdge
		smallStepAncestry, err := ht.findAncestorsInChallenge(start, searchFor)
		if err != nil {
			return nil, err
		}
		if smallStepLevelZeroEdge.ClaimId().IsNone() {
			return nil, errors.New("does not claim any edge")
		}
		claimId := smallStepLevelZeroEdge.ClaimId().Unwrap()
		claimedBigStepEdge, ok := ht.edges.TryGet(protocol.EdgeId(claimId))
		if !ok {
			return nil, errors.New("claimed edge not found")
		}
		var bigStepLevelZero protocol.EdgeSnapshot
		foundLevelZeroEdge = false
		for _, e := range ht.honestBigStepLevelZeroEdges {
			if e.OriginId() == claimedBigStepEdge.OriginId() {
				bigStepLevelZero = e
				foundLevelZeroEdge = true
				break
			}
		}
		if !foundLevelZeroEdge {
			return nil, errors.New("no level zero edge with origin id found")
		}

		start = bigStepLevelZero
		searchFor = claimedBigStepEdge
		bigStepAncestry, err := ht.findAncestorsInChallenge(start, searchFor)
		if err != nil {
			return nil, err
		}

		if bigStepLevelZero.ClaimId().IsNone() {
			return nil, errors.New("does not claim any edge")
		}
		claimId = bigStepLevelZero.ClaimId().Unwrap()
		claimedBlockEdge, ok := ht.edges.TryGet(protocol.EdgeId(claimId))
		if !ok {
			return nil, errors.New("claimed edge not found")
		}

		start = honestLevelZero
		searchFor = claimedBlockEdge
		blockAncestry, err := ht.findAncestorsInChallenge(start, searchFor)
		if err != nil {
			return nil, err
		}
		ancestry = append(ancestry, blockAncestry...)
		ancestry = append(ancestry, claimedBlockEdge.Id())
		ancestry = append(ancestry, bigStepAncestry...)
		ancestry = append(ancestry, claimedBigStepEdge.Id())
		ancestry = append(ancestry, smallStepAncestry...)
		util.Reverse(ancestry)
		return ancestry, nil
	default:
		return nil, errors.New("not found")
	}
}

func (ht *HonestChallengeTree) findAncestorsInChallenge(
	start protocol.EdgeSnapshot,
	queryingFor protocol.EdgeSnapshot,
) ([]protocol.EdgeId, error) {
	found := false
	curr := start
	ancestry := make([]protocol.EdgeId, 0)
	wantedEdgeStart, _ := queryingFor.StartCommitment()
	for {
		if curr.Id() == queryingFor.Id() {
			found = true
			break
		}
		ancestry = append(ancestry, curr.Id())

		currStart, _ := curr.StartCommitment()
		currEnd, _ := curr.EndCommitment()
		bisectTo, _ := util.BisectionPoint(uint64(currStart), uint64(currEnd))
		if uint64(wantedEdgeStart) < bisectTo {
			// Lower child..., increase the list of ancestors.
			lowerSnapshot := curr.LowerChildSnapshot()
			if lowerSnapshot.IsNone() {
				return nil, fmt.Errorf("edge %s had no lower child", curr.Id())
			}
			curr = ht.edges.Get(lowerSnapshot.Unwrap())
		} else {
			upperSnapshot := curr.UpperChildSnapshot()
			if upperSnapshot.IsNone() {
				return nil, fmt.Errorf("edge %s had no upper child", curr.Id())
			}
			curr = ht.edges.Get(upperSnapshot.Unwrap())
		}
	}
	if !found {
		return nil, errors.New("not found")
	}
	return ancestry, nil
}

// Checks if an edge is rivaled by looking up its mutual ids mapping.
func (ht *HonestChallengeTree) isRivaled(edge protocol.EdgeSnapshot) bool {
	mutuals, ok := ht.mutualIds.TryGet(edge.MutualId())
	if !ok {
		return false
	}
	// If the mutual ids mapping has more than 1 item and includes
	// the edge, it is then rivaled.
	return mutuals.NumItems() > 1 && mutuals.Has(edge.Id())
}
