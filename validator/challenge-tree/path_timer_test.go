package challengetree

import (
	"strconv"
	"strings"
	"testing"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/OffchainLabs/challenge-protocol-v2/util/threadsafe"
	"github.com/stretchr/testify/require"
)

// The following tests checks a scenario where the honest
// and dishonest parties take turns making challenge moves,
// and as a result, their edges will be unrivaled for some time,
// contributing to the path timer of edges we will query in this test.
//
// We first setup the following challenge tree, where branch `a` is honest.
//
//	 0-----4a----- 8a-------16a
//		     \------8b-------16b
//
// Here are the creation times of each edge:
//
//	Alice (honest)
//	  0-16a        = T1
//	  0-8a, 8a-16a = T3
//	  0-4a, 4a-8a  = T5
//
//	Bob (evil)
//	  0-16b        = T2
//	  0-8b, 8b-16b = T4
//	  4a-8b        = T6
//
// In this contrived example, Alice and Bob's edges will have
// a time interval of 1 in which they are unrivaled.
func TestPathTimer_FlipFlop(t *testing.T) {
	edges := buildEdges(
		// Alice.
		withCreationTime(t, "assertionA", "blk-0.a-16.a", 1),
		withCreationTime(t, "assertionA", "blk-8.a-16.a", 3),
		withCreationTime(t, "assertionA", "blk-0.a-8.a", 3),
		withCreationTime(t, "assertionA", "blk-4.a-8.a", 5),
		withCreationTime(t, "assertionA", "blk-0.a-4.a", 5),
		// Bob.
		withCreationTime(t, "assertionA", "blk-0.a-16.b", 2),
		withCreationTime(t, "assertionA", "blk-8.b-16.b", 4),
		withCreationTime(t, "assertionA", "blk-0.a-8.b", 4),
		withCreationTime(t, "assertionA", "blk-4.a-8.b", 6),
	)
	// Child-relationship linking.
	// Alice.
	edges["blk-0.a-16.a"].lowerChildId = "blk-0.a-8.a"
	edges["blk-0.a-16.a"].upperChildId = "blk-8.a-16.a"
	edges["blk-0.a-8.a"].lowerChildId = "blk-0.a-4.a"
	edges["blk-0.a-8.a"].upperChildId = "blk-4.a-8.a"
	// Bob.
	edges["blk-0.a-16.b"].lowerChildId = "blk-0.a-8.b"
	edges["blk-0.a-16.b"].upperChildId = "blk-8.b-16.b"
	edges["blk-0.a-8.b"].lowerChildId = "blk-0.a-4.a"
	edges["blk-0.a-8.b"].upperChildId = "blk-4.a-8.b"

	allEdges := threadsafe.NewMapFromItems(edges)
	ct := &challengeTree{
		edges:        allEdges,
		mutualIds:    threadsafe.NewMap[mutualId, *threadsafe.Set[edgeId]](),
		rivaledEdges: threadsafe.NewSet[edgeId](),
	}

	// We then set up the rival relationships in the challenge tree.
	// All edges are rivaled in this example.
	for _, e := range edges {
		ct.rivaledEdges.Insert(e.id)
	}
	mutualA := edges["blk-0.a-16.a"].computeMutualId()
	ct.mutualIds.Put(mutualA, threadsafe.NewSet[edgeId]())
	mutuals := ct.mutualIds.Get(mutualA)
	mutuals.Insert("blk-0.a-16.a")
	mutuals.Insert("blk-0.a-16.b")

	mutualB := edges["blk-0.a-16.b"].computeMutualId()
	require.Equal(t, mutualA, mutualB)

	mutualC := edges["blk-8.b-16.b"].computeMutualId()
	mutualD := edges["blk-8.a-16.a"].computeMutualId()
	require.NotEqual(t, mutualC, mutualD)

	timer, err := ct.pathTimer(ct.edges.Get("blk-0.a-16.a"), uint64(1000))
	require.NoError(t, err)
	t.Log(timer)

	// // Edge was not created before time T5.
	// for i := 0; i < 5; i++ {
	// 	total, err := ct.pathTimer(ct.edges.Get("4a-8a"), uint64(1))
	// 	require.NoError(t, err)
	// 	require.Equal(t, uint64(0), total)
	// }

	// // Test out Alice's timers.
	// total, err := ct.pathTimer(ct.edges.Get("4a-8a"), 5)
	// require.NoError(t, err)

	// require.Equal(t, uint64(6), total)
	// // TODO: Is this correct?
	// total, err = ct.pathTimer(ct.edges.Get("4a-8a"), 6)
	// require.NoError(t, err)
	// require.Equal(t, uint64(9), total)

	// // Test out Bob's timers (was created after Alice).
	// // Given Bob was never unrivaled, its edges should have a timer of 0.
	// total, err = ct.pathTimer(ct.edges.Get("4a-8b"), 6)
	// require.NoError(t, err)
	// require.Equal(t, uint64(0), total)
	// total, err = ct.pathTimer(ct.edges.Get("4a-8b"), 7)
	// require.NoError(t, err)
	// require.Equal(t, uint64(0), total)

	// // Add a in a new level zero edge that will bisect to
	// // merge at height 4 with Alice.
	// //   Charlie
	// //     0-16c        = T10
	// //     0-8b, 8b-16b = T11
	// //     4a-8b        = T12
	// //
	// lateEdges := buildEdges(
	// 	// Charlie.
	// 	withCreationTime("0-16c", 10),
	// 	withCreationTime("8a-16c", 11),
	// 	withCreationTime("0-8c", 11),
	// 	withCreationTime("4a-8c", 12),
	// )
	// for k, v := range lateEdges {
	// 	ct.edges.Put(k, v)
	// }

	// // Ensure Alice's path timer does not change if this occurs.
	// total, err = ct.pathTimer(ct.edges.Get("4a-8a"), 5)
	// require.NoError(t, err)
	// require.Equal(t, uint64(2), total)
	// // TODO: Is this correct?
	// total, err = ct.pathTimer(ct.edges.Get("4a-8a"), 6)
	// require.NoError(t, err)
	// require.Equal(t, uint64(3), total)

	// // Ensure Bob's path timer does not change if this occurs.
	// // Given Bob was never unrivaled, its edges should have a timer of 0.
	// total, err = ct.pathTimer(ct.edges.Get("4a-8b"), 6)
	// require.NoError(t, err)
	// require.Equal(t, uint64(0), total)
	// total, err = ct.pathTimer(ct.edges.Get("4a-8b"), 7)
	// require.NoError(t, err)
	// require.Equal(t, uint64(0), total)
}

