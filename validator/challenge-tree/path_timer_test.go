package challengetree

import (
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
		newEdge(&newCfg{t: t, edgeId: "blk-0.a-16.a", createdAt: 1}),
		newEdge(&newCfg{t: t, edgeId: "blk-0.a-8.a", createdAt: 3}),
		newEdge(&newCfg{t: t, edgeId: "blk-8.a-16.a", createdAt: 3}),
		newEdge(&newCfg{t: t, edgeId: "blk-0.a-4.a", createdAt: 5}),
		newEdge(&newCfg{t: t, edgeId: "blk-4.a-8.a", createdAt: 5}),
		// Bob.
		newEdge(&newCfg{t: t, edgeId: "blk-0.a-16.b", createdAt: 2}),
		newEdge(&newCfg{t: t, edgeId: "blk-0.a-8.b", createdAt: 4}),
		newEdge(&newCfg{t: t, edgeId: "blk-8.b-16.b", createdAt: 4}),
		newEdge(&newCfg{t: t, edgeId: "blk-4.a-8.b", createdAt: 6}),
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

	transformedEdges := make(map[protocol.EdgeId]protocol.EdgeSnapshot)
	timers := make(map[protocol.EdgeId]uint64)
	for _, v := range edges {
		transformedEdges[v.Id()] = v
		timers[v.Id()] = 0
	}
	allEdges := threadsafe.NewMapFromItems(transformedEdges)
	allTimers := threadsafe.NewMapFromItems(timers)
	ht := &HonestChallengeTree{
		edges:                      allEdges,
		mutualIds:                  threadsafe.NewMap[protocol.MutualId, *threadsafe.Map[protocol.EdgeId, creationTime]](),
		cumulativeHonestPathTimers: allTimers,
	}
	// Three pairs of edges are rivaled in this test: 0-16, 0-8, and 4-8.
	mutual := edges["blk-0.a-16.a"].MutualId()

	ht.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
	mutuals := ht.mutualIds.Get(mutual)
	idd := id("blk-0.a-16.a")
	mutuals.Put(idd, creationTime(ht.edges.Get(idd).CreatedAtBlock()))
	idd = id("blk-0.a-16.b")
	mutuals.Put(idd, creationTime(ht.edges.Get(idd).CreatedAtBlock()))

	mutual = edges["blk-0.a-8.a"].MutualId()

	ht.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
	mutuals = ht.mutualIds.Get(mutual)
	idd = id("blk-0.a-8.a")
	mutuals.Put(idd, creationTime(ht.edges.Get(idd).CreatedAtBlock()))
	idd = id("blk-0.a-8.b")
	mutuals.Put(idd, creationTime(ht.edges.Get(idd).CreatedAtBlock()))

	mutual = edges["blk-4.a-8.a"].MutualId()

	ht.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
	mutuals = ht.mutualIds.Get(mutual)
	idd = id("blk-4.a-8.a")
	mutuals.Put(idd, creationTime(ht.edges.Get(idd).CreatedAtBlock()))
	idd = id("blk-4.a-8.b")
	mutuals.Put(idd, creationTime(ht.edges.Get(idd).CreatedAtBlock()))

	ht.honestBlockChalLevelZeroEdge = util.Some(ht.edges.Get(id("blk-0.a-16.a")))

	t.Run("querying path timer before creation should return zero", func(t *testing.T) {
		edge := ht.edges.Get(id("blk-0.a-16.a"))
		err := ht.UpdateCumulativePathTimers(edge.CreatedAtBlock() - 1)
		require.NoError(t, err)
		timer, err := ht.HonestPathTimer(edge.Id())
		require.NoError(t, err)
		require.Equal(t, uint64(0), timer)
	})
	t.Run("at creation time should be zero if no parents", func(t *testing.T) {
		edge := ht.edges.Get(id("blk-0.a-16.a"))
		err := ht.UpdateCumulativePathTimers(edge.CreatedAtBlock())
		require.NoError(t, err)
		timer, err := ht.HonestPathTimer(edge.Id())
		require.NoError(t, err)
		require.Equal(t, uint64(0), timer)
	})
	t.Run("OK", func(t *testing.T) {
		// Top-level edge should have spent 1 second unrivaled
		// as its rival was created 1 second after its creation.
		edge := ht.edges.Get(id("blk-0.a-16.a"))
		err := ht.UpdateCumulativePathTimers(edge.CreatedAtBlock() + 1)
		require.NoError(t, err)
		timer, err := ht.HonestPathTimer(edge.Id())
		require.NoError(t, err)
		require.Equal(t, uint64(1), timer)

		// Now we look at the lower honest child, 0.a-8.a. It will have spent
		// 1 second unrivaled and will inherit the local timers of its honest ancestors.
		// which is 1 for a total of 2.
		edge = ht.edges.Get(id("blk-0.a-8.a"))
		err = ht.UpdateCumulativePathTimers(edge.CreatedAtBlock() + 1)
		require.NoError(t, err)
		timer, err = ht.HonestPathTimer(edge.Id())
		require.NoError(t, err)
		require.Equal(t, uint64(2), timer)

		// Now we look at the upper honest grandchild, 4.a-8.a. It will
		// have spent 1 second unrivaled.
		edge = ht.edges.Get(id("blk-4.a-8.a"))
		err = ht.UpdateCumulativePathTimers(edge.CreatedAtBlock() + 1)
		require.NoError(t, err)
		timer, err = ht.HonestPathTimer(edge.Id())
		require.NoError(t, err)
		require.Equal(t, uint64(3), timer)

		// The lower-most child, which is unrivaled, and is 0.a-4.a,
		// will inherit the path timers of its ancestors AND also increase
		// its local timer each time we query it as it has no rival
		// to contend it.
		edge = ht.edges.Get(id("blk-0.a-4.a"))

		// Querying it at creation time+1 should just have the path timers
		// of its ancestors that count, which is a total of 3.
		err = ht.UpdateCumulativePathTimers(edge.CreatedAtBlock() + 1)
		require.NoError(t, err)
		timer, err = ht.HonestPathTimer(edge.Id())
		require.NoError(t, err)
		require.Equal(t, uint64(3), timer)

		// Continuing to query it at time T+i should increase the timer
		// as it is unrivaled.
		for i := uint64(2); i < 10; i++ {
			err = ht.UpdateCumulativePathTimers(edge.CreatedAtBlock() + i)
			require.NoError(t, err)
			require.NoError(t, err)
			timer, err = ht.HonestPathTimer(edge.Id())
			require.NoError(t, err)
			require.Equal(t, uint64(2)+i, timer)
		}
	})
	t.Run("new ancestors created late", func(t *testing.T) {
		// We add a new set of edges that were created late that rival the lower-most,
		// unrivaled honest edge from before. This means that edge will no longer have
		// an ever-increasing unrivaled timer after these new edges are being tracked.
		edges = buildEdges(
			// Charlie.
			newEdge(&newCfg{t: t, edgeId: "blk-0.a-16.c", createdAt: 7}),
			newEdge(&newCfg{t: t, edgeId: "blk-0.a-8.c", createdAt: 8}),
			newEdge(&newCfg{t: t, edgeId: "blk-8.c-16.c", createdAt: 8}),
			newEdge(&newCfg{t: t, edgeId: "blk-4.a-8.c", createdAt: 9}),
		)
		// Child-relationship linking.
		edges["blk-0.a-16.c"].lowerChildId = "blk-0.a-8.c"
		edges["blk-0.a-16.c"].upperChildId = "blk-8.c-16.c"
		edges["blk-0.a-8.c"].lowerChildId = "blk-0.a-4.a"
		edges["blk-0.a-8.c"].upperChildId = "blk-4.a-8.c"

		// Add the new edges into the mapping.
		for k, v := range edges {
			ht.edges.Put(id(k), v)
			ht.cumulativeHonestPathTimers.Put(id(k), 0)
		}

		// // Three pairs of edges are rivaled in this test: 0-16, 0-8, and 4-8.
		// mutual := edges["blk-0.a-16.c"].MutualId()

		// ct.mutualIds.Put(mutual, threadsafe.NewSet[protocol.EdgeId]())
		// mutuals := ct.mutualIds.Get(mutual)
		// mutuals.Insert(id("blk-0.a-16.c"))

		// mutual = edges["blk-0.a-8.c"].MutualId()

		// ct.mutualIds.Put(mutual, threadsafe.NewSet[protocol.EdgeId]())
		// mutuals = ct.mutualIds.Get(mutual)
		// mutuals.Insert(id("blk-0.a-8.c"))

		// mutual = edges["blk-4.a-8.c"].MutualId()

		// ct.mutualIds.Put(mutual, threadsafe.NewSet[protocol.EdgeId]())
		// mutuals = ct.mutualIds.Get(mutual)
		// mutuals.Insert(id("blk-4.a-8.c"))

		// edge := ct.edges.Get(id("blk-0.a-4.a"))
		// lastCreated := ct.edges.Get(id("blk-4.a-8.c"))

		// // The path timers of the newly created edges should count
		// // towards the unrivaled edge at the lowest level.
		// timer, err := ct.pathTimer(edge, lastCreated.CreatedAtBlock())
		// require.NoError(t, err)
		// require.Equal(t, uint64(15), timer)

		// timer, err = ct.pathTimer(edge, lastCreated.CreatedAtBlock()+1)
		// require.NoError(t, err)
		// require.Equal(t, uint64(16), timer)

		// timer, err = ct.pathTimer(edge, lastCreated.CreatedAtBlock()+2)
		// require.NoError(t, err)
		// require.Equal(t, uint64(17), timer)

		// timer, err = ct.pathTimer(edge, lastCreated.CreatedAtBlock()+3)
		// require.NoError(t, err)
		// require.Equal(t, uint64(18), timer)
	})
}

