// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/challenge-protocol-v2/blob/main/LICENSE

package watcher

import (
	"context"
	"math/big"
	"testing"

	protocol "github.com/OffchainLabs/challenge-protocol-v2/chain-abstraction"
	challengetree "github.com/OffchainLabs/challenge-protocol-v2/challenge-manager/challenge-tree"
	"github.com/OffchainLabs/challenge-protocol-v2/containers/option"
	"github.com/OffchainLabs/challenge-protocol-v2/containers/threadsafe"
	l2stateprovider "github.com/OffchainLabs/challenge-protocol-v2/layer2-state-provider"
	"github.com/OffchainLabs/challenge-protocol-v2/solgen/go/challengeV2gen"
	"github.com/OffchainLabs/challenge-protocol-v2/testing/mocks"
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

	assertionHash := protocol.AssertionHash{Hash: common.BytesToHash([]byte("foo"))}
	edgeId := protocol.EdgeId(common.BytesToHash([]byte("bar")))
	edge := &mocks.MockSpecEdge{}

	mockChallengeManager.On(
		"GetEdge", ctx, edgeId,
	).Return(option.Some(protocol.SpecEdge(edge)), nil)

	edge.On("ClaimId").Return(option.Some(protocol.ClaimId(assertionHash.Hash)))
	edge.On("Id").Return(edgeId)
	edge.On("GetType").Return(protocol.BigStepChallengeEdge)
	edge.On(
		"AssertionHash",
		ctx,
	).Return(assertionHash, nil)

	watcher := &Watcher{
		challenges: threadsafe.NewMap[protocol.AssertionHash, *trackedChallenge](),
		chain:      mockChain,
	}
	watcher.challenges.Put(assertionHash, &trackedChallenge{
		confirmedLevelZeroEdgeClaimIds: threadsafe.NewMap[protocol.ClaimId, protocol.EdgeId](),
	})

	err := watcher.processEdgeConfirmation(ctx, edgeId)
	require.NoError(t, err)

	chal, ok := watcher.challenges.TryGet(assertionHash)
	require.Equal(t, true, ok)
	ok = chal.confirmedLevelZeroEdgeClaimIds.Has(protocol.ClaimId(assertionHash.Hash))
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

	assertionHash := protocol.AssertionHash{Hash: common.BytesToHash([]byte("foo"))}
	edgeId := protocol.EdgeId(common.BytesToHash([]byte("bar")))
	edge := &mocks.MockSpecEdge{}

	mockChain.On(
		"TopLevelAssertion",
		ctx,
		edgeId,
	).Return(assertionHash, nil)

	info := &protocol.AssertionCreatedInfo{
		InboxMaxCount: big.NewInt(1),
	}
	mockChain.On(
		"ReadAssertionCreationInfo",
		ctx,
		assertionHash,
	).Return(info, nil)
	heights := protocol.OriginHeights{}
	mockChain.On(
		"TopLevelClaimHeights",
		ctx,
		edgeId,
	).Return(heights, nil)

	assertionUnrivaledBlocks := uint64(5)
	mockChain.On(
		"AssertionUnrivaledBlocks",
		ctx,
		assertionHash,
	).Return(assertionUnrivaledBlocks, nil)

	mockChallengeManager.On(
		"GetEdge", ctx, edgeId,
	).Return(option.Some(protocol.SpecEdge(edge)), nil)

	edge.On("Id").Return(edgeId)
	edge.On("CreatedAtBlock").Return(uint64(0), nil)
	edge.On("ClaimId").Return(option.Some(protocol.ClaimId(assertionHash.Hash)))
	edge.On("MutualId").Return(protocol.MutualId{})
	edge.On("GetType").Return(protocol.BlockChallengeEdge)
	startCommit := common.BytesToHash([]byte("nyan"))
	endCommit := common.BytesToHash([]byte("nyan2"))
	edge.On("StartCommitment").Return(protocol.Height(0), startCommit)
	edge.On("EndCommitment").Return(protocol.Height(4), endCommit)
	edge.On(
		"AssertionHash",
		ctx,
	).Return(assertionHash, nil)

	mockStateManager := &mocks.MockStateManager{}
	mockStateManager.On(
		"AgreesWithHistoryCommitment",
		ctx,
		common.Hash{},
		uint64(1),
		protocol.BlockChallengeEdge,
		protocol.OriginHeights{
			BlockChallengeOriginHeight: 0,
		},
		l2stateprovider.History{
			Height:     uint64(0),
			MerkleRoot: startCommit,
		},
	).Return(true, nil)
	mockStateManager.On(
		"AgreesWithHistoryCommitment",
		ctx,
		common.Hash{},
		uint64(1),
		protocol.BlockChallengeEdge,
		protocol.OriginHeights{
			BlockChallengeOriginHeight: 0,
		},
		l2stateprovider.History{
			Height:     uint64(4),
			MerkleRoot: endCommit,
		},
	).Return(true, nil)

	mockManager := &mocks.MockEdgeTracker{}
	mockManager.On("TrackEdge", ctx, edge).Return(nil)

	watcher := &Watcher{
		challenges:  threadsafe.NewMap[protocol.AssertionHash, *trackedChallenge](),
		histChecker: mockStateManager,
		chain:       mockChain,
		edgeManager: mockManager,
	}
	err := watcher.processEdgeAddedEvent(ctx, &challengeV2gen.EdgeChallengeManagerEdgeAdded{
		EdgeId:   edgeId,
		OriginId: assertionHash.Hash,
	})
	require.NoError(t, err)

	chal, ok := watcher.challenges.TryGet(assertionHash)
	require.Equal(t, true, ok)

	// Expect it to exist and be unrivaled for 10 blocks if we query at block number = 10,
	// plus the number of blocks the top level assertion was unrivaled (5).
	blockNumber := uint64(10)
	pathTimer, _, err := chal.honestEdgeTree.HonestPathTimer(ctx, edgeId, blockNumber)
	require.NoError(t, err)
	require.Equal(t, pathTimer, challengetree.PathTimer(blockNumber+assertionUnrivaledBlocks))
}
