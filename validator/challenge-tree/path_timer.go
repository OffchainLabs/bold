package challengetree

import (
	"fmt"
	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/OffchainLabs/challenge-protocol-v2/util/threadsafe"
)

// HonestPathTimer returns the cumulative local timers of the
// honest branch of edges starting at a particular edge id.
func (ht *HonestChallengeTree) HonestPathTimer(edgeId protocol.EdgeId) (uint64, error) {
	timer, ok := ht.cumulativeHonestPathTimers.TryGet(edgeId)
	if !ok {
		return 0, fmt.Errorf("edge id %#x not being tracked in honest challenge tree", edgeId)
	}
	return timer, nil
}

func (ht *HonestChallengeTree) UpdateCumulativePathTimers(blockNum uint64) error {
	if ht.honestBlockChalLevelZeroEdge.IsNone() {
		return nil
	}
	blockEdge := ht.honestBlockChalLevelZeroEdge.Unwrap()
	return ht.recursiveTimersUpdate(
		0, // Total honest path timer accumulator, starting at 0.
		blockEdge,
		blockNum,
	)
}

// Recursively updates the path timers of all honest edges tracked in the challenge tree.
func (ht *HonestChallengeTree) recursiveTimersUpdate(
	timerAcc uint64,
	curr protocol.EdgeSnapshot,
	blockNum uint64,
) error {
	timer, err := ht.localTimer(curr, blockNum)
	if err != nil {
		return err
	}
	if !hasChildren(curr) {
		// If the edge has length 1, we then perform a few special checks.
		if edgeLength(curr) == 1 {
			isRivaled := ht.isRivaled(curr)
			// In case the edge is a small step challenge of length 1, or is not rivaled we simply return.
			if curr.GetType() == protocol.SmallStepChallengeEdge || !isRivaled {
				ht.cumulativeHonestPathTimers.Put(curr.Id(), timer+timerAcc)
				return nil
			}

			var lowerLevelEdge protocol.EdgeSnapshot
			// If the edge is a block challenge edge, we continue the recursion starting from the honest
			// big step level zero edge, if it exists.
			if curr.GetType() == protocol.BlockChallengeEdge {
				if ht.honestBigStepChalLevelZeroEdge.IsNone() {
					ht.cumulativeHonestPathTimers.Put(curr.Id(), timer+timerAcc)
					return nil
				}
				lowerLevelEdge = ht.honestBigStepChalLevelZeroEdge.Unwrap()
			}

			// If the edge is a big step challenge edge, we continue the recursion starting from the honest
			// small step level zero edge, if it exists.
			if curr.GetType() == protocol.BigStepChallengeEdge {
				if ht.honestSmallStepChalLevelZeroEdge.IsNone() {
					ht.cumulativeHonestPathTimers.Put(curr.Id(), timer+timerAcc)
					return nil
				}
				lowerLevelEdge = ht.honestSmallStepChalLevelZeroEdge.Unwrap()
			}
			// Defensive check ensuring the honest level zero edge one challenge level below
			// claims the current edge id as its claim id. If it does not, we simply return.
			if !checkEdgeClaim(lowerLevelEdge, protocol.ClaimId(curr.Id())) {
				ht.cumulativeHonestPathTimers.Put(curr.Id(), timer+timerAcc)
				return nil
			}
			// We recurse into the lower, subchallenge levels.
			return ht.recursiveTimersUpdate(timerAcc+timer, lowerLevelEdge, blockNum)
		}
		// Otherwise, the edge has no children and no lower-level edges that claim it,
		// so we simply update its cumulative timer and return.
		ht.cumulativeHonestPathTimers.Put(curr.Id(), timer+timerAcc)
		return nil
	}
	// We recurse into the edge's children, if it has any.
	if !curr.LowerChildSnapshot().IsNone() {
		lowerChildId := curr.LowerChildSnapshot().Unwrap()
		lowerChild := ht.edges.Get(lowerChildId)
		if err := ht.recursiveTimersUpdate(
			timerAcc+timer, lowerChild, blockNum,
		); err != nil {
			return err
		}
	}
	if !curr.UpperChildSnapshot().IsNone() {
		upperChildId := curr.UpperChildSnapshot().Unwrap()
		upperChild := ht.edges.Get(upperChildId)
		if err := ht.recursiveTimersUpdate(
			timerAcc+timer, upperChild, blockNum,
		); err != nil {
			return err
		}
	}
	// We update the edge's cumulative timer and return.
	ht.cumulativeHonestPathTimers.Put(curr.Id(), timer+timerAcc)
	return nil
}

