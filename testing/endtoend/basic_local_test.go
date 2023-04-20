package endtoend

import (
	"context"
	"testing"
	"time"

	solimpl "github.com/OffchainLabs/challenge-protocol-v2/protocol/sol-implementation"
	statemanager "github.com/OffchainLabs/challenge-protocol-v2/state-manager"
	"github.com/OffchainLabs/challenge-protocol-v2/testing/endtoend/internal/backend"
	"github.com/OffchainLabs/challenge-protocol-v2/validator"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/nitro/util/headerreader"
)

type ChallengeScenario struct {
	// Validator knowledge
	AliceStateManager statemanager.Manager
	BobStateManager   statemanager.Manager

	// Expectations
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
			AliceStateManager: func() statemanager.Manager {
				sm, err := statemanager.New(
					[]common.Hash{ // stateRoots
						common.HexToHash("0x1"),
						common.HexToHash("0x2"),
						common.HexToHash("0x3"),
						common.HexToHash("0x4"),
						common.HexToHash("0x5"),
					},
				)
				if err != nil {
					t.Fatal(err)
				}
				return sm
			}(),
			BobStateManager: func() statemanager.Manager {
				sm, err := statemanager.New(
					[]common.Hash{ // stateRoots
						common.HexToHash("0x1"),
						common.HexToHash("0x2"),
						common.HexToHash("0x3"),
						common.HexToHash("0x5"),
						common.HexToHash("0x6"),
					},
				)
				if err != nil {
					t.Fatal(err)
				}
				return sm
			}(),
		},
	} // TODO: Add more scenarios

	for _, scenario := range scenarios {
		testChallengeProtocol_AliceAndBob(t, be, scenario)
	}
}

func testChallengeProtocol_AliceAndBob(t *testing.T, be backend.Backend, scenario *ChallengeScenario) {
	ctx, cancel := context.WithCancel(context.Background())
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

	// TODO: Wait for evaluation?

	t.Log("DEBUG: Sleeping for 30s")
	time.Sleep(30 * time.Second)
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
