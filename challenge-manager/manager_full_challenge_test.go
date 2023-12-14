package challengemanager

import (
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"testing"
	"time"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	solimpl "github.com/OffchainLabs/bold/chain-abstraction/sol-implementation"
	"github.com/OffchainLabs/bold/challenge-manager/types"
	l2stateprovider "github.com/OffchainLabs/bold/layer2-state-provider"
	retry "github.com/OffchainLabs/bold/runtime"
	"github.com/OffchainLabs/bold/solgen/go/rollupgen"
	challenge_testing "github.com/OffchainLabs/bold/testing"
	statemanager "github.com/OffchainLabs/bold/testing/mocks/state-provider"
	"github.com/OffchainLabs/bold/testing/setup"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

// TODO: Support concurrent challenges at the assertion chain level.
// TODO: Support flooding attack
// TODO: Support concurrent assertion level challenges + flooding attack
// TODO: Many evil parties, each with their own claim.
// TODO: Many evil parties, all supporting the same claim.
type integrationTestConfig struct {
	protocol     protocolParams
	timings      timeParams
	inbox        inboxParams
	actors       actorParams
	expectations []expect
}

type actorParams struct {
	numEvilValidators uint64
}

type timeParams struct {
	blockTime                            time.Duration
	challengeMoveInterval                time.Duration
	assertionPostingInterval             time.Duration
	assertionScanningInterval            time.Duration
	assertionConfirmationAttemptInterval time.Duration
}

type inboxParams struct {
	numBatchesPosted uint64
}

type protocolParams struct {
	numBigStepLevels      uint8
	challengePeriodBlocks uint64
	layerZeroHeights      *protocol.LayerZeroHeights
}

func TestChallenge_IntegrationTest_Fast(t *testing.T) {
	runChallengeIntegrationTest(t, &integrationTestConfig{
		protocol: protocolParams{
			challengePeriodBlocks: 150,
			numBigStepLevels:      1,
			layerZeroHeights: &protocol.LayerZeroHeights{
				BlockChallengeHeight:     1 << 6,
				BigStepChallengeHeight:   1 << 5,
				SmallStepChallengeHeight: 1 << 5,
			},
		},
		inbox: inboxParams{
			// Assume 5 batches have been posted to the inbox contract.
			numBatchesPosted: 5,
		},
		actors: actorParams{
			numEvilValidators: 1,
		},
		timings: timeParams{
			// Fast block time.
			blockTime: time.Millisecond * 100,
			// Go as fast as possible.
			challengeMoveInterval:     time.Millisecond,
			assertionPostingInterval:  time.Second,
			assertionScanningInterval: time.Second,
			// An extremely high number so that we never attempt to confirm an assertion by time.
			// We instead expect the assertion to be confirmed by challenge win.
			assertionConfirmationAttemptInterval: time.Hour,
		},
		expectations: []expect{
			// Expect one assertion is confirmed by challenge win.
			expectAssertionConfirmedByChallengeWin,
			// Other ideas:
			// All validators are staked at top-level
			// All subchallenges have mini-stakes
		},
	})
}

func TestChallenge_IntegrationTest_MaxWavmOpcodes(t *testing.T) {
	runChallengeIntegrationTest(t, &integrationTestConfig{
		protocol: protocolParams{
			challengePeriodBlocks: 200,
			// A block can take a max of 2^43 wavm opcodes to validate.
			// With three big step levels of 2^10 each, and one small step level of
			// 2^13 (adding the exponents), we can have full coverage of a block dispute
			// in a challenge game.
			numBigStepLevels: 3,
			layerZeroHeights: &protocol.LayerZeroHeights{
				BlockChallengeHeight:     1 << 6,
				BigStepChallengeHeight:   1 << 10,
				SmallStepChallengeHeight: 1 << 13,
			},
		},
		inbox: inboxParams{
			// Assume 1 batch has been posted to the inbox contract.
			numBatchesPosted: 5,
		},
		actors: actorParams{
			numEvilValidators: 1,
		},
		timings: timeParams{
			// Fast block time.
			blockTime: time.Millisecond * 100,
			// Fast challenge move time.
			challengeMoveInterval:     time.Millisecond,
			assertionPostingInterval:  time.Second,
			assertionScanningInterval: time.Second,
			// An extremely high number so that we never attempt to confirm an assertion by time.
			// We instead expect the assertion to be confirmed by challenge win.
			assertionConfirmationAttemptInterval: time.Hour,
		},
		expectations: []expect{
			// Expect one assertion is confirmed by challenge win.
			expectAssertionConfirmedByChallengeWin,
		},
	})
}

func TestChallenge_IntegrationTest_MultipleValidators(t *testing.T) {
	runChallengeIntegrationTest(t, &integrationTestConfig{
		protocol: protocolParams{
			challengePeriodBlocks: 100,
			numBigStepLevels:      1,
			layerZeroHeights: &protocol.LayerZeroHeights{
				BlockChallengeHeight:     1 << 6,
				BigStepChallengeHeight:   1 << 5,
				SmallStepChallengeHeight: 1 << 5,
			},
		},
		inbox: inboxParams{
			// Assume 1 batch has been posted to the inbox contract.
			numBatchesPosted: 5,
		},
		actors: actorParams{
			numEvilValidators: 2,
		},
		timings: timeParams{
			// Fast block time.
			blockTime: time.Second,
			// Challenge move time.
			challengeMoveInterval:     time.Millisecond * 250,
			assertionPostingInterval:  time.Second,
			assertionScanningInterval: time.Second,
			// An extremely high number so that we never attempt to confirm an assertion by time.
			// We instead expect the assertion to be confirmed by challenge win.
			assertionConfirmationAttemptInterval: time.Hour,
		},
		expectations: []expect{
			// Expect one assertion is confirmed by challenge win.
			expectAssertionConfirmedByChallengeWin,
		},
	})
}

func runChallengeIntegrationTest(t *testing.T, cfg *integrationTestConfig) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Validators include a chain admin, a single honest validators, and any number of evil entities.
	totalValidators := cfg.actors.numEvilValidators + 2
	setup, err := setup.ChainsWithEdgeChallengeManager(
		setup.WithMockBridge(),
		setup.WithMockOneStepProver(),
		setup.WithNumAccounts(totalValidators),
		setup.WithChallengeTestingOpts(
			challenge_testing.WithConfirmPeriodBlocks(cfg.protocol.challengePeriodBlocks),
			challenge_testing.WithLayerZeroHeights(cfg.protocol.layerZeroHeights),
			challenge_testing.WithNumBigStepLevels(cfg.protocol.numBigStepLevels),
		),
	)
	require.NoError(t, err)

	rollupAdminBindings, err := rollupgen.NewRollupAdminLogic(setup.Addrs.Rollup, setup.Backend)
	require.NoError(t, err)
	_, err = rollupAdminBindings.SetMinimumAssertionPeriod(setup.Accounts[0].TxOpts, big.NewInt(1))
	require.NoError(t, err)
	setup.Backend.Commit()

	baseStateManagerOpts := []statemanager.Opt{
		statemanager.WithNumBatchesRead(cfg.inbox.numBatchesPosted),
		statemanager.WithLayerZeroHeights(cfg.protocol.layerZeroHeights, cfg.protocol.numBigStepLevels),
	}
	honestStateManager, err := statemanager.NewForSimpleMachine(baseStateManagerOpts...)
	require.NoError(t, err)

	baseChallengeManagerOpts := []Opt{
		WithEdgeTrackerWakeInterval(cfg.timings.challengeMoveInterval),
		WithMode(types.MakeMode),
		WithAssertionPostingInterval(cfg.timings.assertionPostingInterval),
		WithAssertionScanningInterval(cfg.timings.assertionScanningInterval),
		WithAssertionConfirmingInterval(cfg.timings.assertionConfirmationAttemptInterval),
	}

	name := "honest"
	txOpts := setup.Accounts[1].TxOpts
	honestOpts := append(
		baseChallengeManagerOpts,
		WithAddress(txOpts.From),
		WithName(name),
	)
	honestManager := setupChallengeManager(
		t, ctx, setup.Backend, setup.Addrs.Rollup, honestStateManager, txOpts, name, honestOpts...,
	)

	// Diverge exactly at the last opcode within the block.
	totalOpcodes := totalWasmOpcodes(cfg.protocol.layerZeroHeights, cfg.protocol.numBigStepLevels)
	t.Logf("Total wasm opcodes in test: %d", totalOpcodes)

	assertionDivergenceHeight := uint64(1)
	assertionBlockHeightDifference := int64(1)

	evilChallengeManagers := make([]*Manager, cfg.actors.numEvilValidators)
	for i := uint64(0); i < cfg.actors.numEvilValidators; i++ {
		machineDivergenceStep := randUint64(totalOpcodes)
		evilStateManagerOpts := append(
			baseStateManagerOpts,
			statemanager.WithMachineDivergenceStep(machineDivergenceStep),
			statemanager.WithBlockDivergenceHeight(assertionDivergenceHeight),
			statemanager.WithDivergentBlockHeightOffset(assertionBlockHeightDifference),
		)
		evilStateManager, err := statemanager.NewForSimpleMachine(evilStateManagerOpts...)
		require.NoError(t, err)

		// Honest validator has index 1 in the accounts slice, as 0 is admin, so evil ones should start at 2.
		txOpts = setup.Accounts[2+i].TxOpts
		name = fmt.Sprintf("evil-%d", i)
		evilOpts := append(
			baseChallengeManagerOpts,
			WithAddress(txOpts.From),
			WithName(name),
		)
		evilManager := setupChallengeManager(
			t, ctx, setup.Backend, setup.Addrs.Rollup, evilStateManager, txOpts, name, evilOpts...,
		)
		evilChallengeManagers[i] = evilManager
	}

	honestManager.Start(ctx)

	for _, evilManager := range evilChallengeManagers {
		evilManager.Start(ctx)
	}

	// Advance the blockchain in the background.
	go func() {
		ticker := time.NewTicker(cfg.timings.blockTime)
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

	g, ctx := errgroup.WithContext(ctx)
	for _, e := range cfg.expectations {
		fn := e // loop closure
		g.Go(func() error {
			return fn(t, ctx, setup.Addrs, setup.Backend)
		})
	}
	require.NoError(t, g.Wait())
}

func setupChallengeManager(
	t *testing.T,
	ctx context.Context,
	backend protocol.ChainBackend,
	rollup common.Address,
	sm l2stateprovider.Provider,
	txOpts *bind.TransactOpts,
	name string,
	opts ...Opt,
) *Manager {
	chain, err := solimpl.NewAssertionChain(
		ctx,
		rollup,
		txOpts,
		backend,
	)
	require.NoError(t, err)

	manager, err := New(
		ctx,
		chain,
		backend,
		sm,
		rollup,
		opts...,
	)
	require.NoError(t, err)
	return manager
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

// rand.Uint64() returns a random uint64 value.
// To get a value in the range [0, n), take the modulo n.
func randUint64(n uint64) uint64 {
	if n == 0 {
		return 0
	}
	return rand.Uint64() % n
}
