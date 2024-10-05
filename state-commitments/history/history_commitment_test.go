package history

import (
	"crypto/rand"
	"fmt"
	"testing"

	prefixproofs "github.com/OffchainLabs/bold/state-commitments/prefix-proofs"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func FuzzHistoryCommitter(f *testing.F) {
	simpleHash := crypto.Keccak256Hash([]byte("foo"))
	f.Fuzz(func(t *testing.T, numReal uint64, virtual uint64, limit uint64) {
		// Set some bounds.
		numReal = numReal % (1 << 10)
		virtual = virtual % (1 << 20)
		hashedLeaves := make([]common.Hash, numReal)
		for i := range hashedLeaves {
			hashedLeaves[i] = simpleHash
		}
		committer := NewCommitter()
		_, err := committer.ComputeRoot(hashedLeaves, virtual)
		_ = err
	})
}

func BenchmarkPrefixProofGeneration_Legacy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prefixIndex := 13384
		simpleHash := crypto.Keccak256Hash([]byte("foo"))
		hashes := make([]common.Hash, 1<<14)
		for i := 0; i < len(hashes); i++ {
			hashes[i] = simpleHash
		}

		lowCommitmentNumLeaves := prefixIndex + 1
		hiCommitmentNumLeaves := (1 << 14)
		prefixExpansion, err := prefixproofs.ExpansionFromLeaves(hashes[:lowCommitmentNumLeaves])
		require.NoError(b, err)
		_, err = prefixproofs.GeneratePrefixProof(
			uint64(lowCommitmentNumLeaves),
			prefixExpansion,
			hashes[lowCommitmentNumLeaves:hiCommitmentNumLeaves],
			prefixproofs.RootFetcherFromExpansion,
		)
		require.NoError(b, err)
	}
}

func BenchmarkPrefixProofGeneration_Optimized(b *testing.B) {
	b.StopTimer()
	simpleHash := crypto.Keccak256Hash([]byte("foo"))
	hashes := []common.Hash{crypto.Keccak256Hash(simpleHash[:])}
	prefixIndex := uint64(13384)
	virtual := uint64(1 << 14)
	committer := NewCommitter()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _, err := committer.GeneratePrefixProof(prefixIndex, hashes, virtual)
		require.NoError(b, err)
	}
}

func TestSimpleHistoryCommitment(t *testing.T) {
	aLeaf := common.HexToHash("0xA")
	bLeaf := common.HexToHash("0xB")
	// Level 0
	aHash := crypto.Keccak256Hash(aLeaf[:])
	bHash := crypto.Keccak256Hash(bLeaf[:])
	// Level 1
	abHash := crypto.Keccak256Hash(append(aHash[:], bHash[:]...))
	bzHash := crypto.Keccak256Hash(append(bHash[:], emptyHash[:]...))
	bbHash := crypto.Keccak256Hash(append(bHash[:], bHash[:]...))
	// Level 2
	abbzHash := crypto.Keccak256Hash(append(abHash[:], bzHash[:]...))
	abbbHash := crypto.Keccak256Hash(append(abHash[:], bbHash[:]...))
	tests := []struct {
		name string
		lvs  []common.Hash
		virt uint64
		want common.Hash
	}{
		{
			name: "empty leaves",
			lvs:  []common.Hash{},
			virt: 0,
			want: emptyHash,
		},
		{
			name: "single leaf",
			lvs:  []common.Hash{aLeaf},
			virt: 1,
			want: aHash,
		},
		{
			name: "two leaves",
			lvs:  []common.Hash{aLeaf, bLeaf},
			virt: 2,
			want: abHash,
		},
		{
			name: "two leaves - virtual 3",
			lvs:  []common.Hash{aLeaf, bLeaf},
			virt: 3,
			want: abbzHash,
		},
		{
			name: "two leaves - virtual 4",
			lvs:  []common.Hash{aLeaf, bLeaf},
			virt: 4,
			want: abbbHash,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			hc := NewCommitter()
			got, err := hc.ComputeRoot(tc.lvs, tc.virt)
			if err != nil {
				t.Errorf("ComputeRoot(%v, %d): err %v", tc.lvs, tc.virt, err)
			}
			if got != tc.want {
				t.Errorf("ComputeRoot(%v, %d): got %s, want %s", tc.lvs, tc.virt, got.Hex(), tc.want.Hex())
			}
		})
	}
}

