package statemanager

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/OffchainLabs/new-rollup-exploration/util"
	"github.com/ethereum/go-ethereum/common"
)

type Manager interface {
	HistoryCommitmentAtHeight(ctx context.Context, height uint64) common.Hash
}

type Simulated struct {
	height       *atomic.Uint64
	lock         sync.RWMutex
	stateTree    util.MerkleExpansion
	l2BlockTimes time.Duration
}

type Opt func(*Simulated)

func WithL2BlockTimes(d time.Duration) Opt {
	return func(s *Simulated) {
		s.l2BlockTimes = d
	}
}

func NewSimulatedManager(opts ...Opt) *Simulated {
	s := &Simulated{
		height:       &atomic.Uint64{},
		stateTree:    util.NewEmptyMerkleExpansion(),
		l2BlockTimes: time.Second, // By default, produce an L2 block every second.
	}
	for _, o := range opts {
		o(s)
	}
	return s
}

func (s *Simulated) AdvanceL2State(ctx context.Context) {
	tick := time.NewTicker(s.l2BlockTimes)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			s.height.Add(1)
			// TODO: Compute a new Merkle root.
		case <-ctx.Done():
			return
		}
	}
}

func (s *Simulated) HistoryCommitmentAtHeight(_ context.Context, height uint64) common.Hash {
	s.lock.RLock()
	commit := util.HistoryCommitment{
		Height: height,
		Merkle: s.stateTree.Root(),
	}
	s.lock.RUnlock()
	return commit.Hash()
}
