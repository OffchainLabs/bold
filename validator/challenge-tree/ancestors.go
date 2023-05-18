package challengetree

import (
	"fmt"
	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/OffchainLabs/challenge-protocol-v2/util/threadsafe"
	"github.com/pkg/errors"
)

var (
	ErrNoHonestTopLevelEdge = errors.New("no honest block challenge edge being tracked")
	ErrNotFound             = errors.New("not found in honest challenge tree")
	ErrNoLevelZero          = errors.New("no level zero edge with origin id found")
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
func (ht *HonestChallengeTree) AncestorsForHonestEdge(queryingFor protocol.EdgeId) ([]protocol.EdgeId, error) {
	wantedEdge, ok := ht.edges.TryGet(queryingFor)
	if !ok {
		return nil, errNotFound(queryingFor)
	}
	if ht.honestBlockChalLevelZeroEdge.IsNone() {
		return nil, ErrNoHonestTopLevelEdge
	}
	honestLevelZero := ht.honestBlockChalLevelZeroEdge.Unwrap()

	// Figure out what kind of edge this is, and apply different logic based on it.
	switch wantedEdge.GetType() {
	case protocol.BlockChallengeEdge:
		// If the edge is a block challenge edge, we simply search for the wanted edge's ancestors
		// in the block challenge starting from the honest, level zero edge and return
		// the computed ancestor ids list.
		start := honestLevelZero
		searchFor := wantedEdge
		ancestry, err := ht.findAncestorsInChallenge(start, searchFor)
		if err != nil {
			return nil, err
		}
		// The solidity confirmations function expects a child-to-parent ordering,
		// which is the reverse of our computed list.
		util.Reverse(ancestry)
		return ancestry, nil
	case protocol.BigStepChallengeEdge:
		ancestry := make([]protocol.EdgeId, 0)
		originId := wantedEdge.OriginId()

		// If the edge is a block challenge edge, we simply search for the wanted edge's ancestors
		// in the block challenge starting from the honest, level zero edge and return
		// the computed ancestor ids list.
		bigStepLevelZero, ok := findOriginEdge(originId, ht.honestBigStepLevelZeroEdges)
		if !ok {
			return nil, errNoLevelZero(originId)
		}

		start := bigStepLevelZero
		searchFor := wantedEdge
		bigStepAncestry, err := ht.findAncestorsInChallenge(start, searchFor)
		if err != nil {
			return nil, err
		}

		claimedEdge, err := ht.getClaimedEdge(bigStepLevelZero)
		if err != nil {
			return nil, err
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
		originId := wantedEdge.OriginId()
		smallStepLevelZero, ok := findOriginEdge(originId, ht.honestSmallStepLevelZeroEdges)
		if !ok {
			return nil, errNoLevelZero(originId)
		}
		start := smallStepLevelZero
		searchFor := wantedEdge
		smallStepAncestry, err := ht.findAncestorsInChallenge(start, searchFor)
		if err != nil {
			return nil, err
		}
		claimedBigStepEdge, err := ht.getClaimedEdge(smallStepLevelZero)
		if err != nil {
			return nil, err
		}
		originId = claimedBigStepEdge.OriginId()
		bigStepLevelZero, ok := findOriginEdge(originId, ht.honestBigStepLevelZeroEdges)
		if !ok {
			return nil, errNoLevelZero(originId)
		}

		start = bigStepLevelZero
		searchFor = claimedBigStepEdge
		bigStepAncestry, err := ht.findAncestorsInChallenge(start, searchFor)
		if err != nil {
			return nil, err
		}

		claimedBlockEdge, err := ht.getClaimedEdge(bigStepLevelZero)
		if err != nil {
			return nil, err
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
		return nil, fmt.Errorf("edge with type %v not supported", wantedEdge.GetType())
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
		bisectTo, err := util.BisectionPoint(uint64(currStart), uint64(currEnd))
		if err != nil {
			return nil, errors.Wrapf(err, "could not bisect start=%d, end=%d", currStart, currEnd)
		}
		if uint64(wantedEdgeStart) < bisectTo {
			lowerSnapshot := curr.LowerChildSnapshot()
			if lowerSnapshot.IsNone() {
				return nil, fmt.Errorf("edge %#x had no lower child", curr.Id())
			}
			curr = ht.edges.Get(lowerSnapshot.Unwrap())
		} else {
			upperSnapshot := curr.UpperChildSnapshot()
			if upperSnapshot.IsNone() {
				return nil, fmt.Errorf("edge %#x had no upper child", curr.Id())
			}
			curr = ht.edges.Get(upperSnapshot.Unwrap())
		}
	}
	if !found {
		return nil, errNotFound(queryingFor.Id())
	}
	return ancestry, nil
}

func (ht *HonestChallengeTree) getClaimedEdge(edge protocol.EdgeSnapshot) (protocol.EdgeSnapshot, error) {
	if edge.ClaimId().IsNone() {
		return nil, errors.New("does not claim any edge")
	}
	claimId := edge.ClaimId().Unwrap()
	claimedBlockEdge, ok := ht.edges.TryGet(protocol.EdgeId(claimId))
	if !ok {
		return nil, errors.New("claimed edge not found")
	}
	return claimedBlockEdge, nil
}

func findOriginEdge(originId protocol.OriginId, edges *threadsafe.Slice[protocol.EdgeSnapshot]) (protocol.EdgeSnapshot, bool) {
	var originEdge protocol.EdgeSnapshot
	found := edges.Find(func(_ int, e protocol.EdgeSnapshot) bool {
		if e.OriginId() == originId {
			originEdge = e
			return true
		}
		return false
	})
	return originEdge, found
}

func errNotFound(id protocol.EdgeId) error {
	return errors.Wrapf(ErrNotFound, "id=%#x", id)
}

func errNoLevelZero(originId protocol.OriginId) error {
	return errors.Wrapf(ErrNoLevelZero, "originId=%#x", originId)
}
