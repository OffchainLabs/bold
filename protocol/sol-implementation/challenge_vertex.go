package solimpl

import (
	"context"
	"github.com/OffchainLabs/challenge-protocol-v2/solgen/go/challengeV2gen"
	"math/big"
	"strings"

	"fmt"
	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"time"
)

func (v *ChallengeVertex) Id() [32]byte {
	return v.id
}

func (v *ChallengeVertex) SequenceNum() protocol.VertexSequenceNumber {
	return 0
}

func (v *ChallengeVertex) Prev(ctx context.Context, tx protocol.ActiveTx) (util.Option[*challengeV2gen.ChallengeVertex], error) {
	manager, err := v.assertionChain.CurrentChallengeManager(ctx, tx)
	if err != nil {
		return util.None[*challengeV2gen.ChallengeVertex](), err
	}
	vertex, err := v.fetchChallengeVertex(ctx, tx)
	if err != nil {
		return util.None[*challengeV2gen.ChallengeVertex](), err
	}
	return manager.GetVertex(ctx, tx, vertex.PredecessorId)
}

// Fetches the vertex from chain.
func (v *ChallengeVertex) fetchChallengeVertex(ctx context.Context, tx protocol.ActiveTx) (*challengeV2gen.ChallengeVertex, error) {
	manager, err := v.assertionChain.CurrentChallengeManager(ctx, tx)
	if err != nil {
		return nil, err
	}
	vertex, err := manager.GetVertex(ctx, tx, v.id)
	if err != nil {
		return nil, err
	}
	if vertex.IsNone() {
		return nil, ErrNotFound
	}
	return vertex.Unwrap(), nil
}

func (v *ChallengeVertex) Status(ctx context.Context, tx protocol.ActiveTx) (protocol.AssertionState, error) {
	// TODO: Should be vertex status.
	vertex, err := v.fetchChallengeVertex(ctx, tx)
	if err != nil {
		return 0, err
	}
	return protocol.AssertionState(vertex.Status), nil
}

func (v *ChallengeVertex) HistoryCommitment(ctx context.Context, tx protocol.ActiveTx) (util.HistoryCommitment, error) {
	vertex, err := v.fetchChallengeVertex(ctx, tx)
	if err != nil {
		return util.HistoryCommitment{}, err
	}
	return util.HistoryCommitment{
		Height: vertex.Height.Uint64(),
		Merkle: vertex.HistoryRoot,
	}, nil
}

func (v *ChallengeVertex) MiniStaker(ctx context.Context, tx protocol.ActiveTx) (common.Address, error) {
	vertex, err := v.fetchChallengeVertex(ctx, tx)
	if err != nil {
		return common.Address{}, err
	}
	return vertex.Staker, nil
}

func (v *ChallengeVertex) GetSubChallenge(ctx context.Context, tx protocol.ActiveTx) (util.Option[protocol.Challenge], error) {
	return util.None[protocol.Challenge](), errors.New("unimplemented")
}

func (v *ChallengeVertex) EligibleForNewSuccessor(ctx context.Context, tx protocol.ActiveTx) (bool, error) {
	return false, errors.New("unimplemented")
}

func (v *ChallengeVertex) PresumptiveSuccessor(
	ctx context.Context, tx protocol.ActiveTx,
) (util.Option[protocol.ChallengeVertex], error) {
	return util.None[protocol.ChallengeVertex](), errors.New("unimplemented")
}

func (v *ChallengeVertex) PsTimer(ctx context.Context, tx protocol.ActiveTx) (uint64, error) {
	return 0, errors.New("unimplemented")
}

func (v *ChallengeVertex) ChessClockExpired(
	ctx context.Context,
	tx protocol.ActiveTx,
	challengePeriodSeconds time.Duration,
) (bool, error) {
	return false, errors.New("unimplemented")
}

func (v *ChallengeVertex) ConfirmForChallengeDeadline(ctx context.Context, tx protocol.ActiveTx) error {
	return errors.New("unimplemented")
}

func (v *ChallengeVertex) ConfirmForSubChallengeWin(ctx context.Context, tx protocol.ActiveTx) error {
	return errors.New("unimplemented")
}

