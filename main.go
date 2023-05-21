package main

import (
	"context"
	"math/big"
	"time"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	statemanager "github.com/OffchainLabs/challenge-protocol-v2/state-manager"
	"github.com/OffchainLabs/challenge-protocol-v2/testing/setup"
	"github.com/OffchainLabs/challenge-protocol-v2/validator"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var (
	// The chain id for the backend.
	chainId = big.NewInt(1337)
	// The size of a mini stake that is posted when creating leaf edges in
	// challenges (clarify if gwei?).
	miniStakeSize = big.NewInt(1)
	// The heights at which Alice and Bob diverge at each challenge level.
	divergeHeightAtL2 = uint64(3)
	// How often an edge tracker needs to wake and perform its responsibilities.
	edgeTrackerWakeInterval = time.Millisecond * 500
	// How often the validator polls the chain to see if new assertions have been posted.
	checkForAssertionsInteral = time.Second
	// How often the validator will post its latest assertion to the chain.
	postNewAssertionInterval = time.Second * 5
)

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

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	setupCfg, err := setup.SetupChainsWithEdgeChallengeManager()
	if err != nil {
		panic(err)
	}
	chains := setupCfg.Chains
	accs := setupCfg.Accounts
	addrs := setupCfg.Addrs
	backend := setupCfg.Backend

	// Advance the chain by 100 blocks as there needs to be a minimum period of time
	// before any assertions can be made on-chain.
	for i := 0; i < 100; i++ {
		backend.Commit()
	}

	aliceStateManager, err := statemanager.NewForSimpleMachine()
	if err != nil {
		panic(err)
	}
	cfg := &challengeProtocolTestConfig{
		// The heights at which the validators diverge in histories. In this test,
		// alice and bob start diverging at height 3 at all subchallenge levels.
		assertionDivergenceHeight: 4,
		bigStepDivergenceHeight:   4,
		smallStepDivergenceHeight: 4,
	}
	bobStateManager, err := statemanager.NewForSimpleMachine(
		statemanager.WithMachineDivergenceStep(cfg.bigStepDivergenceHeight*protocol.LevelZeroSmallStepEdgeHeight+cfg.smallStepDivergenceHeight),
		statemanager.WithBlockDivergenceHeight(cfg.assertionDivergenceHeight),
		statemanager.WithDivergentBlockHeightOffset(cfg.assertionBlockHeightDifference),
	)
	if err != nil {
		panic(err)
	}

	a, err := setupValidator(ctx, chains[0], backend, addrs.Rollup, aliceStateManager, "alice", accs[0].TxOpts.From)
	if err != nil {
		panic(err)
	}
	b, err := setupValidator(ctx, chains[1], backend, addrs.Rollup, bobStateManager, "bob", accs[1].TxOpts.From)
	if err != nil {
		panic(err)
	}

	a.Start(ctx)
	b.Start(ctx)
	<-ctx.Done()
}

// setupValidator initializes a validator with the minimum required configuration.
func setupValidator(
	ctx context.Context,
	chain protocol.AssertionChain,
	backend bind.ContractBackend,
	rollup common.Address,
	sm statemanager.Manager,
	name string,
	addr common.Address,
) (*validator.Validator, error) {
	v, err := validator.New(
		ctx,
		chain,
		backend,
		sm,
		rollup,
		validator.WithAddress(addr),
		validator.WithName(name),
		validator.WithEdgeTrackerWakeInterval(time.Second),
		validator.WithNewAssertionCheckInterval(time.Second),
	)
	if err != nil {
		return nil, err
	}

	return v, nil
}
