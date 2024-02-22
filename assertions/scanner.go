// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE

// Package assertions contains testing utilities for posting and scanning for
// assertions on chain, which are useful for simulating the responsibilities
// of Arbitrum Nitro and initiating challenges as needed using our challenge manager.
package assertions

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"time"

	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/metrics"

	"github.com/OffchainLabs/bold/api"
	"github.com/OffchainLabs/bold/api/db"
	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	solimpl "github.com/OffchainLabs/bold/chain-abstraction/sol-implementation"
	"github.com/OffchainLabs/bold/challenge-manager/types"
	"github.com/OffchainLabs/bold/containers/option"
	"github.com/OffchainLabs/bold/containers/threadsafe"
	l2stateprovider "github.com/OffchainLabs/bold/layer2-state-provider"
	retry "github.com/OffchainLabs/bold/runtime"
	"github.com/OffchainLabs/bold/solgen/go/rollupgen"
	"github.com/OffchainLabs/bold/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/pkg/errors"
)

var (
	srvlog                                = log.New("service", "assertions")
	evilAssertionCounter                  = metrics.NewRegisteredCounter("arb/validator/scanner/evil_assertion", nil)
	challengeSubmittedCounter             = metrics.NewRegisteredCounter("arb/validator/scanner/challenge_submitted", nil)
	assertionConfirmedCounter             = metrics.GetOrRegisterCounter("arb/validator/scanner/assertion_confirmed", nil)
	errorConfirmingAssertionByTimeCounter = metrics.NewRegisteredCounter("arb/validator/scanner/error_confirming_assertion_by_time", nil)
	latestConfirmedAssertionGauge         = metrics.NewRegisteredGauge("arb/validator/scanner/latest_confirmed_assertion_block_number", nil)
)

func init() {
	srvlog.SetHandler(log.StreamHandler(os.Stdout, log.LogfmtFormat()))
}

// The Manager struct is responsible for several tasks related to the assertion chain:
// 1. It continuously polls the assertion chain to check for posted, on-chain assertions starting from the latest confirmed assertion up to the newest one.
// 2. As the assertion chain advances, the Manager keeps polling to stay updated.
// 3. Upon observing each new assertion, the Manager evaluates whether it should challenge the assertion or not.
// 4. The Manager frequently posts new assertions to the assertion chain at specific intervals.
// 5. When posting assertions, it relies on the most recent execution state available in its local state manager.
type Manager struct {
	chain                       protocol.AssertionChain
	backend                     bind.ContractBackend
	challengeCreator            types.ChallengeCreator
	challengeReader             types.ChallengeReader
	stateProvider               l2stateprovider.ExecutionStateAgreementChecker
	pollInterval                time.Duration
	confirmationAttemptInterval time.Duration
	averageTimeForBlockCreation time.Duration
	rollupAddr                  common.Address
	challengeManagerAddr        common.Address
	validatorName               string
	forksDetectedCount          uint64
	challengesSubmittedCount    uint64
	assertionsProcessedCount    uint64
	submittedRivalsCount        uint64
	stateManager                l2stateprovider.ExecutionProvider
	postInterval                time.Duration
	submittedAssertions         *threadsafe.LruSet[common.Hash]
	assertionsWithHonestChild   *threadsafe.LruSet[protocol.AssertionHash]
	apiDB                       db.Database
}

