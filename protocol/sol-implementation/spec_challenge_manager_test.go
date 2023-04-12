package solimpl_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/protocol/sol-implementation"
	"github.com/OffchainLabs/challenge-protocol-v2/state-manager"
	"github.com/OffchainLabs/challenge-protocol-v2/testing/setup"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

var (
	_ = protocol.SpecEdge(&solimpl.SpecEdge{})
	_ = protocol.SpecChallengeManager(&solimpl.SpecChallengeManager{})
)

func TestEdgeChallengeManager_IsUnrivaled(t *testing.T) {
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

	// Honest assertion being added.
	leafAdder := func(startCommit, endCommit util.HistoryCommitment, leaf protocol.Assertion) protocol.SpecEdge {
		edge, err := challengeManager.AddBlockChallengeLevelZeroEdge(
			ctx,
			leaf,
			startCommit,
			endCommit,
		)
		require.NoError(t, err)
		return edge
	}

	honestStartCommit, err := honestStateManager.HistoryCommitmentUpTo(ctx, 0)
	require.NoError(t, err)
	honestEndCommit, err := honestStateManager.HistoryCommitmentUpTo(ctx, uint64(height))
	require.NoError(t, err)

	honestEdge := leafAdder(honestStartCommit, honestEndCommit, createdData.Leaf1)
	require.Equal(t, protocol.BlockChallengeEdge, honestEdge.GetType())

	t.Run("first leaf is presumptive", func(t *testing.T) {
		hasRival, err := honestEdge.HasRival(ctx)
		require.NoError(t, err)
		require.Equal(t, true, !hasRival)
	})

	evilStartCommit, err := evilStateManager.HistoryCommitmentUpTo(ctx, 0)
	require.NoError(t, err)
	evilEndCommit, err := evilStateManager.HistoryCommitmentUpTo(ctx, uint64(height))
	require.NoError(t, err)

	evilEdge := leafAdder(evilStartCommit, evilEndCommit, createdData.Leaf2)
	require.Equal(t, protocol.BlockChallengeEdge, evilEdge.GetType())

	t.Run("neither is presumptive if rivals", func(t *testing.T) {
		hasRival, err := honestEdge.HasRival(ctx)
		require.NoError(t, err)
		require.Equal(t, false, !hasRival)

		hasRival, err = evilEdge.HasRival(ctx)
		require.NoError(t, err)
		require.Equal(t, false, !hasRival)
	})

	t.Run("bisected children are presumptive", func(t *testing.T) {
		honestBisectCommit, err := honestStateManager.HistoryCommitmentUpTo(ctx, 2)
		require.NoError(t, err)
		honestProof, err := honestStateManager.PrefixProof(ctx, 2, 3)
		require.NoError(t, err)

		lower, upper, err := honestEdge.Bisect(ctx, honestBisectCommit.Merkle, honestProof)
		require.NoError(t, err)

		hasRival, err := lower.HasRival(ctx)
		require.NoError(t, err)
		require.Equal(t, true, !hasRival)
		hasRival, err = upper.HasRival(ctx)
		require.NoError(t, err)
		require.Equal(t, true, !hasRival)

		hasRival, err = honestEdge.HasRival(ctx)
		require.NoError(t, err)
		require.Equal(t, false, !hasRival)

		hasRival, err = evilEdge.HasRival(ctx)
		require.NoError(t, err)
		require.Equal(t, false, !hasRival)
	})
}

func TestEdgeChallengeManager_HasLengthOneRival(t *testing.T) {
	ctx := context.Background()
	topLevelHeight := protocol.Height(3)
	bisectionScenario := setupBisectionScenario(
		t,
		topLevelHeight,
		statemanager.WithNumOpcodesPerBigStep(1),
		statemanager.WithMaxWavmOpcodesPerBlock(1),
	)
	honestStateManager := bisectionScenario.honestStateManager
	evilStateManager := bisectionScenario.evilStateManager
	honestEdge := bisectionScenario.honestLevelZeroEdge
	evilEdge := bisectionScenario.evilLevelZeroEdge

	t.Run("level zero edge with rivals is not one step fork source", func(t *testing.T) {
		isOSF, err := honestEdge.HasLengthOneRival(ctx)
		require.NoError(t, err)
		require.Equal(t, false, isOSF)
		isOSF, err = evilEdge.HasLengthOneRival(ctx)
		require.NoError(t, err)
		require.Equal(t, false, isOSF)
	})
	t.Run("post bisection, mutual edge is one step fork source", func(t *testing.T) {
		honestBisectCommit, err := honestStateManager.HistoryCommitmentUpTo(ctx, 2)
		require.NoError(t, err)
		honestProof, err := honestStateManager.PrefixProof(ctx, 2, 3)
		require.NoError(t, err)
		lowerHonest, upper, err := honestEdge.Bisect(ctx, honestBisectCommit.Merkle, honestProof)
		require.NoError(t, err)

		isOSF, err := lowerHonest.HasLengthOneRival(ctx)
		require.NoError(t, err)
		require.Equal(t, false, isOSF)
		isOSF, err = upper.HasLengthOneRival(ctx)
		require.NoError(t, err)
		require.Equal(t, false, isOSF)

		evilBisectCommit, err := evilStateManager.HistoryCommitmentUpTo(ctx, 2)
		require.NoError(t, err)
		evilProof, err := evilStateManager.PrefixProof(ctx, 2, 3)
		require.NoError(t, err)
		lower, _, err := evilEdge.Bisect(ctx, evilBisectCommit.Merkle, evilProof)
		require.NoError(t, err)

		evilBisectCommit, err = evilStateManager.HistoryCommitmentUpTo(ctx, 1)
		require.NoError(t, err)
		evilProof, err = evilStateManager.PrefixProof(ctx, 1, 2)
		require.NoError(t, err)
		_, _, err = lower.Bisect(ctx, evilBisectCommit.Merkle, evilProof)
		require.NoError(t, err)

		honestBisectCommit, err = honestStateManager.HistoryCommitmentUpTo(ctx, 1)
		require.NoError(t, err)
		honestProof, err = honestStateManager.PrefixProof(ctx, 1, 2)
		require.NoError(t, err)
		lower, upper, err = lowerHonest.Bisect(ctx, honestBisectCommit.Merkle, honestProof)
		require.NoError(t, err)

		isOSF, err = lower.HasLengthOneRival(ctx)
		require.NoError(t, err)
		require.Equal(t, true, isOSF)

		isOSF, err = upper.HasLengthOneRival(ctx)
		require.NoError(t, err)
		require.Equal(t, false, isOSF)
	})
}

