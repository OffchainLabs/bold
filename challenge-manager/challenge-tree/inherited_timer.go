package challengetree

import (
	"context"
	"fmt"
	"math"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/ethereum/go-ethereum/common"
)

func (ht *RoyalChallengeTree) inheritedTimerForEdge(
	edge protocol.ReadOnlyEdge,
	blockNum uint64,
) (protocol.InheritedTimer, error) {
	localTimer, err := ht.LocalTimer(edge, blockNum)
	if err != nil {
		return 0, err
	}
	inherited, ok := ht.inheritedTimers.TryGet(edge.Id())
	if !ok {
		inherited = protocol.InheritedTimer(localTimer)
		ht.inheritedTimers.Put(edge.Id(), inherited)
	}
	if localTimer > uint64(inherited) {
		inherited = protocol.InheritedTimer(localTimer)
		ht.inheritedTimers.Put(edge.Id(), inherited)
	}
	return inherited, nil
}

func (ht *RoyalChallengeTree) UpdateInheritedTimer(
	ctx context.Context,
	challengedAssertionHash protocol.AssertionHash,
	edgeId protocol.EdgeId,
	blockNum uint64,
) (protocol.InheritedTimer, error) {
	edge, ok := ht.edges.TryGet(edgeId)
	if !ok {
		return 0, fmt.Errorf("edge with id %#x not found", edgeId.Hash)
	}
	status, err := edge.Status(ctx)
	if err != nil {
		return 0, err
	}
	// One step proven edges have an infinite inherited timer.
	if isOneStepProven(ctx, edge, status) {
		return math.MaxUint64, nil
	}
	localTimer, err := ht.LocalTimer(edge, blockNum)
	if err != nil {
		return 0, err
	}
	inheritedTimer := protocol.InheritedTimer(localTimer)

	// If an edge has children, we inherit the minimum timer of its children.
	hasChildren, err := edge.HasChildren(ctx)
	if err != nil {
		return 0, err
	}
	if hasChildren {
		timerFromChildren, err2 := ht.inheritTimerFromChildren(ctx, edge, blockNum)
		if err2 != nil {
			return 0, err2
		}
		inheritedTimer = saturatingSum(inheritedTimer, timerFromChildren)
	}

	// Edges that claim another edge in the level above update the inherited timer onchain
	// if they are able to.
	if IsClaimingAnEdge(edge) {
		claimedEdgeId := edge.ClaimId().Unwrap()
		claimedEdge, ok := ht.edges.TryGet(protocol.EdgeId{Hash: common.Hash(claimedEdgeId)})
		if !ok {
			return 0, fmt.Errorf("claimed edge %#x not found in royal tree", claimedEdgeId)
		}
		claimedEdgeLocalTimer, err := ht.LocalTimer(claimedEdge, blockNum)
		if err != nil {
			return 0, err
		}
		claimedEdgeTimeUnrivaled := saturatingSum(protocol.InheritedTimer(claimedEdgeLocalTimer), inheritedTimer)
		claimedEdgeInheritedTimer, ok := ht.inheritedTimers.TryGet(claimedEdge.Id())
		if !ok {
			claimedEdgeInheritedTimer = claimedEdgeTimeUnrivaled
			ht.inheritedTimers.Put(claimedEdge.Id(), claimedEdgeInheritedTimer)
		}
		if claimedEdgeTimeUnrivaled > claimedEdgeInheritedTimer {
			claimedEdgeInheritedTimer = claimedEdgeTimeUnrivaled
			ht.inheritedTimers.Put(claimedEdge.Id(), claimedEdgeInheritedTimer)
		}
	} else if edge.ClaimId().IsSome() && edge.GetChallengeLevel().IsBlockChallengeLevel() {
		assertionUnrivaledBlocks, err := ht.metadataReader.AssertionUnrivaledBlocks(ctx, challengedAssertionHash, blockNum)
		if err != nil {
			return 0, err
		}
		inheritedTimer = saturatingSum(inheritedTimer, protocol.InheritedTimer(assertionUnrivaledBlocks))
	}
	// Otherwise, the edge does not yet have children, we simply update its timer.
	ht.inheritedTimers.Put(edgeId, inheritedTimer)
	return inheritedTimer, nil
}

// Gets the inherited timer from the children of an edge. The edge
// must have children for this function to be called.
func (ht *RoyalChallengeTree) inheritTimerFromChildren(
	ctx context.Context,
	edge protocol.SpecEdge,
	blockNum uint64,
) (protocol.InheritedTimer, error) {
	lowerChildOpt, err := edge.LowerChild(ctx)
	if err != nil {
		return 0, err
	}
	upperChildOpt, err := edge.UpperChild(ctx)
	if err != nil {
		return 0, err
	}
	lowerChildId := lowerChildOpt.Unwrap()
	upperChildId := upperChildOpt.Unwrap()

	lowerChild, ok := ht.edges.TryGet(lowerChildId)
	if !ok {
		return 0, fmt.Errorf("lower child for edge %#x not yet tracked locally", edge.Id().Hash)
	}
	upperChild, ok := ht.edges.TryGet(upperChildId)
	if !ok {
		return 0, fmt.Errorf("upper child for edge %#x not yet tracked locally", edge.Id().Hash)
	}
	lowerTimer, err := ht.inheritedTimerForEdge(lowerChild, blockNum)
	if err != nil {
		return 0, err
	}
	upperTimer, err := ht.inheritedTimerForEdge(upperChild, blockNum)
	if err != nil {
		return 0, err
	}
	if upperTimer < lowerTimer {
		return upperTimer, nil
	}
	return lowerTimer, nil
}

func IsClaimingAnEdge(edge protocol.ReadOnlyEdge) bool {
	return edge.ClaimId().IsSome() && edge.GetChallengeLevel() != protocol.NewBlockChallengeLevel()
}

func isOneStepProven(
	ctx context.Context, edge protocol.ReadOnlyEdge, status protocol.EdgeStatus,
) bool {
	startHeight, _ := edge.StartCommitment()
	endHeight, _ := edge.EndCommitment()
	diff := endHeight - startHeight
	isSmallStep := edge.GetChallengeLevel() == protocol.ChallengeLevel(edge.GetTotalChallengeLevels(ctx)-1)
	return isSmallStep && status == protocol.EdgeConfirmed && diff == 1
}

func saturatingSum(a, b protocol.InheritedTimer) protocol.InheritedTimer {
	if math.MaxUint64-a < b {
		return math.MaxUint64
	}
	return a + b
}