// NewManager creates a manager from the required dependencies.
func NewManager(
	chain protocol.AssertionChain,
	stateProvider l2stateprovider.Provider,
	backend bind.ContractBackend,
	challengeManager types.ChallengeManager,
	rollupAddr common.Address,
	challengeManagerAddr common.Address,
	validatorName string,
	pollInterval,
	assertionConfirmationAttemptInterval time.Duration,
	stateManager l2stateprovider.ExecutionProvider,
	postInterval time.Duration,
	averageTimeForBlockCreation time.Duration,
	apiDB db.Database,
) (*Manager, error) {
	if pollInterval == 0 {
		return nil, errors.New("assertion scanning interval must be greater than 0")
	}
	if assertionConfirmationAttemptInterval == 0 {
		return nil, errors.New("assertion confirmation attempt interval must be greater than 0")
	}
	return &Manager{
		chain:                       chain,
		apiDB:                       apiDB,
		backend:                     backend,
		stateProvider:               stateProvider,
		challengeCreator:            challengeManager,
		challengeReader:             challengeManager,
		rollupAddr:                  rollupAddr,
		challengeManagerAddr:        challengeManagerAddr,
		validatorName:               validatorName,
		pollInterval:                pollInterval,
		confirmationAttemptInterval: assertionConfirmationAttemptInterval,
		forksDetectedCount:          0,
		challengesSubmittedCount:    0,
		assertionsProcessedCount:    0,
		stateManager:                stateManager,
		postInterval:                postInterval,
		submittedAssertions:         threadsafe.NewLruSet[common.Hash](1000, threadsafe.LruSetWithMetric[common.Hash]("submittedAssertions")),
		assertionsWithHonestChild:   threadsafe.NewLruSet[protocol.AssertionHash](1000, threadsafe.LruSetWithMetric[protocol.AssertionHash]("assertionsWithHonestChild")),
		averageTimeForBlockCreation: averageTimeForBlockCreation,
	}, nil
}

// The Start function begins two main tasks:
// 1. It initiates scanning of the assertion chain for newly created assertions, starting from the latest confirmed assertion. This scanning is done via polling.
// 2. Concurrently, it also starts a routine that is responsible for posting new assertions to the assertion chain.
func (m *Manager) Start(ctx context.Context) {
	go m.postAssertionRoutine(ctx)
	go m.updateLatestConfirmedMetrics(ctx)

	latestConfirmed, err := retry.UntilSucceeds(ctx, func() (protocol.Assertion, error) {
		return m.chain.LatestConfirmed(ctx)
	})
	if err != nil {
		srvlog.Error("Could not get latest confirmed assertion", log.Ctx{"err": err})
		return
	}
	fromBlock := latestConfirmed.CreatedAtBlock()

	filterer, err := retry.UntilSucceeds(ctx, func() (*rollupgen.RollupUserLogicFilterer, error) {
		return rollupgen.NewRollupUserLogicFilterer(m.rollupAddr, m.backend)
	})
	if err != nil {
		srvlog.Error("Could not get rollup user logic filterer", log.Ctx{"err": err})
		return
	}
	latestBlock, err := retry.UntilSucceeds(ctx, func() (*gethtypes.Header, error) {
		return m.backend.HeaderByNumber(ctx, util.GetFinalizedBlockNumber())
	})
	if err != nil {
		srvlog.Error("Could not get header by number", log.Ctx{"err": err})
		return
	}
	if !latestBlock.Number.IsUint64() {
		srvlog.Error("Latest block number was not a uint64")
		return
	}
	toBlock := latestBlock.Number.Uint64()
	if fromBlock != toBlock {
		filterOpts := &bind.FilterOpts{
			Start:   fromBlock,
			End:     &toBlock,
			Context: ctx,
		}
		_, err = retry.UntilSucceeds(ctx, func() (bool, error) {
			return true, m.checkForAssertionAdded(ctx, filterer, filterOpts)
		})
		if err != nil {
			srvlog.Error("Could not check for assertion added event")
			return
		}
		fromBlock = toBlock
	}

	ticker := time.NewTicker(m.pollInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			latestBlock, err := m.backend.HeaderByNumber(ctx, util.GetFinalizedBlockNumber())
			if err != nil {
				srvlog.Error("Could not get header by number", log.Ctx{"err": err})
				continue
			}
			if !latestBlock.Number.IsUint64() {
				srvlog.Error("Latest block number was not a uint64")
				continue
			}
			toBlock := latestBlock.Number.Uint64()
			if fromBlock == toBlock {
				continue
			}
			filterOpts := &bind.FilterOpts{
				Start:   fromBlock,
				End:     &toBlock,
				Context: ctx,
			}
			_, err = retry.UntilSucceeds(ctx, func() (bool, error) {
				return true, m.checkForAssertionAdded(ctx, filterer, filterOpts)
			})
			if err != nil {
				srvlog.Error("Could not check for assertion added", log.Ctx{"err": err})
				return
			}
			fromBlock = toBlock
		case <-ctx.Done():
			return
		}
	}
}

