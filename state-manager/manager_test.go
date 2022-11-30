package statemanager

import (
	"context"
	"testing"

	"github.com/OffchainLabs/new-rollup-exploration/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestSpecificProof(t *testing.T) {
	lo := uint64(4)
	stateRootsInCommon := make([]common.Hash, 0)
	for i := uint64(0); i <= 3; i++ {
		stateRootsInCommon = append(stateRootsInCommon, util.HashForUint(i))
	}
	aliceRoots := append(
		stateRootsInCommon,
		common.BytesToHash([]byte("a4")),
		common.BytesToHash([]byte("a5")),
		common.BytesToHash([]byte("a6")),
	)
	hi := uint64(len(aliceRoots)) - 1
	manager := New(aliceRoots)
	ctx := context.Background()
	historyCommit, err := manager.HistoryCommitmentUpTo(ctx, lo)
	require.NoError(t, err)
	historyCommit2, err := manager.HistoryCommitmentUpTo(ctx, hi)
	require.NoError(t, err)
	proof, err := manager.PrefixProof(ctx, lo, hi)
	require.NoError(t, err)
	t.Log("JAJSDJKAJSJDKAJSKDJAD")
	t.Logf("from %d to %d", lo, hi)
	t.Logf("%d and %#x\n", historyCommit.Height, historyCommit.Merkle)
	t.Logf("%d and %#x\n", historyCommit2.Height, historyCommit2.Merkle)
	err = util.VerifyPrefixProof(historyCommit, historyCommit2, proof)
	require.NoError(t, err)
}
