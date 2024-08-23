package optimized

import (
	"errors"
	"math"

	"github.com/OffchainLabs/bold/containers/option"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func ComputeHistoryCommitment(
	realLeaves []common.Hash,
	fromIndex int,
	toIndex option.Option[int],
	maxSize uint64,
) (common.Hash, error) {
	if len(realLeaves) == 0 {
		return common.Hash{}, errors.New("no leaves provided")
	}
	var emptyHash common.Hash
	from := fromIndex
	var to int
	if toIndex.IsSome() {
		to = toIndex.Unwrap()
		if to <= fromIndex {
			return emptyHash, errors.New("invalid range: end was <= start")
		}
		if uint64(to) >= maxSize {
			return emptyHash, errors.New("invalid range: end was >= max size")
		}
	} else {
		to = len(realLeaves) - 1
	}
	if to < len(realLeaves) {
		// Case 0: the range is entirely within the real leaves' length.
		// we can simply compute a history commitment for it by slicing the list
		// and computing the root of the Merkle tree formed the slice of these leaves.
		// If this slice has a non-power of two length, we use virtual zero hashes
		// to build a sparse Merkle tree and compute its root.
		slicedLeaves := realLeaves[from:to]
		return computeVirtualSparseTree(slicedLeaves, len(slicedLeaves), 1<<26)
	} else if from < len(realLeaves) && to >= len(realLeaves) {
		// Case 1: the `from` index is within the range of the length of the real
		// hashes list, but the `to` index exceeds it.
		// Here, we need to commit to commit to a Merkle tree formed by the following concatenation:
		// realLeaves[from:] ++ (realLeaves[-1] * (to - len(realLeaves)))
		// This means the leaves we are committing to are the real leaves up to the end of the list,
		// and the last leaf padded to the `to` index.
		// If the number of leaves we are committing to is not a power of two, we use virtual zero hashes
		// to compute a sparse Merkle tree and root.
		slicedLeaves := realLeaves[from:]
		return computeVirtualSparseTree(slicedLeaves, to-from, 1<<26)
	} else {
		// Case 2: Both the `from` and `to` indices are out of range of the real hashes list.
		// In this case, we commit to a Merkle tree formed by realLeaves[-1] * (to - from). That is,
		// we commit to a Merkle tree formed by the last leaf of the real leaves list repeated until
		// the specified range.
		leaves := []common.Hash{realLeaves[len(realLeaves)-1]}
		return computeVirtualSparseTree(leaves, to-from, 1<<26)
	}
}

var (
	keccak          = crypto.NewKeccakState()
	lastLeafFillers []common.Hash
)

// precomputeRepeatedHashes returns a slice where built recursively as
// ret[0] = the passed in leaf
// ret[i+1] = Hash(ret[i] + ret[i])
// Allocates n hashes
// Computes n-1 hashes
// Copies 1 hash
func precomputeRepeatedHashes(leaf *common.Hash, n int) []common.Hash {
	if n < 0 {
		return nil
	}
	ret := make([]common.Hash, n)
	copy(ret[0][:], (*leaf)[:])
	for i := 1; i < n; i++ {
		keccak.Write(ret[i-1][:])
		keccak.Write(ret[i-1][:])
		keccak.Read(ret[i][:])
		keccak.Reset()
	}
	return ret
}

// Warning: using ints, don't care about 32 bits systems.
func nextPowerOf2(n int) int {
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

// computeSparseTree returns the htr of a hashtree with the given leaves and
// limit. Any non-allocated leaf is filled with the passed zeroHash of depth 0.
// Recursively, any non allocated intermediate layer at depth i is filled with
// the passed zeroHash of the corresponding depth.
// limit is assumed to be a power of two which is higher or equal than the
// length of the leaves.
// fillers is assumed to be precomputed to the necessary limit, no error
// handling
//
// Zero allocations
// Computes O(len(leaves)) hashes.
func computeSparseTree(leaves []common.Hash, limit int, fillers []common.Hash) common.Hash {
	if limit < 2 {
		return leaves[0]
	}
	m := len(leaves)
	depth := int(math.Log2(float64(limit)))
	for j := 0; j < depth; j++ {
		for i := 0; i < m/2; i++ {
			keccak.Write(leaves[2*i][:])
			keccak.Write(leaves[2*i+1][:])
			keccak.Read(leaves[i][:])
			keccak.Reset()
		}
		if m&1 == 1 {
			keccak.Write(leaves[m-1][:])
			keccak.Write(fillers[j][:])
			keccak.Read(leaves[(m-1)/2][:])
			keccak.Reset()
		}
		m = (m + 1) / 2
	}
	return leaves[0]
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
func computeVirtualSparseTree(leaves []common.Hash, virtual int, limit int) (common.Hash, error) {
	m := len(leaves)
	if m == 0 {
		return common.Hash{}, errors.New("nil leaves")
	}
	if virtual < m {
		return common.Hash{}, errors.New("Don't be nasty")
	}
	if limit == 0 {
		// this is used in the initial case, to signal that the limit
		// needs to be set to be the smallest power of two larger than
		// the virtual size. We also precompute the higher powers of the
		// last leaf, this computes O(virtual) hashes and allocates
		// O(log(virtual)) hashes.
		limit = nextPowerOf2(virtual)
		lastLeafFillers = precomputeRepeatedHashes(&leaves[m-1], int(math.Log2(float64(virtual-m))+1))
	}
	if limit == 1 {
		return leaves[0], nil
	}
	if limit < virtual {
		return common.Hash{}, errors.New("Don't be nasty")
	}
	var left, right common.Hash
	var err error
	if virtual >= limit/2 {
		if m > limit/2 {
			// Leaves are enough to cover the first half of the
			// tree, The first half is then a normal full hashtree
			// and the right side is computed by recursion.
			// It is safe to pass anything here as the fillers since
			// the tree is full
			left = computeSparseTree(leaves[:limit/2], limit/2, nil)
			right, err = computeVirtualSparseTree(leaves[limit/2:], virtual-limit/2, limit/2)
			if err != nil {
				return common.Hash{}, err
			}
		} else {
			// Leaves and virtual fit in the first half of the tree,
			// In this case we need to compute the HTR of the
			// compute then the full first half of the tree as if it
			// were a normal sparse tree but with the virtual
			// fillers
			left = computeSparseTree(leaves, limit/2, lastLeafFillers)
			if virtual == limit {
				right = lastLeafFillers[int(math.Log2(float64(limit/2)))]
			} else if virtual == limit/2 {
				right = zeroHashes[0]
			} else {
				right, err = computeVirtualSparseTree([]common.Hash{lastLeafFillers[0]}, virtual-limit/2, limit/2)
				if err != nil {
					return common.Hash{}, nil
				}
			}
		}
	} else {
		// In this case both leaves and virtual size are in the first
		// half of the tree, so the second half is just a higher zero
		// hash.
		left, err = computeVirtualSparseTree(leaves, virtual, limit/2)
		if err != nil {
			return common.Hash{}, err
		}
		right = zeroHashes[0]
	}
	keccak.Write(left[:])
	keccak.Write(right[:])
	keccak.Read(leaves[0][:])
	keccak.Reset()
	return leaves[0], nil
}
