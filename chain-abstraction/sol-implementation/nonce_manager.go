package solimpl

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/OffchainLabs/bold/containers/threadsafe"
	retry "github.com/OffchainLabs/bold/runtime"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

func init() {
	srvlog.SetHandler(log.StreamHandler(os.Stdout, log.LogfmtFormat()))
}

type nonceManager struct {
	sync.Mutex
	txOpts                 *bind.TransactOpts
	currNonce              uint64
	client                 ChainBackend
	queue                  *threadsafe.Queue[txRequest]
	processRequestInterval time.Duration
	transactor             Transactor
}

func newNonceManager(
	txOpts *bind.TransactOpts,
	processRequestInterval time.Duration,
	client ChainBackend,
	transactor Transactor,
) *nonceManager {
	return &nonceManager{
		txOpts:                 txOpts,
		client:                 client,
		queue:                  threadsafe.NewQueue[txRequest](),
		processRequestInterval: processRequestInterval,
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
		pending, err2 := n.client.PendingNonceAt(ctx, n.txOpts.From)
		if err2 != nil {
			srvlog.Error("Could not get pending nonce", log.Ctx{"error": err2, "validatorName": "honest"})
			return 0, err2
		}
		return pending, nil
	})
	if err != nil {
		srvlog.Error("Could not get pending nonce", log.Ctx{"error": err, "validatorName": "honest"})
		return
	}
	if n.currNonce < pendingNonce {
		n.currNonce = pendingNonce
	}
	srvlog.Info("Submitting queued txs", log.Ctx{
		"numTxs":        n.queue.Len(),
		"nonce":         n.currNonce,
		"validatorName": "honest",
	})
	start := time.Now()
	for n.queue.Len() != 0 {
		opt := n.queue.Pop()
		if opt.IsNone() {
			break
		}
		item := opt.Unwrap()
		txer, ok := n.transactor.(SimpleTransactor)
		if ok {
			newTx := &types.DynamicFeeTx{
				Nonce:      n.currNonce,
				ChainID:    item.tx.ChainId(),
				GasTipCap:  item.tx.GasTipCap(),
				GasFeeCap:  item.tx.GasFeeCap(),
				Gas:        item.tx.Gas(),
				To:         item.tx.To(),
				Value:      item.tx.Value(),
				Data:       item.tx.Data(),
				AccessList: item.tx.AccessList(),
			}
			tx := types.NewTx(newTx)
			signedTx, err := n.txOpts.Signer(n.txOpts.From, tx)
			if err != nil {
				srvlog.Error("Could not sign tx", log.Ctx{"error": err, "hash": item.tx.Hash(), "validatorName": "honest"})
				continue
			}
			if err := txer.SendTransaction(ctx, signedTx); err != nil {
				// TODO: Reinsert into the queue for attempting?
				srvlog.Error("Could not post tx", log.Ctx{"error": err, "hash": signedTx.Hash(), "validatorName": "honest"})
				continue
			}
			if commiter, ok := n.client.(ChainCommitter); ok {
				commiter.Commit()
			}
		} else {
			_, err := n.transactor.PostSimpleTransaction(ctx, n.currNonce, *item.tx.To(), item.tx.Data(), item.gas, item.tx.Value())
			if err != nil {
				// TODO: Reinsert into the queue for attempting?
				srvlog.Error("Could not post simple tx", log.Ctx{"error": err, "hash": item.tx.Hash(), "validatorName": "honest"})
				continue
			}
		}
		n.currNonce += 1
	}
	srvlog.Info("Finished submitting all queued txs", log.Ctx{
		"numTxs":        n.queue.Len(),
		"timeElapsed":   time.Since(start),
		"validatorName": "honest",
	})
}
