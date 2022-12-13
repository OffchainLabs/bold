package validator

import (
	"context"
	"errors"
	"testing"

	"github.com/OffchainLabs/new-rollup-exploration/protocol"
	"github.com/OffchainLabs/new-rollup-exploration/state-manager"
	"github.com/OffchainLabs/new-rollup-exploration/testing/mocks"
	"github.com/OffchainLabs/new-rollup-exploration/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/require"
)

func Test_bisectWhileNonPresumptive(t *testing.T) {
	t.Run("has presumptive successor, no action taken", func(t *testing.T) {
	})
	t.Run("bisects three times until presumptive", func(t *testing.T) {
	})
	t.Run("bisects three times until one-step-fork", func(t *testing.T) {
	})
	t.Run("tries to bisect but bisected vertex already exists", func(t *testing.T) {
	})
}

func Test_loadVertexToActUpon(t *testing.T) {
	t.Run("no vertex with commit height > event commit height", func(t *testing.T) {
		w := &blockChallengeWorker{
			createdVertices: util.NewThreadSafeSlice[*protocol.ChallengeVertex](),
		}
		v1 := &protocol.ChallengeVertex{
			Commitment: util.HistoryCommitment{
				Height: 2,
				Merkle: common.BytesToHash([]byte{2}),
			},
		}
		w.createdVertices.Append(v1)
		got := w.loadVertexToActUpon(100)
		require.Equal(t, v1, got)
	})
	t.Run("gets first vertex with commit height > event commit height", func(t *testing.T) {
		w := &blockChallengeWorker{
			createdVertices: util.NewThreadSafeSlice[*protocol.ChallengeVertex](),
		}
		v1 := &protocol.ChallengeVertex{
			Commitment: util.HistoryCommitment{
				Height: 2,
				Merkle: common.BytesToHash([]byte{2}),
			},
		}
		v2 := &protocol.ChallengeVertex{
			Commitment: util.HistoryCommitment{
				Height: 3,
				Merkle: common.BytesToHash([]byte{3}),
			},
		}
		w.createdVertices.Append(v1)
		w.createdVertices.Append(v2)
		got := w.loadVertexToActUpon(1)
		require.Equal(t, v2, got)
	})
}

func Test_vertexToMergeInto(t *testing.T) {
	challengeCommit := protocol.StateCommitment{
		Height:    0,
		StateRoot: common.Hash{},
	}
	id := protocol.CommitHash(challengeCommit.Hash())
	seqNum := protocol.VertexSequenceNumber(1)

	t.Run("nil vertex", func(t *testing.T) {
		w := &blockChallengeWorker{}
		p := &mocks.MockProtocol{}
		var mergingTo *protocol.ChallengeVertex
		p.On("ChallengeVertexBySequenceNum", &protocol.ActiveTx{}, id, seqNum).Return(
			mergingTo, nil,
		)
		v := &Validator{
			chain: p,
		}
		_, err := w.vertexToMergeInto(v, id, 1)
		require.ErrorContains(t, err, "fetched nil challenge")
	})
	t.Run("fetching error", func(t *testing.T) {
		w := &blockChallengeWorker{}
		p := &mocks.MockProtocol{}
		var mergingTo *protocol.ChallengeVertex
		p.On("ChallengeVertexBySequenceNum", &protocol.ActiveTx{}, id, seqNum).Return(
			mergingTo, errors.New("something went wrong"),
		)
		v := &Validator{
			chain: p,
		}
		_, err := w.vertexToMergeInto(v, id, 1)
		require.ErrorContains(t, err, "something went wrong")
	})
	t.Run("OK", func(t *testing.T) {
		w := &blockChallengeWorker{}
		p := &mocks.MockProtocol{}
		mergingTo := &protocol.ChallengeVertex{
			Commitment: util.HistoryCommitment{
				Height: 1,
				Merkle: common.Hash{},
			},
		}
		p.On("ChallengeVertexBySequenceNum", &protocol.ActiveTx{}, id, seqNum).Return(
			mergingTo, nil,
		)
		v := &Validator{
			chain: p,
		}
		got, err := w.vertexToMergeInto(v, id, 1)
		require.NoError(t, err)
		require.Equal(t, mergingTo, got)
	})
}

func Test_shouldMakeMergeMove(t *testing.T) {
	ctx := context.Background()
	stateRoots := generateStateRoots(10)
	manager := statemanager.New(stateRoots)
	t.Run("merged to ours", func(t *testing.T) {
		logsHook := test.NewGlobal()
		w := &blockChallengeWorker{}
		v := &Validator{}
		commit := util.HistoryCommitment{
			Height: 1,
			Merkle: common.Hash{},
		}
		shouldMerge := w.shouldMakeMergeMove(
			ctx,
			v,
			commit,
			commit,
		)
		require.Equal(t, false, shouldMerge)
		AssertLogsContain(t, logsHook, "Other validator merged to our vertex")
	})
	t.Run("should not merge if no matching history commitment exists", func(t *testing.T) {
		logsHook := test.NewGlobal()
		w := &blockChallengeWorker{}
		v := &Validator{
			stateManager: manager,
		}

		ourCommit := util.HistoryCommitment{
			Height: 6,
			Merkle: common.Hash{},
		}
		incomingCommit := util.HistoryCommitment{
			Height: 1,
			Merkle: common.BytesToHash([]byte("BOGUS COMMIT")),
		}

		shouldMerge := w.shouldMakeMergeMove(
			ctx,
			v,
			incomingCommit,
			ourCommit,
		)
		require.Equal(t, false, shouldMerge)
		AssertLogsDoNotContain(t, logsHook, "Other validator merged to our vertex")

	})
	t.Run("should merge if matching history commitment exists", func(t *testing.T) {
		existingCommit, err := manager.HistoryCommitmentUpTo(ctx, 1)
		require.NoError(t, err)

		logsHook := test.NewGlobal()
		w := &blockChallengeWorker{}
		v := &Validator{
			stateManager: manager,
		}

		ourCommit := util.HistoryCommitment{
			Height: 5,
			Merkle: common.Hash{},
		}

		shouldMerge := w.shouldMakeMergeMove(
			ctx,
			v,
			existingCommit,
			ourCommit,
		)
		require.Equal(t, true, shouldMerge)
		AssertLogsDoNotContain(t, logsHook, "Other validator merged to our vertex")
	})
}
