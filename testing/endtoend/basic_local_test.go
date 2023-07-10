package endtoend

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/OffchainLabs/challenge-protocol-v2/assertions"
	protocol "github.com/OffchainLabs/challenge-protocol-v2/chain-abstraction"
	solimpl "github.com/OffchainLabs/challenge-protocol-v2/chain-abstraction/sol-implementation"
	validator "github.com/OffchainLabs/challenge-protocol-v2/challenge-manager"
	"github.com/OffchainLabs/challenge-protocol-v2/challenge-manager/types"
	l2stateprovider "github.com/OffchainLabs/challenge-protocol-v2/layer2-state-provider"
	challenge_testing "github.com/OffchainLabs/challenge-protocol-v2/testing"
	"github.com/OffchainLabs/challenge-protocol-v2/testing/endtoend/internal/backend"
	statemanager "github.com/OffchainLabs/challenge-protocol-v2/testing/toys/state-provider"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

type ChallengeScenario struct {
	Name string

	// Validator knowledge
	AliceStateManager   l2stateprovider.Provider
	BobStateManager     l2stateprovider.Provider
	CharlieStateManager l2stateprovider.Provider

	// Expectations
	Expectations []expect
}

type challengeProtocolTestConfig struct {
	// The height in the assertion chain at which the validators diverge.
	assertionDivergenceHeight uint64
	// The difference between the malicious assertion block height and the honest assertion block height.
	assertionBlockHeightDifference int64
	// The heights at which the validators diverge in histories at the big step
	// subchallenge level.
	bigStepDivergenceHeight uint64
	// The heights at which the validators diverge in histories at the small step
	// subchallenge level.
	smallStepDivergenceHeight uint64
}

func TestChallengeProtocol_AliceAndBob_AnvilLocal_SameHeight(t *testing.T) {
	be, err := backend.NewAnvilLocal(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if err := be.Start(); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := be.Stop(); err != nil {
			t.Log(fmt.Errorf("error stopping backend: %v", err))
		}
	}()

	scenario := &ChallengeScenario{
		Name: "two forked assertions at the same height",
		AliceStateManager: func() l2stateprovider.Provider {
			sm, err := statemanager.NewForSimpleMachine()
			if err != nil {
				t.Fatal(err)
			}
			return sm
		}(),
		BobStateManager: func() l2stateprovider.Provider {
			cfg := &challengeProtocolTestConfig{
				// The heights at which the validators diverge in histories. In this test,
				// alice and bob start diverging at height 3 at all subchallenge levels.
				assertionDivergenceHeight: 4,
				bigStepDivergenceHeight:   4,
				smallStepDivergenceHeight: 4,
			}
			sm, err := statemanager.NewForSimpleMachine(
				statemanager.WithMachineDivergenceStep(cfg.bigStepDivergenceHeight*challenge_testing.LevelZeroSmallStepEdgeHeight+cfg.smallStepDivergenceHeight),
				statemanager.WithBlockDivergenceHeight(cfg.assertionDivergenceHeight),
				statemanager.WithDivergentBlockHeightOffset(cfg.assertionBlockHeightDifference),
			)
			if err != nil {
				t.Fatal(err)
			}
			return sm
		}(),
		Expectations: []expect{
			expectAssertionConfirmedByChallengeWinner,
			expectAliceAndBobStaked,
		},
	}

	testChallengeProtocol_AliceAndBob(t, be, scenario)
}

func TestChallengeProtocol_AliceAndBob_AnvilLocal_DifferentHeights(t *testing.T) {
	t.Skip()
	be, err := backend.NewAnvilLocal(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if err := be.Start(); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := be.Stop(); err != nil {
			t.Log(fmt.Errorf("error stopping backend: %v", err))
		}
	}()

	scenario := &ChallengeScenario{
		Name: "two forked assertions at the different step heights",
		AliceStateManager: func() l2stateprovider.Provider {
			sm, err := statemanager.NewForSimpleMachine()
			if err != nil {
				t.Fatal(err)
			}
			return sm
		}(),
		BobStateManager: func() l2stateprovider.Provider {
			cfg := &challengeProtocolTestConfig{
				// The heights at which the validators diverge in histories. In this test,
				// alice and bob diverge heights at different subchallenge levels.
				assertionDivergenceHeight: 8,
				bigStepDivergenceHeight:   6,
				smallStepDivergenceHeight: 4,
			}
			sm, err := statemanager.NewForSimpleMachine(
				statemanager.WithMachineDivergenceStep(cfg.bigStepDivergenceHeight*challenge_testing.LevelZeroSmallStepEdgeHeight+cfg.smallStepDivergenceHeight),
				statemanager.WithBlockDivergenceHeight(cfg.assertionDivergenceHeight),
				statemanager.WithDivergentBlockHeightOffset(cfg.assertionBlockHeightDifference),
			)
			if err != nil {
				t.Fatal(err)
			}
			return sm
		}(),
		Expectations: []expect{
			expectAssertionConfirmedByChallengeWinner,
			expectAliceAndBobStaked,
		},
	}
	testChallengeProtocol_AliceAndBob(t, be, scenario)
}

