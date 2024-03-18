package edgetracker

import (
	"context"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
)

type challengeConfirmer struct {
	reader challengeReader
	writer chainWriter
}
type chainWriter interface {
	UpdateInheritedTimers(
		ctx context.Context,
		royalBranch []protocol.EdgeId,
	) error
}

type challengeReader interface {
	BlockChallengeRootEdge(
		ctx context.Context,
		challengedAssertionHash protocol.AssertionHash,
	) (protocol.SpecEdge, error)
	LowerMostRoyalEdges(
		ctx context.Context,
		challengedAssertionHash protocol.AssertionHash,
	) ([]protocol.SpecEdge, error)
	ComputeAncestors(
		ctx context.Context,
		challengedAssertionHash protocol.AssertionHash,
		edgeId protocol.EdgeId,
	) ([]protocol.EdgeId, error)
}

// A challenge confirmation job will attempt to confirm a challenge all the way up to the top,
// block challenge root edge by updating all the inherited timers of royal edges along the way,
// across all open subchallenges, until the onchain timer of the block challenge root edge
// is >= a challenge period.
//
// It operates in branches of logic, starting from the bottom-most, deepest level royal edges
// and for each branch, update the onchain inherited timers of the ancestors along the way.
func (cc *challengeConfirmer) beginConfirmationJob(
	ctx context.Context,
	challengedAssertionHash protocol.AssertionHash,
	challengePeriodBlocks uint64,
) error {
	// Check if the edge is already confirmable onchain and
	// whether or not we can exit this computation early.
	royalRootEdge, err := cc.reader.BlockChallengeRootEdge(ctx, challengedAssertionHash)
	if err != nil {
		return err
	}
	rootTimer, err := royalRootEdge.InheritedTimer(ctx)
	if err != nil {
		return err
	}
	// Short circuit early and confirm right away.
	if uint64(rootTimer) >= challengePeriodBlocks {
		return royalRootEdge.ConfirmByTimer(ctx)
	}

	// Find the bottom-most royal edges that exist in our local challenge tree, each one
	// will be the base of a branch we will update.
	bases, err := cc.reader.LowerMostRoyalEdges(ctx, challengedAssertionHash)
	if err != nil {
		return err
	}
	// For each branch, compute the royal ancestor branch up to the root of the tree
	royalBranches := make([][]protocol.EdgeId, 0)
	for _, edge := range bases {
		branch := []protocol.EdgeId{edge.Id()}
		ancestors, err := cc.reader.ComputeAncestors(
			ctx, challengedAssertionHash, edge.Id(),
		)
		if err != nil {
			return err
		}
		branch = append(branch, ancestors...)
		royalBranches = append(royalBranches, branch)
	}
	// For each branch, update the inherited timers onchain in a single transaction.
	for _, branch := range royalBranches {
		if err := cc.writer.UpdateInheritedTimers(ctx, branch); err != nil {
			return err
		}
		// In each iteration, check if the root edge has a timer >= a challenge period
		rootTimer, err := royalRootEdge.InheritedTimer(ctx)
		if err != nil {
			return err
		}
		// If yes, we confirm the root edge and finish early.
		if uint64(rootTimer) >= challengePeriodBlocks {
			return royalRootEdge.ConfirmByTimer(ctx)
		}
	}
	// If the royal timer is not >= a challenge period by the end of this job,
	// it means the challenge has yet to complete and our local computation was incorrect.
	// In this scenario, we can dump the confirmation job of royal edges for manual
	// inspection and debugging
	return nil
}
