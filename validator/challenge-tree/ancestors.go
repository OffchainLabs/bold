package challengetree

import (
	"fmt"
	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/OffchainLabs/challenge-protocol-v2/util/threadsafe"
)

func (ct *challengeTree) addEdge(eg protocol.EdgeSnapshot) {
	prevAssertionId, err := ct.metadataReader.TopLevelAssertion(eg.Id())
	if err != nil {
		panic(err)
	}
	if ct.topLevelAssertionId != prevAssertionId {
		// Do nothing - this edge should not be part of this challenge tree.
		return
	}

	// Check if the edge id should be added to the rivaled edges set.
	mutualId := eg.MutualId()
	mutuals := ct.mutualIds.Get(mutualId)
	if mutuals == nil {
		ct.mutualIds.Put(mutualId, threadsafe.NewSet[protocol.EdgeId]())
		mutuals = ct.mutualIds.Get(mutualId)
	}
	if mutuals.Has(eg.Id()) {
		ct.rivaledEdges.Insert(eg.Id())
	} else {
		mutuals.Insert(eg.Id())
	}

	// We only need to check that we agree with the edge's start commitment,
	// and then we will necessarily track all edges we care about for the sake
	// of honest edge confirmations.
	_, startCommit := eg.StartCommitment()
	if ct.histChecker.AgreesWithStartHistoryCommitment(
		ct.metadataReader.ClaimHeights(eg.Id()),
		startCommit,
	) {
		ct.edges.Put(eg.Id(), eg)
		if eg.claimId != "" {
			switch eg.GetType() {
			case protocol.BlockChallengeEdge:
				ct.honestBlockChalLevelZeroEdge = util.Some(eg)
			case protocol.BigStepChallengeEdge:
				ct.honestBigStepChalLevelZeroEdge = util.Some(eg)
			case protocol.SmallStepChallengeEdge:
				ct.honestSmallStepChalLevelZeroEdge = util.Some(eg)
			default:
			}
		}
	}
}

// Consider the following set of edges in a challenge where evil
// edges are marked with a ' and a *:
//
//			     /---6---8
//		  0-----4
//			     \---6'--8'
//
//	   0*-------------8*
//
// The honest branch is the one that goes from 0-8. The evil edge is 0-8'.
// The evil edge 0-8' bisects, but agrees with the honest one from 0-4.
// Therefore, there is only a single 0-4 edge in the set.
//
// In this case, the set of ancestors for 4-6 is the following:
//
//	{4-8, 4-8', 0-8, 0-8'}
//
// All of these ancestors will contribute towards the path timer of 4-6 when
// confirmations by time are being attempted for the edge. Note there is another
// evil party that starts at 0*. In this case, it does not agree with the honest
// branch at all, so that party's edges can be ignored for this computation.
//
// In order to retrieve ancestors for an edge with id=I, we start from the honest,
// block challenge level zero edge and recursively traverse its children,
// reducing the ids and their rivals along the way into a slice until we hit a child that
// matches id=I and return the slice.
func (ct *challengeTree) ancestorsForHonestEdge(id protocol.EdgeId) []protocol.EdgeId {
	if ct.honestBlockChalLevelZeroEdge.IsNone() {
		return make([]protocol.EdgeId, 0)
	}
	blockEdge := ct.honestBlockChalLevelZeroEdge.Unwrap()
	ancestors, ok := ct.ancestorQuery(
		make([]protocol.EdgeId, 0),
		blockEdge,
		id,
	)
	if !ok {
		return nil
	}
	// The confirm by time function in Solidity requires ancestors to be specified
	// from earliest to oldest, which is the reverse result of our recursion.
	util.Reverse(ancestors)
	return ancestors
}

func (ct *challengeTree) ancestorQuery(
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
			hasRival := ct.rivaledEdges.Has(curr.Id())
			if !hasRival {
				return accum, false
			}

			rivalIds := ct.rivals(curr)

			// If the edge is a block challenge edge, we continue the recursion starting from the honest
			// big step level zero edge, if it exists.
			if curr.GetType() == protocol.BlockChallengeEdge {
				if ct.honestBigStepChalLevelZeroEdge.IsNone() {
					return accum, false
				}
				honestLowerLevelEdge := ct.honestBigStepChalLevelZeroEdge.Unwrap()

				// Defensive check ensuring the honest level zero edge one challenge level below
				// claims the current edge id as its claim id.
				if honestLowerLevelEdge.claimId != claimId(curr.id) {
					return accum, false
				}
				accum = append(accum, rivalIds...)
				accum = append(accum, curr.Id())
				return ct.ancestorQuery(accum, honestLowerLevelEdge, queryingFor)
			}

			// If the edge is a big step challenge edge, we continue the recursion starting from the honest
			// small step level zero edge, if it exists.
			if curr.edgeType == protocol.BigStepChallengeEdge {
				if ct.honestSmallStepChalLevelZeroEdge.IsNone() {
					return accum, false
				}
				honestLowerLevelEdge := ct.honestSmallStepChalLevelZeroEdge.Unwrap()

				// Defensive check ensuring the honest level zero edge one challenge level below
				// claims the current edge id as its claim id.
				if honestLowerLevelEdge.claimId != protocol.ClaimId(curr.Id()) {
					return accum, false
				}

				accum = append(accum, rivalIds...)
				accum = append(accum, curr.Id())
				return ct.ancestorQuery(accum, honestLowerLevelEdge, queryingFor)
			}
		}
		return accum, false
	}
	rivalIds := ct.rivals(curr)
	accum = append(accum, rivalIds...)
	accum = append(accum, curr.Id())

	// If the edge id we are querying for is a direct child of the current edge, we append
	// the current edge to the ancestors list and return true.
	if isDirectChild(curr, queryingFor) {
		return accum, true
	}
	lowerChild, lowerOk := ct.edges.TryGet(curr.lowerChildId)
	if !lowerOk {
		panic("not lower")
	}
	upperChild, upperOk := ct.edges.TryGet(curr.upperChildId)
	if !upperOk {
		panic(fmt.Sprintf("not upper curr %s, upper=%s", curr.Id(), curr.upperChildId))
	}
	lowerAncestors, foundInLowerChildren := ct.ancestorQuery(
		accum, lowerChild, queryingFor,
	)
	upperAncestors, foundInUpperChildren := ct.ancestorQuery(
		accum,
		upperChild,
		queryingFor,
	)
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
	return parent.lowerChildId == childId || parent.upperChildId == childId
}

// Checks if an edge has any children.
func hasChildren(eg protocol.EdgeSnapshot) bool {
	return !eg.LowerChildSnapshot().IsNone() || !eg.UpperChildSnapshot().IsNone()
}
