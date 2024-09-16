package optimized

import (
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Computes the Merkle proof for a leaf at a given index.
// It uses the last leaf to pad the tree up to the 'virtual' size if needed.
func computeMerkleProof(leafIndex uint64, leaves []common.Hash, virtual uint64) []common.Hash {
	if len(leaves) == 0 {
		return nil
	}
	// TODO: Add all other safety conditions.
	if virtual == 0 {
		return nil
	}
	numRealLeaves := uint64(len(leaves))
	// Last leaf used for padding.
	lastLeaf := leaves[numRealLeaves-1]
	depth := uint64(math.Ceil(math.Log2(float64(virtual))))

	// Precompute virtual hashes
	virtualHashes := precomputeVirtualHashes(lastLeaf, depth)

	var proof []common.Hash
	for level := uint64(0); level < depth; level++ {
		nodeIndex := leafIndex >> level
		siblingHash, exists := computeSiblingHash(nodeIndex, level, numRealLeaves, virtual, leaves, virtualHashes)
		if exists {
			proof = append(proof, siblingHash)
		}
	}
	return proof
}

// Precomputes hashes formed by combining an element with itself at each level.
func precomputeVirtualHashes(item common.Hash, depth uint64) []common.Hash {
	hNHashes := make([]common.Hash, depth+1)
	hNHashes[0] = item
	for level := uint64(1); level <= depth; level++ {
		data := append(hNHashes[level-1].Bytes(), hNHashes[level-1].Bytes()...)
		hNHashes[level] = crypto.Keccak256Hash(data)
	}
	return hNHashes
}

// Computes the hash of a node's sibling at a given index and level.
func computeSiblingHash(
	nodeIndex uint64,
	level uint64,
	N uint64,
	virtual uint64,
	hLeaves []common.Hash,
	hNHashes []common.Hash,
) (common.Hash, bool) {
	siblingIndex := nodeIndex ^ 1
	numNodes := (virtual + (1 << level) - 1) / (1 << level) // Equivalent to ceil(virtual / (2 ** level))
	if siblingIndex >= numNodes {
		// No sibling exists, so use a zero hash.
		return common.Hash{}, false
	} else if siblingIndex >= paddingStartIndexAtLevel(N, level) {
		return hNHashes[level], true
	} else {
		siblingHash := computeNodeHash(siblingIndex, level, N, hLeaves, hNHashes)
		return siblingHash, true
	}
}

// Recursively computes the hash of a node at a given index and level.
func computeNodeHash(nodeIndex uint64, level uint64, numRealLeaves uint64, leaves []common.Hash, virtualHashes []common.Hash) common.Hash {
	if level == 0 {
		if nodeIndex >= numRealLeaves {
			// Node is in padding (the virtual segment of the tree).
			return virtualHashes[0]
		} else {
			return leaves[nodeIndex]
		}
	} else {
		if nodeIndex >= paddingStartIndexAtLevel(numRealLeaves, level) {
			return virtualHashes[level]
		} else {
			leftChild := computeNodeHash(2*nodeIndex, level-1, numRealLeaves, leaves, virtualHashes)
			rightChild := computeNodeHash(2*nodeIndex+1, level-1, numRealLeaves, leaves, virtualHashes)
			data := append(leftChild.Bytes(), rightChild.Bytes()...)
			return crypto.Keccak256Hash(data)
		}
	}
}

// Calculates the index at which padding starts at a given tree level.
func paddingStartIndexAtLevel(N uint64, level uint64) uint64 {
	return N / (1 << level)
}
