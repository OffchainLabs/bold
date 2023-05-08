package challengetree

import (
	"strconv"
	"strings"
	"testing"

	"fmt"
	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/OffchainLabs/challenge-protocol-v2/util/threadsafe"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

var _ = protocol.EdgeSnapshot(&edge{})

type edgeId string
type commit string
type originId string

// Mock edge for challenge tree specific tests, making it easier for test ergonomics.
type edge struct {
	id           edgeId
	edgeType     protocol.EdgeType
	startHeight  uint64
	startCommit  commit
	endHeight    uint64
	endCommit    commit
	originId     originId
	lowerChildId edgeId
	upperChildId edgeId
	creationTime uint64
}

func (e *edge) Id() protocol.EdgeId {
	return protocol.EdgeId(common.BytesToHash([]byte(e.id)))
}

func (e *edge) GetType() protocol.EdgeType {
	return e.edgeType
}

func (e *edge) StartCommitment() (protocol.Height, common.Hash) {
	return protocol.Height(e.startHeight), common.BytesToHash([]byte(e.startCommit))
}

func (e *edge) EndCommitment() (protocol.Height, common.Hash) {
	return protocol.Height(e.endHeight), common.BytesToHash([]byte(e.endCommit))
}

func (e *edge) OriginId() protocol.OriginId {
	return protocol.OriginId(common.BytesToHash([]byte(e.originId)))
}

func (e *edge) LowerChildSnapshot() util.Option[protocol.EdgeId] {
	return util.Some(protocol.EdgeId(common.BytesToHash([]byte(e.lowerChildId))))
}

func (e *edge) UpperChildSnapshot() util.Option[protocol.EdgeId] {
	return util.Some(protocol.EdgeId(common.BytesToHash([]byte(e.upperChildId))))
}

func (e *edge) CreatedAtBlock() uint64 {
	return e.creationTime
}

func (e *edge) MutualId() protocol.MutualId {
	return protocol.MutualId(common.BytesToHash([]byte(fmt.Sprintf(
		"%d-%s-%d-%s-%d",
		e.edgeType,
		e.originId,
		e.startHeight,
		e.startCommit,
		e.endHeight,
	))))
}

// Simple function to go from a mock string edge id to a real protocol EdgeId type.
func id(i edgeId) protocol.EdgeId {
	return protocol.EdgeId(common.BytesToHash([]byte(i)))
}

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
	for _, v := range edges {
		transformedEdges[v.Id()] = v
	}
	allEdges := threadsafe.NewMapFromItems(transformedEdges)
	ct := &challengeTree{
		edges:        allEdges,
		mutualIds:    threadsafe.NewMap[protocol.MutualId, *threadsafe.Set[protocol.EdgeId]](),
		rivaledEdges: threadsafe.NewSet[protocol.EdgeId](),
	}

	// We then set up the rival relationships in the challenge tree.
	// All edges are rivaled in this example.
	for _, e := range edges {
		ct.rivaledEdges.Insert(e.Id())
	}

	// Three pairs of edges are rivaled in this test: 0-16, 0-8, and 4-8.
	mutual := edges["blk-0.a-16.a"].MutualId()

	ct.mutualIds.Put(mutual, threadsafe.NewSet[protocol.EdgeId]())
	mutuals := ct.mutualIds.Get(mutual)
	mutuals.Insert(id("blk-0.a-16.a"))
	mutuals.Insert(id("blk-0.a-16.b"))

	mutual = edges["blk-0.a-8.a"].MutualId()

	ct.mutualIds.Put(mutual, threadsafe.NewSet[protocol.EdgeId]())
	mutuals = ct.mutualIds.Get(mutual)
	mutuals.Insert(id("blk-0.a-8.a"))
	mutuals.Insert(id("blk-0.a-8.b"))

	mutual = edges["blk-4.a-8.a"].MutualId()

	ct.mutualIds.Put(mutual, threadsafe.NewSet[protocol.EdgeId]())
	mutuals = ct.mutualIds.Get(mutual)
	mutuals.Insert(id("blk-4.a-8.a"))
	mutuals.Insert(id("blk-4.a-8.b"))

	t.Run("querying path timer before creation should return zero", func(t *testing.T) {
		edge := ct.edges.Get(id("blk-0.a-16.a"))
		timer, err := ct.pathTimer(edge, edge.CreatedAtBlock()-1)
		require.NoError(t, err)
		require.Equal(t, uint64(0), timer)
	})
	t.Run("at creation time should be zero if no parents", func(t *testing.T) {
		edge := ct.edges.Get(id("blk-0.a-16.a"))
		timer, err := ct.pathTimer(edge, edge.CreatedAtBlock())
		require.NoError(t, err)
		require.Equal(t, uint64(0), timer)
	})
	t.Run("OK", func(t *testing.T) {
		// Top-level edge should have spent 1 second unrivaled
		// as its rival was created 1 second after its creation.
		edge := ct.edges.Get(id("blk-0.a-16.a"))
		timer, err := ct.pathTimer(edge, edge.CreatedAtBlock()+1)
		require.NoError(t, err)
		require.Equal(t, uint64(1), timer)

		// Its rival should have a timer of 0 as was rivaled on creation.
		edge = ct.edges.Get(id("blk-0.a-16.b"))
		timer, err = ct.pathTimer(edge, edge.CreatedAtBlock()+1)
		require.NoError(t, err)
		require.Equal(t, uint64(0), timer)

		// Now we look at the lower honest child, 0.a-8.a. It will have spent
		// 1 second unrivaled and will inherit the max local timer
		// of its parents, which is 1 for a total of 2.
		edge = ct.edges.Get(id("blk-0.a-8.a"))
		timer, err = ct.pathTimer(edge, edge.CreatedAtBlock()+1)
		require.NoError(t, err)
		require.Equal(t, uint64(2), timer)

		// Its rival will have a timer of 0 as was rivaled on creation.
		edge = ct.edges.Get(id("blk-0.a-8.b"))
		timer, err = ct.pathTimer(edge, edge.CreatedAtBlock()+1)
		require.NoError(t, err)
		require.Equal(t, uint64(0), timer)

		// Now we look at the upper honest grandchild, 4.a-8.a. It will have spent
		// 1 second unrivaled.
		edge = ct.edges.Get(id("blk-4.a-8.a"))
		timer, err = ct.pathTimer(edge, edge.CreatedAtBlock()+1)
		require.NoError(t, err)
		require.Equal(t, uint64(3), timer)

		// Its rival will have a timer of 0 as was rivaled on creation.
		edge = ct.edges.Get(id("blk-4.a-8.b"))
		timer, err = ct.pathTimer(edge, edge.CreatedAtBlock()+1)
		require.NoError(t, err)
		require.Equal(t, uint64(0), timer)

		// The lower-most child, which is unrivaled, and is 0.a-4.a,
		// will inherit the path timers of its ancestors AND also increase
		// its local timer each time we query it as it has no rival
		// to contend it.
		edge = ct.edges.Get(id("blk-0.a-4.a"))

		// Querying it at creation time+1 should just have the path timers
		// of its ancestors that count, which is a total of 3.
		timer, err = ct.pathTimer(edge, edge.CreatedAtBlock()+1)
		require.NoError(t, err)
		require.Equal(t, uint64(3), timer)

		// Continuing to query it at time T+i should increase the timer
		// as it is unrivaled.
		for i := uint64(2); i < 10; i++ {
			timer, err = ct.pathTimer(edge, edge.CreatedAtBlock()+i)
			require.NoError(t, err)
			require.Equal(t, uint64(2)+i, timer)
		}
	})
	t.Run("new ancestors created late", func(t *testing.T) {
		// We add a new set of edges that were created late. These will
		// not count towards the path timers of the honest branch
		// as the path timer function will only consider the earliest
		// created rival.
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
			ct.edges.Put(id(k), v)
		}

		// We then set up the rival relationships in the challenge tree.
		// All edges are rivaled in this example.
		for _, e := range edges {
			ct.rivaledEdges.Insert(e.Id())
		}

		// Three pairs of edges are rivaled in this test: 0-16, 0-8, and 4-8.
		mutual := edges["blk-0.a-16.c"].MutualId()

		ct.mutualIds.Put(mutual, threadsafe.NewSet[protocol.EdgeId]())
		mutuals := ct.mutualIds.Get(mutual)
		mutuals.Insert(id("blk-0.a-16.c"))

		mutual = edges["blk-0.a-8.c"].MutualId()

		ct.mutualIds.Put(mutual, threadsafe.NewSet[protocol.EdgeId]())
		mutuals = ct.mutualIds.Get(mutual)
		mutuals.Insert(id("blk-0.a-8.c"))

		mutual = edges["blk-4.a-8.c"].MutualId()

		ct.mutualIds.Put(mutual, threadsafe.NewSet[protocol.EdgeId]())
		mutuals = ct.mutualIds.Get(mutual)
		mutuals.Insert(id("blk-4.a-8.c"))

		edge := ct.edges.Get(id("blk-0.a-4.a"))
		lastCreated := ct.edges.Get(id("blk-4.a-8.c"))

		// The path timers of the newly created edges should count
		// towards the unrivaled edge at the lowest level.
		timer, err := ct.pathTimer(edge, lastCreated.CreatedAtBlock())
		require.NoError(t, err)
		require.Equal(t, uint64(15), timer)

		timer, err = ct.pathTimer(edge, lastCreated.CreatedAtBlock()+1)
		require.NoError(t, err)
		require.Equal(t, uint64(16), timer)

		timer, err = ct.pathTimer(edge, lastCreated.CreatedAtBlock()+2)
		require.NoError(t, err)
		require.Equal(t, uint64(17), timer)

		timer, err = ct.pathTimer(edge, lastCreated.CreatedAtBlock()+3)
		require.NoError(t, err)
		require.Equal(t, uint64(18), timer)
	})
}

