// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/challenge-protocol-v2/blob/main/LICENSE

package assertions_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/OffchainLabs/challenge-protocol-v2/assertions"
	protocol "github.com/OffchainLabs/challenge-protocol-v2/chain-abstraction"
	challengemanager "github.com/OffchainLabs/challenge-protocol-v2/challenge-manager"
	"github.com/OffchainLabs/challenge-protocol-v2/challenge-manager/types"
	"github.com/OffchainLabs/challenge-protocol-v2/solgen/go/rollupgen"
	"github.com/OffchainLabs/challenge-protocol-v2/testing/mocks"
	"github.com/OffchainLabs/challenge-protocol-v2/testing/setup"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestScanner_ProcessAssertionCreation(t *testing.T) {
	ctx := context.Background()
	t.Run("no fork detected", func(t *testing.T) {
		manager, _, mockStateProvider, cfg := setupChallengeManager(t)

		prev := &mocks.MockAssertion{
			MockPrevId:         mockId(1),
			MockId:             mockId(1),
			MockStateHash:      common.Hash{},
			MockHasSecondChild: false,
		}
		ev := &mocks.MockAssertion{
			MockPrevId:         mockId(1),
			MockId:             mockId(2),
			MockStateHash:      common.BytesToHash([]byte("bar")),
			MockHasSecondChild: false,
		}

		p := &mocks.MockProtocol{}
		p.On("SpecChallengeManager", ctx).Return(&mocks.MockSpecChallengeManager{}, nil)
		p.On("ReadAssertionCreationInfo", ctx, mockId(2)).Return(&protocol.AssertionCreatedInfo{
			ParentAssertionHash: common.Hash(mockId(1)),
			AfterState:          rollupgen.ExecutionState{},
		}, nil)
		p.On("GetAssertion", ctx, mockId(2)).Return(ev, nil)
		p.On("GetAssertion", ctx, mockId(1)).Return(prev, nil)
		scanner := assertions.NewScanner(p, mockStateProvider, cfg.Backend, manager, cfg.Addrs.Rollup, "", time.Second)

		err := scanner.ProcessAssertionCreation(ctx, ev.Id())
		require.NoError(t, err)
		require.Equal(t, uint64(1), scanner.AssertionsProcessed())
		require.Equal(t, uint64(0), scanner.ForksDetected())
		require.Equal(t, uint64(0), scanner.ChallengesSubmitted())
	})
	t.Run("fork leads validator to challenge leaf", func(t *testing.T) {
		ctx := context.Background()
		createdData, err := setup.CreateTwoValidatorFork(ctx, &setup.CreateForkConfig{
			DivergeBlockHeight: 5,
		})
		require.NoError(t, err)

		manager, err := challengemanager.New(
			ctx,
			createdData.Chains[1],
			createdData.Backend,
			createdData.HonestStateManager,
			createdData.Addrs.Rollup,
			challengemanager.WithMode(types.MakeMode),
		)
		require.NoError(t, err)
		scanner := assertions.NewScanner(createdData.Chains[1], createdData.HonestStateManager, createdData.Backend, manager, createdData.Addrs.Rollup, "", time.Second)

		err = scanner.ProcessAssertionCreation(ctx, createdData.Leaf1.Id())
		require.NoError(t, err)

		otherManager, err := challengemanager.New(
			ctx,
			createdData.Chains[0],
			createdData.Backend,
			createdData.EvilStateManager,
			createdData.Addrs.Rollup,
			challengemanager.WithMode(types.MakeMode),
		)
		require.NoError(t, err)

		otherScanner := assertions.NewScanner(createdData.Chains[0], createdData.EvilStateManager, createdData.Backend, otherManager, createdData.Addrs.Rollup, "", time.Second)

		err = otherScanner.ProcessAssertionCreation(ctx, createdData.Leaf2.Id())
		require.NoError(t, err)

		err = otherScanner.ProcessAssertionCreation(ctx, createdData.Leaf2.Id())
		require.NoError(t, err)

		require.Equal(t, uint64(2), otherScanner.AssertionsProcessed())
		require.Equal(t, uint64(2), otherScanner.ForksDetected())
		require.Equal(t, uint64(2), otherScanner.ChallengesSubmitted())
	})
	t.Run("defensive validator can still challenge leaf", func(t *testing.T) {
		ctx := context.Background()
		createdData, err := setup.CreateTwoValidatorFork(ctx, &setup.CreateForkConfig{
			DivergeBlockHeight: 5,
		})
		require.NoError(t, err)

		manager, err := challengemanager.New(
			ctx,
			createdData.Chains[1],
			createdData.Backend,
			createdData.HonestStateManager,
			createdData.Addrs.Rollup,
			challengemanager.WithMode(types.DefensiveMode),
		)
		require.NoError(t, err)
		scanner := assertions.NewScanner(createdData.Chains[1], createdData.HonestStateManager, createdData.Backend, manager, createdData.Addrs.Rollup, "", time.Second)

		err = scanner.ProcessAssertionCreation(ctx, createdData.Leaf1.Id())
		require.NoError(t, err)

		otherManager, err := challengemanager.New(
			ctx,
			createdData.Chains[0],
			createdData.Backend,
			createdData.EvilStateManager,
			createdData.Addrs.Rollup,
			challengemanager.WithMode(types.DefensiveMode),
		)
		require.NoError(t, err)

		otherScanner := assertions.NewScanner(createdData.Chains[0], createdData.EvilStateManager, createdData.Backend, otherManager, createdData.Addrs.Rollup, "", time.Second)

		err = otherScanner.ProcessAssertionCreation(ctx, createdData.Leaf2.Id())
		require.NoError(t, err)

		err = otherScanner.ProcessAssertionCreation(ctx, createdData.Leaf2.Id())
		require.NoError(t, err)

		require.Equal(t, uint64(2), otherScanner.AssertionsProcessed())
		require.Equal(t, uint64(2), otherScanner.ForksDetected())
		require.Equal(t, uint64(2), otherScanner.ChallengesSubmitted())
	})
}

func setupChallengeManager(t *testing.T) (*challengemanager.Manager, *mocks.MockProtocol, *mocks.MockStateManager, *setup.ChainSetup) {
	t.Helper()
	p := &mocks.MockProtocol{}
	ctx := context.Background()
	p.On("CurrentChallengeManager", ctx).Return(&mocks.MockChallengeManager{}, nil)
	p.On("SpecChallengeManager", ctx).Return(&mocks.MockSpecChallengeManager{}, nil)
	s := &mocks.MockStateManager{}
	cfg, err := setup.ChainsWithEdgeChallengeManager()
	require.NoError(t, err)
	v, err := challengemanager.New(context.Background(), p, cfg.Backend, s, cfg.Addrs.Rollup, challengemanager.WithMode(types.MakeMode))
	require.NoError(t, err)
	return v, p, s, cfg
}

func mockId(x uint64) protocol.AssertionHash {
	return protocol.AssertionHash(common.BytesToHash([]byte(fmt.Sprintf("%d", x))))
}
