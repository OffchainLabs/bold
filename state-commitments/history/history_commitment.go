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
	// layer is the layer of the tree that the cursor is currently tracking.
	layer uint64
	// index is the index of the leaf in this layer of the tree.
	index uint64
}

func (c *treePosition) copy() treePosition {
	return treePosition{layer: c.layer, index: c.index}
}

type historyCommitter struct {
	lastLeafFillers []common.Hash
	keccak          crypto.KeccakState
	cursor          treePosition
	lastLeafProver  *lastLeafProver
}

func NewCommitter() *historyCommitter {
	return &historyCommitter{
		lastLeafFillers: make([]common.Hash, 0),
		keccak:          crypto.NewKeccakState(),
	}
}

// soughtHash holds the hash that is being sought and whether it has been found.
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
		fmt.Printf("found sibling %v at %v\n", hash, pos)
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

// proof returns the merkle inclusion proof for the last leaf in the virtual
// merkle tree.
//
// If the proof is not complete (i.e. some sibling nodes are missing), the
// sibling nodes are filled in with the lastLeafFillers.
//
// The reason this works, is that the only nodes which are not visited when
// computing the merkle root are those which are in some complete subtree of
// virtual nodes.
func (h *historyCommitter) lastLeafProof() []common.Hash {
	for pos, sibling := range h.lastLeafProver.positions {
		if !sibling.found {
			// fmt.Printf("pos %v not found\n", pos)
			*h.lastLeafProver.positions[pos].hash = h.lastLeafFillers[pos.layer]
			// fmt.Printf("filling with: %v\n", h.lastLeafFillers[pos.layer])
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

// NewCommitment produces a history commitment from a list of leaves that are
// virtually padded using the last leaf in the list to some virtual length.
//
// Virtual must be >= len(leaves).
func NewCommitment(leaves []common.Hash, virtual uint64) (protocol.History, error) {
	if len(leaves) == 0 {
		return emptyHistory, errors.New("must commit to at least one leaf")
	}
	if virtual < uint64(len(leaves)) {
		return emptyHistory, errors.New("virtual size must be greater than or equal to the number of leaves")
	}
	// fmt.Println("leaves", leaves, "virtual", virtual, "len(leaves)", len(leaves))
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

func (h *historyCommitter) ComputeRoot(leaves []common.Hash, virtual uint64) (common.Hash, error) {
	lvLen := uint64(len(leaves))
	if lvLen == 0 {
		return emptyHash, nil
	}
	rehashedLeaves := h.hashLeaves(leaves)
	limit := nextPowerOf2(virtual)
	depth := uint(math.Log2Floor(limit))
	n := uint(1)
	if virtual > lvLen {
		n = depth
	}
	var err error
	h.lastLeafFillers, err = h.precomputeRepeatedHashes(&rehashedLeaves[lvLen-1], n)
	if err != nil {
		return emptyHash, err
	}
	h.cursor = treePosition{layer: uint64(depth), index: 0}
	return h.computeVirtualSparseTree(rehashedLeaves, virtual, limit)
}

func (h *historyCommitter) GeneratePrefixProof(prefixIndex uint64, leaves []common.Hash, virtual uint64) ([]common.Hash, []common.Hash, error) {
	rehashedLeaves := h.hashLeaves(leaves)
	prefixExpansion, proof, err := h.prefixAndProof(prefixIndex, rehashedLeaves, virtual)
	if err != nil {
		return nil, nil, err
	}
	prefixExpansion = trimTrailingZeroHashes(prefixExpansion)
	proof = trimZeroes(proof)
	return prefixExpansion, proof, nil
}

// hashLeaves returns a slich of hashes of the leaves
func (h *historyCommitter) hashLeaves(leaves []common.Hash) []common.Hash {
	hashedLeaves := make([]common.Hash, len(leaves))
	for i := range leaves {
		hashedLeaves[i] = h.hash(&leaves[i])
	}
	return hashedLeaves
}

// computeSparseTree returns the htr of a hashtree with the given leaves and
// limit. Any non-allocated leaf is filled with the passed zeroHash of depth 0.
// Recursively, any non allocated intermediate layer at depth i is filled with
// the passed zeroHash.
// limit is assumed to be a power of two which is higher or equal than the
// length of the leaves.
// fillers is assumed to be precomputed to the necessary limit.
// It is a programming error to call this function with a limit of 0.
//
// Zero allocations
// Computes O(len(leaves)) hashes.
func (h *historyCommitter) computeSparseTree(leaves []common.Hash, limit uint64) (common.Hash, error) {
	if limit == 0 {
		panic("limit must be greater than 0")
	}
	lvLen := len(leaves)
	if lvLen == 0 {
		return emptyHash, nil
	}
	if limit < 2 {
		h.handle(leaves[0])
		return leaves[0], nil
	}
	// Save the current cursor state
	curr := h.cursor.copy()
	depth := math.Log2Floor(limit)
	for j := 0; j < depth; j++ {
		layerIndex := h.cursor.index * 2
		h.cursor.layer = uint64(j)
		// Check to ensure we don't access out of bounds.
		for i := 0; i < lvLen/2; i++ {
			h.cursor.index = layerIndex + uint64(2*i)
			h.handle(leaves[2*i])
			h.cursor.index = layerIndex + uint64(2*i+1)
			h.handle(leaves[2*i+1])
			h.hashInto(&leaves[i], &leaves[2*i], &leaves[2*i+1])
		}
		if lvLen&1 == 1 {
			// Check to ensure m-1 is a valid index.
			if j >= len(h.lastLeafFillers) {
				// Handle the case where j is out of range for fillers.
				return emptyHash, errors.New("insufficient fillers")
			}
			target := (lvLen - 1) / 2
			h.cursor.index = layerIndex + uint64(lvLen-1)
			h.handle(leaves[lvLen-1])
			h.cursor.index = layerIndex + uint64(lvLen)
			h.handle(h.lastLeafFillers[j])
			h.hashInto(&leaves[target], &leaves[lvLen-1], &h.lastLeafFillers[j])
		}
		lvLen = (lvLen + 1) / 2
		h.cursor.index = layerIndex * 2
	}
	// Restore the cursor to the original state
	h.cursor = curr
	h.handle(leaves[0])
	return leaves[0], nil
}

// computeVirtualSparseTree returns the merkle root of a hashtree where the
// first layer is passed as leaves, then completed with the last leaf until it
// reaches virtual and finally completed with zero hashes until it reaches
// limit. limit power of 2 which is greater or equal to virtual.
//
// Implementation note: It is very important that the historyCommitter's
// lastLeafFillers member is populated correctly before calling this method.
// There must be Log2Floor(virtual-len(leaves)) filler nodes to properly pad
// each layer of the tree.
//
// The algorithm is split in three different logical cases:
//  1. If the virtual length is less than half the limit (this can never happen
//     in the first iteration of the algorithm), then the first half of the tree
//     is computed by recursion and the second half is an empty hash.
//  2. If the leaves all fit in the first half, then we can optimize the first
//     half to being a simple sparse tree, just that instead of filling with
//     zero hashes we fill with the precomputed virtual hashes. This is the most
//     common starting scenario. The second part is computed by recursion.
//  3. If the leaves do not fit in the first half, then we can compute the first
//     half of the tree as a normal complete hashtree. The second part is
//     computed by recursion.
func (h *historyCommitter) computeVirtualSparseTree(leaves []common.Hash, virtual, limit uint64) (common.Hash, error) {
	lvLen := uint64(len(leaves))
	if lvLen == 0 {
		return emptyHash, errors.New("nil leaves")
	}
	if virtual < lvLen {
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
		if lvLen > mid {
			// Case 3: The leaves do not fit in the first half
			left, err = h.computeSparseTree(leaves[:mid], mid)
		} else {
			// Case 2: The leaves fit in the first half
			left, err = h.computeSparseTree(leaves, mid)
		}
	} else {
		// Case 1: The virtual size is less than half the limit
		left, err = h.computeVirtualSparseTree(leaves, virtual, mid)
	}
	// Any of the three cases can return an error
	if err != nil {
		return emptyHash, err
	}

	// Deal with the right child
	h.cursor.index = curr.index*2 + 1
	if virtual > mid {
		// Case 2 or 3: The virtual size is greater than half the limit
		if lvLen > mid {
			// Case 3: The leaves do not fit in the first half
			right, err = h.computeVirtualSparseTree(leaves[mid:], virtual-mid, mid)
			if err != nil {
				return emptyHash, err
			}
		} else {
			// Case 2: The leaves fit in the first half
			if virtual == limit {
				// This is a special case where the entire right subtree is
				// made purely of virtual nodes, nand it is a complete tree
				// (because limit is a power of 2).
				if len(h.lastLeafFillers) <= math.Log2Floor(mid) {
					return emptyHash, errors.New("insufficient lastLeafFillers")
				}
				right = h.lastLeafFillers[math.Log2Floor(mid)]
				h.handle(right)
			} else {
				if len(h.lastLeafFillers) <= 0 {
					return emptyHash, errors.New("empty lastLeafFillers")
				}
				right, err = h.computeVirtualSparseTree([]common.Hash{h.lastLeafFillers[0]}, virtual-mid, mid)
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
		left, err2 := h.computeSparseTree(leaves, limit)
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
		left, err2 := h.computeSparseTree(leaves[:mid], mid)
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
		left, err2 := h.computeSparseTree(leaves, mid)
		if err2 != nil {
			return nil, err2
		}
		if len(h.lastLeafFillers) > 0 {
			proof, err = h.subtreeExpansion([]common.Hash{h.lastLeafFillers[0]}, virtual-mid, mid, stripped)
			if err != nil {
				return nil, err
			}
			return append(proof, left), nil
		} else {
			return nil, errors.New("lastLeafFillers is empty")
		}
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
		if len(h.lastLeafFillers) > 0 {
			return h.proof(index-mid, []common.Hash{h.lastLeafFillers[0]}, virtual-mid, mid)
		} else {
			return nil, errors.New("lastLeafFillers is empty")
		}
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
		if len(h.lastLeafFillers) > 0 {
			right, err := h.subtreeExpansion([]common.Hash{h.lastLeafFillers[0]}, virtual-mid, mid, true)
			if err != nil {
				return nil, err
			}
			for i := len(right) - 1; i >= 0; i-- {
				tail = append(tail, right[i])
			}
		} else {
			return nil, errors.New("lastLeafFillers is empty")
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
	h.lastLeafFillers, err = h.precomputeRepeatedHashes(&leaves[lvLen-1], logVirtual)
	if err != nil {
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

// precomputeRepeatedHashes returns a slice where built recursively as
// ret[0] = the passed in leaf
// ret[i+1] = Hash(ret[i] + ret[i])
//
// Allocates n hashes
// Computes n-1 hashes
// Copies 1 hash
func (h *historyCommitter) precomputeRepeatedHashes(leaf *common.Hash, n uint) ([]common.Hash, error) {
	if leaf == nil {
		return nil, errors.New("nil leaf pointer")
	}
	fLen := uint(len(h.lastLeafFillers))
	if fLen > 0 && h.lastLeafFillers[0] == *leaf && fLen >= n {
		return h.lastLeafFillers, nil
	}
	if n == 0 {
		return []common.Hash{*leaf}, nil
	}
	ret := make([]common.Hash, n)
	copy(ret[0][:], (*leaf)[:])
	for i := uint(1); i < n; i++ {
		h.hashInto(&ret[i], &ret[i-1], &ret[i-1])
	}
	return ret, nil
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

// sibling returns the cursor of the sibling of the node at the given layer
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