func Test_localTimer(t *testing.T) {
	ct := &challengeTree{
		edges:        threadsafe.NewMap[edgeId, *edge](),
		mutualIds:    threadsafe.NewMap[mutualId, *threadsafe.Set[edgeId]](),
		rivaledEdges: threadsafe.NewSet[edgeId](),
	}
	edgeA := withCreationTime(t, "assertionA", "blk-0.a-1.a", 3)
	ct.edges.Put(edgeA.id, edgeA)

	t.Run("zero if earlier than creation time", func(t *testing.T) {
		timer, err := ct.localTimer(edgeA, edgeA.creationTime-1)
		require.NoError(t, err)
		require.Equal(t, uint64(0), timer)
	})
	t.Run("no rival is simply difference between T and creation time", func(t *testing.T) {
		timer, err := ct.localTimer(edgeA, edgeA.creationTime)
		require.NoError(t, err)
		require.Equal(t, uint64(0), timer)
		timer, err = ct.localTimer(edgeA, edgeA.creationTime+3)
		require.NoError(t, err)
		require.Equal(t, uint64(3), timer)
		timer, err = ct.localTimer(edgeA, edgeA.creationTime+1000)
		require.NoError(t, err)
		require.Equal(t, uint64(1000), timer)
	})
	t.Run("if rivaled timer is difference between earliest rival and edge creation", func(t *testing.T) {
		edgeB := withCreationTime(t, "assertionA", "blk-0.a-1.b", 5)
		edgeC := withCreationTime(t, "assertionA", "blk-0.a-1.c", 10)
		ct.edges.Put(edgeB.id, edgeB)
		ct.edges.Put(edgeC.id, edgeC)
		ct.rivaledEdges.Insert(edgeA.id)
		ct.rivaledEdges.Insert(edgeB.id)
		ct.rivaledEdges.Insert(edgeC.id)
		ct.mutualIds.Put(edgeA.computeMutualId(), threadsafe.NewSet[edgeId]())
		mutuals := ct.mutualIds.Get(edgeA.computeMutualId())
		mutuals.Insert(edgeA.id)
		mutuals.Insert(edgeB.id)
		mutuals.Insert(edgeC.id)

		// Should get same result regardless of specified time.
		timer, err := ct.localTimer(edgeA, 100)
		require.NoError(t, err)
		require.Equal(t, edgeB.creationTime-edgeA.creationTime, timer)
		timer, err = ct.localTimer(edgeA, 10000)
		require.NoError(t, err)
		require.Equal(t, edgeB.creationTime-edgeA.creationTime, timer)
		timer, err = ct.localTimer(edgeA, 1000000)
		require.NoError(t, err)
		require.Equal(t, edgeB.creationTime-edgeA.creationTime, timer)

		// EdgeB and EdgeC were already rivaled at creation, so they should have
		// a local timer of 0 regardless of specified time.
		timer, err = ct.localTimer(edgeB, 100)
		require.NoError(t, err)
		require.Equal(t, uint64(0), timer)
		timer, err = ct.localTimer(edgeC, 100)
		require.NoError(t, err)
		require.Equal(t, uint64(0), timer)
		timer, err = ct.localTimer(edgeB, 10000)
		require.NoError(t, err)
		require.Equal(t, uint64(0), timer)
		timer, err = ct.localTimer(edgeC, 10000)
		require.NoError(t, err)
		require.Equal(t, uint64(0), timer)
	})
}

