package validator

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	solimpl "github.com/OffchainLabs/challenge-protocol-v2/protocol/sol-implementation"
	statemanager "github.com/OffchainLabs/challenge-protocol-v2/state-manager"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	watcher "github.com/OffchainLabs/challenge-protocol-v2/validator/chain-watcher"
	"github.com/OffchainLabs/challenge-protocol-v2/validator/challenge-tree"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (et *edgeTracker) uniqueTrackerLogFields() logrus.Fields {
	startHeight, startCommit := et.edge.StartCommitment()
	endHeight, endCommit := et.edge.EndCommitment()
	id := et.edge.Id()
	return logrus.Fields{
		"id":            util.Trunc(id[:]),
		"startHeight":   startHeight,
		"startCommit":   util.Trunc(startCommit.Bytes()),
		"endHeight":     endHeight,
		"endCommit":     util.Trunc(endCommit.Bytes()),
		"validatorName": et.cfg.validatorName,
		"challengeType": et.edge.GetType(),
		"address":       util.Trunc(et.cfg.validatorAddress.Bytes()),
	}
}

func (et *edgeTracker) act(ctx context.Context) error {
	fields := et.uniqueTrackerLogFields()
	current := et.fsm.Current()
	switch current.State {
	// Start state.
	case edgeStarted:
		canOsp, err := canOneStepProve(et.edge)
		if err != nil {
			log.WithFields(fields).WithError(err).Error("could not check if edge can be one step proven")
			return et.fsm.Do(edgeBackToStart{})
		}
		if canOsp {
			return et.fsm.Do(edgeHandleOneStepProof{})
		}
		wasConfirmed, err := et.tryToConfirm(ctx)
		if err != nil {
			return errors.Wrap(err, "could not check if can confirm by time")
		}
		if wasConfirmed {
			return et.fsm.Do(edgeConfirm{})
		}
		hasRival, err := et.edge.HasRival(ctx)
		if err != nil {
			return errors.Wrap(err, "could not check presumptive")
		}
		if !hasRival {
			return et.fsm.Do(edgeBackToStart{})
		}
		atOneStepFork, err := et.edge.HasLengthOneRival(ctx)
		if err != nil {
			log.WithFields(fields).WithError(err).Error("could not check if edge has length one rival")
			return et.fsm.Do(edgeBackToStart{})
		}
		if atOneStepFork {
			return et.fsm.Do(edgeHandleOneStepFork{})
		}
		return et.fsm.Do(edgeBisect{})
	// Edge is the source of a one-step-fork.
	case edgeAtOneStepFork:
		startHeight, startCommit := et.edge.StartCommitment()
		log.WithFields(fields).Infof(
			"Reached one-step-fork at start height %d and start history commitment %s",
			startHeight,
			util.Trunc(startCommit.Bytes()),
		)
		return et.fsm.Do(edgeOpenSubchallengeLeaf{})
	// Edge is at a one-step-proof in a small-step challenge.
	case edgeAtOneStepProof:
		if err := et.submitOneStepProof(ctx); err != nil {
			if strings.Contains(err.Error(), "Edge not pending") || strings.Contains(err.Error(), "machine executed to end step") {
				return et.fsm.Do(edgeConfirm{})
			}
			log.WithFields(fields).WithError(err).Error("could not submit one step proof")
			return et.fsm.Do(edgeBackToStart{})
		}
		return et.fsm.Do(edgeConfirm{})
	// Edge tracker should add a subchallenge level zero leaf.
	case edgeAddingSubchallengeLeaf:
		if err := et.openSubchallengeLeaf(ctx); err != nil {
			log.WithFields(fields).WithError(err).Error("could not open subchallenge leaf")
			return et.fsm.Do(edgeBackToStart{})
		}
		return et.fsm.Do(edgeBackToStart{})
	// Edge should bisect.
	case edgeBisecting:
		lowerChild, upperChild, err := et.bisect(ctx)
		if err != nil {
			if errors.Is(err, solimpl.ErrAlreadyExists) {
				return et.fsm.Do(edgeBackToStart{})
			}
			log.WithError(err).WithFields(fields).Error("Could not bisect")
			return et.fsm.Do(edgeBackToStart{})
		}
		firstTracker, err := newEdgeTracker(
			ctx,
			et.cfg,
			lowerChild,
			et.startBlockHeight,
			et.topLevelClaimEndBatchCount,
		)
		if err != nil {
			log.WithError(err).WithFields(fields).Error("Could not create new edge tracker")
			return et.fsm.Do(edgeBackToStart{})
		}
		secondTracker, err := newEdgeTracker(
			ctx,
			et.cfg,
			upperChild,
			et.startBlockHeight,
			et.topLevelClaimEndBatchCount,
		)
		if err != nil {
			log.WithError(err).WithFields(fields).Error("Could not create new edge tracker")
			return et.fsm.Do(edgeBackToStart{})
		}
		go firstTracker.spawn(ctx)
		go secondTracker.spawn(ctx)
		return et.fsm.Do(edgeBackToStart{})
	case edgeConfirmed:
		log.WithFields(fields).Info("Edge reached confirmed state")
		return et.fsm.Do(edgeConfirm{})
	default:
		return fmt.Errorf("invalid state: %s", current.State)
	}
}

