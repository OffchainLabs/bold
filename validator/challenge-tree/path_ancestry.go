package challengetree

import (
	"fmt"
	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/pkg/errors"
)

func (ht *HonestChallengeTree) ancestryPath(queryingFor protocol.EdgeId) ([]protocol.EdgeId, error) {
	wantedEdge, ok := ht.edges.TryGet(queryingFor)
	if !ok {
		return nil, errors.New("not found")
	}
	honestLevelZero := ht.honestBlockChalLevelZeroEdge.Unwrap()
	ancestry := make([]protocol.EdgeId, 0)
	_ = ancestry

	// Figure out what kind of edge this is, and apply different logic based on it.
	var curr protocol.EdgeSnapshot
	switch wantedEdge.GetType() {
	case protocol.BlockChallengeEdge:
		curr = honestLevelZero
		blockChalAncestry, err := ht.findAncestorsInChallenge(curr, wantedEdge)
		if err != nil {
			return nil, err
		}
		return blockChalAncestry, nil
	case protocol.BigStepChallengeEdge:
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
		if bigStepLevelZeroEdge.ClaimId().IsNone() {
			return nil, errors.New("does not claim any edge")
		}
		claimId := bigStepLevelZeroEdge.ClaimId().Unwrap()
		claimedEdge, ok := ht.edges.TryGet(protocol.EdgeId(claimId))
		if !ok {
			return nil, errors.New("claimed edge not found")
		}
		start := honestLevelZero
		searchFor := claimedEdge
		blockChalAncestry, err := ht.findAncestorsInChallenge(start, searchFor)
		if err != nil {
			return nil, err
		}

		start = bigStepLevelZeroEdge
		searchFor = wantedEdge
		bigStepAncestry, err := ht.findAncestorsInChallenge(start, searchFor)
		if err != nil {
			return nil, err
		}
		ancestors := make([]protocol.EdgeId, 0)
		ancestors = append(ancestors, blockChalAncestry...)
		ancestors = append(ancestors, claimedEdge.Id())
		ancestors = append(ancestors, bigStepAncestry...)
		return ancestors, nil
	case protocol.SmallStepChallengeEdge:
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
		if smallStepLevelZeroEdge.ClaimId().IsNone() {
			return nil, errors.New("does not claim any edge")
		}
		claimId := smallStepLevelZeroEdge.ClaimId().Unwrap()
		claimedEdge, ok := ht.edges.TryGet(protocol.EdgeId(claimId))
		if !ok {
			return nil, errors.New("claimed edge not found")
		}
		start := honestLevelZero
		searchFor := claimedEdge
		blockChalAncestry, err := ht.findAncestorsInChallenge(start, searchFor)
		if err != nil {
			return nil, err
		}

		start = bigStepLevelZeroEdge
		searchFor = wantedEdge
		bigStepAncestry, err := ht.findAncestorsInChallenge(start, searchFor)
		if err != nil {
			return nil, err
		}
		ancestors := make([]protocol.EdgeId, 0)
		ancestors = append(ancestors, blockChalAncestry...)
		ancestors = append(ancestors, claimedEdge.Id())
		ancestors = append(ancestors, bigStepAncestry...)
		return ancestors, nil
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
