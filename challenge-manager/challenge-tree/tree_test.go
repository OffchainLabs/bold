// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/challenge-protocol-v2/blob/main/LICENSE

package challengetree

import (
	"context"
	"errors"
	"math/big"
	"strconv"
	"strings"
	"testing"

	protocol "github.com/OffchainLabs/challenge-protocol-v2/chain-abstraction"
	"github.com/OffchainLabs/challenge-protocol-v2/challenge-manager/challenge-tree/mock"
	"github.com/OffchainLabs/challenge-protocol-v2/containers/threadsafe"
	l2stateprovider "github.com/OffchainLabs/challenge-protocol-v2/layer2-state-provider"
	"github.com/OffchainLabs/challenge-protocol-v2/testing/mocks"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestAddEdge(t *testing.T) {
	ht := &HonestChallengeTree{
		edges:                         threadsafe.NewMap[protocol.EdgeId, protocol.SpecEdge](),
		mutualIds:                     threadsafe.NewMap[protocol.MutualId, *threadsafe.Map[protocol.EdgeId, creationTime]](),
		honestBigStepLevelZeroEdges:   threadsafe.NewSlice[protocol.ReadOnlyEdge](),
		honestSmallStepLevelZeroEdges: threadsafe.NewSlice[protocol.ReadOnlyEdge](),
	}
	ht.topLevelAssertionHash = protocol.AssertionHash{Hash: common.BytesToHash([]byte("foo"))}
	ctx := context.Background()
	edge := newEdge(&newCfg{t: t, edgeId: "blk-0.a-16.a", createdAt: 1})

	t.Run("getting top level assertion fails", func(t *testing.T) {
		ht.metadataReader = &mockMetadataReader{
			assertionErr: errors.New("bad request"),
		}
		_, err := ht.AddEdge(ctx, edge)
		require.ErrorContains(t, err, "could not get top level assertion for edge")
	})
	t.Run("ignores if disagrees with top level assertion hash of edge", func(t *testing.T) {
		ht.metadataReader = &mockMetadataReader{
			assertionErr:  nil,
			assertionHash: protocol.AssertionHash{Hash: common.BytesToHash([]byte("bar"))},
		}
		_, err := ht.AddEdge(ctx, edge)
		require.NoError(t, err)
	})
	t.Run("getting claim heights fails", func(t *testing.T) {
		ht.metadataReader = &mockMetadataReader{
			assertionErr:    nil,
			assertionHash:   ht.topLevelAssertionHash,
			claimHeightsErr: errors.New("bad request"),
		}
		_, err := ht.AddEdge(ctx, edge)
		require.ErrorContains(t, err, "could not get claim heights for edge")
	})
	t.Run("checking if agrees with commit fails", func(t *testing.T) {
		ht.metadataReader = &mockMetadataReader{
			assertionErr:  nil,
			assertionHash: ht.topLevelAssertionHash,
		}
		start, startCommit := edge.StartCommitment()
		end, endCommit := edge.EndCommitment()
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
				Height:     uint64(start),
				MerkleRoot: startCommit,
			},
		).Return(false, errors.New("something went wrong"))
		mockStateManager.On(
			"AgreesWithHistoryCommitment",
			ctx,
			uint64(1),
			protocol.BlockChallengeEdge,
			protocol.OriginHeights{
				BlockChallengeOriginHeight: 0,
			},
			l2stateprovider.History{
				Height:     uint64(end),
				MerkleRoot: endCommit,
			},
		).Return(false, errors.New("something went wrong"))
		ht.histChecker = mockStateManager
		_, err := ht.AddEdge(ctx, edge)
		require.ErrorContains(t, err, "could not check if agrees with")
	})
	t.Run("fully disagrees with edge", func(t *testing.T) {
		ht.metadataReader = &mockMetadataReader{
			assertionErr:  nil,
			assertionHash: ht.topLevelAssertionHash,
		}
		badEdge := newEdge(&newCfg{t: t, edgeId: "blk-0.f-16.a", createdAt: 1})
		startHeight, startCommit := badEdge.StartCommitment()
		endHeight, endCommit := badEdge.EndCommitment()
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
				Height:     uint64(startHeight),
				MerkleRoot: startCommit,
			},
		).Return(false, nil)
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
				Height:     uint64(endHeight),
				MerkleRoot: endCommit,
			},
		).Return(false, nil)
		ht.histChecker = mockStateManager
		agreement, err := ht.AddEdge(ctx, badEdge)
		require.NoError(t, err)
		require.Equal(t, protocol.Agreement{
			IsHonestEdge:          false,
			AgreesWithStartCommit: false,
		}, agreement)

		// Check the edge is not kept track of anywhere.
		_, ok := ht.edges.TryGet(badEdge.Id())
		require.Equal(t, false, ok)
		_, ok = ht.mutualIds.TryGet(badEdge.MutualId())
		require.Equal(t, false, ok)
	})
	t.Run("agrees with edge but is not a level zero edge", func(t *testing.T) {
		ht.metadataReader = &mockMetadataReader{
			assertionErr:  nil,
			assertionHash: ht.topLevelAssertionHash,
		}
		edge := newEdge(&newCfg{t: t, edgeId: "blk-0.a-16.a", createdAt: 1})
		startHeight, startCommit := edge.StartCommitment()
		endHeight, endCommit := edge.EndCommitment()
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
				Height:     uint64(startHeight),
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
				Height:     uint64(endHeight),
				MerkleRoot: endCommit,
			},
		).Return(true, nil)
		ht.histChecker = mockStateManager
		agreement, err := ht.AddEdge(ctx, edge)
		require.NoError(t, err)
		require.Equal(t, protocol.Agreement{
			IsHonestEdge:          true,
			AgreesWithStartCommit: true,
		}, agreement)

		// Exists.
		_, ok := ht.edges.TryGet(edge.Id())
		require.Equal(t, true, ok)
		// Exists in the mutual ids mapping.
		_, ok = ht.mutualIds.TryGet(edge.MutualId())
		require.Equal(t, true, ok)

		// However, we should not have a level zero edge being tracked yet.
		require.Equal(t, true, ht.honestBlockChalLevelZeroEdge.IsNone())
		require.Equal(t, true, ht.honestBigStepLevelZeroEdges.Len() == 0)
		require.Equal(t, true, ht.honestSmallStepLevelZeroEdges.Len() == 0)
	})
	t.Run("agrees with edge and is a level zero edge", func(t *testing.T) {
		ht.metadataReader = &mockMetadataReader{
			assertionErr:  nil,
			assertionHash: ht.topLevelAssertionHash,
		}
		edge := newEdge(&newCfg{t: t, edgeId: "blk-0.a-32.a", createdAt: 1, claimId: "foo"})
		startHeight, startCommit := edge.StartCommitment()
		endHeight, endCommit := edge.EndCommitment()
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
				Height:     uint64(startHeight),
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
				Height:     uint64(endHeight),
				MerkleRoot: endCommit,
			},
		).Return(true, nil)
		ht.histChecker = mockStateManager
		_, err := ht.AddEdge(ctx, edge)
		require.NoError(t, err)

		// Exists.
		_, ok := ht.edges.TryGet(edge.Id())
		require.Equal(t, true, ok)
		// Exists in the mutual ids mapping.
		_, ok = ht.mutualIds.TryGet(edge.MutualId())
		require.Equal(t, true, ok)

		// We should have a level zero edge being tracked.
		require.Equal(t, false, ht.honestBlockChalLevelZeroEdge.IsNone())
	})
	t.Run("edge is not honest but we agree with start commit and keep it as a rival", func(t *testing.T) {
		ht.metadataReader = &mockMetadataReader{
			assertionErr:  nil,
			assertionHash: ht.topLevelAssertionHash,
		}
		edge := newEdge(&newCfg{t: t, edgeId: "blk-0.a-32.b", createdAt: 1, claimId: "bar"})
		startHeight, startCommit := edge.StartCommitment()
		endHeight, endCommit := edge.EndCommitment()
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
				Height:     uint64(startHeight),
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
				Height:     uint64(endHeight),
				MerkleRoot: endCommit,
			},
		).Return(false, nil)
		ht.histChecker = mockStateManager
		agreement, err := ht.AddEdge(ctx, edge)
		require.NoError(t, err)
		require.Equal(t, protocol.Agreement{
			IsHonestEdge:          false,
			AgreesWithStartCommit: true,
		}, agreement)

		// Is not being tracked by the honest challenge tree.
		_, ok := ht.edges.TryGet(edge.Id())
		require.Equal(t, false, ok)
		// Exists in the mutual ids mapping.
		mutuals, ok := ht.mutualIds.TryGet(edge.MutualId())
		require.Equal(t, true, ok)
		require.Equal(t, true, mutuals.Has(edge.Id()))
		require.Equal(t, true, mutuals.NumItems() > 0)
	})
}