func TestEdgeChallengeManager_BlockChallengeAddLevelZeroEdge(t *testing.T) {
	ctx := context.Background()
	height := uint64(3)
	createdData, err := setup.CreateTwoValidatorFork(ctx, &setup.CreateForkConfig{
		NumBlocks:     height,
		DivergeHeight: 0,
	})
	require.NoError(t, err)

	chain1 := createdData.Chains[0]
	challengeManager, err := chain1.SpecChallengeManager(ctx)
	require.NoError(t, err)

	t.Run("claim predecessor does not exist", func(t *testing.T) {
		t.Skip("Needs Solidity code")
	})
	t.Run("invalid height", func(t *testing.T) {
		t.Skip("Needs Solidity code")
	})
	t.Run("last state is not assertion claim block hash", func(t *testing.T) {
		t.Skip("Needs Solidity code")
	})
	t.Run("winner already declared", func(t *testing.T) {
		t.Skip("Needs Solidity code")
	})
	t.Run("last state not in history", func(t *testing.T) {
		t.Skip("Needs Solidity code")
	})
	t.Run("first state not in history", func(t *testing.T) {
		t.Skip("Needs Solidity code")
	})

	leaves := make([]common.Hash, 4)
	for i := range leaves {
		leaves[i] = crypto.Keccak256Hash([]byte(fmt.Sprintf("%d", i)))
	}
	opts := []statemanager.Opt{
		statemanager.WithNumOpcodesPerBigStep(1),
		statemanager.WithMaxWavmOpcodesPerBlock(1),
	}

	honestStateManager, err := statemanager.New(
		createdData.HonestValidatorStateRoots,
		opts...,
	)
	require.NoError(t, err)

	start, err := honestStateManager.HistoryCommitmentUpTo(ctx, 0)
	require.NoError(t, err)
	end, err := honestStateManager.HistoryCommitmentUpTo(ctx, height)
	require.NoError(t, err)

	t.Run("OK", func(t *testing.T) {
		_, err = challengeManager.AddBlockChallengeLevelZeroEdge(ctx, createdData.Leaf1, start, end)
		require.NoError(t, err)
	})
	t.Run("already exists", func(t *testing.T) {
		_, err = challengeManager.AddBlockChallengeLevelZeroEdge(ctx, createdData.Leaf1, start, end)
		require.ErrorContains(t, err, "already exists")
	})
}

func TestEdgeChallengeManager_Bisect(t *testing.T) {
	ctx := context.Background()
	topLevelHeight := protocol.Height(3)
	bisectionScenario := setupBisectionScenario(
		t,
		topLevelHeight,
		statemanager.WithNumOpcodesPerBigStep(1),
		statemanager.WithMaxWavmOpcodesPerBlock(1),
	)
	honestStateManager := bisectionScenario.honestStateManager
	honestEdge := bisectionScenario.honestLevelZeroEdge

	t.Run("cannot bisect unrivaled", func(t *testing.T) {
		t.Skip("TODO(RJ): Implement")
	})
	t.Run("invalid prefix proof", func(t *testing.T) {
		t.Skip("TODO(RJ): Implement")
	})
	t.Run("edge has children", func(t *testing.T) {
		t.Skip("TODO(RJ): Implement")
	})
	t.Run("OK", func(t *testing.T) {
		honestBisectCommit, err := honestStateManager.HistoryCommitmentUpTo(ctx, 2)
		require.NoError(t, err)
		honestProof, err := honestStateManager.PrefixProof(ctx, 2, 3)
		require.NoError(t, err)
		_, _, err = honestEdge.Bisect(ctx, honestBisectCommit.Merkle, honestProof)
		require.NoError(t, err)
	})
}

func TestEdgeChallengeManager_SubChallenges(t *testing.T) {
	t.Run("leaf cannot be a fork candidate", func(t *testing.T) {
		t.Skip("TODO(RJ): Implement")
	})
	t.Run("lowest height not one step fork", func(t *testing.T) {
		t.Skip("TODO(RJ): Implement")
	})
	t.Run("has presumptive successor", func(t *testing.T) {
		t.Skip("TODO(RJ): Implement")
	})
	t.Run("empty history root", func(t *testing.T) {
		t.Skip("TODO(RJ): Implement")
	})
	t.Run("OK", func(t *testing.T) {
		t.Skip("TODO(RJ): Implement")
	})
}

