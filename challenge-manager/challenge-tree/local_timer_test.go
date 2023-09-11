// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE

package challengetree

import (
	"testing"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/OffchainLabs/bold/challenge-manager/challenge-tree/mock"
	"github.com/OffchainLabs/bold/containers/option"
	"github.com/OffchainLabs/bold/containers/threadsafe"
	"github.com/stretchr/testify/require"
)

func Test_earliestCreatedRivalBlockNumber(t *testing.T) {
	ct := &HonestChallengeTree{
		edges:     threadsafe.NewMap[protocol.EdgeId, protocol.SpecEdge](),
		mutualIds: threadsafe.NewMap[protocol.MutualId, *threadsafe.Map[protocol.EdgeId, creationTime]](),
	}
	edgeA := newEdge(&newCfg{t: t, edgeId: "blk-0.a-1.a", createdAt: 3})
	edgeB := newEdge(&newCfg{t: t, edgeId: "blk-0.a-1.b", createdAt: 5})
	edgeC := newEdge(&newCfg{t: t, edgeId: "blk-0.a-1.c", createdAt: 10})
	ct.edges.Put(edgeA.Id(), edgeA)
	t.Run("no rivals", func(t *testing.T) {
		res := ct.earliestCreatedRivalBlockNumber(edgeA)

		require.Equal(t, option.None[uint64](), res)
	})
	t.Run("one rival", func(t *testing.T) {
		mutual := edgeA.MutualId()
		ct.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
		mutuals := ct.mutualIds.Get(mutual)
		mutuals.Put(edgeA.Id(), creationTime(edgeA.CreationBlock))
		mutuals.Put(edgeB.Id(), creationTime(edgeB.CreationBlock))
		ct.edges.Put(edgeB.Id(), edgeB)

		res := ct.earliestCreatedRivalBlockNumber(edgeA)

		require.Equal(t, uint64(5), res.Unwrap())
	})
	t.Run("multiple rivals", func(t *testing.T) {
		ct.edges.Put(edgeC.Id(), edgeC)
		mutual := edgeC.MutualId()

		mutuals := ct.mutualIds.Get(mutual)
		mutuals.Put(edgeC.Id(), creationTime(edgeC.CreationBlock))

		res := ct.earliestCreatedRivalBlockNumber(edgeA)

		require.Equal(t, uint64(5), res.Unwrap())
	})
}

func Test_unrivaledAtBlockNum(t *testing.T) {
	ct := &HonestChallengeTree{
		edges:     threadsafe.NewMap[protocol.EdgeId, protocol.SpecEdge](),
		mutualIds: threadsafe.NewMap[protocol.MutualId, *threadsafe.Map[protocol.EdgeId, creationTime]](),
	}
	edgeA := newEdge(&newCfg{t: t, edgeId: "blk-0.a-1.a", createdAt: 3})
	edgeB := newEdge(&newCfg{t: t, edgeId: "blk-0.a-1.b", createdAt: 5})
	ct.edges.Put(edgeA.Id(), edgeA)
	t.Run("less than specified time", func(t *testing.T) {
		_, err := ct.unrivaledAtBlockNum(edgeA, 0)
		require.ErrorContains(t, err, "less than specified")
	})
	t.Run("no rivals", func(t *testing.T) {
		unrivaled, err := ct.unrivaledAtBlockNum(edgeA, 3)
		require.NoError(t, err)
		require.Equal(t, true, unrivaled)
		unrivaled, err = ct.unrivaledAtBlockNum(edgeA, 1000)
		require.NoError(t, err)
		require.Equal(t, true, unrivaled)
	})
	t.Run("with rivals but unrivaled at creation time", func(t *testing.T) {
		mutual := edgeA.MutualId()
		ct.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
		mutuals := ct.mutualIds.Get(mutual)
		mutuals.Put(edgeA.Id(), creationTime(edgeA.CreationBlock))
		mutuals.Put(edgeB.Id(), creationTime(edgeB.CreationBlock))
		ct.edges.Put(edgeB.Id(), edgeB)

		unrivaled, err := ct.unrivaledAtBlockNum(edgeA, 3)
		require.NoError(t, err)
		require.Equal(t, true, unrivaled)
	})
	t.Run("rivaled at first rival creation time", func(t *testing.T) {
		unrivaled, err := ct.unrivaledAtBlockNum(edgeA, 5)
		require.NoError(t, err)
		require.Equal(t, false, unrivaled)
		unrivaled, err = ct.unrivaledAtBlockNum(edgeB, 5)
		require.NoError(t, err)
		require.Equal(t, false, unrivaled)
	})
}

func Test_rivalsWithCreationTimes(t *testing.T) {
	ct := &HonestChallengeTree{
		edges:     threadsafe.NewMap[protocol.EdgeId, protocol.SpecEdge](),
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
		mutuals.Put(edgeB.Id(), creationTime(edgeB.CreationBlock))
		mutuals.Put(edgeA.Id(), creationTime(edgeA.CreationBlock))
		ct.edges.Put(edgeB.Id(), edgeB)
		rivals := ct.rivalsWithCreationTimes(edgeA)

		want := []*rival{
			{id: edgeB.Id(), createdAtBlock: creationTime(edgeB.CreationBlock)},
		}
		require.Equal(t, want, rivals)
		rivals = ct.rivalsWithCreationTimes(edgeB)

		want = []*rival{
			{id: edgeA.Id(), createdAtBlock: creationTime(edgeA.CreationBlock)},
		}
		require.Equal(t, want, rivals)
	})
	t.Run("multiple rivals", func(t *testing.T) {
		ct.edges.Put(edgeC.Id(), edgeC)
		mutual := edgeC.MutualId()
		mutuals := ct.mutualIds.Get(mutual)
		mutuals.Put(edgeC.Id(), creationTime(edgeC.CreationBlock))
		want := []mock.EdgeId{edgeA.ID, edgeB.ID}
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
