package validator

import (
	"context"
	"fmt"

	protocol "github.com/OffchainLabs/challenge-protocol-v2/chain-abstraction"
	"github.com/OffchainLabs/challenge-protocol-v2/containers"
	inclusionproofs "github.com/OffchainLabs/challenge-protocol-v2/state-commitments/inclusion-proofs"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// ChallengeCreator defines a struct which can initiate a challenge on an assertion id
// by creating a level zero, block challenge edge onchain.
type ChallengeCreator interface {
	ChallengeAssertion(ctx context.Context, id protocol.AssertionId) error
}

// Initiates a challenge on an assertion added to the protocol by finding its parent assertion
// and starting a challenge transaction. If the challenge creation is successful, we add a leaf
// with an associated history commitment to it and spawn a challenge tracker in the background.
func (v *Manager) ChallengeAssertion(ctx context.Context, id protocol.AssertionId) error {
	assertion, err := v.chain.GetAssertion(ctx, id)
	if err != nil {
		return errors.Wrapf(err, "could not get assertion to challenge with id %#x", id)
	}

	// We then add a level zero edge to initiate a challenge.
	levelZeroEdge, err := v.addBlockChallengeLevelZeroEdge(ctx, assertion)
	if err != nil {
		return fmt.Errorf("could not add block challenge level zero eddge %v: %w", v.name, err)
	}

	prevCreationInfo, err := v.chain.ReadAssertionCreationInfo(ctx, id)
	if err != nil {
		return err
	}
	assertionPrevHeight, ok := v.stateManager.ExecutionStateBlockHeight(
		ctx,
		protocol.GoExecutionStateFromSolidity(prevCreationInfo.AfterState),
	)
	if !ok {
		return fmt.Errorf(
			"missing previous assertion %v after execution %+v in local state manager",
			id,
			prevCreationInfo.AfterState,
		)
	}

	// Start tracking the challenge.
	tracker, err := newEdgeTracker(
		ctx,
		&edgeTrackerConfig{
			timeRef:          v.timeRef,
			actEveryNSeconds: v.edgeTrackerWakeInterval,
			chain:            v.chain,
			stateManager:     v.stateManager,
			validatorName:    v.name,
			validatorAddress: v.address,
			chainWatcher:     v.watcher,
		},
		levelZeroEdge,
		assertionPrevHeight,
		prevCreationInfo.InboxMaxCount.Uint64(),
	)
	if err != nil {
		return err
	}
	go tracker.spawn(ctx)

	logFields := logrus.Fields{}
	logFields["name"] = v.name
	logFields["assertionId"] = containers.Trunc(id[:])
	log.WithFields(logFields).Info("Successfully created level zero edge for block challenge")
	return nil
}

func (v *Manager) addBlockChallengeLevelZeroEdge(
	ctx context.Context,
	assertion protocol.Assertion,
) (protocol.SpecEdge, error) {
	prevId := assertion.PrevId()
	prevCreationInfo, err := v.chain.ReadAssertionCreationInfo(ctx, prevId)
	if err != nil {
		return nil, errors.Wrap(err, "could not get assertion creation info")
	}
	startCommit, err := v.stateManager.HistoryCommitmentUpTo(ctx, 0)
	if err != nil {
		return nil, err
	}
	endCommit, err := v.stateManager.HistoryCommitmentUpToBatch(
		ctx,
		0,
		protocol.LevelZeroBlockEdgeHeight,
		prevCreationInfo.InboxMaxCount.Uint64(),
	)
	if err != nil {
		return nil, err
	}

	computed, err := inclusionproofs.CalculateRootFromProof(endCommit.LastLeafProof, protocol.LevelZeroBlockEdgeHeight, endCommit.LastLeaf)
	if err != nil {
		return nil, err
	}
	if computed != endCommit.Merkle {
		return nil, fmt.Errorf("got %#x, wanted %#x", computed, endCommit.Merkle)
	}
	log.Infof("Passing inclusion checks: %v", v.name)

	startEndPrefixProof, err := v.stateManager.PrefixProofUpToBatch(
		ctx,
		0,
		0,
		protocol.LevelZeroBlockEdgeHeight,
		prevCreationInfo.InboxMaxCount.Uint64(),
	)
	if err != nil {
		return nil, err
	}
	manager, err := v.chain.SpecChallengeManager(ctx)
	if err != nil {
		return nil, err
	}
	return manager.AddBlockChallengeLevelZeroEdge(ctx, assertion, startCommit, endCommit, startEndPrefixProof)
}