func TestSync_HonestBobStopsCharlieJoins(t *testing.T) {
	be, err := backend.NewAnvilLocal(context.Background())
	require.NoError(t, err)
	require.NoError(t, be.Start())
	defer func() {
		require.NoError(t, be.Stop(), "error stopping backend")
	}()

	scenario := &ChallengeScenario{
		Name: "two forked assertions at the same height",
		AliceStateManager: func() l2stateprovider.Provider {
			cfg := &challengeProtocolTestConfig{
				// The heights at which the validators diverge in histories. In this test,
				// alice and bob start diverging at height 3 at all subchallenge levels.
				assertionDivergenceHeight: 4,
				bigStepDivergenceHeight:   4,
				smallStepDivergenceHeight: 4,
			}
			sm, err := statemanager.NewForSimpleMachine(
				statemanager.WithMachineDivergenceStep(cfg.bigStepDivergenceHeight*challenge_testing.LevelZeroSmallStepEdgeHeight+cfg.smallStepDivergenceHeight),
				statemanager.WithBlockDivergenceHeight(cfg.assertionDivergenceHeight),
				statemanager.WithDivergentBlockHeightOffset(cfg.assertionBlockHeightDifference),
			)
			if err != nil {
				t.Fatal(err)
			}
			return sm
		}(),
		BobStateManager: func() l2stateprovider.Provider {
			sm, err := statemanager.NewForSimpleMachine()
			if err != nil {
				t.Fatal(err)
			}
			return sm
		}(),
		CharlieStateManager: func() l2stateprovider.Provider {
			sm, err := statemanager.NewForSimpleMachine()
			if err != nil {
				t.Fatal(err)
			}
			return sm
		}(),
		Expectations: []expect{
			expectAssertionConfirmedByChallengeWinner,
		},
	}

	testSyncBobStopsCharlieJoins(t, be, scenario)
}

func testChallengeProtocol_AliceAndBob(t *testing.T, be backend.Backend, scenario *ChallengeScenario) {
	t.Run(scenario.Name, func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
		defer cancel()

		rollup, err := be.DeployRollup()
		if err != nil {
			t.Fatal(err)
		}

		a, aChain, err := setupValidator(ctx, be, rollup, scenario.AliceStateManager, be.Alice(), "alice")
		if err != nil {
			t.Fatal(err)
		}
		b, bChain, err := setupValidator(ctx, be, rollup, scenario.BobStateManager, be.Bob(), "bob")
		if err != nil {
			t.Fatal(err)
		}

		// Post assertions.
		alicePoster := assertions.NewPoster(aChain, scenario.AliceStateManager, "alice", time.Hour)
		bobPoster := assertions.NewPoster(bChain, scenario.BobStateManager, "bob", time.Hour)

		aliceLeaf, err := alicePoster.PostLatestAssertion(ctx)
		if err != nil {
			t.Fatal(err)
		}
		bobLeaf, err := bobPoster.PostLatestAssertion(ctx)
		if err != nil {
			t.Fatal(err)
		}

		// Scan for created assertions.
		aliceScanner := assertions.NewScanner(aChain, scenario.AliceStateManager, be.Client(), a, rollup, "alice", time.Hour)
		bobScanner := assertions.NewScanner(bChain, scenario.BobStateManager, be.Client(), b, rollup, "bob", time.Hour)

		if err := aliceScanner.ProcessAssertionCreation(ctx, aliceLeaf.Id()); err != nil {
			panic(err)
		}
		if err := bobScanner.ProcessAssertionCreation(ctx, bobLeaf.Id()); err != nil {
			panic(err)
		}

		a.Start(ctx)
		b.Start(ctx)

		g, ctx := errgroup.WithContext(ctx)
		for _, e := range scenario.Expectations {
			fn := e // loop closure
			g.Go(func() error {
				return fn(t, ctx, be)
			})
		}

		if err := g.Wait(); err != nil {
			t.Fatal(err)
		}
	})
}