func (m *Manager) ForksDetected() uint64 {
	return m.forksDetectedCount
}

func (m *Manager) ChallengesSubmitted() uint64 {
	return m.challengesSubmittedCount
}

func (m *Manager) SubmittedAssertions() *threadsafe.LruSet[common.Hash] {
	return m.submittedAssertions
}
func (m *Manager) AssertionsProcessed() uint64 {
	return m.assertionsProcessedCount
}

func (m *Manager) SubmittedRivals() uint64 {
	return m.submittedRivalsCount
}

func (m *Manager) AssertionHasHonestChild(hash protocol.AssertionHash) bool {
	return m.assertionsWithHonestChild.Has(hash)
}

func (m *Manager) AssertionsSubmittedInProcess() []common.Hash {
	hashes := make([]common.Hash, 0)
	m.submittedAssertions.ForEach(func(elem common.Hash) {
		hashes = append(hashes, elem)
	})
	return hashes
}

func (m *Manager) checkForAssertionAdded(
	ctx context.Context,
	filterer *rollupgen.RollupUserLogicFilterer,
	filterOpts *bind.FilterOpts,
) error {
	it, err := filterer.FilterAssertionCreated(filterOpts, nil, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err = it.Close(); err != nil {
			srvlog.Error("Could not close filter iterator", log.Ctx{"err": err})
		}
	}()
	for it.Next() {
		if it.Error() != nil {
			return errors.Wrapf(
				err,
				"got iterator error when scanning assertion creations from block %d to %d",
				filterOpts.Start,
				*filterOpts.End,
			)
		}
		assertionHash := protocol.AssertionHash{Hash: it.Event.AssertionHash}

		// Try to confirm the assertion in the background.
		go m.keepTryingAssertionConfirmation(ctx, assertionHash)

		// Try to process the assertion creation event in the background
		// to not block the processing of other incoming events.
		go func() {
			_, processErr := retry.UntilSucceeds(ctx, func() (bool, error) {
				assertionCreationErr := m.ProcessAssertionCreationEvent(ctx, assertionHash)
				if assertionCreationErr != nil {
					log.Error(fmt.Sprintf("Could not process assertion creation event: %v", assertionCreationErr))
				}
				return true, assertionCreationErr
			}, retry.WithInterval(time.Minute))
			if processErr != nil {
				srvlog.Error(
					"Could not process assertion creation after retries",
					log.Ctx{"err": processErr},
				)
			}
		}()
	}
	return nil
}

