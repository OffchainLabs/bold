// The history package provides functions for computing merkele tree roots
// and proofs needed for the BoLD protocol's history commitments.
//
// Throughout this package, the following terms are used:
//
//   - leaf: a leaf node in a merkle tree, which is a hash of some data.
//   - virtual: the length of the desired number of leaf nodes. In the BoLD
//     protocol, it is important that all history commitments which for a given
//     challenge edge have the same length, even if the participants disagree
//     about the number of blocks or steps to which they are committing. To
//     solve this, history commitments must have fixed lengths at different
//     challenge levels. Callers only need to provide the leaves they to which
//     they commit, and the virtual length. The last leaf in the list is used
//     to pad the tree to the virtual length.
//   - limit: the length of the leaves that would be in a complete subtree
//     of the depth required to hold the virtual leaves in a tree (or subtree)
//   - pure tree: a tree where len(leaves) == virtual
//   - complete tree: a tree where the number of leaves is a power of 2
//   - complete virtual tree: a tree where the number of leaves including the
//     virtual padding is a power of 2
//   - partial tree: a tree where the number of leaves is not a power of 2
//   - partial virtual tree: a tree where the number of leaves including the
//     virtual padding is not a power of 2
//   - empty hash: common.Hash{}
//     Any time the root of a partial tree (either virtual or pure) is computed,
//     the sibling node of the last node in a layer may be missing. In this case
//     an empty hash (common.Hash{}) is used as the sibling node.
//     Note: This is not the same as padding the leaves of the tree with
//     common.Hash{} values. If that approach were taken, then the higher-level
//     layers would contain the hash of the empty hash, or the hash of multiple
//     empty hashes. This would be less efficient to calculate, and would not
//     change expressiveness or security of the data structure, but it would
//     produce a different root hash.
//   - virtual node: a node in a virtual tree which is not one of the real
//     leaves and not computed from the data in the real leaves.
package history

import (
	"errors"
	"fmt"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/OffchainLabs/bold/math"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	emptyHash    = common.Hash{}
	emptyHistory = protocol.History{}
)

type nonZero uint64

func newNonZero(n uint64) (nonZero, error) {
	if n == 0 {
		return 0, errors.New("zero is not a valid non-zero value")
	}
	return nonZero(n), nil
}

// treePosition tracks the current position in the merkele tree.
type treePosition struct {
	// layer is the layer of the tree.
	layer uint64
	// index is the index of the leaf in this layer of the tree.
	index uint64
}

func (c *treePosition) copy() treePosition {
	return treePosition{layer: c.layer, index: c.index}
}

type historyCommitter struct {
	fillers        []common.Hash
	keccak         crypto.KeccakState
	cursor         treePosition
	lastLeafProver *lastLeafProver
}

func NewCommitter() *historyCommitter {
	return &historyCommitter{
		fillers: make([]common.Hash, 0),
		keccak:  crypto.NewKeccakState(),
	}
}

// soughtHash holds a pointer to the hash and whether it has been found.
//
// Without this type, it would be impossible to distinguish between a hash which
// has not been found and a hash which is the value of common.Hash{}.
// That's because the lastLeafProver's postions map is initialized with pointers
// to common.Hash{} values in a pre-allocated slice.
type soughtHash struct {
	found bool
	hash  *common.Hash
}

// lastLeafProver finds the siblings needed to produce a merkle inclusion
// proof for the last leaf in a virtual merkle tree.
//
// The prover maintains a map of treePositions where sibling nodes live
// and fills them in as the historyCommitter calculates them.
type lastLeafProver struct {
	positions map[treePosition]*soughtHash
	proof     []common.Hash
}

func newLastLeafProver(virtual nonZero) *lastLeafProver {
	positions := lastLeafProofPositions(virtual)
	posMap := make(map[treePosition]*soughtHash, len(positions))
	proof := make([]common.Hash, len(positions))
	for i, pos := range positions {
		posMap[pos] = &soughtHash{false, &proof[i]}
	}
	return &lastLeafProver{
		positions: posMap,
		proof:     proof,
	}
}

// handle filters the hashes found while computing the merkle root looking for
// the sibling nodes needed to produce the merkle inclusion proof, and fills
// them in the proof slice.
func (p *lastLeafProver) handle(hash common.Hash, pos treePosition) {
	if sibling, ok := p.positions[pos]; ok {
		sibling.found = true
		*sibling.hash = hash
	}
}