type mockMetadataReader struct {
	assertionHash            protocol.AssertionHash
	assertionErr             error
	claimHeights             protocol.OriginHeights
	claimHeightsErr          error
	unrivaledAssertionBlocks uint64
}

func (m *mockMetadataReader) TopLevelAssertion(
	_ context.Context, _ protocol.EdgeId,
) (protocol.AssertionHash, error) {
	return m.assertionHash, m.assertionErr
}

func (m *mockMetadataReader) AssertionUnrivaledBlocks(
	_ context.Context, _ protocol.AssertionHash,
) (uint64, error) {
	return m.unrivaledAssertionBlocks, nil
}

func (m *mockMetadataReader) TopLevelClaimHeights(
	_ context.Context, _ protocol.EdgeId,
) (protocol.OriginHeights, error) {
	return m.claimHeights, m.claimHeightsErr
}

func (m *mockMetadataReader) SpecChallengeManager(_ context.Context) (protocol.SpecChallengeManager, error) {
	return nil, nil
}
func (m *mockMetadataReader) ReadAssertionCreationInfo(
	_ context.Context, _ protocol.AssertionHash,
) (*protocol.AssertionCreatedInfo, error) {
	return &protocol.AssertionCreatedInfo{InboxMaxCount: big.NewInt(1)}, nil
}

