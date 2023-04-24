package validator

import (
	"context"
	"testing"
	"time"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/state-manager"
	"github.com/OffchainLabs/challenge-protocol-v2/testing/setup"
	"github.com/stretchr/testify/require"
)

func TestWatcher(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	createdData, err := setup.CreateTwoValidatorFork(ctx, &setup.CreateForkConfig{
		NumBlocks:     4,
		DivergeHeight: 0,
	})
	require.NoError(t, err)

	opts := []statemanager.Opt{
		statemanager.WithNumOpcodesPerBigStep(1),
		statemanager.WithMaxWavmOpcodesPerBlock(1),
	}

	honestStateManager, err := statemanager.NewWithAssertionStates(
		createdData.HonestValidatorStates,
		createdData.HonestValidatorInboxCounts,
		opts...,
	)
	require.NoError(t, err)
	evilStateManager, err := statemanager.NewWithAssertionStates(
		createdData.EvilValidatorStates,
		createdData.EvilValidatorInboxCounts,
		opts...,
	)
	require.NoError(t, err)

	challengeManager, err := createdData.Chains[0].SpecChallengeManager(ctx)
	require.NoError(t, err)

	leafAdder := func(stateManager statemanager.Manager, leaf protocol.Assertion) protocol.SpecEdge {
		startCommit, err := stateManager.HistoryCommitmentUpToBatch(ctx, 0, 0, 1)
		require.NoError(t, err)
		endCommit, err := stateManager.HistoryCommitmentUpToBatch(ctx, 0, protocol.LevelZeroBlockEdgeHeight, 1)
		require.NoError(t, err)
		prefixProof, err := stateManager.PrefixProofUpToBatch(ctx, 0, 0, protocol.LevelZeroBlockEdgeHeight, 1)
		require.NoError(t, err)
		edge, err := challengeManager.AddBlockChallengeLevelZeroEdge(
			ctx,
			leaf,
			startCommit,
			endCommit,
			prefixProof,
		)
		require.NoError(t, err)
		return edge
	}

	honestEdge := leafAdder(honestStateManager, createdData.Leaf1)
	require.Equal(t, protocol.BlockChallengeEdge, honestEdge.GetType())

	evilEdge := leafAdder(evilStateManager, createdData.Leaf2)
	require.Equal(t, protocol.BlockChallengeEdge, evilEdge.GetType())

	watcher := NewWatcher(createdData.Chains[0], createdData.Backend, time.Millisecond*100)
	go func() {
		watcher.Watch(ctx)
	}()
	time.Sleep(time.Millisecond * 500)
	cancel()
}

func TestMai(t *testing.T) {
	ctx := context.Background()
	createdData, err := setup.CreateTwoValidatorFork(ctx, &setup.CreateForkConfig{
		NumBlocks:     4,
		DivergeHeight: 0,
	})
	require.NoError(t, err)

	opts := []statemanager.Opt{
		statemanager.WithNumOpcodesPerBigStep(1),
		statemanager.WithMaxWavmOpcodesPerBlock(1),
	}
	_ = opts
	_ = createdData
	_ = opts
	_ = createdData
	_ = opts
	_ = createdData
	_ = opts
	_ = createdData
	_ = opts
	_ = createdData
	_ = opts
	_ = createdData
	_ = opts
	_ = createdData
	_ = opts
	_ = createdData
	_ = opts
	_ = createdData
	_ = opts
	_ = createdData
	_ = opts
	_ = createdData
	_ = opts
	_ = createdData
	_ = opts
	_ = createdData
	_ = opts
	_ = createdData
	_ = opts
	_ = createdData
	_ = opts
	_ = createdData
	_ = opts
	_ = createdData
	_ = opts
	_ = createdData
	_ = opts
	_ = createdData
	_ = opts
	_ = createdData
	_ = opts
	_ = createdData
}
