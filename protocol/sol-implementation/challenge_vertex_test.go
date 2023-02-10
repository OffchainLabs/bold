package solimpl

import (
	"context"
	"testing"
	"time"

	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestChallengeVertex_ConfirmPsTimer(t *testing.T) {
	chain, acc := setupAssertionChainWithChallengeManager(t)
	height1 := uint64(6)
	height2 := uint64(7)
	a1, _, challenge := setupTopLevelFork(t, chain, height1, height2)

	genesis, err := chain.AssertionByID(common.Hash{})
	require.NoError(t, err)

	v1, err := challenge.AddLeaf(
		a1,
		util.HistoryCommitment{
			Height:    height1,
			Merkle:    common.BytesToHash([]byte("nyan")),
			FirstLeaf: genesis.inner.StateHash,
		},
	)
	require.NoError(t, err)

	t.Run("vertex ps timer has not exceeded challenge duration", func(t *testing.T) {
		require.ErrorIs(t, v1.ConfirmPsTimer(context.Background()), ErrPsTimerNotYet)
	})
	t.Run("vertex ps timer has exceeded challenge duration", func(t *testing.T) {
		require.NoError(t, acc.backend.AdjustTime(time.Second*2000))
		require.NoError(t, v1.ConfirmPsTimer(context.Background()))
	})
}

func TestChallengeVertex_ConfirmForSuccessionChallengeWin(t *testing.T) {
	chain, acc := setupAssertionChainWithChallengeManager(t)
	height1 := uint64(6)
	height2 := uint64(7)
	a1, _, challenge := setupTopLevelFork(t, chain, height1, height2)

	genesis, err := chain.AssertionByID(common.Hash{})
	require.NoError(t, err)

	v1, err := challenge.AddLeaf(
		a1,
		util.HistoryCommitment{
			Height:    height1,
			Merkle:    common.BytesToHash([]byte("nyan")),
			FirstLeaf: genesis.inner.StateHash,
		},
	)
	require.NoError(t, err)

	t.Run("vertex does not exist", func(t *testing.T) {
		oldId := v1.id
		v1.id = common.Hash{1}
		require.ErrorIs(t, v1.ConfirmForSuccessionChallengeWin(context.Background()), ErrNotFound)
		v1.id = oldId
	})
	t.Run("succession does not exist", func(t *testing.T) {
		require.NoError(t, acc.backend.AdjustTime(time.Second*2000))
		require.NoError(t, v1.ConfirmPsTimer(context.Background()))
		acc.backend.Commit()
		require.ErrorIs(t, v1.ConfirmForSuccessionChallengeWin(context.Background()), ErrSuccessionNotFound)
	})
	t.Run("can confirm through succession", func(t *testing.T) {
		t.Skip("TODO: Need one step proof")
	})
}

func TestChallengeVertex_Bisect(t *testing.T) {
	chain, acc := setupAssertionChainWithChallengeManager(t)
	height1 := uint64(6)
	height2 := uint64(7)
	a1, a2, challenge := setupTopLevelFork(t, chain, height1, height2)

	genesis, err := chain.AssertionByID(common.Hash{})
	require.NoError(t, err)

	// We add two leaves to the challenge.
	v1, err := challenge.AddLeaf(
		a1,
		util.HistoryCommitment{
			Height:    height1,
			Merkle:    common.BytesToHash([]byte("nyan")),
			FirstLeaf: genesis.inner.StateHash,
		},
	)
	require.NoError(t, err)
	v2, err := challenge.AddLeaf(
		a2,
		util.HistoryCommitment{
			Height:    height2,
			Merkle:    common.BytesToHash([]byte("nyan2")),
			FirstLeaf: genesis.inner.StateHash,
		},
	)
	require.NoError(t, err)

	t.Run("vertex does not exist", func(t *testing.T) {
		vertex := &ChallengeVertex{
			id:      common.BytesToHash([]byte("junk")),
			manager: challenge.manager,
		}
		_, err = vertex.Bisect(
			util.HistoryCommitment{
				Height:    4,
				Merkle:    common.BytesToHash([]byte("nyan2")),
				FirstLeaf: genesis.inner.StateHash,
			},
			make([]common.Hash, 0),
		)
		require.ErrorContains(t, err, "does not exist")
	})
	t.Run("winner already declared", func(t *testing.T) {
		t.Skip("Need to add winner capabilities in order to test")
	})
	t.Run("cannot bisect presumptive successor", func(t *testing.T) {
		// V1 should be the presumptive successor here.
		_, err = v1.Bisect(
			util.HistoryCommitment{
				Height:    4,
				Merkle:    common.BytesToHash([]byte("nyan2")),
				FirstLeaf: genesis.inner.StateHash,
			},
			make([]common.Hash, 0),
		)
		require.ErrorContains(t, err, "Cannot bisect presumptive")
	})
	t.Run("presumptive successor already confirmable", func(t *testing.T) {
		chalPeriod, err := chain.ChallengePeriodSeconds()
		require.NoError(t, err)
		err = acc.backend.AdjustTime(chalPeriod)
		require.NoError(t, err)
		// We make a challenge period pass.
		_, err = v2.Bisect(
			util.HistoryCommitment{
				Height:    4,
				Merkle:    common.BytesToHash([]byte("nyan2")),
				FirstLeaf: genesis.inner.StateHash,
			},
			make([]common.Hash, 0),
		)
		require.ErrorContains(t, err, "cannot set lower ps")
	})
	t.Run("invalid prefix history", func(t *testing.T) {
		t.Skip("Need to add proof capabilities in solidity in order to test")
	})
	t.Run("OK", func(t *testing.T) {
		chain, _ = setupAssertionChainWithChallengeManager(t)
		height1 = uint64(6)
		height2 = uint64(7)
		a1, a2, challenge = setupTopLevelFork(t, chain, height1, height2)

		// We add two leaves to the challenge.
		v1, err := challenge.AddLeaf(
			a1,
			util.HistoryCommitment{
				Height:    height1,
				Merkle:    common.BytesToHash([]byte("nyan")),
				FirstLeaf: genesis.inner.StateHash,
			},
		)
		require.NoError(t, err)
		v2, err = challenge.AddLeaf(
			a2,
			util.HistoryCommitment{
				Height:    height2,
				Merkle:    common.BytesToHash([]byte("nyan2")),
				FirstLeaf: genesis.inner.StateHash,
			},
		)
		require.NoError(t, err)
		wantCommit := common.BytesToHash([]byte("nyan2"))
		bisectedTo, err := v2.Bisect(
			util.HistoryCommitment{
				Height:    4,
				Merkle:    wantCommit,
				FirstLeaf: genesis.inner.StateHash,
			},
			make([]common.Hash, 0),
		)
		require.NoError(t, err)
		require.Equal(t, uint64(4), bisectedTo.inner.Height.Uint64())
		require.Equal(t, wantCommit[:], bisectedTo.inner.HistoryRoot[:])

		_, err = v1.Bisect(
			util.HistoryCommitment{
				Height:    4,
				Merkle:    wantCommit,
				FirstLeaf: genesis.inner.StateHash,
			},
			make([]common.Hash, 0),
		)
		require.ErrorContains(t, err, "already exists")
	})
}