// ProcessAssertionCreationEvent by checking if we agree with its claimed state.
// If we do not, we attempt to post a rival assertion along the fork and initiate a challenge
// if we are configured to do so. If we have not yet caught up to the claimed state,
// this function will then return an error.
func (m *Manager) ProcessAssertionCreationEvent(
	ctx context.Context,
	assertionHash protocol.AssertionHash,
) error {
	// Save the assertion creation event to the DB if possible.
	if err := m.saveAssertionToDB(ctx, assertionHash); err != nil {
		return err
	}
	// Ignore assertions we have submitted ourselves.
	if m.submittedAssertions.Has(assertionHash.Hash) {
		return nil
	}
	if assertionHash.Hash == (common.Hash{}) {
		return nil // Assertions cannot have a zero hash, not even genesis.
	}
	creationInfo, err := m.chain.ReadAssertionCreationInfo(ctx, assertionHash)
	if err != nil {
		return errors.Wrapf(err, "could not read assertion creation info for %#x", assertionHash.Hash)
	}
	if creationInfo.ParentAssertionHash == (common.Hash{}) {
		return nil // Skip processing genesis, as it has a parent assertion hash of 0x0.
	}
	// Check if we agree with the assertion's claimed state.
	claimedState := protocol.GoExecutionStateFromSolidity(creationInfo.AfterState)
	err = m.stateProvider.AgreesWithExecutionState(ctx, claimedState)
	switch {
	case errors.Is(err, l2stateprovider.ErrNoExecutionState):
		// If we disagree with the execution state, we should try to post the rival
		// assertion that we believe is correct and initiate a challenge if possible.
		if postRivalErr := m.postRivalAssertionAndChallenge(ctx, creationInfo); postRivalErr != nil {
			return postRivalErr
		}
		return nil
	case errors.Is(err, l2stateprovider.ErrChainCatchingUp):
		// Otherwise, we return the error that we are still catching up to the
		// execution state claimed by the assertion, and this function will be retried
		// by the caller if wrapped in a retryable call.
		chainCatchingUpCounter.Inc(1)
		return fmt.Errorf(
			"chain still catching up to processed execution state - "+
				"will reattempt assertion processing when caught up: %w",
			l2stateprovider.ErrChainCatchingUp,
		)
	case err != nil:
		return err
	}
	// If no error, this means we agree with the claimed assertion state
	// so there is no action to take.
	machineFinishedHash := crypto.Keccak256Hash([]byte("Machine finished:"), claimedState.GlobalState.Hash().Bytes())
	srvlog.Info("Agreed with incoming assertion", log.Ctx{
		"validatorName":       m.validatorName,
		"claimedState":        fmt.Sprintf("%+v", claimedState),
		"machineFinishedHash": machineFinishedHash,
		"assertionHash":       assertionHash,
	})
	m.assertionsProcessedCount++
	return nil
}

// Attempts to post a rival assertion to a given assertion and then attempts to
// open a challenge on that fork in the chain if configured to do so.
func (m *Manager) postRivalAssertionAndChallenge(
	ctx context.Context,
	creationInfo *protocol.AssertionCreatedInfo,
) error {
	m.assertionsProcessedCount++
	if !creationInfo.InboxMaxCount.IsUint64() {
		return errors.New("inbox max count not a uint64")
	}
	batchCount := creationInfo.InboxMaxCount.Uint64()
	claimedState := protocol.GoExecutionStateFromSolidity(creationInfo.AfterState)
	logFields := log.Ctx{
		"validatorName":         m.validatorName,
		"parentAssertionHash":   creationInfo.ParentAssertionHash,
		"detectedAssertionHash": creationInfo.AssertionHash,
		"batchCount":            batchCount,
		"claimedExecutionState": fmt.Sprintf("%+v", claimedState),
	}
	if !m.canPostRivalAssertion() {
		srvlog.Warn("Detected invalid assertion, but not configured to post a rival stake", logFields)
		evilAssertionCounter.Inc(1)
		return nil
	}

	latestAgreedWithAncestor, err := m.findLastAgreedWithAncestor(ctx, creationInfo)
	if err != nil {
		return err
	}
	// If the latest agreed with ancestor already has an honest child, we can exit early here.
	if m.assertionsWithHonestChild.Has(protocol.AssertionHash{Hash: latestAgreedWithAncestor.AssertionHash}) {
		return nil
	}

	srvlog.Info("Disagreed with execution state from observed assertion", logFields)
	evilAssertionCounter.Inc(1)

	// Post what we believe is the correct rival assertion that follows the ancestor we agree with.
	correctRivalAssertion, err := m.maybePostRivalAssertion(ctx, creationInfo, latestAgreedWithAncestor)
	if err != nil {
		return err
	}
	if correctRivalAssertion.IsNone() {
		srvlog.Warn(fmt.Sprintf("Expected to post a rival assertion to %#x, but did not post anything", creationInfo.AssertionHash))
		return nil
	}
	if !m.canPostChallenge() {
		srvlog.Warn("Attempted to post rival assertion and stake, but not configured to initiate a challenge", logFields)
		return nil
	}

	if creationInfo.ChallengeManager != m.challengeManagerAddr {
		var correctRivalAssertionCreatedInfo *protocol.AssertionCreatedInfo
		correctRivalAssertionCreatedInfo, err = m.chain.ReadAssertionCreationInfo(ctx, correctRivalAssertion.Unwrap().Id())
		if err != nil {
			return errors.Wrapf(err, "could not read assertion creation info for %#x", correctRivalAssertion.Unwrap().Id())
		}
		srvlog.Warn("Posted rival assertion, but could not challenge as challenge manager address did not match, "+
			"start a new server with the right challenge manager address", log.Ctx{
			"correctAssertion":                 correctRivalAssertionCreatedInfo.AssertionHash,
			"correctAssertionChallengeManager": correctRivalAssertionCreatedInfo.ChallengeManager,
			"evilAssertion":                    creationInfo.AssertionHash,
			"evilAssertionChallengeManager":    creationInfo.ChallengeManager,
			"expectedChallengeManager":         m.challengeManagerAddr,
		})
		return nil
	}

	// Generating a random integer between 0 and max delay second to wait before challenging.
	// This is to avoid all validators challenging at the same time.
	// TODO: Abstract into a smaller function.
	mds := 1 // default max delay seconds to 1 to avoid panic
	if m.challengeReader.MaxDelaySeconds() > 1 {
		mds = m.challengeReader.MaxDelaySeconds()
	}
	randSecs, err := randUint64(uint64(mds))
	if err != nil {
		return err
	}
	srvlog.Info("Waiting before submitting challenge on assertion", log.Ctx{"delay": randSecs})
	time.Sleep(time.Duration(randSecs) * time.Second)
	correctClaimedAssertionHash := correctRivalAssertion.Unwrap().Id()
	challengeSubmitted, err := m.challengeCreator.ChallengeAssertion(ctx, correctClaimedAssertionHash)
	if err != nil {
		return err
	}
	if challengeSubmitted {
		challengeSubmittedCounter.Inc(1)
		m.challengesSubmittedCount++
	}

	if err := m.logChallengeConfigs(ctx); err != nil {
		srvlog.Error("Could not log challenge configs", log.Ctx{"err": err})
	}

	return nil
}

