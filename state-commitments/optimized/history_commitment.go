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
	fromIndex uint64,
	toIndex option.Option[uint64],
	maxSize uint64,
) (common.Hash, error) {
	if len(realLeaves) == 0 {
		return common.Hash{}, errors.New("no leaves provided")
	}
	var emptyHash common.Hash
	from := fromIndex
	var to uint64
	if toIndex.IsSome() {
		to = toIndex.Unwrap()
		if to <= fromIndex {
			return emptyHash, errors.New("invalid range: end was <= start")
		}
		if to >= maxSize {
			return emptyHash, errors.New("invalid range: end was >= max size")
		}
	} else {
		to = uint64(len(realLeaves) - 1)
	}
	if to < uint64(len(realLeaves)) {
		// Case 0: the range is entirely within the real leaves' length.
		// we can simply compute a history commitment for it by slicing the list
		// and computing the root of the Merkle tree formed the slice of these leaves.
		// If this slice has a non-power of two length, we use virtual zero hashes
		// to build a sparse Merkle tree and compute its root.
		slicedLeaves := realLeaves[from:to]
		depth := uint64(math.Log2(float64(nextPowerOf2(uint64(len(slicedLeaves))))))
		return computeSparseRoot(slicedLeaves, depth)
	} else if from < uint64(len(realLeaves)) && to >= uint64(len(realLeaves)) {
		// Case 1: the `from` index is within the range of the length of the real
		// hashes list, but the `to` index exceeds it.
		// Here, we need to commit to commit to a Merkle tree formed by the following concatenation:
		// realLeaves[from:] ++ (realLeaves[-1] * (to - len(realLeaves)))
		// This means the leaves we are committing to are the real leaves up to the end of the list,
		// and the last leaf padded to the `to` index.
		return emptyHash, nil
	} else {
		// Case 2: Both the `from` and `to` indices are out of range of the real hashes list.
		// In this case, we commit to a Merkle tree formed by realLeaves[-1] * (to - from). That is,
		// we commit to a Merkle tree formed by the last leaf of the real leaves list repeated until
		// the specified range.
		slicedLeaves := []common.Hash{realLeaves[len(realLeaves)-1]}
		depth := uint64(math.Log2(float64(nextPowerOf2(to - from))))
		return computeSparseRoot(slicedLeaves, depth)
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
			layers[i] = append(layers[i], emptyHash) // TODO: Use zerohashes[i]
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

func virtualHashAtDepth(depth uint64, leaf common.Hash) common.Hash {
	currentHash := leaf
	for i := uint64(1); i <= depth; i++ {
		currentHash = crypto.Keccak256Hash(currentHash[:], currentHash[:])
	}
	return currentHash
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

// def first_and_last_hash_inclusion_proofs(
// 	real_hashes: [Hash],
// 	from_idx: usize,
// 	to_idx: usize,
// 	challenge_level: usize,
// ) -> ([Hash], [Hash]):
// 	max_size = max_size_at_challenge_level(challenge_level)

//   # Basic input validation
// 	assert(to_idx > from_idx)
// 	assert(to_idx <= max_size)

// 	# If the commitment we are asking for is in the range of the length
// 	# of the real hashes list, we can compute a tree and simple Merkle proofs
// 	if to_idx < len(real_hashes):
// 		 first_hash_proof = simple_merkle_proof(real_hashes, idx=0)
// 	   last_hash_proof = simple_merkle_proof(real_hashes, idx=len(real_hashes)-1)
// 	   return first_hash_proof, last_hash_proof

// 	# However, if the from index is within the range of the length of the real
// 	# hashes list, but the to index exceeds it, we need to commit to real hashes
// 	# before we hit the end of the list, and then commit to the last hash
// 	# repeated until we reach the to_idx
// 	else if from_idx < len(real_hashes) && to_idx >= len(real_hashes):
// 		first_hash_proof = simple_merkle_proof(real_hashes, idx=0)
// 		last_hash_proof = padded_subtree_proof(n, real_hashes[-1], n)
// 		return first_hash_proof, last_hash_proof
// 	# If both the from index and to index are out of range, we know the history
// 	# commitment consists of a Merkle tree of the same hash padded to a
// 	# specified length
// 	else:
// 	   n = (to_idx - from_idx) + 1
// 	   first_hash_proof = padded_subtree_proof(n, real_hashes[-1], 0)
// 	   last_hash_proof = padded_subtree_proof(n, real_hashes[-1], n)
// 	   return first_hash_proof, last_hash_proof

// def padded_subtree_root(
// 	hash: Hash,
// 	depth: usize,
// ) -> Hash:
// 	curr = hash
// 	for i = 0; i < depth; i++:
// 		curr = keccak256(curr, curr)
// 	return curr

// # Computes a Merkle proof for a padded subtree where all leaves are the same
// # element. This avoids the need to build a whole Merkle tree and allows us
// # to compute proofs more easily
// def padded_subtree_proof(padding_size, leaf, leaf_index):
//     proof = []
//     current_hash = keccak256(leaf)

//     # Compute all necessary hashes at each level
//     current_level_hashes = [current_hash] * padding_size
//     while len(current_level_hashes) > 1:
//         next_level_hashes = []
//         for i in range(0, len(current_level_hashes), 2):
//             if i + 1 < len(current_level_hashes):
//                 next_level_hashes.append(hash_pair(current_level_hashes[i], current_level_hashes[i + 1]))
//             else:
//                 next_level_hashes.append(hash_pair(current_level_hashes[i], current_level_hashes[i]))
//         sibling_index = leaf_index ^ 1
//         if sibling_index < len(current_level_hashes):
//             proof.append(current_level_hashes[sibling_index])
//         current_level_hashes = next_level_hashes
//         leaf_index //= 2
//     return proof
