package validator

import (
	"context"
	"sync"
	"time"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
)

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
