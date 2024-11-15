// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE

package challengemanager

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/log"

	protocol "github.com/offchainlabs/bold/chain-abstraction"
	edgetracker "github.com/offchainlabs/bold/challenge-manager/edge-tracker"
	"github.com/offchainlabs/bold/containers"
	"github.com/offchainlabs/bold/containers/option"
	l2stateprovider "github.com/offchainlabs/bold/layer2-state-provider"
)

// HandleCorrectRival is called when the assertion manager has posted a correct
// rival assertion on the chain and the chllenge manager needs to create a
// challenge committing to the correct assertion to rival one or more incorrect
// assertions.
func (m *Manager) HandleCorrectRival(ctx context.Context, riv protocol.AssertionHash) error {
	// Generate a random integer between 0 and max delay seconds to wait before
	// challenging.
	// This is to avoid all validators challenging at the same time.
	mds := 1 // default max delay seconds to 1 to avoid panic
	if m.MaxDelaySeconds() > 1 {
		mds = m.MaxDelaySeconds()
	}
	randSecs, err := randUint64(uint64(mds))
	if err != nil {
		return err
	}
	time.Sleep(time.Duration(randSecs) * time.Second)
	challengeSubmitted, err := m.ChallengeAssertion(ctx, riv)
	if err != nil {
		return err
	}
	if challengeSubmitted {
		challengeSubmittedCounter.Inc(1)
	}
	if err := m.logChallengeConfigs(ctx); err != nil {
		log.Error("Could not log challenge configs", "err", err)
	}
	return nil
}

// ChallengeAssertion initiates a challenge committing to an assertion added to
// the protocol by finding its parent assertion and starting a challenge
// transaction. If the challenge creation is successful, the challenge manager
// adds a leaf with an associated history commitment to it and spawns a
// challenge tracker in the background.
//
// id is the id of the assertion that this validator agrees with.
func (m *Manager) ChallengeAssertion(ctx context.Context, id protocol.AssertionHash) (bool, error) {
	assertion, err := m.chain.GetAssertion(ctx, &bind.CallOpts{Context: ctx}, id)
	if err != nil {
		return false, errors.Wrapf(err, "could not get assertion to challenge with id %#x", id)
	}
	if m.claimedAssertionsInChallenge.Has(id) {
		log.Debug(fmt.Sprintf("Already challenged assertion with id %#x, skipping", id.Hash))
		return false, nil
	}
	assertionStatus, err := m.chain.AssertionStatus(ctx, assertion.Id())
	if err != nil {
		return false, errors.Wrapf(err, "could not get assertion status with id %#x", id)
	}
	if assertionStatus == protocol.AssertionConfirmed {
		log.Info("Skipping challenge submission on already confirmed assertion", "assertionHash", id.Hash)
		return false, nil
	}
	// We then add a level zero edge to initiate a challenge.
	levelZeroEdge, shouldTrack, edgeTrackerAssertionInfo, alreadyExists, err := m.addBlockChallengeLevelZeroEdge(ctx, assertion)
	if err != nil {
		return false, fmt.Errorf("could not add block challenge level zero edge %v: %w", m.name, err)
	}
	if !shouldTrack {
		log.Info("Challenge not in list of specified challenges to track, skipping", "assertionHash", id.Hash)
		return false, nil
	}
	log.Info("Opening a challenge on an observed assertion",
		"assertionHash", id.Hash,
		"validatorName", m.name,
	)
	if alreadyExists {
		log.Info("Challenge on assertion already exists, now tracking it locally", "assertionHash", id.Hash)
		m.claimedAssertionsInChallenge.Insert(id)
		return false, nil
	}
	if verifiedErr := m.watcher.AddVerifiedHonestEdge(ctx, levelZeroEdge); verifiedErr != nil {
		fields := []any{
			"edgeId", levelZeroEdge.Id(),
			"err", verifiedErr,
		}
		log.Error("could not add verified honest edge to chain watcher", fields...)
	}
	// Start tracking the challenge.
	tracker, err := edgetracker.New(
		ctx,
		levelZeroEdge,
		m.chain,
		m.stateManager,
		m.watcher,
		m,
		edgeTrackerAssertionInfo,
		edgetracker.WithTimeReference(m.timeRef),
		edgetracker.WithValidatorName(m.name),
	)
	if err != nil {
		return false, err
	}
	m.LaunchThread(tracker.Spawn)

	log.Info("Successfully opened a challenge on an invalid assertion",
		"name", m.name,
		"assertionHash", containers.Trunc(id.Bytes()),
		"fromBatch", edgeTrackerAssertionInfo.FromState.Batch,
		"fromPosInBatch", edgeTrackerAssertionInfo.FromState.PosInBatch,
		"batchLimit", edgeTrackerAssertionInfo.BatchLimit,
	)
	return true, nil
}

