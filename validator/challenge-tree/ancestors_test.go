package challengetree

import (
	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"testing"
)

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
	// tree := &challengeTree{
	// 	edges:        threadsafe.NewMap[protocol.EdgeId, *edge](),
	// 	rivaledEdges: threadsafe.NewSet[protocol.EdgeId](),
	// }
	// setupBlockChallengeTreeSnapshot(t, tree)
	// claimId := protocol.ClaimId(id("blk-5-6"))
	// setupBigStepChallengeSnapshot(t, tree, claimId)
	// claimId = protocol.ClaimId(id("big-5-6"))
	// setupSmallStepChallengeSnapshot(t, tree, claimId)

	// // We start by querying for ancestors for a big step edge id.
	// ancestors := tree.ancestorsForHonestEdge(id("big-4-5"))

	// // Edge ids that belong to block challenges are prefixed with "blk".
	// // For big step, prefixed with "big", and small step, prefixed with "smol".
	// wanted := []protocol.EdgeId{
	// 	id("big-4-6"),
	// 	id("big-4-8"),
	// 	id("big-0-8"),
	// 	id("big-0-16"),
	// 	id("blk-5-6"), // TODO: Should the claim id be part of the ancestors as well?
	// 	id("blk-4-6"),
	// 	id("blk-4-8"),
	// 	id("blk-0-8"),
	// 	id("blk-0-16"),
	// }
	// require.Equal(t, wanted, ancestors)

	// // We start query the ancestors of the lowest level, length one, small step edge.
	// ancestors = tree.ancestorsForHonestEdge(id("smol-5-6"))
	// wanted = []protocol.EdgeId{
	// 	id("smol-4-6"),
	// 	id("smol-4-8"),
	// 	id("smol-0-8"),
	// 	id("smol-0-16"),
	// 	id("big-5-6"),
	// 	id("big-4-6"),
	// 	id("big-4-8"),
	// 	id("big-0-8"),
	// 	id("big-0-16"),
	// 	id("blk-5-6"), // TODO: Should the claim id be part of the ancestors as well?
	// 	id("blk-4-6"),
	// 	id("blk-4-8"),
	// 	id("blk-0-8"),
	// 	id("blk-0-16"),
	// }
	// require.Equal(t, wanted, ancestors)

	// // Query the level zero edge at each challenge type.
	// ancestors = tree.ancestorsForHonestEdge(id("blk-0-16"))
	// require.Equal(t, 0, len(ancestors))

	// ancestors = tree.ancestorsForHonestEdge(id("big-0-16"))
	// require.Equal(t, ancestors, []protocol.EdgeId{
	// 	id("blk-5-6"),
	// 	id("blk-4-6"),
	// 	id("blk-4-8"),
	// 	id("blk-0-8"),
	// 	id("blk-0-16"),
	// })

	// ancestors = tree.ancestorsForHonestEdge(id("smol-0-16"))
	// require.Equal(t, ancestors, []protocol.EdgeId{
	// 	id("big-5-6"),
	// 	id("big-4-6"),
	// 	id("big-4-8"),
	// 	id("big-0-8"),
	// 	id("big-0-16"),
	// 	id("blk-5-6"),
	// 	id("blk-4-6"),
	// 	id("blk-4-8"),
	// 	id("blk-0-8"),
	// 	id("blk-0-16"),
	// })
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
func setupBlockChallengeTreeSnapshot(t *testing.T, tree *HonestChallengeTree) {
	t.Helper()
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
func setupBigStepChallengeSnapshot(t *testing.T, tree *HonestChallengeTree, claimId protocol.ClaimId) {
	t.Helper()
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
func setupSmallStepChallengeSnapshot(t *testing.T, tree *HonestChallengeTree, claimId protocol.ClaimId) {
	t.Helper()
}
