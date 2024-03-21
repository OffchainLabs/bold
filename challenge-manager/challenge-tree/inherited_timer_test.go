package challengetree

import (
	"context"
	"math"
	"testing"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/OffchainLabs/bold/containers/threadsafe"
	"github.com/stretchr/testify/require"
)

func TestUpdateInheritedTimer(t *testing.T) {
	ctx := context.Background()
	edge := newEdge(&newCfg{t: t, edgeId: "smol-0.a-1.a", createdAt: 0})
	edge.TotalChallengeLevels = 3
	edge.InnerStatus = protocol.EdgeConfirmed
	unrivaledAssertionBlocks := uint64(30)
	ht := &RoyalChallengeTree{
		edges:                 threadsafe.NewMap[protocol.EdgeId, protocol.SpecEdge](),
		edgeCreationTimes:     threadsafe.NewMap[OriginPlusMutualId, *threadsafe.Map[protocol.EdgeId, creationTime]](),
		royalRootEdgesByLevel: threadsafe.NewMap[protocol.ChallengeLevel, *threadsafe.Slice[protocol.SpecEdge]](),
		totalChallengeLevels:  3,
		metadataReader: &mockMetadataReader{
			assertionErr:             nil,
			assertionHash:            protocol.AssertionHash{},
			unrivaledAssertionBlocks: unrivaledAssertionBlocks,
		},
		inheritedTimers: threadsafe.NewMap[protocol.EdgeId, protocol.InheritedTimer](),
	}
	ht.edges.Put(edge.Id(), edge)

	t.Run("one step proven edge returns max uint64", func(t *testing.T) {
		timer, err := ht.UpdateInheritedTimer(ctx, protocol.AssertionHash{}, edge.Id(), 1)
		require.NoError(t, err)
		require.Equal(t, protocol.InheritedTimer(math.MaxUint64), timer)
	})
	t.Run("edge without children and not subchallenged returns time unrivaled", func(t *testing.T) {
		edge := newEdge(&newCfg{t: t, edgeId: "big-0.a-16.a", createdAt: 1})
		ht.edges.Put(edge.Id(), edge)
		timer, err := ht.UpdateInheritedTimer(ctx, protocol.AssertionHash{}, edge.Id(), 10)
		require.NoError(t, err)
		require.Equal(t, protocol.InheritedTimer(9), timer)
		timer, err = ht.UpdateInheritedTimer(ctx, protocol.AssertionHash{}, edge.Id(), 20)
		require.NoError(t, err)
		require.Equal(t, protocol.InheritedTimer(19), timer)
	})
	t.Run("edge with children inherits min of the children", func(t *testing.T) {
		edge := newEdge(&newCfg{t: t, edgeId: "big-0.a-16.a", createdAt: 1})
		lowerChild := newEdge(&newCfg{t: t, edgeId: "big-0.a-8.a", createdAt: 5})
		upperChild := newEdge(&newCfg{t: t, edgeId: "big-8.a-16.a", createdAt: 2})
		edge.LowerChildID = lowerChild.ID
		edge.UpperChildID = upperChild.ID
		ht.edges.Put(edge.Id(), edge)
		ht.edges.Put(lowerChild.Id(), lowerChild)
		ht.edges.Put(upperChild.Id(), upperChild)

		blockNum := uint64(10)
		timer, err := ht.UpdateInheritedTimer(ctx, protocol.AssertionHash{}, edge.Id(), blockNum)
		require.NoError(t, err)
		expectedEdgeLocalTimer := blockNum - edge.CreationBlock

		// Before updating the children, the edge should simply have a timer equal
		// to its local timer.
		require.Equal(t, expectedEdgeLocalTimer, uint64(timer))

		// Then, we update the children. We check the children's inherited timers.
		timer, err = ht.UpdateInheritedTimer(ctx, protocol.AssertionHash{}, lowerChild.Id(), blockNum)
		require.NoError(t, err)
		expectedLowerChild := blockNum - lowerChild.CreationBlock
		require.Equal(t, expectedLowerChild, uint64(timer))
		timer, err = ht.UpdateInheritedTimer(ctx, protocol.AssertionHash{}, upperChild.Id(), blockNum)
		require.NoError(t, err)
		expectedUpperChild := blockNum - upperChild.CreationBlock
		require.Equal(t, expectedUpperChild, uint64(timer))

		// Now, we update the parent again and should see it add to its local timer the minimum
		// of its children's inherited timers.
		timer, err = ht.UpdateInheritedTimer(ctx, protocol.AssertionHash{}, edge.Id(), blockNum)
		require.NoError(t, err)
		expected := expectedEdgeLocalTimer + expectedLowerChild
		require.Equal(t, expected, uint64(timer))
	})
	t.Run("edge with both children having maxuint64 timers inherits maxuint64", func(t *testing.T) {
		edge := newEdge(&newCfg{t: t, edgeId: "smol-4.a-6.a", createdAt: 1})
		lowerChild := newEdge(&newCfg{t: t, edgeId: "smol-4.a-5.a", createdAt: 5})
		upperChild := newEdge(&newCfg{t: t, edgeId: "smol-5.a-6.a", createdAt: 2})
		lowerChild.InnerStatus = protocol.EdgeConfirmed
		upperChild.InnerStatus = protocol.EdgeConfirmed
		edge.LowerChildID = lowerChild.ID
		edge.UpperChildID = upperChild.ID
		ht.edges.Put(edge.Id(), edge)
		ht.edges.Put(lowerChild.Id(), lowerChild)
		ht.edges.Put(upperChild.Id(), upperChild)

		// Before updating children.
		blockNum := uint64(10)
		timer, err := ht.UpdateInheritedTimer(ctx, protocol.AssertionHash{}, edge.Id(), blockNum)
		require.NoError(t, err)
		expected := blockNum - edge.CreationBlock
		require.Equal(t, expected, uint64(timer))

		timer, err = ht.UpdateInheritedTimer(ctx, protocol.AssertionHash{}, lowerChild.Id(), blockNum)
		require.NoError(t, err)
		require.Equal(t, uint64(math.MaxUint64), uint64(timer))
		timer, err = ht.UpdateInheritedTimer(ctx, protocol.AssertionHash{}, upperChild.Id(), blockNum)
		require.NoError(t, err)
		require.Equal(t, uint64(math.MaxUint64), uint64(timer))

		// After updating children.
		timer, err = ht.UpdateInheritedTimer(ctx, protocol.AssertionHash{}, edge.Id(), blockNum)
		require.NoError(t, err)
		require.Equal(t, uint64(math.MaxUint64), uint64(timer))
	})
	t.Run("edge with only one child having maxuint64 timers inherits the lower timer", func(t *testing.T) {
		edge := newEdge(&newCfg{t: t, edgeId: "smol-4.a-6.a", createdAt: 1})
		lowerChild := newEdge(&newCfg{t: t, edgeId: "smol-4.a-5.a", createdAt: 5})
		upperChild := newEdge(&newCfg{t: t, edgeId: "smol-5.a-6.a", createdAt: 2})
		lowerChild.InnerStatus = protocol.EdgeConfirmed
		edge.LowerChildID = lowerChild.ID
		edge.UpperChildID = upperChild.ID
		ht.edges.Put(edge.Id(), edge)
		ht.edges.Put(lowerChild.Id(), lowerChild)
		ht.edges.Put(upperChild.Id(), upperChild)

		// Before updating children.
		blockNum := uint64(10)
		timer, err := ht.UpdateInheritedTimer(ctx, protocol.AssertionHash{}, edge.Id(), blockNum)
		require.NoError(t, err)
		expected := blockNum - edge.CreationBlock
		require.Equal(t, expected, uint64(timer))

		timer, err = ht.UpdateInheritedTimer(ctx, protocol.AssertionHash{}, lowerChild.Id(), blockNum)
		require.NoError(t, err)
		require.Equal(t, uint64(math.MaxUint64), uint64(timer))
		timer, err = ht.UpdateInheritedTimer(ctx, protocol.AssertionHash{}, upperChild.Id(), blockNum)
		require.NoError(t, err)
		expectedUpperChild := blockNum - upperChild.CreationBlock
		require.Equal(t, expectedUpperChild, uint64(timer))

		// After updating children.
		timer, err = ht.UpdateInheritedTimer(ctx, protocol.AssertionHash{}, edge.Id(), blockNum)
		require.NoError(t, err)
		require.Equal(t, expected+expectedUpperChild, uint64(timer))
	})
	t.Run("edge with length one returns max of either its cached inherited or local timer or onchain timer", func(t *testing.T) {
		edge := newEdge(&newCfg{t: t, edgeId: "big-16.a-17.a", createdAt: 2})
		ht.edges.Put(edge.Id(), edge)
		blockNum := uint64(10)

		// Edge just returns its time unrivaled.
		timer, err := ht.UpdateInheritedTimer(ctx, protocol.AssertionHash{}, edge.Id(), blockNum)
		require.NoError(t, err)
		expected := blockNum - edge.CreationBlock
		require.Equal(t, expected, uint64(timer))

		// Locally cached value > time unrivaled, returns the former.
		ht.inheritedTimers.Put(edge.Id(), 20)
		timer, err = ht.UpdateInheritedTimer(ctx, protocol.AssertionHash{}, edge.Id(), blockNum)
		require.NoError(t, err)
		require.Equal(t, uint64(20), uint64(timer))

		// Onchain value is greater than all of them, returns onchain value
		// and checks we updated our local cache after that.
		edge.InnerInheritedTimer = 100
		ht.edges.Put(edge.Id(), edge)
		timer, err = ht.UpdateInheritedTimer(ctx, protocol.AssertionHash{}, edge.Id(), blockNum)
		require.NoError(t, err)
		require.Equal(t, uint64(100), uint64(timer))
		require.Equal(t, protocol.InheritedTimer(100), ht.inheritedTimers.Get(edge.Id()))
	})
	t.Run("edge that claims another edge updates that claimed edge's inherited timer", func(t *testing.T) {
		edge := newEdge(&newCfg{t: t, edgeId: "big-0.a-32.a", createdAt: 2})
		claimedEdge := newEdge(&newCfg{t: t, edgeId: "blk-0.a-1.a", createdAt: 1})
		edge.ClaimID = string(claimedEdge.ID)
		ht.edges.Put(edge.Id(), edge)
		ht.edges.Put(claimedEdge.Id(), claimedEdge)

		blockNum := uint64(10)
		timer, err := ht.UpdateInheritedTimer(ctx, protocol.AssertionHash{}, edge.Id(), blockNum)
		require.NoError(t, err)
		expected := blockNum - edge.CreationBlock
		require.Equal(t, expected, uint64(timer))
		require.Equal(t, expected, uint64(ht.inheritedTimers.Get(claimedEdge.Id())))
	})
	t.Run("claimed edge with onchain timer greater than locally computed value updates cache", func(t *testing.T) {
		edge := newEdge(&newCfg{t: t, edgeId: "big-0.a-32.a", createdAt: 2})
		claimedEdge := newEdge(&newCfg{t: t, edgeId: "blk-0.a-1.a", createdAt: 1})
		claimedEdge.InnerInheritedTimer = 100
		edge.ClaimID = string(claimedEdge.ID)
		ht.edges.Put(edge.Id(), edge)
		ht.edges.Put(claimedEdge.Id(), claimedEdge)

		blockNum := uint64(10)
		timer, err := ht.UpdateInheritedTimer(ctx, protocol.AssertionHash{}, edge.Id(), blockNum)
		require.NoError(t, err)
		expected := blockNum - edge.CreationBlock
		require.Equal(t, expected, uint64(timer))
		require.Equal(t, uint64(100), uint64(ht.inheritedTimers.Get(claimedEdge.Id())))
	})
	t.Run("root level block challenge edge computed inherited timer includes assertion unrivaled blocks", func(t *testing.T) {
		edge := newEdge(&newCfg{t: t, edgeId: "blk-0.a-32.a", createdAt: 2})
		edge.ClaimID = string("assertion")
		ht.edges.Put(edge.Id(), edge)

		blockNum := uint64(10)
		timer, err := ht.UpdateInheritedTimer(ctx, protocol.AssertionHash{}, edge.Id(), blockNum)
		require.NoError(t, err)
		expected := blockNum - edge.CreationBlock + unrivaledAssertionBlocks
		require.Equal(t, expected, uint64(timer))
	})
}

