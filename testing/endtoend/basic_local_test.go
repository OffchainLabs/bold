package endtoend

import (
	"context"
	"fmt"
	"testing"
	"time"

	protocol "github.com/OffchainLabs/challenge-protocol-v2/chain-abstraction"
	solimpl "github.com/OffchainLabs/challenge-protocol-v2/chain-abstraction/sol-implementation"
	validator "github.com/OffchainLabs/challenge-protocol-v2/challenge-manager"
	l2stateprovider "github.com/OffchainLabs/challenge-protocol-v2/layer2-state-provider"
	retry "github.com/OffchainLabs/challenge-protocol-v2/runtime"
	challenge_testing "github.com/OffchainLabs/challenge-protocol-v2/testing"
	"github.com/OffchainLabs/challenge-protocol-v2/testing/endtoend/internal/backend"
	"github.com/OffchainLabs/challenge-protocol-v2/testing/toys/assertions"
	statemanager "github.com/OffchainLabs/challenge-protocol-v2/testing/toys/state-provider"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/nitro/util/headerreader"
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
	require.NoError(t, err)

	require.NoError(t, be.Start())
	defer func() {
		if err := be.Stop(); err != nil {
			t.Log(fmt.Errorf("error stopping backend: %v", err))
		}
	}()

	scenario := &ChallengeScenario{
		Name: "two forked assertions at the same height",
		AliceStateManager: func() l2stateprovider.Provider {
			sm, err := statemanager.NewForSimpleMachine()
			require.NoError(t, err)
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
			expectOneStepProofSuccessful,
			expectAssertionConfirmedByChallengeWinner,
			expectLevelZeroBlockEdgeConfirmed,
			expectAliceAndBobStaked,
		},
	}

	testChallengeProtocol_AliceAndBob(t, be, scenario)
}

func TestChallengeProtocol_AliceAndBob_AnvilLocal_DifferentHeights(t *testing.T) {
	t.Skip()
	be, err := backend.NewAnvilLocal(context.Background())
	require.NoError(t, err)

	require.NoError(t, be.Start())
	defer func() {
		require.NoError(t, be.Stop(), "error stopping backend")
	}()

	scenario := &ChallengeScenario{
		Name: "two forked assertions at the different step heights",
		AliceStateManager: func() l2stateprovider.Provider {
			sm, err := statemanager.NewForSimpleMachine()
			require.NoError(t, err)
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
			require.NoError(t, err)
			return sm
		}(),
		Expectations: []expect{
			expectOneStepProofSuccessful,
			expectAssertionConfirmedByChallengeWinner,
			expectLevelZeroBlockEdgeConfirmed,
			expectAliceAndBobStaked,
		},
	}
	testChallengeProtocol_AliceAndBob(t, be, scenario)
}

func TestSync_HonestBobStopsCharlieJoins(t *testing.T) {
	be, err := backend.NewAnvilLocal(context.Background())
	require.NoError(t, err)
	require.NoError(t, be.Start())

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
			require.NoError(t, err)
			return sm
		}(),
		BobStateManager: func() l2stateprovider.Provider {
			sm, err := statemanager.NewForSimpleMachine()
			require.NoError(t, err)
			return sm
		}(),
		CharlieStateManager: func() l2stateprovider.Provider {
			sm, err := statemanager.NewForSimpleMachine()
			require.NoError(t, err)
			return sm
		}(),
		Expectations: []expect{
			expectLevelZeroBlockEdgeConfirmed,
			expectAssertionConfirmedByChallengeWinner,
			expectOneStepProofSuccessful,
		},
	}

	testSyncBobStopsCharlieJoins(t, be, scenario)
}

