package challengetree

import (
	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/OffchainLabs/challenge-protocol-v2/util/threadsafe"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDirectPathAncestry(t *testing.T) {
	tree := &HonestChallengeTree{
		edges:     threadsafe.NewMap[protocol.EdgeId, protocol.EdgeSnapshot](),
		mutualIds: threadsafe.NewMap[protocol.MutualId, *threadsafe.Set[protocol.EdgeId]](),
	}
	// Edge ids that belong to block challenges are prefixed with "blk".
	// For big step, prefixed with "big", and small step, prefixed with "smol".
	setupBlockChallengeTreeSnapshot(t, tree)
	tree.honestBlockChalLevelZeroEdge = util.Some(tree.edges.Get(id("blk-0.a-16.a")))
	claimId := "blk-4.a-5.a"
	setupBigStepChallengeSnapshot(t, tree, claimId)
	tree.honestBigStepLevelZeroEdges = append(tree.honestBigStepLevelZeroEdges, tree.edges.Get(id("big-0.a-16.a")))
	claimId = "big-4.a-5.a"
	setupSmallStepChallengeSnapshot(t, tree, claimId)
	tree.honestSmallStepLevelZeroEdges = append(tree.honestSmallStepLevelZeroEdges, tree.edges.Get(id("smol-0.a-16.a")))

	ancestors, err := tree.ancestryPath(id("blk-5.a-6.a"))
	require.NoError(t, err)
	for _, an := range ancestors {
		t.Logf("%s", an)
	}
	t.Log(ancestors)

	t.Log("********************")
	ancestors, err = tree.ancestryPath(id("big-5.a-6.a"))
	require.NoError(t, err)
	for _, an := range ancestors {
		t.Logf("%s", an)
	}
	t.Log(ancestors)
}