func Test_localTimer(t *testing.T) {
	ct := &challengeTree{
		edges:        threadsafe.NewMap[protocol.EdgeId, protocol.EdgeSnapshot](),
		mutualIds:    threadsafe.NewMap[protocol.MutualId, *threadsafe.Set[protocol.EdgeId]](),
		rivaledEdges: threadsafe.NewSet[protocol.EdgeId](),
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
		ct.rivaledEdges.Insert(edgeA.Id())
		ct.rivaledEdges.Insert(edgeB.Id())
		ct.rivaledEdges.Insert(edgeC.Id())
		mutual := edgeA.MutualId()

		ct.mutualIds.Put(mutual, threadsafe.NewSet[protocol.EdgeId]())
		mutuals := ct.mutualIds.Get(mutual)
		mutuals.Insert(edgeA.Id())
		mutuals.Insert(edgeB.Id())
		mutuals.Insert(edgeC.Id())

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
		edges:        threadsafe.NewMap[protocol.EdgeId, protocol.EdgeSnapshot](),
		mutualIds:    threadsafe.NewMap[protocol.MutualId, *threadsafe.Set[protocol.EdgeId]](),
		rivaledEdges: threadsafe.NewSet[protocol.EdgeId](),
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
		ct.rivaledEdges.Insert(edgeA.Id())
		ct.rivaledEdges.Insert(edgeB.Id())
		mutual := edgeA.MutualId()
		ct.mutualIds.Put(mutual, threadsafe.NewSet[protocol.EdgeId]())
		mutuals := ct.mutualIds.Get(mutual)
		mutuals.Insert(edgeA.Id())
		mutuals.Insert(edgeB.Id())
		ct.edges.Put(edgeB.Id(), edgeB)

		res := ct.earliestCreatedRivalTimestamp(edgeA)

		require.Equal(t, uint64(5), res.Unwrap())
	})
	t.Run("multiple rivals", func(t *testing.T) {
		ct.edges.Put(edgeC.Id(), edgeC)
		ct.rivaledEdges.Insert(edgeC.Id())
		mutual := edgeC.MutualId()

		mutuals := ct.mutualIds.Get(mutual)
		mutuals.Insert(edgeC.Id())

		res := ct.earliestCreatedRivalTimestamp(edgeA)

		require.Equal(t, uint64(5), res.Unwrap())
	})
}

