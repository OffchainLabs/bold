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
	levelZeroEdge, creationInfo, err := v.addBlockChallengeLevelZeroEdge(ctx, assertion)
	if err != nil {
		return fmt.Errorf("could not add block challenge level zero edge %v: %w", v.name, err)
	}
	assertionHeight, ok := v.stateManager.ExecutionStateBlockHeight(
		ctx,
		protocol.GoExecutionStateFromSolidity(creationInfo.AfterState),
	)
	if !ok {
		return fmt.Errorf(
			"missing previous assertion %v after execution %+v in local state manager",
			id,
			creationInfo.AfterState,
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
		assertionHeight,
		creationInfo.InboxMaxCount.Uint64(),
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
) (protocol.SpecEdge, *protocol.AssertionCreatedInfo, error) {
	creationInfo, err := v.chain.ReadAssertionCreationInfo(ctx, assertion.Id())
	if err != nil {
		return nil, nil, errors.Wrap(err, "could not get assertion creation info")
	}
	startCommit, err := v.stateManager.HistoryCommitmentUpTo(ctx, 0)
	if err != nil {
		return nil, nil, err
	}
	endCommit, err := v.stateManager.HistoryCommitmentUpToBatch(
		ctx,
		0,
		protocol.LevelZeroBlockEdgeHeight,
		creationInfo.InboxMaxCount.Uint64(),
	)
	if err != nil {
		return nil, nil, err
	}

	computed, err := inclusionproofs.CalculateRootFromProof(endCommit.LastLeafProof, protocol.LevelZeroBlockEdgeHeight, endCommit.LastLeaf)
	if err != nil {
		return nil, nil, err
	}
	if computed != endCommit.Merkle {
		return nil, nil, fmt.Errorf("got %#x, wanted %#x", computed, endCommit.Merkle)
	}
	log.Infof("Passing inclusion checks: %v", v.name)

	startEndPrefixProof, err := v.stateManager.PrefixProofUpToBatch(
		ctx,
		0,
		0,
		protocol.LevelZeroBlockEdgeHeight,
		creationInfo.InboxMaxCount.Uint64(),
	)
	if err != nil {
		return nil, nil, err
	}
	manager, err := v.chain.SpecChallengeManager(ctx)
	if err != nil {
		return nil, nil, err
	}
	edge, err := manager.AddBlockChallengeLevelZeroEdge(ctx, assertion, startCommit, endCommit, startEndPrefixProof)
	if err != nil {
		return nil, nil, err
	}
	return edge, creationInfo, nil
}
