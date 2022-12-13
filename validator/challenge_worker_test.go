package validator

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/OffchainLabs/new-rollup-exploration/protocol"
	"github.com/OffchainLabs/new-rollup-exploration/state-manager"
	"github.com/OffchainLabs/new-rollup-exploration/testing/mocks"
	"github.com/OffchainLabs/new-rollup-exploration/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/require"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(io.Discard)
}

func Test_actOnBlockChallenge(t *testing.T) {
	logsHook := test.NewGlobal()
	alice := common.BytesToAddress([]byte("alice"))
	w := &blockChallengeWorker{
		createdVertices:  util.NewThreadSafeSlice[*protocol.ChallengeVertex](),
		validatorAddress: alice,
	}
	v := &Validator{}
	ctx := context.Background()
	err := w.actOnBlockChallenge(
		ctx,
		v,
		alice, // Self.
		util.HistoryCommitment{},
		protocol.VertexSequenceNumber(0),
	)
	require.NoError(t, err)

	bob := common.BytesToAddress([]byte("bob"))
	err = w.actOnBlockChallenge(
		ctx,
		v,
		bob,
		util.HistoryCommitment{},
		protocol.VertexSequenceNumber(0),
	)
	require.NoError(t, err)
	AssertLogsContain(t, logsHook, "No created vertices, not acting")
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

func Test_bisectWhileNonPresumptive(t *testing.T) {
	ctx := context.Background()
	t.Run("has presumptive successor no action taken", func(t *testing.T) {
		logsHook := test.NewGlobal()
		stateRoots := generateStateRoots(10)
		manager := statemanager.New(stateRoots)
		leaf1, leaf2, validator := createTwoValidatorFork(t, ctx, manager, stateRoots)
		challenge, _, _ := prepareBisectionGame(
			t,
			ctx,
			logsHook,
			validator,
			leaf1,
			leaf2,
		)

		var vertex1 *protocol.ChallengeVertex
		var err error
		err = validator.chain.Tx(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
			commit := util.HistoryCommitment{
				Height: 1,
				Merkle: common.BytesToHash([]byte{1}),
			}
			seqNum := protocol.AssertionSequenceNumber(1)
			assertion, fetchErr := p.AssertionBySequenceNum(tx, seqNum)
			if fetchErr != nil {
				return fetchErr
			}
			vertex1, err = challenge.AddLeaf(tx, assertion, commit, validator.address)
			if err != nil {
				return err
			}
			return nil
		})
		require.NoError(t, err)

		w := &blockChallengeWorker{
			createdVertices: util.NewThreadSafeSlice[*protocol.ChallengeVertex](),
		}
		err = w.bisectWhileNonPresumptive(ctx, validator, vertex1)
		require.NoError(t, err)
		AssertLogsContain(t, logsHook, "Has presumptive successor, not acting")
	})
	t.Run("bisects twice until hitting vertex that already exists", func(t *testing.T) {
		logsHook := test.NewGlobal()
		stateRoots := generateStateRoots(10)
		manager := statemanager.New(stateRoots)
		leaf1, leaf2, validator := createTwoValidatorFork(t, ctx, manager, stateRoots)
		challenge, _, vertex6 := prepareBisectionGame(
			t,
			ctx,
			logsHook,
			validator,
			leaf1,
			leaf2,
		)

		// We create a challenge vertex at height 1 with a history commitment we have.
		commit, err := manager.HistoryCommitmentUpTo(ctx, 1)
		require.NoError(t, err)

		err = validator.chain.Tx(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
			seqNum := protocol.AssertionSequenceNumber(1)
			assertion, fetchErr := p.AssertionBySequenceNum(tx, seqNum)
			if fetchErr != nil {
				return fetchErr
			}
			if _, err = challenge.AddLeaf(tx, assertion, commit, validator.address); err != nil {
				return err
			}
			return nil
		})

		require.NoError(t, err)
		w := &blockChallengeWorker{
			createdVertices: util.NewThreadSafeSlice[*protocol.ChallengeVertex](),
		}

		// Expect two bisections, from 6 to 4 to 2, then failing from 2 to 1 because 1 already exists, as
		// we have previously created vertex with commitment height 1.
		err = w.bisectWhileNonPresumptive(ctx, validator, vertex6)
		require.ErrorIs(t, err, protocol.ErrVertexAlreadyExists)
		AssertLogsContain(t, logsHook, "Successfully bisected to vertex with height 4")
		AssertLogsContain(t, logsHook, "Successfully bisected to vertex with height 2")
		AssertLogsDoNotContain(t, logsHook, "Reached one-step-fork")
	})
	t.Run("bisects twice until one-step-fork", func(t *testing.T) {
		logsHook := test.NewGlobal()
		stateRoots := generateStateRoots(10)
		manager := statemanager.New(stateRoots)
		leaf1, leaf2, validator := createTwoValidatorFork(t, ctx, manager, stateRoots)
		challenge, _, vertex6 := prepareBisectionGame(
			t,
			ctx,
			logsHook,
			validator,
			leaf1,
			leaf2,
		)

		// We create a challenge vertex at height 1 with a history commitment we DO NOT have.
		var err error
		err = validator.chain.Tx(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
			seqNum := protocol.AssertionSequenceNumber(1)
			assertion, fetchErr := p.AssertionBySequenceNum(tx, seqNum)
			if fetchErr != nil {
				return fetchErr
			}
			commit := util.HistoryCommitment{
				Height: 1,
				Merkle: common.BytesToHash([]byte("BOGUS COMMIT")),
			}
			if _, err = challenge.AddLeaf(tx, assertion, commit, validator.address); err != nil {
				return err
			}
			return nil
		})

		require.NoError(t, err)
		w := &blockChallengeWorker{
			createdVertices:    util.NewThreadSafeSlice[*protocol.ChallengeVertex](),
			reachedOneStepFork: make(chan struct{}, 1),
		}

		// Expect two bisections, from 6 to 4 to 2, then 2 to 1 and stopping because we have a one-step-fork.
		err = w.bisectWhileNonPresumptive(ctx, validator, vertex6)
		require.NoError(t, err)
		AssertLogsContain(t, logsHook, "Successfully bisected to vertex with height 4")
		AssertLogsContain(t, logsHook, "Successfully bisected to vertex with height 2")
		AssertLogsContain(t, logsHook, "Successfully bisected to vertex with height 1")
		AssertLogsContain(t, logsHook, "Reached one-step-fork")
	})
}

