package challengetree

import "github.com/OffchainLabs/challenge-protocol-v2/util"

type unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type edgeId string

type edg struct {
	id           edgeId
	creationTime uint64
	lowerChild   edgeId
	upperChild   edgeId
}

type helper struct {
	rivalCreationTimes map[edgeId][]uint64
	edges              map[edgeId]*edg
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
	maxTimer := max(parentTimers).Unwrap()
	return local + maxTimer
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
	if p.unrivaledAtTime(e, t) {
		return t - e.creationTime
	}
	// Else we return tRival minus the edge's creation time.
	tRival := p.tRival(e).Unwrap()
	if e.creationTime >= tRival {
		return 0
	}
	return tRival - e.creationTime
}

func (h *helper) tRival(e *edg) util.Option[uint64] {
	rivalTimes := h.rivalCreationTimes[e.id]
	return min(rivalTimes)
}

func (h *helper) unrivaledAtTime(e *edg, t uint64) bool {
	rivalTimes, ok := h.rivalCreationTimes[e.id]
	if !ok {
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
