package main

import (
	"context"
	"math/big"

	"encoding/binary"
	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/protocol/sol-implementation"
	"github.com/OffchainLabs/challenge-protocol-v2/state-manager"
	chalTesting "github.com/OffchainLabs/challenge-protocol-v2/testing"
	"github.com/OffchainLabs/challenge-protocol-v2/testing/setup"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/OffchainLabs/challenge-protocol-v2/validator"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/offchainlabs/nitro/util/headerreader"
	"github.com/sirupsen/logrus"
	"math"
	"time"
)

var (
	log                        = logrus.WithField("prefix", "main")
	chainId                    = big.NewInt(1337)
	challengePeriodSeconds     = big.NewInt(100)
	miniStakeSize              = big.NewInt(1)
	currentChainHeight         = uint64(7)
	maxWavmOpcodesPerBlock     = uint64(49)
	numOpcodesPerBigStep       = uint64(7)
	divergenceHeight           = uint64(3)
	edgeTrackerWakeInterval    = 100 * time.Millisecond
	checkForChallengesInterval = 100 * time.Millisecond
	checkForAssertionsInteral  = 100 * time.Millisecond
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	// Setup an admin account, Alice and Bob.
	accs, backend, err := setup.SetupAccounts(3)
	if err != nil {
		log.Fatal(err)
	}
	addresses := deployStack(ctx, accs[0], backend)

	headerReader := headerreader.New(util.SimulatedBackendWrapper{
		SimulatedBackend: backend,
	}, func() *headerreader.Config {
		return &headerreader.TestConfig
	})
	headerReader.Start(ctx)

	// Setup the chain abstractions for Alice and Bob.
	aliceChain := setupChainAbstraction(ctx, headerReader, backend, accs[1], addresses)
	bobChain := setupChainAbstraction(ctx, headerReader, backend, accs[2], addresses)

	// Advance the chain by 100 blocks as there needs to be a minimum period of time
	// before any assertions can be made on-chain.
	for i := 0; i < 100; i++ {
		backend.Commit()
	}

	honestHashes := honestHashesForUints(0, currentChainHeight+1)
	evilHashes := evilHashesForUints(0, currentChainHeight+1)

	honestStates, honestInboxCounts := prepareHonestStates(
		ctx,
		aliceChain,
		backend,
		honestHashes,
		currentChainHeight,
	)

	maliciousStates, maliciousInboxCounts := prepareMaliciousStates(
		divergenceHeight,
		evilHashes,
		honestStates,
		honestInboxCounts,
	)

	// Initialize Alice and Bob's respective state managers.
	managerOpts := []statemanager.Opt{
		statemanager.WithMaxWavmOpcodesPerBlock(maxWavmOpcodesPerBlock),
		statemanager.WithNumOpcodesPerBigStep(numOpcodesPerBigStep),
	}
	aliceStateManager, err := statemanager.NewWithAssertionStates(
		honestStates,
		honestInboxCounts,
		managerOpts...,
	)
	if err != nil {
		log.Fatal(err)
	}

	managerOpts = append(
		managerOpts,
		statemanager.WithBigStepStateDivergenceHeight(divergenceHeight),
		statemanager.WithSmallStepStateDivergenceHeight(divergenceHeight),
	)
	bobStateManager, err := statemanager.NewWithAssertionStates(
		maliciousStates,
		maliciousInboxCounts,
		managerOpts...,
	)
	if err != nil {
		log.Fatal(err)
	}

	timeReference := util.NewRealTimeReference()
	commonValidatorOpts := []validator.Opt{
		validator.WithTimeReference(timeReference),
		validator.WithChallengeVertexWakeInterval(edgeTrackerWakeInterval),
		validator.WithNewAssertionCheckInterval(checkForAssertionsInteral),
		validator.WithNewChallengeCheckInterval(checkForAssertionsInteral),
	}
	aliceOpts := []validator.Opt{
		validator.WithName("alice"),
		validator.WithAddress(accs[1].AccountAddr),
	}
	alice, err := validator.New(
		ctx,
		aliceChain,
		backend,
		aliceStateManager,
		addresses.Rollup,
		append(aliceOpts, commonValidatorOpts...)...,
	)
	if err != nil {
		log.Fatal(err)
	}

	bobOpts := []validator.Opt{
		validator.WithName("bob"),
		validator.WithAddress(accs[2].AccountAddr),
	}
	bob, err := validator.New(
		ctx,
		bobChain,
		backend,
		bobStateManager,
		addresses.Rollup,
		append(bobOpts, commonValidatorOpts...)...,
	)
	if err != nil {
		log.Fatal(err)
	}

	go alice.Start(ctx)
	go bob.Start(ctx)

	<-ctx.Done()
}

