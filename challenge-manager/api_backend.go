package challengemanager

import (
	"context"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/ethereum/go-ethereum/common"
)

// 1. Get ongoing challenges (list of parent assertion hashes)
// 2. Get challenge activity by hash
// 3. Within the challenge by hash:
//     1. For each edge tracked by the chain watcher service in the BOLD validator:
//     2. Extract edges that are rivaled and unrivaled
//     3. Recover the Ethereum addresses of the tx creators of all these honest edges
//     4. Check how long it has been since each rivaled edge was created
//     5. Check how long it has been since each honest actor by Ethereum address has made a move
//     6. From this, the operator can determine if there are some rivaled edges not being taken care of, can also determine which honest actor is most recent and which ones have dropped off

type HonestEdgeJson struct {
	Id                      common.Hash    `json:"id"`
	CreatorAddress          common.Address `json:"creator_address"`
	FirstRivalCreationBlock uint64         `json:"first_rival_creation_block"`
	CreationBlock           uint64         `json:"creation_block"`
	UpperChild              common.Hash    `json:"upper_child"`
	LowerChild              common.Hash    `json:"lower_child"`
	UpperChildCreationBlock common.Hash    `json:"upper_child_creation_block"`
	LowerChildCreationBlock common.Hash    `json:"lower_child_creation_block"`
	HasClaimingSubchallenge bool           `json:"has_claiming_subchallenge"`
}

type HonestActivityJson struct {
	ChallengedParentAssertionHash common.Hash       `json:"challenged_parent_assertion_hash"`
	RivaledEdges                  []*HonestEdgeJson `json:"rivaled_edges"`
	UnrivaledEdges                []*HonestEdgeJson `json:"unrivaled_edges"`
}

func (m *Manager) queryHonestActivity(
	ctx context.Context, challengedParentAssertionHash protocol.AssertionHash,
) (*HonestActivityJson, error) {
	honestEdges, err := m.watcher.GetHonestEdgesByChallenge(challengedParentAssertionHash)
	if err != nil {
		return nil, err
	}
	response := &HonestActivityJson{
		ChallengedParentAssertionHash: challengedParentAssertionHash.Hash,
		RivaledEdges:                  make([]*HonestEdgeJson, 0),
		UnrivaledEdges:                make([]*HonestEdgeJson, 0),
	}
	for _, edge := range honestEdges {
		isRivaled, err := edge.HasRival(ctx)
		if err != nil {
			return nil, err
		}
		if isRivaled {
			// rivaled = append(rivaled, edge)
		} else {
			// unrivaled = append(unrivaled, edge)
		}
	}
	return response, nil
}