func TestEdgeChallengeManager_ConfirmByOneStepProof(t *testing.T) {
	ctx := context.Background()
	topLevelHeight := protocol.Height(2)
	// t.Run("edge does not exist", func(t *testing.T) {
	// 	bisectionScenario := setupBisectionScenario(
	// 		t,
	// 		topLevelHeight,
	// 		statemanager.WithNumOpcodesPerBigStep(1),
	// 		statemanager.WithMaxWavmOpcodesPerBlock(1),
	// 	)
	// 	challengeManager, err := bisectionScenario.topLevelFork.Chains[1].SpecChallengeManager(ctx)
	// 	require.NoError(t, err)
	// 	err = challengeManager.ConfirmEdgeByOneStepProof(
	// 		ctx,
	// 		protocol.EdgeId(common.BytesToHash([]byte("foo"))),
	// 		&protocol.OneStepData{
	// 			BridgeAddr:           common.Address{},
	// 			MaxInboxMessagesRead: 2,
	// 			MachineStep:          0,
	// 			BeforeHash:           common.Hash{},
	// 			Proof:                make([]byte, 0),
	// 		},
	// 		make([]common.Hash, 0),
	// 		make([]common.Hash, 0),
	// 	)
	// 	require.ErrorContains(t, err, "Edge does not exist")
	// })
	// t.Run("edge not pending", func(t *testing.T) {
	// 	bisectionScenario := setupBisectionScenario(
	// 		t,
	// 		topLevelHeight,
	// 		statemanager.WithNumOpcodesPerBigStep(1),
	// 		statemanager.WithMaxWavmOpcodesPerBlock(1),
	// 	)
	// 	honestStateManager := bisectionScenario.honestStateManager
	// 	honestEdge := bisectionScenario.honestLevelZeroEdge
	// 	challengeManager, err := bisectionScenario.topLevelFork.Chains[1].SpecChallengeManager(ctx)
	// 	require.NoError(t, err)

	// 	honestBisectCommit, err := honestStateManager.HistoryCommitmentUpTo(ctx, 1)
	// 	require.NoError(t, err)
	// 	honestProof, err := honestStateManager.PrefixProof(ctx, 1, 2)
	// 	require.NoError(t, err)
	// 	honestChildren1, honestChildren2, err := honestEdge.Bisect(ctx, honestBisectCommit.Merkle, honestProof)
	// 	require.NoError(t, err)

	// 	s1, err := honestChildren1.Status(ctx)
	// 	require.NoError(t, err)
	// 	require.Equal(t, protocol.EdgePending, s1)
	// 	s2, err := honestChildren2.Status(ctx)
	// 	require.NoError(t, err)
	// 	require.Equal(t, protocol.EdgePending, s2)

	// 	// Adjust well beyond a challenge period.
	// 	for i := 0; i < 100; i++ {
	// 		bisectionScenario.topLevelFork.Backend.Commit()
	// 	}

	// 	require.NoError(t, honestChildren1.ConfirmByTimer(ctx, []protocol.EdgeId{}))
	// 	require.NoError(t, honestChildren2.ConfirmByTimer(ctx, []protocol.EdgeId{}))
	// 	s1, err = honestChildren1.Status(ctx)
	// 	require.NoError(t, err)
	// 	require.Equal(t, protocol.EdgeConfirmed, s1)
	// 	s2, err = honestChildren2.Status(ctx)
	// 	require.NoError(t, err)
	// 	require.Equal(t, protocol.EdgeConfirmed, s2)

	// 	err = challengeManager.ConfirmEdgeByOneStepProof(
	// 		ctx,
	// 		honestChildren1.Id(),
	// 		&protocol.OneStepData{
	// 			BridgeAddr:           common.Address{},
	// 			MaxInboxMessagesRead: 2,
	// 			MachineStep:          0,
	// 			BeforeHash:           common.Hash{},
	// 			Proof:                make([]byte, 0),
	// 		},
	// 		make([]common.Hash, 0),
	// 		make([]common.Hash, 0),
	// 	)
	// 	require.ErrorContains(t, err, "Edge not pending")
	// 	err = challengeManager.ConfirmEdgeByOneStepProof(
	// 		ctx,
	// 		honestChildren2.Id(),
	// 		&protocol.OneStepData{
	// 			BridgeAddr:           common.Address{},
	// 			MaxInboxMessagesRead: 2,
	// 			MachineStep:          0,
	// 			BeforeHash:           common.Hash{},
	// 			Proof:                make([]byte, 0),
	// 		},
	// 		make([]common.Hash, 0),
	// 		make([]common.Hash, 0),
	// 	)
	// 	require.ErrorContains(t, err, "Edge not pending")
	// })
	// t.Run("edge not small step type", func(t *testing.T) {
	// 	bisectionScenario := setupBisectionScenario(
	// 		t,
	// 		topLevelHeight,
	// 		statemanager.WithNumOpcodesPerBigStep(1),
	// 		statemanager.WithMaxWavmOpcodesPerBlock(1),
	// 	)
	// 	honestStateManager := bisectionScenario.honestStateManager
	// 	honestEdge := bisectionScenario.honestLevelZeroEdge
	// 	challengeManager, err := bisectionScenario.topLevelFork.Chains[1].SpecChallengeManager(ctx)
	// 	require.NoError(t, err)

	// 	honestBisectCommit, err := honestStateManager.HistoryCommitmentUpTo(ctx, 1)
	// 	require.NoError(t, err)
	// 	honestProof, err := honestStateManager.PrefixProof(ctx, 1, 2)
	// 	require.NoError(t, err)
	// 	honestChildren1, honestChildren2, err := honestEdge.Bisect(ctx, honestBisectCommit.Merkle, honestProof)
	// 	require.NoError(t, err)

	// 	s1, err := honestChildren1.Status(ctx)
	// 	require.NoError(t, err)
	// 	require.Equal(t, protocol.EdgePending, s1)
	// 	s2, err := honestChildren2.Status(ctx)
	// 	require.NoError(t, err)
	// 	require.Equal(t, protocol.EdgePending, s2)

	// 	err = challengeManager.ConfirmEdgeByOneStepProof(
	// 		ctx,
	// 		honestChildren1.Id(),
	// 		&protocol.OneStepData{
	// 			BridgeAddr:           common.Address{},
	// 			MaxInboxMessagesRead: 2,
	// 			MachineStep:          0,
	// 			BeforeHash:           common.Hash{},
	// 			Proof:                make([]byte, 0),
	// 		},
	// 		make([]common.Hash, 0),
	// 		make([]common.Hash, 0),
	// 	)
	// 	require.ErrorContains(t, err, "Edge is not a small step")
	// 	err = challengeManager.ConfirmEdgeByOneStepProof(
	// 		ctx,
	// 		honestChildren2.Id(),
	// 		&protocol.OneStepData{
	// 			BridgeAddr:           common.Address{},
	// 			MaxInboxMessagesRead: 2,
	// 			MachineStep:          0,
	// 			BeforeHash:           common.Hash{},
	// 			Proof:                make([]byte, 0),
	// 		},
	// 		make([]common.Hash, 0),
	// 		make([]common.Hash, 0),
	// 	)
	// 	require.ErrorContains(t, err, "Edge is not a small step")
	// })
	// t.Run("edge does not have a single step rival", func(t *testing.T) {
	// 	ctx := context.Background()
	// 	bisectionScenario := setupBisectionScenario(
	// 		t,
	// 		topLevelHeight,
	// 		// We want four opcodes per big step.
	// 		statemanager.WithNumOpcodesPerBigStep(4),
	// 		// We want a total of one big step.
	// 		statemanager.WithMaxWavmOpcodesPerBlock(4),
	// 	)
	// 	honestStateManager := bisectionScenario.honestStateManager
	// 	evilStateManager := bisectionScenario.evilStateManager
	// 	honestEdge := bisectionScenario.honestLevelZeroEdge
	// 	evilEdge := bisectionScenario.evilLevelZeroEdge
	// 	honestStartCommit := bisectionScenario.honestStartCommit
	// 	evilStartCommit := bisectionScenario.evilStartCommit

	// 	challengeManager, err := bisectionScenario.topLevelFork.Chains[1].SpecChallengeManager(ctx)
	// 	require.NoError(t, err)

	// 	// Attempt bisections down to one step fork.
	// 	honestBisectCommit, err := honestStateManager.HistoryCommitmentUpTo(ctx, 1)
	// 	require.NoError(t, err)
	// 	honestProof, err := honestStateManager.PrefixProof(ctx, 1, 2)
	// 	require.NoError(t, err)

	// 	_, _, err = honestEdge.Bisect(ctx, honestBisectCommit.Merkle, honestProof)
	// 	require.NoError(t, err)

	// 	evilBisectCommit, err := evilStateManager.HistoryCommitmentUpTo(ctx, 1)
	// 	require.NoError(t, err)
	// 	evilProof, err := evilStateManager.PrefixProof(ctx, 1, 2)
	// 	require.NoError(t, err)

	// 	oneStepForkSourceEdge, _, err := evilEdge.Bisect(ctx, evilBisectCommit.Merkle, evilProof)
	// 	require.NoError(t, err)

	// 	isAtOneStepFork, err := oneStepForkSourceEdge.HasLengthOneRival(ctx)
	// 	require.NoError(t, err)
	// 	require.Equal(t, true, isAtOneStepFork)

	// 	// Now opening big step level zero leaves
	// 	bigStepAdder := func(startCommit, endCommit util.HistoryCommitment) protocol.SpecEdge {
	// 		leaf, err := challengeManager.AddSubChallengeLevelZeroEdge(
	// 			ctx,
	// 			oneStepForkSourceEdge,
	// 			startCommit,
	// 			endCommit,
	// 		)
	// 		require.NoError(t, err)
	// 		return leaf
	// 	}

	// 	honestBigStepCommit, err := honestStateManager.BigStepCommitmentUpTo(
	// 		ctx, 0 /* from assertion */, 1 /* to assertion */, 1, /* to big step */
	// 	)
	// 	require.NoError(t, err)

	// 	honestEdge = bigStepAdder(honestStartCommit, honestBigStepCommit)
	// 	require.Equal(t, protocol.BigStepChallengeEdge, honestEdge.GetType())
	// 	hasRival, err := honestEdge.HasRival(ctx)
	// 	require.NoError(t, err)
	// 	require.Equal(t, true, !hasRival)

	// 	evilBigStepCommit, err := evilStateManager.BigStepCommitmentUpTo(
	// 		ctx, 0 /* from assertion */, 1 /* to assertion */, 1, /* to big step */
	// 	)
	// 	require.NoError(t, err)

	// 	evilEdge = bigStepAdder(evilStartCommit, evilBigStepCommit)
	// 	require.Equal(t, protocol.BigStepChallengeEdge, evilEdge.GetType())

	// 	hasRival, err = honestEdge.HasRival(ctx)
	// 	require.NoError(t, err)
	// 	require.Equal(t, false, !hasRival)
	// 	hasRival, err = evilEdge.HasRival(ctx)
	// 	require.NoError(t, err)
	// 	require.Equal(t, false, !hasRival)

	// 	isAtOneStepFork, err = honestEdge.HasLengthOneRival(ctx)
	// 	require.NoError(t, err)
	// 	require.Equal(t, true, isAtOneStepFork)

	// 	// Now opening small step level zero leaves
	// 	smallStepAdder := func(startCommit, endCommit util.HistoryCommitment) protocol.SpecEdge {
	// 		leaf, err := challengeManager.AddSubChallengeLevelZeroEdge(
	// 			ctx,
	// 			honestEdge,
	// 			startCommit,
	// 			endCommit,
	// 		)
	// 		require.NoError(t, err)
	// 		return leaf
	// 	}

	// 	honestSmallStepCommit, err := honestStateManager.SmallStepCommitmentUpTo(
	// 		ctx, 0 /* from assertion */, 1 /* to assertion */, 3, /* to pc */
	// 	)
	// 	require.NoError(t, err)

	// 	smallStepHonest := smallStepAdder(honestStartCommit, honestSmallStepCommit)
	// 	require.Equal(t, protocol.SmallStepChallengeEdge, smallStepHonest.GetType())
	// 	hasRival, err = smallStepHonest.HasRival(ctx)
	// 	require.NoError(t, err)
	// 	require.Equal(t, true, !hasRival)

	// 	evilSmallStepCommit, err := evilStateManager.SmallStepCommitmentUpTo(
	// 		ctx, 0 /* from assertion */, 1 /* to assertion */, 3, /* to pc */
	// 	)
	// 	require.NoError(t, err)

	// 	smallStepEvil := smallStepAdder(evilStartCommit, evilSmallStepCommit)
	// 	require.Equal(t, protocol.SmallStepChallengeEdge, smallStepEvil.GetType())

	// 	hasRival, err = smallStepHonest.HasRival(ctx)
	// 	require.NoError(t, err)
	// 	require.Equal(t, false, !hasRival)
	// 	hasRival, err = smallStepEvil.HasRival(ctx)
	// 	require.NoError(t, err)
	// 	require.Equal(t, false, !hasRival)

	// 	isAtOneStepFork, err = smallStepHonest.HasLengthOneRival(ctx)
	// 	require.NoError(t, err)
	// 	require.Equal(t, false, isAtOneStepFork)

	// 	err = challengeManager.ConfirmEdgeByOneStepProof(
	// 		ctx,
	// 		honestEdge.Id(),
	// 		&protocol.OneStepData{
	// 			BridgeAddr:           common.Address{},
	// 			MaxInboxMessagesRead: 2,
	// 			MachineStep:          0,
	// 			BeforeHash:           honestSmallStepCommit.FirstLeaf,
	// 			Proof:                honestSmallStepCommit.LastLeaf[:],
	// 		},
	// 		honestSmallStepCommit.FirstLeafProof,
	// 		honestSmallStepCommit.LastLeafProof,
	// 	)
	// 	require.ErrorContains(t, err, "Edge is not a small step")
	// })
	// t.Run("before state not in history", func(t *testing.T) {
	// 	scenario := setupOneStepProofScenario(
	// 		t,
	// 		topLevelHeight,
	// 		statemanager.WithNumOpcodesPerBigStep(1),
	// 		statemanager.WithMaxWavmOpcodesPerBlock(1),
	// 	)
	// 	honestEdge := scenario.smallStepHonestEdge

	// 	challengeManager, err := scenario.topLevelFork.Chains[1].SpecChallengeManager(ctx)
	// 	require.NoError(t, err)

	// 	honestStateManager := scenario.honestStateManager
	// 	fromAssertion := uint64(0)
	// 	toAssertion := uint64(1)
	// 	honestCommit, err := honestStateManager.SmallStepCommitmentUpTo(
	// 		ctx,
	// 		fromAssertion,
	// 		toAssertion,
	// 		1,
	// 	)
	// 	require.NoError(t, err)

	// 	err = challengeManager.ConfirmEdgeByOneStepProof(
	// 		ctx,
	// 		honestEdge.Id(),
	// 		&protocol.OneStepData{
	// 			BridgeAddr:           common.Address{},
	// 			MaxInboxMessagesRead: 2,
	// 			MachineStep:          0,
	// 			BeforeHash:           honestCommit.FirstLeaf,
	// 			Proof:                honestCommit.LastLeaf[:],
	// 		},
	// 		honestCommit.FirstLeafProof,
	// 		honestCommit.LastLeafProof,
	// 	)
	// 	require.ErrorContains(t, err, "Before state not in history")
	// })
	// t.Run("one step proof fails", func(t *testing.T) {
	// 	scenario := setupOneStepProofScenario(
	// 		t,
	// 		topLevelHeight,
	// 		statemanager.WithNumOpcodesPerBigStep(1),
	// 		statemanager.WithMaxWavmOpcodesPerBlock(1),
	// 	)
	// 	honestEdge := scenario.smallStepHonestEdge

	// 	challengeManager, err := scenario.topLevelFork.Chains[1].SpecChallengeManager(ctx)
	// 	require.NoError(t, err)

	// 	honestStateManager := scenario.honestStateManager
	// 	fromAssertion := uint64(0)
	// 	toAssertion := uint64(1)
	// 	honestStartCommit, err := honestStateManager.SmallStepCommitmentUpTo(
	// 		ctx,
	// 		fromAssertion,
	// 		toAssertion,
	// 		0,
	// 	)
	// 	require.NoError(t, err)
	// 	honestEndCommit, err := honestStateManager.SmallStepCommitmentUpTo(
	// 		ctx,
	// 		fromAssertion,
	// 		toAssertion,
	// 		1,
	// 	)
	// 	require.NoError(t, err)

	// 	err = challengeManager.ConfirmEdgeByOneStepProof(
	// 		ctx,
	// 		honestEdge.Id(),
	// 		&protocol.OneStepData{
	// 			BridgeAddr:           common.Address{},
	// 			MaxInboxMessagesRead: 2,
	// 			MachineStep:          0,
	// 			BeforeHash:           honestStartCommit.FirstLeaf,
	// 			Proof:                make([]byte, 0),
	// 		},
	// 		honestStartCommit.FirstLeafProof,
	// 		honestEndCommit.LastLeafProof,
	// 	)
	// 	require.ErrorContains(t, err, "After state not in history")
	// })
	t.Run("OK", func(t *testing.T) {
		scenario := setupOneStepProofScenario(
			t,
			topLevelHeight,
			statemanager.WithNumOpcodesPerBigStep(4),
			statemanager.WithMaxWavmOpcodesPerBlock(4),
		)
		honestEdge := scenario.smallStepHonestEdge

		challengeManager, err := scenario.topLevelFork.Chains[1].SpecChallengeManager(ctx)
		require.NoError(t, err)

		honestStateManager := scenario.honestStateManager
		fromAssertion := uint64(0)
		toAssertion := uint64(1)
		honestStartCommit, err := honestStateManager.SmallStepCommitmentUpTo(
			ctx,
			fromAssertion,
			toAssertion,
			0,
		)
		require.NoError(t, err)
		honestEndCommit, err := honestStateManager.SmallStepCommitmentUpTo(
			ctx,
			fromAssertion,
			toAssertion,
			1,
		)
		require.NoError(t, err)

		err = challengeManager.ConfirmEdgeByOneStepProof(
			ctx,
			honestEdge.Id(),
			&protocol.OneStepData{
				BridgeAddr:           common.Address{},
				MaxInboxMessagesRead: 2,
				MachineStep:          0,
				BeforeHash:           honestStartCommit.FirstLeaf,
				Proof:                honestEndCommit.LastLeaf[:],
			},
			honestStartCommit.FirstLeafProof,
			honestEndCommit.LastLeafProof,
		)
		require.NoError(t, err)
		edgeStatus, err := honestEdge.Status(ctx)
		require.NoError(t, err)
		require.Equal(t, protocol.EdgeConfirmed, edgeStatus)
	})
}