// handle is called each time a hash is computed in the merkle tree.
//
// The cursor is kept in sync with tree traversal. The implementation of
// handle can therefore assume that the currsor is pointing to the node which
// has the value of the hash.
func (h *historyCommitter) handle(hash common.Hash) {
	if h.lastLeafProver != nil {
		h.lastLeafProver.handle(hash, h.cursor)
	}
}

// hash hashes the passed item into a common.Hash.
func (h *historyCommitter) hash(item ...*common.Hash) common.Hash {
	var result common.Hash
	h.hashInto(&result, item...)
	return result
}

// proof returns the merkle inclusion proof for the last leaf in a virtual tree.
//
// If the proof is not complete (i.e. some sibling nodes are missing), the
// sibling nodes are filled in with the fillers.
//
// The reason this works, is that the only nodes which are not visited when
// computing the merkle root are those which are in some complete virtual
// subtree.
func (h *historyCommitter) lastLeafProof() []common.Hash {
	for pos, sibling := range h.lastLeafProver.positions {
		if !sibling.found {
			*h.lastLeafProver.positions[pos].hash = h.fillers[pos.layer]
		}
	}
	if len(h.lastLeafProver.proof) == 0 {
		return nil
	}
	return h.lastLeafProver.proof
}

// hashInto hashes the concatenation of the passed items into the result.
// nolint:errcheck
func (h *historyCommitter) hashInto(result *common.Hash, items ...*common.Hash) {
	defer h.keccak.Reset()
	for _, item := range items {
		h.keccak.Write(item[:]) // #nosec G104 - KeccakState.Write never errors
	}
	h.keccak.Read(result[:]) // #nosec G104 - KeccakState.Read never errors
}

// NewCommitment produces a history commitment from a list of real leaves that
// are virtually padded using the last leaf in the list to some virtual length.
//
// Virtual must be >= len(leaves).
func NewCommitment(leaves []common.Hash, virtual uint64) (protocol.History, error) {
	if len(leaves) == 0 {
		return emptyHistory, errors.New("must commit to at least one leaf")
	}
	if virtual < uint64(len(leaves)) {
		return emptyHistory, errors.New("virtual size must be >= len(leaves)")
	}
	comm := NewCommitter()
	firstLeaf := leaves[0]
	lastLeaf := leaves[len(leaves)-1]
	nzVirtual, err := newNonZero(virtual)
	if err != nil {
		return emptyHistory, err
	}
	comm.lastLeafProver = newLastLeafProver(nzVirtual)
	root, err := comm.ComputeRoot(leaves, virtual)
	if err != nil {
		return emptyHistory, err
	}
	lastLeafProof := comm.lastLeafProof()
	return protocol.History{
		Height:        virtual - 1,
		Merkle:        root,
		FirstLeaf:     firstLeaf,
		LastLeaf:      lastLeaf,
		LastLeafProof: lastLeafProof,
	}, nil
}

// ComputeRoot computes the merkle root of a virtual merkle tree.
func (h *historyCommitter) ComputeRoot(leaves []common.Hash, virtual uint64) (common.Hash, error) {
	lvLen := uint64(len(leaves))
	if lvLen == 0 {
		return emptyHash, nil
	}
	hashed := h.hashLeaves(leaves)
	limit := nextPowerOf2(virtual)
	depth := uint(math.Log2Floor(limit))
	n := uint(1)
	if virtual > lvLen {
		n = depth
	}
	nzVirt, err := newNonZero(virtual)
	if err != nil {
		return emptyHash, err
	}
	nzLimit, err := newNonZero(limit)
	if err != nil {
		return emptyHash, err
	}
	if err := h.populateFillers(&hashed[lvLen-1], n); err != nil {
		return emptyHash, err
	}
	h.cursor = treePosition{layer: uint64(depth), index: 0}
	return h.partialRoot(hashed, nzVirt, nzLimit)
}

// GeneratePrefixProof generates a prefix proof for a given prefix index.
//
// A prefix proof consists of the data needed to prove that a merkle root
// created from the leaves upto the prefix index represents a merkle tree which
// spans a specific prefix of the virtual merkle tree.
func (h *historyCommitter) GeneratePrefixProof(prefixIndex uint64, leaves []common.Hash, virtual uint64) ([]common.Hash, []common.Hash, error) {
	hashed := h.hashLeaves(leaves)
	prefixExpansion, proof, err := h.prefixAndProof(prefixIndex, hashed, virtual)
	if err != nil {
		return nil, nil, err
	}
	prefixExpansion = trimTrailingZeroHashes(prefixExpansion)
	proof = trimZeroes(proof)
	return prefixExpansion, proof, nil
}

