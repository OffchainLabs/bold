package challengetree

import "github.com/OffchainLabs/challenge-protocol-v2/util"
import "fmt"

func (ct *challengeTree) pathTimer(e *edge, t uint64) (uint64, error) {
	if t < e.creationTime {
		return 0, nil
	}
	local, err := ct.localTimer(e, t)
	if err != nil {
		return 0, err
	}
	edgeParents := ct.parents(e)
	parentTimers := make([]uint64, len(edgeParents))
	for i, parent := range edgeParents {
		parentEdge, ok := ct.edges.Get(parent)
		if !ok {
			return 0, fmt.Errorf("parent edge with id %#x not found in challenge tree", parent)
		}
		computed, err := ct.pathTimer(parentEdge, t)
		if err != nil {
			return 0, err
		}
		parentTimers[i] = computed
	}
	maxTimerOpt := util.Max(parentTimers)
	if maxTimerOpt.IsNone() {
		return local, nil
	}
	return local + maxTimerOpt.Unwrap(), nil
}

// Naive parent lookup just for testing purposes.
func (ct *challengeTree) parents(e *edge) []edgeId {
	p := make([]edgeId, 0)
	for _, k := range ct.edges.Keys() {
		edge, _ := ct.edges.Get(k)
		if edge.lowerChildId == e.id || edge.upperChildId == e.id {
			p = append(p, edge.id)
		}
	}
	return p
}

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
	// Else we return tRival minus the edge's creation time.
	someRival, err := ct.tRival(e)
	if err != nil {
		return 0, err
	}
	tRival := someRival.Unwrap()
	if e.creationTime >= tRival {
		return 0, nil
	}
	return tRival - e.creationTime, nil
}

func (ct *challengeTree) tRival(e *edge) (util.Option[uint64], error) {
	rivalTimes, err := ct.rivalCreationTimes(e)
	if err != nil {
		return util.None[uint64](), err
	}
	return util.Min(rivalTimes), nil
}

func (ct *challengeTree) unrivaledAtTime(e *edge, t uint64) (bool, error) {
	rivalTimes, err := ct.rivalCreationTimes(e)
	if err != nil {
		return false, err
	}
	if len(rivalTimes) == 0 {
		return true, nil
	}
	for _, rTime := range rivalTimes {
		// If a rival existed before or at the time of the edge's
		// creation, we then return false.
		if rTime <= t {
			return false, nil
		}
	}
	return true, nil
}

func (ct *challengeTree) rivalCreationTimes(e *edge) ([]uint64, error) {
	rivals := ct.rivals(e)
	if len(rivals) == 0 {
		return make([]uint64, 0), nil
	}
	timers := make([]uint64, len(rivals))
	for i, rivalId := range rivals {
		rival, ok := ct.edges.Get(rivalId)
		if !ok {
			return nil, fmt.Errorf("rival with id %#x not found in challenge tree", rivalId)
		}
		timers[i] = rival.creationTime
	}
	return timers, nil
}