func TestEdgeChallengeManager_ConfirmByTimerAndChildren(t *testing.T) {
	ctx := context.Background()
	topLevelHeight := protocol.Height(3)
	bisectionScenario := setupBisectionScenario(
		t,
		topLevelHeight,
		statemanager.WithNumOpcodesPerBigStep(1),
		statemanager.WithMaxWavmOpcodesPerBlock(1),
	)
	honestStateManager := bisectionScenario.honestStateManager
	honestEdge := bisectionScenario.honestLevelZeroEdge

	honestBisectCommit, err := honestStateManager.HistoryCommitmentUpTo(ctx, 2)
	require.NoError(t, err)
	honestProof, err := honestStateManager.PrefixProof(ctx, 2, 3)
	require.NoError(t, err)
	honestChildren1, honestChildren2, err := honestEdge.Bisect(ctx, honestBisectCommit.Merkle, honestProof)
	require.NoError(t, err)

	s1, err := honestChildren1.Status(ctx)
	require.NoError(t, err)
	require.Equal(t, protocol.EdgePending, s1)
	s2, err := honestChildren2.Status(ctx)
	require.NoError(t, err)
	require.Equal(t, protocol.EdgePending, s2)

	// Adjust well beyond a challenge period.
	for i := 0; i < 100; i++ {
		bisectionScenario.topLevelFork.Backend.Commit()
	}

	require.NoError(t, honestChildren1.ConfirmByTimer(ctx, []protocol.EdgeId{}))
	require.NoError(t, honestChildren2.ConfirmByTimer(ctx, []protocol.EdgeId{}))
	s1, err = honestChildren1.Status(ctx)
	require.NoError(t, err)
	require.Equal(t, protocol.EdgeConfirmed, s1)
	s2, err = honestChildren2.Status(ctx)
	require.NoError(t, err)
	require.Equal(t, protocol.EdgeConfirmed, s2)

	require.NoError(t, honestEdge.ConfirmByChildren(ctx))
	s0, err := honestEdge.Status(ctx)
	require.NoError(t, err)
	require.Equal(t, protocol.EdgeConfirmed, s0)
}