func Test_unrivaledAtTime(t *testing.T) {
	ct := &challengeTree{
		edges:        threadsafe.NewMap[protocol.EdgeId, protocol.EdgeSnapshot](),
		mutualIds:    threadsafe.NewMap[protocol.MutualId, *threadsafe.Set[protocol.EdgeId]](),
		rivaledEdges: threadsafe.NewSet[protocol.EdgeId](),
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
		ct.rivaledEdges.Insert(edgeA.Id())
		ct.rivaledEdges.Insert(edgeB.Id())
		mutual := edgeA.MutualId()
		ct.mutualIds.Put(mutual, threadsafe.NewSet[protocol.EdgeId]())
		mutuals := ct.mutualIds.Get(mutual)
		mutuals.Insert(edgeA.Id())
		mutuals.Insert(edgeB.Id())
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
	ct := &challengeTree{
		edges:        threadsafe.NewMap[protocol.EdgeId, protocol.EdgeSnapshot](),
		mutualIds:    threadsafe.NewMap[protocol.MutualId, *threadsafe.Set[protocol.EdgeId]](),
		rivaledEdges: threadsafe.NewSet[protocol.EdgeId](),
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
		ct.rivaledEdges.Insert(edgeA.Id())
		ct.rivaledEdges.Insert(edgeB.Id())
		mutual := edgeA.MutualId()
		ct.mutualIds.Put(mutual, threadsafe.NewSet[protocol.EdgeId]())
		mutuals := ct.mutualIds.Get(mutual)
		mutuals.Insert(edgeA.Id())
		mutuals.Insert(edgeB.Id())
		ct.edges.Put(edgeB.Id(), edgeB)
		rivals := ct.rivalsWithCreationTimes(edgeA)

		want := []*rival{
			{id: edgeB.Id(), creationTime: edgeB.creationTime},
		}
		require.Equal(t, want, rivals)
		rivals = ct.rivalsWithCreationTimes(edgeB)

		want = []*rival{
			{id: edgeA.Id(), creationTime: edgeA.creationTime},
		}
		require.Equal(t, want, rivals)
	})
	t.Run("multiple rivals", func(t *testing.T) {
		ct.edges.Put(edgeC.Id(), edgeC)
		ct.rivaledEdges.Insert(edgeC.Id())
		mutual := edgeC.MutualId()
		mutuals := ct.mutualIds.Get(mutual)
		mutuals.Insert(edgeC.Id())
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

func Test_parents(t *testing.T) {
	ct := &challengeTree{
		edges: threadsafe.NewMap[protocol.EdgeId, protocol.EdgeSnapshot](),
	}
	childId := id("foo")
	t.Run("no parents", func(t *testing.T) {
		parents := ct.parents(childId)

		require.Equal(t, 0, len(parents))
	})
	t.Run("one parent", func(t *testing.T) {
		ct.edges.Put(id("a"), &edge{
			id:           "a",
			lowerChildId: "foo",
		})
		parents := ct.parents(childId)

		require.Equal(t, []protocol.EdgeId{id("a")}, parents)
	})
	t.Run("two parents", func(t *testing.T) {
		ct.edges.Put(id("b"), &edge{
			id:           "b",
			upperChildId: "foo",
		})
		parents := ct.parents(childId)

		require.Equal(t, 2, len(parents))
		got := make(map[protocol.EdgeId]bool)
		for _, p := range parents {
			got[p] = true
		}
		require.Equal(t, true, got[id("a")])
		require.Equal(t, true, got[id("b")])
	})
}

func buildEdges(allEdges ...*edge) map[edgeId]*edge {
	m := make(map[edgeId]*edge)
	for _, e := range allEdges {
		m[e.id] = e
	}
	return m
}

type newCfg struct {
	t         *testing.T
	originId  originId
	edgeId    edgeId
	createdAt uint64
}

func newEdge(cfg *newCfg) *edge {
	cfg.t.Helper()
	items := strings.Split(string(cfg.edgeId), "-")
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
	require.NoError(cfg.t, err)
	startCommit := startData[1]

	endData := strings.Split(items[2], ".")
	endHeight, err := strconv.ParseUint(endData[0], 10, 64)
	require.NoError(cfg.t, err)
	endCommit := endData[1]

	return &edge{
		edgeType:     typ,
		originId:     cfg.originId,
		id:           cfg.edgeId,
		startHeight:  startHeight,
		startCommit:  commit(startCommit),
		endHeight:    endHeight,
		endCommit:    commit(endCommit),
		lowerChildId: "",
		upperChildId: "",
		creationTime: cfg.createdAt,
	}
}
