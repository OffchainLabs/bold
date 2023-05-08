package challengetree

import (
	"context"
	"fmt"
	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/OffchainLabs/challenge-protocol-v2/util/threadsafe"
)

// Computes the path timer for an edge at time T. A path timer is defined recursively
// via min/maxing edges creation times and their rivals along a path of ancestors
// within a challenge. The mathematical definition is as follows:
//
//	path_timer(e: edge) -> local_timer(e) + max{ path_timer(p) | p in parents(e) }
//
// This definition captures the sum of all local timers of the maximum-contributing
// edges along an edge e's ancestor path.
func (ct *challengeTree) pathTimer(ctx context.Context, e protocol.EdgeGetter, t uint64) (uint64, error) {
	if t < e.CreatedAtBlock() {
		return 0, nil
	}
	local, err := ct.localTimer(ctx, e, t)
	if err != nil {
		return 0, err
	}
	edgeParents, err := ct.parents(ctx, e.Id())
	if err != nil {
		return 0, err
	}
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
		computed, err := ct.pathTimer(ctx, parentEdge, t)
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
func (ct *challengeTree) localTimer(ctx context.Context, e protocol.EdgeGetter, t uint64) (uint64, error) {
	if t < e.CreatedAtBlock() {
		return 0, nil
	}
	// If no rival at time t, then the local timer is defined
	// as t - t_creation(e).
	unrivaled, err := ct.unrivaledAtTime(ctx, e, t)
	if err != nil {
		return 0, err
	}
	if unrivaled {
		return t - e.CreatedAtBlock(), nil
	}
	// Else we return the earliest created rival's time: t_rival - t_creation(e).
	// This unwrap is safe because the edge has rivals at this point due to the check above.
	earliest, err := ct.earliestCreatedRivalTimestamp(ctx, e)
	if err != nil {
		return 0, err
	}
	tRival := earliest.Unwrap()
	if e.CreatedAtBlock() >= tRival {
		return 0, nil
	}
	return tRival - e.CreatedAtBlock(), nil
}

// Gets all edges that claim a specified edge as their lower or upper child.
func (ct *challengeTree) parents(ctx context.Context, childId protocol.EdgeId) ([]protocol.EdgeId, error) {
	p := make([]protocol.EdgeId, 0)
	if err := ct.edges.ForEach(func(_ protocol.EdgeId, edge protocol.EdgeGetter) error {
		lowerChild, err := edge.LowerChild(ctx)
		if err != nil {
			return err
		}
		if !lowerChild.IsNone() {
			if lowerChild.Unwrap() == childId {
				p = append(p, edge.Id())
				return nil
			}
		}
		upperChild, err := edge.UpperChild(ctx)
		if err != nil {
			return err
		}
		if !upperChild.IsNone() {
			if upperChild.Unwrap() == childId {
				p = append(p, edge.Id())
				return nil
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return p, nil
}

// Gets the minimum creation timestamp across all of an edge's rivals. If an edge
// has no rivals, this minimum is undefined.
func (ct *challengeTree) earliestCreatedRivalTimestamp(ctx context.Context, e protocol.EdgeGetter) (util.Option[uint64], error) {
	rivals, err := ct.rivalsWithCreationTimes(ctx, e)
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
func (ct *challengeTree) unrivaledAtTime(ctx context.Context, e protocol.EdgeGetter, t uint64) (bool, error) {
	if t < e.CreatedAtBlock() {
		return false, fmt.Errorf(
			"edge creation block %d less than specified %d",
			e.CreatedAtBlock(),
			t,
		)
	}
	rivals, err := ct.rivalsWithCreationTimes(ctx, e)
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
	id           protocol.EdgeId
	creationTime uint64
}

// Computes the set of rivals with their creation timestamp for an edge being tracked
// by the challenge tree. We do this by computing the mutual id of the edge and fetching
// all edge ids that share the same one from a set the challenge tree keeps track of.
// We exclude the specified edge from the returned list of rivals.
func (ct *challengeTree) rivalsWithCreationTimes(ctx context.Context, eg protocol.EdgeGetter) ([]*rival, error) {
	rivals := make([]*rival, 0)
	// If the edge is unrivaled, we simply return.
	if !ct.rivaledEdges.Has(eg.Id()) {
		return rivals, nil
	}
	mutualId, err := eg.ComputeMutualId(ctx)
	if err != nil {
		return nil, err
	}
	mutuals := ct.mutualIds.Get(mutualId)
	if mutuals == nil {
		ct.mutualIds.Put(mutualId, threadsafe.NewSet[protocol.EdgeId]())
		return rivals, nil
	}
	if err := mutuals.ForEach(func(rivalId protocol.EdgeId) error {
		if rivalId == eg.Id() {
			return nil
		}
		rivals = append(rivals, &rival{
			id:           rivalId,
			creationTime: ct.edges.Get(rivalId).CreatedAtBlock(),
		})
		return nil
	}); err != nil {
		return nil, err
	}
	return rivals, nil
}