func (m *Manager) logChallengeConfigs(ctx context.Context) error {
	cm, err := m.chain.SpecChallengeManager(ctx)
	if err != nil {
		return err
	}
	bigStepNum, err := cm.NumBigSteps(ctx)
	if err != nil {
		return err
	}
	challengePeriodBlocks, err := cm.ChallengePeriodBlocks(ctx)
	if err != nil {
		return err
	}
	layerZeroHeights, err := cm.LayerZeroHeights(ctx)
	if err != nil {
		return err
	}
	srvlog.Info("Challenge configs", log.Ctx{
		"address":               cm.Address(),
		"bigStepNumber":         bigStepNum,
		"challengePeriodBlocks": challengePeriodBlocks,
		"layerZeroHeights":      layerZeroHeights,
	})
	return nil
}

// Attempt to post a rival assertion based on the last agreed with ancestor
// of a given assertion.
//
// If this parent assertion already has a rival we agree with that arleady exists
// then this function will return that assertion.
func (m *Manager) maybePostRivalAssertion(
	ctx context.Context,
	creationInfo,
	latestAgreedWithAncestor *protocol.AssertionCreatedInfo,
) (option.Option[protocol.Assertion], error) {
	// Post what we believe is the correct assertion that follows the ancestor we agree with.
	staked, err := m.chain.IsStaked(ctx)
	if err != nil {
		return option.None[protocol.Assertion](), err
	}
	// If the validator is already staked, we post an assertion and move existing stake to it.
	var assertionOpt option.Option[protocol.Assertion]
	var postErr error
	if staked {
		assertionOpt, postErr = m.PostAssertionBasedOnParent(
			ctx, latestAgreedWithAncestor, m.chain.StakeOnNewAssertion,
		)
	} else {
		// Otherwise, we post a new assertion and place a new stake on it.
		assertionOpt, postErr = m.PostAssertionBasedOnParent(
			ctx, latestAgreedWithAncestor, m.chain.NewStakeOnNewAssertion,
		)
	}
	if postErr != nil {
		return option.None[protocol.Assertion](), postErr
	}
	if assertionOpt.IsSome() {
		m.submittedAssertions.Insert(assertionOpt.Unwrap().Id().Hash)
		m.assertionsWithHonestChild.Insert(protocol.AssertionHash{Hash: latestAgreedWithAncestor.AssertionHash})
		m.submittedRivalsCount++
		if err2 := m.saveAssertionToDB(ctx, assertionOpt.Unwrap().Id()); err2 != nil {
			return option.None[protocol.Assertion](), err2
		}
	}
	return assertionOpt, nil
}