func Test_earliestCreatedRivalTimestamp(t *testing.T) {
	ct := &challengeTree{
		edges:        threadsafe.NewMap[edgeId, *edge](),
		mutualIds:    threadsafe.NewMap[mutualId, *threadsafe.Set[edgeId]](),
		rivaledEdges: threadsafe.NewSet[edgeId](),
	}
	e := &edge{
		id:           "0-1a",
		creationTime: 3,
	}
	ct.edges.Put("0-1a", e)
	t.Run("no rivals", func(t *testing.T) {
		res := ct.earliestCreatedRivalTimestamp(e)
		require.Equal(t, util.None[uint64](), res)
	})
	t.Run("one rival", func(t *testing.T) {
		ct.rivaledEdges.Insert("0-1a")
		ct.rivaledEdges.Insert("0-1b")
		ct.mutualIds.Put("0-1", threadsafe.NewSet[edgeId]())
		mutuals := ct.mutualIds.Get("0-1")
		mutuals.Insert("0-1a")
		mutuals.Insert("0-1b")
		ct.edges.Put("0-1b", &edge{
			id:           "0-1b",
			creationTime: 5,
		})

		res := ct.earliestCreatedRivalTimestamp(e)
		require.Equal(t, uint64(5), res.Unwrap())
	})
	t.Run("multiple rivals", func(t *testing.T) {
		ct.edges.Put("0-1c", &edge{
			id:           "0-1c",
			creationTime: 10,
		})
		ct.rivaledEdges.Insert("0-1c")
		mutuals := ct.mutualIds.Get("0-1")
		mutuals.Insert("0-1c")

		res := ct.earliestCreatedRivalTimestamp(e)
		require.Equal(t, uint64(5), res.Unwrap())
	})
}

func Test_unrivaledAtTime(t *testing.T) {
	ct := &challengeTree{
		edges:        threadsafe.NewMap[edgeId, *edge](),
		mutualIds:    threadsafe.NewMap[mutualId, *threadsafe.Set[edgeId]](),
		rivaledEdges: threadsafe.NewSet[edgeId](),
	}
	ct.edges.Put("0-1a", &edge{
		id:           "0-1a",
		creationTime: 3,
	})
	t.Run("less than specified time", func(t *testing.T) {
		_, err := ct.unrivaledAtTime(ct.edges.Get("0-1a"), 0)
		require.ErrorContains(t, err, "less than specified time")
	})
	t.Run("no rivals", func(t *testing.T) {
		unrivaled, err := ct.unrivaledAtTime(ct.edges.Get("0-1a"), 3)
		require.NoError(t, err)
		require.Equal(t, true, unrivaled)
		unrivaled, err = ct.unrivaledAtTime(ct.edges.Get("0-1a"), 1000)
		require.NoError(t, err)
		require.Equal(t, true, unrivaled)
	})
	t.Run("with rivals but unrivaled at creation time", func(t *testing.T) {
		ct.rivaledEdges.Insert("0-1a")
		ct.rivaledEdges.Insert("0-1b")
		ct.mutualIds.Put("0-1", threadsafe.NewSet[edgeId]())
		mutuals := ct.mutualIds.Get("0-1")
		mutuals.Insert("0-1a")
		mutuals.Insert("0-1b")
		ct.edges.Put("0-1b", &edge{
			id:           "0-1b",
			creationTime: 5,
		})

		unrivaled, err := ct.unrivaledAtTime(ct.edges.Get("0-1a"), 3)
		require.NoError(t, err)
		require.Equal(t, true, unrivaled)
	})
	t.Run("rivaled at first rival creation time", func(t *testing.T) {
		unrivaled, err := ct.unrivaledAtTime(ct.edges.Get("0-1a"), 5)
		require.NoError(t, err)
		require.Equal(t, false, unrivaled)
		unrivaled, err = ct.unrivaledAtTime(ct.edges.Get("0-1b"), 5)
		require.NoError(t, err)
		require.Equal(t, false, unrivaled)
	})
}