// HasConfirmedSibling checks if the vertex has a confirmed sibling in the protocol.
func (v *ChallengeVertex) HasConfirmedSibling(ctx context.Context, tx protocol.ActiveTx) (bool, error) {
	manager, err := v.assertionChain.CurrentChallengeManager(ctx, tx)
	if err != nil {
		return false, err
	}
	caller, err := manager.GetCaller(ctx, tx)
	if err != nil {
		return false, err
	}
	return caller.HasConfirmedSibling(v.assertionChain.callOpts, v.id)
}

// IsPresumptiveSuccessor checks if a vertex is the presumptive successor
// within its challenge.
func (v *ChallengeVertex) IsPresumptiveSuccessor(ctx context.Context, tx protocol.ActiveTx) (bool, error) {
	manager, err := v.assertionChain.CurrentChallengeManager(ctx, tx)
	if err != nil {
		return false, err
	}
	caller, err := manager.GetCaller(ctx, tx)
	if err != nil {
		return false, err
	}
	return caller.IsPresumptiveSuccessor(v.assertionChain.callOpts, v.id)
}

// ChildrenAreAtOneStepFork checks if child vertices are at a one-step-fork in the challenge
// it is contained in.
func (v *ChallengeVertex) ChildrenAreAtOneStepFork(ctx context.Context, tx protocol.ActiveTx) (bool, error) {
	manager, err := v.assertionChain.CurrentChallengeManager(ctx, tx)
	if err != nil {
		return false, err
	}
	caller, err := manager.GetCaller(ctx, tx)
	if err != nil {
		return false, err
	}
	atFork, err := caller.ChildrenAreAtOneStepFork(v.assertionChain.callOpts, v.id)
	if err != nil {
		errS := err.Error()
		switch {
		case strings.Contains(errS, "Lowest height not one above"):
			return false, nil
		case strings.Contains(errS, "Has presumptive successor"):
			return false, nil
		default:
			return false, err
		}
	}
	return atFork, nil
}

// Merge a challenge vertex to another by providing its history
// commitment and a prefix proof.
func (v *ChallengeVertex) Merge(
	ctx context.Context,
	tx protocol.ActiveTx,
	mergingToHistory util.HistoryCommitment,
	proof []common.Hash,
) (protocol.ChallengeVertex, error) {
	// Flatten the last leaf proof for submission to the chain.
	flatProof := make([]byte, 0)
	for _, h := range proof {
		flatProof = append(flatProof, h[:]...)
	}
	manager, err := v.assertionChain.CurrentChallengeManager(ctx, tx)
	if err != nil {
		return nil, err
	}
	writer, err := manager.GetWriter(ctx, tx)
	if err != nil {
		return nil, err
	}
	_, err = transact(ctx, v.assertionChain.backend, v.assertionChain.headerReader, func() (*types.Transaction, error) {
		return writer.Merge(
			v.assertionChain.txOpts,
			v.id,
			mergingToHistory.Merkle,
			flatProof,
		)
	})
	if err != nil {
		return nil, err
	}
	vertex, err := v.fetchChallengeVertex(ctx, tx)
	if err != nil {
		return nil, err
	}
	return getVertexFromComponents(
		ctx,
		tx,
		manager,
		v.assertionChain,
		vertex.ChallengeId,
		mergingToHistory,
	)
}

// Bisect a challenge vertex by providing a history commitment.
func (v *ChallengeVertex) Bisect(
	ctx context.Context,
	tx protocol.ActiveTx,
	history util.HistoryCommitment,
	proof []common.Hash,
) (protocol.ChallengeVertex, error) {
	// Flatten the last leaf proof for submission to the chain.
	flatProof := make([]byte, 0)
	for _, h := range proof {
		flatProof = append(flatProof, h[:]...)
	}
	manager, err := v.assertionChain.CurrentChallengeManager(ctx, tx)
	if err != nil {
		return nil, err
	}
	writer, err := manager.GetWriter(ctx, tx)
	if err != nil {
		return nil, err
	}
	_, err = transact(ctx, v.assertionChain.backend, v.assertionChain.headerReader, func() (*types.Transaction, error) {
		return writer.Bisect(
			v.assertionChain.txOpts,
			v.id,
			history.Merkle,
			flatProof,
		)
	})
	if err != nil {
		errS := err.Error()
		switch {
		case strings.Contains(errS, "Bisection vertex already exists"):
			return nil, ErrAlreadyExists
		default:
			return nil, err
		}
	}
	vertex, err := v.fetchChallengeVertex(ctx, tx)
	if err != nil {
		return nil, err
	}
	return getVertexFromComponents(
		ctx,
		tx,
		manager,
		v.assertionChain,
		vertex.ChallengeId,
		history,
	)
}