func TestLegacyVsOptimized(t *testing.T) {
	t.Parallel()
	end := uint64(1 << 9)
	simpleHash := crypto.Keccak256Hash([]byte("foo"))
	for i := uint64(1); i < end; i++ {
		limit := nextPowerOf2(i)
		for j := i; j < limit; j++ {
			inputLeaves := make([]common.Hash, i)
			for i := range inputLeaves {
				inputLeaves[i] = simpleHash
			}
			committer := NewCommitter()
			computedRoot, err := committer.ComputeRoot(inputLeaves, uint64(j))
			require.NoError(t, err)

			legacyInputLeaves := make([]common.Hash, j)
			for i := range legacyInputLeaves {
				legacyInputLeaves[i] = simpleHash
			}
			histCommit, err := NewLegacy(legacyInputLeaves)
			require.NoError(t, err)
			require.Equal(t, computedRoot, histCommit.Merkle)
		}
	}
}

func TestLegacyVsOptimizedEdgeCases(t *testing.T) {
	t.Parallel()
	simpleHash := crypto.Keccak256Hash([]byte("foo"))

	tests := []struct {
		realLength    int
		virtualLength int
	}{
		{12, 14},
		{8, 10},
		{6, 6},
		{10, 16},
		{4, 8},
		{1, 5},
		{3, 5},
		{5, 5},
		{1023, 1024},
		{(1 << 14) - 7, (1 << 14) - 7},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("real length %d, virtual %d", tt.realLength, tt.virtualLength), func(t *testing.T) {
			inputLeaves := make([]common.Hash, tt.realLength)
			for i := range inputLeaves {
				inputLeaves[i] = simpleHash
			}
			committer := NewCommitter()
			computedRoot, err := committer.ComputeRoot(inputLeaves, uint64(tt.virtualLength))
			require.NoError(t, err)

			leaves := make([]common.Hash, tt.virtualLength)
			for i := range leaves {
				leaves[i] = simpleHash
			}
			histCommit, err := NewLegacy(leaves)
			require.NoError(t, err)
			require.Equal(t, computedRoot, histCommit.Merkle)
		})
	}
}