func testChallengeProtocol_AliceAndBob(t *testing.T, be backend.Backend, scenario *ChallengeScenario) {
	t.Run(scenario.Name, func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
		defer cancel()

		rollup, err := retry.UntilSucceeds(ctx, func() (common.Address, error) {
			return be.DeployRollup()
		})
		require.NoError(t, err)

		a, aChain, err := setupValidator(ctx, be, rollup, scenario.AliceStateManager, be.Alice(), "alice")
		require.NoError(t, err)
		b, bChain, err := setupValidator(ctx, be, rollup, scenario.BobStateManager, be.Bob(), "bob")
		require.NoError(t, err)

		// Post assertions.
		alicePoster := assertions.NewPoster(aChain, scenario.AliceStateManager, "alice", time.Hour)
		bobPoster := assertions.NewPoster(bChain, scenario.BobStateManager, "bob", time.Hour)

		aliceLeaf, err := retry.UntilSucceeds(ctx, func() (protocol.Assertion, error) {
			return alicePoster.PostLatestAssertion(ctx)
		})
		require.NoError(t, err)
		bobLeaf, err := retry.UntilSucceeds(ctx, func() (protocol.Assertion, error) {
			return bobPoster.PostLatestAssertion(ctx)
		})
		require.NoError(t, err)

		// Scan for created assertions.
		aliceScanner := assertions.NewScanner(aChain, scenario.AliceStateManager, be.Client(), a, rollup, "alice", time.Hour)
		bobScanner := assertions.NewScanner(bChain, scenario.BobStateManager, be.Client(), b, rollup, "bob", time.Hour)

		_, err = retry.UntilSucceeds(ctx, func() (bool, error) {
			processErr := aliceScanner.ProcessAssertionCreation(ctx, aliceLeaf.Id())
			return false, processErr
		})
		require.NoError(t, err)
		_, err = retry.UntilSucceeds(ctx, func() (bool, error) {
			processErr := bobScanner.ProcessAssertionCreation(ctx, bobLeaf.Id())
			return false, processErr
		})
		require.NoError(t, err)

		a.Start(ctx)
		b.Start(ctx)

		g, ctx := errgroup.WithContext(ctx)
		for _, e := range scenario.Expectations {
			fn := e // loop closure
			g.Go(func() error {
				return fn(t, ctx, be)
			})
		}

		require.NoError(t, g.Wait())
	})
}

// testSyncBobStopsCharlieJoins tests the scenario where Bob stops and Charlie joins.
func testSyncBobStopsCharlieJoins(t *testing.T, be backend.Backend, s *ChallengeScenario) {
	t.Run(s.Name, func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
		defer cancel()
		rollup, err := retry.UntilSucceeds(ctx, func() (common.Address, error) {
			return be.DeployRollup()
		})
		require.NoError(t, err)

		// Bad Alice
		alice, aChain, err := setupValidator(ctx, be, rollup, s.AliceStateManager, be.Alice(), "alice")
		require.NoError(t, err)

		// Good Bob
		bobCtx, bobCancelCtx := context.WithCancel(ctx)
		bob, bChain, err := setupValidator(bobCtx, be, rollup, s.BobStateManager, be.Bob(), "bob")
		require.NoError(t, err)

		alicePoster := assertions.NewPoster(aChain, s.AliceStateManager, "alice", time.Hour)
		bobPoster := assertions.NewPoster(bChain, s.BobStateManager, "bob", time.Hour)

		aliceLeaf, err := retry.UntilSucceeds(ctx, func() (protocol.Assertion, error) {
			return alicePoster.PostLatestAssertion(ctx)
		})
		require.NoError(t, err)
		bobLeaf, err := retry.UntilSucceeds(ctx, func() (protocol.Assertion, error) {
			return bobPoster.PostLatestAssertion(ctx)
		})
		require.NoError(t, err)

		aliceScanner := assertions.NewScanner(aChain, s.AliceStateManager, be.Client(), alice, rollup, "alice", time.Hour)
		bobScanner := assertions.NewScanner(bChain, s.BobStateManager, be.Client(), bob, rollup, "bob", time.Hour)

		_, err = retry.UntilSucceeds(ctx, func() (bool, error) {
			processErr := aliceScanner.ProcessAssertionCreation(ctx, aliceLeaf.Id())
			return false, processErr
		})
		require.NoError(t, err)
		_, err = retry.UntilSucceeds(ctx, func() (bool, error) {
			processErr := bobScanner.ProcessAssertionCreation(ctx, bobLeaf.Id())
			return false, processErr
		})
		require.NoError(t, err)

		// Alice and bob starts to challenge each other.
		alice.Start(ctx)
		bob.Start(bobCtx)

		// 10s later, bob shuts down
		time.Sleep(10 * time.Second)
		bobCancelCtx()

		// Good Charlie joins
		charlie, _, err := setupValidator(ctx, be, rollup, s.CharlieStateManager, be.Charlie(), "charlie")
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
	hr := headerreader.New(be.Client(), func() *headerreader.Config {
		return &headerreader.DefaultConfig
	})

	chain, err := retry.UntilSucceeds(ctx, func() (*solimpl.AssertionChain, error) {
		return solimpl.NewAssertionChain(
			ctx,
			rollup,
			txOpts,
			be.Client(),
			hr,
		)
	})
	if err != nil {
		return nil, nil, err
	}

	v, err := retry.UntilSucceeds(ctx, func() (*validator.Manager, error) {
		return validator.New(
			ctx,
			chain,
			be.Client(),
			sm,
			rollup,
			validator.WithAddress(txOpts.From),
			validator.WithName(name),
			validator.WithEdgeTrackerWakeInterval(time.Millisecond*250),
		)
	})
	if err != nil {
		return nil, nil, err
	}

	return v, chain, nil
}