func Test_localTimer(t *testing.T) {
	ct := &HonestChallengeTree{
		edges:     threadsafe.NewMap[protocol.EdgeId, protocol.EdgeSnapshot](),
		mutualIds: threadsafe.NewMap[protocol.MutualId, *threadsafe.Map[protocol.EdgeId, creationTime]](),
	}
	edgeA := newEdge(&newCfg{t: t, edgeId: "blk-0.a-1.a", createdAt: 3})
	ct.edges.Put(edgeA.Id(), edgeA)

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
		edgeB := newEdge(&newCfg{t: t, edgeId: "blk-0.a-1.b", createdAt: 5})
		edgeC := newEdge(&newCfg{t: t, edgeId: "blk-0.a-1.c", createdAt: 10})
		ct.edges.Put(edgeB.Id(), edgeB)
		ct.edges.Put(edgeC.Id(), edgeC)
		mutual := edgeA.MutualId()

		ct.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
		mutuals := ct.mutualIds.Get(mutual)
		mutuals.Put(edgeA.Id(), creationTime(edgeA.creationTime))
		mutuals.Put(edgeB.Id(), creationTime(edgeB.creationTime))
		mutuals.Put(edgeC.Id(), creationTime(edgeC.creationTime))

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
	ct := &HonestChallengeTree{
		edges:     threadsafe.NewMap[protocol.EdgeId, protocol.EdgeSnapshot](),
		mutualIds: threadsafe.NewMap[protocol.MutualId, *threadsafe.Map[protocol.EdgeId, creationTime]](),
	}
	edgeA := newEdge(&newCfg{t: t, edgeId: "blk-0.a-1.a", createdAt: 3})
	edgeB := newEdge(&newCfg{t: t, edgeId: "blk-0.a-1.b", createdAt: 5})
	edgeC := newEdge(&newCfg{t: t, edgeId: "blk-0.a-1.c", createdAt: 10})
	ct.edges.Put(edgeA.Id(), edgeA)
	t.Run("no rivals", func(t *testing.T) {
		res := ct.earliestCreatedRivalTimestamp(edgeA)

		require.Equal(t, util.None[uint64](), res)
	})
	t.Run("one rival", func(t *testing.T) {
		mutual := edgeA.MutualId()
		ct.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
		mutuals := ct.mutualIds.Get(mutual)
		mutuals.Put(edgeA.Id(), creationTime(edgeA.creationTime))
		mutuals.Put(edgeB.Id(), creationTime(edgeB.creationTime))
		ct.edges.Put(edgeB.Id(), edgeB)

		res := ct.earliestCreatedRivalTimestamp(edgeA)

		require.Equal(t, uint64(5), res.Unwrap())
	})
	t.Run("multiple rivals", func(t *testing.T) {
		ct.edges.Put(edgeC.Id(), edgeC)
		mutual := edgeC.MutualId()

		mutuals := ct.mutualIds.Get(mutual)
		mutuals.Put(edgeC.Id(), creationTime(edgeC.creationTime))

		res := ct.earliestCreatedRivalTimestamp(edgeA)

		require.Equal(t, uint64(5), res.Unwrap())
	})
}

