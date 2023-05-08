package challengetree

import (
	"testing"

	"github.com/OffchainLabs/challenge-protocol-v2/util/threadsafe"
	"github.com/stretchr/testify/require"
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
	edges["0-16a"].lowerChildId = "0-8a"
	edges["0-16a"].upperChildId = "8a-16a"
	edges["0-8a"].lowerChildId = "0-4a"
	edges["0-8a"].upperChildId = "4a-8a"
	// Bob.
	edges["0-16b"].lowerChildId = "0-8b"
	edges["0-16b"].upperChildId = "8b-16b"
	edges["0-8b"].lowerChildId = "0-4a"
	edges["0-8b"].upperChildId = "4a-8b"

	allEdges := threadsafe.NewMapFromItems(edges)
	ct := &challengeTree{
		edges:        allEdges,
		mutualIds:    threadsafe.NewMap[mutualId, *threadsafe.Set[edgeId]](),
		rivaledEdges: threadsafe.NewSet[edgeId](),
	}

	// Edge was not created before time T5.
	for i := 0; i < 5; i++ {
		total, err := ct.pathTimer(ct.edges.Get("4a-8a"), uint64(1))
		require.NoError(t, err)
		require.Equal(t, uint64(0), total)
	}

	// Test out Alice's timers.
	total, err := ct.pathTimer(ct.edges.Get("4a-8a"), 5)
	require.NoError(t, err)

	require.Equal(t, uint64(6), total)
	// TODO: Is this correct?
	total, err = ct.pathTimer(ct.edges.Get("4a-8a"), 6)
	require.NoError(t, err)
	require.Equal(t, uint64(9), total)

	// Test out Bob's timers (was created after Alice).
	// Given Bob was never unrivaled, its edges should have a timer of 0.
	total, err = ct.pathTimer(ct.edges.Get("4a-8b"), 6)
	require.NoError(t, err)
	require.Equal(t, uint64(0), total)
	total, err = ct.pathTimer(ct.edges.Get("4a-8b"), 7)
	require.NoError(t, err)
	require.Equal(t, uint64(0), total)

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
		ct.edges.Put(k, v)
	}

	// Ensure Alice's path timer does not change if this occurs.
	total, err = ct.pathTimer(ct.edges.Get("4a-8a"), 5)
	require.NoError(t, err)
	require.Equal(t, uint64(2), total)
	// TODO: Is this correct?
	total, err = ct.pathTimer(ct.edges.Get("4a-8a"), 6)
	require.NoError(t, err)
	require.Equal(t, uint64(3), total)

	// Ensure Bob's path timer does not change if this occurs.
	// Given Bob was never unrivaled, its edges should have a timer of 0.
	total, err = ct.pathTimer(ct.edges.Get("4a-8b"), 6)
	require.NoError(t, err)
	require.Equal(t, uint64(0), total)
	total, err = ct.pathTimer(ct.edges.Get("4a-8b"), 7)
	require.NoError(t, err)
	require.Equal(t, uint64(0), total)
}

func buildEdges(allEdges ...*edge) map[edgeId]*edge {
	m := make(map[edgeId]*edge)
	for _, e := range allEdges {
		m[e.id] = e
	}
	return m
}

func withCreationTime(id string, createdAt uint64) *edge {
	return &edge{
		id: edgeId(id),
		//mutualId:     id[:len(id)-1], // Strip off the last char.
		lowerChildId: "",
		upperChildId: "",
		creationTime: createdAt,
	}
}
