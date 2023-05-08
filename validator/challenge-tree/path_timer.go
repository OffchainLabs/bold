package challengetree

import (
	"fmt"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
)

// Computes the path timer for an edge at time T. A path timer is defined recursively
// via min/maxing edges creation times and their rivals along a path of ancestors
// within a challenge. The mathematical definition is as follows:
//
//	path_timer(e: edge) -> local_timer(e) + max{ path_timer(p) | p in parents(e) }
//
// This definition captures the sum of all local timers of the maximum-contributing
// edges along an edge e's ancestor path.
func (ct *challengeTree) pathTimer(e *edge, t uint64) (uint64, error) {
	if t < e.creationTime {
		return 0, nil
	}
	local, err := ct.localTimer(e, t)
	if err != nil {
		return 0, err
	}
	edgeParents := ct.parents(e.id)
	parentTimers := make([]uint64, len(edgeParents))

	// We compute a recursion over all of an edge's parents.
	for i, parent := range edgeParents {
		parentEdge, ok := ct.edges.TryGet(parent)
		if !ok {
			return 0, fmt.Errorf(
				"parent edge with id %#x not found in challenge tree",
				parent,
			)
		}
		computed, err := ct.pathTimer(parentEdge, t)
		if err != nil {
			return 0, err
		}
		parentTimers[i] = computed
	}

	// If a maximum is not defined, we return the local timer of the current edge.
	maxTimerOpt := util.Max(parentTimers)
	if maxTimerOpt.IsNone() {
		return local, nil
	}
	// Else, we return the sum of the edge's local timer plus the maximum path
	// timer of all its ancestors.
	return local + maxTimerOpt.Unwrap(), nil
}

// Gets the local timer of an edge at time T. If T is earlier than the edge's creation,
// this function will return 0.
func (ct *challengeTree) localTimer(e *edge, t uint64) (uint64, error) {
	if t < e.creationTime {
		return 0, nil
	}
	// If no rival at time t, then the local timer is defined
	// as t - t_creation(e).
	unrivaled, err := ct.unrivaledAtTime(e, t)
	if err != nil {
		return 0, err
	}
	if unrivaled {
		return t - e.creationTime, nil
	}
	// Else we return the earliest created rival's time: t_rival - t_creation(e).
	earliest, err := ct.earliestCreatedRivalTimestamp(e)
	if err != nil {
		return 0, nil
	}
	tRival := earliest.Unwrap()
	if e.creationTime >= tRival {
		return 0, nil
	}
	return tRival - e.creationTime, nil
}

// Gets all edges that claim a specified edge as their lower or upper child.
func (ct *challengeTree) parents(childId edgeId) []edgeId {
	p := make([]edgeId, 0)
	ct.edges.ForEach(func(_ edgeId, edge *edge) {
		if edge.lowerChildId == childId || edge.upperChildId == childId {
			p = append(p, edge.id)
		}
	})
	return p
}

// Gets the minimum creation timestamp across all of an edge's rivals. If an edge
// has no rivals, this minimum is undefined.
func (ct *challengeTree) earliestCreatedRivalTimestamp(e *edge) (util.Option[uint64], error) {
	rivals, err := ct.rivalsWithCreationTimes(e)
	if err != nil {
		return util.None[uint64](), err
	}
	timestamps := make([]uint64, len(rivals))
	for i, r := range rivals {
		timestamps[i] = r.creationTime
	}
	return util.Min(timestamps), nil
}

// Determines if an edge was unrivaled at timestamp T. If any rival existed
// for the edge at T, this function will return false.
func (ct *challengeTree) unrivaledAtTime(e *edge, t uint64) (bool, error) {
	if t < e.creationTime {
		return false, fmt.Errorf(
			"edge creation timestamp %d less than specified time %d",
			e.creationTime,
			t,
		)
	}
	rivals, err := ct.rivalsWithCreationTimes(e)
	if err != nil {
		return false, err
	}
	if len(rivals) == 0 {
		return true, nil
	}
	for _, r := range rivals {
		// If a rival existed before or at the time of the edge's
		// creation, we then return false.
		if r.creationTime <= t {
			return false, nil
		}
	}
	return true, nil
}

// Contains a rival edge's id and its creation time.
type rival struct {
	id           edgeId
	creationTime uint64
}

// Computes the set of rivals with their creation timestamp for an edge being tracked
// by the challenge tree. We do this by computing the mutual id of the edge and fetching
// all edge ids that share the same one from a set the challenge tree keeps track of.
// We exclude the specified edge from the returned list of rivals.
func (ct *challengeTree) rivalsWithCreationTimes(eg *edge) ([]*rival, error) {
	rivals := make([]*rival, 0)
	// If the edge is unrivaled, we simply return.
	if !ct.rivaledEdges.Has(eg.id) {
		return rivals, nil
	}
	mutualId := eg.computeMutualId()
	mutuals, ok := ct.mutualIds.TryGet(mutualId)
	if !ok {
		// Defensive check: this should not happen, as the challenge tree
		// should ensure to track the mutual ids on edge insertion.
		return nil, fmt.Errorf("mutual id %#x not found in challenge tree", mutualId)
	}
	mutuals.ForEach(func(rivalId edgeId) {
		if rivalId == eg.id {
			return
		}
		rivals = append(rivals, &rival{
			id:           rivalId,
			creationTime: ct.edges.Get(rivalId).creationTime,
		})
	})
	return rivals, nil
}
