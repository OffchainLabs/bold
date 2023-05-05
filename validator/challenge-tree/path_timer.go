package challengetree

import "github.com/OffchainLabs/challenge-protocol-v2/util"

type unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func (ct *challengeTree) pathTimer(e *edge, t uint64) uint64 {
	if t < e.creationTime {
		return 0
	}
	local := ct.localTimer(e, t)
	edgeParents := ct.parents(e)
	parentTimers := make([]uint64, len(edgeParents))
	for i, parent := range edgeParents {
		parentEdge, ok := ct.edges.Get(parent)
		if !ok {
			panic("should not happen")
		}
		parentTimers[i] = ct.pathTimer(
			parentEdge,
			t,
		)
	}
	maxTimerOpt := max(parentTimers)
	if maxTimerOpt.IsNone() {
		return local
	}
	return local + maxTimerOpt.Unwrap()
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

func (ct *challengeTree) localTimer(e *edge, t uint64) uint64 {
	if t < e.creationTime {
		return 0
	}
	// If no rival at time t, then the local timer is defined
	// as t - t_creation(e).
	if ct.unrivaledAtTime(e, t) {
		return t - e.creationTime
	}
	// Else we return tRival minus the edge's creation time.
	tRival := ct.tRival(e).Unwrap()
	if e.creationTime >= tRival {
		return 0
	}
	return tRival - e.creationTime
}

func (ct *challengeTree) tRival(e *edge) util.Option[uint64] {
	rivalTimes := ct.rivalCreationTimes(e)
	return min(rivalTimes)
}

func (ct *challengeTree) unrivaledAtTime(e *edge, t uint64) bool {
	rivalTimes := ct.rivalCreationTimes(e)
	if len(rivalTimes) == 0 {
		return true
	}
	for _, rTime := range rivalTimes {
		// If a rival existed before or at the time of the edge's
		// creation, we then return false.
		if rTime <= t {
			return false
		}
	}
	return true
}

func (ct *challengeTree) rivalCreationTimes(e *edge) []uint64 {
	rivals := ct.rivals(e)
	if len(rivals) == 0 {
		return make([]uint64, 0)
	}
	timers := make([]uint64, len(rivals))
	for i, rivalId := range rivals {
		rival, ok := ct.edges.Get(rivalId)
		if !ok {
			panic("should not happen")
		}
		timers[i] = rival.creationTime
	}
	return timers
}

func min[T unsigned](items []T) util.Option[T] {
	if len(items) == 0 {
		return util.None[T]()
	}
	var m T
	if len(items) > 0 {
		m = items[0]
	}
	for i := 1; i < len(items); i++ {
		if items[i] < m {
			m = items[i]
		}
	}
	return util.Some(m)
}

func max[T unsigned](items []T) util.Option[T] {
	if len(items) == 0 {
		return util.None[T]()
	}
	var m T
	if len(items) > 0 {
		m = items[0]
	}
	for i := 1; i < len(items); i++ {
		if items[i] > m {
			m = items[i]
		}
	}
	return util.Some(m)
}