func Test_rivalsWithCreationTimes(t *testing.T) {
	ct := &challengeTree{
		edges:        threadsafe.NewMap[edgeId, *edge](),
		mutualIds:    threadsafe.NewMap[mutualId, *threadsafe.Set[edgeId]](),
		rivaledEdges: threadsafe.NewSet[edgeId](),
	}
	ct.edges.Put("0-1a", &edge{
		id:           "0-1a",
		creationTime: 3,
	})
	t.Run("no rivals", func(t *testing.T) {
		rivals := ct.rivalsWithCreationTimes(ct.edges.Get("0-1a"))
		require.Equal(t, 0, len(rivals))
	})
	t.Run("single rival", func(t *testing.T) {
		ct.rivaledEdges.Insert("0-1a")
		ct.rivaledEdges.Insert("0-1b")
		ct.mutualIds.Put("0-1", threadsafe.NewSet[edgeId]())
		mutuals := ct.mutualIds.Get("0-1")
		mutuals.Insert("0-1a")
		mutuals.Insert("0-1b")
		ct.edges.Put("0-1b", &edge{
			id:           "0-1b",
			creationTime: 5,
		})
		rivals := ct.rivalsWithCreationTimes(ct.edges.Get("0-1a"))
		want := []*rival{
			{id: "0-1b", creationTime: 5},
		}
		require.Equal(t, want, rivals)
		rivals = ct.rivalsWithCreationTimes(ct.edges.Get("0-1b"))
		want = []*rival{
			{id: "0-1a", creationTime: 3},
		}
		require.Equal(t, want, rivals)
	})
	t.Run("multiple rivals", func(t *testing.T) {
		ct.edges.Put("0-1c", &edge{
			id:           "0-1c",
			creationTime: 10,
		})
		ct.rivaledEdges.Insert("0-1c")
		mutuals := ct.mutualIds.Get("0-1")
		mutuals.Insert("0-1c")
		want := []edgeId{"0-1a", "0-1b"}
		rivals := ct.rivalsWithCreationTimes(ct.edges.Get("0-1c"))
		require.Equal(t, true, len(rivals) > 0)
		got := make(map[edgeId]bool)
		for _, r := range rivals {
			got[r.id] = true
		}
		for _, w := range want {
			require.Equal(t, true, got[w])
		}
	})
}

func Test_parents(t *testing.T) {
	ct := &challengeTree{
		edges: threadsafe.NewMap[edgeId, *edge](),
	}
	childId := edgeId("foo")
	t.Run("no parents", func(t *testing.T) {
		parents := ct.parents(childId)
		require.Equal(t, 0, len(parents))
	})
	t.Run("one parent", func(t *testing.T) {
		ct.edges.Put("a", &edge{
			id:           "a",
			lowerChildId: childId,
		})
		parents := ct.parents(childId)
		require.Equal(t, []edgeId{edgeId("a")}, parents)
	})
	t.Run("two parents", func(t *testing.T) {
		ct.edges.Put("b", &edge{
			id:           "b",
			upperChildId: childId,
		})
		parents := ct.parents(childId)
		require.Equal(t, 2, len(parents))
		got := make(map[edgeId]bool)
		for _, p := range parents {
			got[p] = true
		}
		require.Equal(t, true, got["a"])
		require.Equal(t, true, got["b"])
	})
}

func buildEdges(allEdges ...*edge) map[edgeId]*edge {
	m := make(map[edgeId]*edge)
	for _, e := range allEdges {
		m[e.id] = e
	}
	return m
}

func withCreationTime(t *testing.T, origin originId, id edgeId, createdAt uint64) *edge {
	t.Helper()
	items := strings.Split(string(id), "-")
	var typ protocol.EdgeType
	switch items[0] {
	case "blk":
		typ = protocol.BlockChallengeEdge
	case "big":
		typ = protocol.BigStepChallengeEdge
	case "smol":
		typ = protocol.SmallStepChallengeEdge
	}
	startData := strings.Split(items[1], ".")
	startHeight, err := strconv.ParseUint(startData[0], 10, 64)
	require.NoError(t, err)
	startCommit := startData[1]

	endData := strings.Split(items[2], ".")
	endHeight, err := strconv.ParseUint(endData[0], 10, 64)
	require.NoError(t, err)
	endCommit := endData[1]

	return &edge{
		edgeType:     typ,
		originId:     origin,
		id:           id,
		startHeight:  startHeight,
		startCommit:  commit(startCommit),
		endHeight:    endHeight,
		endCommit:    commit(endCommit),
		lowerChildId: "",
		upperChildId: "",
		creationTime: createdAt,
	}
}
