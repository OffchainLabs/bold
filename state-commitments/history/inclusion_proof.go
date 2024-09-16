package history

import (
	"errors"
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Computes the Merkle proof for a leaf at a given index.
// It uses the last leaf to pad the tree up to the 'virtual' size if needed.
func (h *HistoryCommitter) computeMerkleProof(leafIndex uint64, leaves []common.Hash, virtual uint64) ([]common.Hash, error) {
	if len(leaves) == 0 {
		return nil, nil
	}
	// TODO: Add all other safety conditions.
	if virtual == 0 {
		return nil, errors.New("virtual size must be greater than 0")
	}
	numRealLeaves := uint64(len(leaves))
	// Last leaf used for padding.
	lastLeaf := leaves[numRealLeaves-1]
	depth := int(math.Ceil(math.Log2(float64(virtual))))

	// Precompute virtual hashes
	virtualHashes, err := h.precomputeRepeatedHashes(&lastLeaf, depth)
	if err != nil {
		return nil, err
	}

	var proof []common.Hash
	for level := 0; level < depth; level++ {
		nodeIndex := leafIndex >> level
		siblingHash, exists := computeSiblingHash(nodeIndex, uint64(level), numRealLeaves, virtual, leaves, virtualHashes)
		if exists {
			proof = append(proof, siblingHash)
		}
	}
	return proof, nil
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
func computeNodeHash(
	nodeIndex uint64, level uint64, numRealLeaves uint64, leaves []common.Hash, virtualHashes []common.Hash,
) common.Hash {
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
