package validator

import (
	"context"

	"github.com/OffchainLabs/new-rollup-exploration/protocol"
	"github.com/OffchainLabs/new-rollup-exploration/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Performs a bisection move during a BlockChallenge in the assertion protocol given
// a validator challenge vertex.
func (v *Validator) bisect(
	ctx context.Context,
	validatorChallengeVertex *protocol.ChallengeVertex,
) (*protocol.ChallengeVertex, error) {

	// TODO: Use optionals for safety around nil.
	toHeight := validatorChallengeVertex.Commitment.Height
	parentHeight := validatorChallengeVertex.Prev.Commitment.Height

	bisectTo, err := util.BisectionPoint(parentHeight, toHeight)
	if err != nil {
		return nil, err
	}
	historyCommit, err := v.stateManager.HistoryCommitmentUpTo(ctx, bisectTo)
	if err != nil {
		return nil, err
	}
	proof, err := v.stateManager.PrefixProof(ctx, bisectTo, toHeight)
	if err != nil {
		return nil, err
	}
	if err = util.VerifyPrefixProof(historyCommit, validatorChallengeVertex.Commitment, proof); err != nil {
		return nil, err
	}
	// Otherwise, we must bisect to our own historical commitment and produce
	// a proof of the vertex we want to bisect to.
	var bisectedVertex *protocol.ChallengeVertex
	err = v.chain.Tx(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		bisectedVertex, err = validatorChallengeVertex.Bisect(tx, historyCommit, proof, v.address)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"could not bisect vertex with sequence %d and challenger %#x to height %d with history %d and %#x",
			validatorChallengeVertex.SequenceNum,
			validatorChallengeVertex.Challenger,
			bisectTo,
			historyCommit.Height,
			historyCommit.Merkle,
		)
	}
	log.WithFields(logrus.Fields{
		"name":                   v.name,
		"IsPresumptiveSuccessor": bisectedVertex.IsPresumptiveSuccessor(),
	}).Infof(
		"Successfully bisected to vertex with height %d and commit %#x",
		bisectedVertex.Commitment.Height,
		bisectedVertex.Commitment.Merkle,
	)
	return bisectedVertex, nil
}

// Performs a merge move during a BlockChallenge in the assertion protocol given
// a challenge vertex and the sequence number we should be merging into. To do this, we
// also need to fetch vertex we are are merging to by reading it from the protocol.
func (v *Validator) merge(
	ctx context.Context,
	challenge *protocol.Challenge,
	validatorChallengeVertex *protocol.ChallengeVertex,
	newPrevSeqNum protocol.SequenceNum,
) error {
	var mergingTo *protocol.ChallengeVertex
	var err error
	err = v.chain.Call(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		id := protocol.AssertionStateCommitHash(challenge.ParentStateCommitment().Hash())
		mergingTo, err = p.ChallengeVertexBySequenceNum(tx, id, newPrevSeqNum)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "could not read challenge vertex from protocol")
	}
	mergingToHeight := mergingTo.Commitment.Height
	historyCommit, err := v.stateManager.HistoryCommitmentUpTo(ctx, mergingToHeight)
	if err != nil {
		return err
	}
	currentCommit := validatorChallengeVertex.Commitment
	proof, err := v.stateManager.PrefixProof(ctx, mergingToHeight, currentCommit.Height)
	if err != nil {
		return err
	}
	if err := util.VerifyPrefixProof(historyCommit, currentCommit, proof); err != nil {
		return err
	}
	if err := v.chain.Tx(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		return validatorChallengeVertex.Merge(tx, mergingTo, proof, v.address)
	}); err != nil {
		return errors.Wrapf(
			err,
			"could not merge vertex with height %d and commit %#x to height %x and commit %#x",
			currentCommit.Height,
			currentCommit.Merkle,
			mergingToHeight,
			mergingTo.Commitment.Merkle,
		)
	}
	log.WithFields(logrus.Fields{
		"name": v.name,
	}).Infof(
		"Successfully merged to vertex with height %d and commit %#x",
		mergingTo.Commitment.Height,
		mergingTo.Commitment.Merkle,
	)
	return nil
}
