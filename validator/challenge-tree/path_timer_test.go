package challengetree

import (
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
	// Here are the creation times of each edge.
	//
	// 8a-16a = T20
	// 8a-16a = T20
	// 8b-16b = T15
	// 8c-16b = T15

	creationTimes := map[edgeId]uint64{
		// Alice's edges (honest).
		"0-16a":  1,
		"0-8a":   2,
		"8a-16a": 2,
		"0-4a":   3,
		"4a-8a":  3,
	}

	h := &helper{
		edges: map[edgeId]*edg{
			// Alice's edges (honest).
			"0-16a": {
				id:         "0-16a",
				mutualId:   "0-16",
				lowerChild: "0-8a",
				upperChild: "8a-16a",
			},
			"0-8a": {
				id:         "0-8a",
				mutualId:   "0-8",
				lowerChild: "0-4a",
				upperChild: "4a-8a",
			},
			"0-4a": {
				id:         "0-4a",
				mutualId:   "0-4",
				lowerChild: "",
				upperChild: "",
			},
			"4a-8a": {
				id:         "4a-8a",
				mutualId:   "4a-8",
				lowerChild: "",
				upperChild: "",
			},
			"8a-16a": {
				id:         "8a-16a",
				mutualId:   "8a-16",
				lowerChild: "",
				upperChild: "",
			},
		},
		creationTimes: creationTimes,
	}
	total := h.pathTimer(h.edges["4a-8a"], 3)
	t.Log(total)
}
