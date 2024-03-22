package challengetree

import (
	"context"
	"fmt"
	"math"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	edgetracker "github.com/OffchainLabs/bold/challenge-manager/edge-tracker"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"golang.org/x/exp/constraints"
)

func (ht *RoyalChallengeTree) TimeCacheUpdate(
	ctx context.Context,
	challengedAssertionHash protocol.AssertionHash,
	blockNum uint64,
) (protocol.InheritedTimer, error) {
	royalRootEdge, err := ht.BlockChallengeRootEdge(ctx)
	if err != nil {
		return 0, err
	}
	assertionUnrivaledBlocks, err := ht.metadataReader.AssertionUnrivaledBlocks(
		ctx,
		protocol.AssertionHash{
			Hash: common.Hash(royalRootEdge.ClaimId().Unwrap()),
		},
	)
	if err != nil {
		return 0, err
	}
	inheritedTimer, err := ht.recursiveCacheUpdate(ctx, royalRootEdge.Id(), blockNum)
	if err != nil {
		return 0, err
	}
	return saturatingSum(inheritedTimer, protocol.InheritedTimer(assertionUnrivaledBlocks)), nil
}

func (ht *RoyalChallengeTree) recursiveCacheUpdate(
	ctx context.Context,
	edgeId protocol.EdgeId,
	blockNum uint64,
) (protocol.InheritedTimer, error) {
	edge, ok := ht.edges.TryGet(edgeId)
	if !ok {
		return 0, nil
	}
	status, err := edge.Status(ctx)
	if err != nil {
		return 0, err
	}
	if isOneStepProven(ctx, edge, status) {
		return math.MaxUint64, nil
	}
	localTimer, err := ht.LocalTimer(edge, blockNum)
	if err != nil {
		return 0, err
	}
	// If length one, find the edge that claims it,
	// compute the recursive timer for it. If the onchain is bigger, return the onchain here.
	if hasLengthOne(edge) {
		onchainTimer, err := edge.InheritedTimer(ctx)
		if err != nil {
			return 0, err
		}
		claimingEdgeTimer := protocol.InheritedTimer(0)
		claimingEdge, ok := ht.findClaimingEdge(ctx, edge.Id())
		if ok {
			claimingEdgeTimer, err = ht.recursiveCacheUpdate(
				ctx,
				claimingEdge.Id(),
				blockNum,
			)
			if err != nil {
				return 0, err
			}
		}
		claimedEdgeInheritedTimer := protocol.InheritedTimer(localTimer + uint64(claimingEdgeTimer))
		if onchainTimer > claimedEdgeInheritedTimer {
			return onchainTimer, nil
		}
		return claimedEdgeInheritedTimer, nil
	}
	hasChildren, err := edge.HasChildren(ctx)
	if err != nil {
		return 0, err
	}
	if !hasChildren {
		return protocol.InheritedTimer(localTimer), nil
	}
	lowerChildId, err := edge.LowerChild(ctx)
	if err != nil {
		return 0, err
	}
	upperChildId, err := edge.UpperChild(ctx)
	if err != nil {
		return 0, err
	}
	lowerChildTimer, err := ht.recursiveCacheUpdate(ctx, lowerChildId.Unwrap(), blockNum)
	if err != nil {
		return 0, err
	}
	upperChildTimer, err := ht.recursiveCacheUpdate(ctx, upperChildId.Unwrap(), blockNum)
	if err != nil {
		return 0, err
	}
	minTimer := lowerChildTimer
	if upperChildTimer < lowerChildTimer {
		minTimer = lowerChildTimer
	}
	return protocol.InheritedTimer(localTimer) + minTimer, nil
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
		ht.inheritedTimers.Put(edgeId, protocol.InheritedTimer(math.MaxUint64))
		return math.MaxUint64, nil
	}
	localTimer, err := ht.LocalTimer(edge, blockNum)
	if err != nil {
		return 0, err
	}

	// If the edge has a length one, we return return the max of the following:
	// - the local timer of the edge
	// - the cached inherited timer from any lower-level edges that may claim it
	// - or the actual onchain timer of the edge
	if hasLengthOne(edge) {
		var onchainTimer protocol.InheritedTimer
		onchainTimer, err = edge.InheritedTimer(ctx)
		if err != nil {
			return 0, err
		}
		cachedInheritedTimer, ok := ht.inheritedTimers.TryGet(edgeId)
		if !ok {
			cachedInheritedTimer = 0
		}
		maxVal := max(cachedInheritedTimer, onchainTimer, protocol.InheritedTimer(localTimer))
		if maxVal > cachedInheritedTimer {
			ht.inheritedTimers.Put(edgeId, maxVal)
		}
		return maxVal, nil
	}

	inheritedTimer := protocol.InheritedTimer(localTimer)

	// If an edge has children, we inherit the minimum timer of its children.
	hasChildren, err := edge.HasChildren(ctx)
	if err != nil {
		return 0, err
	}
	if hasChildren {
		timerFromChildren, err2 := ht.inheritTimerFromChildren(ctx, edge)
		if err2 != nil {
			return 0, err2
		}
		inheritedTimer = saturatingSum(timerFromChildren, inheritedTimer)
	}

	// Edges that claim another edge in the level above need to read the onchain inherited timer
	// of the claimed edge, as it could have been updated by some malicious adversary outside of
	// our view. If this happens, the onchain inherited timer of the claimed edge will be greater
	// than we locally computed, and then we cache the larger value locally.
	if IsClaimingAnEdge(edge) {
		claimedEdgeId := edge.ClaimId().Unwrap()
		claimId := protocol.EdgeId{Hash: common.Hash(claimedEdgeId)}
		claimedEdge, ok := ht.edges.TryGet(claimId)
		if !ok {
			return 0, fmt.Errorf("claimed edge %#x not found in royal tree", claimedEdgeId)
		}
		onchainTimer, err := claimedEdge.InheritedTimer(ctx)
		if err != nil {
			return 0, errors.Wrapf(err, "could not get inherited timer for claimed edge %#x", claimedEdgeId)
		}
		if onchainTimer > inheritedTimer {
			ht.inheritedTimers.Put(claimId, onchainTimer)
		} else if onchainTimer < inheritedTimer {
			ht.inheritedTimers.Put(claimId, inheritedTimer)
		}
	} else if edgetracker.IsRootBlockChallengeEdge(edge) {
		assertionUnrivaledBlocks, err := ht.metadataReader.AssertionUnrivaledBlocks(ctx, protocol.AssertionHash{Hash: common.Hash(edge.ClaimId().Unwrap())})
		if err != nil {
			return 0, err
		}
		ht.inheritedTimers.Put(edgeId, inheritedTimer)
		return saturatingSum(inheritedTimer, protocol.InheritedTimer(assertionUnrivaledBlocks)), nil
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
	lowerTimer, ok := ht.inheritedTimers.TryGet(lowerChildId)
	if !ok {
		lowerTimer = 0
	}
	upperTimer, ok := ht.inheritedTimers.TryGet(upperChildId)
	if !ok {
		upperTimer = 0
	}
	if upperTimer < lowerTimer {
		return upperTimer, nil
	}
	return lowerTimer, nil
}

func IsClaimingAnEdge(edge protocol.ReadOnlyEdge) bool {
	return edge.ClaimId().IsSome() && edge.GetChallengeLevel() != protocol.NewBlockChallengeLevel()
}

func hasLengthOne(edge protocol.ReadOnlyEdge) bool {
	startHeight, _ := edge.StartCommitment()
	endHeight, _ := edge.EndCommitment()
	return endHeight-startHeight == 1
}

func isOneStepProven(
	ctx context.Context, edge protocol.ReadOnlyEdge, status protocol.EdgeStatus,
) bool {
	isSmallStep := edge.GetChallengeLevel() == protocol.ChallengeLevel(edge.GetTotalChallengeLevels(ctx)-1)
	return isSmallStep && status == protocol.EdgeConfirmed && hasLengthOne(edge)
}

func saturatingSum(a, b protocol.InheritedTimer) protocol.InheritedTimer {
	if math.MaxUint64-a < b {
		return math.MaxUint64
	}
	return a + b
}

func max[T constraints.Ordered](values ...T) T {
	var zeroVal T
	if len(values) == 0 {
		return zeroVal
	}
	maxVal := values[0]
	for _, val := range values[1:] {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}
