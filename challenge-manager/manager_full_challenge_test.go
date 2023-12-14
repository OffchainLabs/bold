package challengemanager

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/OffchainLabs/bold/assertions"
	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	solimpl "github.com/OffchainLabs/bold/chain-abstraction/sol-implementation"
	"github.com/OffchainLabs/bold/challenge-manager/types"
	l2stateprovider "github.com/OffchainLabs/bold/layer2-state-provider"
	retry "github.com/OffchainLabs/bold/runtime"
	"github.com/OffchainLabs/bold/solgen/go/mocksgen"
	"github.com/OffchainLabs/bold/solgen/go/rollupgen"
	challenge_testing "github.com/OffchainLabs/bold/testing"
	statemanager "github.com/OffchainLabs/bold/testing/mocks/state-provider"
	"github.com/OffchainLabs/bold/testing/setup"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

func TestFullChallenge_IntegrationTest(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	setup, err := setup.ChainsWithEdgeChallengeManager(
		setup.WithMockBridge(),
		setup.WithMockOneStepProver(),
		setup.WithChallengeTestingOpts(
			challenge_testing.WithConfirmPeriodBlocks(25),
			challenge_testing.WithLayerZeroHeights(&protocol.LayerZeroHeights{
				BlockChallengeHeight:     64,
				BigStepChallengeHeight:   32,
				SmallStepChallengeHeight: 32,
			}),
		),
	)
	require.NoError(t, err)

	bridgeBindings, err := mocksgen.NewBridgeStub(setup.Addrs.Bridge, setup.Backend)
	require.NoError(t, err)

	rollupAdminBindings, err := rollupgen.NewRollupAdminLogic(setup.Addrs.Rollup, setup.Backend)
	require.NoError(t, err)
	_, err = rollupAdminBindings.SetMinimumAssertionPeriod(setup.Accounts[0].TxOpts, big.NewInt(1))
	require.NoError(t, err)
	setup.Backend.Commit()

	msgCount, err := bridgeBindings.SequencerMessageCount(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, uint64(1), msgCount.Uint64())

	genesisHash, err := setup.Chains[1].GenesisAssertionHash(ctx)
	require.NoError(t, err)
	genesisCreationInfo, err := setup.Chains[1].ReadAssertionCreationInfo(ctx, protocol.AssertionHash{Hash: genesisHash})
	require.NoError(t, err)
	_ = genesisCreationInfo

	stateManagerOpts := []statemanager.Opt{
		statemanager.WithNumBatchesRead(5),
	}
	honestStateManager, err := statemanager.NewForSimpleMachine(stateManagerOpts...)
	require.NoError(t, err)

	// Bob diverges from Alice at batch 1.
	// assertionDivergenceHeight := uint64(4)
	// assertionBlockHeightDifference := int64(4)
	stateManagerOpts = append(
		stateManagerOpts,
		statemanager.WithBlockDivergenceHeight(1),
		statemanager.WithMachineDivergenceStep(1),
	// 	statemanager.WithMachineDivergenceStep(machineDivergenceStep),
	// 	statemanager.WithBlockDivergenceHeight(assertionDivergenceHeight),
	// 	statemanager.WithDivergentBlockHeightOffset(assertionBlockHeightDifference),
	)
	evilStateManager, err := statemanager.NewForSimpleMachine(stateManagerOpts...)
	require.NoError(t, err)

	honestPostState, err := honestStateManager.ExecutionStateAfterBatchCount(ctx, 1)
	require.NoError(t, err)

	t.Logf("New stake from alice at post state %+v\n", honestPostState)
	honestManager, honestChain := setupChallengeManager(
		t, ctx, setup.Backend, setup.Addrs.Rollup, honestStateManager, setup.Accounts[1].TxOpts, "honest",
	)
	evilManager, evilChain := setupChallengeManager(
		t, ctx, setup.Backend, setup.Addrs.Rollup, evilStateManager, setup.Accounts[2].TxOpts, "evil",
	)

	honestPoster, err := assertions.NewManager(
		honestChain,
		honestStateManager,
		setup.Backend,
		honestManager,
		setup.Addrs.Rollup,
		"honest",
		time.Second,
		time.Second,
		honestStateManager,
		time.Second,
		time.Millisecond,
	)
	require.NoError(t, err)

	evilPoster, err := assertions.NewManager(
		evilChain,
		evilStateManager,
		setup.Backend,
		evilManager,
		setup.Addrs.Rollup,
		"evil",
		time.Second,
		time.Second,
		evilStateManager,
		time.Second,
		time.Millisecond,
	)
	require.NoError(t, err)

	honestAssertion, err := honestPoster.PostAssertion(ctx)
	require.NoError(t, err)

	evilAssertion, err := evilPoster.PostAssertion(ctx)
	require.NoError(t, err)
	_ = honestAssertion
	_ = evilAssertion

	err = honestPoster.ProcessAssertionCreationEvent(ctx, honestAssertion.Unwrap().Id())
	require.NoError(t, err)
	err = evilPoster.ProcessAssertionCreationEvent(ctx, evilAssertion.Unwrap().Id())
	require.NoError(t, err)
	// _ = evilAssertion

	honestManager.Start(ctx)
	evilManager.Start(ctx)

	//time.Sleep(time.Minute * 10)

	// Advance the blockchain in the background.
	blockTime := time.Millisecond * 100
	go func() {
		ticker := time.NewTicker(blockTime)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				setup.Backend.Commit()
			case <-ctx.Done():
				return
			}
		}
	}()

	expectations := []expect{
		expectAssertionConfirmedByChallengeWin,
	}
	g, ctx := errgroup.WithContext(ctx)
	for _, e := range expectations {
		fn := e // loop closure
		g.Go(func() error {
			return fn(t, ctx, setup.Addrs, setup.Backend)
		})
	}

	if err := g.Wait(); err != nil {
		t.Fatal(err)
	}

	// trackedBackend, ok := aChain.Backend().(*solimpl.TrackedContractBackend)
	// if !ok {
	// 	t.Fatal("Not a tracked contract backend")
	// }
	// t.Log("Printing Alice's ethclient metrics at the end of a challenge")
	// trackedBackend.PrintMetrics()
}

