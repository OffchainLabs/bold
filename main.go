package main

import (
	"context"
	"math/big"

	"encoding/binary"
	"math"
	"time"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	statemanager "github.com/OffchainLabs/challenge-protocol-v2/state-manager"
	"github.com/OffchainLabs/challenge-protocol-v2/testing/setup"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/OffchainLabs/challenge-protocol-v2/validator"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type challengeProtocolTestConfig struct {
	// The latest heights by index at the assertion chain level.
	aliceHeight uint64
	bobHeight   uint64
	// The height in the assertion chain at which the validators diverge.
	assertionDivergenceHeight uint64
	// The heights at which the validators diverge in histories at the big step
	// subchallenge level.
	bigStepDivergenceHeight uint64
	// The heights at which the validators diverge in histories at the small step
	// subchallenge level.
	smallStepDivergenceHeight uint64
	// Events we want to assert are fired from the goimpl.
	expectedBisections  uint64
	expectedLeavesAdded uint64
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	ref := util.NewRealTimeReference()
	setupCfg, err := setup.SetupChainsWithEdgeChallengeManager()
	if err != nil {
		panic(err)
	}
	chains := setupCfg.Chains
	accs := setupCfg.Accounts
	addrs := setupCfg.Addrs
	backend := setupCfg.Backend
	prevInboxMaxCount := big.NewInt(1)

	// Advance the chain by 100 blocks as there needs to be a minimum period of time
	// before any assertions can be made on-chain.
	for i := 0; i < 100; i++ {
		backend.Commit()
	}

	cfg := &challengeProtocolTestConfig{
		// The latest assertion height each validator has seen.
		aliceHeight: 7,
		bobHeight:   7,
		// The heights at which the validators diverge in histories. In this test,
		// alice and bob start diverging at height 3 at all subchallenge levels.
		assertionDivergenceHeight: 4,
		bigStepDivergenceHeight:   4,
		smallStepDivergenceHeight: 4,
	}

	honestHashes := honestHashesForUints(0, cfg.aliceHeight+1)
	evilHashes := evilHashesForUints(0, cfg.bobHeight+1)

	honestStates, honestInboxCounts := prepareHonestStates(
		ctx,
		chains[0],
		backend,
		honestHashes,
		cfg.aliceHeight,
		prevInboxMaxCount,
	)

	maliciousStates, maliciousInboxCounts := prepareMaliciousStates(
		cfg,
		evilHashes,
		honestStates,
		honestInboxCounts,
	)

	// Initialize each validator.
	honestManager, err := statemanager.NewWithAssertionStates(
		honestStates,
		honestInboxCounts,
		statemanager.WithNumOpcodesPerBigStep(protocol.LevelZeroSmallStepEdgeHeight),
		statemanager.WithMaxWavmOpcodesPerBlock(protocol.LevelZeroBigStepEdgeHeight*protocol.LevelZeroSmallStepEdgeHeight),
	)
	if err != nil {
		panic(err)
	}
	aliceWatcher := validator.NewWatcher(chains[0], backend, time.Millisecond*500)
	aliceAddr := accs[0].AccountAddr
	alice, err := validator.New(
		ctx,
		chains[0],
		backend,
		honestManager,
		addrs.Rollup,
		validator.WithName("alice"),
		validator.WithAddress(aliceAddr),
		validator.WithTimeReference(ref),
		validator.WithEdgeTrackerWakeInterval(time.Millisecond*200),
		validator.WithNewAssertionCheckInterval(time.Millisecond*50),
		validator.WithWatcher(aliceWatcher),
	)
	if err != nil {
		panic(err)
	}

	maliciousManager, err := statemanager.NewWithAssertionStates(
		maliciousStates,
		maliciousInboxCounts,
		statemanager.WithMaliciousIntent(),
		statemanager.WithNumOpcodesPerBigStep(protocol.LevelZeroSmallStepEdgeHeight),
		statemanager.WithMaxWavmOpcodesPerBlock(protocol.LevelZeroBigStepEdgeHeight*protocol.LevelZeroSmallStepEdgeHeight),
		statemanager.WithBigStepStateDivergenceHeight(cfg.bigStepDivergenceHeight),
		statemanager.WithSmallStepStateDivergenceHeight(cfg.smallStepDivergenceHeight),
	)
	if err != nil {
		panic(err)
	}
	bobWatcher := validator.NewWatcher(chains[1], backend, time.Millisecond*500)
	bobAddr := accs[1].AccountAddr
	bob, err := validator.New(
		ctx,
		chains[1],
		backend,
		maliciousManager,
		addrs.Rollup,
		validator.WithName("bob"),
		validator.WithAddress(bobAddr),
		validator.WithTimeReference(ref),
		validator.WithEdgeTrackerWakeInterval(time.Millisecond*200),
		validator.WithNewAssertionCheckInterval(time.Millisecond*50),
		validator.WithWatcher(bobWatcher),
	)
	if err != nil {
		panic(err)
	}

	challengeManager, err := chains[0].SpecChallengeManager(ctx)
	if err != nil {
		panic(err)
	}

	genesis, err := chains[0].AssertionBySequenceNum(ctx, protocol.GenesisAssertionSeqNum)
	if err != nil {
		panic(err)
	}
	genesisStateHash, err := genesis.StateHash()
	if err != nil {
		panic(err)
	}

	// Submit leaf creation manually for each validator.
	genesisState, err := honestManager.AssertionExecutionState(ctx, genesisStateHash)
	if err != nil {
		panic(err)
	}
	latestHonest, err := honestManager.LatestAssertionCreationData(ctx)
	if err != nil {
		panic(err)
	}
	leaf1, err := chains[0].CreateAssertion(
		ctx,
		genesisState,
		latestHonest.State,
		latestHonest.InboxMaxCount,
	)
	if err != nil {
		panic(err)
	}

	latestEvil, err := maliciousManager.LatestAssertionCreationData(ctx)
	if err != nil {
		panic(err)
	}
	leaf2, err := chains[1].CreateAssertion(
		ctx,
		genesisState,
		latestEvil.State,
		latestEvil.InboxMaxCount,
	)
	if err != nil {
		panic(err)
	}

	// Honest assertion being added.
	leafAdder := func(startCommit, endCommit util.HistoryCommitment, prefixProof []byte, leaf protocol.Assertion) protocol.SpecEdge {
		edge, err := challengeManager.AddBlockChallengeLevelZeroEdge(
			ctx,
			leaf,
			startCommit,
			endCommit,
			prefixProof,
		)
		if err != nil {
			panic(err)
		}
		return edge
	}

	honestStartCommit, err := honestManager.HistoryCommitmentUpTo(ctx, 0)
	if err != nil {
		panic(err)
	}
	honestEndCommit, err := honestManager.HistoryCommitmentUpToBatch(ctx, 0, protocol.LevelZeroBlockEdgeHeight, 1)
	if err != nil {
		panic(err)
	}
	honestPrefixProof, err := honestManager.PrefixProofUpToBatch(ctx, 0, 0, protocol.LevelZeroBlockEdgeHeight, 1)
	if err != nil {
		panic(err)
	}

	honestEdge := leafAdder(honestStartCommit, honestEndCommit, honestPrefixProof, leaf1)

	evilStartCommit, err := maliciousManager.HistoryCommitmentUpTo(ctx, 0)
	if err != nil {
		panic(err)
	}
	evilEndCommit, err := maliciousManager.HistoryCommitmentUpToBatch(ctx, 0, protocol.LevelZeroBlockEdgeHeight, 1)
	if err != nil {
		panic(err)
	}
	evilPrefixProof, err := maliciousManager.PrefixProofUpToBatch(ctx, 0, 0, protocol.LevelZeroBlockEdgeHeight, 1)
	if err != nil {
		panic(err)
	}

	evilEdge := leafAdder(evilStartCommit, evilEndCommit, evilPrefixProof, leaf2)

	go aliceWatcher.Watch(ctx)
	go bobWatcher.Watch(ctx)

	alice.SpawnEdgeTracker(
		ctx,
		honestEdge,
		0,
		prevInboxMaxCount.Uint64(),
	)
	bob.SpawnEdgeTracker(
		ctx,
		evilEdge,
		0,
		prevInboxMaxCount.Uint64(),
	)
	<-ctx.Done()
}