func TestVirtualSparse(t *testing.T) {
	t.Parallel()
	simpleHash := crypto.Keccak256Hash([]byte("foo"))
	t.Run("real length 1, virtual length 3", func(t *testing.T) {
		committer := NewCommitter()
		computedRoot, err := committer.ComputeRoot([]common.Hash{simpleHash}, 3)
		require.NoError(t, err)

		leaves := []common.Hash{
			simpleHash,
			simpleHash,
			simpleHash,
		}
		histCommit, err := NewLegacy(leaves)
		require.NoError(t, err)
		require.Equal(t, histCommit.Merkle, computedRoot)
	})
	t.Run("real length 2, virtual length 3", func(t *testing.T) {
		hashedLeaves := []common.Hash{
			simpleHash,
			simpleHash,
		}
		committer := NewCommitter()
		computedRoot, err := committer.ComputeRoot(hashedLeaves, 3)
		require.NoError(t, err)
		leaves := []common.Hash{
			simpleHash,
			simpleHash,
			simpleHash,
		}
		histCommit, err := NewLegacy(leaves)
		require.NoError(t, err)
		require.Equal(t, histCommit.Merkle, computedRoot)
	})
	t.Run("real length 3, virtual length 3", func(t *testing.T) {
		hashedLeaves := []common.Hash{
			simpleHash,
			simpleHash,
			simpleHash,
		}
		committer := NewCommitter()
		computedRoot, err := committer.ComputeRoot(hashedLeaves, 3)
		require.NoError(t, err)
		leaves := []common.Hash{
			simpleHash,
			simpleHash,
			simpleHash,
		}
		histCommit, err := NewLegacy(leaves)
		require.NoError(t, err)
		require.Equal(t, histCommit.Merkle, computedRoot)
	})
	t.Run("real length 4, virtual length 4", func(t *testing.T) {
		hashedLeaves := []common.Hash{
			simpleHash,
			simpleHash,
			simpleHash,
			simpleHash,
		}
		committer := NewCommitter()
		computedRoot, err := committer.ComputeRoot(hashedLeaves, 4)
		require.NoError(t, err)
		leaves := []common.Hash{
			simpleHash,
			simpleHash,
			simpleHash,
			simpleHash,
		}
		histCommit, err := NewLegacy(leaves)
		require.NoError(t, err)
		require.Equal(t, histCommit.Merkle, computedRoot)
	})
	t.Run("real length 1, virtual length 5", func(t *testing.T) {
		hashedLeaves := []common.Hash{
			simpleHash,
		}
		committer := NewCommitter()
		computedRoot, err := committer.ComputeRoot(hashedLeaves, 5)
		require.NoError(t, err)

		leaves := []common.Hash{
			simpleHash,
			simpleHash,
			simpleHash,
			simpleHash,
			simpleHash,
		}
		histCommit, err := NewLegacy(leaves)
		require.NoError(t, err)
		require.Equal(t, computedRoot, histCommit.Merkle)
	})
	t.Run("real length 12, virtual length 14", func(t *testing.T) {
		hashedLeaves := make([]common.Hash, 12)
		for i := range hashedLeaves {
			hashedLeaves[i] = simpleHash
		}
		committer := NewCommitter()
		computedRoot, err := committer.ComputeRoot(hashedLeaves, 14)
		require.NoError(t, err)

		leaves := make([]common.Hash, 14)
		for i := range leaves {
			leaves[i] = simpleHash
		}
		histCommit, err := NewLegacy(leaves)
		require.NoError(t, err)
		require.Equal(t, computedRoot, histCommit.Merkle)
	})
}

func TestMaximumDepthHistoryCommitment(t *testing.T) {
	t.Parallel()
	simpleHash := crypto.Keccak256Hash([]byte("foo"))
	hashedLeaves := []common.Hash{
		simpleHash,
	}
	committer := NewCommitter()
	_, err := committer.ComputeRoot(hashedLeaves, 1<<26)
	require.NoError(t, err)
}

func BenchmarkMaximumDepthHistoryCommitment(b *testing.B) {
	b.StopTimer()
	simpleHash := crypto.Keccak256Hash([]byte("foo"))
	hashedLeaves := []common.Hash{
		simpleHash,
	}
	committer := NewCommitter()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, err := committer.ComputeRoot(hashedLeaves, 1<<26)
		_ = err
	}
}

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
	oldCommit, err := NewLegacy(oldLeaves)
	require.NoError(t, err)
	require.Equal(t, commit.Merkle, oldCommit.Merkle)
}

// Utility function to generate random bytes for benchmarking
func randomBytes(size int) []byte {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		panic(err) // Handle error in production code
	}
	return b
}

// Benchmark for your custom hash function
func BenchmarkHashFunction(b *testing.B) {
	comm := NewCommitter()
	data := randomBytes(32) // Change size as necessary
	for i := 0; i < b.N; i++ {
		comm.hash(data)
	}
}

// Benchmark for the Keccak256Hash function
func BenchmarkKeccak256Hash(b *testing.B) {
	data := randomBytes(32) // Change size as necessary
	for i := 0; i < b.N; i++ {
		crypto.Keccak256Hash(data)
	}
}