// setupValidator initializes a validator with the minimum required configuration.
func setupChallengeManager(
	t *testing.T,
	ctx context.Context,
	backend protocol.ChainBackend,
	rollup common.Address,
	sm l2stateprovider.Provider,
	txOpts *bind.TransactOpts,
	name string,
) (*Manager, protocol.Protocol) {
	chain, err := solimpl.NewAssertionChain(
		ctx,
		rollup,
		txOpts,
		backend,
		//solimpl.WithTrackedContractBackend(),
	)
	require.NoError(t, err)

	v, err := New(
		ctx,
		chain,
		backend,
		sm,
		rollup,
		WithAddress(txOpts.From),
		WithName(name),
		WithEdgeTrackerWakeInterval(time.Millisecond*250),
		WithMode(types.MakeMode),
	)
	require.NoError(t, err)
	return v, chain
}

func totalWasmOpcodes(heights *protocol.LayerZeroHeights, numBigSteps uint8) uint64 {
	totalWasmOpcodes := uint64(1)
	for i := uint8(0); i < numBigSteps; i++ {
		totalWasmOpcodes *= heights.BigStepChallengeHeight
	}
	totalWasmOpcodes *= heights.SmallStepChallengeHeight
	return totalWasmOpcodes
}

// expect is a function that will be called asynchronously to verify some success criteria
// for the given scenario.
type expect func(t *testing.T, ctx context.Context, addresses *setup.RollupAddresses, be protocol.ChainBackend) error

// Expects that an assertion is confirmed by challenge win.
func expectAssertionConfirmedByChallengeWin(
	t *testing.T,
	ctx context.Context,
	addresses *setup.RollupAddresses,
	backend protocol.ChainBackend,
) error {
	t.Run("assertion confirmed by challenge win", func(t *testing.T) {
		rc, err := rollupgen.NewRollupCore(addresses.Rollup, backend)
		require.NoError(t, err)

		var confirmed bool
		for ctx.Err() == nil && !confirmed {
			i, err := retry.UntilSucceeds(ctx, func() (*rollupgen.RollupCoreAssertionConfirmedIterator, error) {
				return rc.FilterAssertionConfirmed(nil, nil)
			})
			require.NoError(t, err)
			for i.Next() {
				assertionNode, err := retry.UntilSucceeds(ctx, func() (rollupgen.AssertionNode, error) {
					return rc.GetAssertion(&bind.CallOpts{Context: ctx}, i.Event.AssertionHash)
				})
				require.NoError(t, err)
				if assertionNode.Status != uint8(protocol.AssertionConfirmed) {
					t.Fatal("Confirmed assertion with unfinished state")
				}
				confirmed = true
				break
			}
			time.Sleep(500 * time.Millisecond) // Don't spam the backend.
		}

		if !confirmed {
			t.Fatal("assertion was not confirmed")
		}
	})
	return nil
}

// // expectAliceAndBobStaked monitors EdgeAdded events until Alice and Bob are observed adding edges
// // with a stake.
// func expectAliceAndBobStaked(t *testing.T, ctx context.Context, be backend.Backend) error {
// 	t.Run("alice and bob staked", func(t *testing.T) {
// 		ecm, err := retry.UntilSucceeds(ctx, func() (*challengeV2gen.EdgeChallengeManager, error) {
// 			return edgeManager(be)
// 		})
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		var aliceStaked, bobStaked bool
// 		for ctx.Err() == nil && (!aliceStaked || !bobStaked) {
// 			i, err := retry.UntilSucceeds(ctx, func() (*challengeV2gen.EdgeChallengeManagerEdgeAddedIterator, error) {
// 				return ecm.FilterEdgeAdded(nil, nil, nil, nil)
// 			})
// 			if err != nil {
// 				t.Fatal(err)
// 			}
// 			for i.Next() {
// 				e, err := retry.UntilSucceeds(ctx, func() (challengeV2gen.ChallengeEdge, error) {
// 					return ecm.GetEdge(nil, i.Event.EdgeId)
// 				})
// 				if err != nil {
// 					t.Fatal(err)
// 				}

// 				switch e.Staker {
// 				case be.Alice().From:
// 					aliceStaked = true
// 				case be.Bob().From:
// 					bobStaked = true
// 				}

// 				time.Sleep(500 * time.Millisecond) // Don't spam the backend.
// 			}
// 		}

// 		if !aliceStaked {
// 			t.Error("alice did not stake")
// 		}
// 		if !bobStaked {
// 			t.Error("bob did not stake")
// 		}
// 	})

// 	return nil
// }