func TestIsOneStepProven(t *testing.T) {
	ctx := context.Background()
	edge := newEdge(&newCfg{t: t, edgeId: "big-0.a-32.a", createdAt: 0})
	require.Equal(t, false, isOneStepProven(ctx, edge, protocol.EdgePending))
	require.Equal(t, false, isOneStepProven(ctx, edge, protocol.EdgeConfirmed))
	edge = newEdge(&newCfg{t: t, edgeId: "big-0.a-1.a", createdAt: 0})
	require.Equal(t, false, isOneStepProven(ctx, edge, protocol.EdgeConfirmed))
	edge = newEdge(&newCfg{t: t, edgeId: "smol-0.a-1.a", createdAt: 0})
	edge.TotalChallengeLevels = 3
	require.Equal(t, true, isOneStepProven(ctx, edge, protocol.EdgeConfirmed))
}

func TestSaturatingSum(t *testing.T) {
	tests := []struct {
		a, b, expected protocol.InheritedTimer
	}{
		{10, 20, 30},
		{0, 0, 0},
		{math.MaxUint64, 0, math.MaxUint64},
		{0, math.MaxUint64, math.MaxUint64},
		{math.MaxUint64 - 1, 2, math.MaxUint64},
		{math.MaxUint64, math.MaxUint64, math.MaxUint64},
	}
	for _, test := range tests {
		result := saturatingSum(test.a, test.b)
		require.Equal(t, test.expected, result)
	}
}
