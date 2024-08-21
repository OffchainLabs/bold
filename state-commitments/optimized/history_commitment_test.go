package optimized

import (
	"errors"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestVirtualSparse(t *testing.T) {
	simpleHash := crypto.Keccak256Hash([]byte("foo"))
	leaves := []common.Hash{
		simpleHash,
		simpleHash,
		simpleHash,
		{},
	}
	root, err := computeSparseRoot(leaves, 2)
	require.NoError(t, err)
	fmt.Println(root.Hex())

	// Manual root:
	firstHalf := crypto.Keccak256Hash(simpleHash[:], simpleHash[:])
	secondHalf := crypto.Keccak256Hash(simpleHash[:], (common.Hash{}).Bytes())
	manualRoot := crypto.Keccak256Hash(firstHalf[:], secondHalf[:])
	fmt.Println(manualRoot.Hex())

	// Virtual + sparse.
	fullRoot, err := computeVirtualSparseTree([]common.Hash{simpleHash}, 3)
	require.NoError(t, err)
	fmt.Println(fullRoot.Hex())

	// fullRoot, err = computeVirtualSparseTree([]common.Hash{simpleHash}, (1<<26)-1)
	// require.NoError(t, err)
	// fmt.Println(fullRoot.Hex())
	t.Fatal(1)
}

func TestMaximumDepthHistoryCommitment(t *testing.T) {
	simpleHash := crypto.Keccak256Hash([]byte("foo"))
	_, err := computeVirtualSparseTree([]common.Hash{simpleHash}, (1<<26)-1)
	require.NoError(t, err)
}

func BenchmarkMaximumDepthHistoryCommitment(b *testing.B) {
	b.StopTimer()
	simpleHash := crypto.Keccak256Hash([]byte("foo"))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		computeVirtualSparseTree([]common.Hash{simpleHash}, (1<<26)-1)
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
