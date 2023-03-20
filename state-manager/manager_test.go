package statemanager

import (
	"context"
	"math/big"
	"testing"

	"fmt"

	"encoding/binary"
	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/util/prefix-proofs"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
	"math/rand"
)

var _ = Manager(&Simulated{})

func TestPrefixProof(t *testing.T) {
	ctx := context.Background()
	hashes := make([]common.Hash, 10)
	for i := 0; i < len(hashes); i++ {
		hashes[i] = crypto.Keccak256Hash([]byte(fmt.Sprintf("%d", i)))
	}
	manager := New(hashes)

	loCommit, err := manager.HistoryCommitmentUpTo(ctx, 3)
	require.NoError(t, err)
	hiCommit, err := manager.HistoryCommitmentUpTo(ctx, 7)
	require.NoError(t, err)
	packedProof, err := manager.PrefixProof(ctx, 3, 7)
	require.NoError(t, err)

	data, err := ProofArgs.Unpack(packedProof)
	require.NoError(t, err)
	preExpansion := data[0].([][32]byte)
	proof := data[1].([][32]byte)

	preExpansionHashes := make([]common.Hash, len(preExpansion))
	for i := 0; i < len(preExpansion); i++ {
		preExpansionHashes[i] = preExpansion[i]
	}
	prefixProof := make([]common.Hash, len(proof))
	for i := 0; i < len(proof); i++ {
		prefixProof[i] = proof[i]
	}

	err = prefixproofs.VerifyPrefixProof(&prefixproofs.VerifyPrefixProofConfig{
		PreRoot:      loCommit.Merkle,
		PreSize:      4,
		PostRoot:     hiCommit.Merkle,
		PostSize:     8,
		PreExpansion: preExpansionHashes,
		PrefixProof:  prefixProof,
	})
	require.NoError(t, err)
}

func TestDivergenceGranularity(t *testing.T) {
	ctx := context.Background()
	numStates := uint64(10)
	bigStepSize := uint64(10)
	maxOpcodesPerBlock := uint64(100)

	honestStates, _, honestCounts := setupStates(t, numStates, 0 /* honest */)
	honestManager, err := NewWithAssertionStates(
		honestStates,
		honestCounts,
		WithMaxWavmOpcodesPerBlock(maxOpcodesPerBlock),
		WithNumOpcodesPerBigStep(bigStepSize),
	)
	require.NoError(t, err)

	fromBlock := uint64(1)
	toBlock := uint64(2)
	honestCommit, err := honestManager.BigStepLeafCommitment(
		ctx,
		fromBlock,
		toBlock,
	)
	require.NoError(t, err)

	t.Log("Big step leaf commitment height", honestCommit.Height)

	divergenceHeight := uint64(3)
	evilStates, _, evilCounts := setupStates(t, numStates, divergenceHeight)

	evilManager, err := NewWithAssertionStates(
		evilStates,
		evilCounts,
		WithBigStepStateDivergenceHeight(divergenceHeight),   // Diverges at the 3rd big step.
		WithSmallStepStateDivergenceHeight(divergenceHeight), // Diverges at the 3rd small step, within the 3rd big step.
		WithMaxWavmOpcodesPerBlock(maxOpcodesPerBlock),
		WithNumOpcodesPerBigStep(bigStepSize),
	)
	require.NoError(t, err)

	// Big step challenge granularity.
	evilCommit, err := evilManager.BigStepLeafCommitment(
		ctx,
		fromBlock,
		toBlock,
	)
	require.NoError(t, err)

	require.Equal(t, honestCommit.Height, evilCommit.Height)
	require.Equal(t, honestCommit.FirstLeaf, evilCommit.FirstLeaf)
	require.NotEqual(t, honestCommit.Merkle, evilCommit.Merkle)

	// Check if big step commitments between the validators agree before the divergence height.
	checkHeight := divergenceHeight - 1
	honestCommit, err = honestManager.BigStepCommitmentUpTo(
		ctx,
		fromBlock,
		toBlock,
		checkHeight,
	)
	require.NoError(t, err)
	evilCommit, err = evilManager.BigStepCommitmentUpTo(
		ctx,
		fromBlock,
		toBlock,
		checkHeight,
	)
	require.NoError(t, err)
	require.Equal(t, honestCommit, evilCommit)

	// Check if big step commitments between the validators disagree starting at the divergence height.
	honestCommit, err = honestManager.BigStepCommitmentUpTo(
		ctx,
		fromBlock,
		toBlock,
		divergenceHeight,
	)
	require.NoError(t, err)
	evilCommit, err = evilManager.BigStepCommitmentUpTo(
		ctx,
		fromBlock,
		toBlock,
		divergenceHeight,
	)
	require.NoError(t, err)

	require.Equal(t, honestCommit.Height, evilCommit.Height)
	require.Equal(t, honestCommit.FirstLeaf, evilCommit.FirstLeaf)
	require.NotEqual(t, honestCommit.Merkle, evilCommit.Merkle)

	// Small step challenge granularity.
	honestCommit, err = honestManager.SmallStepLeafCommitment(
		ctx,
		fromBlock,
		toBlock,
	)
	require.NoError(t, err)

	evilCommit, err = evilManager.SmallStepLeafCommitment(
		ctx,
		fromBlock,
		toBlock,
	)
	require.NoError(t, err)

	t.Log("Small step leaf commitment height", honestCommit.Height)
	require.Equal(t, honestCommit.Height, evilCommit.Height)
	require.Equal(t, honestCommit.FirstLeaf, evilCommit.FirstLeaf)
	require.NotEqual(t, honestCommit.Merkle, evilCommit.Merkle)

	// Check if small step commitments between the validators agree before the divergence height.
	checkHeight = divergenceHeight - 1
	honestCommit, err = honestManager.SmallStepCommitmentUpTo(
		ctx,
		fromBlock,
		toBlock,
		checkHeight,
	)
	require.NoError(t, err)
	evilCommit, err = evilManager.SmallStepCommitmentUpTo(
		ctx,
		fromBlock,
		toBlock,
		checkHeight,
	)
	require.NoError(t, err)
	require.Equal(t, honestCommit, evilCommit)

	// Check if small step commitments between the validators disagree starting at the divergence height.
	honestCommit, err = honestManager.SmallStepCommitmentUpTo(
		ctx,
		fromBlock,
		toBlock,
		divergenceHeight,
	)
	require.NoError(t, err)
	evilCommit, err = evilManager.SmallStepCommitmentUpTo(
		ctx,
		fromBlock,
		toBlock,
		divergenceHeight,
	)
	require.NoError(t, err)

	require.Equal(t, honestCommit.Height, evilCommit.Height)
	require.Equal(t, honestCommit.FirstLeaf, evilCommit.FirstLeaf)
	require.NotEqual(t, honestCommit.Merkle, evilCommit.Merkle)
}

