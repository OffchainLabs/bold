// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE

package challengetree

import (
	"context"
	"testing"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/OffchainLabs/bold/challenge-manager/challenge-tree/mock"
	"github.com/OffchainLabs/bold/containers/option"
	"github.com/OffchainLabs/bold/containers/threadsafe"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

// Tests the following tree, all the way down to the small
// step subchallenge level.
//
//		Block challenge:
//
//			      /--5---6-----8-----------16 = Alice
//			0-----4
//			      \--5'--6'----8'----------16' = Bob
//
//		Big step challenge:
//
//			      /--5---6-----8-----------16 = Alice (claim_id = id(5, 6) in the level above)
//			0-----4
//			      \--5'--6'----8'----------16' = Bob
//
//	 Small step challenge:
//
//			      /--5---6-----8-----------16 = Alice (claim_id = id(5, 6) in the level above)
//			0-----4
//			      \--5'--6'----8'----------16' = Bob
//
// From here, the list of ancestors can be determined all the way to the top across
// challenge levels successfully, linked by claimed edges.
func TestComputeAncestorsWithTimers(t *testing.T) {
	ctx := context.Background()
	tree := &HonestChallengeTree{
		edges:                         threadsafe.NewMap[protocol.EdgeId, protocol.SpecEdge](),
		mutualIds:                     threadsafe.NewMap[protocol.MutualId, *threadsafe.Map[protocol.EdgeId, creationTime]](),
		honestBigStepLevelZeroEdges:   threadsafe.NewSlice[protocol.ReadOnlyEdge](),
		honestSmallStepLevelZeroEdges: threadsafe.NewSlice[protocol.ReadOnlyEdge](),
		metadataReader:                &mockMetadataReader{},
		totalChallengeLevels:          3,
		honestRootEdgesByLevel:        threadsafe.NewMap[protocol.ChallengeLevel, *threadsafe.Slice[protocol.ReadOnlyEdge]](),
	}
	tree.honestRootEdgesByLevel.Put(2, threadsafe.NewSlice[protocol.ReadOnlyEdge]())
	tree.honestRootEdgesByLevel.Put(1, threadsafe.NewSlice[protocol.ReadOnlyEdge]())
	tree.honestRootEdgesByLevel.Put(0, threadsafe.NewSlice[protocol.ReadOnlyEdge]())

	// Edge ids that belong to block challenges are prefixed with "blk".
	// For big step, prefixed with "big", and small step, prefixed with "smol".
	setupBlockChallengeTreeSnapshot(t, tree)
	tree.honestBlockChalLevelZeroEdge = option.Some(protocol.ReadOnlyEdge(tree.edges.Get(id("blk-0.a-16.a"))))
	claimId := "blk-4.a-5.a"
	setupBigStepChallengeSnapshot(t, tree, claimId)
	bigStepRootEdges := tree.honestRootEdgesByLevel.Get(1 /* big step level */)
	bigStepRootEdges.Push(tree.edges.Get(id("big-0.a-16.a")))
	claimId = "big-4.a-5.a"
	setupSmallStepChallengeSnapshot(t, tree, claimId)
	smallStepRootEdges := tree.honestRootEdgesByLevel.Get(0 /* small step level */)
	smallStepRootEdges.Push(tree.edges.Get(id("smol-0.a-16.a")))
	blockNum := uint64(30)

	t.Run("junk edge fails", func(t *testing.T) {
		// We start by querying for ancestors for a block edge id.
		_, err := tree.ComputeAncestorsWithTimers(ctx, id("foo"), blockNum)
		require.ErrorContains(t, err, "not found in honest challenge tree")
	})
	t.Run("dishonest edge lookup fails", func(t *testing.T) {
		_, err := tree.ComputeAncestorsWithTimers(ctx, id("blk-0.a-16.b"), blockNum)
		require.ErrorContains(t, err, "not found in honest challenge tree")
	})
	t.Run("block challenge: level zero edge has no ancestors", func(t *testing.T) {
		resp, err := tree.ComputeAncestorsWithTimers(ctx, id("blk-0.a-16.a"), blockNum)
		require.NoError(t, err)
		require.Equal(t, 0, len(resp.AncestorEdgeIds))
	})
	t.Run("block challenge: single ancestor", func(t *testing.T) {
		resp, err := tree.ComputeAncestorsWithTimers(ctx, id("blk-0.a-8.a"), blockNum)
		require.NoError(t, err)
		require.Equal(t, HonestAncestors{id("blk-0.a-16.a")}, resp.AncestorEdgeIds)
		resp, err = tree.ComputeAncestorsWithTimers(ctx, id("blk-8.a-16.a"), blockNum)
		require.NoError(t, err)
		require.Equal(t, HonestAncestors{id("blk-0.a-16.a")}, resp.AncestorEdgeIds)
	})
	t.Run("block challenge: many ancestors", func(t *testing.T) {
		resp, err := tree.ComputeAncestorsWithTimers(ctx, id("blk-4.a-5.a"), blockNum)
		require.NoError(t, err)
		wanted := HonestAncestors{
			id("blk-4.a-6.a"),
			id("blk-4.a-8.a"),
			id("blk-0.a-8.a"),
			id("blk-0.a-16.a"),
		}
		require.Equal(t, wanted, resp.AncestorEdgeIds)
	})
	t.Run("big step challenge: level zero edge has ancestors from block challenge", func(t *testing.T) {
		resp, err := tree.ComputeAncestorsWithTimers(ctx, id("big-0.a-16.a"), blockNum)
		require.NoError(t, err)
		wanted := HonestAncestors{
			id("blk-4.a-5.a"),
			id("blk-4.a-6.a"),
			id("blk-4.a-8.a"),
			id("blk-0.a-8.a"),
			id("blk-0.a-16.a"),
		}
		for i, ancestor := range resp.AncestorEdgeIds {
			if ancestor != wanted[i] {
				t.Errorf("ancestor %d: wanted %s, got %s", i, wanted[i].Bytes(), ancestor.Bytes())
			}
		}
	})
	t.Run("big step challenge: many ancestors plus block challenge ancestors", func(t *testing.T) {
		resp, err := tree.ComputeAncestorsWithTimers(ctx, id("big-5.a-6.a"), blockNum)
		require.NoError(t, err)
		wanted := HonestAncestors{
			// Big step chal.
			id("big-4.a-6.a"),
			id("big-4.a-8.a"),
			id("big-0.a-8.a"),
			id("big-0.a-16.a"),
			// Block chal.
			id("blk-4.a-5.a"),
			id("blk-4.a-6.a"),
			id("blk-4.a-8.a"),
			id("blk-0.a-8.a"),
			id("blk-0.a-16.a"),
		}
		require.Equal(t, wanted, resp.AncestorEdgeIds)
	})
	t.Run("small step challenge: level zero edge has ancestors from big and block challenge", func(t *testing.T) {
		resp, err := tree.ComputeAncestorsWithTimers(ctx, id("smol-0.a-16.a"), blockNum)
		require.NoError(t, err)
		wanted := HonestAncestors{
			// Big step chal.
			id("big-4.a-5.a"),
			id("big-4.a-6.a"),
			id("big-4.a-8.a"),
			id("big-0.a-8.a"),
			id("big-0.a-16.a"),
			// Block chal.
			id("blk-4.a-5.a"),
			id("blk-4.a-6.a"),
			id("blk-4.a-8.a"),
			id("blk-0.a-8.a"),
			id("blk-0.a-16.a"),
		}
		require.Equal(t, wanted, resp.AncestorEdgeIds)
	})
	t.Run("small step challenge: lowest level edge has full ancestry", func(t *testing.T) {
		resp, err := tree.ComputeAncestorsWithTimers(ctx, id("smol-5.a-6.a"), blockNum)
		require.NoError(t, err)
		wanted := HonestAncestors{
			// Small step chal.
			id("smol-4.a-6.a"),
			id("smol-4.a-8.a"),
			id("smol-0.a-8.a"),
			id("smol-0.a-16.a"),
			// Big step chal.
			id("big-4.a-5.a"),
			id("big-4.a-6.a"),
			id("big-4.a-8.a"),
			id("big-0.a-8.a"),
			id("big-0.a-16.a"),
			// Block chal.
			id("blk-4.a-5.a"),
			id("blk-4.a-6.a"),
			id("blk-4.a-8.a"),
			id("blk-0.a-8.a"),
			id("blk-0.a-16.a"),
		}
		require.Equal(t, wanted, resp.AncestorEdgeIds)
	})
}

func TestHasConfirmableAncestor(t *testing.T) {
	ctx := context.Background()
	challengePeriodBlocks := uint64(10)
	t.Run("empty ancestor timers", func(t *testing.T) {
		ht := &HonestChallengeTree{
			metadataReader: &mockMetadataReader{},
		}
		has, err := ht.HasConfirmableAncestor(ctx, nil, challengePeriodBlocks)
		require.NoError(t, err)
		require.Equal(t, false, has)
	})
	t.Run("single ancestor not enough timer", func(t *testing.T) {
		ht := &HonestChallengeTree{
			metadataReader: &mockMetadataReader{},
		}
		has, err := ht.HasConfirmableAncestor(
			ctx,
			[]EdgeLocalTimer{
				EdgeLocalTimer(challengePeriodBlocks) - 1,
			},
			challengePeriodBlocks,
		)
		require.NoError(t, err)
		require.Equal(t, false, has)
	})
	t.Run("single ancestor timer plus assertion unrivaled blocks is enough to be confirmable", func(t *testing.T) {
		ht := &HonestChallengeTree{
			metadataReader: &mockMetadataReader{
				unrivaledAssertionBlocks: 1,
			},
		}
		has, err := ht.HasConfirmableAncestor(
			ctx,
			[]EdgeLocalTimer{
				EdgeLocalTimer(challengePeriodBlocks) - 1,
			},
			challengePeriodBlocks,
		)
		require.NoError(t, err)
		require.Equal(t, true, has)
	})
	t.Run("multiple ancestor timers not enough", func(t *testing.T) {
		ht := &HonestChallengeTree{
			metadataReader: &mockMetadataReader{
				unrivaledAssertionBlocks: 0,
			},
		}
		has, err := ht.HasConfirmableAncestor(
			ctx,
			[]EdgeLocalTimer{
				1,
				2,
				3,
				3,
			}, // Total of 9, just shy of 10.
			challengePeriodBlocks,
		)
		require.NoError(t, err)
		require.Equal(t, false, has)
	})
	t.Run("multiple ancestor timers plus assertion unrivaled blocks enough to be confirmable", func(t *testing.T) {
		ht := &HonestChallengeTree{
			metadataReader: &mockMetadataReader{
				unrivaledAssertionBlocks: 1,
			},
		}
		has, err := ht.HasConfirmableAncestor(
			ctx,
			[]EdgeLocalTimer{
				1,
				2,
				3,
				3,
			}, // Total of 10 including the unrivaled assertion blocks.
			challengePeriodBlocks,
		)
		require.NoError(t, err)
		require.Equal(t, true, has)
	})
	t.Run("many ancestor timers is not enough", func(t *testing.T) {
		ht := &HonestChallengeTree{
			metadataReader: &mockMetadataReader{
				unrivaledAssertionBlocks: 0,
			},
		}
		has, err := ht.HasConfirmableAncestor(
			ctx,
			[]EdgeLocalTimer{
				1,
				2,
				3,
				4,
				5,
				6,
				7,
				8,
				9,
				10, // Sum of 55, not enough.
			},
			56, /* challenge period blocks */
		)
		require.NoError(t, err)
		require.Equal(t, false, has)
	})
	t.Run("all ancestors are confirmable", func(t *testing.T) {
		ht := &HonestChallengeTree{
			metadataReader: &mockMetadataReader{
				unrivaledAssertionBlocks: 0,
			},
		}
		has, err := ht.HasConfirmableAncestor(
			ctx,
			[]EdgeLocalTimer{
				10,
				20,
				30,
				40,
				50,
				60,
			},
			challengePeriodBlocks,
		)
		require.NoError(t, err)
		require.Equal(t, true, has)
	})
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
func TestComputeHonestPathTimer(t *testing.T) {
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
	edges["blk-0.a-16.a"].LowerChildID = "blk-0.a-8.a"
	edges["blk-0.a-16.a"].UpperChildID = "blk-8.a-16.a"
	edges["blk-0.a-8.a"].LowerChildID = "blk-0.a-4.a"
	edges["blk-0.a-8.a"].UpperChildID = "blk-4.a-8.a"
	// Bob.
	edges["blk-0.a-16.b"].LowerChildID = "blk-0.a-8.b"
	edges["blk-0.a-16.b"].UpperChildID = "blk-8.b-16.b"
	edges["blk-0.a-8.b"].LowerChildID = "blk-0.a-4.a"
	edges["blk-0.a-8.b"].UpperChildID = "blk-4.a-8.b"

	transformedEdges := make(map[protocol.EdgeId]protocol.SpecEdge)
	timers := make(map[protocol.EdgeId]uint64)
	for _, v := range edges {
		transformedEdges[v.Id()] = v
		timers[v.Id()] = 0
	}
	allEdges := threadsafe.NewMapFromItems(transformedEdges)
	ht := &HonestChallengeTree{
		edges:                         allEdges,
		mutualIds:                     threadsafe.NewMap[protocol.MutualId, *threadsafe.Map[protocol.EdgeId, creationTime]](),
		honestBigStepLevelZeroEdges:   threadsafe.NewSlice[protocol.ReadOnlyEdge](),
		honestSmallStepLevelZeroEdges: threadsafe.NewSlice[protocol.ReadOnlyEdge](),
		metadataReader:                &mockMetadataReader{},
		totalChallengeLevels:          3,
		honestRootEdgesByLevel:        threadsafe.NewMap[protocol.ChallengeLevel, *threadsafe.Slice[protocol.ReadOnlyEdge]](),
	}
	ht.honestRootEdgesByLevel.Put(2, threadsafe.NewSlice[protocol.ReadOnlyEdge]())
	ht.honestRootEdgesByLevel.Put(1, threadsafe.NewSlice[protocol.ReadOnlyEdge]())
	ht.honestRootEdgesByLevel.Put(0, threadsafe.NewSlice[protocol.ReadOnlyEdge]())

	// Three pairs of edges are rivaled in this test: 0-16, 0-8, and 4-8.
	mutual := edges["blk-0.a-16.a"].MutualId()

	ht.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
	mutuals := ht.mutualIds.Get(mutual)
	idd := id("blk-0.a-16.a")
	iddCreation, err := ht.edges.Get(idd).CreatedAtBlock()
	require.NoError(t, err)
	mutuals.Put(idd, creationTime(iddCreation))
	idd = id("blk-0.a-16.b")
	iddCreation, err = ht.edges.Get(idd).CreatedAtBlock()
	require.NoError(t, err)
	mutuals.Put(idd, creationTime(iddCreation))

	mutual = edges["blk-0.a-8.a"].MutualId()

	ht.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
	mutuals = ht.mutualIds.Get(mutual)
	idd = id("blk-0.a-8.a")
	iddCreation, err = ht.edges.Get(idd).CreatedAtBlock()
	require.NoError(t, err)
	mutuals.Put(idd, creationTime(iddCreation))
	idd = id("blk-0.a-8.b")
	iddCreation, err = ht.edges.Get(idd).CreatedAtBlock()
	require.NoError(t, err)
	mutuals.Put(idd, creationTime(iddCreation))

	mutual = edges["blk-4.a-8.a"].MutualId()

	ht.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
	mutuals = ht.mutualIds.Get(mutual)
	idd = id("blk-4.a-8.a")
	iddCreation, err = ht.edges.Get(idd).CreatedAtBlock()
	require.NoError(t, err)
	mutuals.Put(idd, creationTime(iddCreation))
	idd = id("blk-4.a-8.b")
	iddCreation, err = ht.edges.Get(idd).CreatedAtBlock()
	require.NoError(t, err)
	mutuals.Put(idd, creationTime(iddCreation))

	ht.honestBlockChalLevelZeroEdge = option.Some(protocol.ReadOnlyEdge(ht.edges.Get(id("blk-0.a-16.a"))))
	ctx := context.Background()

	t.Run("querying path timer before creation should return zero", func(t *testing.T) {
		edge := ht.edges.Get(id("blk-0.a-16.a"))
		resp, respErr := ht.ComputeAncestorsWithTimers(ctx, edge.Id(), 0)
		require.NoError(t, respErr)
		timer, pathErr := ht.ComputeHonestPathTimer(ctx, edge.Id(), resp.AncestorLocalTimers, 0)
		require.NoError(t, pathErr)
		require.Equal(t, PathTimer(0), timer)
	})
	t.Run("at creation time should be zero if no parents", func(t *testing.T) {
		edge := ht.edges.Get(id("blk-0.a-16.a"))
		creation, createErr := edge.CreatedAtBlock()
		require.NoError(t, createErr)
		resp, respErr := ht.ComputeAncestorsWithTimers(ctx, edge.Id(), creation)
		require.NoError(t, respErr)
		timer, timeErr := ht.ComputeHonestPathTimer(ctx, edge.Id(), resp.AncestorLocalTimers, creation)
		require.NoError(t, timeErr)
		require.Equal(t, PathTimer(0), timer)
	})
	t.Run("OK", func(t *testing.T) {
		// Top-level edge should have spent 1 second unrivaled
		// as its rival was created 1 second after its creation.
		edge := ht.edges.Get(id("blk-0.a-16.a"))
		creation, createErr := edge.CreatedAtBlock()
		require.NoError(t, createErr)
		resp, respErr := ht.ComputeAncestorsWithTimers(ctx, edge.Id(), creation+1)
		require.NoError(t, respErr)
		timer, timeErr := ht.ComputeHonestPathTimer(ctx, edge.Id(), resp.AncestorLocalTimers, creation+1)
		require.NoError(t, timeErr)
		require.Equal(t, PathTimer(1), timer)

		// Now we look at the lower honest child, 0.a-8.a. It will have spent
		// 1 second unrivaled and will inherit the local timers of its honest ancestors.
		// which is 1 for a total of 2.
		edge = ht.edges.Get(id("blk-0.a-8.a"))
		creation, err = edge.CreatedAtBlock()
		require.NoError(t, err)
		resp, err = ht.ComputeAncestorsWithTimers(ctx, edge.Id(), creation+1)
		require.NoError(t, err)
		timer, err = ht.ComputeHonestPathTimer(ctx, edge.Id(), resp.AncestorLocalTimers, creation+1)
		require.NoError(t, err)
		require.Equal(t, PathTimer(2), timer)

		// Now we look at the upper honest grandchild, 4.a-8.a. It will
		// have spent 1 second unrivaled.
		edge = ht.edges.Get(id("blk-4.a-8.a"))
		creation, err = edge.CreatedAtBlock()
		require.NoError(t, err)
		resp, err = ht.ComputeAncestorsWithTimers(ctx, edge.Id(), creation+1)
		require.NoError(t, err)
		timer, err = ht.ComputeHonestPathTimer(ctx, edge.Id(), resp.AncestorLocalTimers, creation+1)
		require.NoError(t, err)
		require.Equal(t, PathTimer(3), timer)

		// The lower-most child, which is unrivaled, and is 0.a-4.a,
		// will inherit the path timers of its ancestors AND also increase
		// its local timer each time we query it as it has no rival
		// to contend it.
		edge = ht.edges.Get(id("blk-0.a-4.a"))

		// Querying it at creation time+1 should just have the path timers
		// of its ancestors that count, which is a total of 3.
		creation, err = edge.CreatedAtBlock()
		require.NoError(t, err)
		resp, err = ht.ComputeAncestorsWithTimers(ctx, edge.Id(), creation+1)
		require.NoError(t, err)
		timer, err = ht.ComputeHonestPathTimer(ctx, edge.Id(), resp.AncestorLocalTimers, creation+1)
		require.NoError(t, err)
		require.Equal(t, PathTimer(3), timer)

		// Continuing to query it at time T+i should increase the timer
		// as it is unrivaled.
		creation, err = edge.CreatedAtBlock()
		require.NoError(t, err)
		for i := uint64(2); i < 20; i++ {
			resp, err = ht.ComputeAncestorsWithTimers(ctx, edge.Id(), creation+i)
			require.NoError(t, err)
			timer, err = ht.ComputeHonestPathTimer(ctx, edge.Id(), resp.AncestorLocalTimers, creation+i)
			require.NoError(t, err)
			require.Equal(t, PathTimer(2)+PathTimer(i), timer)
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
			newEdge(&newCfg{t: t, edgeId: "blk-0.a-4.c", createdAt: 9}),
			newEdge(&newCfg{t: t, edgeId: "blk-4.c-8.c", createdAt: 9}),
		)
		// Child-relationship linking.
		edges["blk-0.a-16.c"].LowerChildID = "blk-0.a-8.c"
		edges["blk-0.a-16.c"].UpperChildID = "blk-8.c-16.c"
		edges["blk-0.a-8.c"].LowerChildID = "blk-0.a-4.c"
		edges["blk-0.a-8.c"].UpperChildID = "blk-4.c-8.c"

		// Add the new edges into the mapping.
		for k, v := range edges {
			ht.edges.Put(id(k), v)
		}

		// Three pairs of edges are rivaled in this test: 0-16, 0-8, 0-4
		mutual := edges["blk-0.a-16.c"].MutualId()
		mutuals := ht.mutualIds.Get(mutual)
		idd := id("blk-0.a-16.c")
		iddCreation, err = ht.edges.Get(idd).CreatedAtBlock()
		require.NoError(t, err)
		mutuals.Put(idd, creationTime(iddCreation))

		mutual = edges["blk-0.a-8.c"].MutualId()

		mutuals = ht.mutualIds.Get(mutual)
		idd = id("blk-0.a-8.c")
		iddCreation, err = ht.edges.Get(idd).CreatedAtBlock()
		require.NoError(t, err)
		mutuals.Put(idd, creationTime(iddCreation))

		mutual = edges["blk-0.a-4.c"].MutualId()

		ht.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
		mutuals = ht.mutualIds.Get(mutual)
		idd = id("blk-0.a-4.a")
		iddCreation, err = ht.edges.Get(idd).CreatedAtBlock()
		require.NoError(t, err)
		mutuals.Put(idd, creationTime(iddCreation))
		idd = id("blk-0.a-4.c")
		iddCreation, err = ht.edges.Get(idd).CreatedAtBlock()
		require.NoError(t, err)
		mutuals.Put(idd, creationTime(iddCreation))

		// The path timer of the old, unrivaled edge should no longer increase
		// as it is now rivaled as of the time of the last created edge above.
		lastCreated := ht.edges.Get(id("blk-0.a-4.c"))
		edge := ht.edges.Get(id("blk-0.a-4.a"))
		latestCreation, err := lastCreated.CreatedAtBlock()
		require.NoError(t, err)
		edgeCreation, err := edge.CreatedAtBlock()
		require.NoError(t, err)
		resp, err := ht.ComputeAncestorsWithTimers(ctx, edge.Id(), latestCreation)
		require.NoError(t, err)
		timer, err := ht.ComputeHonestPathTimer(ctx, edge.Id(), resp.AncestorLocalTimers, latestCreation)
		require.NoError(t, err)
		ancestorTimers := uint64(2)
		require.Equal(t, PathTimer(latestCreation-edgeCreation+ancestorTimers), timer)

		// Should no longer increase.
		for i := 0; i < 10; i++ {
			resp, err = ht.ComputeAncestorsWithTimers(ctx, edge.Id(), latestCreation+uint64(i))
			require.NoError(t, err)
			timer, err := ht.ComputeHonestPathTimer(ctx, edge.Id(), resp.AncestorLocalTimers, latestCreation+uint64(i))
			require.NoError(t, err)
			require.Equal(t, PathTimer(latestCreation-edgeCreation+ancestorTimers), timer)
		}
	})
}

// Tests out path timers across all challenge levels, checking the lowest, small step challenge
// edge inherits the local timers of all its honest ancestors through a cumulative update
// for confirmation purposes.
func TestComputePathTimer_AllChallengeLevels(t *testing.T) {
	unrivaledAssertionBlocks := uint64(10) // Should incorporate the assertion's unrivaled blocks into the total timer.
	ht := &HonestChallengeTree{
		edges:                         threadsafe.NewMap[protocol.EdgeId, protocol.SpecEdge](),
		mutualIds:                     threadsafe.NewMap[protocol.MutualId, *threadsafe.Map[protocol.EdgeId, creationTime]](),
		honestBigStepLevelZeroEdges:   threadsafe.NewSlice[protocol.ReadOnlyEdge](),
		honestSmallStepLevelZeroEdges: threadsafe.NewSlice[protocol.ReadOnlyEdge](),
		metadataReader: &mockMetadataReader{
			unrivaledAssertionBlocks: unrivaledAssertionBlocks,
		},
		totalChallengeLevels:   3,
		honestRootEdgesByLevel: threadsafe.NewMap[protocol.ChallengeLevel, *threadsafe.Slice[protocol.ReadOnlyEdge]](),
	}
	ht.honestRootEdgesByLevel.Put(2, threadsafe.NewSlice[protocol.ReadOnlyEdge]())
	ht.honestRootEdgesByLevel.Put(1, threadsafe.NewSlice[protocol.ReadOnlyEdge]())
	ht.honestRootEdgesByLevel.Put(0, threadsafe.NewSlice[protocol.ReadOnlyEdge]())

	// Edge ids that belong to block challenges are prefixed with "blk".
	// For big step, prefixed with "big", and small step, prefixed with "smol".
	setupBlockChallengeTreeSnapshot(t, ht)
	ht.honestBlockChalLevelZeroEdge = option.Some(protocol.ReadOnlyEdge(ht.edges.Get(id("blk-0.a-16.a"))))
	claimId := "blk-4.a-5.a"
	setupBigStepChallengeSnapshot(t, ht, claimId)
	bigStepRootEdges := ht.honestRootEdgesByLevel.Get(1 /* big step level */)
	bigStepRootEdges.Push(ht.edges.Get(id("big-0.a-16.a")))
	claimId = "big-4.a-5.a"
	setupSmallStepChallengeSnapshot(t, ht, claimId)
	smallStepRootEdges := ht.honestRootEdgesByLevel.Get(0 /* small step level */)
	smallStepRootEdges.Push(ht.edges.Get(id("smol-0.a-16.a")))

	ctx := context.Background()
	lastCreated := ht.edges.Get(id("smol-4.a-5.a"))
	lastCreatedTime, err := lastCreated.CreatedAtBlock()
	require.NoError(t, err)
	resp, err := ht.ComputeAncestorsWithTimers(ctx, lastCreated.Id(), lastCreatedTime+1)
	require.NoError(t, err)
	timer, err := ht.ComputeHonestPathTimer(ctx, lastCreated.Id(), resp.AncestorLocalTimers, lastCreatedTime+1)
	require.NoError(t, err)

	// Should be the sum of the unrivaled timers of honest edges along the path
	// all the way to the block challenge level. There are 15 edges in total, including the one
	// we are querying for. The assertion was unrivaled for 0 seconds. However, due to a merge move
	// made into edge with commit 4a, the edge blk-4.a-6.b from the malicious party was created
	// before blk-4.a-6.a, so 4.a-6.a was rivaled at time of creation. This means the total time
	// unrivaled is 15 - 1, which is 14.
	wantedAncestors := HonestAncestors{
		id("smol-4.a-6.a"),
		id("smol-4.a-8.a"),
		id("smol-0.a-8.a"),
		id("smol-0.a-16.a"),

		id("big-4.a-5.a"),
		id("big-4.a-6.a"),
		id("big-4.a-8.a"),
		id("big-0.a-8.a"),
		id("big-0.a-16.a"),

		id("blk-4.a-5.a"),
		id("blk-4.a-6.a"),
		id("blk-4.a-8.a"),
		id("blk-0.a-8.a"),
		id("blk-0.a-16.a"),
	}
	require.Equal(t, wantedAncestors, resp.AncestorEdgeIds)

	// This gives a total of 14 seconds unrivaled along the honest path plus the top-level assertion's
	// total amount of blocks unrivaled.
	require.Equal(t, PathTimer(14+unrivaledAssertionBlocks), timer)
}

func Test_findOriginEdge(t *testing.T) {
	edges := threadsafe.NewSlice[protocol.ReadOnlyEdge]()
	origin := protocol.OriginId(common.BytesToHash([]byte("foo")))
	_, ok := findOriginEdge(origin, edges)
	require.Equal(t, false, ok)
	edges.Push(newEdge(&newCfg{
		t:         t,
		originId:  "bar",
		edgeId:    "blk-0.a-4.a",
		claimId:   "",
		createdAt: 2,
	}))

	_, ok = findOriginEdge(origin, edges)
	require.Equal(t, false, ok)

	origin = protocol.OriginId(common.BytesToHash([]byte("bar")))
	got, ok := findOriginEdge(origin, edges)
	require.Equal(t, true, ok)
	require.Equal(t, got.Id(), protocol.EdgeId{Hash: common.BytesToHash([]byte("blk-0.a-4.a"))})

	edges.Push(newEdge(&newCfg{
		t:         t,
		originId:  "baz",
		edgeId:    "blk-0.b-4.b",
		claimId:   "",
		createdAt: 2,
	}))

	origin = protocol.OriginId(common.BytesToHash([]byte("baz")))
	got, ok = findOriginEdge(origin, edges)
	require.Equal(t, true, ok)
	require.Equal(t, got.Id(), protocol.EdgeId{Hash: common.BytesToHash([]byte("blk-0.b-4.b"))})
}

func buildEdges(allEdges ...*mock.Edge) map[mock.EdgeId]*mock.Edge {
	m := make(map[mock.EdgeId]*mock.Edge)
	for _, e := range allEdges {
		m[e.ID] = e
	}
	return m
}

// Sets up the following block challenge snapshot:
//
//	      /--5---6-----8-----------16 = Alice
//	0-----4
//	      \--5'--6'----8'----------16' = Bob
//
// and then inserts the respective edges into a challenge tree.
func setupBlockChallengeTreeSnapshot(t *testing.T, tree *HonestChallengeTree) {
	t.Helper()
	aliceEdges := buildEdges(
		// Alice.
		newEdge(&newCfg{t: t, edgeId: "blk-0.a-16.a", createdAt: 1}),
		newEdge(&newCfg{t: t, edgeId: "blk-0.a-8.a", createdAt: 3}),
		newEdge(&newCfg{t: t, edgeId: "blk-8.a-16.a", createdAt: 3}),
		newEdge(&newCfg{t: t, edgeId: "blk-0.a-4.a", createdAt: 5}),
		newEdge(&newCfg{t: t, edgeId: "blk-4.a-8.a", createdAt: 5}),
		newEdge(&newCfg{t: t, edgeId: "blk-4.a-6.a", createdAt: 7}),
		newEdge(&newCfg{t: t, edgeId: "blk-6.a-8.a", createdAt: 7}),
		newEdge(&newCfg{t: t, edgeId: "blk-4.a-5.a", createdAt: 9}),
		newEdge(&newCfg{t: t, edgeId: "blk-5.a-6.a", createdAt: 9}),
	)
	bobEdges := buildEdges(
		// Bob.
		newEdge(&newCfg{t: t, edgeId: "blk-0.a-16.b", createdAt: 2}),
		newEdge(&newCfg{t: t, edgeId: "blk-0.a-8.b", createdAt: 4}),
		newEdge(&newCfg{t: t, edgeId: "blk-8.b-16.b", createdAt: 4}),
		newEdge(&newCfg{t: t, edgeId: "blk-4.a-8.b", createdAt: 6}),
		newEdge(&newCfg{t: t, edgeId: "blk-4.a-6.b", createdAt: 6}),
		newEdge(&newCfg{t: t, edgeId: "blk-6.b-8.b", createdAt: 8}),
		newEdge(&newCfg{t: t, edgeId: "blk-4.a-5.b", createdAt: 10}),
		newEdge(&newCfg{t: t, edgeId: "blk-5.b-6.b", createdAt: 10}),
	)
	// Child-relationship linking.
	// Alice.
	aliceEdges["blk-0.a-16.a"].LowerChildID = "blk-0.a-8.a"
	aliceEdges["blk-0.a-16.a"].UpperChildID = "blk-8.a-16.a"
	aliceEdges["blk-0.a-8.a"].LowerChildID = "blk-0.a-4.a"
	aliceEdges["blk-0.a-8.a"].UpperChildID = "blk-4.a-8.a"
	aliceEdges["blk-4.a-8.a"].LowerChildID = "blk-4.a-6.a"
	aliceEdges["blk-4.a-8.a"].UpperChildID = "blk-6.a-8.a"
	aliceEdges["blk-4.a-6.a"].LowerChildID = "blk-4.a-5.a"
	aliceEdges["blk-4.a-6.a"].UpperChildID = "blk-5.a-6.a"
	// Bob.
	bobEdges["blk-0.a-16.b"].LowerChildID = "blk-0.a-8.b"
	bobEdges["blk-0.a-16.b"].UpperChildID = "blk-8.b-16.b"
	bobEdges["blk-0.a-8.b"].LowerChildID = "blk-0.a-4.a"
	bobEdges["blk-0.a-8.b"].UpperChildID = "blk-4.a-8.b"
	bobEdges["blk-4.a-8.b"].LowerChildID = "blk-4.a-6.b"
	bobEdges["blk-4.a-8.b"].UpperChildID = "blk-6.b-6.8"
	bobEdges["blk-4.a-6.b"].LowerChildID = "blk-4.a-5.b"
	bobEdges["blk-4.a-6.b"].UpperChildID = "blk-5.b-6.b"

	transformedEdges := make(map[protocol.EdgeId]protocol.SpecEdge)
	for _, v := range aliceEdges {
		transformedEdges[v.Id()] = v
	}
	allEdges := threadsafe.NewMapFromItems(transformedEdges)
	tree.edges = allEdges

	// Set up rivaled edges.
	mutual := aliceEdges["blk-0.a-16.a"].MutualId()
	tree.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
	mutuals := tree.mutualIds.Get(mutual)
	a := aliceEdges["blk-0.a-16.a"]
	b := bobEdges["blk-0.a-16.b"]
	aCreation, err := a.CreatedAtBlock()
	require.NoError(t, err)
	bCreation, err := b.CreatedAtBlock()
	require.NoError(t, err)
	mutuals.Put(a.Id(), creationTime(aCreation))
	mutuals.Put(b.Id(), creationTime(bCreation))

	mutual = aliceEdges["blk-0.a-8.a"].MutualId()
	tree.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
	mutuals = tree.mutualIds.Get(mutual)
	a = aliceEdges["blk-0.a-8.a"]
	b = bobEdges["blk-0.a-8.b"]
	aCreation, err = a.CreatedAtBlock()
	require.NoError(t, err)
	bCreation, err = b.CreatedAtBlock()
	require.NoError(t, err)
	mutuals.Put(a.Id(), creationTime(aCreation))
	mutuals.Put(b.Id(), creationTime(bCreation))

	mutual = aliceEdges["blk-4.a-8.a"].MutualId()
	tree.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
	mutuals = tree.mutualIds.Get(mutual)
	a = aliceEdges["blk-4.a-8.a"]
	b = bobEdges["blk-4.a-8.b"]
	aCreation, err = a.CreatedAtBlock()
	require.NoError(t, err)
	bCreation, err = b.CreatedAtBlock()
	require.NoError(t, err)
	mutuals.Put(a.Id(), creationTime(aCreation))
	mutuals.Put(b.Id(), creationTime(bCreation))

	mutual = aliceEdges["blk-4.a-6.a"].MutualId()
	tree.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
	mutuals = tree.mutualIds.Get(mutual)
	aCreation, err = a.CreatedAtBlock()
	require.NoError(t, err)
	bCreation, err = b.CreatedAtBlock()
	require.NoError(t, err)
	mutuals.Put(a.Id(), creationTime(aCreation))
	mutuals.Put(b.Id(), creationTime(bCreation))

	mutual = aliceEdges["blk-4.a-5.a"].MutualId()
	tree.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
	mutuals = tree.mutualIds.Get(mutual)
	a = aliceEdges["blk-4.a-5.a"]
	b = bobEdges["blk-4.a-5.b"]
	aCreation, err = a.CreatedAtBlock()
	require.NoError(t, err)
	bCreation, err = b.CreatedAtBlock()
	require.NoError(t, err)
	mutuals.Put(a.Id(), creationTime(aCreation))
	mutuals.Put(b.Id(), creationTime(bCreation))
}

func id(eId mock.EdgeId) protocol.EdgeId {
	return protocol.EdgeId{Hash: common.BytesToHash([]byte(eId))}
}

// Sets up the following big step challenge snapshot:
//
//	      /--5---6-----8-----------16 = Alice
//	0-----4
//	      \--5'--6'----8'----------16' = Bob
//
// and then inserts the respective edges into a challenge tree.
func setupBigStepChallengeSnapshot(t *testing.T, tree *HonestChallengeTree, claimId string) {
	t.Helper()
	originEdge := tree.edges.Get(id(mock.EdgeId(claimId))).(*mock.Edge)
	origin := mock.OriginId(originEdge.ComputeMutualId())
	aliceEdges := buildEdges(
		// Alice.
		newEdge(&newCfg{t: t, edgeId: "big-0.a-16.a", originId: origin, claimId: claimId, createdAt: 11}),
		newEdge(&newCfg{t: t, edgeId: "big-0.a-8.a", originId: origin, createdAt: 13}),
		newEdge(&newCfg{t: t, edgeId: "big-8.a-16.a", originId: origin, createdAt: 13}),
		newEdge(&newCfg{t: t, edgeId: "big-0.a-4.a", originId: origin, createdAt: 15}),
		newEdge(&newCfg{t: t, edgeId: "big-4.a-8.a", originId: origin, createdAt: 15}),
		newEdge(&newCfg{t: t, edgeId: "big-4.a-6.a", originId: origin, createdAt: 17}),
		newEdge(&newCfg{t: t, edgeId: "big-6.a-8.a", originId: origin, createdAt: 17}),
		newEdge(&newCfg{t: t, edgeId: "big-4.a-5.a", originId: origin, createdAt: 19}),
		newEdge(&newCfg{t: t, edgeId: "big-5.a-6.a", originId: origin, createdAt: 19}),
	)
	bobEdges := buildEdges(
		// Bob.
		newEdge(&newCfg{t: t, edgeId: "big-0.a-16.b", originId: origin, createdAt: 12}),
		newEdge(&newCfg{t: t, edgeId: "big-0.a-8.b", originId: origin, createdAt: 14}),
		newEdge(&newCfg{t: t, edgeId: "big-8.b-16.b", originId: origin, createdAt: 14}),
		newEdge(&newCfg{t: t, edgeId: "big-4.a-8.b", originId: origin, createdAt: 16}),
		newEdge(&newCfg{t: t, edgeId: "big-4.a-6.b", originId: origin, createdAt: 18}),
		newEdge(&newCfg{t: t, edgeId: "big-6.b-8.b", originId: origin, createdAt: 18}),
		newEdge(&newCfg{t: t, edgeId: "big-4.a-5.b", originId: origin, createdAt: 20}),
		newEdge(&newCfg{t: t, edgeId: "big-5.b-6.b", originId: origin, createdAt: 20}),
	)
	// Child-relationship linking.
	// Alice.
	aliceEdges["big-0.a-16.a"].LowerChildID = "big-0.a-8.a"
	aliceEdges["big-0.a-16.a"].UpperChildID = "big-8.a-16.a"
	aliceEdges["big-0.a-8.a"].LowerChildID = "big-0.a-4.a"
	aliceEdges["big-0.a-8.a"].UpperChildID = "big-4.a-8.a"
	aliceEdges["big-4.a-8.a"].LowerChildID = "big-4.a-6.a"
	aliceEdges["big-4.a-8.a"].UpperChildID = "big-6.a-8.a"
	aliceEdges["big-4.a-6.a"].LowerChildID = "big-4.a-5.a"
	aliceEdges["big-4.a-6.a"].UpperChildID = "big-5.a-6.a"
	// Bob.
	bobEdges["big-0.a-16.b"].LowerChildID = "big-0.a-8.b"
	bobEdges["big-0.a-16.b"].UpperChildID = "big-8.b-16.b"
	bobEdges["big-0.a-8.b"].LowerChildID = "big-0.a-4.a"
	bobEdges["big-0.a-8.b"].UpperChildID = "big-4.a-8.b"
	bobEdges["big-4.a-8.b"].LowerChildID = "big-4.a-6.b"
	bobEdges["big-4.a-8.b"].UpperChildID = "big-6.b-6.8"
	bobEdges["big-4.a-6.b"].LowerChildID = "big-4.a-5.b"
	bobEdges["big-4.a-6.b"].UpperChildID = "big-5.b-6.b"

	for _, v := range aliceEdges {
		tree.edges.Put(v.Id(), v)
	}

	// Set up rivaled edges.
	mutual := aliceEdges["big-0.a-16.a"].MutualId()
	tree.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
	mutuals := tree.mutualIds.Get(mutual)
	a := aliceEdges["big-0.a-16.a"]
	b := bobEdges["big-0.a-16.b"]
	aCreation, err := a.CreatedAtBlock()
	require.NoError(t, err)
	bCreation, err := b.CreatedAtBlock()
	require.NoError(t, err)
	mutuals.Put(a.Id(), creationTime(aCreation))
	mutuals.Put(b.Id(), creationTime(bCreation))

	mutual = aliceEdges["big-0.a-8.a"].MutualId()
	tree.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
	mutuals = tree.mutualIds.Get(mutual)
	a = aliceEdges["big-0.a-8.a"]
	b = bobEdges["big-0.a-8.b"]
	aCreation, err = a.CreatedAtBlock()
	require.NoError(t, err)
	bCreation, err = b.CreatedAtBlock()
	require.NoError(t, err)
	mutuals.Put(a.Id(), creationTime(aCreation))
	mutuals.Put(b.Id(), creationTime(bCreation))

	mutual = aliceEdges["big-4.a-8.a"].MutualId()
	tree.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
	mutuals = tree.mutualIds.Get(mutual)
	a = aliceEdges["big-4.a-8.a"]
	b = bobEdges["big-4.a-8.b"]
	aCreation, err = a.CreatedAtBlock()
	require.NoError(t, err)
	bCreation, err = b.CreatedAtBlock()
	require.NoError(t, err)
	mutuals.Put(a.Id(), creationTime(aCreation))
	mutuals.Put(b.Id(), creationTime(bCreation))

	mutual = aliceEdges["big-4.a-6.a"].MutualId()
	tree.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
	mutuals = tree.mutualIds.Get(mutual)
	a = aliceEdges["big-4.a-6.a"]
	b = bobEdges["big-4.a-6.b"]
	aCreation, err = a.CreatedAtBlock()
	require.NoError(t, err)
	bCreation, err = b.CreatedAtBlock()
	require.NoError(t, err)
	mutuals.Put(a.Id(), creationTime(aCreation))
	mutuals.Put(b.Id(), creationTime(bCreation))

	mutual = aliceEdges["big-4.a-5.a"].MutualId()
	tree.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
	mutuals = tree.mutualIds.Get(mutual)
	a = aliceEdges["big-4.a-5.a"]
	b = bobEdges["big-4.a-5.b"]
	aCreation, err = a.CreatedAtBlock()
	require.NoError(t, err)
	bCreation, err = b.CreatedAtBlock()
	require.NoError(t, err)
	mutuals.Put(a.Id(), creationTime(aCreation))
	mutuals.Put(b.Id(), creationTime(bCreation))
}

// Sets up the following small step challenge snapshot:
//
//	      /--5---6-----8-----------16 = Alice
//	0-----4
//	      \--5'--6'----8'----------16' = Bob
//
// and then inserts the respective edges into a challenge tree.
//
// and then inserts the respective edges into a challenge tree.
func setupSmallStepChallengeSnapshot(t *testing.T, tree *HonestChallengeTree, claimId string) {
	t.Helper()
	originEdge := tree.edges.Get(id(mock.EdgeId(claimId))).(*mock.Edge)
	origin := mock.OriginId(originEdge.ComputeMutualId())
	aliceEdges := buildEdges(
		// Alice.
		newEdge(&newCfg{t: t, edgeId: "smol-0.a-16.a", originId: origin, claimId: claimId, createdAt: 21}),
		newEdge(&newCfg{t: t, edgeId: "smol-0.a-8.a", originId: origin, createdAt: 23}),
		newEdge(&newCfg{t: t, edgeId: "smol-8.a-16.a", originId: origin, createdAt: 23}),
		newEdge(&newCfg{t: t, edgeId: "smol-0.a-4.a", originId: origin, createdAt: 25}),
		newEdge(&newCfg{t: t, edgeId: "smol-4.a-8.a", originId: origin, createdAt: 25}),
		newEdge(&newCfg{t: t, edgeId: "smol-4.a-6.a", originId: origin, createdAt: 27}),
		newEdge(&newCfg{t: t, edgeId: "smol-6.a-8.a", originId: origin, createdAt: 27}),
		newEdge(&newCfg{t: t, edgeId: "smol-4.a-5.a", originId: origin, createdAt: 29}),
		newEdge(&newCfg{t: t, edgeId: "smol-5.a-6.a", originId: origin, createdAt: 29}),
	)
	bobEdges := buildEdges(
		// Bob.
		newEdge(&newCfg{t: t, edgeId: "smol-0.a-16.b", originId: origin, createdAt: 22}),
		newEdge(&newCfg{t: t, edgeId: "smol-0.a-8.b", originId: origin, createdAt: 24}),
		newEdge(&newCfg{t: t, edgeId: "smol-8.b-16.b", originId: origin, createdAt: 24}),
		newEdge(&newCfg{t: t, edgeId: "smol-4.a-8.b", originId: origin, createdAt: 26}),
		newEdge(&newCfg{t: t, edgeId: "smol-4.a-6.b", originId: origin, createdAt: 28}),
		newEdge(&newCfg{t: t, edgeId: "smol-6.b-8.b", originId: origin, createdAt: 28}),
		newEdge(&newCfg{t: t, edgeId: "smol-4.a-5.b", originId: origin, createdAt: 30}),
		newEdge(&newCfg{t: t, edgeId: "smol-5.b-6.b", originId: origin, createdAt: 30}),
	)
	// Child-relationship linking.
	// Alice.
	aliceEdges["smol-0.a-16.a"].LowerChildID = "smol-0.a-8.a"
	aliceEdges["smol-0.a-16.a"].UpperChildID = "smol-8.a-16.a"
	aliceEdges["smol-0.a-8.a"].LowerChildID = "smol-0.a-4.a"
	aliceEdges["smol-0.a-8.a"].UpperChildID = "smol-4.a-8.a"
	aliceEdges["smol-4.a-8.a"].LowerChildID = "smol-4.a-6.a"
	aliceEdges["smol-4.a-8.a"].UpperChildID = "smol-6.a-8.a"
	aliceEdges["smol-4.a-6.a"].LowerChildID = "smol-4.a-5.a"
	aliceEdges["smol-4.a-6.a"].UpperChildID = "smol-5.a-6.a"
	// Bob.
	bobEdges["smol-0.a-16.b"].LowerChildID = "smol-0.a-8.b"
	bobEdges["smol-0.a-16.b"].UpperChildID = "smol-8.b-16.b"
	bobEdges["smol-0.a-8.b"].LowerChildID = "smol-0.a-4.a"
	bobEdges["smol-0.a-8.b"].UpperChildID = "smol-4.a-8.b"
	bobEdges["smol-4.a-8.b"].LowerChildID = "smol-4.a-6.b"
	bobEdges["smol-4.a-8.b"].UpperChildID = "smol-6.b-6.8"
	bobEdges["smol-4.a-6.b"].LowerChildID = "smol-4.a-5.b"
	bobEdges["smol-4.a-6.b"].UpperChildID = "smol-5.b-6.b"

	for _, v := range aliceEdges {
		tree.edges.Put(v.Id(), v)
	}

	// Set up rivaled edges.
	mutual := aliceEdges["smol-0.a-16.a"].MutualId()
	tree.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
	mutuals := tree.mutualIds.Get(mutual)
	a := aliceEdges["smol-0.a-16.a"]
	b := bobEdges["smol-0.a-16.b"]
	aCreation, err := a.CreatedAtBlock()
	require.NoError(t, err)
	bCreation, err := b.CreatedAtBlock()
	require.NoError(t, err)
	mutuals.Put(a.Id(), creationTime(aCreation))
	mutuals.Put(b.Id(), creationTime(bCreation))

	mutual = aliceEdges["smol-0.a-8.a"].MutualId()
	tree.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
	mutuals = tree.mutualIds.Get(mutual)
	a = aliceEdges["smol-0.a-8.a"]
	b = bobEdges["smol-0.a-8.b"]
	aCreation, err = a.CreatedAtBlock()
	require.NoError(t, err)
	bCreation, err = b.CreatedAtBlock()
	require.NoError(t, err)
	mutuals.Put(a.Id(), creationTime(aCreation))
	mutuals.Put(b.Id(), creationTime(bCreation))

	mutual = aliceEdges["smol-4.a-8.a"].MutualId()
	tree.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
	mutuals = tree.mutualIds.Get(mutual)
	a = aliceEdges["smol-4.a-8.a"]
	b = bobEdges["smol-4.a-8.b"]
	aCreation, err = a.CreatedAtBlock()
	require.NoError(t, err)
	bCreation, err = b.CreatedAtBlock()
	require.NoError(t, err)
	mutuals.Put(a.Id(), creationTime(aCreation))
	mutuals.Put(b.Id(), creationTime(bCreation))

	mutual = aliceEdges["smol-4.a-6.a"].MutualId()
	tree.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
	mutuals = tree.mutualIds.Get(mutual)
	a = aliceEdges["smol-4.a-6.a"]
	b = bobEdges["smol-4.a-6.b"]
	aCreation, err = a.CreatedAtBlock()
	require.NoError(t, err)
	bCreation, err = b.CreatedAtBlock()
	require.NoError(t, err)
	mutuals.Put(a.Id(), creationTime(aCreation))
	mutuals.Put(b.Id(), creationTime(bCreation))

	mutual = aliceEdges["smol-4.a-5.a"].MutualId()
	tree.mutualIds.Put(mutual, threadsafe.NewMap[protocol.EdgeId, creationTime]())
	mutuals = tree.mutualIds.Get(mutual)
	a = aliceEdges["smol-4.a-5.a"]
	b = bobEdges["smol-4.a-5.b"]
	aCreation, err = a.CreatedAtBlock()
	require.NoError(t, err)
	bCreation, err = b.CreatedAtBlock()
	require.NoError(t, err)
	mutuals.Put(a.Id(), creationTime(aCreation))
	mutuals.Put(b.Id(), creationTime(bCreation))
}