func Test_unrivaledAtTime(t *testing.T) {
	ct := &HonestChallengeTree{
		edges:     threadsafe.NewMap[protocol.EdgeId, protocol.EdgeSnapshot](),
		mutualIds: threadsafe.NewMap[protocol.MutualId, *threadsafe.Map[protocol.EdgeId, creationTime]](),
	}
	edgeA := newEdge(&newCfg{t: t, edgeId: "blk-0.a-1.a", createdAt: 3})
	edgeB := newEdge(&newCfg{t: t, edgeId: "blk-0.a-1.b", createdAt: 5})
	ct.edges.Put(edgeA.Id(), edgeA)
	t.Run("less than specified time", func(t *testing.T) {
		_, err := ct.unrivaledAtTime(edgeA, 0)
		require.ErrorContains(t, err, "less than specified")
	})
	t.Run("no rivals", func(t *testing.T) {
		unrivaled, err := ct.unrivaledAtTime(edgeA, 3)
		require.NoError(t, err)
		require.Equal(t, true, unrivaled)
		unrivaled, err = ct.unrivaledAtTime(edgeA, 1000)
		require.NoError(t, err)
		require.Equal(t, true, unrivaled)
	})
	t.Run("with rivals but unrivaled at creation time", func(t *testing.T) {
		mutual := edgeA.MutualId()
		ct.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
		mutuals := ct.mutualIds.Get(mutual)
		mutuals.Put(edgeA.Id(), creationTime(edgeA.creationTime))
		mutuals.Put(edgeB.Id(), creationTime(edgeB.creationTime))
		ct.edges.Put(edgeB.Id(), edgeB)

		unrivaled, err := ct.unrivaledAtTime(edgeA, 3)
		require.NoError(t, err)
		require.Equal(t, true, unrivaled)
	})
	t.Run("rivaled at first rival creation time", func(t *testing.T) {
		unrivaled, err := ct.unrivaledAtTime(edgeA, 5)
		require.NoError(t, err)
		require.Equal(t, false, unrivaled)
		unrivaled, err = ct.unrivaledAtTime(edgeB, 5)
		require.NoError(t, err)
		require.Equal(t, false, unrivaled)
	})
}

