package challengetree

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPathTimer(t *testing.T) {
	// Setup the following challenge tree, where
	// branch `a` is honest.
	//
	// 0-----4----- 8a-------16a
	//        \-----8b-------16b
	//        \-----8c-------16c
	//        \-----8d-------16d
	//
	// Here are the creation times of each edge:
	edges := buildEdges(
		withCreationTime("0-16a", 1),
		withCreationTime("8a-16a", 2),
		withCreationTime("0-8a", 2),
		withCreationTime("4a-8a", 3),
		withCreationTime("0-4a", 3),
	)
	edges["0-16a"].lowerChild = "0-8a"
	edges["0-16a"].upperChild = "8a-16a"
	edges["0-8a"].lowerChild = "0-4a"
	edges["0-8a"].upperChild = "4a-8a"

	h := &helper{
		edges: edges,
	}

	// Edge was not created at time 1 nor 2.
	total := h.pathTimer(h.edges["4a-8a"], 1)
	require.Equal(t, uint64(0), total)
	total = h.pathTimer(h.edges["4a-8a"], 2)
	require.Equal(t, uint64(0), total)
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
