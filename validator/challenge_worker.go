package validator

import (
	"context"

	"github.com/OffchainLabs/new-rollup-exploration/protocol"
	"github.com/OffchainLabs/new-rollup-exploration/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type challengeWorker struct {
	challenge          *protocol.Challenge
	validatorAddress   common.Address
	validatorName      string
	createdVertices    util.ThreadSafeSlice[*protocol.ChallengeVertex]
	reachedOneStepFork chan struct{}
}

// Performs the actions required by a validator when a ChallengeEvent is fired during
// a block challenge. The basic algorith is as follows:
//
// 1. We fetch our last created vertex that is higher than the commitment of the vertex seen
// during the challenge event.
// 2. If we have the same history commitment, and the vertex is not ours, we make a merge move.
// 3. While we are not presumptive, we keep trying to merge until we are presumptive or we reach a one-step-fork.
func (w *challengeWorker) actOnBlockChallenge(
	ctx context.Context,
	validator *Validator,
	eventActor common.Address,
	eventHistoryCommit util.HistoryCommitment,
	eventSequenceNum protocol.VertexSequenceNumber,
) error {
	if isFromSelf(w.validatorAddress, eventActor) {
		return nil
	}
	if w.createdVertices.Empty() {
		return nil
	}
	// Go down the tree to find the first vertex we created that has a commitment height >
	// the vertex seen from the merge event.
	vertexToActUpon := w.createdVertices.Last().Unwrap()
	numVertices := w.createdVertices.Len()
	for i := numVertices - 1; i > 0; i-- {
		vertex := w.createdVertices.Get(i).Unwrap()
		if vertex.Commitment.Height > eventHistoryCommit.Height {
			vertexToActUpon = vertex
			break
		}
	}

	mergedToOurs := eventHistoryCommit.Hash() == vertexToActUpon.Commitment.Hash()
	if mergedToOurs {
		log.WithFields(logrus.Fields{
			"name":                w.validatorName,
			"mergedHeight":        eventHistoryCommit.Height,
			"mergedHistoryMerkle": eventHistoryCommit.Merkle,
		}).Info("Other validator merged to our vertex")
	}

	// Make a merge move.
	if validator.stateManager.HasHistoryCommitment(ctx, eventHistoryCommit) && !mergedToOurs {
		challengeID := protocol.CommitHash(w.challenge.ParentStateCommitment().Hash())
		if err := validator.merge(ctx, challengeID, vertexToActUpon, eventSequenceNum); err != nil {
			// TODO: Find a better way to exit if a merge is invalid than showing a scary log to the user.
			// Validators currently try to make merge moves they should not during the challenge game.
			if errors.Is(err, protocol.ErrInvalidOp) {
				return nil
			}
			return errors.Wrap(err, "failed to merge")
		}
	}

	hasPresumptiveSuccessor := vertexToActUpon.IsPresumptiveSuccessor()
	currentVertex := vertexToActUpon

	// While we do not have the presumptive successor, we keep trying to bisect and
	// break from the loop if reach a one step fork.
	for !hasPresumptiveSuccessor {
		if currentVertex.Commitment.Height == currentVertex.Prev.Commitment.Height+1 {
			w.reachedOneStepFork <- struct{}{}
			break
		}
		bisectedVertex, err := validator.bisect(ctx, currentVertex)
		if err != nil {
			// TODO: Find another way of cleanly ending the bisection process so that we do not
			// end on a scary "state did not allow this operation" log.
			if errors.Is(err, protocol.ErrWrongState) {
				log.WithError(err).Debug("State incorrect for bisection")
				return nil
			}
			if errors.Is(err, protocol.ErrVertexAlreadyExists) {
				return nil
			}
			return err
		}
		currentVertex = bisectedVertex
		w.createdVertices.Append(currentVertex)
	}
	return nil
}
