// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE

package assertions_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/OffchainLabs/bold/assertions"
	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	challengemanager "github.com/OffchainLabs/bold/challenge-manager"
	"github.com/OffchainLabs/bold/challenge-manager/types"
	"github.com/OffchainLabs/bold/solgen/go/rollupgen"
	"github.com/OffchainLabs/bold/testing/mocks"
	"github.com/OffchainLabs/bold/testing/setup"
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
		cm := &mocks.MockSpecChallengeManager{}
		p.On("SpecChallengeManager", ctx).Return(cm, nil)
		p.On("ReadAssertionCreationInfo", ctx, mockId(2)).Return(&protocol.AssertionCreatedInfo{
			ParentAssertionHash: mockId(1).Hash,
			AfterState:          rollupgen.ExecutionState{},
		}, nil)
		p.On("GetAssertion", ctx, mockId(2)).Return(ev, nil)
		p.On("GetAssertion", ctx, mockId(1)).Return(prev, nil)
		mockStateProvider.On("AgreesWithExecutionState", ctx, &protocol.ExecutionState{}).Return(nil)
		scanner, err := assertions.NewManager(p, mockStateProvider, cfg.Backend, manager, cfg.Addrs.Rollup, "", time.Second, time.Second, &mocks.MockStateManager{}, time.Second, time.Second)
		require.NoError(t, err)

		err = scanner.ProcessAssertionCreationEvent(ctx, ev.Id())
		require.NoError(t, err)
		require.Equal(t, uint64(1), scanner.AssertionsProcessed())
		require.Equal(t, uint64(0), scanner.ForksDetected())
		require.Equal(t, uint64(0), scanner.ChallengesSubmitted())
	})
	t.Run("fork leads validator to challenge leaf", func(t *testing.T) {
		ctx := context.Background()
		createdData, err := setup.CreateTwoValidatorFork(ctx, &setup.CreateForkConfig{
			DivergeBlockHeight: 5,
		}, setup.WithMockOneStepProver())
		require.NoError(t, err)

		manager, err := challengemanager.New(
			ctx,
			createdData.Chains[1],
			createdData.Backend,
			createdData.HonestStateManager,
			createdData.Addrs.Rollup,
			challengemanager.WithMode(types.MakeMode),
			challengemanager.WithEdgeTrackerWakeInterval(100*time.Millisecond),
		)
		require.NoError(t, err)

		scanner, err := assertions.NewManager(createdData.Chains[1], createdData.HonestStateManager, createdData.Backend, manager, createdData.Addrs.Rollup, "", time.Second, time.Second, createdData.HonestStateManager, time.Second, time.Second)
		require.NoError(t, err)

		err = scanner.ProcessAssertionCreationEvent(ctx, createdData.Leaf2.Id())
		require.NoError(t, err)

		otherManager, err := challengemanager.New(
			ctx,
			createdData.Chains[0],
			createdData.Backend,
			createdData.EvilStateManager,
			createdData.Addrs.Rollup,
			challengemanager.WithMode(types.MakeMode),
			challengemanager.WithEdgeTrackerWakeInterval(100*time.Millisecond),
		)
		require.NoError(t, err)

		otherScanner, err := assertions.NewManager(createdData.Chains[0], createdData.EvilStateManager, createdData.Backend, otherManager, createdData.Addrs.Rollup, "", time.Second, time.Second, createdData.EvilStateManager, time.Second, time.Second)
		require.NoError(t, err)

		err = otherScanner.ProcessAssertionCreationEvent(ctx, createdData.Leaf1.Id())
		require.NoError(t, err)

		require.Equal(t, uint64(1), otherScanner.AssertionsProcessed())
		require.Equal(t, uint64(1), otherScanner.ChallengesSubmitted())
		require.Equal(t, uint64(1), scanner.AssertionsProcessed())
		require.Equal(t, uint64(1), scanner.ChallengesSubmitted())
	})
	t.Run("defensive validator can still challenge leaf", func(t *testing.T) {
		ctx := context.Background()
		createdData, err := setup.CreateTwoValidatorFork(ctx, &setup.CreateForkConfig{
			DivergeBlockHeight: 5,
		}, setup.WithMockOneStepProver())
		require.NoError(t, err)

		manager, err := challengemanager.New(
			ctx,
			createdData.Chains[1],
			createdData.Backend,
			createdData.HonestStateManager,
			createdData.Addrs.Rollup,
			challengemanager.WithMode(types.DefensiveMode),
			challengemanager.WithEdgeTrackerWakeInterval(100*time.Millisecond),
		)
		require.NoError(t, err)
		scanner, err := assertions.NewManager(createdData.Chains[1], createdData.HonestStateManager, createdData.Backend, manager, createdData.Addrs.Rollup, "", time.Second, time.Second, createdData.HonestStateManager, time.Second, time.Second)
		require.NoError(t, err)

		err = scanner.ProcessAssertionCreationEvent(ctx, createdData.Leaf2.Id())
		require.NoError(t, err)

		otherManager, err := challengemanager.New(
			ctx,
			createdData.Chains[0],
			createdData.Backend,
			createdData.EvilStateManager,
			createdData.Addrs.Rollup,
			challengemanager.WithMode(types.DefensiveMode),
			challengemanager.WithEdgeTrackerWakeInterval(100*time.Millisecond),
		)
		require.NoError(t, err)

		otherScanner, err := assertions.NewManager(createdData.Chains[0], createdData.EvilStateManager, createdData.Backend, otherManager, createdData.Addrs.Rollup, "", time.Second, time.Second, createdData.EvilStateManager, time.Second, time.Second)
		require.NoError(t, err)

		err = otherScanner.ProcessAssertionCreationEvent(ctx, createdData.Leaf1.Id())
		require.NoError(t, err)

		require.Equal(t, uint64(1), otherScanner.AssertionsProcessed())
		require.Equal(t, uint64(1), otherScanner.ChallengesSubmitted())
		require.Equal(t, uint64(1), scanner.AssertionsProcessed())
		require.Equal(t, uint64(1), scanner.ChallengesSubmitted())
	})
}

func setupChallengeManager(t *testing.T) (*challengemanager.Manager, *mocks.MockProtocol, *mocks.MockStateManager, *setup.ChainSetup) {
	t.Helper()
	p := &mocks.MockProtocol{}
	ctx := context.Background()
	cm := &mocks.MockSpecChallengeManager{}
	cm.On("NumBigSteps", ctx).Return(uint8(1), nil)
	p.On("CurrentChallengeManager", ctx).Return(cm, nil)
	p.On("SpecChallengeManager", ctx).Return(cm, nil)
	s := &mocks.MockStateManager{}
	cfg, err := setup.ChainsWithEdgeChallengeManager(setup.WithMockOneStepProver())
	require.NoError(t, err)
	v, err := challengemanager.New(context.Background(), p, cfg.Backend, s, cfg.Addrs.Rollup, challengemanager.WithMode(types.MakeMode), challengemanager.WithEdgeTrackerWakeInterval(100*time.Millisecond))
	require.NoError(t, err)
	return v, p, s, cfg
}

func mockId(x uint64) protocol.AssertionHash {
	return protocol.AssertionHash{Hash: common.BytesToHash([]byte(fmt.Sprintf("%d", x)))}
}
