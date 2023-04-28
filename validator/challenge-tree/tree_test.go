package challengetree

import (
	"testing"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/OffchainLabs/challenge-protocol-v2/util/threadsafe"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

// TODO: Add benchmark of worst case scenario updates, add some latency to the chain calls.
// TODO: Test adding edges to the challenge tree.
// TODO: Rename challenge tree to HonestTreeForChallenge

type mockChain struct {
	rivaledEdges         *threadsafe.Set[protocol.EdgeId]
	unrivaledTimerByEdge map[protocol.EdgeId]uint64
}

func (m *mockChain) timeUnrivaled(edgeId protocol.EdgeId) uint64 {
	return m.unrivaledTimerByEdge[edgeId]
}

func (m *mockChain) advanceBlocks(numBlocks uint64) {
	for k, v := range m.unrivaledTimerByEdge {
		// Do not update unrivaled timer for rivaled edges.
		if m.rivaledEdges.Has(k) {
			continue
		}
		m.unrivaledTimerByEdge[k] = v + numBlocks
	}
}

func TestAddEdge(t *testing.T) {
}

func TestCumulativeUnrivaledTimeUpdates_UniformInitialUnrivaledTimes(t *testing.T) {
	t.Skip()
	tree := &challengeTree{
		edges:                           threadsafe.NewMap[protocol.EdgeId, *edge](),
		rivaledEdges:                    threadsafe.NewSet[protocol.EdgeId](),
		honestUnrivaledCumulativeTimers: threadsafe.NewMap[protocol.EdgeId, uint64](),
	}
	setupBlockChallengeTreeSnapshot(t, tree)
	claimId := protocol.ClaimId(id("blk-5-6"))
	setupBigStepChallengeSnapshot(t, tree, claimId)
	claimId = protocol.ClaimId(id("big-5-6"))
	setupSmallStepChallengeSnapshot(t, tree, claimId)

	mChain := &mockChain{
		rivaledEdges: tree.rivaledEdges,
		unrivaledTimerByEdge: map[protocol.EdgeId]uint64{
			// Initial, unrivaled timers for block challenge edges.
			id("blk-0-16"): 1,
			id("blk-8-16"): 1,
			id("blk-0-8"):  1,
			id("blk-0-4"):  1,
			id("blk-4-8"):  1,
			id("blk-4-6"):  1,
			id("blk-6-8"):  1,
			id("blk-4-5"):  1,
			id("blk-5-6"):  1,
			// Initial, unrivaled timers for big step challenge edges.
			id("big-0-16"): 1,
			id("big-8-16"): 1,
			id("big-0-8"):  1,
			id("big-0-4"):  1,
			id("big-4-8"):  1,
			id("big-4-6"):  1,
			id("big-6-8"):  1,
			id("big-4-5"):  1,
			id("big-5-6"):  1,
			// Initial, unrivaled timers for small step challenge edges.
			id("smol-0-16"): 1,
			id("smol-8-16"): 1,
			id("smol-0-8"):  1,
			id("smol-0-4"):  1,
			id("smol-4-8"):  1,
			id("smol-4-6"):  1,
			id("smol-6-8"):  1,
			id("smol-4-5"):  1,
			id("smol-5-6"):  1,
		},
	}
	// Advance by 4 blocks and expect that only unrivaled edges, blk-0-4, big-0-4,
	// and smol-0-4, were the only ones to change their timers.
	mChain.advanceBlocks(4)

	for k, v := range mChain.unrivaledTimerByEdge {
		if k == id("blk-0-4") || k == id("big-0-4") || k == id("smol-0-4") {
			require.Equal(t, uint64(5), v)
		} else {
			require.Equal(t, uint64(1), v)
		}
	}

	tree.chain = mChain

	// We now recursively update the cumulative timers for each edge being tracked
	// in our challenge tree, and are able to expect certain results based on the setup above.
	tree.updateCumulativeTimers()

	t.Run("level zero block challenge edge", func(t *testing.T) {
		edgeId := id("blk-0-16")
		got, err := tree.CumulativeTimeUnrivaled(edgeId)
		require.NoError(t, err)
		require.Equal(t, uint64(1), got)
	})

	t.Run("unrivaled block challenge edge", func(t *testing.T) {
		unrivaledBlockEdgeId := id("blk-0-4")
		// Expect the edge's cumulative time unrivaled is the time unrivaled of all its ancestors
		// plus its own time unrivaled.
		ancestors := tree.ancestorsForHonestEdge(unrivaledBlockEdgeId)
		var wantedAncestorsTotal uint64
		for _, an := range ancestors {
			timeUnrivaled := tree.chain.timeUnrivaled(an)
			wantedAncestorsTotal += timeUnrivaled
		}
		edgeUnrivaledTimer := tree.chain.timeUnrivaled(unrivaledBlockEdgeId)
		got, err := tree.CumulativeTimeUnrivaled(unrivaledBlockEdgeId)
		require.NoError(t, err)
		require.Equal(t, edgeUnrivaledTimer+wantedAncestorsTotal, got)
	})

	t.Run("rivaled block challenge edge", func(t *testing.T) {
		edgeId := id("blk-5-6")
		ancestors := tree.ancestorsForHonestEdge(edgeId)
		var wantedAncestorsTotal uint64
		for _, an := range ancestors {
			timeUnrivaled := tree.chain.timeUnrivaled(an)
			wantedAncestorsTotal += timeUnrivaled
		}
		edgeUnrivaledTimer := tree.chain.timeUnrivaled(edgeId)
		got, err := tree.CumulativeTimeUnrivaled(edgeId)
		require.NoError(t, err)
		require.Equal(t, edgeUnrivaledTimer+wantedAncestorsTotal, got)
	})

	t.Run("level zero big step challenge edge", func(t *testing.T) {
		edgeId := id("big-0-16")
		ancestors := tree.ancestorsForHonestEdge(edgeId)
		var wantedAncestorsTotal uint64
		for _, an := range ancestors {
			timeUnrivaled := tree.chain.timeUnrivaled(an)
			wantedAncestorsTotal += timeUnrivaled
		}
		edgeUnrivaledTimer := tree.chain.timeUnrivaled(edgeId)
		got, err := tree.CumulativeTimeUnrivaled(edgeId)
		require.NoError(t, err)
		require.Equal(t, edgeUnrivaledTimer+wantedAncestorsTotal, got)
	})

	t.Run("unrivaled big step challenge edge", func(t *testing.T) {
		edgeId := id("big-0-4")
		ancestors := tree.ancestorsForHonestEdge(edgeId)
		var wantedAncestorsTotal uint64
		for _, an := range ancestors {
			timeUnrivaled := tree.chain.timeUnrivaled(an)
			wantedAncestorsTotal += timeUnrivaled
		}
		edgeUnrivaledTimer := tree.chain.timeUnrivaled(edgeId)
		got, err := tree.CumulativeTimeUnrivaled(edgeId)
		require.NoError(t, err)
		require.Equal(t, edgeUnrivaledTimer+wantedAncestorsTotal, got)
	})

	t.Run("rivaled big step challenge edge", func(t *testing.T) {
		edgeId := id("big-5-6")
		ancestors := tree.ancestorsForHonestEdge(edgeId)
		var wantedAncestorsTotal uint64
		for _, an := range ancestors {
			timeUnrivaled := tree.chain.timeUnrivaled(an)
			wantedAncestorsTotal += timeUnrivaled
		}
		edgeUnrivaledTimer := tree.chain.timeUnrivaled(edgeId)
		got, err := tree.CumulativeTimeUnrivaled(edgeId)
		require.NoError(t, err)
		require.Equal(t, edgeUnrivaledTimer+wantedAncestorsTotal, got)
	})

	t.Run("level zero small step challenge edge", func(t *testing.T) {
		edgeId := id("smol-0-16")
		ancestors := tree.ancestorsForHonestEdge(edgeId)
		var wantedAncestorsTotal uint64
		for _, an := range ancestors {
			timeUnrivaled := tree.chain.timeUnrivaled(an)
			wantedAncestorsTotal += timeUnrivaled
		}
		edgeUnrivaledTimer := tree.chain.timeUnrivaled(edgeId)
		got, err := tree.CumulativeTimeUnrivaled(edgeId)
		require.NoError(t, err)
		require.Equal(t, edgeUnrivaledTimer+wantedAncestorsTotal, got)
	})

	t.Run("unrivaled small step challenge edge", func(t *testing.T) {
		edgeId := id("smol-0-4")
		ancestors := tree.ancestorsForHonestEdge(edgeId)
		var wantedAncestorsTotal uint64
		for _, an := range ancestors {
			timeUnrivaled := tree.chain.timeUnrivaled(an)
			wantedAncestorsTotal += timeUnrivaled
		}
		edgeUnrivaledTimer := tree.chain.timeUnrivaled(edgeId)
		got, err := tree.CumulativeTimeUnrivaled(edgeId)
		require.NoError(t, err)
		require.Equal(t, edgeUnrivaledTimer+wantedAncestorsTotal, got)
	})

	t.Run("rivaled small step challenge edge", func(t *testing.T) {
		edgeId := id("smol-5-6")
		ancestors := tree.ancestorsForHonestEdge(edgeId)
		var wantedAncestorsTotal uint64
		for _, an := range ancestors {
			timeUnrivaled := tree.chain.timeUnrivaled(an)
			wantedAncestorsTotal += timeUnrivaled
		}
		edgeUnrivaledTimer := tree.chain.timeUnrivaled(edgeId)
		got, err := tree.CumulativeTimeUnrivaled(edgeId)
		require.NoError(t, err)
		require.Equal(t, edgeUnrivaledTimer+wantedAncestorsTotal, got)
	})

	t.Run("rivaled edge cumulative does not change but unrivaled does after advancing blocks", func(t *testing.T) {
		// We advance blocks, and check that the only cumulative time that changed
		// was that of the unrivaled edges blk-0-4, big-0-4, smol-0-4 and the deepest.
		unrivaledEdgeTimer, err := tree.CumulativeTimeUnrivaled(id("blk-0-4"))
		require.NoError(t, err)
		rivaledEdgeTimer, err := tree.CumulativeTimeUnrivaled(id("blk-4-5"))
		require.NoError(t, err)

		mChain.advanceBlocks(4)
		tree.updateCumulativeTimers()

		unrivaledEdgeTimerAfterAdvancing, err := tree.CumulativeTimeUnrivaled(id("blk-0-4"))
		require.NoError(t, err)
		rivaledEdgeTimerAfterAdvancing, err := tree.CumulativeTimeUnrivaled(id("blk-4-5"))
		require.NoError(t, err)

		require.Equal(t, rivaledEdgeTimer, rivaledEdgeTimerAfterAdvancing)
		require.Equal(t, unrivaledEdgeTimer+4, unrivaledEdgeTimerAfterAdvancing)
	})
}

func TestAncestors_BlockChallengeOnly(t *testing.T) {
	tree := &challengeTree{
		edges:        threadsafe.NewMap[protocol.EdgeId, *edge](),
		rivaledEdges: threadsafe.NewSet[protocol.EdgeId](),
		mutualIds:    threadsafe.NewMap[protocol.MutualId, *threadsafe.Set[protocol.EdgeId]](),
	}
	setupBlockChallengeTreeSnapshot(t, tree)

	// Edge ids that belong to block challenges are prefixed with "blk".
	// For big step, prefixed with "big", and small step, prefixed with "smol".
	// ancestors := tree.ancestorsForHonestEdge(id("blk-6-8"))
	// require.Equal(t, ancestors, []protocol.EdgeId{id("blk-4-8"), id("blk-0-8"), id("blk-0-16")})

	// ancestors = tree.ancestorsForHonestEdge(id("blk-4-6"))
	// require.Equal(t, ancestors, []protocol.EdgeId{id("blk-4-8"), id("blk-0-8"), id("blk-0-16")})

	ancestors := tree.ancestorsForHonestEdge(id("blk-0-4"))
	for _, an := range ancestors {
		t.Logf("%s", an)
	}
	require.Equal(t, ancestors, []protocol.EdgeId{id("blk-0-8"), id("blk-0-16")})

	// ancestors = tree.ancestorsForHonestEdge(id("blk-4-8"))
	// require.Equal(t, ancestors, []protocol.EdgeId{id("blk-0-8"), id("blk-0-16")})

	// ancestors = tree.ancestorsForHonestEdge(id("blk-5-6"))
	// require.Equal(t, ancestors, []protocol.EdgeId{id("blk-4-6"), id("blk-4-8"), id("blk-0-8"), id("blk-0-16")})

	ancestors = tree.ancestorsForHonestEdge(id("blk-0-16"))
	require.Equal(t, 0, len(ancestors))
}

// Tests the following tree, all the way down to the small
// step subchallenge level.
//
//	 Block challenge tree:
//		0-----------------------8, 8----------------------16
//		0-----------------------8
//		0-------4, 4------------8
//		4----6, 6----8
//		4--5, 5--6
//
//	 Big step challenge tree:
//		0-----------------------8, 8----------------------16 (claim_id = id(5,6) in the level above)
//		0-----------------------8
//		0-------4, 4------------8
//		4----6, 6----8
//		4--5, 5--6
//
//	 Small step challenge tree:
//		0-----------------------8, 8----------------------16 (claim_id = id(5,6) in the level above)
//		0-----------------------8
//		0-------4, 4------------8
//		4----6, 6----8
//		4--5, 5--6
func TestAncestors_AllChallengeLevels(t *testing.T) {
	tree := &challengeTree{
		edges:        threadsafe.NewMap[protocol.EdgeId, *edge](),
		rivaledEdges: threadsafe.NewSet[protocol.EdgeId](),
	}
	setupBlockChallengeTreeSnapshot(t, tree)
	claimId := protocol.ClaimId(id("blk-5-6"))
	setupBigStepChallengeSnapshot(t, tree, claimId)
	claimId = protocol.ClaimId(id("big-5-6"))
	setupSmallStepChallengeSnapshot(t, tree, claimId)

	// We start by querying for ancestors for a big step edge id.
	ancestors := tree.ancestorsForHonestEdge(id("big-4-5"))

	// Edge ids that belong to block challenges are prefixed with "blk".
	// For big step, prefixed with "big", and small step, prefixed with "smol".
	wanted := []protocol.EdgeId{
		id("big-4-6"),
		id("big-4-8"),
		id("big-0-8"),
		id("big-0-16"),
		id("blk-5-6"), // TODO: Should the claim id be part of the ancestors as well?
		id("blk-4-6"),
		id("blk-4-8"),
		id("blk-0-8"),
		id("blk-0-16"),
	}
	require.Equal(t, wanted, ancestors)

	// We start query the ancestors of the lowest level, length one, small step edge.
	ancestors = tree.ancestorsForHonestEdge(id("smol-5-6"))
	wanted = []protocol.EdgeId{
		id("smol-4-6"),
		id("smol-4-8"),
		id("smol-0-8"),
		id("smol-0-16"),
		id("big-5-6"),
		id("big-4-6"),
		id("big-4-8"),
		id("big-0-8"),
		id("big-0-16"),
		id("blk-5-6"), // TODO: Should the claim id be part of the ancestors as well?
		id("blk-4-6"),
		id("blk-4-8"),
		id("blk-0-8"),
		id("blk-0-16"),
	}
	require.Equal(t, wanted, ancestors)

	// Query the level zero edge at each challenge type.
	ancestors = tree.ancestorsForHonestEdge(id("blk-0-16"))
	require.Equal(t, 0, len(ancestors))

	ancestors = tree.ancestorsForHonestEdge(id("big-0-16"))
	require.Equal(t, ancestors, []protocol.EdgeId{
		id("blk-5-6"),
		id("blk-4-6"),
		id("blk-4-8"),
		id("blk-0-8"),
		id("blk-0-16"),
	})

	ancestors = tree.ancestorsForHonestEdge(id("smol-0-16"))
	require.Equal(t, ancestors, []protocol.EdgeId{
		id("big-5-6"),
		id("big-4-6"),
		id("big-4-8"),
		id("big-0-8"),
		id("big-0-16"),
		id("blk-5-6"),
		id("blk-4-6"),
		id("blk-4-8"),
		id("blk-0-8"),
		id("blk-0-16"),
	})
}

// Sets up the following block challenge snapshot:
//
//	0-----------------------8, 8----------------------16
//	0-----------------------8
//	0-------4, 4------------8
//	4----6, 6----8
//	4--5, 5--6
//
// and then inserts the respective edges into a challenge tree.
func setupBlockChallengeTreeSnapshot(t *testing.T, tree *challengeTree) {
	t.Helper()
	topLevel := &edge{
		id:          id("blk-0-16"),
		edgeType:    protocol.BlockChallengeEdge,
		startHeight: 0,
		endHeight:   16,
	}
	tree.honestBlockChalLevelZeroEdge = util.Some(topLevel)
	child08 := &edge{
		id:          id("blk-0-8"),
		edgeType:    protocol.BlockChallengeEdge,
		startHeight: 0,
		endHeight:   8,
	}
	child816 := &edge{
		id:          id("blk-8-16"),
		edgeType:    protocol.BlockChallengeEdge,
		startHeight: 8,
		endHeight:   16,
	}
	topLevel.lowerChildId = common.Hash(child08.id)
	topLevel.upperChildId = common.Hash(child816.id)

	child04 := &edge{
		id:          id("blk-0-4"),
		edgeType:    protocol.BlockChallengeEdge,
		startHeight: 0,
		endHeight:   4,
	}
	child48 := &edge{
		id:          id("blk-4-8"),
		edgeType:    protocol.BlockChallengeEdge,
		startHeight: 4,
		endHeight:   8,
	}
	child08.lowerChildId = common.Hash(child04.id)
	child08.upperChildId = common.Hash(child48.id)

	child46 := &edge{
		id:          id("blk-4-6"),
		edgeType:    protocol.BlockChallengeEdge,
		startHeight: 4,
		endHeight:   6,
	}
	child68 := &edge{
		id:          id("blk-6-8"),
		edgeType:    protocol.BlockChallengeEdge,
		startHeight: 6,
		endHeight:   8,
	}
	child48.lowerChildId = common.Hash(child46.id)
	child48.upperChildId = common.Hash(child68.id)

	child45 := &edge{
		id:          id("blk-4-5"),
		edgeType:    protocol.BlockChallengeEdge,
		startHeight: 4,
		endHeight:   5,
	}
	child56 := &edge{
		id:          id("blk-5-6"),
		edgeType:    protocol.BlockChallengeEdge,
		startHeight: 5,
		endHeight:   6,
	}
	child46.lowerChildId = common.Hash(child45.id)
	child46.upperChildId = common.Hash(child56.id)

	tree.edges.Insert(topLevel.id, topLevel)
	tree.edges.Insert(child08.id, child08)
	tree.edges.Insert(child816.id, child816)
	tree.edges.Insert(child04.id, child04)
	tree.edges.Insert(child48.id, child48)
	tree.edges.Insert(child46.id, child46)
	tree.edges.Insert(child68.id, child68)
	tree.edges.Insert(child45.id, child45)
	tree.edges.Insert(child56.id, child56)

	// Rival all edges except for 0-4
	tree.rivaledEdges.Insert(topLevel.id)
	tree.rivaledEdges.Insert(child08.id)
	tree.rivaledEdges.Insert(child816.id)
	tree.rivaledEdges.Insert(child48.id)
	tree.rivaledEdges.Insert(child46.id)
	tree.rivaledEdges.Insert(child68.id)
	tree.rivaledEdges.Insert(child45.id)
	tree.rivaledEdges.Insert(child56.id)
}

// Sets up the following big step challenge snapshot:
//
//	0-----------------------8, 8----------------------16
//	0-----------------------8
//	0-------4, 4------------8
//	4----6, 6----8
//	4--5, 5--6
//
// and then inserts the respective edges into a challenge tree.
func setupBigStepChallengeSnapshot(t *testing.T, tree *challengeTree, claimId protocol.ClaimId) {
	t.Helper()
	topLevel := &edge{
		id:          id("big-0-16"),
		edgeType:    protocol.BigStepChallengeEdge,
		claimId:     common.Hash(claimId),
		startHeight: 0,
		endHeight:   16,
	}
	tree.honestBigStepChalLevelZeroEdge = util.Some(topLevel)
	child08 := &edge{
		id:          id("big-0-8"),
		edgeType:    protocol.BigStepChallengeEdge,
		startHeight: 0,
		endHeight:   8,
	}
	child816 := &edge{
		id:          id("big-8-16"),
		edgeType:    protocol.BigStepChallengeEdge,
		startHeight: 8,
		endHeight:   16,
	}
	topLevel.lowerChildId = common.Hash(child08.id)
	topLevel.upperChildId = common.Hash(child816.id)

	child04 := &edge{
		id:          id("big-0-4"),
		edgeType:    protocol.BigStepChallengeEdge,
		startHeight: 0,
		endHeight:   4,
	}
	child48 := &edge{
		id:          id("big-4-8"),
		edgeType:    protocol.BigStepChallengeEdge,
		startHeight: 4,
		endHeight:   8,
	}
	child08.lowerChildId = common.Hash(child04.id)
	child08.upperChildId = common.Hash(child48.id)

	child46 := &edge{
		id:          id("big-4-6"),
		edgeType:    protocol.BigStepChallengeEdge,
		startHeight: 4,
		endHeight:   6,
	}
	child68 := &edge{
		id:          id("big-6-8"),
		edgeType:    protocol.BigStepChallengeEdge,
		startHeight: 6,
		endHeight:   8,
	}
	child48.lowerChildId = common.Hash(child46.id)
	child48.upperChildId = common.Hash(child68.id)

	child45 := &edge{
		id:          id("big-4-5"),
		edgeType:    protocol.BigStepChallengeEdge,
		startHeight: 4,
		endHeight:   5,
	}
	child56 := &edge{
		id:          id("big-5-6"),
		edgeType:    protocol.BigStepChallengeEdge,
		startHeight: 5,
		endHeight:   6,
	}
	child46.lowerChildId = common.Hash(child45.id)
	child46.upperChildId = common.Hash(child56.id)

	tree.edges.Insert(topLevel.id, topLevel)
	tree.edges.Insert(child08.id, child08)
	tree.edges.Insert(child816.id, child816)
	tree.edges.Insert(child04.id, child04)
	tree.edges.Insert(child48.id, child48)
	tree.edges.Insert(child46.id, child46)
	tree.edges.Insert(child68.id, child68)
	tree.edges.Insert(child45.id, child45)
	tree.edges.Insert(child56.id, child56)

	// Rival all edges except for 0-4
	tree.rivaledEdges.Insert(topLevel.id)
	tree.rivaledEdges.Insert(child08.id)
	tree.rivaledEdges.Insert(child816.id)
	tree.rivaledEdges.Insert(child48.id)
	tree.rivaledEdges.Insert(child46.id)
	tree.rivaledEdges.Insert(child68.id)
	tree.rivaledEdges.Insert(child45.id)
	tree.rivaledEdges.Insert(child56.id)
}

// Sets up the following small step challenge snapshot:
//
//	0-----------------------8, 8----------------------16
//	0-----------------------8
//	0-------4, 4------------8
//	4----6, 6----8
//	4--5, 5--6
//
// and then inserts the respective edges into a challenge tree.
func setupSmallStepChallengeSnapshot(t *testing.T, tree *challengeTree, claimId protocol.ClaimId) {
	t.Helper()
	topLevel := &edge{
		id:          id("smol-0-16"),
		edgeType:    protocol.SmallStepChallengeEdge,
		claimId:     common.Hash(claimId),
		startHeight: 0,
		endHeight:   16,
	}
	tree.honestSmallStepChalLevelZeroEdge = util.Some(topLevel)
	child08 := &edge{
		id:          id("smol-0-8"),
		edgeType:    protocol.SmallStepChallengeEdge,
		startHeight: 0,
		endHeight:   8,
	}
	child816 := &edge{
		id:          id("smol-8-16"),
		edgeType:    protocol.SmallStepChallengeEdge,
		startHeight: 8,
		endHeight:   16,
	}
	topLevel.lowerChildId = common.Hash(child08.id)
	topLevel.upperChildId = common.Hash(child816.id)

	child04 := &edge{
		id:          id("smol-0-4"),
		edgeType:    protocol.SmallStepChallengeEdge,
		startHeight: 0,
		endHeight:   4,
	}
	child48 := &edge{
		id:          id("smol-4-8"),
		edgeType:    protocol.SmallStepChallengeEdge,
		startHeight: 4,
		endHeight:   8,
	}
	child08.lowerChildId = common.Hash(child04.id)
	child08.upperChildId = common.Hash(child48.id)

	child46 := &edge{
		id:          id("smol-4-6"),
		edgeType:    protocol.SmallStepChallengeEdge,
		startHeight: 4,
		endHeight:   6,
	}
	child68 := &edge{
		id:          id("smol-6-8"),
		edgeType:    protocol.SmallStepChallengeEdge,
		startHeight: 6,
		endHeight:   8,
	}
	child48.lowerChildId = common.Hash(child46.id)
	child48.upperChildId = common.Hash(child68.id)

	child45 := &edge{
		id:          id("smol-4-5"),
		edgeType:    protocol.SmallStepChallengeEdge,
		startHeight: 4,
		endHeight:   5,
	}
	child56 := &edge{
		id:          id("smol-5-6"),
		edgeType:    protocol.SmallStepChallengeEdge,
		startHeight: 5,
		endHeight:   6,
	}
	child46.lowerChildId = common.Hash(child45.id)
	child46.upperChildId = common.Hash(child56.id)

	tree.edges.Insert(topLevel.id, topLevel)
	tree.edges.Insert(child08.id, child08)
	tree.edges.Insert(child816.id, child816)
	tree.edges.Insert(child04.id, child04)
	tree.edges.Insert(child48.id, child48)
	tree.edges.Insert(child46.id, child46)
	tree.edges.Insert(child68.id, child68)
	tree.edges.Insert(child45.id, child45)
	tree.edges.Insert(child56.id, child56)

	// Rival all edges except for 0-4
	tree.rivaledEdges.Insert(topLevel.id)
	tree.rivaledEdges.Insert(child08.id)
	tree.rivaledEdges.Insert(child816.id)
	tree.rivaledEdges.Insert(child48.id)
	tree.rivaledEdges.Insert(child46.id)
	tree.rivaledEdges.Insert(child68.id)
	tree.rivaledEdges.Insert(child45.id)
	tree.rivaledEdges.Insert(child56.id)
}

// Helper function for readable ids. Ids that belong to block challenges should be prefixed with "blk".
// For big step, prefixed with "big", and small step, prefixed with "smol".
func id(s string) protocol.EdgeId {
	return protocol.EdgeId(common.BytesToHash([]byte(s)))
}