func (et *edgeTracker) childrenAreConfirmed(
	ctx context.Context,
	chalManager protocol.SpecChallengeManager,
) (bool, error) {
	lower, err := et.edge.LowerChild(ctx)
	if err != nil {
		return false, err
	}
	upper, err := et.edge.UpperChild(ctx)
	if err != nil {
		return false, err
	}
	if lower.IsNone() || upper.IsNone() {
		return false, nil
	}
	someLowerEdge, err := chalManager.GetEdge(ctx, lower.Unwrap())
	if err != nil {
		return false, err
	}
	someUpperEdge, err := chalManager.GetEdge(ctx, upper.Unwrap())
	if err != nil {
		return false, err
	}
	if someLowerEdge.IsNone() || someUpperEdge.IsNone() {
		return false, nil
	}
	lowerStatus, err := someLowerEdge.Unwrap().Status(ctx)
	if err != nil {
		return false, err
	}
	upperStatus, err := someUpperEdge.Unwrap().Status(ctx)
	if err != nil {
		return false, err
	}
	return lowerStatus == protocol.EdgeConfirmed && upperStatus == protocol.EdgeConfirmed, nil
}

func (et *edgeTracker) tryToConfirm(ctx context.Context) (bool, error) {
	status, err := et.edge.Status(ctx)
	if err != nil {
		return false, err
	}
	if status == protocol.EdgeConfirmed {
		return true, nil
	}
	prevAssertionId, err := et.edge.PrevAssertionId(ctx)
	if err != nil {
		return false, err
	}
	manager, err := et.cfg.chain.SpecChallengeManager(ctx)
	if err != nil {
		return false, err
	}

	// Check if we can confirm by children.
	childrenConfirmed, err := et.childrenAreConfirmed(ctx, manager)
	if err != nil {
		return false, err
	}
	if childrenConfirmed {
		if err := et.edge.ConfirmByChildren(ctx); err != nil {
			return false, err
		}
		log.WithFields(et.uniqueTrackerLogFields()).Info("Confirmed by children")
		return true, nil
	}

	// Check if we can confirm by claim.
	edgeWithClaimConfirmed, err := et.cfg.chainWatcher.ConfirmedEdgeWithClaimExists(
		prevAssertionId,
		protocol.ClaimId(et.edge.Id()),
	)
	if err != nil {
		return false, err
	}
	if edgeWithClaimConfirmed {
		if err := et.edge.ConfirmByClaim(ctx, protocol.ClaimId(et.edge.Id())); err != nil {
			return false, err
		}
		log.WithFields(et.uniqueTrackerLogFields()).Info("Confirmed by claim")
		return true, nil
	}

	// Check if we can confirm by time.
	timer, ancestors, err := et.cfg.chainWatcher.ComputeHonestPathTimer(ctx, prevAssertionId, et.edge.Id())
	if err != nil {
		return false, err
	}
	chalPeriod, err := manager.ChallengePeriodBlocks(ctx)
	start, _ := et.edge.StartCommitment()
	end, _ := et.edge.EndCommitment()
	if et.cfg.validatorName == "alice" && start == 0 && end == 2 && et.edge.GetType() == protocol.SmallStepChallengeEdge {
		log.WithFields(et.uniqueTrackerLogFields()).Infof("Path timer %d", timer)
	}
	if et.cfg.validatorName == "bob" && start == 16 && end == 32 && et.edge.GetType() == protocol.BlockChallengeEdge {
		log.WithFields(et.uniqueTrackerLogFields()).Infof("Bob Path timer %d", timer)
	}
	if et.cfg.validatorName == "alice" && start == 16 && end == 32 && et.edge.GetType() == protocol.BlockChallengeEdge {
		log.WithFields(et.uniqueTrackerLogFields()).Infof("Alice Path timer %d", timer)
	}
	if timer >= challengetree.PathTimer(chalPeriod) {
		if err := et.edge.ConfirmByTimer(ctx, ancestors); err != nil {
			return false, err
		}
		log.WithFields(et.uniqueTrackerLogFields()).Info("Confirmed by time")
		return true, nil
	}
	return false, nil
}

