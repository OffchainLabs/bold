package challengetree

import "github.com/OffchainLabs/challenge-protocol-v2/util"

type unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type edgeId string

type edg struct {
	id           edgeId
	mutualId     string // maybe id minus last char for testing?
	lowerChild   edgeId
	upperChild   edgeId
	creationTime uint64
}

type helper struct {
	edges map[edgeId]*edg
}

func (h *helper) pathTimer(e *edg, t uint64) uint64 {
	if t < e.creationTime {
		return 0
	}
	local := h.localTimer(e, t)
	edgeParents := h.parents(e)
	parentTimers := make([]uint64, len(edgeParents))
	for i, parent := range edgeParents {
		parentEdge, ok := h.edges[parent]
		if !ok {
			panic("should not happen")
		}
		parentTimers[i] = h.pathTimer(
			parentEdge,
			e.creationTime,
		)
	}
	maxTimerOpt := max(parentTimers)
	if maxTimerOpt.IsNone() {
		return local
	}
	return local + maxTimerOpt.Unwrap()
}

// Naive parent lookup just for testing purposes.
func (h *helper) parents(e *edg) []edgeId {
	p := make([]edgeId, 0)
	for _, edge := range h.edges {
		if edge.lowerChild == e.id || edge.upperChild == e.id {
			p = append(p, edge.id)
		}
	}
	return p
}

func (h *helper) localTimer(e *edg, t uint64) uint64 {
	if t < e.creationTime {
		return 0
	}
	// If no rival at time t, then the local timer is defined
	// as t - t_creation(e).
	if h.unrivaledAtTime(e, t) {
		return t - e.creationTime
	}
	// Else we return tRival minus the edge's creation time.
	tRival := h.tRival(e).Unwrap()
	if e.creationTime >= tRival {
		return 0
	}
	return tRival - e.creationTime
}

func (h *helper) tRival(e *edg) util.Option[uint64] {
	rivalTimes := h.rivalCreationTimes(e)
	return min(rivalTimes)
}

func (h *helper) unrivaledAtTime(e *edg, t uint64) bool {
	rivalTimes := h.rivalCreationTimes(e)
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

func (h *helper) rivalCreationTimes(e *edg) []uint64 {
	rivals := h.rivals(e)
	if len(rivals) == 0 {
		return make([]uint64, 0)
	}
	timers := make([]uint64, len(rivals))
	for i, rivalId := range rivals {
		rival, ok := h.edges[rivalId]
		if !ok {
			panic("should not happen")
		}
		timers[i] = rival.creationTime
	}
	return timers
}

func (h *helper) rivals(e *edg) []edgeId {
	rivals := make([]edgeId, 0)
	for edgeId, potentialRival := range h.edges {
		if edgeId == e.id {
			continue
		}
		if potentialRival.mutualId == e.mutualId {
			rivals = append(rivals, edgeId)
		}
	}
	return rivals
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
