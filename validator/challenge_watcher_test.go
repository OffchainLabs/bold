package validator

import (
	"context"
	"testing"
	"time"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/state-manager"
	"github.com/OffchainLabs/challenge-protocol-v2/testing/setup"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/stretchr/testify/require"
)

func TestFullConfirmations(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	scenario := setupOneStepProofScenario(t)

	chain := scenario.topLevelFork.Chains[0]
	challengeManager, err := scenario.topLevelFork.Chains[1].SpecChallengeManager(ctx)
	require.NoError(t, err)

	honestStateManager := scenario.honestStateManager
	fromBlockChallengeHeight := uint64(0)
	toBlockChallengeHeight := uint64(1)
	fromBigStep := uint64(0)
	toBigStep := uint64(1)

	prevId, err := scenario.smallStepHonestLowerChild.PrevAssertionId(ctx)
	require.NoError(t, err)
	assertionNum, err := chain.GetAssertionNum(ctx, prevId)
	require.NoError(t, err)
	prevAssertion, err := chain.AssertionBySequenceNum(ctx, assertionNum)
	require.NoError(t, err)
	parentAssertionStateHash, err := prevAssertion.StateHash()
	require.NoError(t, err)
	assertionCreationInfo, err := chain.ReadAssertionCreationInfo(ctx, assertionNum)
	require.NoError(t, err)

	ospConfirm := func(edge protocol.SpecEdge, fromSmallStep, toSmallStep uint64) {
		data, startInclusionProof, endInclusionProof, err := honestStateManager.OneStepProofData(
			ctx,
			parentAssertionStateHash,
			assertionCreationInfo,
			fromBlockChallengeHeight,
			toBlockChallengeHeight,
			fromBigStep,
			toBigStep,
			fromSmallStep,
			toSmallStep,
		)
		require.NoError(t, err)

		err = challengeManager.ConfirmEdgeByOneStepProof(
			ctx,
			edge.Id(),
			data,
			startInclusionProof,
			endInclusionProof,
		)
		require.NoError(t, err)
		edgeStatus, err := edge.Status(ctx)
		require.NoError(t, err)
		require.Equal(t, protocol.EdgeConfirmed, edgeStatus)
	}

	fromSmallStep := uint64(0)
	toSmallStep := uint64(1)
	ospConfirm(scenario.smallStepHonestLowerChild, fromSmallStep, toSmallStep)

	t.Log("Small step lower honest child one step proven")

	fromSmallStep = uint64(1)
	toSmallStep = uint64(2)
	ospConfirm(scenario.smallStepHonestUpperChild, fromSmallStep, toSmallStep)

	t.Log("Small step upper honest child one step proven")

	createdData := scenario.topLevelFork
	watcher := NewWatcher(createdData.Chains[0], createdData.Backend, time.Millisecond*100)
	go watcher.Watch(ctx)
	time.Sleep(time.Second)
	cancel()
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
	commonStateManagerOpts ...statemanager.Opt,
) *bisectionScenario {
	ctx := context.Background()

	createdData, err := setup.CreateTwoValidatorFork(ctx, &setup.CreateForkConfig{
		NumBlocks:     8,
		DivergeHeight: 0,
	})
	require.NoError(t, err)

	honestStateManager, err := statemanager.NewWithAssertionStates(
		createdData.HonestValidatorStates,
		createdData.HonestValidatorInboxCounts,
		commonStateManagerOpts...,
	)
	require.NoError(t, err)

	commonStateManagerOpts = append(
		commonStateManagerOpts,
		statemanager.WithMaliciousIntent(),
		statemanager.WithBigStepStateDivergenceHeight(1),
		statemanager.WithSmallStepStateDivergenceHeight(1),
	)
	evilStateManager, err := statemanager.NewWithAssertionStates(
		createdData.EvilValidatorStates,
		createdData.EvilValidatorInboxCounts,
		commonStateManagerOpts...,
	)
	require.NoError(t, err)

	challengeManager, err := createdData.Chains[0].SpecChallengeManager(ctx)
	require.NoError(t, err)

	// Honest assertion being added.
	leafAdder := func(stateManager statemanager.Manager, leaf protocol.Assertion) (util.HistoryCommitment, protocol.SpecEdge) {
		startCommit, err := stateManager.HistoryCommitmentUpToBatch(ctx, 0, 0, 1)
		require.NoError(t, err)
		endCommit, err := stateManager.HistoryCommitmentUpToBatch(ctx, 0, protocol.LevelZeroBlockEdgeHeight, 1)
		require.NoError(t, err)
		prefixProof, err := stateManager.PrefixProofUpToBatch(ctx, 0, 0, protocol.LevelZeroBlockEdgeHeight, 1)
		require.NoError(t, err)

		edge, err := challengeManager.AddBlockChallengeLevelZeroEdge(
			ctx,
			leaf,
			startCommit,
			endCommit,
			prefixProof,
		)
		require.NoError(t, err)
		return startCommit, edge
	}

	honestStartCommit, honestEdge := leafAdder(honestStateManager, createdData.Leaf1)
	require.Equal(t, protocol.BlockChallengeEdge, honestEdge.GetType())
	hasRival, err := honestEdge.HasRival(ctx)
	require.NoError(t, err)
	require.Equal(t, true, !hasRival)

	isOSF, err := honestEdge.HasLengthOneRival(ctx)
	require.NoError(t, err)
	require.Equal(t, false, isOSF)

	evilStartCommit, evilEdge := leafAdder(evilStateManager, createdData.Leaf2)
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
	topLevelFork              *setup.CreatedValidatorFork
	honestStateManager        statemanager.Manager
	evilStateManager          statemanager.Manager
	smallStepHonestLowerChild protocol.SpecEdge
	smallStepHonestUpperChild protocol.SpecEdge
	smallStepEvilEdge         protocol.SpecEdge
}