func prepareBisectionGame(
	t *testing.T,
	ctx context.Context,
	logsHook *test.Hook,
	validator *Validator,
	leaf1 *protocol.CreateLeafEvent,
	leaf2 *protocol.CreateLeafEvent,
) (*protocol.Challenge, *protocol.ChallengeVertex, *protocol.ChallengeVertex) {
	// Should first process leaf creation through the validator
	// which should result in the validator taking challenge actions on those
	// leaves by interacting with the chain. We assert an challenge log
	// has indeed been emitted.
	err := validator.onLeafCreated(ctx, leaf1)
	require.NoError(t, err)
	err = validator.onLeafCreated(ctx, leaf2)
	require.NoError(t, err)
	AssertLogsContain(t, logsHook, "New leaf appended")
	AssertLogsContain(t, logsHook, "New leaf appended")
	AssertLogsContain(t, logsHook, "Successfully created challenge and added leaf")

	require.NoError(t, err)

	genesisCommit := protocol.StateCommitment{
		Height:    0,
		StateRoot: common.Hash{},
	}

	// Upon creating a challenge, we should have added a challenge vertex to it.
	// However, because this test only has a single validator, we will
	// add the second challenge vertex as well.
	id := protocol.CommitHash(genesisCommit.Hash())
	var challenge *protocol.Challenge
	var vertexHeight5 *protocol.ChallengeVertex
	var vertexHeight6 *protocol.ChallengeVertex
	err = validator.chain.Tx(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		vertexHeight6, err = p.ChallengeVertexBySequenceNum(
			tx,
			id,
			protocol.VertexSequenceNumber(1),
		)
		if err != nil {
			return err
		}
		height5SeqNum := protocol.AssertionSequenceNumber(1)
		assertion, fetchErr := p.AssertionBySequenceNum(tx, height5SeqNum)
		if fetchErr != nil {
			return fetchErr
		}
		challenge, err = p.ChallengeByCommitHash(tx, id)
		if err != nil {
			return err
		}
		historyCommitUpTo5, err := validator.stateManager.HistoryCommitmentUpTo(
			ctx,
			uint64(5),
		)
		vertexHeight5, err = challenge.AddLeaf(
			tx, assertion, historyCommitUpTo5, validator.address,
		)
		if err != nil {
			return err
		}
		return nil
	})
	require.NoError(t, err)
	require.NotNil(t, challenge)
	require.NotNil(t, vertexHeight5)
	require.NotNil(t, vertexHeight6)
	return challenge, vertexHeight5, vertexHeight6
}