// Gets the local timer of an edge at a block number, T. If T is earlier than the edge's creation,
// this function will return 0.
func (ht *HonestChallengeTree) localTimer(e protocol.EdgeSnapshot, blockNum uint64) (uint64, error) {
	if blockNum < e.CreatedAtBlock() {
		return 0, nil
	}
	// If no rival at a block num, then the local timer is defined
	// as t - t_creation(e).
	unrivaled, err := ht.unrivaledAtTime(e, blockNum)
	if err != nil {
		return 0, err
	}
	if unrivaled {
		return blockNum - e.CreatedAtBlock(), nil
	}
	// Else we return the earliest created rival's time: t_rival - t_creation(e).
	// This unwrap is safe because the edge has rivals at this point due to the check above.
	earliest := ht.earliestCreatedRivalTimestamp(e)
	tRival := earliest.Unwrap()
	if e.CreatedAtBlock() >= tRival {
		return 0, nil
	}
	return tRival - e.CreatedAtBlock(), nil
}

// Gets the minimum creation timestamp across all of an edge's rivals. If an edge
// has no rivals, this minimum is undefined.
func (ht *HonestChallengeTree) earliestCreatedRivalTimestamp(e protocol.EdgeSnapshot) util.Option[uint64] {
	rivals := ht.rivalsWithCreationTimes(e)
	creationBlocks := make([]uint64, len(rivals))
	for i, r := range rivals {
		creationBlocks[i] = uint64(r.createdAtBlock)
	}
	return util.Min(creationBlocks)
}

// Determines if an edge was unrivaled at a block num T. If any rival existed
// for the edge at T, this function will return false.
func (ht *HonestChallengeTree) unrivaledAtTime(e protocol.EdgeSnapshot, blockNum uint64) (bool, error) {
	if blockNum < e.CreatedAtBlock() {
		return false, fmt.Errorf(
			"edge creation block %d less than specified %d",
			e.CreatedAtBlock(),
			blockNum,
		)
	}
	rivals := ht.rivalsWithCreationTimes(e)
	if len(rivals) == 0 {
		return true, nil
	}
	for _, r := range rivals {
		// If a rival existed before or at the time of the edge's
		// creation, we then return false.
		if uint64(r.createdAtBlock) <= blockNum {
			return false, nil
		}
	}
	return true, nil
}

// Contains a rival edge's id and its creation time.
type rival struct {
	id             protocol.EdgeId
	createdAtBlock creationTime
}

// Computes the set of rivals with their creation timestamp for an edge being tracked
// by the challenge tree. We do this by computing the mutual id of the edge and fetching
// all edge ids that share the same one from a set the challenge tree keeps track of.
// We exclude the specified edge from the returned list of rivals.
func (ht *HonestChallengeTree) rivalsWithCreationTimes(eg protocol.EdgeSnapshot) []*rival {
	rivals := make([]*rival, 0)
	mutualId := eg.MutualId()
	mutuals := ht.mutualIds.Get(mutualId)
	if mutuals == nil {
		ht.mutualIds.Put(mutualId, threadsafe.NewMap[protocol.EdgeId, creationTime]())
		return rivals
	}
	_ = mutuals.ForEach(func(rivalId protocol.EdgeId, t creationTime) error {
		if rivalId == eg.Id() {
			return nil
		}
		rivals = append(rivals, &rival{
			id:             rivalId,
			createdAtBlock: t,
		})
		return nil
	})
	return rivals
}
