package endtoend

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/OffchainLabs/bold/api"
	challengemanager "github.com/OffchainLabs/bold/challenge-manager"
	"github.com/OffchainLabs/bold/challenge-manager/types"
	"github.com/OffchainLabs/bold/solgen/go/challengeV2gen"
	"github.com/OffchainLabs/bold/solgen/go/rollupgen"
	challenge_testing "github.com/OffchainLabs/bold/testing"
	"github.com/OffchainLabs/bold/testing/endtoend/backend"
	statemanager "github.com/OffchainLabs/bold/testing/mocks/state-provider"
	"github.com/OffchainLabs/bold/testing/setup"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

func TestEndToEnd_ChallengeManagerContractUpgradeMidChallenge(t *testing.T) {
	runEndToEndTestMidChallengeUpgrade(t, &e2eConfig{
		backend:  simulated,
		protocol: defaultProtocolParams(),
		inbox:    defaultInboxParams(),
		actors: actorParams{
			numEvilValidators: 1,
		},
		timings: defaultTimeParams(),
		expectations: []expect{
			// Expect one assertion is confirmed by challenge win.
			expectAssertionConfirmedByChallengeWin,
		},
	})
}

func runEndToEndTestMidChallengeUpgrade(t *testing.T, cfg *e2eConfig) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Validators include a chain admin, a single honest validators, and any number of evil entities.
	totalValidators := cfg.actors.numEvilValidators + 2

	challengeTestingOpts := []challenge_testing.Opt{
		challenge_testing.WithConfirmPeriodBlocks(cfg.protocol.challengePeriodBlocks),
		challenge_testing.WithLayerZeroHeights(&cfg.protocol.layerZeroHeights),
		challenge_testing.WithNumBigStepLevels(cfg.protocol.numBigStepLevels),
	}
	deployOpts := []setup.Opt{
		setup.WithMockBridge(),
		setup.WithMockOneStepProver(),
		setup.WithNumAccounts(totalValidators),
		setup.WithChallengeTestingOpts(challengeTestingOpts...),
	}

	var bk backend.Backend
	switch cfg.backend {
	case simulated:
		simBackend, err := backend.NewSimulated(cfg.timings.blockTime, deployOpts...)
		require.NoError(t, err)
		bk = simBackend
	case anvil:
		anvilBackend, err := backend.NewAnvilLocal(ctx)
		require.NoError(t, err)
		bk = anvilBackend
	default:
		t.Fatalf("Backend kind for e2e test not supported: %s", cfg.backend)
	}

	rollupAddr, err := bk.DeployRollup(ctx, challengeTestingOpts...)
	require.NoError(t, err)

	require.NoError(t, bk.Start(ctx))

	rollupAdminBindings, err := rollupgen.NewRollupAdminLogic(rollupAddr, bk.Client())
	require.NoError(t, err)
	accounts := bk.Accounts()
	_, err = rollupAdminBindings.SetMinimumAssertionPeriod(accounts[0], big.NewInt(1))
	require.NoError(t, err)

	bk.Commit()

	baseStateManagerOpts := []statemanager.Opt{
		statemanager.WithNumBatchesRead(cfg.inbox.numBatchesPosted),
		statemanager.WithLayerZeroHeights(&cfg.protocol.layerZeroHeights, cfg.protocol.numBigStepLevels),
	}
	honestStateManager, err := statemanager.NewForSimpleMachine(baseStateManagerOpts...)
	require.NoError(t, err)

	baseChallengeManagerOpts := []challengemanager.Opt{
		challengemanager.WithEdgeTrackerWakeInterval(cfg.timings.challengeMoveInterval),
		challengemanager.WithMode(types.MakeMode),
		challengemanager.WithAssertionPostingInterval(cfg.timings.assertionPostingInterval),
		challengemanager.WithAssertionScanningInterval(cfg.timings.assertionScanningInterval),
		challengemanager.WithAssertionConfirmingInterval(cfg.timings.assertionConfirmationAttemptInterval),
	}

	name := "honest"
	txOpts := accounts[1]
	honestOpts := append(
		baseChallengeManagerOpts,
		challengemanager.WithAddress(txOpts.From),
		challengemanager.WithName(name),
	)
	honestManager := setupChallengeManager(
		t, ctx, bk.Client(), rollupAddr, honestStateManager, txOpts, name, honestOpts...,
	)
	if !api.IsNil(honestManager.Database()) {
		honestStateManager.UpdateAPIDatabase(honestManager.Database())
	}

	// Diverge exactly at the last opcode within the block.
	totalOpcodes := totalWasmOpcodes(&cfg.protocol.layerZeroHeights, cfg.protocol.numBigStepLevels)
	t.Logf("Total wasm opcodes in test: %d", totalOpcodes)

	assertionDivergenceHeight := uint64(1)
	assertionBlockHeightDifference := int64(1)

	evilChallengeManagers := make([]*challengemanager.Manager, cfg.actors.numEvilValidators)
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
		txOpts = accounts[2+i]
		name = fmt.Sprintf("evil-%d", i)
		evilOpts := append(
			baseChallengeManagerOpts,
			challengemanager.WithAddress(txOpts.From),
			challengemanager.WithName(name),
		)
		evilManager := setupChallengeManager(
			t, ctx, bk.Client(), rollupAddr, evilStateManager, txOpts, name, evilOpts...,
		)
		evilChallengeManagers[i] = evilManager
	}

	honestManager.Start(ctx)

	for _, evilManager := range evilChallengeManagers {
		evilManager.Start(ctx)
	}

	g, ctx := errgroup.WithContext(ctx)
	for _, e := range cfg.expectations {
		fn := e // loop closure
		g.Go(func() error {
			return fn(t, ctx, bk.ContractAddresses(), bk.Client())
		})
	}

	time.Sleep(time.Second * 5)
	g.Go(func() error {
		existingChalManagerAddr, err := rollupAdminBindings.ChallengeManager(&bind.CallOpts{})
		require.NoError(t, err)
		chalManagerAddr, tx, _, err := challengeV2gen.DeployEdgeChallengeManager(
			accounts[0],
			bk.Client(),
		)
		require.NoError(t, err)
		require.NoError(t, challenge_testing.TxSucceeded(ctx, tx, chalManagerAddr, bk.Client(), err))
		tx, err = rollupAdminBindings.SetChallengeManager(accounts[0], chalManagerAddr)
		require.NoError(t, err)
		require.NoError(t, challenge_testing.TxSucceeded(ctx, tx, chalManagerAddr, bk.Client(), err))
		t.Logf("Challenge manager address updated from %#x to %#x", existingChalManagerAddr, chalManagerAddr)
		return nil
	})
	require.NoError(t, g.Wait())
}
