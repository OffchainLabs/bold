package solimpl

import (
	"context"
	"sync"
	"time"

	"github.com/OffchainLabs/bold/containers/threadsafe"
	retry "github.com/OffchainLabs/bold/runtime"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

type nonceManager struct {
	sync.Mutex
	senderAddress          common.Address
	client                 ChainBackend
	queue                  *threadsafe.Queue[txRequest]
	processRequestInterval time.Duration
	transactor             Transactor
}

func newNonceManager(
	senderAddress common.Address,
	processRequestInterval time.Duration,
	transactor Transactor,
) *nonceManager {
	return &nonceManager{
		queue:                  threadsafe.NewQueue[txRequest](),
		processRequestInterval: processRequestInterval,
		senderAddress:          senderAddress,
		transactor:             transactor,
	}
}

func (n *nonceManager) start(ctx context.Context) {
	ticker := time.NewTicker(n.processRequestInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			n.submitQueuedTxs(ctx)
		case <-ctx.Done():
			return
		}
	}
}

func (n *nonceManager) push(req txRequest) {
	n.Lock()
	defer n.Unlock()
	n.queue.Push(req)
}

func (n *nonceManager) submitQueuedTxs(ctx context.Context) {
	n.Lock()
	defer n.Unlock()
	if n.queue.Len() == 0 {
		return
	}
	pendingNonce, err := retry.UntilSucceeds[uint64](ctx, func() (uint64, error) {
		return n.client.PendingNonceAt(ctx, n.senderAddress)
	})
	if err != nil {
		srvlog.Error("Could not get pending nonce", log.Ctx{"error": err})
		return
	}
	nonce := pendingNonce
	srvlog.Info("Submitting queued txs", log.Ctx{
		"numTxs":       n.queue.Len(),
		"pendingNonce": pendingNonce,
	})
	start := time.Now()
	for n.queue.Len() != 0 {
		opt := n.queue.Pop()
		if opt.IsNone() {
			break
		}
		item := opt.Unwrap()
		_, err := n.transactor.PostSimpleTransaction(ctx, nonce, n.senderAddress, item.tx.Data(), item.gas, item.tx.Value())
		if err != nil {
			// TODO: Reinsert into the queue for attempting?
			srvlog.Error("Could not post tx", log.Ctx{"error": err, "hash": item.tx.Hash()})
			continue
		}
		nonce += 1
	}
	srvlog.Info("Finished submitting all queued txs", log.Ctx{
		"numTxs":      n.queue.Len(),
		"timeElapsed": time.Since(start),
	})
}
