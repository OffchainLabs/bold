package validator

import (
	"context"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/solgen/go/challengeV2gen"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// Sync edges from challenges from confirmed block height to latest block height.
// - Get all edges from challenges (retry on fail)
// - Build edge trackers for every edge (retry on fail)
// - Given block still advances while building all the edges trackers. At the end, it checks if it's on the latest block, or loop from the start
// - Once gathered all the sync edges from all the blocks, spin of all the edge trackers as part of go routine.
func (v *Validator) syncEdges(ctx context.Context) error {
	latestBlockNum, err := v.getLatestBlockNum(ctx)
	if err != nil {
		return err
	}

	currentBlockNum, err := v.getConfirmedBlockNum(ctx)
	if err != nil {
		return err
	}

	cm, err := v.chain.SpecChallengeManager(ctx)
	if err != nil {
		return err
	}
	filterer, err := challengeV2gen.NewEdgeChallengeManagerFilterer(cm.Address(), v.backend)
	if err != nil {
		return err
	}

	var edgeTrackers []*edgeTracker
	for {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		filterOpts := &bind.FilterOpts{
			Start:   currentBlockNum,
			End:     &latestBlockNum,
			Context: ctx,
		}

		// Retry if there's an error on filtering edge added events.
		it, err := filterer.FilterEdgeAdded(filterOpts, nil, nil, nil)
		if err != nil {
			log.WithError(err).Error("error filtering edge added events")
			continue
		}

		// Get all the edges.
		edges := v.getEdges(ctx, cm, it)

		// Build edge trackers for every edge.
		trackers := v.getEdgeTrackers(ctx, edges)
		edgeTrackers = append(edgeTrackers, trackers...)

		// latest block will keep advance. We shouldn't be done until we've processed all blocks.
		lbn, err := v.getLatestBlockNum(ctx)
		if err != nil {
			return err
		}
		if latestBlockNum == lbn {
			break
		}
		currentBlockNum = latestBlockNum
		latestBlockNum = lbn
	}

	// Spin off all the edge trackers as part of go routine.
	for _, tracker := range edgeTrackers {
		go tracker.spawn(ctx)
	}

	return nil
}

// get latest block number from the chain.
func (v *Validator) getLatestBlockNum(ctx context.Context) (uint64, error) {
	latestBlock, err := v.backend.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}
	return latestBlock.Number.Uint64(), nil
}

// get confirmed block number from the chain.
func (v *Validator) getConfirmedBlockNum(ctx context.Context) (uint64, error) {
	latestConfirmed, err := v.chain.LatestConfirmed(ctx)
	if err != nil {
		return 0, err
	}
	return latestConfirmed.CreatedAtBlock()
}

// getEdges gets all the edges from edge added events.
// If fails to get an edge given the edge ID, it'll retry until it succeeds.
func (v *Validator) getEdges(ctx context.Context, cm protocol.SpecChallengeManager, it *challengeV2gen.EdgeChallengeManagerEdgeAddedIterator) []util.Option[protocol.SpecEdge] {
	edges := make([]util.Option[protocol.SpecEdge], 0)
	for it.Next() {
		edgeAdded := it.Event
		e, err := cm.GetEdge(ctx, edgeAdded.EdgeId)
		if err != nil {
			log.WithError(err).Error("error getting edge") // Retry on error.
			continue
		}
		edges = append(edges, e)
	}
	return edges
}

// getEdgeTrackers builds edge trackers for every edge.
// If fails on getting assertion number or creation info, it'll retry until it succeeds.
func (v *Validator) getEdgeTrackers(ctx context.Context, edges []util.Option[protocol.SpecEdge]) []*edgeTracker {
	var assertionIdMap = make(map[protocol.AssertionId][2]uint64)
	edgeTrackers := make([]*edgeTracker, len(edges))
	for i, edge := range edges {
		assertionId, err := edge.Unwrap().PrevAssertionId(ctx)
		if err != nil {
			log.WithError(err).Error("error getting prev assertion id")
			continue
		}

		// Smart caching to avoid querying the same assertion number and creation info multiple times.
		// Edges in the same challenge should have the same creation info.
		cachedHeightAndInboxMsgCount, ok := assertionIdMap[assertionId]
		var assertionHeight uint64
		var inboxMsgCount uint64
		if !ok {
			n, err := v.chain.GetAssertionNum(ctx, assertionId)
			if err != nil {
				log.WithError(err).Error("error getting prev assertion id")
				continue
			}
			creationInfo, err := v.chain.ReadAssertionCreationInfo(ctx, n)
			if err != nil {
				log.WithError(err).Error("error getting creation info")
				continue
			}
			height, ok := v.stateManager.ExecutionStateBlockHeight(ctx, protocol.GoExecutionStateFromSolidity(creationInfo.AfterState))
			if !ok {
				log.Errorf("missing previous assertion %v after execution %+v in local state manager", n, creationInfo.AfterState)
				continue
			}
			assertionHeight = height
			inboxMsgCount = creationInfo.InboxMaxCount.Uint64()
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
			edge.Unwrap(),
			assertionHeight,
			inboxMsgCount,
		)
		if err != nil {
			log.WithError(err).Error("error creating edge tracker")
			continue
		}
	}
	return edgeTrackers
}