// hashLeaves returns a slice of hashes of the leaves
func (h *historyCommitter) hashLeaves(leaves []common.Hash) []common.Hash {
	hashedLeaves := make([]common.Hash, len(leaves))
	for i := range leaves {
		hashedLeaves[i] = h.hash(&leaves[i])
	}
	return hashedLeaves
}

// completeRoot returns the root hash of a complete tree given leaves and limit.
//
// In the case of a complete virtual tree (when len(leaves) < limit)
// non-allocated leaves are filled with the filler at the corresponding layer
// in the tree.
//
// limit must be a power of two which is higher or equal to len(leaves).
//
// The historyCommitter's fillers must be precomputed to the necessary depth.
//
// Zero allocations (other than the local copy of the cursor)
// Computes O(len(leaves)) hashes.
func (h *historyCommitter) completeRoot(leaves []common.Hash, limit uint64) (common.Hash, error) {
	lvLen := len(leaves)
	if lvLen == 0 {
		return emptyHash, nil
	}
	if limit < uint64(lvLen) {
		return emptyHash, errors.New("limit must be >= len(leaves)")
	}
	if limit == 1 {
		h.handle(leaves[0])
		return leaves[0], nil
	}
	// Save the current cursor state
	curr := h.cursor.copy()
	depth := math.Log2Floor(limit)
	// From the bottom up, the "real" leaves are hashed into their parent
	// nodes in the next higher layer of the tree.
	for j := 0; j < depth; j++ {
		layerIndex := h.cursor.index * 2
		h.cursor.layer = uint64(j)
		// If the number of real leaves is even, then this loop will create
		// all of the hashes needed for calculating the next layer of the tree.
		// Even if the limit is not reached, it is okay, because the fillers are
		// precomputed for the non-virtual nodes higher up the tree.
		for i := 0; i < lvLen/2; i++ {
			h.cursor.index = layerIndex + uint64(2*i)
			h.handle(leaves[2*i])
			h.cursor.index = layerIndex + uint64(2*i+1)
			h.handle(leaves[2*i+1])
			h.hashInto(&leaves[i], &leaves[2*i], &leaves[2*i+1])
		}
		// If the number of leaves is odd, then the last real leaf needs to be
		// hashed with the filler at the same depth in the tree.
		if lvLen&1 == 1 {
			if j >= len(h.fillers) {
				return emptyHash, fmt.Errorf("programming error: insufficient fillers, want %d, got %d", j, len(h.fillers))
			}
			pIdx := (lvLen - 1) / 2
			h.cursor.index = layerIndex + uint64(lvLen-1)
			h.handle(leaves[lvLen-1])
			h.cursor.index = layerIndex + uint64(lvLen)
			h.handle(h.fillers[j])
			h.hashInto(&leaves[pIdx], &leaves[lvLen-1], &h.fillers[j])
		}
		lvLen = (lvLen + 1) / 2
		h.cursor.index = layerIndex * 2
	}
	// Restore the cursor to the original state
	h.cursor = curr
	h.handle(leaves[0])
	return leaves[0], nil
}