// Determines the bisection point from parentHeight to toHeight and returns a history
// commitment with a prefix proof for the action based on the challenge type.
func (et *edgeTracker) determineBisectionHistoryWithProof(
	ctx context.Context,
) (util.HistoryCommitment, []byte, error) {
	startHeight, _ := et.edge.StartCommitment()
	endHeight, _ := et.edge.EndCommitment()
	bisectTo, err := util.BisectionPoint(uint64(startHeight), uint64(endHeight))
	if err != nil {
		return util.HistoryCommitment{}, nil, errors.Wrapf(err, "determining bisection point failed for %d and %d", startHeight, endHeight)
	}
	if et.edge.GetType() == protocol.BlockChallengeEdge {
		historyCommit, commitErr := et.cfg.stateManager.HistoryCommitmentUpToBatch(ctx, et.startBlockHeight, et.startBlockHeight+bisectTo, et.topLevelClaimEndBatchCount)
		if commitErr != nil {
			return util.HistoryCommitment{}, nil, commitErr
		}
		proof, proofErr := et.cfg.stateManager.PrefixProofUpToBatch(ctx, et.startBlockHeight, bisectTo, uint64(endHeight), et.topLevelClaimEndBatchCount)
		if proofErr != nil {
			return util.HistoryCommitment{}, nil, proofErr
		}
		return historyCommit, proof, nil
	}
	var historyCommit util.HistoryCommitment
	var commitErr error
	var proof []byte
	var proofErr error

	originHeights, err := et.edge.TopLevelClaimHeight(ctx)
	if err != nil {
		return util.HistoryCommitment{}, nil, err
	}

	fromAssertionHeight := uint64(originHeights.BlockChallengeOriginHeight)
	toAssertionHeight := fromAssertionHeight + 1

	switch et.edge.GetType() {
	case protocol.BigStepChallengeEdge:
		historyCommit, commitErr = et.cfg.stateManager.BigStepCommitmentUpTo(ctx, fromAssertionHeight, toAssertionHeight, bisectTo)
		proof, proofErr = et.cfg.stateManager.BigStepPrefixProof(ctx, fromAssertionHeight, toAssertionHeight, bisectTo, uint64(endHeight))
	case protocol.SmallStepChallengeEdge:
		fromBigStep := uint64(originHeights.BigStepChallengeOriginHeight)
		toBigStep := fromBigStep + 1

		historyCommit, commitErr = et.cfg.stateManager.SmallStepCommitmentUpTo(ctx, fromAssertionHeight, toAssertionHeight, fromBigStep, toBigStep, bisectTo)
		proof, proofErr = et.cfg.stateManager.SmallStepPrefixProof(ctx, fromAssertionHeight, toAssertionHeight, fromBigStep, toBigStep, bisectTo, uint64(endHeight))
	default:
		return util.HistoryCommitment{}, nil, fmt.Errorf("unsupported challenge type: %s", et.edge.GetType())
	}
	if commitErr != nil {
		return util.HistoryCommitment{}, nil, errors.Wrap(commitErr, "could not produce history commitment")
	}
	if proofErr != nil {
		return util.HistoryCommitment{}, nil, errors.Wrap(proofErr, "could not produce prefix proof")
	}
	return historyCommit, proof, nil
}