func setupChainAbstraction(
	ctx context.Context,
	headerReader *headerreader.HeaderReader,
	backend *backends.SimulatedBackend,
	account *setup.TestAccount,
	addrs *setup.RollupAddresses,
) *solimpl.AssertionChain {
	chain, err := solimpl.NewAssertionChain(
		ctx,
		addrs.Rollup,
		account.TxOpts,
		&bind.CallOpts{},
		account.AccountAddr,
		backend,
		headerReader,
		addrs.EdgeChallengeManager,
	)
	if err != nil {
		log.Fatal(err)
	}
	return chain
}

func deployStack(
	ctx context.Context,
	adminAccount *setup.TestAccount,
	backend *backends.SimulatedBackend,
) *setup.RollupAddresses {
	prod := false
	wasmModuleRoot := common.Hash{}
	rollupOwner := adminAccount
	loserStakeEscrow := common.Address{}
	cfg := chalTesting.GenerateRollupConfig(
		prod,
		wasmModuleRoot,
		rollupOwner.AccountAddr,
		chainId,
		loserStakeEscrow,
		challengePeriodSeconds,
		miniStakeSize,
	)
	addrs, err := setup.DeployFullRollupStack(
		ctx,
		backend,
		adminAccount.TxOpts,
		common.Address{}, // Sequencer addr.
		cfg,
	)
	if err != nil {
		log.Fatal(err)
	}
	return addrs
}

func prepareHonestStates(
	ctx context.Context,
	chain protocol.Protocol,
	backend *backends.SimulatedBackend,
	honestHashes []common.Hash,
	chainHeight uint64,
) ([]*protocol.ExecutionState, []*big.Int) {
	genesisState := &protocol.ExecutionState{
		GlobalState: protocol.GoGlobalState{
			BlockHash: common.Hash{},
		},
		MachineStatus: protocol.MachineStatusFinished,
	}

	// Initialize each validator associated state roots which diverge
	// at specified points in the test config.
	honestStates := make([]*protocol.ExecutionState, chainHeight+1)
	honestInboxCounts := make([]*big.Int, chainHeight+1)
	honestStates[0] = genesisState
	honestInboxCounts[0] = big.NewInt(1)

	for i := uint64(1); i <= chainHeight; i++ {
		backend.Commit()
		state := &protocol.ExecutionState{
			GlobalState: protocol.GoGlobalState{
				BlockHash: honestHashes[i],
				Batch:     1,
			},
			MachineStatus: protocol.MachineStatusFinished,
		}

		honestStates[i] = state
		honestInboxCounts[i] = big.NewInt(1)
	}
	return honestStates, honestInboxCounts
}

func prepareMaliciousStates(
	assertionDivergenceHeight uint64,
	evilHashes []common.Hash,
	honestStates []*protocol.ExecutionState,
	honestInboxCounts []*big.Int,
) ([]*protocol.ExecutionState, []*big.Int) {
	divergenceHeight := assertionDivergenceHeight
	numRoots := currentChainHeight + 1
	states := make([]*protocol.ExecutionState, numRoots)
	inboxCounts := make([]*big.Int, numRoots)

	for j := uint64(0); j < numRoots; j++ {
		if divergenceHeight == 0 || j < divergenceHeight {
			states[j] = honestStates[j]
			inboxCounts[j] = honestInboxCounts[j]
		} else {
			evilState := &protocol.ExecutionState{
				GlobalState: protocol.GoGlobalState{
					BlockHash: evilHashes[j],
					Batch:     1,
				},
				MachineStatus: protocol.MachineStatusFinished,
			}
			states[j] = evilState
			inboxCounts[j] = big.NewInt(1)
		}
	}
	return states, inboxCounts
}

func evilHashesForUints(lo, hi uint64) []common.Hash {
	ret := []common.Hash{}
	for i := lo; i < hi; i++ {
		ret = append(ret, hashForUint(math.MaxUint64-i))
	}
	return ret
}

func honestHashesForUints(lo, hi uint64) []common.Hash {
	ret := []common.Hash{}
	for i := lo; i < hi; i++ {
		ret = append(ret, hashForUint(i))
	}
	return ret
}

func hashForUint(x uint64) common.Hash {
	return crypto.Keccak256Hash(binary.BigEndian.AppendUint64([]byte{}, x))
}
