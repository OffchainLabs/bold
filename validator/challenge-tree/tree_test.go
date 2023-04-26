package challengetree

import (
	"testing"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/OffchainLabs/challenge-protocol-v2/util/threadsafe"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func id(s string) protocol.EdgeId {
	return protocol.EdgeId(common.BytesToHash([]byte(s)))
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
//		0-----------------------8, 8----------------------16 (claim_id = id(4,5) in the level above)
//		0-----------------------8
//		0-------4, 4------------8
//		4----6, 6----8
//		4--5, 5--6
func TestAncestors(t *testing.T) {
	tree := &challengeTree{
		edges: threadsafe.NewMap[protocol.EdgeId, *edge](),
	}
	setupBlockChallengeTreeSnapshot(t, tree)

	// Edge ids that belong to block challenges are prefixed with "blk".
	// For big step, prefixed with "big", and small step, prefixed with "smol".
	ancestors := tree.ancestorsForHonestEdge(id("blk-6-8"))
	require.Equal(t, ancestors, []protocol.EdgeId{id("blk-4-8"), id("blk-0-8"), id("blk-0-16")})

	ancestors = tree.ancestorsForHonestEdge(id("blk-4-6"))
	require.Equal(t, ancestors, []protocol.EdgeId{id("blk-4-8"), id("blk-0-8"), id("blk-0-16")})

	ancestors = tree.ancestorsForHonestEdge(id("blk-0-4"))
	require.Equal(t, ancestors, []protocol.EdgeId{id("blk-0-8"), id("blk-0-16")})

	ancestors = tree.ancestorsForHonestEdge(id("blk-4-8"))
	require.Equal(t, ancestors, []protocol.EdgeId{id("blk-0-8"), id("blk-0-16")})

	ancestors = tree.ancestorsForHonestEdge(id("blk-0-16"))
	require.Equal(t, 0, len(ancestors))

	ancestors = tree.ancestorsForHonestEdge(id("foo"))
	require.Equal(t, 0, len(ancestors))
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
		id:       id("0-16"),
		edgeType: protocol.BlockChallengeEdge,
	}
	tree.honestBlockChalLevelZeroEdge = util.Some(topLevel)
	child08 := &edge{
		id:       id("0-8"),
		edgeType: protocol.BlockChallengeEdge,
	}
	child816 := &edge{
		id:       id("8-16"),
		edgeType: protocol.BlockChallengeEdge,
	}
	topLevel.lowerChildId = common.Hash(child08.id)
	topLevel.upperChildId = common.Hash(child816.id)

	child04 := &edge{
		id:       id("0-4"),
		edgeType: protocol.BlockChallengeEdge,
	}
	child48 := &edge{
		id:       id("4-8"),
		edgeType: protocol.BlockChallengeEdge,
	}
	child08.lowerChildId = common.Hash(child04.id)
	child08.upperChildId = common.Hash(child48.id)

	child46 := &edge{
		id:       id("4-6"),
		edgeType: protocol.BlockChallengeEdge,
	}
	child68 := &edge{
		id:       id("6-8"),
		edgeType: protocol.BlockChallengeEdge,
	}
	child48.lowerChildId = common.Hash(child46.id)
	child48.upperChildId = common.Hash(child68.id)

	tree.edges.Insert(topLevel.id, topLevel)
	tree.edges.Insert(child08.id, child08)
	tree.edges.Insert(child816.id, child816)
	tree.edges.Insert(child04.id, child04)
	tree.edges.Insert(child48.id, child48)
	tree.edges.Insert(child46.id, child46)
	tree.edges.Insert(child68.id, child68)
}