// testSyncBobStopsCharlieJoins tests the scenario where Bob stops and Charlie joins.
func testSyncBobStopsCharlieJoins(t *testing.T, be backend.Backend, s *ChallengeScenario) {
	t.Run(s.Name, func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
		defer cancel()

		rollup, err := be.DeployRollup()
		require.NoError(t, err)

		// Bad Alice
		aChain, err := solimpl.NewAssertionChain(ctx, rollup, be.Alice(), be.Client())
		require.NoError(t, err)
		alice, err := validator.New(ctx, aChain, be.Client(), s.AliceStateManager, rollup, validator.WithAddress(be.Alice().From), validator.WithName("alice"), validator.WithMode(types.MakeMode))
		require.NoError(t, err)

		// Good Bob
		bobCtx, bobCancelCtx := context.WithCancel(ctx)
		bChain, err := solimpl.NewAssertionChain(bobCtx, rollup, be.Bob(), be.Client())
		require.NoError(t, err)
		bob, err := validator.New(bobCtx, bChain, be.Client(), s.BobStateManager, rollup, validator.WithAddress(be.Bob().From), validator.WithName("bob"), validator.WithMode(types.MakeMode))
		require.NoError(t, err)

		alicePoster := assertions.NewPoster(aChain, s.AliceStateManager, "alice", time.Hour)
		bobPoster := assertions.NewPoster(bChain, s.BobStateManager, "bob", time.Hour)
		aliceLeaf, err := alicePoster.PostLatestAssertion(ctx)
		require.NoError(t, err)
		bobLeaf, err := bobPoster.PostLatestAssertion(bobCtx)
		require.NoError(t, err)
		aliceScanner := assertions.NewScanner(aChain, s.AliceStateManager, be.Client(), alice, rollup, "alice", time.Hour)
		bobScanner := assertions.NewScanner(bChain, s.BobStateManager, be.Client(), bob, rollup, "bob", time.Hour)
		require.NoError(t, aliceScanner.ProcessAssertionCreation(ctx, aliceLeaf.Id()))
		require.NoError(t, bobScanner.ProcessAssertionCreation(bobCtx, bobLeaf.Id()))

		// Alice and bob starts to challenge each other.
		alice.Start(ctx)
		bob.Start(bobCtx)

		// 10s later, bob shuts down
		time.Sleep(10 * time.Second)
		bobCancelCtx()

		// Good Charlie joins
		cChain, err := solimpl.NewAssertionChain(ctx, rollup, be.Charlie(), be.Client())
		require.NoError(t, err)
		charlie, err := validator.New(ctx, cChain, be.Client(), s.CharlieStateManager, rollup, validator.WithAddress(be.Charlie().From), validator.WithName("charlie"), validator.WithMode(types.DefensiveMode)) // Defensive is good enough here.
		require.NoError(t, err)
		charlie.Start(ctx)

		g, ctx := errgroup.WithContext(ctx)
		for _, e := range s.Expectations {
			fn := e // loop closure
			g.Go(func() error {
				return fn(t, ctx, be)
			})
		}

		if err := g.Wait(); err != nil {
			t.Fatal(err)
		}

	})
}

// setupValidator initializes a validator with the minimum required configuration.
func setupValidator(
	ctx context.Context,
	be backend.Backend,
	rollup common.Address,
	sm l2stateprovider.Provider,
	txOpts *bind.TransactOpts,
	name string,
) (*validator.Manager, protocol.Protocol, error) {
	chain, err := solimpl.NewAssertionChain(
		ctx,
		rollup,
		txOpts,
		be.Client(),
	)
	if err != nil {
		return nil, nil, err
	}

	v, err := validator.New(
		ctx,
		chain,
		be.Client(),
		sm,
		rollup,
		validator.WithAddress(txOpts.From),
		validator.WithName(name),
		validator.WithEdgeTrackerWakeInterval(time.Millisecond*250),
		validator.WithMode(types.MakeMode),
	)
	if err != nil {
		return nil, nil, err
	}

	return v, chain, nil
}
