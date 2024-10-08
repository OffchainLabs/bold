package history

import (
	"errors"

	"github.com/OffchainLabs/bold/math"
	"github.com/ethereum/go-ethereum/common"
)

// Computes the Merkle proof for a leaf at a given index.
// It uses the last leaf to pad the tree up to the 'virtual' size if needed.
func (h *historyCommitter) computeMerkleProof(leafIndex uint64, leaves []common.Hash, virtual uint64) ([]common.Hash, error) {
	if len(leaves) == 0 {
		return nil, nil
	}
	if leafIndex >= virtual {
		return nil, errors.New("leaf index out of bounds")
	}
	if virtual < uint64(len(leaves)) {
		return nil, errors.New("virtual size must be greater than or equal to the number of leaves")
	}
	numRealLeaves := uint64(len(leaves))
	lastLeaf := h.hash(&leaves[numRealLeaves-1])
	depth := uint(math.Log2Ceil(virtual))

	// Precompute virtual hashes
	virtualHashes, err := h.precomputeRepeatedHashes(&lastLeaf, depth)
	if err != nil {
		return nil, err
	}
	var proof []common.Hash
	for level := uint(0); level < depth; level++ {
		nodeIndex := leafIndex >> level
		siblingHash, exists, err := h.computeSiblingHash(nodeIndex, uint64(level), numRealLeaves, virtual, leaves, virtualHashes)
		if err != nil {
			return nil, err
		}
		if exists {
			proof = append(proof, siblingHash)
		}
	}
	return proof, nil
}

// Computes the hash of a node's sibling at a given index and level.
func (h *historyCommitter) computeSiblingHash(
	nodeIndex uint64,
	level uint64,
	N uint64,
	virtual uint64,
	hLeaves []common.Hash,
	hNHashes []common.Hash,
) (common.Hash, bool, error) {
	siblingIndex := nodeIndex ^ 1
	// Essentially ceil(virtual / (2 ** level))
	numNodes := (virtual + (1 << level) - 1) / (1 << level)
	if siblingIndex >= numNodes {
		// No sibling exists, so use a zero hash.
		return emptyHash, true, nil
	} else if siblingIndex >= paddingStartIndexAtLevel(N, level) && siblingIndex <= paddingEndIndexAtLevel(virtual, level) {
		return hNHashes[level], true, nil
	} else {
		siblingHash, err := h.computeNodeHash(siblingIndex, level, N, hLeaves, hNHashes, virtual)
		if err != nil {
			return emptyHash, false, err
		}
		return siblingHash, true, nil
	}
}

// Recursively computes the hash of a node at a given index and level.
func (h *historyCommitter) computeNodeHash(
	nodeIndex uint64, level uint64, numRealLeaves uint64, leaves []common.Hash, virtualHashes []common.Hash, virtual uint64,
) (common.Hash, error) {
	if level == 0 {
		if nodeIndex >= virtual && virtual > numRealLeaves {
			// Beyond the virtual or real size of the tree, so use a zero hash.
			return emptyHash, nil
		}
		if nodeIndex >= numRealLeaves {
			// Node is in padding (the virtual segment of the tree).
			return virtualHashes[0], nil
		} else {
			return h.hash(&leaves[nodeIndex]), nil
		}
	} else {
		if nodeIndex >= paddingStartIndexAtLevel(numRealLeaves, level) && nodeIndex <= paddingEndIndexAtLevel(virtual, level) {
			return virtualHashes[level], nil
		} else {
			leftChild, err := h.computeNodeHash(2*nodeIndex, level-1, numRealLeaves, leaves, virtualHashes, virtual)
			if err != nil {
				return emptyHash, err
			}
			rightChild, err := h.computeNodeHash(2*nodeIndex+1, level-1, numRealLeaves, leaves, virtualHashes, virtual)
			if err != nil {
				return emptyHash, err
			}
			return h.hash(&leftChild, &rightChild), nil
		}
	}
}

// Calculates the highest index in a level which is purely padding.
func paddingEndIndexAtLevel(virtual uint64, level uint64) uint64 {
	return virtual/(1<<level) - 1
}

// Calculates the index at which padding starts at a given tree level.
func paddingStartIndexAtLevel(N uint64, level uint64) uint64 {
	return N / (1 << level)
}