func (m *Manager) saveAssertionToDB(ctx context.Context, assertionHash protocol.AssertionHash) error {
	if api.IsNil(m.apiDB) {
		return nil
	}
	creationInfo, err := m.chain.ReadAssertionCreationInfo(ctx, assertionHash)
	if err != nil {
		return err
	}
	beforeState := protocol.GoExecutionStateFromSolidity(creationInfo.BeforeState)
	afterState := protocol.GoExecutionStateFromSolidity(creationInfo.AfterState)
	status, err := m.chain.AssertionStatus(ctx, assertionHash)
	if err != nil {
		return err
	}
	assertion, err := m.chain.GetAssertion(ctx, assertionHash)
	if err != nil {
		return err
	}
	isFirstChild, err := assertion.IsFirstChild()
	if err != nil {
		return err
	}
	firstChildBlock, err := assertion.SecondChildCreationBlock()
	if err != nil {
		return err
	}
	secondChildBlock, err := assertion.SecondChildCreationBlock()
	if err != nil {
		return err
	}
	return m.apiDB.InsertAssertion(&api.JsonAssertion{
		Hash:                     assertionHash.Hash,
		ConfirmPeriodBlocks:      creationInfo.ConfirmPeriodBlocks,
		RequiredStake:            creationInfo.RequiredStake.String(),
		ParentAssertionHash:      creationInfo.ParentAssertionHash,
		InboxMaxCount:            creationInfo.InboxMaxCount.String(),
		AfterInboxBatchAcc:       creationInfo.AfterInboxBatchAcc,
		WasmModuleRoot:           creationInfo.WasmModuleRoot,
		ChallengeManager:         creationInfo.ChallengeManager,
		CreationBlock:            creationInfo.CreationBlock,
		TransactionHash:          creationInfo.TransactionHash,
		BeforeStateBlockHash:     beforeState.GlobalState.BlockHash,
		BeforeStateSendRoot:      beforeState.GlobalState.SendRoot,
		BeforeStateBatch:         beforeState.GlobalState.Batch,
		BeforeStatePosInBatch:    beforeState.GlobalState.PosInBatch,
		BeforeStateMachineStatus: beforeState.MachineStatus,
		AfterStateBlockHash:      afterState.GlobalState.BlockHash,
		AfterStateSendRoot:       afterState.GlobalState.SendRoot,
		AfterStateBatch:          afterState.GlobalState.Batch,
		AfterStatePosInBatch:     afterState.GlobalState.PosInBatch,
		AfterStateMachineStatus:  afterState.MachineStatus,
		FirstChildBlock:          &firstChildBlock,
		SecondChildBlock:         &secondChildBlock,
		IsFirstChild:             isFirstChild,
		Status:                   status.String(),
	})
}

// Look back until we find the ancestor we agree with for the given assertion.
func (m *Manager) findLastAgreedWithAncestor(
	ctx context.Context, assertionCreationInfo *protocol.AssertionCreatedInfo,
) (*protocol.AssertionCreatedInfo, error) {
	latestConfirmed, err := m.chain.LatestConfirmed(ctx)
	if err != nil {
		return nil, err
	}
	latestConfirmedInfo, err := m.chain.ReadAssertionCreationInfo(ctx, latestConfirmed.Id())
	if err != nil {
		return nil, err
	}
	agreedWithAncestor := latestConfirmed.Id().Hash
	cursor := assertionCreationInfo.ParentAssertionHash
	for cursor != agreedWithAncestor {
		// Get the cursor's creation info.
		parentCreationInfo, err := m.chain.ReadAssertionCreationInfo(
			ctx, protocol.AssertionHash{Hash: cursor},
		)
		if err != nil {
			return nil, err
		}
		parentExecState := protocol.GoExecutionStateFromSolidity(parentCreationInfo.AfterState)
		if err = m.stateProvider.AgreesWithExecutionState(ctx, parentExecState); err != nil {
			if errors.Is(err, l2stateprovider.ErrNoExecutionState) {
				// Disagreed with parent. This means we should look at the
				// grandparent and continue our loop.
				cursor = parentCreationInfo.ParentAssertionHash
				continue
			}
			return nil, err
		}
		// No error means we agree with this parent. We can break the loop.
		return parentCreationInfo, nil
	}
	return latestConfirmedInfo, nil
}

