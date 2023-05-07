package validator

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
)

// Sync challenges from last finalized assertion to current assertion.
// If an assertion is not the first child of parent, call `syncChallengeVertex`.
func (v *Validator) syncChallenges(ctx context.Context) error {
	numberOfAssertions, err := v.chain.NumAssertions(ctx)
	if err != nil {
		return err
	}
	latestConfirmedAssertion, err := v.chain.LatestConfirmed(ctx)
	if err != nil {
		return err
	}
	for i := latestConfirmedAssertion.SeqNum() + 1; i < protocol.AssertionSequenceNumber(numberOfAssertions); i++ {
		a, err := v.chain.AssertionBySequenceNum(ctx, i)
		if err != nil {
			log.WithError(err).Error("failed to get assertion")
			continue
		}
		firstChild, err := a.IsFirstChild()
		if err != nil {
			log.WithError(err).Error("failed to get first child")
			continue
		}
		if firstChild {
			continue
		}
		if err := v.syncChallengeVertex(ctx, a); err != nil {
			log.WithError(err).Error("failed to sync challenge vertex")
		}
	}
	return nil
}

// Sync a top level challenge edge, if success, spin  up a new edge tracker routine.
// - Get the parent assertion ID
// - Get the start and end commitments
// - Calculate top level edge ID
// - Get the edge
// - Initialize an edge tracker
// - Start the edge tracker routine
func (v *Validator) syncChallengeVertex(ctx context.Context, a protocol.Assertion) error {
	n, err := a.PrevSeqNum()
	if err != nil {
		return err
	}
	cm, err := v.chain.SpecChallengeManager(ctx)
	if err != nil {
		return err
	}
	startCommit, err := v.stateManager.HistoryCommitmentUpTo(ctx, 0)
	if err != nil {
		return err
	}
	prevCreationInfo, err := v.chain.ReadAssertionCreationInfo(ctx, n)
	if err != nil {
		return err
	}
	endCommit, err := v.stateManager.HistoryCommitmentUpToBatch(
		ctx,
		0,
		protocol.LevelZeroBlockEdgeHeight,
		prevCreationInfo.InboxMaxCount.Uint64(),
	)
	if err != nil {
		return err
	}
	parentAssertionId, err := v.chain.GetAssertionId(ctx, n)
	if err != nil {
		return err
	}
	edgeId, err := cm.CalculateEdgeId(
		ctx,
		protocol.BlockChallengeEdge,
		protocol.OriginId(parentAssertionId),
		protocol.Height(startCommit.Height),
		startCommit.Merkle,
		protocol.Height(endCommit.Height),
		endCommit.Merkle,
	)
	if err != nil {
		return err
	}
	edge, err := cm.GetEdge(ctx, edgeId)
	if err != nil {
		return err
	}
	assertionPrevHeight, ok := v.stateManager.ExecutionStateBlockHeight(ctx, protocol.GoExecutionStateFromSolidity(prevCreationInfo.AfterState))
	if !ok {
		return fmt.Errorf("missing previous assertion %v after execution %+v in local state manager", n, prevCreationInfo.AfterState)
	}
	tracker, err := newEdgeTracker(
		ctx,
		&edgeTrackerConfig{
			timeRef:          v.timeRef,
			actEveryNSeconds: v.edgeTrackerWakeInterval,
			chain:            v.chain,
			stateManager:     v.stateManager,
			validatorName:    v.name,
			validatorAddress: v.address,
		},
		edge.Unwrap(),
		assertionPrevHeight,
		prevCreationInfo.InboxMaxCount.Uint64(),
	)
	if err != nil {
		return err
	}
	go tracker.spawn(ctx)
	return nil
}

func (v *Validator) pollForAssertions(ctx context.Context) {
	ticker := time.NewTicker(v.newAssertionCheckInterval)
	defer ticker.Stop()
	var nextAssertion protocol.AssertionSequenceNumber

	var onLeafCreatedLock sync.Mutex

	for {
		select {
		case <-ticker.C:
			onLeafCreatedLock.Lock()

			numberOfAssertions, err := v.chain.NumAssertions(ctx)
			if err != nil {
				log.Error(err)
				continue
			}
			latestConfirmedAssertion, err := v.chain.LatestConfirmed(ctx)
			if err != nil {
				log.Error(err)
				continue
			}
			latestConfirmedSeqNum := latestConfirmedAssertion.SeqNum()
			if latestConfirmedSeqNum >= nextAssertion {
				nextAssertion = latestConfirmedSeqNum + 1
			}

			for ; nextAssertion < protocol.AssertionSequenceNumber(numberOfAssertions); nextAssertion++ {
				assertion, err := v.chain.AssertionBySequenceNum(ctx, nextAssertion)
				if err != nil {
					log.Error(err)
					continue
				}
				if err := v.onLeafCreated(ctx, assertion); err != nil {
					log.Error(err)
				}
			}

			onLeafCreatedLock.Unlock()
		case <-ctx.Done():
			return
		}
	}
}
