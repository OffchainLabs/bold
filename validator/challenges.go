package validator

import (
	"context"
	"fmt"

	"github.com/OffchainLabs/new-rollup-exploration/protocol"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Processes new challenge creation events from the protocol that were not created by self.
// This will fetch the challenge, its parent assertion, and create a challenge leaf that is
// relevant towards resolving the challenge. We then spawn a challenge tracker in the background.
func (v *Validator) onChallengeStarted(
	ctx context.Context, ev *protocol.StartChallengeEvent,
) error {
	if ev == nil {
		return nil
	}
	// Ignore challenges initiated by self.
	if isFromSelf(v.address, ev.Validator) {
		return nil
	}

	challenge, err := v.submitOrFetchProtocolChallenge(
		ctx,
		ev.ParentSeqNum,
		ev.ParentStateCommitment,
	)
	if err != nil {
		return err
	}

	// We then add a challenge vertex to the challenge.
	challengeVertex, err := v.addChallengeVertex(ctx, challenge)
	if err != nil {
		if errors.Is(err, protocol.ErrVertexAlreadyExists) {
			log.Infof(
				"Attempted to add a challenge leaf that already exists to challenge with "+
					"parent state commit: height=%d, stateRoot=%#x",
				challenge.ParentStateCommitment().Height,
				challenge.ParentStateCommitment().StateRoot,
			)
			return nil
		}
		return err
	}

	challengerName := "unknown-name"
	staker := challengeVertex.Challenger
	if name, ok := v.knownValidatorNames[staker]; ok {
		challengerName = name
	}
	log.WithFields(logrus.Fields{
		"name":                 v.name,
		"challenger":           challengerName,
		"challengingStateRoot": fmt.Sprintf("%#x", challenge.ParentStateCommitment().StateRoot),
		"challengingHeight":    challenge.ParentStateCommitment().Height,
	}).Warn("Received challenge for a created leaf, added own leaf with history commitment")

	// TODO: Start tracking the challenge.
	_ = challengeVertex

	return nil
}

// Initiates a challenge on an assertion added to the protocol by finding its parent assertion
// and starting a challenge transaction. If the challenge creation is successful, we add a leaf
// with an associated history commitment to it and spawn a challenge tracker in the background.
func (v *Validator) challengeAssertion(ctx context.Context, ev *protocol.CreateLeafEvent) error {
	challenge, err := v.submitOrFetchProtocolChallenge(ctx, ev.PrevSeqNum, ev.PrevStateCommitment)
	if err != nil {
		return err
	}

	// We then add a challenge vertex to the challenge.
	challengeVertex, err := v.addChallengeVertex(ctx, challenge)
	if err != nil {
		return err
	}
	if errors.Is(err, protocol.ErrVertexAlreadyExists) {
		log.Infof(
			"Attempted to add a challenge leaf that already exists to challenge with "+
				"parent state commit: height=%d, stateRoot=%#x",
			challenge.ParentStateCommitment().Height,
			challenge.ParentStateCommitment().StateRoot,
		)
		return nil
	}

	// TODO: Start tracking the challenge.
	_ = challengeVertex

	logFields := logrus.Fields{}
	logFields["name"] = v.name
	logFields["parentAssertionSeqNum"] = ev.PrevSeqNum
	logFields["parentAssertionStateRoot"] = fmt.Sprintf("%#x", ev.PrevStateCommitment.StateRoot)
	logFields["challengeID"] = fmt.Sprintf("%#x", ev.PrevStateCommitment.Hash())
	log.WithFields(logFields).Info("Successfully created challenge and added leaf, now tracking events")

	return nil
}

func (v *Validator) addChallengeVertex(
	ctx context.Context,
	challenge *protocol.Challenge,
) (*protocol.ChallengeVertex, error) {
	historyCommit, err := v.stateManager.LatestHistoryCommitment(ctx)
	if err != nil {
		return nil, err
	}
	var challengeVertex *protocol.ChallengeVertex
	if err = v.chain.Tx(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		numAssertions := p.NumAssertions(tx)
		currentAssertion, readErr := p.AssertionBySequenceNum(tx, protocol.SequenceNum(numAssertions-1))
		if readErr != nil {
			return readErr
		}
		challengeVertex, err = challenge.AddLeaf(tx, currentAssertion, historyCommit, v.address)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(
			err,
			"could add challenge vertex to challenge with parent state commitment: height=%d, stateRoot=%#x",
			challenge.ParentStateCommitment().Height,
			challenge.ParentStateCommitment().StateRoot,
		)
	}
	return challengeVertex, nil
}

// Tries to submit a challenge to the protocol or retrieve it if it already exists.
// based on the parent assertion's state commitment hash.
func (v *Validator) submitOrFetchProtocolChallenge(
	ctx context.Context,
	parentAssertionSeqNum protocol.SequenceNum,
	parentAssertionCommit protocol.StateCommitment,
) (*protocol.Challenge, error) {
	var challenge *protocol.Challenge
	var err error
	err = v.chain.Tx(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		parentAssertion, readErr := p.AssertionBySequenceNum(tx, parentAssertionSeqNum)
		if readErr != nil {
			return readErr
		}
		challenge, err = parentAssertion.CreateChallenge(tx, ctx, v.address)
		if err != nil {
			return errors.Wrap(err, "cannot make challenge")
		}
		return nil
	})
	switch {
	case errors.Is(err, protocol.ErrChallengeAlreadyExists):
		log.Info("Challenge on leaf already exists, reading existing challenge from protocol")
		if err = v.chain.Call(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
			challenge, err = p.ChallengeByAssertionStateCommit(
				tx,
				protocol.AssertionStateCommitHash(parentAssertionCommit.Hash()),
			)
			if err != nil {
				return errors.Wrap(err, "cannot make challenge")
			}
			return nil
		}); err != nil {
			return nil, errors.Wrap(err, "could not get challenge by ID")
		}
	case err != nil:
		return nil, errors.Wrapf(
			err,
			"could not initiate challenge on assertion with state commit: height=%d and stateRoot=%#x",
			parentAssertionCommit.Height,
			parentAssertionCommit.StateRoot,
		)
	default:
	}
	if challenge == nil {
		return nil, errors.New("got nil challenge from protocol")
	}
	return challenge, nil
}
