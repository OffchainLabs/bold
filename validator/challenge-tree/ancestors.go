package challengetree

import (
	"fmt"
	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
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
func (ht *HonestChallengeTree) AncestorsForHonestEdge(id protocol.EdgeId) ([]protocol.EdgeId, error) {
	if _, ok := ht.edges.TryGet(id); !ok {
		return nil, fmt.Errorf("edge with id %#x not found in honest challenge tree", id)
	}
	if ht.honestBlockChalLevelZeroEdge.IsNone() {
		return make([]protocol.EdgeId, 0), nil
	}
	blockEdge := ht.honestBlockChalLevelZeroEdge.Unwrap()
	ancestors, ok := ht.ancestorQuery(
		make([]protocol.EdgeId, 0),
		blockEdge,
		id,
	)
	if !ok {
		return make([]protocol.EdgeId, 0), nil
	}
	// The confirm by time function in Solidity requires ancestors to be specified
	// from earliest to oldest, which is the reverse result of our recursion.
	util.Reverse(ancestors)
	return ancestors, nil
}

func (ht *HonestChallengeTree) ancestorQuery(
	accum []protocol.EdgeId,
	curr protocol.EdgeSnapshot,
	queryingFor protocol.EdgeId,
) ([]protocol.EdgeId, bool) {
	if curr.Id() == queryingFor {
		return accum, true
	}
	if !hasChildren(curr) {
		// If the edge has length 1, we then perform a few special checks.
		if edgeLength(curr) == 1 {
			// In case the edge is a small step challenge of length 1, we simply return.
			if curr.GetType() == protocol.SmallStepChallengeEdge {
				return accum, false
			}

			// If the edge is unrivaled, we return.
			isRivaled := ht.isRivaled(curr)
			if !isRivaled {
				return accum, false
			}

			var lowerLevelEdge protocol.EdgeSnapshot
			// If the edge is a block challenge edge, we continue the recursion starting from the honest
			// big step level zero edge, if it exists.
			if curr.GetType() == protocol.BlockChallengeEdge {
				if ht.honestBigStepChalLevelZeroEdge.IsNone() {
					return accum, false
				}
				lowerLevelEdge = ht.honestBigStepChalLevelZeroEdge.Unwrap()
			}

			// If the edge is a big step challenge edge, we continue the recursion starting from the honest
			// small step level zero edge, if it exists.
			if curr.GetType() == protocol.BigStepChallengeEdge {
				if ht.honestSmallStepChalLevelZeroEdge.IsNone() {
					return accum, false
				}
				lowerLevelEdge = ht.honestSmallStepChalLevelZeroEdge.Unwrap()
			}
			// Defensive check ensuring the honest level zero edge one challenge level below
			// claims the current edge id as its claim id. If it does not, we simply return.
			if !checkEdgeClaim(lowerLevelEdge, protocol.ClaimId(curr.Id())) {
				return accum, false
			}
			accum = append(accum, curr.Id())
			return ht.ancestorQuery(accum, lowerLevelEdge, queryingFor)
		}
		return accum, false
	}
	accum = append(accum, curr.Id())

	// If the edge id we are querying for is a direct child of the current edge, we append
	// the current edge to the ancestors list and return true.
	if isDirectChild(curr, queryingFor) {
		return accum, true
	}
	var lowerAncestors []protocol.EdgeId
	var foundInLowerChildren bool
	if !curr.LowerChildSnapshot().IsNone() {
		lowerChildId := curr.LowerChildSnapshot().Unwrap()
		lowerChild := ht.edges.Get(lowerChildId)
		lowerAncestors, foundInLowerChildren = ht.ancestorQuery(
			accum, lowerChild, queryingFor,
		)
	}
	var upperAncestors []protocol.EdgeId
	var foundInUpperChildren bool
	if !curr.UpperChildSnapshot().IsNone() {
		upperChildId := curr.UpperChildSnapshot().Unwrap()
		upperChild := ht.edges.Get(upperChildId)
		upperAncestors, foundInUpperChildren = ht.ancestorQuery(
			accum, upperChild, queryingFor,
		)
	}
	// If the edge we are querying for is found in the lower children,
	// we return the ancestry along such path.
	if foundInLowerChildren {
		return lowerAncestors, true
	}
	// If the edge we are querying for is found in the upper children,
	// we return the ancestry along such path.
	if foundInUpperChildren {
		return upperAncestors, true
	}
	return accum, false
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

// Checks if an edge claims a certain claim id.
func checkEdgeClaim(edge protocol.EdgeSnapshot, claimId protocol.ClaimId) bool {
	return !edge.ClaimId().IsNone() && edge.ClaimId().Unwrap() == claimId
}

// Computes the length of an edge by taking the difference between
// its end and start heights.
//
// SAFETY: We will never receive a malformed edge, as the challenge tree is
// created from events emitted by successful challenge addition events
// in the protocol smart contracts.
func edgeLength(eg protocol.EdgeSnapshot) protocol.Height {
	startHeight, _ := eg.StartCommitment()
	endHeight, _ := eg.EndCommitment()
	return endHeight - startHeight
}

// Checks if an edge id is a direct child of a specified parent edge.
func isDirectChild(parent protocol.EdgeSnapshot, childId protocol.EdgeId) bool {
	lowerChild := parent.LowerChildSnapshot()
	upperChild := parent.UpperChildSnapshot()
	if lowerChild.IsNone() && lowerChild.Unwrap() == childId {
		return true
	}
	if upperChild.IsNone() && upperChild.Unwrap() == childId {
		return true
	}
	return false
}

// Checks if an edge has any children.
func hasChildren(eg protocol.EdgeSnapshot) bool {
	return !eg.LowerChildSnapshot().IsNone() || !eg.UpperChildSnapshot().IsNone()
}
