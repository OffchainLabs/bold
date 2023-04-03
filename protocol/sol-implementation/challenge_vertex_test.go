package solimpl_test

import (
	"context"
	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/state-manager"
	"github.com/OffchainLabs/challenge-protocol-v2/testing/setup"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEdgeChallengeManager_IsPresumptive(t *testing.T) {
	ctx := context.Background()
	height := protocol.Height(3)

	createdData, err := setup.CreateTwoValidatorFork(ctx, &setup.CreateForkConfig{
		NumBlocks:     uint64(height) + 1,
		DivergeHeight: 0,
	})
	require.NoError(t, err)

	opts := []statemanager.Opt{
		statemanager.WithNumOpcodesPerBigStep(1),
		statemanager.WithMaxWavmOpcodesPerBlock(1),
	}

	honestStateManager, err := statemanager.New(
		createdData.HonestValidatorStateRoots,
		opts...,
	)
	require.NoError(t, err)
	evilStateManager, err := statemanager.New(
		createdData.EvilValidatorStateRoots,
		opts...,
	)
	require.NoError(t, err)

	challengeManager, err := createdData.Chains[0].SpecChallengeManager(ctx)
	require.NoError(t, err)
	genesis, err := createdData.Chains[0].AssertionBySequenceNum(ctx, 0)
	require.NoError(t, err)

	// Honest assertion being added.
	leafAdder := func(endCommit util.HistoryCommitment) protocol.SpecEdge {
		leaf, err := challengeManager.AddBlockChallengeLevelZeroEdge(
			ctx,
			genesis,
			util.HistoryCommitment{Merkle: common.Hash{}},
			endCommit,
		)
		require.NoError(t, err)
		return leaf
	}
	honestEndCommit, err := honestStateManager.HistoryCommitmentUpTo(ctx, uint64(height))
	require.NoError(t, err)

	honestEdge := leafAdder(honestEndCommit)
	require.Equal(t, protocol.BlockChallengeEdge, honestEdge.GetType())

	t.Run("first leaf is presumptive", func(t *testing.T) {
		isPs, err := honestEdge.IsPresumptive(ctx)
		require.NoError(t, err)
		require.Equal(t, true, isPs)
	})

	evilEndCommit, err := evilStateManager.HistoryCommitmentUpTo(ctx, uint64(height))
	require.NoError(t, err)

	evilEdge := leafAdder(evilEndCommit)
	require.Equal(t, protocol.BlockChallengeEdge, evilEdge.GetType())

	t.Run("neither is presumptive if rivals", func(t *testing.T) {
		isPs, err := honestEdge.IsPresumptive(ctx)
		require.NoError(t, err)
		require.Equal(t, false, isPs)

		isPs, err = evilEdge.IsPresumptive(ctx)
		require.NoError(t, err)
		require.Equal(t, false, isPs)
	})

	t.Run("bisected children are presumptive", func(t *testing.T) {
		honestBisectCommit, err := honestStateManager.HistoryCommitmentUpTo(ctx, 1)
		require.NoError(t, err)
		honestProof, err := honestStateManager.PrefixProof(ctx, 1, 3)
		require.NoError(t, err)
		lower, upper, err := honestEdge.Bisect(ctx, honestBisectCommit.Merkle, honestProof)
		require.NoError(t, err)

		isPs, err := lower.IsPresumptive(ctx)
		require.NoError(t, err)
		require.Equal(t, true, isPs)
		isPs, err = upper.IsPresumptive(ctx)
		require.NoError(t, err)
		require.Equal(t, true, isPs)

		isPs, err = honestEdge.IsPresumptive(ctx)
		require.NoError(t, err)
		require.Equal(t, false, isPs)

		isPs, err = evilEdge.IsPresumptive(ctx)
		require.NoError(t, err)
		require.Equal(t, false, isPs)
	})
}

func TestSpecChallengeManager_IsOneStepForkSource(t *testing.T) {
	ctx := context.Background()
	height := protocol.Height(3)

	createdData, err := setup.CreateTwoValidatorFork(ctx, &setup.CreateForkConfig{
		NumBlocks:     uint64(height) + 1,
		DivergeHeight: 0,
	})
	require.NoError(t, err)

	opts := []statemanager.Opt{
		statemanager.WithNumOpcodesPerBigStep(1),
		statemanager.WithMaxWavmOpcodesPerBlock(1),
	}

	honestStateManager, err := statemanager.New(
		createdData.HonestValidatorStateRoots,
		opts...,
	)
	require.NoError(t, err)
	evilStateManager, err := statemanager.New(
		createdData.EvilValidatorStateRoots,
		opts...,
	)
	require.NoError(t, err)

	challengeManager, err := createdData.Chains[0].SpecChallengeManager(ctx)
	require.NoError(t, err)
	genesis, err := createdData.Chains[0].AssertionBySequenceNum(ctx, 0)
	require.NoError(t, err)

	// Honest assertion being added.
	leafAdder := func(endCommit util.HistoryCommitment) protocol.SpecEdge {
		leaf, err := challengeManager.AddBlockChallengeLevelZeroEdge(
			ctx,
			genesis,
			util.HistoryCommitment{Merkle: common.Hash{}},
			endCommit,
		)
		require.NoError(t, err)
		return leaf
	}
	honestEndCommit, err := honestStateManager.HistoryCommitmentUpTo(ctx, uint64(height))
	require.NoError(t, err)

	honestEdge := leafAdder(honestEndCommit)
	require.Equal(t, protocol.BlockChallengeEdge, honestEdge.GetType())

	t.Run("lone level zero edge is not one step fork source", func(t *testing.T) {
		isOSF, err := honestEdge.IsOneStepForkSource(ctx)
		require.NoError(t, err)
		require.Equal(t, false, isOSF)
	})

	evilEndCommit, err := evilStateManager.HistoryCommitmentUpTo(ctx, uint64(height))
	require.NoError(t, err)
	evilEdge := leafAdder(evilEndCommit)
	require.Equal(t, protocol.BlockChallengeEdge, evilEdge.GetType())

	t.Run("level zero edge with rivals is not one step fork source", func(t *testing.T) {
		isOSF, err := honestEdge.IsOneStepForkSource(ctx)
		require.NoError(t, err)
		require.Equal(t, false, isOSF)
		isOSF, err = evilEdge.IsOneStepForkSource(ctx)
		require.NoError(t, err)
		require.Equal(t, false, isOSF)
	})
	t.Run("single bisected edge is not one step fork source", func(t *testing.T) {
		honestBisectCommit, err := honestStateManager.HistoryCommitmentUpTo(ctx, 1)
		require.NoError(t, err)
		honestProof, err := honestStateManager.PrefixProof(ctx, 1, 3)
		require.NoError(t, err)
		lower, upper, err := honestEdge.Bisect(ctx, honestBisectCommit.Merkle, honestProof)
		require.NoError(t, err)

		isOSF, err := lower.IsOneStepForkSource(ctx)
		require.NoError(t, err)
		require.Equal(t, false, isOSF)
		isOSF, err = upper.IsOneStepForkSource(ctx)
		require.NoError(t, err)
		require.Equal(t, false, isOSF)
	})
	t.Run("post bisection, mutual edge is one step fork source", func(t *testing.T) {
		evilBisectCommit, err := evilStateManager.HistoryCommitmentUpTo(ctx, 1)
		require.NoError(t, err)
		evilProof, err := evilStateManager.PrefixProof(ctx, 1, 3)
		require.NoError(t, err)
		lower, upper, err := evilEdge.Bisect(ctx, evilBisectCommit.Merkle, evilProof)
		require.NoError(t, err)

		isOSF, err := lower.IsOneStepForkSource(ctx)
		require.NoError(t, err)
		require.Equal(t, true, isOSF)

		isOSF, err = upper.IsOneStepForkSource(ctx)
		require.NoError(t, err)
		require.Equal(t, false, isOSF)
	})
}

// func TestChallengeVertex_Bisect(t *testing.T) {
// 	ctx := context.Background()
// 	height1 := uint64(3)
// 	height2 := uint64(7)
// 	a1, a2, challenge, chain1, chain2 := setupTopLevelFork(t, ctx, height1, height2)

// 	honestHashes := honestHashesUpTo(10)
// 	evilHashes := evilHashesUpTo(10)
// 	honestManager, err := statemanager.New(honestHashes)
// 	require.NoError(t, err)

// 	evilManager, err := statemanager.New(evilHashes)
// 	require.NoError(t, err)
// 	honestCommit, err := honestManager.HistoryCommitmentUpTo(ctx, height1)
// 	require.NoError(t, err)
// 	evilCommit, err := evilManager.HistoryCommitmentUpTo(ctx, height2)
// 	require.NoError(t, err)

// 	// We add two leaves to the challenge.
// 	challenge.chain = chain1
// 	v1, err := challenge.AddBlockChallengeLeaf(
// 		ctx,
// 		a1,
// 		honestCommit,
// 	)
// 	require.NoError(t, err)

// 	challenge.chain = chain2
// 	v2, err := challenge.AddBlockChallengeLeaf(
// 		ctx,
// 		a2,
// 		evilCommit,
// 	)
// 	require.NoError(t, err)

// 	t.Run("vertex does not exist", func(t *testing.T) {
// 		vertex := &ChallengeVertex{
// 			id:    common.BytesToHash([]byte("junk")),
// 			chain: challenge.chain,
// 		}
// 		_, err = vertex.Bisect(ctx, util.HistoryCommitment{
// 			Height: 4,
// 			Merkle: common.BytesToHash([]byte("nyan4")),
// 		}, make([]byte, 0))
// 		require.ErrorContains(t, err, "does not exist")
// 	})
// 	t.Run("winner already declared", func(t *testing.T) {
// 		t.Skip("Need to add winner capabilities in order to test")
// 	})
// 	t.Run("cannot bisect presumptive successor", func(t *testing.T) {
// 		// V1 should be the presumptive successor here.
// 		_, err = v1.Bisect(ctx, util.HistoryCommitment{
// 			Height: 4,
// 			Merkle: common.BytesToHash([]byte("nyan4")),
// 		}, make([]byte, 0))
// 		require.ErrorContains(t, err, "Cannot bisect presumptive")
// 	})
// 	t.Run("presumptive successor already confirmable", func(t *testing.T) {
// 		manager, err := chain1.CurrentChallengeManager(ctx)
// 		require.NoError(t, err)
// 		chalPeriod, err := manager.ChallengePeriodSeconds(ctx)
// 		require.NoError(t, err)
// 		backend, ok := chain1.backend.(*backends.SimulatedBackend)
// 		require.Equal(t, true, ok)
// 		err = backend.AdjustTime(chalPeriod)
// 		require.NoError(t, err)

// 		preCommit, err := evilManager.HistoryCommitmentUpTo(ctx, 3)
// 		require.NoError(t, err)
// 		prefixProof, err := evilManager.PrefixProof(ctx, 3, 7)
// 		require.NoError(t, err)

// 		// We make a challenge period pass.
// 		_, err = v2.Bisect(ctx, preCommit, prefixProof)
// 		require.ErrorContains(t, err, "cannot set same height ps")
// 	})
// 	t.Run("invalid prefix history", func(t *testing.T) {
// 		t.Skip("Need to add proof capabilities in solidity in order to test")
// 	})
// 	t.Run("OK", func(t *testing.T) {
// 		height1 := uint64(3)
// 		height2 := uint64(7)
// 		a1, a2, challenge, chain1, chain2 := setupTopLevelFork(t, ctx, height1, height2)

// 		// We add two leaves to the challenge.
// 		challenge.chain = chain1
// 		v1, err := challenge.AddBlockChallengeLeaf(
// 			ctx,
// 			a1,
// 			honestCommit,
// 		)
// 		require.NoError(t, err)

// 		challenge.chain = chain2
// 		v2, err := challenge.AddBlockChallengeLeaf(
// 			ctx,
// 			a2,
// 			evilCommit,
// 		)
// 		require.NoError(t, err)

// 		preCommit, err := evilManager.HistoryCommitmentUpTo(ctx, 3)
// 		require.NoError(t, err)
// 		prefixProof, err := evilManager.PrefixProof(ctx, 3, 7)
// 		require.NoError(t, err)

// 		bisectedToV, err := v2.Bisect(ctx, preCommit, prefixProof)
// 		require.NoError(t, err)
// 		bisectedTo := bisectedToV.(*ChallengeVertex)
// 		bisectedToInner, err := bisectedTo.inner(ctx)
// 		require.NoError(t, err)
// 		require.Equal(t, uint64(3), bisectedToInner.Height.Uint64())

// 		bisectTo, err := util.BisectionPoint(0, 3)
// 		require.NoError(t, err)

// 		preCommit, err = honestManager.HistoryCommitmentUpTo(ctx, bisectTo)
// 		require.NoError(t, err)
// 		prefixProof, err = honestManager.PrefixProof(ctx, bisectTo, 3)
// 		require.NoError(t, err)

// 		bisectedToV, err = v1.Bisect(ctx, preCommit, prefixProof)
// 		require.NoError(t, err)
// 		bisectedTo = bisectedToV.(*ChallengeVertex)
// 		bisectedToInner, err = bisectedTo.inner(ctx)
// 		require.NoError(t, err)
// 		require.Equal(t, uint64(1), bisectedToInner.Height.Uint64())
// 	})
// }

// func TestChallengeVertex_CreateSubChallenge(t *testing.T) {
// 	ctx := context.Background()
// 	height1 := uint64(7)
// 	height2 := uint64(7)

// 	t.Run("Error: vertex does not exist", func(t *testing.T) {
// 		_, _, challenge, _, _ := setupTopLevelFork(t, ctx, height1, height2)

// 		vertex := &ChallengeVertex{
// 			id:    common.BytesToHash([]byte("junk")),
// 			chain: challenge.chain,
// 		}
// 		_, err := vertex.CreateSubChallenge(ctx)
// 		require.ErrorContains(t, err, "execution reverted: Vertex does not exist")
// 	})

// 	honestHashes := honestHashesUpTo(10)
// 	evilHashes := divergingHashesStartingAt(t, 1, honestHashes)
// 	honestManager, err := statemanager.New(honestHashes)
// 	require.NoError(t, err)

// 	evilManager, err := statemanager.New(evilHashes)
// 	require.NoError(t, err)
// 	honestCommit, err := honestManager.HistoryCommitmentUpTo(ctx, height1)
// 	require.NoError(t, err)
// 	evilCommit, err := evilManager.HistoryCommitmentUpTo(ctx, height2)
// 	require.NoError(t, err)

// 	t.Run("Error: leaf can never be a fork candidate", func(t *testing.T) {
// 		a1, _, challenge, _, _ := setupTopLevelFork(t, ctx, height1, height2)

// 		v1, err := challenge.AddBlockChallengeLeaf(ctx, a1, honestCommit)
// 		require.NoError(t, err)
// 		_, err = v1.CreateSubChallenge(ctx)
// 		require.ErrorContains(t, err, "execution reverted: Leaf can never be a fork candidate")
// 	})
// 	t.Run("Error: lowest height not one above the current height", func(t *testing.T) {
// 		a1, a2, challenge, _, _ := setupTopLevelFork(t, ctx, height1, height2)

// 		// We add two leaves to the challenge.
// 		_, err := challenge.AddBlockChallengeLeaf(ctx, a1, honestCommit)
// 		require.NoError(t, err)
// 		v2, err := challenge.AddBlockChallengeLeaf(ctx, a2, evilCommit)
// 		require.NoError(t, err)

// 		preCommit, err := evilManager.HistoryCommitmentUpTo(ctx, 3)
// 		require.NoError(t, err)
// 		prefixProof, err := evilManager.PrefixProof(ctx, 3, 7)
// 		require.NoError(t, err)
// 		bisectedToV, err := v2.Bisect(ctx, preCommit, prefixProof)
// 		require.NoError(t, err)
// 		bisectedTo := bisectedToV.(*ChallengeVertex)
// 		bisectedToInner, err := bisectedTo.inner(ctx)
// 		require.NoError(t, err)
// 		require.Equal(t, uint64(3), bisectedToInner.Height.Uint64())

// 		// Vertex must be in the protocol.
// 		challengeManager, err := challenge.manager(ctx)
// 		require.NoError(t, err)
// 		_, err = challengeManager.caller.GetVertex(challenge.chain.callOpts, bisectedTo.id)
// 		require.NoError(t, err)
// 		_, err = bisectedTo.CreateSubChallenge(ctx)
// 		require.ErrorContains(t, err, "execution reverted: Lowest height not one above the current height")
// 	})
// 	t.Run("Error: has presumptive successor", func(t *testing.T) {
// 		height1 = uint64(2)
// 		height2 = uint64(2)
// 		a1, a2, challenge, _, _ := setupTopLevelFork(t, ctx, height1, height2)
// 		honestHashes := honestHashesUpTo(10)
// 		evilHashes := divergingHashesStartingAt(t, 1, honestHashes)
// 		honestManager, err := statemanager.New(honestHashes)
// 		require.NoError(t, err)

// 		evilManager, err := statemanager.New(evilHashes)
// 		require.NoError(t, err)
// 		honestCommit, err := honestManager.HistoryCommitmentUpTo(ctx, height1)
// 		require.NoError(t, err)
// 		evilCommit, err := evilManager.HistoryCommitmentUpTo(ctx, height2)
// 		require.NoError(t, err)

// 		// We add two leaves to the challenge.
// 		v1, err := challenge.AddBlockChallengeLeaf(ctx, a1, honestCommit)
// 		require.NoError(t, err)

// 		v2, err := challenge.AddBlockChallengeLeaf(ctx, a2, evilCommit)
// 		require.NoError(t, err)

// 		preCommit, err := evilManager.HistoryCommitmentUpTo(ctx, 1)
// 		require.NoError(t, err)
// 		prefixProof, err := evilManager.PrefixProof(ctx, 1, 2)
// 		require.NoError(t, err)

// 		_, err = v2.Bisect(ctx, preCommit, prefixProof)
// 		require.NoError(t, err)

// 		rootVertex, err := v1.Prev(ctx)
// 		require.NoError(t, err)
// 		_, err = rootVertex.Unwrap().CreateSubChallenge(ctx)
// 		require.ErrorContains(t, err, "Has presumptive successor")
// 	})
// }

// func TestChallengeVertex_AddSubChallengeLeaf(t *testing.T) {
// 	t.Skip("TODO: replace this test with edge base design")
// 	ctx := context.Background()
// 	bigStepChal, parent, firstChild, chalManager := setupBigStepSubChallenge(t)

// 	subChalHashes := make([]common.Hash, 8)
// 	for i := range subChalHashes {
// 		subChalHashes[i] = crypto.Keccak256Hash([]byte(fmt.Sprintf("foo-%d", i)))
// 	}
// 	bigStepManager, err := statemanager.New(subChalHashes)
// 	require.NoError(t, err)

// 	firstChildHistoryCommitment := firstChild.HistoryCommitment()
// 	bigStepCommit, err := bigStepManager.HistoryCommitmentUpTo(ctx, firstChildHistoryCommitment.Height)
// 	require.NoError(t, err)

// 	leaf := &mocks.MockChallengeVertex{
// 		MockId: firstChild.Id(),
// 		MockPrev: util.Some(protocol.ChallengeVertex(&mocks.MockChallengeVertex{
// 			MockHistory: util.HistoryCommitment{
// 				Merkle: subChalHashes[0],
// 			},
// 		})),
// 	}

// 	t.Run("empty history root", func(t *testing.T) {
// 		_, err = bigStepChal.AddSubChallengeLeaf(ctx, firstChild, util.HistoryCommitment{})
// 		require.ErrorContains(t, err, "execution reverted: Empty historyRoot")
// 	})
// 	t.Run("vertex does not exist", func(t *testing.T) {
// 		_, err = bigStepChal.AddSubChallengeLeaf(ctx, &ChallengeVertex{
// 			id:    [32]byte{},
// 			chain: chalManager.assertionChain,
// 		}, bigStepCommit)
// 		require.ErrorContains(t, err, "execution reverted: Claim does not exist")
// 	})
// 	t.Run("claim has invalid succession challenge", func(t *testing.T) {
// 		_, err = bigStepChal.AddSubChallengeLeaf(ctx, parent, bigStepCommit)
// 		require.ErrorContains(t, err, "execution reverted: Claim has invalid succession challenge")
// 	})
// 	t.Run("OK", func(t *testing.T) {
// 		bigStepLeaf, err := bigStepChal.AddSubChallengeLeaf(ctx, leaf, bigStepCommit)
// 		require.NoError(t, err)
// 		require.False(t, bigStepLeaf.Id() == [32]byte{}) // Should have a non-empty ID
// 	})
// }

// func TestChallengeVertex_CanConfirmSubChallenge(t *testing.T) {
// 	t.Skip("TODO: replace this test with edge base design")
// 	ctx := context.Background()
// 	bigStepChal, _, firstChild, chalManager := setupBigStepSubChallenge(t)

// 	subChalHashes := make([]common.Hash, 8)
// 	for i := range subChalHashes {
// 		subChalHashes[i] = crypto.Keccak256Hash([]byte(fmt.Sprintf("foo-%d", i)))
// 	}
// 	bigStepManager, err := statemanager.New(subChalHashes)
// 	require.NoError(t, err)

// 	firstChildHistoryCommitment := firstChild.HistoryCommitment()
// 	bigStepCommit, err := bigStepManager.HistoryCommitmentUpTo(ctx, firstChildHistoryCommitment.Height)
// 	require.NoError(t, err)

// 	leaf := &mocks.MockChallengeVertex{
// 		MockId: firstChild.Id(),
// 		MockPrev: util.Some(protocol.ChallengeVertex(&mocks.MockChallengeVertex{
// 			MockHistory: util.HistoryCommitment{
// 				Merkle: subChalHashes[0],
// 			},
// 		})),
// 	}
// 	bigStepLeaf, err := bigStepChal.AddSubChallengeLeaf(ctx, leaf, bigStepCommit)
// 	require.NoError(t, err)

// 	t.Run("can't confirm sub challenge", func(t *testing.T) {
// 		require.ErrorContains(t, bigStepLeaf.ConfirmForPsTimer(ctx), "ps timer has not exceeded challenge period")
// 	})
// 	t.Run("can confirm sub challenge", func(t *testing.T) {
// 		backend, ok := chalManager.assertionChain.backend.(*backends.SimulatedBackend)
// 		require.Equal(t, true, ok)
// 		for i := 0; i < 1000; i++ {
// 			backend.Commit()
// 		}
// 		require.NoError(t, bigStepLeaf.ConfirmForPsTimer(ctx))
// 	})
// }

// func setupBigStepSubChallenge(t *testing.T) (
// 	subChal protocol.Challenge,
// 	parent protocol.ChallengeVertex,
// 	firstChild protocol.ChallengeVertex,
// 	chalManager *ChallengeManager,
// ) {
// 	t.Helper()
// 	ctx := context.Background()
// 	height1 := uint64(7)
// 	height2 := uint64(7)
// 	a1, a2, challenge, chain, _ := setupTopLevelFork(t, ctx, height1, height2)

// 	honestHashes := honestHashesUpTo(10)
// 	evilHashes := divergingHashesStartingAt(t, 3, honestHashes)
// 	honestManager, err := statemanager.New(honestHashes)
// 	require.NoError(t, err)

// 	evilManager, err := statemanager.New(evilHashes)
// 	require.NoError(t, err)

// 	honestCommit, err := honestManager.HistoryCommitmentUpTo(ctx, height1)
// 	require.NoError(t, err)
// 	evilCommit, err := evilManager.HistoryCommitmentUpTo(ctx, height2)
// 	require.NoError(t, err)

// 	// We add two leaves to the challenge.
// 	v1, err := challenge.AddBlockChallengeLeaf(ctx, a1, honestCommit)
// 	require.NoError(t, err)

// 	v2, err := challenge.AddBlockChallengeLeaf(ctx, a2, evilCommit)
// 	require.NoError(t, err)

// 	preCommit, err := evilManager.HistoryCommitmentUpTo(ctx, 3)
// 	require.NoError(t, err)
// 	prefixProof, err := evilManager.PrefixProof(ctx, 3, 7)
// 	require.NoError(t, err)

// 	v2Height3V, err := v2.Bisect(ctx, preCommit, prefixProof)
// 	require.NoError(t, err)
// 	v2Height3 := v2Height3V.(*ChallengeVertex)
// 	v2Height3Inner, err := v2Height3.inner(ctx)
// 	require.NoError(t, err)
// 	require.Equal(t, uint64(3), v2Height3Inner.Height.Uint64())

// 	preCommit, err = honestManager.HistoryCommitmentUpTo(ctx, 3)
// 	require.NoError(t, err)
// 	prefixProof, err = honestManager.PrefixProof(ctx, 3, 7)
// 	require.NoError(t, err)

// 	v1Height3V, err := v1.Bisect(ctx, preCommit, prefixProof)
// 	require.NoError(t, err)
// 	v1Height3 := v1Height3V.(*ChallengeVertex)
// 	v1Height3Inner, err := v1Height3.inner(ctx)
// 	require.NoError(t, err)
// 	require.Equal(t, uint64(3), v1Height3Inner.Height.Uint64())

// 	preCommit, err = evilManager.HistoryCommitmentUpTo(ctx, 1)
// 	require.NoError(t, err)
// 	prefixProof, err = evilManager.PrefixProof(ctx, 1, 3)
// 	require.NoError(t, err)
// 	v2Height1V, err := v2Height3.Bisect(ctx, preCommit, prefixProof)
// 	require.NoError(t, err)
// 	v2Height1 := v2Height1V.(*ChallengeVertex)
// 	v2Height1Inner, err := v2Height1.inner(ctx)
// 	require.NoError(t, err)
// 	require.Equal(t, uint64(1), v2Height1Inner.Height.Uint64())

// 	preCommit, err = evilManager.HistoryCommitmentUpTo(ctx, 2)
// 	require.NoError(t, err)
// 	prefixProof, err = evilManager.PrefixProof(ctx, 2, 3)
// 	require.NoError(t, err)
// 	v2Height2V, err := v2Height3.Bisect(ctx, preCommit, prefixProof)
// 	require.NoError(t, err)
// 	v2Height2 := v2Height2V.(*ChallengeVertex)
// 	v2Height2Inner, err := v2Height2.inner(ctx)
// 	require.NoError(t, err)
// 	require.Equal(t, uint64(2), v2Height2Inner.Height.Uint64())

// 	cm, err := chain.CurrentChallengeManager(ctx)
// 	require.NoError(t, err)
// 	chalManager = cm.(*ChallengeManager)
// 	return
// }

// func setupAssertionChainWithChallengeManager(t *testing.T) (*AssertionChain, []*testAccount, *rollupAddresses, *backends.SimulatedBackend, *headerreader.HeaderReader) {
// 	t.Helper()
// 	ctx := context.Background()
// 	accs, backend := setupAccounts(t, 3)
// 	prod := false
// 	wasmModuleRoot := common.Hash{}
// 	rollupOwner := accs[0].accountAddr
// 	chainId := big.NewInt(1337)
// 	loserStakeEscrow := common.Address{}
// 	challengePeriodSeconds := big.NewInt(100)
// 	miniStake := big.NewInt(1)
// 	cfg := generateRollupConfig(prod, wasmModuleRoot, rollupOwner, chainId, loserStakeEscrow, challengePeriodSeconds, miniStake)
// 	addresses := deployFullRollupStack(
// 		t,
// 		ctx,
// 		backend,
// 		accs[0].txOpts,
// 		common.Address{}, // Sequencer addr.
// 		cfg,
// 	)
// 	headerReader := headerreader.New(util.SimulatedBackendWrapper{SimulatedBackend: backend}, func() *headerreader.Config { return &headerreader.TestConfig })
// 	headerReader.Start(ctx)
// 	chain, err := NewAssertionChain(
// 		ctx,
// 		addresses.Rollup,
// 		accs[1].txOpts,
// 		&bind.CallOpts{},
// 		accs[1].accountAddr,
// 		backend,
// 		headerReader,
// 		common.Address{},
// 	)
// 	require.NoError(t, err)
// 	return chain, accs, addresses, backend, headerReader
// }

// func TestCopyTxOpts(t *testing.T) {
// 	a := &bind.TransactOpts{
// 		From:      common.BigToAddress(big.NewInt(1)),
// 		Nonce:     big.NewInt(2),
// 		Value:     big.NewInt(3),
// 		GasPrice:  big.NewInt(4),
// 		GasFeeCap: big.NewInt(5),
// 		GasTipCap: big.NewInt(6),
// 		GasLimit:  7,
// 		Context:   context.TODO(),
// 		NoSend:    false,
// 	}

// 	b := copyTxOpts(a)

// 	require.Equal(t, a.From, b.From)
// 	require.Equal(t, a.Nonce, b.Nonce)
// 	require.Equal(t, a.Value, b.Value)
// 	require.Equal(t, a.GasPrice, b.GasPrice)
// 	require.Equal(t, a.GasFeeCap, b.GasFeeCap)
// 	require.Equal(t, a.GasTipCap, b.GasTipCap)
// 	require.Equal(t, a.GasLimit, b.GasLimit)
// 	require.Equal(t, a.Context, b.Context)
// 	require.Equal(t, a.NoSend, b.NoSend)

// 	// Make changes like SetBytes which modify the underlying values.

// 	b.From.SetBytes([]byte("foobar"))
// 	b.Nonce.SetBytes([]byte("foobar"))
// 	b.Value.SetBytes([]byte("foobar"))
// 	b.GasPrice.SetBytes([]byte("foobar"))
// 	b.GasFeeCap.SetBytes([]byte("foobar"))
// 	b.GasTipCap.SetBytes([]byte("foobar"))
// 	b.GasLimit = 123456789
// 	type foo string // custom type for linter.
// 	b.Context = context.WithValue(context.TODO(), foo("bar"), foo("baz"))
// 	b.NoSend = true

// 	// Everything should be different.
// 	// Note: signer is not evaluated because function comparison is not possible.
// 	require.NotEqual(t, a.From, b.From)
// 	require.NotEqual(t, a.Nonce, b.Nonce)
// 	require.NotEqual(t, a.Value, b.Value)
// 	require.NotEqual(t, a.GasPrice, b.GasPrice)
// 	require.NotEqual(t, a.GasFeeCap, b.GasFeeCap)
// 	require.NotEqual(t, a.GasTipCap, b.GasTipCap)
// 	require.NotEqual(t, a.GasLimit, b.GasLimit)
// 	require.NotEqual(t, a.Context, b.Context)
// 	require.NotEqual(t, a.NoSend, b.NoSend)
// }

// func setupTopLevelFork(
// 	t *testing.T,
// 	ctx context.Context,
// 	height1,
// 	height2 uint64,
// ) (*Assertion, *Assertion, *Challenge, *AssertionChain, *AssertionChain) {
// 	t.Helper()
// 	chain1, accs, addresses, backend, headerReader := setupAssertionChainWithChallengeManager(t)
// 	prev := uint64(0)

// 	minAssertionPeriod, err := chain1.userLogic.MinimumAssertionPeriod(chain1.callOpts)
// 	require.NoError(t, err)

// 	latestBlockHash := common.Hash{}
// 	for i := uint64(0); i < minAssertionPeriod.Uint64(); i++ {
// 		latestBlockHash = backend.Commit()
// 	}

// 	prevState := &protocol.ExecutionState{
// 		GlobalState:   protocol.GoGlobalState{},
// 		MachineStatus: protocol.MachineStatusFinished,
// 	}
// 	postState := &protocol.ExecutionState{
// 		GlobalState: protocol.GoGlobalState{
// 			BlockHash:  latestBlockHash,
// 			SendRoot:   common.Hash{},
// 			Batch:      1,
// 			PosInBatch: 0,
// 		},
// 		MachineStatus: protocol.MachineStatusFinished,
// 	}
// 	prevInboxMaxCount := big.NewInt(1)
// 	a1, err := chain1.CreateAssertion(ctx, height1, protocol.AssertionSequenceNumber(prev), prevState, postState, prevInboxMaxCount)
// 	require.NoError(t, err)

// 	chain2, err := NewAssertionChain(
// 		ctx,
// 		addresses.Rollup,
// 		accs[2].txOpts,
// 		&bind.CallOpts{},
// 		accs[2].accountAddr,
// 		backend,
// 		headerReader,
// 		common.Address{},
// 	)
// 	require.NoError(t, err)

// 	postState.GlobalState.BlockHash = common.BytesToHash([]byte("evil"))
// 	a2, err := chain2.CreateAssertion(ctx, height2, protocol.AssertionSequenceNumber(prev), prevState, postState, prevInboxMaxCount)
// 	require.NoError(t, err)

// 	challenge, err := chain2.CreateSuccessionChallenge(ctx, 0)
// 	require.NoError(t, err)
// 	return a1.(*Assertion), a2.(*Assertion), challenge.(*Challenge), chain1, chain2
// }