func (et *edgeTracker) bisect(ctx context.Context) (protocol.SpecEdge, protocol.SpecEdge, error) {
	hasChildren, err := et.edge.HasChildren(ctx)
	if err != nil {
		return nil, nil, err
	}
	if hasChildren {
		return nil, nil, solimpl.ErrAlreadyExists
	}
	historyCommit, proof, err := et.determineBisectionHistoryWithProof(ctx)
	if err != nil {
		return nil, nil, err
	}
	endHeight, endCommit := et.edge.EndCommitment()
	bisectTo := historyCommit.Height
	firstChild, secondChild, err := et.edge.Bisect(ctx, historyCommit.Merkle, proof)
	if err != nil {
		return nil, nil, errors.Wrapf(
			err,
			"%s could not bisect to height=%d,commit=%s from height=%d,commit=%s",
			et.cfg.validatorName,
			bisectTo,
			util.Trunc(historyCommit.Merkle.Bytes()),
			endHeight,
			util.Trunc(endCommit.Bytes()),
		)
	}
	log.WithFields(logrus.Fields{
		"name":               et.cfg.validatorName,
		"challengeType":      et.edge.GetType(),
		"bisectedFrom":       endHeight,
		"bisectedFromMerkle": util.Trunc(endCommit.Bytes()),
		"bisectedTo":         bisectTo,
		"bisectedToMerkle":   util.Trunc(historyCommit.Merkle.Bytes()),
	}).Info("Successfully bisected edge")
	return firstChild, secondChild, nil
}

