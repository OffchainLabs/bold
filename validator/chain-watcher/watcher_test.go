package watcher

import (
	"context"
	"testing"

	"math/big"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/solgen/go/challengeV2gen"
	"github.com/OffchainLabs/challenge-protocol-v2/testing/mocks"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/OffchainLabs/challenge-protocol-v2/util/threadsafe"
	challengetree "github.com/OffchainLabs/challenge-protocol-v2/validator/challenge-tree"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestWatcher_processEdgeConfirmation(t *testing.T) {
	ctx := context.Background()
	mockChain := &mocks.MockProtocol{}
	mockChallengeManager := &mocks.MockSpecChallengeManager{}
	mockChain.On(
		"SpecChallengeManager",
		ctx,
	).Return(mockChallengeManager, nil)

	assertionId := protocol.AssertionId(common.BytesToHash([]byte("foo")))
	edgeId := protocol.EdgeId(common.BytesToHash([]byte("bar")))
	edge := &mocks.MockSpecEdge{}

	mockChallengeManager.On(
		"GetEdge", ctx, edgeId,
	).Return(util.Some(protocol.SpecEdge(edge)), nil)

	edge.On("ClaimId").Return(util.Some(protocol.ClaimId(assertionId)))
	edge.On(
		"PrevAssertionId",
		ctx,
	).Return(assertionId, nil)

	watcher := &Watcher{
		challenges: threadsafe.NewMap[protocol.AssertionId, *trackedChallenge](),
		chain:      mockChain,
	}
	watcher.challenges.Put(assertionId, &trackedChallenge{
		confirmedLevelZeroEdgeClaimIds: threadsafe.NewSet[protocol.ClaimId](),
	})

	err := watcher.processEdgeConfirmation(ctx, edgeId)
	require.NoError(t, err)

	chal, ok := watcher.challenges.TryGet(assertionId)
	require.Equal(t, true, ok)
	ok = chal.confirmedLevelZeroEdgeClaimIds.Has(protocol.ClaimId(assertionId))
	require.Equal(t, true, ok)
}

func TestWatcher_processEdgeAddedEvent(t *testing.T) {
	ctx := context.Background()
	mockChain := &mocks.MockProtocol{}
	mockChallengeManager := &mocks.MockSpecChallengeManager{}
	mockChain.On(
		"SpecChallengeManager",
		ctx,
	).Return(mockChallengeManager, nil)

	assertionId := protocol.AssertionId(common.BytesToHash([]byte("foo")))
	edgeId := protocol.EdgeId(common.BytesToHash([]byte("bar")))
	edge := &mocks.MockSpecEdge{}

	mockChain.On(
		"TopLevelAssertion",
		ctx,
		edgeId,
	).Return(assertionId, nil)
	mockChain.On(
		"GetAssertionNum",
		ctx,
		assertionId,
	).Return(protocol.AssertionSequenceNumber(0), nil)

	info := &protocol.AssertionCreatedInfo{
		InboxMaxCount: big.NewInt(1),
	}
	mockChain.On(
		"ReadAssertionCreationInfo",
		ctx,
		protocol.AssertionSequenceNumber(0),
	).Return(info, nil)
	heights := &protocol.OriginHeights{}
	mockChain.On(
		"TopLevelClaimHeights",
		ctx,
		edgeId,
	).Return(heights, nil)
	mockChain.On(
		"AssertionUnrivaledTime",
		ctx,
		assertionId,
	).Return(uint64(0), nil)

	mockChallengeManager.On(
		"GetEdge", ctx, edgeId,
	).Return(util.Some(protocol.SpecEdge(edge)), nil)

	edge.On("Id").Return(edgeId)
	edge.On("CreatedAtBlock").Return(uint64(0))
	edge.On("ClaimId").Return(util.Some(protocol.ClaimId(assertionId)))
	edge.On("MutualId").Return(protocol.MutualId{})
	edge.On("GetType").Return(protocol.BlockChallengeEdge)
	startCommit := common.BytesToHash([]byte("nyan"))
	endCommit := common.BytesToHash([]byte("nyan2"))
	edge.On("StartCommitment").Return(protocol.Height(0), startCommit)
	edge.On("EndCommitment").Return(protocol.Height(4), endCommit)
	edge.On(
		"PrevAssertionId",
		ctx,
	).Return(assertionId, nil)

	mockStateManager := &mocks.MockStateManager{
		Agreement: protocol.Agreement{
			IsHonestEdge:          true,
			AgreesWithStartCommit: true,
		},
	}

	watcher := &Watcher{
		challenges:   threadsafe.NewMap[protocol.AssertionId, *trackedChallenge](),
		stateManager: mockStateManager,
		chain:        mockChain,
	}
	err := watcher.processEdgeAddedEvent(ctx, &challengeV2gen.EdgeChallengeManagerEdgeAdded{
		EdgeId:   edgeId,
		OriginId: assertionId,
	})
	require.NoError(t, err)

	chal, ok := watcher.challenges.TryGet(assertionId)
	require.Equal(t, true, ok)

	// Expect it to exist and be unrivaled for 10 blocks if we query at block number = 10.
	blockNumber := uint64(10)
	pathTimer, _, err := chal.honestEdgeTree.HonestPathTimer(ctx, edgeId, blockNumber)
	require.NoError(t, err)
	require.Equal(t, pathTimer, challengetree.PathTimer(blockNumber))
}
