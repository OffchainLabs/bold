package solimpl

import (
	"context"
	"math/big"
	"strings"

	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

// Bisect a challenge vertex by providing a history commitment.
func (v *ChallengeVertex) Bisect(
	history util.HistoryCommitment,
	proof []common.Hash,
) (*ChallengeVertex, error) {
	// Flatten the last leaf proof for submission to the chain.
	flatProof := make([]byte, 0)
	for _, h := range proof {
		flatProof = append(flatProof, h[:]...)
	}
	if err2 := withChainCommitment(v.manager.assertionChain.backend, func() error {
		_, err3 := v.manager.writer.Bisect(
			v.manager.assertionChain.txOpts,
			v.id,
			history.Merkle,
			flatProof,
		)
		return err3
	}); err2 != nil {
		return nil, err2
	}
	bisectedToId, err := v.manager.caller.CalculateChallengeVertexId(
		v.manager.assertionChain.callOpts,
		v.inner.ChallengeId,
		history.Merkle,
		big.NewInt(int64(history.Height)),
	)
	if err != nil {
		return nil, err
	}
	bisectedTo, err := v.manager.caller.GetVertex(
		v.manager.assertionChain.callOpts,
		bisectedToId,
	)
	if err != nil {
		return nil, err
	}
	return &ChallengeVertex{
		inner:   bisectedTo,
		manager: v.manager,
	}, nil
}

func (v *ChallengeVertex) ConfirmPsTimer(ctx context.Context) error {
	err := withChainCommitment(v.manager.assertionChain.backend, func() error {
		_, err := v.manager.writer.ConfirmForPsTimer(
			v.manager.assertionChain.txOpts,
			v.id,
		)
		return err
	})
	switch {
	case err == nil:
	case strings.Contains(err.Error(), "PsTimer not greater than challenge period"):
		return errors.Wrapf(ErrPsTimerNotYet, "vertex id %#v", v.id)
	default:
		return err
	}
	return nil
}

func (v *ChallengeVertex) ConfirmForSuccessionChallengeWin(ctx context.Context) error {
	err := withChainCommitment(v.manager.assertionChain.backend, func() error {
		_, err := v.manager.writer.ConfirmForSucessionChallengeWin(
			v.manager.assertionChain.txOpts,
			v.id,
		)
		return err
	})
	switch {
	case err == nil:
	case strings.Contains(err.Error(), "Vertex does not exist"):
		return errors.Wrapf(ErrNotFound, "vertex id %#v", v.id)
	case strings.Contains(err.Error(), "Vertex is not pending"):
		return errors.Wrapf(ErrNonPendingAssertion, "vertex id %#v", v.id)
	case strings.Contains(err.Error(), "Predecessor vertex does not exist"):
		return errors.Wrapf(ErrNotFound, "vertex predecessor id %#v", v.inner.PredecessorId)
	case strings.Contains(err.Error(), "Predecessor vertex is not confirmed"):
		return errors.Wrapf(ErrNonConfirmedAssertion, "vertex predecessor id %#v", v.inner.PredecessorId)
	case strings.Contains(err.Error(), "Succession challenge does not exist"):
		return errors.Wrapf(ErrSuccessionNotFound, "succession challenge id %#v", v.inner.SuccessionChallenge)
	case strings.Contains(err.Error(), "Succession challenge did not declare this vertex the winner"):
		return errors.Wrapf(ErrNotFound, "succession challenge id %#v", v.inner.SuccessionChallenge)
	default:
		return err
	}
	return nil
}
