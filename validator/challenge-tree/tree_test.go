package challengetree

import (
	"testing"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/OffchainLabs/challenge-protocol-v2/util/threadsafe"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestAncestors(t *testing.T) {
	// Tests the following set of honest edges for ancestor queries.
	//
	// 0-----------------------8, 8----------------------16
	// 0-----------------------8
	// 0-------4, 4------------8
	//            4----6, 6----8
	//
	topLevel := &edge{
		id:       protocol.EdgeId(common.BytesToHash([]byte("0-16"))),
		edgeType: protocol.BlockChallengeEdge,
	}
	child08 := &edge{
		id:       protocol.EdgeId(common.BytesToHash([]byte("0-8"))),
		edgeType: protocol.BlockChallengeEdge,
	}
	child816 := &edge{
		id:       protocol.EdgeId(common.BytesToHash([]byte("8-16"))),
		edgeType: protocol.BlockChallengeEdge,
	}
	topLevel.lowerChildId = common.Hash(child08.id)
	topLevel.upperChildId = common.Hash(child816.id)

	child04 := &edge{
		id:       protocol.EdgeId(common.BytesToHash([]byte("0-4"))),
		edgeType: protocol.BlockChallengeEdge,
	}
	child48 := &edge{
		id:       protocol.EdgeId(common.BytesToHash([]byte("4-8"))),
		edgeType: protocol.BlockChallengeEdge,
	}
	child08.lowerChildId = common.Hash(child04.id)
	child08.upperChildId = common.Hash(child48.id)

	child46 := &edge{
		id:       protocol.EdgeId(common.BytesToHash([]byte("4-6"))),
		edgeType: protocol.BlockChallengeEdge,
	}
	child68 := &edge{
		id:       protocol.EdgeId(common.BytesToHash([]byte("6-8"))),
		edgeType: protocol.BlockChallengeEdge,
	}
	child48.lowerChildId = common.Hash(child46.id)
	child48.upperChildId = common.Hash(child68.id)

	tree := &challengeTree{
		honestBlockChalLevelZeroEdge: util.Some(topLevel),
		edges:                        threadsafe.NewMap[protocol.EdgeId, *edge](),
	}

	tree.edges.Insert(topLevel.id, topLevel)
	tree.edges.Insert(child08.id, child08)
	tree.edges.Insert(child816.id, child816)
	tree.edges.Insert(child04.id, child04)
	tree.edges.Insert(child48.id, child48)
	tree.edges.Insert(child46.id, child46)
	tree.edges.Insert(child68.id, child68)

	ancestors := tree.ancestorsForHonestEdge(child68.id)
	require.Equal(t, ancestors, []protocol.EdgeId{child48.id, child08.id, topLevel.id})

	ancestors = tree.ancestorsForHonestEdge(child46.id)
	require.Equal(t, ancestors, []protocol.EdgeId{child48.id, child08.id, topLevel.id})

	ancestors = tree.ancestorsForHonestEdge(child04.id)
	require.Equal(t, ancestors, []protocol.EdgeId{child08.id, topLevel.id})

	ancestors = tree.ancestorsForHonestEdge(child48.id)
	require.Equal(t, ancestors, []protocol.EdgeId{child08.id, topLevel.id})

	ancestors = tree.ancestorsForHonestEdge(topLevel.id)
	require.Equal(t, 0, len(ancestors))

	ancestors = tree.ancestorsForHonestEdge(protocol.EdgeId(common.BytesToHash([]byte("foo"))))
	require.Equal(t, 0, len(ancestors))
}
