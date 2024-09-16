package optimized

import (
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// computeMerkleProof computes the Merkle proof for a leaf at a given index.
// It uses the last leaf in hLeaves to pad the tree up to the 'virtual' size if needed.
func computeMerkleProof(index int, virtual int, hLeaves []common.Hash) []common.Hash {
	N := len(hLeaves)                                // Number of real leaves
	hN := hLeaves[N-1]                               // Last leaf used for padding
	D := int(math.Ceil(math.Log2(float64(virtual)))) // Tree depth

	// Precompute hNHashes
	hNHashes := precomputeHNHashes(hN, D)

	var proof []common.Hash
	for level := 0; level < D; level++ {
		nodeIndex := index >> level
		siblingHash, exists := computeSiblingHash(nodeIndex, level, N, virtual, hLeaves, hNHashes)
		if exists {
			proof = append(proof, siblingHash)
		}
	}
	return proof
}

// precomputeHNHashes precomputes the hashes formed by combining hN with itself at each level.
func precomputeHNHashes(hN common.Hash, D int) []common.Hash {
	hNHashes := make([]common.Hash, D+1)
	hNHashes[0] = hN
	for level := 1; level <= D; level++ {
		data := append(hNHashes[level-1].Bytes(), hNHashes[level-1].Bytes()...)
		hNHashes[level] = crypto.Keccak256Hash(data)
	}
	return hNHashes
}

// computeSiblingHash computes the hash of a node's sibling at a given index and level.
func computeSiblingHash(nodeIndex int, level int, N int, virtual int, hLeaves []common.Hash, hNHashes []common.Hash) (common.Hash, bool) {
	siblingIndex := nodeIndex ^ 1
	numNodes := (virtual + (1 << level) - 1) / (1 << level) // Equivalent to ceil(virtual / (2 ** level))

	if siblingIndex >= numNodes {
		// No sibling exists; handle according to your tree's rules
		return common.Hash{}, false
	} else if siblingIndex >= paddingStartIndexAtLevel(N, level) {
		return hNHashes[level], true
	} else {
		siblingHash := computeNodeHash(siblingIndex, level, N, hLeaves, hNHashes)
		return siblingHash, true
	}
}

// computeNodeHash recursively computes the hash of a node at a given index and level.
func computeNodeHash(nodeIndex int, level int, N int, hLeaves []common.Hash, hNHashes []common.Hash) common.Hash {
	if level == 0 {
		if nodeIndex >= N {
			// Node is in padding
			return hNHashes[0]
		} else {
			return hLeaves[nodeIndex]
		}
	} else {
		if nodeIndex >= paddingStartIndexAtLevel(N, level) {
			return hNHashes[level]
		} else {
			leftChild := computeNodeHash(2*nodeIndex, level-1, N, hLeaves, hNHashes)
			rightChild := computeNodeHash(2*nodeIndex+1, level-1, N, hLeaves, hNHashes)
			data := append(leftChild.Bytes(), rightChild.Bytes()...)
			return crypto.Keccak256Hash(data)
		}
	}
}

// paddingStartIndexAtLevel calculates the index at which padding starts at a given tree level.
func paddingStartIndexAtLevel(N int, level int) int {
	return N / (1 << level)
}
