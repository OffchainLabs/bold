package validator

import (
	"context"
	"fmt"

	"github.com/OffchainLabs/new-rollup-exploration/protocol"
	"github.com/OffchainLabs/new-rollup-exploration/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Performs a bisection move during a BlockChallenge in the assertion protocol given
// a validator challenge vertex. It will create a historical commitment for the vertex
// the validator wants to bisect to and an associated proof for submitting to the protocol.
func (v *Validator) bisect(
	ctx context.Context,
	validatorChallengeVertex *protocol.ChallengeVertex,
) (*protocol.ChallengeVertex, error) {
	currentHeight := validatorChallengeVertex.Commitment.Height
	parentHeight := validatorChallengeVertex.Prev.Commitment.Height
	bisectTo, err := util.BisectionPoint(parentHeight, currentHeight)
	if err != nil {
		return nil, errors.Wrapf(err, "determining bisection point failed for %d and %d", parentHeight, currentHeight)
	}
	historyCommit := validatorChallengeVertex.Commitment
	if err := v.verifyPrefixProofWithHeights(ctx, historyCommit, bisectTo, currentHeight); err != nil {
		return nil, err
	}

	var bisectedVertex *protocol.ChallengeVertex
	err = v.chain.Tx(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		proof, err := v.stateManager.PrefixProof(ctx, bisectTo, currentHeight)
		if err != nil {
			return err
		}
		bisectedVertex, err = validatorChallengeVertex.Bisect(tx, historyCommit, proof, v.address)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		if errors.Is(protocol.ErrVertexAlreadyExists, err) {
			log.Infof(
				"Bisected vertex with height %d and commit %#x already exists",
				historyCommit.Height,
				historyCommit.Merkle,
			)
			return nil, nil
		}
		return nil, errors.Wrapf(
			err,
			"could not bisect vertex with sequence %d and validator %#x to height %d with history %d and %#x",
			validatorChallengeVertex.SequenceNum,
			validatorChallengeVertex.Challenger,
			bisectTo,
			historyCommit.Height,
			historyCommit.Merkle,
		)
	}
	log.WithFields(logrus.Fields{
		"name":                   v.name,
		"isPresumptiveSuccessor": bisectedVertex.IsPresumptiveSuccessor(),
		"historyCommitHeight":    bisectedVertex.Commitment.Height,
		"historyCommitMerkle":    fmt.Sprintf("%#x", bisectedVertex.Commitment.Height),
	}).Info("Successfully bisected to vertex")
	return bisectedVertex, nil
}

// Performs a merge move during a BlockChallenge in the assertion protocol given
// a challenge vertex and the sequence number we should be merging into. To do this, we
// also need to fetch vertex we are are merging to by reading it from the protocol.
func (v *Validator) merge(
	ctx context.Context,
	mergingTo *protocol.ChallengeVertex,
	mergingFrom *protocol.ChallengeVertex,
) error {
	mergingToHeight := mergingTo.Commitment.Height
	currentCommit := mergingFrom.Commitment
	if err := v.verifyPrefixProofWithHeights(ctx, currentCommit, mergingTo.Commitment.Height, currentCommit.Height); err != nil {
		return err
	}

	if err := v.chain.Tx(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		proof, err := v.stateManager.PrefixProof(ctx, mergingToHeight, currentCommit.Height)
		if err != nil {
			return err
		}
		return mergingFrom.Merge(tx, mergingTo, proof, v.address)
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

// verifies prefix proofs with heights and commitment.
func (v *Validator) verifyPrefixProofWithHeights(ctx context.Context, commitment util.HistoryCommitment, fromHeight uint64, toHeight uint64) error {
	historyCommit, err := v.stateManager.HistoryCommitmentUpTo(ctx, fromHeight)
	if err != nil {
		return errors.Wrapf(err, "could not rertieve history commitment up to height %d", fromHeight)
	}
	proof, err := v.stateManager.PrefixProof(ctx, fromHeight, toHeight)
	if err != nil {
		return errors.Wrapf(err, "generating prefix proof failed from height %d to %d", fromHeight, toHeight)
	}
	// Perform an extra safety check to ensure our proof verifies against the specified commitment
	// before we make an on-chain transaction.
	if err = util.VerifyPrefixProof(historyCommit, commitment, proof); err != nil {
		return errors.Wrapf(
			err,
			"prefix proof failed to verify for commit %+v to commit %+v",
			historyCommit,
			commitment,
		)
	}
	return nil
}
