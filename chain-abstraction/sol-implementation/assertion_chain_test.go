package solimpl_test

import (
	"context"
	"testing"

	protocol "github.com/OffchainLabs/challenge-protocol-v2/chain-abstraction"
	solimpl "github.com/OffchainLabs/challenge-protocol-v2/chain-abstraction/sol-implementation"
	"github.com/OffchainLabs/challenge-protocol-v2/testing/setup"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestCreateAssertion(t *testing.T) {
	ctx := context.Background()
	cfg, err := setup.SetupChainsWithEdgeChallengeManager()
	require.NoError(t, err)
	chain := cfg.Chains[0]
	backend := cfg.Backend

	t.Run("OK", func(t *testing.T) {

		latestBlockHash := common.Hash{}
		for i := uint64(0); i < 100; i++ {
			latestBlockHash = backend.Commit()
		}

		prevState := &protocol.ExecutionState{
			GlobalState:   protocol.GoGlobalState{},
			MachineStatus: protocol.MachineStatusFinished,
		}
		postState := &protocol.ExecutionState{
			GlobalState: protocol.GoGlobalState{
				BlockHash:  latestBlockHash,
				SendRoot:   common.Hash{},
				Batch:      1,
				PosInBatch: 0,
			},
			MachineStatus: protocol.MachineStatusFinished,
		}
		_, err := chain.CreateAssertion(ctx, prevState, postState)
		require.NoError(t, err)

		_, err = chain.CreateAssertion(ctx, prevState, postState)
		require.ErrorContains(t, err, "ALREADY_STAKED")
	})
	t.Run("can create fork", func(t *testing.T) {
		assertionChain := cfg.Chains[1]

		for i := uint64(0); i < 100; i++ {
			backend.Commit()
		}

		prevState := &protocol.ExecutionState{
			GlobalState:   protocol.GoGlobalState{},
			MachineStatus: protocol.MachineStatusFinished,
		}
		postState := &protocol.ExecutionState{
			GlobalState: protocol.GoGlobalState{
				BlockHash:  common.BytesToHash([]byte("evil hash")),
				SendRoot:   common.Hash{},
				Batch:      1,
				PosInBatch: 0,
			},
			MachineStatus: protocol.MachineStatusFinished,
		}
		_, err := assertionChain.CreateAssertion(ctx, prevState, postState)
		require.NoError(t, err)
	})
}

func TestAssertionBySequenceNum(t *testing.T) {
	ctx := context.Background()
	cfg, err := setup.SetupChainsWithEdgeChallengeManager()
	require.NoError(t, err)
	chain := cfg.Chains[0]
	_, err = chain.AssertionBySequenceNum(ctx, 1)
	require.NoError(t, err)

	_, err = chain.AssertionBySequenceNum(ctx, 2)
	require.ErrorIs(t, err, solimpl.ErrNotFound)
}

func TestAssertion_Confirm(t *testing.T) {
	ctx := context.Background()
	t.Run("OK", func(t *testing.T) {
		cfg, err := setup.SetupChainsWithEdgeChallengeManager()
		require.NoError(t, err)

		chain := cfg.Chains[0]
		backend := cfg.Backend

		assertionBlockHash := common.Hash{}
		for i := uint64(0); i < 100; i++ {
			assertionBlockHash = backend.Commit()
		}

		prevState := &protocol.ExecutionState{
			GlobalState:   protocol.GoGlobalState{},
			MachineStatus: protocol.MachineStatusFinished,
		}
		postState := &protocol.ExecutionState{
			GlobalState: protocol.GoGlobalState{
				BlockHash:  assertionBlockHash,
				SendRoot:   common.Hash{},
				Batch:      1,
				PosInBatch: 0,
			},
			MachineStatus: protocol.MachineStatusFinished,
		}
		_, err = chain.CreateAssertion(ctx, prevState, postState)
		require.NoError(t, err)

		err = chain.Confirm(ctx, assertionBlockHash, common.Hash{})
		require.ErrorIs(t, err, solimpl.ErrTooSoon)

		for i := uint64(0); i < 100; i++ {
			backend.Commit()
		}
		require.NoError(t, chain.Confirm(ctx, assertionBlockHash, common.Hash{}))
		require.ErrorIs(t, solimpl.ErrNoUnresolved, chain.Confirm(ctx, assertionBlockHash, common.Hash{}))
	})
}

func TestAssertion_Reject(t *testing.T) {
	ctx := context.Background()

	t.Run("Can reject assertion", func(t *testing.T) {
		t.Skip("TODO: Can't reject assertion. Blocked by one step proof")
	})

	t.Run("Already confirmed assertion", func(t *testing.T) {
		cfg, err := setup.SetupChainsWithEdgeChallengeManager()
		require.NoError(t, err)

		chain := cfg.Chains[0]
		backend := cfg.Backend

		assertionBlockHash := common.Hash{}
		for i := uint64(0); i < 100; i++ {
			assertionBlockHash = backend.Commit()
		}

		prevState := &protocol.ExecutionState{
			GlobalState:   protocol.GoGlobalState{},
			MachineStatus: protocol.MachineStatusFinished,
		}
		postState := &protocol.ExecutionState{
			GlobalState: protocol.GoGlobalState{
				BlockHash:  assertionBlockHash,
				SendRoot:   common.Hash{},
				Batch:      1,
				PosInBatch: 0,
			},
			MachineStatus: protocol.MachineStatusFinished,
		}
		_, err = chain.CreateAssertion(ctx, prevState, postState)
		require.NoError(t, err)

		for i := uint64(0); i < 100; i++ {
			backend.Commit()
		}
		require.NoError(t, chain.Confirm(ctx, assertionBlockHash, common.Hash{}))
		require.ErrorIs(t, solimpl.ErrNoUnresolved, chain.Reject(ctx, cfg.Accounts[0].AccountAddr))
	})
}

func TestChallengePeriodBlocks(t *testing.T) {
	ctx := context.Background()
	cfg, err := setup.SetupChainsWithEdgeChallengeManager()
	require.NoError(t, err)
	chain := cfg.Chains[0]

	manager, err := chain.SpecChallengeManager(ctx)
	require.NoError(t, err)

	chalPeriod, err := manager.ChallengePeriodBlocks(ctx)
	require.NoError(t, err)
	require.Equal(t, uint64(20), chalPeriod)
}
