package optimized

import (
	"fmt"
	"testing"

	"github.com/OffchainLabs/bold/state-commitments/history"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestLegacyVsOptimized(t *testing.T) {
	end := 1 << 9
	simpleHash := crypto.Keccak256Hash([]byte("foo"))
	for i := 1; i < end; i++ {
		limit := nextPowerOf2(i)
		for j := i; j < limit; j++ {
			hashedLeaves := make([]common.Hash, i)
			for i := range hashedLeaves {
				hashedLeaves[i] = crypto.Keccak256Hash(simpleHash[:])
			}
			_, err := computeVirtualSparseTree(hashedLeaves, j, 0)
			require.NoError(t, err)

			hashedLeaves = make([]common.Hash, i)
			for i := range hashedLeaves {
				hashedLeaves[i] = crypto.Keccak256Hash(simpleHash[:])
			}
			computedRoot, err := computeVirtualSparseTree(hashedLeaves, j, limit)
			require.NoError(t, err)
			legacyInputLeaves := make([]common.Hash, j)
			for i := range legacyInputLeaves {
				legacyInputLeaves[i] = simpleHash
			}
			histCommit, err := history.New(legacyInputLeaves)
			require.NoError(t, err)
			require.Equal(t, computedRoot, histCommit.Merkle)
		}
	}
}

func TestLegacyVsOptimizedEdgeCases(t *testing.T) {
	simpleHash := crypto.Keccak256Hash([]byte("foo"))

	tests := []struct {
		realLength    int
		virtualLength int
		limit         int
	}{
		{12, 14, 16},
		{8, 10, 16},
		{6, 6, 8},
		{10, 16, 16},
		{4, 8, 8},
		{1, 5, 8},
		{3, 5, 8},
		{5, 5, 8},
		{1023, 1024, 1024},
		{(1 << 14) - 7, (1 << 14) - 7, 1 << 14},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("real length %d, virtual %d, limit %d", tt.realLength, tt.virtualLength, tt.limit), func(t *testing.T) {
			hashedLeaves := make([]common.Hash, tt.realLength)
			for i := range hashedLeaves {
				hashedLeaves[i] = crypto.Keccak256Hash(simpleHash[:])
			}
			_, err := computeVirtualSparseTree(hashedLeaves, tt.virtualLength, 0)
			require.NoError(t, err)

			hashedLeaves = make([]common.Hash, tt.realLength)
			for i := range hashedLeaves {
				hashedLeaves[i] = crypto.Keccak256Hash(simpleHash[:])
			}
			computedRoot, err := computeVirtualSparseTree(hashedLeaves, tt.virtualLength, tt.limit)
			require.NoError(t, err)

			leaves := make([]common.Hash, tt.virtualLength)
			for i := range leaves {
				leaves[i] = simpleHash
			}
			histCommit, err := history.New(leaves)
			require.NoError(t, err)
			require.Equal(t, computedRoot, histCommit.Merkle)
		})
	}
}
func TestVirtualSparse(t *testing.T) {
	simpleHash := crypto.Keccak256Hash([]byte("foo"))
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
		require.Equal(t, computedRoot, histCommit.Merkle)
	})
	t.Run("real length 12, virtual length 14, limit 16", func(t *testing.T) {
		hashedLeaves := make([]common.Hash, 12)
		for i := range hashedLeaves {
			hashedLeaves[i] = crypto.Keccak256Hash(simpleHash[:])
		}
		_, err := computeVirtualSparseTree(hashedLeaves, 14, 0)
		require.NoError(t, err)
		hashedLeaves = make([]common.Hash, 12)
		for i := range hashedLeaves {
			hashedLeaves[i] = crypto.Keccak256Hash(simpleHash[:])
		}
		computedRoot, err := computeVirtualSparseTree(hashedLeaves, 14, 16)
		require.NoError(t, err)
		leaves := make([]common.Hash, 14)
		for i := range leaves {
			leaves[i] = simpleHash
		}
		histCommit, err := history.New(leaves)
		require.NoError(t, err)
		require.Equal(t, computedRoot, histCommit.Merkle)
	})
}

func TestMaximumDepthHistoryCommitment(t *testing.T) {
	simpleHash := crypto.Keccak256Hash([]byte("foo"))
	hashedLeaves := []common.Hash{
		crypto.Keccak256Hash(simpleHash[:]),
	}
	_, err := computeVirtualSparseTree(hashedLeaves, 1<<26, 0)
	require.NoError(t, err)
	hashedLeaves = []common.Hash{
		crypto.Keccak256Hash(simpleHash[:]),
	}
	_, err = computeVirtualSparseTree(hashedLeaves, 1<<26, 1<<26)
	require.NoError(t, err)
}

func BenchmarkMaximumDepthHistoryCommitment(b *testing.B) {
	b.StopTimer()
	simpleHash := crypto.Keccak256Hash([]byte("foo"))
	hashedLeaves := []common.Hash{
		crypto.Keccak256Hash(simpleHash[:]),
	}
	_, err := computeVirtualSparseTree(hashedLeaves, 1<<26, 0)
	require.NoError(b, err)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		hashedLeaves = []common.Hash{
			crypto.Keccak256Hash(simpleHash[:]),
		}
		computeVirtualSparseTree(hashedLeaves, 1<<26, 1<<26)
	}
}