func setupStates(t *testing.T, numStates, divergenceHeight uint64) ([]*protocol.ExecutionState, []common.Hash, []*big.Int) {
	t.Helper()
	states := make([]*protocol.ExecutionState, numStates)
	roots := make([]common.Hash, numStates)
	inboxCounts := make([]*big.Int, numStates)
	for i := uint64(0); i < numStates; i++ {
		var blockHash common.Hash
		if divergenceHeight == 0 || i < divergenceHeight {
			blockHash = crypto.Keccak256Hash([]byte(fmt.Sprintf("%d", i)))
		} else {
			junkRoot := make([]byte, 32)
			_, err := rand.Read(junkRoot)
			require.NoError(t, err)
			blockHash = crypto.Keccak256Hash(junkRoot)
		}
		state := &protocol.ExecutionState{
			GlobalState: protocol.GoGlobalState{
				BlockHash: blockHash,
				Batch:     1,
			},
			MachineStatus: protocol.MachineStatusFinished,
		}
		states[i] = state
		roots[i] = protocol.ComputeStateHash(state, big.NewInt(1))
		inboxCounts[i] = big.NewInt(1)
	}
	return states, roots, inboxCounts
}

func TestPrefixProofs(t *testing.T) {
	ctx := context.Background()
	for _, c := range []struct {
		lo uint64
		hi uint64
	}{
		{0, 1},
		{0, 3},
		{1, 2},
		{1, 3},
		{1, 15},
		{17, 255},
		{23, 255},
		{20, 511},
	} {
		leaves := hashesForUints(0, c.hi+1)
		manager := New(leaves)
		packedProof, err := manager.PrefixProof(ctx, c.lo, c.hi)
		require.NoError(t, err)

		data, err := ProofArgs.Unpack(packedProof)
		require.NoError(t, err)
		preExpansion := data[0].([][32]byte)
		proof := data[1].([][32]byte)

		preExpansionHashes := make([]common.Hash, len(preExpansion))
		for i := 0; i < len(preExpansion); i++ {
			preExpansionHashes[i] = preExpansion[i]
		}
		prefixProof := make([]common.Hash, len(proof))
		for i := 0; i < len(proof); i++ {
			prefixProof[i] = proof[i]
		}

		postExpansion, err := manager.HistoryCommitmentUpTo(ctx, c.hi)
		require.NoError(t, err)

		cfg := &prefixproofs.VerifyPrefixProofConfig{
			PreRoot:      prefixproofs.Root(preExpansionHashes),
			PreSize:      c.lo + 1,
			PostRoot:     postExpansion.Merkle,
			PostSize:     c.hi + 1,
			PreExpansion: preExpansionHashes,
			PrefixProof:  prefixProof,
		}
		err = prefixproofs.VerifyPrefixProof(cfg)
		require.NoError(t, err)
	}
}

func hashesForUints(lo, hi uint64) []common.Hash {
	ret := []common.Hash{}
	for i := lo; i < hi; i++ {
		ret = append(ret, hashForUint(i))
	}
	return ret
}

func hashForUint(x uint64) common.Hash {
	return crypto.Keccak256Hash(binary.BigEndian.AppendUint64([]byte{}, x))
}