func (et *edgeTracker) openSubchallengeLeaf(ctx context.Context) error {
	originHeights, err := et.edge.TopLevelClaimHeight(ctx)
	if err != nil {
		return errors.Wrap(err, "could not get top level claim height")
	}

	fromAssertionHeight := uint64(originHeights.BlockChallengeOriginHeight)
	toAssertionHeight := fromAssertionHeight + 1

	startHeight, _ := et.edge.StartCommitment()
	endHeight, _ := et.edge.EndCommitment()

	fields := logrus.Fields{
		"name":                et.cfg.validatorName,
		"edgeStartHeight":     startHeight,
		"edgeEndHeight":       endHeight,
		"fromAssertionHeight": fromAssertionHeight,
		"toAssertionHeight":   toAssertionHeight,
	}

	var startHistory util.HistoryCommitment
	var endHistory util.HistoryCommitment
	var startParentCommitment util.HistoryCommitment
	var endParentCommitment util.HistoryCommitment
	var startEndPrefixProof []byte
	switch et.edge.GetType() {
	case protocol.BlockChallengeEdge:
		log.WithFields(fields).Info("Big step leaf commit")
		fromBlock := fromAssertionHeight + et.startBlockHeight
		toBlock := toAssertionHeight + et.startBlockHeight
		startHistory, err = et.cfg.stateManager.BigStepCommitmentUpTo(ctx, fromBlock, toBlock, 0)
		if err != nil {
			return err
		}
		endHistory, err = et.cfg.stateManager.BigStepLeafCommitment(ctx, fromBlock, toBlock)
		if err != nil {
			return err
		}
		startParentCommitment, err = et.cfg.stateManager.HistoryCommitmentUpToBatch(ctx, et.startBlockHeight, fromBlock, et.topLevelClaimEndBatchCount)
		if err != nil {
			return err
		}
		endParentCommitment, err = et.cfg.stateManager.HistoryCommitmentUpToBatch(ctx, et.startBlockHeight, toBlock, et.topLevelClaimEndBatchCount)
		if err != nil {
			return err
		}
		startEndPrefixProof, err = et.cfg.stateManager.BigStepPrefixProof(ctx, fromBlock, toBlock, 0, endHistory.Height)
		if err != nil {
			return err
		}
	case protocol.BigStepChallengeEdge:
		log.WithFields(fields).Info("Small step leaf commit")
		fromBlock := fromAssertionHeight + et.startBlockHeight
		toBlock := toAssertionHeight + et.startBlockHeight
		startHistory, err = et.cfg.stateManager.SmallStepCommitmentUpTo(ctx, fromBlock, toBlock, uint64(startHeight), uint64(endHeight), 0)
		if err != nil {
			return err
		}
		endHistory, err = et.cfg.stateManager.SmallStepLeafCommitment(ctx, fromBlock, toBlock, uint64(startHeight), uint64(endHeight))
		if err != nil {
			return err
		}
		startParentCommitment, err = et.cfg.stateManager.BigStepCommitmentUpTo(ctx, fromBlock, toBlock, uint64(startHeight))
		if err != nil {
			return err
		}
		endParentCommitment, err = et.cfg.stateManager.BigStepCommitmentUpTo(ctx, fromBlock, toBlock, uint64(endHeight))
		if err != nil {
			return err
		}
		startEndPrefixProof, err = et.cfg.stateManager.SmallStepPrefixProof(ctx, fromBlock, toBlock, uint64(startHeight), uint64(endHeight), 0, endHistory.Height)
		if err != nil {
			return err
		}
	default:
		return errors.New("unsupported subchallenge type for creating leaf commitment")
	}
	manager, err := et.cfg.chain.SpecChallengeManager(ctx)
	if err != nil {
		return err
	}
	addedLeaf, err := manager.AddSubChallengeLevelZeroEdge(
		ctx,
		et.edge,
		startHistory,
		endHistory,
		startParentCommitment.LastLeafProof,
		endParentCommitment.LastLeafProof,
		startEndPrefixProof,
	)
	if err != nil {
		return err
	}
	fields["firstLeaf"] = util.Trunc(startHistory.FirstLeaf.Bytes())
	fields["endHeight"] = endHistory.Height
	fields["startCommitment"] = util.Trunc(startHistory.Merkle.Bytes())
	fields["subChallengeType"] = addedLeaf.GetType()
	log.WithFields(fields).Info("Added subchallenge level zero edge, now tracking it")
	tracker, err := newEdgeTracker(
		ctx,
		et.cfg,
		addedLeaf,
		et.startBlockHeight,
		et.topLevelClaimEndBatchCount,
	)
	if err != nil {
		return err
	}
	go tracker.spawn(ctx)
	return nil
}

