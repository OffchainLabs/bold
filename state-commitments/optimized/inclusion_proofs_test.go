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
	rehashed := crypto.Keccak256Hash(simpleHash.Bytes())
	left := crypto.Keccak256Hash(rehashed.Bytes(), rehashed.Bytes())
	right := crypto.Keccak256Hash(rehashed.Bytes(), rehashed.Bytes())
	total := crypto.Keccak256Hash(left.Bytes(), right.Bytes())
	_ = total
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