// Sets up a challenge between two validators in which they make challenge moves
// to reach a one-step-proof in a small step subchallenge. It returns the data needed
// to then confirm the winner by one-step-proof execution.
func setupOneStepProofScenario(
	t *testing.T,
	commonStateManagerOpts ...statemanager.Opt,
) *oneStepProofScenario {
	ctx := context.Background()
	commonStateManagerOpts = append(
		commonStateManagerOpts,
		statemanager.WithNumOpcodesPerBigStep(protocol.LevelZeroSmallStepEdgeHeight),
		statemanager.WithMaxWavmOpcodesPerBlock(protocol.LevelZeroBigStepEdgeHeight*protocol.LevelZeroSmallStepEdgeHeight),
	)
	bisectionScenario := setupBisectionScenario(t, commonStateManagerOpts...)
	honestStateManager := bisectionScenario.honestStateManager
	evilStateManager := bisectionScenario.evilStateManager
	honestEdge := bisectionScenario.honestLevelZeroEdge
	evilEdge := bisectionScenario.evilLevelZeroEdge

	challengeManager, err := bisectionScenario.topLevelFork.Chains[1].SpecChallengeManager(ctx)
	require.NoError(t, err)

	var blockHeight uint64 = protocol.LevelZeroBlockEdgeHeight
	for blockHeight > 1 {
		honestBisectCommit, err := honestStateManager.HistoryCommitmentUpToBatch(ctx, 0, blockHeight/2, 1)
		require.NoError(t, err)
		honestProof, err := honestStateManager.PrefixProofUpToBatch(ctx, 0, blockHeight/2, blockHeight, 1)
		require.NoError(t, err)
		honestEdge, _, err = honestEdge.Bisect(ctx, honestBisectCommit.Merkle, honestProof)
		require.NoError(t, err)

		evilBisectCommit, err := evilStateManager.HistoryCommitmentUpToBatch(ctx, 0, blockHeight/2, 1)
		require.NoError(t, err)
		evilProof, err := evilStateManager.PrefixProofUpToBatch(ctx, 0, blockHeight/2, blockHeight, 1)
		require.NoError(t, err)
		evilEdge, _, err = evilEdge.Bisect(ctx, evilBisectCommit.Merkle, evilProof)
		require.NoError(t, err)

		blockHeight /= 2

		isOSF, err := honestEdge.HasLengthOneRival(ctx)
		require.NoError(t, err)
		require.Equal(t, blockHeight == 1, isOSF)
		isOSF, err = evilEdge.HasLengthOneRival(ctx)
		require.NoError(t, err)
		require.Equal(t, blockHeight == 1, isOSF)
	}

	// Now opening big step level zero leaves at index 0
	bigStepAdder := func(stateManager statemanager.Manager, sourceEdge protocol.SpecEdge) protocol.SpecEdge {
		startCommit, err := stateManager.BigStepCommitmentUpTo(ctx, 0, 1, 0)
		require.NoError(t, err)
		endCommit, err := stateManager.BigStepLeafCommitment(ctx, 0, 1)
		require.NoError(t, err)
		require.Equal(t, startCommit.LastLeaf, endCommit.FirstLeaf)
		startParentCommitment, err := stateManager.HistoryCommitmentUpToBatch(ctx, 0, 0, 1)
		require.NoError(t, err)
		endParentCommitment, err := stateManager.HistoryCommitmentUpToBatch(ctx, 0, 1, 1)
		require.NoError(t, err)
		startEndPrefixProof, err := stateManager.BigStepPrefixProof(ctx, 0, 1, 0, endCommit.Height)
		require.NoError(t, err)
		leaf, err := challengeManager.AddSubChallengeLevelZeroEdge(
			ctx,
			sourceEdge,
			startCommit,
			endCommit,
			startParentCommitment.LastLeafProof,
			endParentCommitment.LastLeafProof,
			startEndPrefixProof,
		)
		require.NoError(t, err)
		return leaf
	}

	honestEdge = bigStepAdder(honestStateManager, honestEdge)
	require.Equal(t, protocol.BigStepChallengeEdge, honestEdge.GetType())
	hasRival, err := honestEdge.HasRival(ctx)
	require.NoError(t, err)
	require.Equal(t, true, !hasRival)

	evilEdge = bigStepAdder(evilStateManager, evilEdge)
	require.Equal(t, protocol.BigStepChallengeEdge, evilEdge.GetType())

	var bigStepHeight uint64 = protocol.LevelZeroBigStepEdgeHeight
	for bigStepHeight > 1 {
		honestBisectCommit, err := honestStateManager.BigStepCommitmentUpTo(ctx, 0, 1, bigStepHeight/2)
		require.NoError(t, err)
		honestProof, err := honestStateManager.BigStepPrefixProof(ctx, 0, 1, bigStepHeight/2, bigStepHeight)
		require.NoError(t, err)
		honestEdge, _, err = honestEdge.Bisect(ctx, honestBisectCommit.Merkle, honestProof)
		require.NoError(t, err)

		evilBisectCommit, err := evilStateManager.BigStepCommitmentUpTo(ctx, 0, 1, bigStepHeight/2)
		require.NoError(t, err)
		evilProof, err := evilStateManager.BigStepPrefixProof(ctx, 0, 1, bigStepHeight/2, bigStepHeight)
		require.NoError(t, err)
		evilEdge, _, err = evilEdge.Bisect(ctx, evilBisectCommit.Merkle, evilProof)
		require.NoError(t, err)

		bigStepHeight /= 2

		isOSF, err := honestEdge.HasLengthOneRival(ctx)
		require.NoError(t, err)
		require.Equal(t, bigStepHeight == 1, isOSF)
		isOSF, err = evilEdge.HasLengthOneRival(ctx)
		require.NoError(t, err)
		require.Equal(t, bigStepHeight == 1, isOSF)
	}

	hasRival, err = honestEdge.HasRival(ctx)
	require.NoError(t, err)
	require.Equal(t, false, !hasRival)
	hasRival, err = evilEdge.HasRival(ctx)
	require.NoError(t, err)
	require.Equal(t, false, !hasRival)

	isAtOneStepFork, err := honestEdge.HasLengthOneRival(ctx)
	require.NoError(t, err)
	require.Equal(t, true, isAtOneStepFork)

	// Now opening small step level zero leaves at index 0
	smallStepAdder := func(stateManager statemanager.Manager, edge protocol.SpecEdge) protocol.SpecEdge {
		startCommit, err := stateManager.SmallStepCommitmentUpTo(ctx, 0, 1, 0, 1, 0)
		require.NoError(t, err)
		endCommit, err := stateManager.SmallStepLeafCommitment(ctx, 0, 1, 0, 1)
		require.NoError(t, err)
		startParentCommitment, err := stateManager.BigStepCommitmentUpTo(ctx, 0, 1, 0)
		require.NoError(t, err)
		endParentCommitment, err := stateManager.BigStepCommitmentUpTo(ctx, 0, 1, 1)
		require.NoError(t, err)
		startEndPrefixProof, err := stateManager.SmallStepPrefixProof(ctx, 0, 1, 0, 1, 0, endCommit.Height)
		require.NoError(t, err)
		leaf, err := challengeManager.AddSubChallengeLevelZeroEdge(
			ctx,
			edge,
			startCommit,
			endCommit,
			startParentCommitment.LastLeafProof,
			endParentCommitment.LastLeafProof,
			startEndPrefixProof,
		)
		require.NoError(t, err)
		return leaf
	}

	honestEdge = smallStepAdder(honestStateManager, honestEdge)
	require.Equal(t, protocol.SmallStepChallengeEdge, honestEdge.GetType())
	hasRival, err = honestEdge.HasRival(ctx)
	require.NoError(t, err)
	require.Equal(t, true, !hasRival)

	evilEdge = smallStepAdder(evilStateManager, evilEdge)
	require.Equal(t, protocol.SmallStepChallengeEdge, evilEdge.GetType())

	hasRival, err = honestEdge.HasRival(ctx)
	require.NoError(t, err)
	require.Equal(t, false, !hasRival)
	hasRival, err = evilEdge.HasRival(ctx)
	require.NoError(t, err)
	require.Equal(t, false, !hasRival)

	// Get the lower-level edge of either vertex we just bisected.
	require.Equal(t, protocol.SmallStepChallengeEdge, honestEdge.GetType())

	honestLowerChild := honestEdge
	honestUpperChild := honestEdge
	evilLowerChild := evilEdge

	var smallStepHeight uint64 = protocol.LevelZeroBigStepEdgeHeight
	for smallStepHeight > 1 {
		t.Log("Running")
		honestBisectCommit, err := honestStateManager.SmallStepCommitmentUpTo(ctx, 0, 1, 0, 1, smallStepHeight/2)
		require.NoError(t, err)
		honestProof, err := honestStateManager.SmallStepPrefixProof(ctx, 0, 1, 0, 1, smallStepHeight/2, smallStepHeight)
		require.NoError(t, err)
		honestLowerChild, honestUpperChild, err = honestLowerChild.Bisect(ctx, honestBisectCommit.Merkle, honestProof)
		require.NoError(t, err)

		evilBisectCommit, err := evilStateManager.SmallStepCommitmentUpTo(ctx, 0, 1, 0, 1, smallStepHeight/2)
		require.NoError(t, err)
		evilProof, err := evilStateManager.SmallStepPrefixProof(ctx, 0, 1, 0, 1, smallStepHeight/2, smallStepHeight)
		require.NoError(t, err)
		evilLowerChild, _, err = evilLowerChild.Bisect(ctx, evilBisectCommit.Merkle, evilProof)
		require.NoError(t, err)

		smallStepHeight /= 2

		isOSF, err := honestLowerChild.HasLengthOneRival(ctx)
		require.NoError(t, err)
		require.Equal(t, smallStepHeight == 1, isOSF)
		isOSF, err = evilLowerChild.HasLengthOneRival(ctx)
		require.NoError(t, err)
		require.Equal(t, smallStepHeight == 1, isOSF)
	}

	isAtOneStepFork, err = honestLowerChild.HasLengthOneRival(ctx)
	require.NoError(t, err)
	require.Equal(t, true, isAtOneStepFork)
	isAtOneStepFork, err = evilLowerChild.HasLengthOneRival(ctx)
	require.NoError(t, err)
	require.Equal(t, true, isAtOneStepFork)

	return &oneStepProofScenario{
		topLevelFork:              bisectionScenario.topLevelFork,
		honestStateManager:        honestStateManager,
		evilStateManager:          evilStateManager,
		smallStepHonestLowerChild: honestLowerChild,
		smallStepHonestUpperChild: honestUpperChild,
	}
}