// partialRoot returns the merkle root of a possibly partial hashtree where the
// first layer is passed as leaves, then padded by repeating the last leaf
// until it reaches virtual and terminated with a single common.Hash{}.
//
// limit is a power of 2 which is greater or equal to virtual, and defines how
// deep the complete tree analogous to this partial one would be.
//
// Implementation note: It is very important that the historyCommitter's
// fillers member is populated correctly before calling this method. There must
// be at least Log2Floor(virtual-len(leaves)) filler nodes to properly pad each
// layer of the tree if it is a partial virtual tree.
//
// The algorithm is split in three different logical cases:
//
//  1. If the virtual length is less than or equal to half the limit (this can
//     never happen in the first iteration of the algorithm), the left half of
//     the tree is computed by recursion and the right half is an empty hash.
//  2. If the leaves all fit in the left half, then the left half can be
//     optimized as computing a complete virtual tree. This is the most common
//     starting scenario. The right half is computed by recursion.
//  3. If the leaves do not fit in the left half, then the left half can be
//     optimized as computing a complete tree. The right half is computed by
//     recursion.
func (h *historyCommitter) partialRoot(leaves []common.Hash, virtual, limit nonZero) (common.Hash, error) {
	lvLen := uint64(len(leaves))
	if lvLen == 0 {
		return emptyHash, errors.New("nil leaves")
	}
	if uint64(virtual) < lvLen {
		return emptyHash, fmt.Errorf("virtual %d should be >= num leaves %d", virtual, lvLen)
	}
	if limit < virtual {
		return emptyHash, fmt.Errorf("limit %d should be >= virtual %d", limit, virtual)
	}
	if limit == 1 {
		h.handle(leaves[0])
		return leaves[0], nil
	}
	// Save the current cursor state
	curr := h.cursor.copy()
	h.cursor.layer--
	var left, right common.Hash
	var err error
	mid := limit / 2

	// Deal with the left child first
	h.cursor.index = curr.index * 2
	if virtual > mid {
		// Case 2 or 3: The virtual size is greater than half the limit
		if lvLen > uint64(mid) {
			// Case 3: A pure complete subtree can be computed
			left, err = h.completeRoot(leaves[:mid], uint64(mid))
			if err != nil {
				return emptyHash, err
			}
		} else {
			// Case 2: A virtual complete subtree can be computed
			left, err = h.completeRoot(leaves, uint64(mid))
			if err != nil {
				return emptyHash, err
			}
		}
	} else {
		// Case 1: The virtual size is less than half the limit
		left, err = h.partialRoot(leaves, virtual, mid)
		if err != nil {
			return emptyHash, err
		}
	}

	// Deal with the right child
	h.cursor.index = curr.index*2 + 1
	if virtual > mid {
		// Case 2 or 3: The virtual size is greater than half the limit
		if lvLen > uint64(mid) {
			// Case 3: The leaves do not fit in the first half
			right, err = h.partialRoot(leaves[mid:], virtual-mid, mid)
			if err != nil {
				return emptyHash, err
			}
		} else {
			// Case 2: The leaves fit in the first half
			layer := math.Log2Floor(uint64(mid))
			if len(h.fillers) <= layer {
				return emptyHash, fmt.Errorf("programming error: insufficient fillers, want %d, got %d", layer, len(h.fillers))
			}
			if virtual == limit {
				// This is a special case where the entire right subtree is
				// made purely of virtual nodes, and it is a complete tree.
				// So, the root of the subtree will be the precomputed filler
				// at the current layer.
				right = h.fillers[layer]
				h.handle(right)
			} else {
				right, err = h.partialRoot([]common.Hash{h.fillers[0]}, virtual-mid, mid)
				if err != nil {
					return emptyHash, err
				}
			}
		}
	} else {
		// Case 1: The virtual size is less than half the limit
		h.handle(emptyHash)
		right = emptyHash
	}

	// Restore the cursor to the state for this level of recursion
	h.cursor = curr
	h.hashInto(&leaves[0], &left, &right)
	h.handle(leaves[0])
	return leaves[0], nil
}

func (h *historyCommitter) subtreeExpansion(leaves []common.Hash, virtual, limit uint64, stripped bool) (proof []common.Hash, err error) {
	lvLen := uint64(len(leaves))
	if lvLen == 0 {
		return make([]common.Hash, 0), nil
	}
	if virtual == 0 {
		for i := limit; i > 1; i /= 2 {
			proof = append(proof, emptyHash)
		}
		return
	}
	if limit == 0 {
		limit = nextPowerOf2(virtual)
	}
	if limit == virtual {
		left, err2 := h.completeRoot(leaves, limit)
		if err2 != nil {
			return nil, err2
		}
		if !stripped {
			for i := limit; i > 1; i /= 2 {
				proof = append(proof, emptyHash)
			}
		}
		return append(proof, left), nil
	}
	mid := limit / 2
	if lvLen > mid {
		left, err2 := h.completeRoot(leaves[:mid], mid)
		if err2 != nil {
			return nil, err2
		}
		proof, err = h.subtreeExpansion(leaves[mid:], virtual-mid, mid, stripped)
		if err != nil {
			return nil, err
		}
		return append(proof, left), nil
	}
	if virtual >= mid {
		left, err2 := h.completeRoot(leaves, mid)
		if err2 != nil {
			return nil, err2
		}
		if len(h.fillers) == 0 {
			return nil, errors.New("fillers is empty")
		}
		proof, err = h.subtreeExpansion([]common.Hash{h.fillers[0]}, virtual-mid, mid, stripped)
		if err != nil {
			return nil, err
		}
		return append(proof, left), nil
	}
	if stripped {
		return h.subtreeExpansion(leaves, virtual, mid, stripped)
	}
	expac, err := h.subtreeExpansion(leaves, virtual, mid, stripped)
	if err != nil {
		return nil, err
	}
	return append(expac, emptyHash), nil
}

