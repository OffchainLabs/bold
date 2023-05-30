package validator

import (
	"context"
	"fmt"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/solgen/go/rollupgen"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
)

// Sync edges from challenges from confirmed block height to latest block height.
// - Get all edges from watcher (retry on fail)
// - Build edge trackers for every edge (retry on fail)
// - Spin of all the edge trackers as part of go routine.
// nolint:unused
func (v *Validator) syncEdges(ctx context.Context) error {
	if err := v.waitForSync(v.initialSyncComplete); err != nil {
		return err
	}
	edges, err := v.watcher.GetEdges()
	if err != nil {
		return err
	}

	trackers, err := v.getEdgeTrackers(ctx, edges)
	if err != nil {
		return err
	}

	// Spin off all the edge trackers as part of go routine.
	for _, tracker := range trackers {
		go tracker.spawn(ctx)
	}

	return nil
}

// nolint:unused
func (v *Validator) getExecutionStateBlockHeight(ctx context.Context, st rollupgen.ExecutionState) (uint64, error) {
	height, ok := v.stateManager.ExecutionStateBlockHeight(ctx, protocol.GoExecutionStateFromSolidity(st))
	if !ok {
		return 0, fmt.Errorf("missing previous assertion after execution %+v in local state manager", st)
	}
	return height, nil
}

// getEdgeTrackers builds edge trackers for every edge.
// If fails on getting assertion number or creation info, it'll retry until it succeeds.
// nolint:unused
func (v *Validator) getEdgeTrackers(ctx context.Context, edges []protocol.SpecEdge) ([]*edgeTracker, error) {
	var assertionIdMap = make(map[protocol.AssertionId][2]uint64)
	edgeTrackers := make([]*edgeTracker, len(edges))
	var err error
	var assertionId protocol.AssertionId
	for i, edge := range edges {
		// Retry until you get the previous assertion ID.
		assertionId, err = util.RetryUntilSucceeds(ctx, func() (protocol.AssertionId, error) {
			return edge.PrevAssertionId(ctx)
		})
		if err != nil {
			return nil, err
		}

		// Smart caching to avoid querying the same assertion number and creation info multiple times.
		// Edges in the same challenge should have the same creation info.
		cachedHeightAndInboxMsgCount, ok := assertionIdMap[assertionId]
		var assertionHeight uint64
		var inboxMsgCount uint64
		if !ok {
			// Retry until you get the assertion number.
			assertionNum, assertionErr := util.RetryUntilSucceeds(ctx, func() (protocol.AssertionSequenceNumber, error) {
				return v.chain.GetAssertionNum(ctx, assertionId)
			})
			if assertionErr != nil {
				return nil, assertionErr
			}

			// Retry until you get the assertion creation info.
			assertionCreationInfo, creationErr := util.RetryUntilSucceeds(ctx, func() (*protocol.AssertionCreatedInfo, error) {
				return v.chain.ReadAssertionCreationInfo(ctx, assertionNum)
			})
			if creationErr != nil {
				return nil, creationErr
			}

			// Retry until you get the execution state block height.
			height, heightErr := util.RetryUntilSucceeds(ctx, func() (uint64, error) {
				return v.getExecutionStateBlockHeight(ctx, assertionCreationInfo.AfterState)
			})
			if heightErr != nil {
				return nil, heightErr
			}
			assertionHeight = height
			inboxMsgCount = assertionCreationInfo.InboxMaxCount.Uint64()
			assertionIdMap[assertionId] = [2]uint64{assertionHeight, inboxMsgCount}
		} else {
			assertionHeight, inboxMsgCount = cachedHeightAndInboxMsgCount[0], cachedHeightAndInboxMsgCount[1]
		}
		edgeTrackers[i], err = newEdgeTracker(
			ctx,
			&edgeTrackerConfig{
				timeRef:          v.timeRef,
				actEveryNSeconds: v.edgeTrackerWakeInterval,
				chain:            v.chain,
				stateManager:     v.stateManager,
				validatorName:    v.name,
				validatorAddress: v.address,
			},
			edge,
			assertionHeight,
			inboxMsgCount,
		)
		if err != nil {
			log.WithError(err).Error("error creating edge tracker")
			continue
		}
	}
	return edgeTrackers, nil
}
