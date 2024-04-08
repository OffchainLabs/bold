package challengetree

import (
	"context"
	"math"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/ethereum/go-ethereum/common"
)

func (ht *RoyalChallengeTree) ComputeRootInheritedTimer(
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
	inheritedTimer, err := ht.iterativeInheritedTimerCompute(ctx, royalRootEdge.Id(), blockNum)
	if err != nil {
		return 0, err
	}
	return saturatingSum(inheritedTimer, protocol.InheritedTimer(assertionUnrivaledBlocks)), nil
}

func (ht *RoyalChallengeTree) iterativeInheritedTimerCompute(
	ctx context.Context,
	edgeId protocol.EdgeId,
	blockNum uint64,
) (protocol.InheritedTimer, error) {
	// inheritedTimers maps edgeId to the inherited timer of that edge
	// This is used to avoid recomputing the inherited timer of an edge
	inheritedTimers := make(map[protocol.EdgeId]protocol.InheritedTimer)
	// stack is used to keep track of the edges whose inherited timer needs to be computed
	stack := []protocol.EdgeId{edgeId}

	var result protocol.InheritedTimer
	for len(stack) > 0 {
		top := len(stack) - 1
		item := stack[top]
		stack = stack[:top]

		edge, ok := ht.edges.TryGet(item)
		if !ok {
			continue
		}
		status, err := edge.Status(ctx)
		if err != nil {
			return 0, err
		}
		if isOneStepProven(ctx, edge, status) {
			result = math.MaxUint64
			inheritedTimers[edge.Id()] = result
			continue
		}
		localTimer, err := ht.LocalTimer(edge, blockNum)
		if err != nil {
			return 0, err
		}
		// If length one, find the edge that claims it,
		// compute the recursive timer for it. If the onchain is bigger, return the onchain here.
		if hasLengthOne(edge) {
			onchainTimer, innerErr := edge.SafeHeadInheritedTimer(ctx)
			if innerErr != nil {
				return 0, innerErr
			}
			claimingEdgeTimer := protocol.InheritedTimer(0)
			claimingEdge, ok := ht.findClaimingEdge(ctx, edge.Id())
			if ok {
				claimingEdgeTimerFound := false
				claimingEdgeTimer, claimingEdgeTimerFound = inheritedTimers[claimingEdge.Id()]
				if !claimingEdgeTimerFound {
					stack = append(stack, item)
					stack = append(stack, claimingEdge.Id())
					continue
				}
			}
			claimedEdgeInheritedTimer := saturatingSum(protocol.InheritedTimer(localTimer), claimingEdgeTimer)
			if onchainTimer > claimedEdgeInheritedTimer {
				result = onchainTimer
				inheritedTimers[edge.Id()] = result
			} else {
				result = claimedEdgeInheritedTimer
				inheritedTimers[edge.Id()] = result
			}
			continue
		}
		hasChildren, err := edge.HasChildren(ctx)
		if err != nil {
			return 0, err
		}
		if !hasChildren {
			result = protocol.InheritedTimer(localTimer)
			inheritedTimers[edge.Id()] = result
			continue
		}
		lowerChildId, err := edge.LowerChild(ctx)
		if err != nil {
			return 0, err
		}
		upperChildId, err := edge.UpperChild(ctx)
		if err != nil {
			return 0, err
		}
		upperChildTimer, upperChildTimerFound := inheritedTimers[upperChildId.Unwrap()]
		lowerChildTimer, lowerChildTimerFound := inheritedTimers[lowerChildId.Unwrap()]
		if !upperChildTimerFound || !lowerChildTimerFound {
			stack = append(stack, item)
			stack = append(stack, upperChildId.Unwrap())
			stack = append(stack, lowerChildId.Unwrap())
			continue
		}
		minTimer := lowerChildTimer
		if upperChildTimer < lowerChildTimer {
			minTimer = upperChildTimer
		}
		result = saturatingSum(protocol.InheritedTimer(localTimer), minTimer)
		inheritedTimers[edge.Id()] = result
	}
	return result, nil
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
