package optimized

import (
	"errors"
	"fmt"
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type HistoryCommitter struct {
	lastLeafFillers []common.Hash
	keccak          crypto.KeccakState
	limit           uint64
	virtual         uint64
}

type CommitmentBuilder struct {
	limit   *uint64
	virtual *uint64
}

func NewBuilder() *CommitmentBuilder {
	return &CommitmentBuilder{}
}

func (cb *CommitmentBuilder) Limit(n uint64) *CommitmentBuilder {
	cb.limit = &n
	return cb
}

func (cb *CommitmentBuilder) Virtual(n uint64) *CommitmentBuilder {
	cb.virtual = &n
	return cb
}

func (cb *CommitmentBuilder) Build() (*HistoryCommitter, error) {
	if cb.limit == nil {
		return nil, errors.New("limit not set")
	}
	if cb.virtual == nil {
		return nil, errors.New("virtual not set")
	}
	return &HistoryCommitter{
		limit:   *cb.limit,
		virtual: *cb.virtual,
		keccak:  crypto.NewKeccakState(),
	}, nil
}

func (h *HistoryCommitter) ComputeRoot(leaves []common.Hash) (common.Hash, error) {
	// Called with 0 limit first to compute the last leaf fillers for the commitment.
	copiedLeaves := make([]common.Hash, len(leaves))
	for i, leaf := range leaves {
		var copied common.Hash
		copy(copied[:], leaf[:])
		copiedLeaves[i] = copied
	}
	_, err := h.computeVirtualSparseTree(copiedLeaves, h.virtual, 0)
	if err != nil {
		return common.Hash{}, err
	}
	return h.computeVirtualSparseTree(leaves, h.virtual, h.limit)
}

func (h *HistoryCommitter) GeneratePrefixProof(prefixIndex uint64, leaves []common.Hash) ([]common.Hash, []common.Hash, error) {
	prefixExpansion, proof, err := h.prefixAndProof(prefixIndex, leaves, h.virtual)
	if err != nil {
		return nil, nil, err
	}
	prefixExpansion = trimTrailingZeroHashes(prefixExpansion)
	proof = trimZeroes(proof)
	return prefixExpansion, proof, nil
}

// computeSparseTree returns the htr of a hashtree with the given leaves and
// limit. Any non-allocated leaf is filled with the passed zeroHash of depth 0.
// Recursively, any non allocated intermediate layer at depth i is filled with
// the passed zeroHash.
// limit is assumed to be a power of two which is higher or equal than the
// length of the leaves.
// fillers is assumed to be precomputed to the necessary limit, no error
// handling
//
// Zero allocations
// Computes O(len(leaves)) hashes.
func (h *HistoryCommitter) computeSparseTree(leaves []common.Hash, limit uint64, fillers []common.Hash) (common.Hash, error) {
	if limit < 2 {
		return leaves[0], nil
	}
	m := len(leaves)
	depth := int(math.Log2(float64(limit)))
	for j := 0; j < depth; j++ {
		for i := 0; i < m/2; i++ {
			if _, err := h.keccak.Write(leaves[2*i][:]); err != nil {
				return common.Hash{}, err
			}
			if _, err := h.keccak.Write(leaves[2*i+1][:]); err != nil {
				return common.Hash{}, err
			}
			if _, err := h.keccak.Read(leaves[i][:]); err != nil {
				return common.Hash{}, err
			}
			h.keccak.Reset()
		}
		if m&1 == 1 {
			if _, err := h.keccak.Write(leaves[m-1][:]); err != nil {
				return common.Hash{}, err
			}
			if _, err := h.keccak.Write(fillers[j][:]); err != nil {
				return common.Hash{}, err
			}
			if _, err := h.keccak.Read(leaves[(m-1)/2][:]); err != nil {
				return common.Hash{}, err
			}
			h.keccak.Reset()
		}
		m = (m + 1) / 2
	}
	return leaves[0], nil
}

// computeVirtualSparseTree returns the htr of a hashtree where the first layer
// is passed as leaves, the completed with the last leaf until it reaches
// virtual and finally completed with zero hashes until it reaches limit.
// limit is assumed to be either 0 or a power of 2 which is greater or equal to
// virtual. If limit is zero it behaves as if it were the smallest power of two
// that is greater or equal than virtual.
//
// The algorithm is split in three different logic parts:
//
//  1. If the virtual length is less than half the limit (this can never happen
//     in the first iteration of the algorithm), then the first half of the tree
//     is computed by recursion and the second half is a zero hash of a given
//     depth.
//  2. If the leaves all fit in the first half, then we can optimize the first
//     half to being a simple sparse tree, just that instead of filling with zero
//     hashes we fill with the precomputed virtual hashes. This is the most common
//     starting scenario. The second part is computed by recursion.
//  3. If the leaves do not fit in the first half, then we can compute the first half of
//     the tree as a normal full hashtree. The second part is computed by recursion.
func (h *HistoryCommitter) computeVirtualSparseTree(leaves []common.Hash, virtual, limit uint64) (common.Hash, error) {
	m := uint64(len(leaves))
	if m == 0 {
		return common.Hash{}, errors.New("nil leaves")
	}
	if virtual < m {
		return common.Hash{}, fmt.Errorf("virtual %d should be >= num leaves %d", virtual, m)
	}
	var err error
	if limit == 0 {
		// this is used in the initial case, to signal that the limit
		// needs to be set to be the smallest power of two larger than
		// the virtual size. We also precompute the higher powers of the
		// last leaf, this computes O(virtual) hashes and allocates
		// O(log(virtual)) hashes.
		limit = nextPowerOf2(virtual)
		h.lastLeafFillers, err = h.precomputeRepeatedHashes(&leaves[m-1], int(math.Log2(float64(virtual-m))+1))
		if err != nil {
			return common.Hash{}, err
		}
	}
	if limit == 1 {
		return leaves[0], nil
	}
	if limit < virtual {
		return common.Hash{}, fmt.Errorf("limit %d should be >= virtual %d", limit, virtual)
	}
	var left, right common.Hash
	if virtual > limit/2 {
		if m > limit/2 {
			// Leaves are enough to cover the first half of the
			// tree, The first half is then a normal full hashtree
			// and the right side is computed by recursion.
			// It is safe to pass anything here as the fillers since
			// the tree is full
			left, err = h.computeSparseTree(leaves[:limit/2], limit/2, nil)
			if err != nil {
				return common.Hash{}, err
			}
			right, err = h.computeVirtualSparseTree(leaves[limit/2:], virtual-limit/2, limit/2)
			if err != nil {
				return common.Hash{}, err
			}
		} else {
			// Leaves and virtual fit in the first half of the tree,
			// In this case we need to compute the HTR of the
			// compute then the full first half of the tree as if it
			// were a normal sparse tree but with the virtual
			// fillers
			left, err = h.computeSparseTree(leaves, limit/2, h.lastLeafFillers)
			if err != nil {
				return common.Hash{}, err
			}
			if virtual == limit {
				right = h.lastLeafFillers[int(math.Log2(float64(limit/2)))]
			} else {
				right, err = h.computeVirtualSparseTree([]common.Hash{h.lastLeafFillers[0]}, virtual-limit/2, limit/2)
				if err != nil {
					return common.Hash{}, err
				}
			}
		}
	} else {
		// In this case both leaves and virtual size are in the first
		// half of the tree, so the second half is just a higher zero
		// hash.
		left, err = h.computeVirtualSparseTree(leaves, virtual, limit/2)
		if err != nil {
			return common.Hash{}, err
		}
		right = zeroHashes[0]
	}
	if _, err = h.keccak.Write(left[:]); err != nil {
		return common.Hash{}, err
	}
	if _, err = h.keccak.Write(right[:]); err != nil {
		return common.Hash{}, err
	}
	if _, err = h.keccak.Read(leaves[0][:]); err != nil {
		return common.Hash{}, err
	}
	h.keccak.Reset()
	return leaves[0], nil
}

func (h *HistoryCommitter) subtreeExpansion(leaves []common.Hash, virtual, limit uint64, stripped bool) (proof []common.Hash, err error) {
	m := uint64(len(leaves))
	if virtual == 0 {
		for i := limit; i > 1; i /= 2 {
			proof = append(proof, zeroHashes[0])
		}
		return
	}
	if limit == 0 {
		limit = nextPowerOf2(virtual)
	}
	if limit == virtual {
		left, err2 := h.computeSparseTree(leaves, limit, h.lastLeafFillers)
		if err2 != nil {
			return nil, err2
		}
		if !stripped {
			for i := limit; i > 1; i /= 2 {
				proof = append(proof, zeroHashes[0])
			}
		}
		return append(proof, left), nil
	}
	if m > limit/2 {
		left, err2 := h.computeSparseTree(leaves[:limit/2], limit/2, nil)
		if err2 != nil {
			return nil, err2
		}
		proof, err = h.subtreeExpansion(leaves[limit/2:], virtual-limit/2, limit/2, stripped)
		if err != nil {
			return nil, err
		}
		return append(proof, left), nil
	}
	if virtual >= limit/2 {
		left, err2 := h.computeSparseTree(leaves, limit/2, h.lastLeafFillers)
		if err2 != nil {
			return nil, err2
		}
		proof, err = h.subtreeExpansion([]common.Hash{h.lastLeafFillers[0]}, virtual-limit/2, limit/2, stripped)
		if err != nil {
			return nil, err
		}
		return append(proof, left), nil
	}
	if stripped {
		return h.subtreeExpansion(leaves, virtual, limit/2, stripped)
	}
	expac, err := h.subtreeExpansion(leaves, virtual, limit/2, stripped)
	if err != nil {
		return nil, err
	}
	return append(expac, zeroHashes[0]), nil
}

func (h *HistoryCommitter) proof(index uint64, leaves []common.Hash, virtual, limit uint64) (tail []common.Hash, err error) {
	m := uint64(len(leaves))
	if limit == 0 {
		limit = nextPowerOf2(virtual)
	}
	if limit == 1 {
		// Can only reach this with index == 0
		return
	}
	if index >= limit/2 {
		if m > limit/2 {
			return h.proof(index-limit/2, leaves[limit/2:], virtual-limit/2, limit/2)
		}
		return h.proof(index-limit/2, []common.Hash{h.lastLeafFillers[0]}, virtual-limit/2, limit/2)
	}
	if m > limit/2 {
		tail, err = h.proof(index, leaves[:limit/2], limit/2, limit/2)
		if err != nil {
			return nil, err
		}
		right, err2 := h.subtreeExpansion(leaves[limit/2:], virtual-limit/2, limit/2, true)
		if err2 != nil {
			return nil, err2
		}
		for i := len(right) - 1; i >= 0; i-- {
			tail = append(tail, right[i])
		}
		return tail, nil
	}
	if virtual > limit/2 {
		tail, err = h.proof(index, leaves, limit/2, limit/2)
		if err != nil {
			return nil, err
		}
		right, err := h.subtreeExpansion([]common.Hash{h.lastLeafFillers[0]}, virtual-limit/2, limit/2, true)
		if err != nil {
			return nil, err
		}
		for i := len(right) - 1; i >= 0; i-- {
			tail = append(tail, right[i])
		}
		return tail, nil
	}
	return h.proof(index, leaves, virtual, limit/2)
}

func (h *HistoryCommitter) prefixAndProof(index uint64, leaves []common.Hash, virtual uint64) (prefix []common.Hash, tail []common.Hash, err error) {
	m := uint64(len(leaves))
	if m == 0 {
		return nil, nil, errors.New("nil leaves")
	}
	if m > virtual {
		return nil, nil, fmt.Errorf("num leaves %d should be <= virtual %d", m, virtual)
	}
	if index+1 > virtual {
		return nil, nil, fmt.Errorf("index %d + 1 should be <= virtual %d", index, virtual)
	}
	h.lastLeafFillers, err = h.precomputeRepeatedHashes(&leaves[m-1], int(math.Log2(float64(virtual))+1))
	if err != nil {
		return nil, nil, err
	}
	if index+1 > m {
		prefix, err = h.subtreeExpansion(leaves, index+1, 0, false)
	} else {
		prefix, err = h.subtreeExpansion(leaves[:index+1], index+1, 0, false)
	}
	if err != nil {
		return nil, nil, err
	}
	tail, err = h.proof(index, leaves, virtual, 0)
	return
}

// precomputeRepeatedHashes returns a slice where built recursively as
// ret[0] = the passed in leaf
// ret[i+1] = Hash(ret[i] + ret[i])
// Allocates n hashes
// Computes n-1 hashes
// Copies 1 hash
func (h *HistoryCommitter) precomputeRepeatedHashes(leaf *common.Hash, n int) ([]common.Hash, error) {
	if n < 0 {
		return nil, nil
	}
	ret := make([]common.Hash, n)
	copy(ret[0][:], (*leaf)[:])
	for i := 1; i < n; i++ {
		if _, err := h.keccak.Write(ret[i-1][:]); err != nil {
			return nil, err
		}
		if _, err := h.keccak.Write(ret[i-1][:]); err != nil {
			return nil, err
		}
		if _, err := h.keccak.Read(ret[i][:]); err != nil {
			return nil, err
		}
		h.keccak.Reset()
	}
	return ret, nil
}

func nextPowerOf2(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	n--         // Decrement n to handle the case where n is already a power of 2
	n |= n >> 1 // Propagate the highest bit set
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n |= n >> 32
	return n + 1 // Increment n to get the next power of 2
}

func trimTrailingZeroHashes(hashes []common.Hash) []common.Hash {
	// Start from the end of the slice
	for i := len(hashes) - 1; i >= 0; i-- {
		// If we find a non-zero hash, return the slice up to and including this element
		if hashes[i] != (common.Hash{}) {
			return hashes[:i+1]
		}
	}
	// If all elements are zero, return an empty slice
	return []common.Hash{}
}

func trimZeroes(hashes []common.Hash) []common.Hash {
	newHashes := make([]common.Hash, 0, len(hashes))
	for _, h := range hashes {
		if h == (common.Hash{}) {
			continue
		}
		newHashes = append(newHashes, h)
	}
	return newHashes
}
