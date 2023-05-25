package challengetree

import (
	"context"
	"testing"

	"errors"
	"fmt"
	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/OffchainLabs/challenge-protocol-v2/util/threadsafe"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"strconv"
	"strings"
)

func TestAddEdge(t *testing.T) {
	ht := &HonestChallengeTree{
		edges:                         threadsafe.NewMap[protocol.EdgeId, protocol.ReadOnlyEdge](),
		mutualIds:                     threadsafe.NewMap[protocol.MutualId, *threadsafe.Map[protocol.EdgeId, creationTime]](),
		honestBigStepLevelZeroEdges:   threadsafe.NewSlice[protocol.ReadOnlyEdge](),
		honestSmallStepLevelZeroEdges: threadsafe.NewSlice[protocol.ReadOnlyEdge](),
	}
	ht.topLevelAssertionId = protocol.AssertionId(common.BytesToHash([]byte("foo")))
	ctx := context.Background()
	edge := newEdge(&newCfg{t: t, edgeId: "blk-0.a-16.a", createdAt: 1})

	t.Run("getting top level assertion fails", func(t *testing.T) {
		ht.metadataReader = &mockMetadataReader{
			assertionErr: errors.New("bad request"),
		}
		err := ht.AddEdge(ctx, edge)
		require.ErrorContains(t, err, "could not get top level assertion for edge")
	})
	t.Run("ignores if disagrees with top level assertion id of edge", func(t *testing.T) {
		ht.metadataReader = &mockMetadataReader{
			assertionErr: nil,
			assertionId:  protocol.AssertionId(common.BytesToHash([]byte("bar"))),
		}
		err := ht.AddEdge(ctx, edge)
		require.NoError(t, err)
	})
	t.Run("getting claim heights fails", func(t *testing.T) {
		ht.metadataReader = &mockMetadataReader{
			assertionErr:    nil,
			assertionId:     ht.topLevelAssertionId,
			claimHeightsErr: errors.New("bad request"),
		}
		err := ht.AddEdge(ctx, edge)
		require.ErrorContains(t, err, "could not get claim heights for edge")
	})
	t.Run("checking if agrees with commit fails", func(t *testing.T) {
		ht.metadataReader = &mockMetadataReader{
			assertionErr: nil,
			assertionId:  ht.topLevelAssertionId,
		}
		ht.histChecker = &mockHistChecker{
			agreesErr: errors.New("bad request"),
		}
		err := ht.AddEdge(ctx, edge)
		require.ErrorContains(t, err, "could not check if agrees with")
	})
	t.Run("fully disagrees with edge", func(t *testing.T) {
		ht.histChecker = &mockHistChecker{
			agreement: Agreement{
				IsHonestEdge:          false,
				AgreesWithStartCommit: false,
			},
		}
		badEdge := newEdge(&newCfg{t: t, edgeId: "blk-0.f-16.a", createdAt: 1})
		err := ht.AddEdge(ctx, badEdge)
		require.NoError(t, err)

		// Check the edge is not kept track of anywhere.
		_, ok := ht.edges.TryGet(badEdge.Id())
		require.Equal(t, false, ok)
		_, ok = ht.mutualIds.TryGet(badEdge.MutualId())
		require.Equal(t, false, ok)
	})
	t.Run("agrees with edge but is not a level zero edge", func(t *testing.T) {
		ht.histChecker = &mockHistChecker{
			agreement: Agreement{
				IsHonestEdge: true,
			},
		}
		edge := newEdge(&newCfg{t: t, edgeId: "blk-0.a-16.a", createdAt: 1})
		err := ht.AddEdge(ctx, edge)
		require.NoError(t, err)

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
		edge := newEdge(&newCfg{t: t, edgeId: "blk-0.a-32.a", createdAt: 1, claimId: "foo"})
		err := ht.AddEdge(ctx, edge)
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
		ht.histChecker = &mockHistChecker{
			agreement: Agreement{
				IsHonestEdge:          false,
				AgreesWithStartCommit: true,
			},
		}
		edge := newEdge(&newCfg{t: t, edgeId: "blk-0.a-32.b", createdAt: 1, claimId: "bar"})
		err := ht.AddEdge(ctx, edge)
		require.NoError(t, err)

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
	assertionId     protocol.AssertionId
	assertionErr    error
	claimHeights    *ClaimHeights
	claimHeightsErr error
}

func (m *mockMetadataReader) TopLevelAssertion(
	_ context.Context, _ protocol.EdgeId,
) (protocol.AssertionId, error) {
	return m.assertionId, m.assertionErr
}

func (*mockMetadataReader) AssertionUnrivaledTime(
	_ context.Context, _ protocol.AssertionId,
) (uint64, error) {
	return 0, nil
}

func (m *mockMetadataReader) ClaimHeights(
	_ context.Context, _ protocol.EdgeId,
) (*ClaimHeights, error) {
	return m.claimHeights, m.claimHeightsErr
}

type mockHistChecker struct {
	agreement Agreement
	agreesErr error
}

func (m *mockHistChecker) AgreesWithHistoryCommitment(
	_ context.Context,
	_ *ClaimHeights,
	_,
	_ util.HistoryCommitment,
) (Agreement, error) {
	return m.agreement, m.agreesErr
}

var _ = protocol.ReadOnlyEdge(&edge{})

type edgeId string
type commit string
type originId string

// Mock edge for challenge tree specific tests, making it easier for test ergonomics.
type edge struct {
	id            edgeId
	edgeType      protocol.EdgeType
	startHeight   uint64
	startCommit   commit
	endHeight     uint64
	endCommit     commit
	originId      originId
	claimId       string
	lowerChildId  edgeId
	upperChildId  edgeId
	creationBlock uint64
}

func (e *edge) Id() protocol.EdgeId {
	return protocol.EdgeId(common.BytesToHash([]byte(e.id)))
}

func (e *edge) GetType() protocol.EdgeType {
	return e.edgeType
}

func (e *edge) StartCommitment() (protocol.Height, common.Hash) {
	return protocol.Height(e.startHeight), common.BytesToHash([]byte(e.startCommit))
}

func (e *edge) EndCommitment() (protocol.Height, common.Hash) {
	return protocol.Height(e.endHeight), common.BytesToHash([]byte(e.endCommit))
}

func (e *edge) CreatedAtBlock() uint64 {
	return e.creationBlock
}

func (e *edge) OriginId() protocol.OriginId {
	return protocol.OriginId(common.BytesToHash([]byte(e.originId)))
}

func (e *edge) MutualId() protocol.MutualId {
	return protocol.MutualId(common.BytesToHash([]byte(e.computeMutualId())))
}

func (e *edge) computeMutualId() string {
	return fmt.Sprintf(
		"%d-%s-%d-%s-%d",
		e.edgeType,
		e.originId,
		e.startHeight,
		e.startCommit,
		e.endHeight,
	)
}

// The claim id of the edge, if any
func (e *edge) ClaimId() util.Option[protocol.ClaimId] {
	if e.claimId == "" {
		return util.None[protocol.ClaimId]()
	}
	return util.Some(protocol.ClaimId(common.BytesToHash([]byte(e.claimId))))
}

// The lower child of the edge, if any.
func (e *edge) LowerChild(_ context.Context) (util.Option[protocol.EdgeId], error) {
	if e.lowerChildId == "" {
		return util.None[protocol.EdgeId](), nil
	}
	return util.Some(protocol.EdgeId(common.BytesToHash([]byte(e.lowerChildId)))), nil
}

// The upper child of the edge, if any.
func (e *edge) UpperChild(_ context.Context) (util.Option[protocol.EdgeId], error) {
	if e.upperChildId == "" {
		return util.None[protocol.EdgeId](), nil
	}
	return util.Some(protocol.EdgeId(common.BytesToHash([]byte(e.upperChildId)))), nil
}

// The ministaker of an edge. Only existing for level zero edges.
func (*edge) MiniStaker() util.Option[common.Address] {
	return util.None[common.Address]()
}

// The assertion id of the parent assertion that originated the challenge
// at the top-level.
func (*edge) PrevAssertionId(_ context.Context) (protocol.AssertionId, error) {
	return protocol.AssertionId{}, errors.New("unimplemented")
}

// The time in seconds an edge has been unrivaled.
func (*edge) TimeUnrivaled(_ context.Context) (uint64, error) {
	return 0, errors.New("unimplemented")
}

// The status of an edge.
func (*edge) Status(_ context.Context) (protocol.EdgeStatus, error) {
	return 0, errors.New("unimplemented")
}

// Whether or not an edge has rivals.
func (*edge) HasRival(_ context.Context) (bool, error) {
	return false, errors.New("unimplemented")
}

// Checks if an edge has a length one rival.
func (*edge) HasLengthOneRival(_ context.Context) (bool, error) {
	return false, errors.New("unimplemented")
}

// The history commitment for the top-level edge the current edge's challenge is made upon.
// This is used at subchallenge creation boundaries.
func (*edge) TopLevelClaimHeight(_ context.Context) (*protocol.OriginHeights, error) {
	return nil, errors.New("unimplemented")
}

type newCfg struct {
	t         *testing.T
	originId  originId
	edgeId    edgeId
	claimId   string
	createdAt uint64
}

func newEdge(cfg *newCfg) *edge {
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

	return &edge{
		edgeType:      typ,
		originId:      cfg.originId,
		id:            cfg.edgeId,
		startHeight:   startHeight,
		claimId:       cfg.claimId,
		startCommit:   commit(startCommit),
		endHeight:     endHeight,
		endCommit:     commit(endCommit),
		lowerChildId:  "",
		upperChildId:  "",
		creationBlock: cfg.createdAt,
	}
}