func prepareHonestStates(
	ctx context.Context,
	chain protocol.Protocol,
	backend *backends.SimulatedBackend,
	honestHashes []common.Hash,
	chainHeight uint64,
	prevInboxMaxCount *big.Int,
) ([]*protocol.ExecutionState, []*big.Int) {
	// Initialize each validator's associated state roots which diverge
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
				BlockHash:  honestHashes[i],
				Batch:      0,
				PosInBatch: i,
			},
			MachineStatus: protocol.MachineStatusFinished,
		}
		if i == chainHeight {
			state.GlobalState.Batch = 1
			state.GlobalState.PosInBatch = 0
		}

		honestStates[i] = state
		honestInboxCounts[i] = big.NewInt(2)
	}
	return honestStates, honestInboxCounts
}

func prepareMaliciousStates(
	cfg *challengeProtocolTestConfig,
	evilHashes []common.Hash,
	honestStates []*protocol.ExecutionState,
	honestInboxCounts []*big.Int,
) ([]*protocol.ExecutionState, []*big.Int) {
	divergenceHeight := cfg.assertionDivergenceHeight
	numRoots := cfg.bobHeight + 1
	states := make([]*protocol.ExecutionState, numRoots)
	inboxCounts := make([]*big.Int, numRoots)

	for j := uint64(0); j < numRoots; j++ {
		if divergenceHeight == 0 || j < divergenceHeight {
			evilState := *honestStates[j]
			if j < cfg.bobHeight {
				evilState.GlobalState.Batch = 0
				evilState.GlobalState.PosInBatch = j
			}
			states[j] = &evilState
			inboxCounts[j] = honestInboxCounts[j]
		} else {
			evilState := &protocol.ExecutionState{
				GlobalState: protocol.GoGlobalState{
					BlockHash:  evilHashes[j],
					Batch:      0,
					PosInBatch: j,
				},
				MachineStatus: protocol.MachineStatusFinished,
			}
			if j == cfg.bobHeight {
				evilState.GlobalState.Batch = 1
				evilState.GlobalState.PosInBatch = 0
			}
			states[j] = evilState
			inboxCounts[j] = big.NewInt(2)
		}
	}
	return states, inboxCounts
}

func evilHashesForUints(lo, hi uint64) []common.Hash {
	var ret []common.Hash
	for i := lo; i < hi; i++ {
		ret = append(ret, hashForUint(math.MaxUint64-i))
	}
	return ret
}

func honestHashesForUints(lo, hi uint64) []common.Hash {
	var ret []common.Hash
	for i := lo; i < hi; i++ {
		ret = append(ret, hashForUint(i))
	}
	return ret
}

func hashForUint(x uint64) common.Hash {
	return crypto.Keccak256Hash(binary.BigEndian.AppendUint64([]byte{}, x))
}