func (h *historyCommitter) proof(index uint64, leaves []common.Hash, virtual, limit uint64) (tail []common.Hash, err error) {
	lvLen := uint64(len(leaves))
	if lvLen == 0 {
		return nil, errors.New("empty leaves slice")
	}
	if limit == 0 {
		limit = nextPowerOf2(virtual)
	}
	if limit == 1 {
		// Can only reach this with index == 0
		return
	}
	mid := limit / 2
	if index >= mid {
		if lvLen > mid {
			return h.proof(index-mid, leaves[mid:], virtual-mid, mid)
		}
		if len(h.fillers) == 0 {
			return nil, errors.New("fillers is empty")
		}
		return h.proof(index-mid, []common.Hash{h.fillers[0]}, virtual-mid, mid)
	}
	if lvLen > mid {
		tail, err = h.proof(index, leaves[:mid], mid, mid)
		if err != nil {
			return nil, err
		}
		right, err2 := h.subtreeExpansion(leaves[mid:], virtual-mid, mid, true)
		if err2 != nil {
			return nil, err2
		}
		for i := len(right) - 1; i >= 0; i-- {
			tail = append(tail, right[i])
		}
		return tail, nil
	}
	if virtual > mid {
		tail, err = h.proof(index, leaves, mid, mid)
		if err != nil {
			return nil, err
		}
		if len(h.fillers) == 0 {
			return nil, errors.New("fillers is empty")
		}
		right, err := h.subtreeExpansion([]common.Hash{h.fillers[0]}, virtual-mid, mid, true)
		if err != nil {
			return nil, err
		}
		for i := len(right) - 1; i >= 0; i-- {
			tail = append(tail, right[i])
		}
		return tail, nil
	}
	return h.proof(index, leaves, virtual, mid)
}

func (h *historyCommitter) prefixAndProof(index uint64, leaves []common.Hash, virtual uint64) (prefix []common.Hash, tail []common.Hash, err error) {
	lvLen := uint64(len(leaves))
	if lvLen == 0 {
		return nil, nil, errors.New("nil leaves")
	}
	if virtual == 0 {
		return nil, nil, errors.New("virtual size cannot be zero")
	}
	if lvLen > virtual {
		return nil, nil, fmt.Errorf("num leaves %d should be <= virtual %d", lvLen, virtual)
	}
	if index+1 > virtual {
		return nil, nil, fmt.Errorf("index %d + 1 should be <= virtual %d", index, virtual)
	}
	logVirtual := uint(math.Log2Floor(virtual) + 1)
	if err = h.populateFillers(&leaves[lvLen-1], logVirtual); err != nil {
		return nil, nil, err
	}

	if index+1 > lvLen {
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

// populateFillers returns a slice built recursively as
// ret[0] = the passed in leaf
// ret[i+1] = Hash(ret[i] + ret[i])
//
// Allocates n hashes
// Computes n-1 hashes
// Copies 1 hash
func (h *historyCommitter) populateFillers(leaf *common.Hash, n uint) error {
	if leaf == nil {
		return errors.New("nil leaf pointer")
	}
	h.fillers = make([]common.Hash, n)
	copy(h.fillers[0][:], (*leaf)[:])
	for i := uint(1); i < n; i++ {
		h.hashInto(&h.fillers[i], &h.fillers[i-1], &h.fillers[i-1])
	}
	return nil
}

// lastLeafProofPositions returns the positions in a virtual merkle tree
// of the sibling nodes that need to be hashed with the last leaf at each
// layer to compute the root of the tree.
func lastLeafProofPositions(virtual nonZero) []treePosition {
	if virtual == 1 {
		return []treePosition{}
	}
	limit := nextPowerOf2(uint64(virtual))
	depth := math.Log2Floor(limit)
	positions := make([]treePosition, depth)
	idx := uint64(virtual) - 1
	for l := range positions {
		positions[l] = sibling(idx, uint64(l))
		idx = parent(idx)
	}
	return positions
}

// sibling returns the position of the sibling of the node at the given layer
func sibling(index, layer uint64) treePosition {
	return treePosition{layer: layer, index: index ^ 1}
}

// parent returns the index of the parent of the node in the next higher layer
func parent(index uint64) uint64 {
	return index >> 1
}

func nextPowerOf2(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	n--         // Decrement n to handle the case where n is a power of 2
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
		if hashes[i] != emptyHash {
			return hashes[:i+1]
		}
	}
	// If all elements are zero, return an empty slice
	return []common.Hash{}
}

func trimZeroes(hashes []common.Hash) []common.Hash {
	newHashes := make([]common.Hash, 0, len(hashes))
	for _, h := range hashes {
		if h == emptyHash {
			continue
		}
		newHashes = append(newHashes, h)
	}
	return newHashes
}