func (et *edgeTracker) submitOneStepProof(ctx context.Context) error {
	fields := et.uniqueTrackerLogFields()
	log.WithFields(fields).Info("Submitting one-step-proof to protocol")
	originHeights, err := et.edge.TopLevelClaimHeight(ctx)
	if err != nil {
		return errors.Wrap(err, "could not get top level claim height")
	}
	fromAssertionHeight := uint64(originHeights.BlockChallengeOriginHeight)
	toAssertionHeight := fromAssertionHeight + 1
	fromBigStep := uint64(originHeights.BigStepChallengeOriginHeight)
	toBigStep := fromBigStep + 1
	pc, _ := et.edge.StartCommitment()

	prevAssertionId, err := et.edge.PrevAssertionId(ctx)
	if err != nil {
		return err
	}
	prevAssertionNum, err := et.cfg.chain.GetAssertionNum(ctx, prevAssertionId)
	if err != nil {
		return err
	}
	parentAssertionCreationInfo, err := et.cfg.chain.ReadAssertionCreationInfo(ctx, prevAssertionNum)
	if err != nil {
		return err
	}
	data, beforeStateInclusionProof, afterStateInclusionProof, err := et.cfg.stateManager.OneStepProofData(
		ctx,
		parentAssertionCreationInfo,
		fromAssertionHeight,
		toAssertionHeight,
		fromBigStep,
		toBigStep,
		uint64(pc),
		uint64(pc)+1,
	)
	if err != nil {
		return err
	}

	manager, err := et.cfg.chain.SpecChallengeManager(ctx)
	if err != nil {
		return err
	}
	if err = manager.ConfirmEdgeByOneStepProof(
		ctx,
		et.edge.Id(),
		data,
		beforeStateInclusionProof,
		afterStateInclusionProof,
	); err != nil {
		return errors.Wrap(err, "could not confirm one step proof against protocol")
	}
	log.WithFields(fields).Info("Succeeded one-step-proof for edge and confirmed it as winner")
	return nil
}

type edgeTrackerConfig struct {
	actEveryNSeconds time.Duration
	timeRef          util.TimeReference
	chain            protocol.Protocol
	stateManager     statemanager.Manager
	validatorName    string
	validatorAddress common.Address
	chainWatcher     *watcher.Watcher
}

type edgeTracker struct {
	cfg                        *edgeTrackerConfig
	edge                       protocol.SpecEdge
	fsm                        *util.Fsm[edgeTrackerAction, edgeTrackerState]
	startBlockHeight           uint64
	topLevelClaimEndBatchCount uint64
}

func newEdgeTracker(
	_ context.Context,
	cfg *edgeTrackerConfig,
	edge protocol.SpecEdge,
	startHeightOffset uint64,
	topLevelClaimEndBatchCount uint64,
	fsmOpts ...util.FsmOpt[edgeTrackerAction, edgeTrackerState],
) (*edgeTracker, error) {
	fsmOpts = append(fsmOpts, util.WithTrackedTransitions[edgeTrackerAction, edgeTrackerState]())
	fsm, err := newEdgeTrackerFsm(
		edgeStarted,
		fsmOpts...,
	)
	if err != nil {
		return nil, err
	}
	return &edgeTracker{
		cfg:                        cfg,
		edge:                       edge,
		fsm:                        fsm,
		startBlockHeight:           startHeightOffset,
		topLevelClaimEndBatchCount: topLevelClaimEndBatchCount,
	}, nil
}

func (et *edgeTracker) spawn(ctx context.Context) {
	fields := et.uniqueTrackerLogFields()
	log.WithFields(fields).Info("Tracking edge")

	t := et.cfg.timeRef.NewTicker(et.cfg.actEveryNSeconds)
	defer t.Stop()
	for {
		select {
		case <-t.C():
			if et.shouldComplete() {
				log.WithFields(fields).Infof("Edge tracker received notice of a confirmation, exiting")
				return
			}
			if err := et.act(ctx); err != nil {
				log.Error(err)
			}
		case <-ctx.Done():
			log.WithFields(fields).Debug("Edge tracker goroutine exiting")
			return
		}
	}
}

func (et *edgeTracker) shouldComplete() bool {
	return et.fsm.Current().State == edgeConfirmed
}

func canOneStepProve(edge protocol.SpecEdge) (bool, error) {
	start, _ := edge.StartCommitment()
	end, _ := edge.EndCommitment()
	// Can never happen in the protocol, but added as an additional defensive check.
	if start >= end {
		return false, fmt.Errorf("start height %d cannot be >= end height %d", start, end)
	}
	return end-start == 1 && edge.GetType() == protocol.SmallStepChallengeEdge, nil
}
