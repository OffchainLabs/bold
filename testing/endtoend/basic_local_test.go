package endtoend

import (
	"context"
	"testing"
	"time"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	solimpl "github.com/OffchainLabs/challenge-protocol-v2/protocol/sol-implementation"
	"github.com/OffchainLabs/challenge-protocol-v2/solgen/go/challengeV2gen"
	"github.com/OffchainLabs/challenge-protocol-v2/solgen/go/rollupgen"
	statemanager "github.com/OffchainLabs/challenge-protocol-v2/state-manager"
	"github.com/OffchainLabs/challenge-protocol-v2/testing/endtoend/internal/backend"
	"github.com/OffchainLabs/challenge-protocol-v2/validator"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/nitro/util/headerreader"
)

type ChallengeScenario struct {
	Name string

	// Validator knowledge
	AliceStateManager statemanager.Manager
	BobStateManager   statemanager.Manager

	// Expectations
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

func TestChallengeProtocol_AliceAndBob_AnvilLocal(t *testing.T) {
	be, err := backend.NewAnvilLocal(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if err := be.Start(); err != nil {
		t.Fatal(err)
	}

	scenarios := []*ChallengeScenario{
		{
			Name: "two forked assertions at the same height",
			AliceStateManager: func() statemanager.Manager {
				sm, err := statemanager.NewForSimpleMachine()
				if err != nil {
					t.Fatal(err)
				}
				return sm
			}(),
			BobStateManager: func() statemanager.Manager {
				cfg := &challengeProtocolTestConfig{
					// The heights at which the validators diverge in histories. In this test,
					// alice and bob start diverging at height 3 at all subchallenge levels.
					assertionDivergenceHeight: 4,
					bigStepDivergenceHeight:   4,
					smallStepDivergenceHeight: 4,
				}
				sm, err := statemanager.NewForSimpleMachine(
					statemanager.WithMachineDivergenceStep(cfg.bigStepDivergenceHeight*protocol.LevelZeroSmallStepEdgeHeight+cfg.smallStepDivergenceHeight),
					statemanager.WithBlockDivergenceHeight(cfg.assertionDivergenceHeight),
					statemanager.WithDivergentBlockHeightOffset(cfg.assertionBlockHeightDifference),
				)
				if err != nil {
					t.Fatal(err)
				}
				return sm
			}(),
			// TODO: Alice should win this challenge.
		},
	} // TODO: Add more scenarios

	for _, scenario := range scenarios {
		testChallengeProtocol_AliceAndBob(t, be, scenario)
	}

}

func testChallengeProtocol_AliceAndBob(t *testing.T, be backend.Backend, scenario *ChallengeScenario) {
	t.Run(scenario.Name, func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 70*time.Second)
		defer cancel()

		rollup, err := be.DeployRollup()
		if err != nil {
			t.Fatal(err)
		}

		a, err := setupValidator(ctx, be, rollup, scenario.AliceStateManager, be.Alice(), "alice")
		if err != nil {
			t.Fatal(err)
		}
		b, err := setupValidator(ctx, be, rollup, scenario.BobStateManager, be.Bob(), "bob")
		if err != nil {
			t.Fatal(err)
		}

		a.Start(ctx)
		b.Start(ctx)

		t.Log("DEBUG: Sleeping for 15s")
		time.Sleep(15 * time.Second)

		// TODO: Abstract this to be part of the scenario success criteria.

		// Read contract events to ensure that Alice and Bob did stuff.
		rc, err := rollupgen.NewRollupCore(rollup, be.Client())
		if err != nil {
			t.Fatal(err)
		}
		cmAddr, err := rc.ChallengeManager(nil)
		if err != nil {
			t.Fatal(err)
		}
		ecm, err := challengeV2gen.NewEdgeChallengeManager(cmAddr, be.Client())
		if err != nil {
			t.Fatal(err)
		}
		i, err := ecm.FilterEdgeAdded(nil, nil, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		for i.Next() {
			t.Logf("DEBUG: Edge %#x added with origin ID %#x", i.Event.EdgeId, i.Event.OriginId)

			edge, err := ecm.GetEdge(nil, i.Event.EdgeId)
			if err != nil {
				t.Fatal(err)
			}
			t.Logf("DEBUG: Edge staker %#x", edge.Staker)
			switch edge.Staker {
			case be.Alice().From:
				t.Log("DEBUG: Alice staker")
			case be.Bob().From:
				t.Log("DEBUG: Bob staker")
			default:
				t.Log("unexpected staker")
			}
		}

		// Ensure a challenge has ended with a one step proof.
		var edgeConfirmed bool
		i2, err2 := ecm.FilterEdgeConfirmedByOneStepProof(nil, nil, nil)
		if err2 != nil {
			t.Fatal(err2)
		}
		for i2.Next() {
			t.Logf("DEBUG: Edge %#x confirmed by one step proof", i2.Event.EdgeId)
			edgeConfirmed = true
		}
		if !edgeConfirmed {
			t.Fatal("FAIL: edge not confirmed by one step proof")
		}

		t.Fail() // Temporary until some success criteria are added.
	})
}

func setupValidator(ctx context.Context, be backend.Backend, rollup common.Address, sm statemanager.Manager, txOpts *bind.TransactOpts, name string) (*validator.Validator, error) {
	hr := headerreader.New(be.Client(), func() *headerreader.Config {
		return &headerreader.DefaultConfig
	})

	chain, err := solimpl.NewAssertionChain(
		ctx,
		rollup,
		txOpts,
		be.Client(),
		hr,
	)
	if err != nil {
		return nil, err
	}

	v, err := validator.New(
		ctx,
		chain,
		be.Client(),
		sm,
		rollup,
		validator.WithAddress(txOpts.From),
		validator.WithName(name),
		validator.WithNewAssertionCheckInterval(500*time.Millisecond),
	)
	if err != nil {
		return nil, err
	}

	return v, nil
}
