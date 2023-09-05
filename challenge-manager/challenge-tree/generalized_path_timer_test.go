// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE

package challengetree

import (
	"context"
	"testing"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/OffchainLabs/bold/containers/option"
	"github.com/OffchainLabs/bold/containers/threadsafe"
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
func Test_computeAncestorsWithTimers(t *testing.T) {
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
		_, err := tree.computeAncestorsWithTimers(ctx, id("foo"), blockNum)
		require.ErrorContains(t, err, "not found in honest challenge tree")
	})
	t.Run("dishonest edge lookup fails", func(t *testing.T) {
		_, err := tree.computeAncestorsWithTimers(ctx, id("blk-0.a-16.b"), blockNum)
		require.ErrorContains(t, err, "not found in honest challenge tree")
	})
	t.Run("block challenge: level zero edge has no ancestors", func(t *testing.T) {
		resp, err := tree.computeAncestorsWithTimers(ctx, id("blk-0.a-16.a"), blockNum)
		require.NoError(t, err)
		require.Equal(t, 0, len(resp.ancestorEdgeIds))
	})
	t.Run("block challenge: single ancestor", func(t *testing.T) {
		resp, err := tree.computeAncestorsWithTimers(ctx, id("blk-0.a-8.a"), blockNum)
		require.NoError(t, err)
		require.Equal(t, HonestAncestors{id("blk-0.a-16.a")}, resp.ancestorEdgeIds)
		resp, err = tree.computeAncestorsWithTimers(ctx, id("blk-8.a-16.a"), blockNum)
		require.NoError(t, err)
		require.Equal(t, HonestAncestors{id("blk-0.a-16.a")}, resp.ancestorEdgeIds)
	})
	t.Run("block challenge: many ancestors", func(t *testing.T) {
		_, ancestors, err := tree.HonestPathTimer(ctx, id("blk-4.a-5.a"), blockNum)
		require.NoError(t, err)
		wanted := HonestAncestors{
			id("blk-4.a-6.a"),
			id("blk-4.a-8.a"),
			id("blk-0.a-8.a"),
			id("blk-0.a-16.a"),
		}
		require.Equal(t, wanted, ancestors)
	})
	t.Run("big step challenge: level zero edge has ancestors from block challenge", func(t *testing.T) {
		_, ancestors, err := tree.HonestPathTimer(ctx, id("big-0.a-16.a"), blockNum)
		require.NoError(t, err)
		wanted := HonestAncestors{
			id("blk-4.a-5.a"),
			id("blk-4.a-6.a"),
			id("blk-4.a-8.a"),
			id("blk-0.a-8.a"),
			id("blk-0.a-16.a"),
		}
		require.Equal(t, wanted, ancestors)
	})
	// t.Run("big step challenge: many ancestors plus block challenge ancestors", func(t *testing.T) {
	// 	_, ancestors, err := tree.HonestPathTimer(ctx, id("big-5.a-6.a"), blockNum)
	// 	require.NoError(t, err)
	// 	wanted := HonestAncestors{
	// 		// Big step chal.
	// 		id("big-4.a-6.a"),
	// 		id("big-4.a-8.a"),
	// 		id("big-0.a-8.a"),
	// 		id("big-0.a-16.a"),
	// 		// Block chal.
	// 		id("blk-4.a-5.a"),
	// 		id("blk-4.a-6.a"),
	// 		id("blk-4.a-8.a"),
	// 		id("blk-0.a-8.a"),
	// 		id("blk-0.a-16.a"),
	// 	}
	// 	require.Equal(t, wanted, ancestors)
	// })
	// t.Run("small step challenge: level zero edge has ancestors from big and block challenge", func(t *testing.T) {
	// 	_, ancestors, err := tree.HonestPathTimer(ctx, id("smol-0.a-16.a"), blockNum)
	// 	require.NoError(t, err)
	// 	wanted := HonestAncestors{
	// 		// Big step chal.
	// 		id("big-4.a-5.a"),
	// 		id("big-4.a-6.a"),
	// 		id("big-4.a-8.a"),
	// 		id("big-0.a-8.a"),
	// 		id("big-0.a-16.a"),
	// 		// Block chal.
	// 		id("blk-4.a-5.a"),
	// 		id("blk-4.a-6.a"),
	// 		id("blk-4.a-8.a"),
	// 		id("blk-0.a-8.a"),
	// 		id("blk-0.a-16.a"),
	// 	}
	// 	require.Equal(t, wanted, ancestors)
	// })
	// t.Run("small step challenge: lowest level edge has full ancestry", func(t *testing.T) {
	// 	_, ancestors, err := tree.HonestPathTimer(ctx, id("smol-5.a-6.a"), blockNum)
	// 	require.NoError(t, err)
	// 	wanted := HonestAncestors{
	// 		// Small step chal.
	// 		id("smol-4.a-6.a"),
	// 		id("smol-4.a-8.a"),
	// 		id("smol-0.a-8.a"),
	// 		id("smol-0.a-16.a"),
	// 		// Big step chal.
	// 		id("big-4.a-5.a"),
	// 		id("big-4.a-6.a"),
	// 		id("big-4.a-8.a"),
	// 		id("big-0.a-8.a"),
	// 		id("big-0.a-16.a"),
	// 		// Block chal.
	// 		id("blk-4.a-5.a"),
	// 		id("blk-4.a-6.a"),
	// 		id("blk-4.a-8.a"),
	// 		id("blk-0.a-8.a"),
	// 		id("blk-0.a-16.a"),
	// 	}
	// 	require.Equal(t, wanted, ancestors)
	// })
}
