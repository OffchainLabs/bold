// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE

package inclusionproofs

import (
	"fmt"
	"testing"

	"github.com/OffchainLabs/bold/mmap"
	prefixproofs "github.com/OffchainLabs/bold/state-commitments/prefix-proofs"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestInclusionProof(t *testing.T) {
	leavesMmap, err := mmap.NewMmap(8)
	require.NoError(t, err)
	defer leavesMmap.Free()
	for i := 0; i < leavesMmap.Length(); i++ {
		leavesMmap.Set(i, common.BytesToHash([]byte(fmt.Sprintf("%d", i))))
	}
	index := uint64(0)
	proof, err := GenerateInclusionProof(leavesMmap, index)
	require.NoError(t, err)
	require.Equal(t, true, len(proof) > 0)

	computedRoot, err := CalculateRootFromProof(proof, index, leavesMmap.Get(int(index)))
	require.NoError(t, err)

	exp := prefixproofs.NewEmptyMerkleExpansion()
	for i := 0; i < leavesMmap.Length(); i++ {
		exp, err = prefixproofs.AppendLeaf(exp, leavesMmap.Get(i))
		require.NoError(t, err)
	}

	root, err := prefixproofs.Root(exp)
	require.NoError(t, err)

	t.Run("proof verifies", func(t *testing.T) {
		require.Equal(t, root, computedRoot)
	})
	t.Run("first leaf proof", func(t *testing.T) {
		index = uint64(0)
		proof, err = GenerateInclusionProof(leavesMmap, index)
		require.NoError(t, err)
		require.Equal(t, true, len(proof) > 0)
		computedRoot, err = CalculateRootFromProof(proof, index, leavesMmap.Get(int(index)))
		require.NoError(t, err)
		require.Equal(t, root, computedRoot)
	})
	t.Run("last leaf proof", func(t *testing.T) {
		index = uint64(leavesMmap.Length() - 1)
		proof, err = GenerateInclusionProof(leavesMmap, index)
		require.NoError(t, err)
		require.Equal(t, true, len(proof) > 0)
		computedRoot, err = CalculateRootFromProof(proof, index, leavesMmap.Get(int(index)))
		require.NoError(t, err)
		require.Equal(t, root, computedRoot)
	})
	t.Run("Invalid inputs", func(t *testing.T) {
		// Empty tree should not generate a proof.
		_, err := GenerateInclusionProof(mmap.Mmap{}, 0)
		require.Equal(t, ErrInvalidLeaves, err)

		// Index greater than the number of leaves should not generate a proof.
		_, err = GenerateInclusionProof(leavesMmap, uint64(leavesMmap.Length()))
		require.Equal(t, ErrInvalidLeaves, err)

		// Proof with more than 256 elements should not calculate a root...
		_, err = CalculateRootFromProof(make([]common.Hash, 257), 0, common.Hash{})
		require.Equal(t, ErrProofTooLong, err)

		// ... but proof with exactly 256 elements should be OK.
		_, err = CalculateRootFromProof(make([]common.Hash, 256), 0, common.Hash{})
		require.NotEqual(t, ErrProofTooLong, err)
	})
}
