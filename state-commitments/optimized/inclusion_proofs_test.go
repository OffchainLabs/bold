package optimized

import (
	"testing"

	"github.com/OffchainLabs/bold/state-commitments/history"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestInclusionProofEquivalence(t *testing.T) {
	simpleHash := crypto.Keccak256Hash([]byte("foo"))
	leaves := []common.Hash{
		simpleHash,
		simpleHash,
		simpleHash,
		simpleHash,
	}
	commit, err := NewCommitment(leaves, 4)
	require.NoError(t, err)
	oldLeaves := []common.Hash{
		simpleHash,
		simpleHash,
		simpleHash,
		simpleHash,
	}
	oldCommit, err := history.New(oldLeaves)
	require.NoError(t, err)
	require.Equal(t, commit.Merkle, oldCommit.Merkle)
}