func (m *Manager) addBlockChallengeLevelZeroEdge(
	ctx context.Context,
	assertion protocol.Assertion,
) (protocol.VerifiedRoyalEdge, bool, *l2stateprovider.AssociatedAssertionMetadata, bool, error) {
	creationInfo, err := m.chain.ReadAssertionCreationInfo(ctx, assertion.Id())
	if err != nil {
		return nil, false, nil, false, errors.Wrap(err, "could not get assertion creation info")
	}
	if !m.watcher.AllowTrackingEdgeWithParentHash(protocol.AssertionHash{Hash: creationInfo.ParentAssertionHash}) {
		return nil, false, nil, false, nil
	}
	prevCreationInfo, err := m.chain.ReadAssertionCreationInfo(ctx, protocol.AssertionHash{Hash: creationInfo.ParentAssertionHash})
	if err != nil {
		return nil, false, nil, false, errors.Wrap(err, "could not get assertion creation info")
	}
	if prevCreationInfo.InboxMaxCount == nil {
		return nil, false, nil, false, errors.New("prevCreationInfo.InboxMaxCount is nil")
	}
	if !prevCreationInfo.InboxMaxCount.IsUint64() {
		return nil, false, nil, false, fmt.Errorf("inbox max count is not a uint64: %v", prevCreationInfo.InboxMaxCount)
	}
	fromState := protocol.GoGlobalStateFromSolidity(creationInfo.BeforeState.GlobalState)
	assertionMetadata := &l2stateprovider.AssociatedAssertionMetadata{
		FromState:            fromState,
		BatchLimit:           l2stateprovider.Batch(prevCreationInfo.InboxMaxCount.Uint64()),
		WasmModuleRoot:       prevCreationInfo.WasmModuleRoot,
		ClaimedAssertionHash: creationInfo.AssertionHash,
	}

	startCommit, err := m.stateManager.HistoryCommitment(
		ctx,
		&l2stateprovider.HistoryCommitmentRequest{
			AssertionMetadata:           assertionMetadata,
			UpperChallengeOriginHeights: []l2stateprovider.Height{},
			UpToHeight:                  option.Some(l2stateprovider.Height(0)),
		},
	)
	if err != nil {
		return nil, false, nil, false, err
	}
	manager := m.chain.SpecChallengeManager()
	layerZeroHeights, err := manager.LayerZeroHeights(ctx)
	if err != nil {
		return nil, false, nil, false, err
	}
	req := &l2stateprovider.HistoryCommitmentRequest{
		AssertionMetadata:           assertionMetadata,
		UpperChallengeOriginHeights: []l2stateprovider.Height{},
		UpToHeight:                  option.Some(l2stateprovider.Height(layerZeroHeights.BlockChallengeHeight)),
	}
	endCommit, err := m.stateManager.HistoryCommitment(
		ctx,
		req,
	)
	if err != nil {
		return nil, false, nil, false, err
	}
	precomputedEdgeId, err := manager.CalculateEdgeId(
		ctx,
		protocol.NewBlockChallengeLevel(),
		protocol.OriginId(creationInfo.ParentAssertionHash),
		protocol.Height(startCommit.Height),
		startCommit.Merkle,
		protocol.Height(endCommit.Height),
		endCommit.Merkle,
	)
	if err != nil {
		return nil, false, nil, false, errors.Wrap(err, "could not calculate edge id")
	}
	someLevelZeroEdge, err := manager.GetEdge(ctx, precomputedEdgeId)

	// If the edge already exists, we return true and everything else nil.
	if err == nil && !someLevelZeroEdge.IsNone() {
		return nil, true, nil, true, nil
	}
	startEndPrefixProof, err := m.stateManager.PrefixProof(
		ctx,
		req,
		l2stateprovider.Height(0),
	)
	if err != nil {
		return nil, false, nil, false, err
	}
	edge, err := manager.AddBlockChallengeLevelZeroEdge(ctx, assertion, startCommit, endCommit, startEndPrefixProof)
	if err != nil {
		return nil, false, nil, false, errors.Wrap(err, "could not post block challenge root edge")
	}
	return edge, true, assertionMetadata, false, nil
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
