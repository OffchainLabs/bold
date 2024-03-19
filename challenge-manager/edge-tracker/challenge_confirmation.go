package edgetracker

import (
	"context"
	"fmt"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	retry "github.com/OffchainLabs/bold/runtime"
)

type challengeConfirmer struct {
	reader challengeReader
	writer chainWriter
}
type chainWriter interface {
	MultiUpdateInheritedTimers(
		ctx context.Context,
		challengeBranch []protocol.EdgeId,
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
// is greater than or equal to a challenge period.
//
// It works by updating royal branches of the challenge tree, starting from the bottom-most,
// deepest level royal edges. For each branch, update the onchain inherited timers
// of the ancestors along the way.
//
// This function must only be called once the locally computed value of the block challenge, royal root
// edge has an inherited timer that is confirmable. This function MUST complete, and it will retry
// any external call if it errors during its execution.
func (cc *challengeConfirmer) beginConfirmationJob(
	ctx context.Context,
	challengedAssertionHash protocol.AssertionHash,
	royalRootEdge protocol.SpecEdge,
	challengePeriodBlocks uint64,
) error {
	// Find the bottom-most royal edges that exist in our local challenge tree, each one
	// will be the base of a branch we will update.
	royalTreeLeaves, err := retry.UntilSucceeds(ctx, func() ([]protocol.SpecEdge, error) {
		return cc.reader.LowerMostRoyalEdges(ctx, challengedAssertionHash)
	})
	if err != nil {
		return err
	}
	// For each branch, compute the royal ancestor branch up to the root of the tree.
	// The branch should contain royal ancestors ordered from a bottom-most leaf edge to the root edge
	// of the block level challenge, meaning it should also include claim id links.
	royalBranches := make([][]protocol.EdgeId, 0)
	for _, edge := range royalTreeLeaves {
		branch := []protocol.EdgeId{edge.Id()}
		ancestors, err := retry.UntilSucceeds(ctx, func() ([]protocol.EdgeId, error) {
			return cc.reader.ComputeAncestors(
				ctx, challengedAssertionHash, edge.Id(),
			)
		})
		if err != nil {
			return err
		}
		branch = append(branch, ancestors...)
		royalBranches = append(royalBranches, branch)
	}
	// For each branch, update the inherited timers onchain in a single transaction.
	for _, branch := range royalBranches {
		if err := cc.writer.MultiUpdateInheritedTimers(ctx, branch); err != nil {
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
	onchainInheritedTimer, err := retry.UntilSucceeds(ctx, func() (protocol.InheritedTimer, error) {
		return royalRootEdge.InheritedTimer(ctx)
	})
	if err != nil {
		return err
	}
	// If the onchain timer is not >= a challenge period by the end of this job,
	// it means the challenge has yet to complete and our local computation was incorrect.
	// In this scenario, we can dump the confirmation job of royal edges for manual
	// inspection and debugging
	if onchainInheritedTimer < protocol.InheritedTimer(challengePeriodBlocks) {
		return fmt.Errorf(
			"onchain timer %d after confirmation job was executed < challenge period %d",
			onchainInheritedTimer,
			challengePeriodBlocks,
		)
	}
	return nil
}