func (m *Manager) keepTryingAssertionConfirmation(ctx context.Context, assertionHash protocol.AssertionHash) {
	// Only resolve mode strategies or higher should be confirming assertions.
	if m.challengeReader.Mode() < types.ResolveMode {
		return
	}
	creationInfo, err := retry.UntilSucceeds(ctx, func() (*protocol.AssertionCreatedInfo, error) {
		return m.chain.ReadAssertionCreationInfo(ctx, assertionHash)
	})
	if err != nil {
		log.Error("Could not get assertion creation info", log.Ctx{"error": err})
		return
	}
	prevCreationInfo, err := retry.UntilSucceeds(ctx, func() (*protocol.AssertionCreatedInfo, error) {
		return m.chain.ReadAssertionCreationInfo(ctx, protocol.AssertionHash{Hash: creationInfo.ParentAssertionHash})
	})
	if err != nil {
		log.Error("Could not get prev assertion creation info", log.Ctx{"error": err})
		return
	}
	ticker := time.NewTicker(m.confirmationAttemptInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			parentAssertion, err := m.chain.GetAssertion(ctx, protocol.AssertionHash{Hash: creationInfo.ParentAssertionHash})
			if err != nil {
				log.Error("Could not get parent assertion", log.Ctx{"error": err})
				continue
			}
			parentAssertionHasSecondChild, err := parentAssertion.HasSecondChild()
			if err != nil {
				log.Error("Could not confirm if parent assertion has second child", log.Ctx{"error": err})
				continue
			}
			// Assertions that have a rival assertion cannot be confirmed by time.
			if parentAssertionHasSecondChild {
				return
			}
			confirmed, err := solimpl.TryConfirmingAssertion(ctx, creationInfo.AssertionHash, prevCreationInfo.ConfirmPeriodBlocks+creationInfo.CreationBlock, m.chain, m.averageTimeForBlockCreation, option.None[protocol.EdgeId]())
			if err != nil {
				srvlog.Error("Could not confirm assertion", log.Ctx{"err": err, "assertionHash": assertionHash.Hash})
				errorConfirmingAssertionByTimeCounter.Inc(1)
				continue
			}
			if confirmed {
				assertionConfirmedCounter.Inc(1)
				srvlog.Info("Confirmed assertion by time", log.Ctx{"assertionHash": creationInfo.AssertionHash})
				return
			}
		}
	}
}

// Returns true if the manager can respond to an assertion with a challenge.
func (m *Manager) canPostRivalAssertion() bool {
	return m.challengeReader.Mode() >= types.DefensiveMode
}

func (m *Manager) canPostChallenge() bool {
	return m.challengeReader.Mode() >= types.DefensiveMode
}

func (m *Manager) updateLatestConfirmedMetrics(ctx context.Context) {
	ticker := time.NewTicker(m.confirmationAttemptInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			latestConfirmed, err := m.chain.LatestConfirmed(ctx)
			if err != nil {
				srvlog.Debug("Could not fetch latest confirmed assertion", log.Ctx{"error": err})
				continue
			}
			latestConfirmedAssertionGauge.Update(int64(latestConfirmed.CreatedAtBlock()))
		case <-ctx.Done():
			return
		}
	}
}

func randUint64(max uint64) (uint64, error) {
	n, err := rand.Int(rand.Reader, new(big.Int).SetUint64(max))
	if err != nil {
		return 0, err
	}
	if !n.IsUint64() {
		return 0, errors.New("not a uint64")
	}
	return n.Uint64(), nil
}