func TestEdgeChallengeManager_ConfirmByTimer(t *testing.T) {
	ctx := context.Background()
	height := protocol.Height(3)

	createdData, err := setup.CreateTwoValidatorFork(ctx, &setup.CreateForkConfig{
		NumBlocks:     uint64(height) + 1,
		DivergeHeight: 0,
	})
	require.NoError(t, err)

	honestStateManager, err := statemanager.New(
		createdData.HonestValidatorStateRoots,
		statemanager.WithNumOpcodesPerBigStep(1),
		statemanager.WithMaxWavmOpcodesPerBlock(1),
	)
	require.NoError(t, err)

	challengeManager, err := createdData.Chains[0].SpecChallengeManager(ctx)
	require.NoError(t, err)

	// Honest assertion being added.
	leafAdder := func(startCommit, endCommit util.HistoryCommitment, leaf protocol.Assertion) protocol.SpecEdge {
		edge, err := challengeManager.AddBlockChallengeLevelZeroEdge(
			ctx,
			leaf,
			startCommit,
			endCommit,
		)
		require.NoError(t, err)
		return edge
	}
	honestStartCommit, err := honestStateManager.HistoryCommitmentUpTo(ctx, 0)
	require.NoError(t, err)
	honestEndCommit, err := honestStateManager.HistoryCommitmentUpTo(ctx, uint64(height))
	require.NoError(t, err)
	honestEdge := leafAdder(honestStartCommit, honestEndCommit, createdData.Leaf1)
	s0, err := honestEdge.Status(ctx)
	require.NoError(t, err)
	require.Equal(t, protocol.EdgePending, s0)

	hasRival, err := honestEdge.HasRival(ctx)
	require.NoError(t, err)
	require.Equal(t, false, hasRival)

	// Adjust well beyond a challenge period.
	for i := 0; i < 100; i++ {
		createdData.Backend.Commit()
	}

	t.Run("confirmed by timer", func(t *testing.T) {
		require.ErrorContains(t, honestEdge.ConfirmByTimer(ctx, []protocol.EdgeId{protocol.EdgeId(common.Hash{1})}), "execution reverted: Edge does not exist")
	})
	t.Run("confirmed by timer", func(t *testing.T) {
		require.NoError(t, honestEdge.ConfirmByTimer(ctx, []protocol.EdgeId{}))
		status, err := honestEdge.Status(ctx)
		require.NoError(t, err)
		require.Equal(t, protocol.EdgeConfirmed, status)
	})
	t.Run("cannot confirm again", func(t *testing.T) {
		status, err := honestEdge.Status(ctx)
		require.NoError(t, err)
		require.Equal(t, protocol.EdgeConfirmed, status)
		require.ErrorContains(t, honestEdge.ConfirmByTimer(ctx, []protocol.EdgeId{}), "execution reverted: Edge not pending")
	})
}

