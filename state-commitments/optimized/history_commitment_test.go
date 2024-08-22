package optimized

import (
	"errors"
	"testing"

	"github.com/OffchainLabs/bold/state-commitments/history"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestVirtualSparse(t *testing.T) {
	simpleHash := crypto.Keccak256Hash([]byte("foo"))
	// Now compare against a history commitment implementation.
	t.Run("real length 1, virtual length 3, limit 4", func(t *testing.T) {
		_, err := computeVirtualSparseTree([]common.Hash{crypto.Keccak256Hash(simpleHash[:])}, 3, 0)
		require.NoError(t, err)
		computedRoot, err := computeVirtualSparseTree([]common.Hash{crypto.Keccak256Hash(simpleHash[:])}, 3, 4)
		require.NoError(t, err)
		leaves := []common.Hash{
			simpleHash,
			simpleHash,
			simpleHash,
		}
		histCommit, err := history.New(leaves)
		require.NoError(t, err)
		require.Equal(t, histCommit.Merkle, computedRoot)
	})
	t.Run("real length 2, virtual length 3, limit 4", func(t *testing.T) {
		hashedLeaves := []common.Hash{
			crypto.Keccak256Hash(simpleHash[:]),
			crypto.Keccak256Hash(simpleHash[:]),
		}
		_, err := computeVirtualSparseTree(hashedLeaves, 3, 0)
		require.NoError(t, err)
		hashedLeaves = []common.Hash{
			crypto.Keccak256Hash(simpleHash[:]),
			crypto.Keccak256Hash(simpleHash[:]),
		}
		computedRoot, err := computeVirtualSparseTree(hashedLeaves, 3, 4)
		require.NoError(t, err)
		leaves := []common.Hash{
			simpleHash,
			simpleHash,
			simpleHash,
		}
		histCommit, err := history.New(leaves)
		require.NoError(t, err)
		require.Equal(t, histCommit.Merkle, computedRoot)
	})
	t.Run("real length 3, virtual length 3, limit 4", func(t *testing.T) {
		hashedLeaves := []common.Hash{
			crypto.Keccak256Hash(simpleHash[:]),
			crypto.Keccak256Hash(simpleHash[:]),
			crypto.Keccak256Hash(simpleHash[:]),
		}
		_, err := computeVirtualSparseTree(hashedLeaves, 3, 0)
		require.NoError(t, err)
		hashedLeaves = []common.Hash{
			crypto.Keccak256Hash(simpleHash[:]),
			crypto.Keccak256Hash(simpleHash[:]),
			crypto.Keccak256Hash(simpleHash[:]),
		}
		computedRoot, err := computeVirtualSparseTree(hashedLeaves, 3, 4)
		require.NoError(t, err)
		leaves := []common.Hash{
			simpleHash,
			simpleHash,
			simpleHash,
		}
		histCommit, err := history.New(leaves)
		require.NoError(t, err)
		require.Equal(t, histCommit.Merkle, computedRoot)
	})
	t.Run("real length 4, virtual length 4, limit 4", func(t *testing.T) {
		hashedLeaves := []common.Hash{
			crypto.Keccak256Hash(simpleHash[:]),
			crypto.Keccak256Hash(simpleHash[:]),
			crypto.Keccak256Hash(simpleHash[:]),
			crypto.Keccak256Hash(simpleHash[:]),
		}
		_, err := computeVirtualSparseTree(hashedLeaves, 4, 0)
		require.NoError(t, err)
		hashedLeaves = []common.Hash{
			crypto.Keccak256Hash(simpleHash[:]),
			crypto.Keccak256Hash(simpleHash[:]),
			crypto.Keccak256Hash(simpleHash[:]),
			crypto.Keccak256Hash(simpleHash[:]),
		}
		computedRoot, err := computeVirtualSparseTree(hashedLeaves, 4, 4)
		require.NoError(t, err)
		leaves := []common.Hash{
			simpleHash,
			simpleHash,
			simpleHash,
			simpleHash,
		}
		histCommit, err := history.New(leaves)
		require.NoError(t, err)
		require.Equal(t, histCommit.Merkle, computedRoot)
	})
	t.Run("real length 1, virtual length 5, limit 8", func(t *testing.T) {
		hashedLeaves := []common.Hash{
			crypto.Keccak256Hash(simpleHash[:]),
		}
		_, err := computeVirtualSparseTree(hashedLeaves, 5, 0)
		require.NoError(t, err)
		hashedLeaves = []common.Hash{
			crypto.Keccak256Hash(simpleHash[:]),
		}
		computedRoot, err := computeVirtualSparseTree(hashedLeaves, 5, 8)
		require.NoError(t, err)
		leaves := []common.Hash{
			simpleHash,
			simpleHash,
			simpleHash,
			simpleHash,
			simpleHash,
		}
		histCommit, err := history.New(leaves)
		require.NoError(t, err)
		t.Log(computedRoot, histCommit.Merkle)
		require.Equal(t, histCommit.Merkle, computedRoot)
	})
}

func TestMaximumDepthHistoryCommitment(t *testing.T) {
	simpleHash := crypto.Keccak256Hash([]byte("foo"))
	_, err := computeVirtualSparseTree([]common.Hash{simpleHash}, 1<<26, 0)
	require.NoError(t, err)
	_, err = computeVirtualSparseTree([]common.Hash{simpleHash}, 1<<26, 1<<26)
	require.NoError(t, err)
}

func BenchmarkMaximumDepthHistoryCommitment(b *testing.B) {
	b.StopTimer()
	simpleHash := crypto.Keccak256Hash([]byte("foo"))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		computeVirtualSparseTree([]common.Hash{simpleHash}, 1<<26, 1<<26)
	}
}

// GenerateTrieFromItems constructs a Merkle trie from a sequence of byte slices.
func computeSparseRoot(leaves []common.Hash, depth uint64) (common.Hash, error) {
	var emptyHash common.Hash
	if len(leaves) == 0 {
		return emptyHash, errors.New("no items provided to generate Merkle trie")
	}
	if depth >= 26 {
		return emptyHash, errors.New("supported merkle trie depth exceeded (max depth is 26)")
	}
	layers := make([][]common.Hash, depth+1)
	layers[0] = leaves
	for i := uint64(0); i < depth; i++ {
		if len(layers[i])%2 == 1 {
			layers[i] = append(layers[i], zeroHashes[i])
		}
		updatedValues := make([]common.Hash, 0)
		for j := 0; j < len(layers[i]); j += 2 {
			concat := crypto.Keccak256Hash(layers[i][j][:], layers[i][j+1][:])
			updatedValues = append(updatedValues, concat)
		}
		layers[i+1] = updatedValues
	}
	return layers[len(layers)-1][0], nil
}
