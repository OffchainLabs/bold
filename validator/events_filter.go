package validator

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/solgen/go/challengeV2gen"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// syncEdges from last confirmed edge block to the latest block.
// - loop through the edge added events
// - get the edge id
// - get the prev assertion
// - construct the new edge tracker and start tracking it via a goroutine
func (v *Validator) syncEdges(ctx context.Context, a protocol.Assertion) error {
	latestConfirmed, err := v.chain.LatestConfirmed(ctx)
	if err != nil {
		return err
	}
	fromBlock, err := latestConfirmed.CreatedAtBlock()
	if err != nil {
		return err
	}

	latestBlock, err := v.backend.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}
	toBlock := latestBlock.Number.Uint64()

	filterOpts := &bind.FilterOpts{
		Start:   fromBlock,
		End:     &toBlock,
		Context: ctx,
	}

	cm, err := v.chain.SpecChallengeManager(ctx)
	if err != nil {
		return err
	}
	filterer, err := challengeV2gen.NewEdgeChallengeManagerFilterer(cm.Address(), v.backend)
	if err != nil {
		return err
	}
	it, err := filterer.FilterEdgeAdded(filterOpts, nil, nil, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err = it.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	for it.Next() {
		edgeAdded := it.Event
		edge, err := cm.GetEdge(ctx, edgeAdded.EdgeId)
		if err != nil {
			return err
		}
		id, err := edge.Unwrap().PrevAssertionId(ctx)
		if err != nil {
			return err
		}
		n, err := v.chain.GetAssertionNum(ctx, id)
		if err != nil {
			return err
		}
		prevCreationInfo, err := v.chain.ReadAssertionCreationInfo(ctx, n)
		if err != nil {
			return err
		}
		assertionPrevHeight, ok := v.stateManager.ExecutionStateBlockHeight(ctx, protocol.GoExecutionStateFromSolidity(prevCreationInfo.AfterState))
		if !ok {
			return fmt.Errorf("missing previous assertion %v after execution %+v in local state manager", n, prevCreationInfo.AfterState)
		}
		tracker, err := newEdgeTracker(
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
			assertionPrevHeight,
			prevCreationInfo.InboxMaxCount.Uint64(),
		)
		if err != nil {
			return err
		}
		go tracker.spawn(ctx)
	}
	return nil
}

func (v *Validator) pollForAssertions(ctx context.Context) {
	ticker := time.NewTicker(v.newAssertionCheckInterval)
	defer ticker.Stop()
	var nextAssertion protocol.AssertionSequenceNumber

	var onLeafCreatedLock sync.Mutex

	for {
		select {
		case <-ticker.C:
			onLeafCreatedLock.Lock()

			numberOfAssertions, err := v.chain.NumAssertions(ctx)
			if err != nil {
				log.Error(err)
				continue
			}
			latestConfirmedAssertion, err := v.chain.LatestConfirmed(ctx)
			if err != nil {
				log.Error(err)
				continue
			}
			latestConfirmedSeqNum := latestConfirmedAssertion.SeqNum()
			if latestConfirmedSeqNum >= nextAssertion {
				nextAssertion = latestConfirmedSeqNum + 1
			}

			for ; nextAssertion < protocol.AssertionSequenceNumber(numberOfAssertions); nextAssertion++ {
				assertion, err := v.chain.AssertionBySequenceNum(ctx, nextAssertion)
				if err != nil {
					log.Error(err)
					continue
				}
				if err := v.onLeafCreated(ctx, assertion); err != nil {
					log.Error(err)
				}
			}

			onLeafCreatedLock.Unlock()
		case <-ctx.Done():
			return
		}
	}
}