func TestEdgeChallengeManager_ReachesOneStepProof(t *testing.T) {
	setupOneStepProofScenario(
		t,
		protocol.Height(2),
		statemanager.WithNumOpcodesPerBigStep(4),
		statemanager.WithMaxWavmOpcodesPerBlock(4),
	)
}

// Returns a snapshot of the data for a scenario in which both honest
// and evil validator validators have created level zero edges in a top-level
// challenge and are ready to bisect.
type bisectionScenario struct {
	topLevelFork        *setup.CreatedValidatorFork
	honestStateManager  statemanager.Manager
	evilStateManager    statemanager.Manager
	honestLevelZeroEdge protocol.SpecEdge
	evilLevelZeroEdge   protocol.SpecEdge
	honestStartCommit   util.HistoryCommitment
	evilStartCommit     util.HistoryCommitment
}

func setupBisectionScenario(
	t *testing.T,
	topLevelHeight protocol.Height,
	commonStateManagerOpts ...statemanager.Opt,
) *bisectionScenario {
	ctx := context.Background()

	createdData, err := setup.CreateTwoValidatorFork(ctx, &setup.CreateForkConfig{
		NumBlocks:     uint64(topLevelHeight) + 1,
		DivergeHeight: 0,
	})
	require.NoError(t, err)

	honestStateManager, err := statemanager.New(
		createdData.HonestValidatorStateRoots,
		commonStateManagerOpts...,
	)
	require.NoError(t, err)

	commonStateManagerOpts = append(
		commonStateManagerOpts,
		statemanager.WithBigStepStateDivergenceHeight(1),
		statemanager.WithSmallStepStateDivergenceHeight(3),
	)
	evilStateManager, err := statemanager.New(
		createdData.EvilValidatorStateRoots,
		commonStateManagerOpts...,
	)
	require.NoError(t, err)

	challengeManager, err := createdData.Chains[0].SpecChallengeManager(ctx)
	require.NoError(t, err)

	// Honest assertion being added.
	leafAdder := func(startCommit, endCommit util.HistoryCommitment, leaf protocol.Assertion) protocol.SpecEdge {
		edge, err := challengeManager.AddBlockChallengeLevelZeroEdge(
			ctx,
			leaf,
			startCommit,
			endCommit,
		)
		require.NoError(t, err)
		return edge
	}
	honestStartCommit, err := honestStateManager.HistoryCommitmentUpTo(ctx, 0)
	require.NoError(t, err)
	honestEndCommit, err := honestStateManager.HistoryCommitmentUpTo(ctx, uint64(topLevelHeight))
	require.NoError(t, err)

	honestEdge := leafAdder(honestStartCommit, honestEndCommit, createdData.Leaf1)
	require.Equal(t, protocol.BlockChallengeEdge, honestEdge.GetType())
	hasRival, err := honestEdge.HasRival(ctx)
	require.NoError(t, err)
	require.Equal(t, true, !hasRival)

	t.Run("unrivaled level zero edge is not one step fork source", func(t *testing.T) {
		isOSF, err := honestEdge.HasLengthOneRival(ctx)
		require.NoError(t, err)
		require.Equal(t, false, isOSF)
	})

	evilStartCommit, err := evilStateManager.HistoryCommitmentUpTo(ctx, 0)
	require.NoError(t, err)
	evilEndCommit, err := evilStateManager.HistoryCommitmentUpTo(ctx, uint64(topLevelHeight))
	require.NoError(t, err)

	evilEdge := leafAdder(evilStartCommit, evilEndCommit, createdData.Leaf2)
	require.Equal(t, protocol.BlockChallengeEdge, evilEdge.GetType())

	// Honest and evil edge are rivals, neither is presumptive.
	hasRival, err = honestEdge.HasRival(ctx)
	require.NoError(t, err)
	require.Equal(t, false, !hasRival)

	hasRival, err = evilEdge.HasRival(ctx)
	require.NoError(t, err)
	require.Equal(t, false, !hasRival)

	return &bisectionScenario{
		topLevelFork:        createdData,
		honestStateManager:  honestStateManager,
		evilStateManager:    evilStateManager,
		honestLevelZeroEdge: honestEdge,
		evilLevelZeroEdge:   evilEdge,
		honestStartCommit:   honestStartCommit,
		evilStartCommit:     evilStartCommit,
	}
}