func Test_rivalsWithCreationTimes(t *testing.T) {
	ct := &HonestChallengeTree{
		edges:     threadsafe.NewMap[protocol.EdgeId, protocol.EdgeSnapshot](),
		mutualIds: threadsafe.NewMap[protocol.MutualId, *threadsafe.Map[protocol.EdgeId, creationTime]](),
	}
	edgeA := newEdge(&newCfg{t: t, edgeId: "blk-0.a-1.a", createdAt: 5})
	edgeB := newEdge(&newCfg{t: t, edgeId: "blk-0.a-1.b", createdAt: 5})
	edgeC := newEdge(&newCfg{t: t, edgeId: "blk-0.a-1.c", createdAt: 10})
	ct.edges.Put(edgeA.Id(), edgeA)
	t.Run("no rivals", func(t *testing.T) {
		rivals := ct.rivalsWithCreationTimes(edgeA)

		require.Equal(t, 0, len(rivals))
	})
	t.Run("single rival", func(t *testing.T) {
		mutual := edgeA.MutualId()
		ct.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
		mutuals := ct.mutualIds.Get(mutual)
		_ = mutuals
		ct.edges.Put(edgeB.Id(), edgeB)
		rivals := ct.rivalsWithCreationTimes(edgeA)

		want := []*rival{
			{id: edgeB.Id(), createdAtBlock: creationTime(edgeB.creationTime)},
		}
		require.Equal(t, want, rivals)
		rivals = ct.rivalsWithCreationTimes(edgeB)

		want = []*rival{
			{id: edgeA.Id(), createdAtBlock: creationTime(edgeA.creationTime)},
		}
		require.Equal(t, want, rivals)
	})
	t.Run("multiple rivals", func(t *testing.T) {
		ct.edges.Put(edgeC.Id(), edgeC)
		mutual := edgeC.MutualId()
		mutuals := ct.mutualIds.Get(mutual)
		mutuals.Put(edgeC.Id(), creationTime(edgeC.creationTime))
		want := []edgeId{edgeA.id, edgeB.id}
		rivals := ct.rivalsWithCreationTimes(edgeC)

		require.Equal(t, true, len(rivals) > 0)
		got := make(map[protocol.EdgeId]bool)
		for _, r := range rivals {
			got[r.id] = true
		}
		for _, w := range want {
			require.Equal(t, true, got[id(w)])
		}
	})
}
