package challengetree

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPathTimer_FlipFlop(t *testing.T) {
	// Setup the following challenge tree, where
	// branch `a` is honest.
	//
	// 0-----4a----- 8a-------16a
	//        \------8b-------16b
	//
	// Here are the creation times of each edge:
	//
	//   Alice
	//     0-16a        = T1
	//     0-8a, 8a-16a = T3
	//     0-4a, 4a-8a  = T5
	//
	//   Bob
	//     0-16b        = T2
	//     0-8b, 8b-16b = T4
	//     4a-8b        = T6
	//
	edges := buildEdges(
		// Alice.
		withCreationTime("0-16a", 1),
		withCreationTime("8a-16a", 3),
		withCreationTime("0-8a", 3),
		withCreationTime("4a-8a", 5),
		withCreationTime("0-4a", 5),
		// Bob.
		withCreationTime("0-16b", 2),
		withCreationTime("8b-16b", 4),
		withCreationTime("0-8b", 4),
		withCreationTime("4a-8b", 6),
	)
	// Child-relationship linking.
	// Alice.
	edges["0-16a"].lowerChild = "0-8a"
	edges["0-16a"].upperChild = "8a-16a"
	edges["0-8a"].lowerChild = "0-4a"
	edges["0-8a"].upperChild = "4a-8a"
	// Bob.
	edges["0-16b"].lowerChild = "0-8b"
	edges["0-16b"].upperChild = "8b-16b"
	edges["0-8b"].lowerChild = "0-4a"
	edges["0-8b"].upperChild = "4a-8b"

	h := &helper{
		edges: edges,
	}

	// Edge was not created before time T5.
	for i := 0; i < 5; i++ {
		total := h.pathTimer(h.edges["4a-8a"], uint64(1))
		require.Equal(t, uint64(0), total)
	}

	total := h.pathTimer(h.edges["4a-8a"], 5)
	require.Equal(t, uint64(2), total)
	// TODO: Is this correct?
	total = h.pathTimer(h.edges["4a-8a"], 6)
	require.Equal(t, uint64(3), total)

	// Add a in a new level zero edge that will bisect to
	// merge at height 4 with Alice.
	//   Charlie
	//     0-16c        = T10
	//     0-8b, 8b-16b = T11
	//     4a-8b        = T12
	//
	lateEdges := buildEdges(
		// Charlie.
		withCreationTime("0-16c", 10),
		withCreationTime("8a-16c", 11),
		withCreationTime("0-8c", 11),
		withCreationTime("4a-8c", 12),
	)
	for k, v := range lateEdges {
		h.edges[k] = v
	}

	// Ensure Alice's path timer does not change if this occurs.
	total = h.pathTimer(h.edges["4a-8a"], 5)
	require.Equal(t, uint64(2), total)
	// TODO: Is this correct?
	total = h.pathTimer(h.edges["4a-8a"], 6)
	require.Equal(t, uint64(3), total)
}

func buildEdges(allEdges ...*edg) map[edgeId]*edg {
	m := make(map[edgeId]*edg)
	for _, e := range allEdges {
		m[e.id] = e
	}
	return m
}

func withCreationTime(id string, createdAt uint64) *edg {
	return &edg{
		id:           edgeId(id),
		mutualId:     id[:len(id)-1], // Strip off the last char.
		lowerChild:   "",
		upperChild:   "",
		creationTime: createdAt,
	}
}