// Returns a snapshot of the data for a one-step-proof scenario in which
// an evil validator has reached a one step fork against an honest validator
// in a small step subchallenge. Their disagreement must then be resolved via
// a one-step-proof to declare a winner.
type oneStepProofScenario struct {
	topLevelFork        *setup.CreatedValidatorFork
	honestStateManager  statemanager.Manager
	evilStateManager    statemanager.Manager
	smallStepHonestEdge protocol.SpecEdge
	smallStepEvilEdge   protocol.SpecEdge
}

// Sets up a challenge between two validators in which they make challenge moves
// to reach a one-step-proof in a small step subchallenge. It returns the data needed
// to then confirm the winner by one-step-proof execution.
func setupOneStepProofScenario(
	t *testing.T,
	topLevelAssertionChainHeight protocol.Height,
	commonStateManagerOpts ...statemanager.Opt,
) *oneStepProofScenario {
	ctx := context.Background()
	bisectionScenario := setupBisectionScenario(t, topLevelAssertionChainHeight, commonStateManagerOpts...)
	honestStateManager := bisectionScenario.honestStateManager
	evilStateManager := bisectionScenario.evilStateManager
	honestEdge := bisectionScenario.honestLevelZeroEdge
	evilEdge := bisectionScenario.evilLevelZeroEdge
	honestStartCommit := bisectionScenario.honestStartCommit
	evilStartCommit := bisectionScenario.evilStartCommit

	challengeManager, err := bisectionScenario.topLevelFork.Chains[1].SpecChallengeManager(ctx)
	require.NoError(t, err)

	// Attempt bisections down to one step fork.
	honestBisectCommit, err := honestStateManager.HistoryCommitmentUpTo(ctx, 1)
	require.NoError(t, err)
	honestProof, err := honestStateManager.PrefixProof(ctx, 1, 2)
	require.NoError(t, err)

	_, _, err = honestEdge.Bisect(ctx, honestBisectCommit.Merkle, honestProof)
	require.NoError(t, err)

	evilBisectCommit, err := evilStateManager.HistoryCommitmentUpTo(ctx, 1)
	require.NoError(t, err)
	evilProof, err := evilStateManager.PrefixProof(ctx, 1, 2)
	require.NoError(t, err)

	oneStepForkSourceEdge, _, err := evilEdge.Bisect(ctx, evilBisectCommit.Merkle, evilProof)
	require.NoError(t, err)

	isAtOneStepFork, err := oneStepForkSourceEdge.HasLengthOneRival(ctx)
	require.NoError(t, err)
	require.Equal(t, true, isAtOneStepFork)

	// Now opening big step level zero leaves
	bigStepAdder := func(startCommit, endCommit util.HistoryCommitment) protocol.SpecEdge {
		leaf, err := challengeManager.AddSubChallengeLevelZeroEdge(
			ctx,
			oneStepForkSourceEdge,
			startCommit,
			endCommit,
		)
		require.NoError(t, err)
		return leaf
	}

	honestBigStepCommit, err := honestStateManager.BigStepCommitmentUpTo(
		ctx, 0 /* from assertion */, 1 /* to assertion */, 1, /* to big step */
	)
	require.NoError(t, err)

	honestEdge = bigStepAdder(honestStartCommit, honestBigStepCommit)
	require.Equal(t, protocol.BigStepChallengeEdge, honestEdge.GetType())
	hasRival, err := honestEdge.HasRival(ctx)
	require.NoError(t, err)
	require.Equal(t, true, !hasRival)

	evilBigStepCommit, err := evilStateManager.BigStepCommitmentUpTo(
		ctx, 0 /* from assertion */, 1 /* to assertion */, 1, /* to big step */
	)
	require.NoError(t, err)

	evilEdge = bigStepAdder(evilStartCommit, evilBigStepCommit)
	require.Equal(t, protocol.BigStepChallengeEdge, evilEdge.GetType())

	hasRival, err = honestEdge.HasRival(ctx)
	require.NoError(t, err)
	require.Equal(t, false, !hasRival)
	hasRival, err = evilEdge.HasRival(ctx)
	require.NoError(t, err)
	require.Equal(t, false, !hasRival)

	isAtOneStepFork, err = honestEdge.HasLengthOneRival(ctx)
	require.NoError(t, err)
	require.Equal(t, true, isAtOneStepFork)

	// Now opening small step level zero leaves
	smallStepAdder := func(startCommit, endCommit util.HistoryCommitment) protocol.SpecEdge {
		leaf, err := challengeManager.AddSubChallengeLevelZeroEdge(
			ctx,
			honestEdge,
			startCommit,
			endCommit,
		)
		require.NoError(t, err)
		return leaf
	}

	honestSmallStepCommit, err := honestStateManager.SmallStepCommitmentUpTo(
		ctx, 0 /* from assertion */, 1 /* to assertion */, 3, /* to pc */
	)
	require.NoError(t, err)

	smallStepHonest := smallStepAdder(honestStartCommit, honestSmallStepCommit)
	require.Equal(t, protocol.SmallStepChallengeEdge, smallStepHonest.GetType())
	hasRival, err = smallStepHonest.HasRival(ctx)
	require.NoError(t, err)
	require.Equal(t, true, !hasRival)

	evilSmallStepCommit, err := evilStateManager.SmallStepCommitmentUpTo(
		ctx, 0 /* from assertion */, 1 /* to assertion */, 3, /* to pc */
	)
	require.NoError(t, err)

	smallStepEvil := smallStepAdder(evilStartCommit, evilSmallStepCommit)
	require.Equal(t, protocol.SmallStepChallengeEdge, smallStepEvil.GetType())

	hasRival, err = smallStepHonest.HasRival(ctx)
	require.NoError(t, err)
	require.Equal(t, false, !hasRival)
	hasRival, err = smallStepEvil.HasRival(ctx)
	require.NoError(t, err)
	require.Equal(t, false, !hasRival)

	// Get the lower-level edge of either vertex we just bisected.
	require.Equal(t, protocol.SmallStepChallengeEdge, smallStepHonest.GetType())

	t.Log("GETS HERE")

	// Attempt bisections down to one step fork.
	honestBisectCommit, err = honestStateManager.SmallStepCommitmentUpTo(ctx, 0, 1, 2)
	require.NoError(t, err)
	honestProof, err = honestStateManager.PrefixProof(ctx, 2, 3)
	require.NoError(t, err)

	_, _, err = smallStepHonest.Bisect(ctx, honestBisectCommit.Merkle, honestProof)
	require.NoError(t, err)

	evilBisectCommit, err = evilStateManager.SmallStepCommitmentUpTo(ctx, 0, 1, 2)
	require.NoError(t, err)
	evilProof, err = evilStateManager.PrefixProof(ctx, 2, 3)
	require.NoError(t, err)

	oneStepForkSourceEdge, _, err = smallStepEvil.Bisect(ctx, evilBisectCommit.Merkle, evilProof)
	require.NoError(t, err)

	isAtOneStepFork, err = oneStepForkSourceEdge.HasLengthOneRival(ctx)
	require.NoError(t, err)
	require.Equal(t, true, isAtOneStepFork)
	return &oneStepProofScenario{
		topLevelFork:        bisectionScenario.topLevelFork,
		honestStateManager:  honestStateManager,
		evilStateManager:    evilStateManager,
		smallStepHonestEdge: smallStepHonest,
		smallStepEvilEdge:   smallStepEvil,
	}
}