type newCfg struct {
	t         *testing.T
	originId  mock.OriginId
	edgeId    mock.EdgeId
	claimId   string
	createdAt uint64
}

func newEdge(cfg *newCfg) *mock.Edge {
	cfg.t.Helper()
	items := strings.Split(string(cfg.edgeId), "-")
	var typ protocol.EdgeType
	switch items[0] {
	case "blk":
		typ = protocol.BlockChallengeEdge
	case "big":
		typ = protocol.BigStepChallengeEdge
	case "smol":
		typ = protocol.SmallStepChallengeEdge
	}
	startData := strings.Split(items[1], ".")
	startHeight, err := strconv.ParseUint(startData[0], 10, 64)
	require.NoError(cfg.t, err)
	startCommit := startData[1]

	endData := strings.Split(items[2], ".")
	endHeight, err := strconv.ParseUint(endData[0], 10, 64)
	require.NoError(cfg.t, err)
	endCommit := endData[1]

	return &mock.Edge{
		EdgeType:      typ,
		OriginID:      cfg.originId,
		ID:            cfg.edgeId,
		StartHeight:   startHeight,
		ClaimID:       cfg.claimId,
		StartCommit:   mock.Commit(startCommit),
		EndHeight:     endHeight,
		EndCommit:     mock.Commit(endCommit),
		LowerChildID:  "",
		UpperChildID:  "",
		CreationBlock: cfg.createdAt,
	}
}
