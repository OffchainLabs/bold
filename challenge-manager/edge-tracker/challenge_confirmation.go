package edgetracker

import (
	"context"
	"fmt"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	retry "github.com/OffchainLabs/bold/runtime"
	"github.com/ethereum/go-ethereum/log"
)

// Defines a struct which can handle confirming of an entire challenge tree
// in the BOLD protocol. It does so by updating the inherited timers of royal edges
// onchain until the root of the tree has a timer >= a challenge period. At that point,
// it ensures to confirm that edge. If this is not the case, it will return an error
// and write data to disk to help with debugging the issue.
type challengeConfirmer struct {
	reader        RoyalChallengeReader
	writer        ChainWriter
	validatorName string
}

// Defines a chain writer interface that is
// used to update the cached inherited timers of edges
// onchain.
type ChainWriter interface {
	MultiUpdateInheritedTimers(
		ctx context.Context,
		challengeBranch []protocol.ReadOnlyEdge,
	) error
}

type RoyalChallengeReader interface {
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
	) ([]protocol.ReadOnlyEdge, error)
}

func newChallengeConfirmer(
	challengeReader RoyalChallengeReader,
	chainWriter ChainWriter,
	validatorName string,
) *challengeConfirmer {
	return &challengeConfirmer{
		reader:        challengeReader,
		writer:        chainWriter,
		validatorName: validatorName,
	}
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
	fields := log.Ctx{
		"validatorName": cc.validatorName,
	}
	srvlog.Info("Starting challenge confirmation job", fields)
	// Find the bottom-most royal edges that exist in our local challenge tree, each one
	// will be the base of a branch we will update.
	royalTreeLeaves, err := retry.UntilSucceeds(ctx, func() ([]protocol.SpecEdge, error) {
		return cc.reader.LowerMostRoyalEdges(ctx, challengedAssertionHash)
	})
	if err != nil {
		return err
	}
	srvlog.Info("Got all the royal tree leaves", log.Ctx{
		"validatorName": cc.validatorName,
		"numLeaves":     len(royalTreeLeaves),
	})
	// For each branch, compute the royal ancestor branch up to the root of the tree.
	// The branch should contain royal ancestors ordered from a bottom-most leaf edge to the root edge
	// of the block level challenge, meaning it should also include claim id links.
	royalBranches := make([][]protocol.ReadOnlyEdge, 0)
	for _, edge := range royalTreeLeaves {
		branch := []protocol.ReadOnlyEdge{edge}
		ancestors, err2 := retry.UntilSucceeds(ctx, func() ([]protocol.ReadOnlyEdge, error) {
			return cc.reader.ComputeAncestors(
				ctx, challengedAssertionHash, edge.Id(),
			)
		})
		if err2 != nil {
			return err2
		}
		srvlog.Info("Got ancestors for branch with base", log.Ctx{
			"validatorName": cc.validatorName,
			"numAncestors":  len(ancestors),
			"baseEdge":      edge.Id(),
		})
		branch = append(branch, ancestors...)
		royalBranches = append(royalBranches, branch)
	}
	// For each branch, update the inherited timers onchain in a single transaction.
	for _, branch := range royalBranches {
		if _, err2 := retry.UntilSucceeds(ctx, func() (bool, error) {
			innerErr := cc.writer.MultiUpdateInheritedTimers(ctx, branch)
			return false, innerErr
		}); err2 != nil {
			return err2
		}
		// In each iteration, check if the root edge has a timer >= a challenge period
		rootTimer, err2 := retry.UntilSucceeds(ctx, func() (protocol.InheritedTimer, error) {
			return royalRootEdge.InheritedTimer(ctx)
		})
		if err2 != nil {
			return err2
		}
		srvlog.Info("Updated the inherited timer for branch with base", log.Ctx{
			"validatorName": cc.validatorName,
			"rootTimer":     rootTimer,
			"baseEdge":      branch[0].Id().Hash,
		})
		// If yes, we confirm the root edge and finish early.
		if uint64(rootTimer) >= challengePeriodBlocks {
			srvlog.Info("Branch was confirmable by time", log.Ctx{
				"validatorName": cc.validatorName,
				"rootTimer":     rootTimer,
				"baseEdge":      branch[0].Id().Hash,
			})
			_, err2 = retry.UntilSucceeds(ctx, func() (bool, error) {
				innerErr := royalRootEdge.ConfirmByTimer(ctx)
				return false, innerErr
			})
			return err2
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
		srvlog.Error("Onchain timer differed after confirmation job", log.Ctx{
			"validatorName":   cc.validatorName,
			"onchainTimer":    onchainInheritedTimer,
			"challengePeriod": challengePeriodBlocks,
		})
		return fmt.Errorf(
			"onchain timer %d after confirmation job was executed < challenge period %d",
			onchainInheritedTimer,
			challengePeriodBlocks,
		)
	}
	srvlog.Info("Confirming edge by time", log.Ctx{
		"validatorName":   cc.validatorName,
		"onchainTimer":    onchainInheritedTimer,
		"challengePeriod": challengePeriodBlocks,
	})
	if _, err = retry.UntilSucceeds(ctx, func() (bool, error) {
		innerErr := royalRootEdge.ConfirmByTimer(ctx)
		return false, innerErr
	}); err != nil {
		return err
	}
	srvlog.Info("Challenge completed", log.Ctx{
		"validatorName":   cc.validatorName,
		"onchainTimer":    onchainInheritedTimer,
		"challengePeriod": challengePeriodBlocks,
	})
	return nil
}
