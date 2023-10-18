// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE

package history

import (
	"fmt"
	"testing"

	"github.com/OffchainLabs/bold/mmap"
	inclusionproofs "github.com/OffchainLabs/bold/state-commitments/inclusion-proofs"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestHistoryCommitment_LeafProofs(t *testing.T) {
	leavesMmap, err := mmap.NewMmap(8)
	defer leavesMmap.Free()
	require.NoError(t, err)
	for i := 0; i < leavesMmap.Length(); i++ {
		leavesMmap.Set(i, common.BytesToHash([]byte(fmt.Sprintf("%d", i))))
	}
	history, err := New(leavesMmap)
	require.NoError(t, err)
	require.Equal(t, history.FirstLeaf, leavesMmap.Get(0))
	require.Equal(t, history.LastLeaf, leavesMmap.Get(leavesMmap.Length()-1))

	computed, err := inclusionproofs.CalculateRootFromProof(history.LastLeafProof, history.Height, history.LastLeaf)
	require.NoError(t, err)
	require.Equal(t, history.Merkle, computed)
	computed, err = inclusionproofs.CalculateRootFromProof(history.FirstLeafProof, 0, history.FirstLeaf)
	require.NoError(t, err)
	require.Equal(t, history.Merkle, computed)
}