func getVertexFromComponents(
	ctx context.Context,
	tx protocol.ActiveTx,
	manager protocol.ChallengeManager,
	assertionChain *AssertionChain,
	challengeId [32]byte,
	history util.HistoryCommitment,
) (protocol.ChallengeVertex, error) {
	caller, err := manager.GetCaller(ctx, tx)
	if err != nil {
		return nil, err
	}
	vertexId, err := caller.CalculateChallengeVertexId(
		assertionChain.callOpts,
		challengeId,
		history.Merkle,
		big.NewInt(int64(history.Height)),
	)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &ChallengeVertex{
		id:             vertexId,
		assertionChain: assertionChain,
	}, nil
}

func (v *ChallengeVertex) ConfirmForPsTimer(ctx context.Context, tx protocol.ActiveTx) error {
	manager, err := v.assertionChain.CurrentChallengeManager(ctx, tx)
	if err != nil {
		return err
	}
	writer, err := manager.GetWriter(ctx, tx)
	if err != nil {
		return err
	}
	_, err = transact(ctx, v.assertionChain.backend, v.assertionChain.headerReader, func() (*types.Transaction, error) {
		return writer.ConfirmForPsTimer(
			v.assertionChain.txOpts,
			v.id,
		)
	})
	if err == nil {
		return nil
	}
	switch {
	case strings.Contains(err.Error(), "PsTimer not greater than challenge period"):
		return errors.Wrapf(ErrPsTimerNotYet, "vertex id %#v", v.id)
	default:
		return err
	}
}

func (v *ChallengeVertex) CreateSubChallenge(ctx context.Context, tx protocol.ActiveTx) (protocol.Challenge, error) {
	manager, err := v.assertionChain.CurrentChallengeManager(ctx, tx)
	if err != nil {
		return nil, err
	}
	vertex, err := v.fetchChallengeVertex(ctx, tx)
	if err != nil {
		return nil, err
	}
	currentChallenge, err := manager.GetChallenge(ctx, tx, vertex.ChallengeId)
	if err != nil {
		return nil, err
	}
	if currentChallenge.IsNone() {
		return nil, errors.New("no challenge exists found for vertex")
	}
	challenge := currentChallenge.Unwrap()
	var subChallengeType protocol.ChallengeType
	switch protocol.ChallengeType(challenge.ChallengeType) {
	case protocol.BlockChallenge:
		subChallengeType = protocol.BigStepChallenge
	case protocol.BigStepChallenge:
		subChallengeType = protocol.SmallStepChallenge
	default:
		return nil, fmt.Errorf("cannot make subchallenge for challenge type %d", challenge.ChallengeType)
	}

	writer, err := manager.GetWriter(ctx, tx)
	if err != nil {
		return nil, err
	}

	if _, err = transact(ctx, v.assertionChain.backend, v.assertionChain.headerReader, func() (*types.Transaction, error) {
		return writer.CreateSubChallenge(
			v.assertionChain.txOpts,
			v.id,
		)
	}); err != nil {
		return nil, err
	}

	challengeId, err := manager.CalculateChallengeHash(ctx, tx, v.id, subChallengeType)
	if err != nil {
		return nil, err
	}
	chal, err := manager.GetChallenge(ctx, tx, challengeId)
	if err != nil {
		return nil, err
	}
	if chal.IsNone() {
		return nil, errors.New("no challenge found after subchallenge creation")
	}
	return &Challenge{
		id:             challengeId,
		assertionChain: v.assertionChain,
	}, nil
}
