package statemanager

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/OffchainLabs/new-rollup-exploration/protocol"
	"github.com/OffchainLabs/new-rollup-exploration/util"
	"github.com/ethereum/go-ethereum/common"
)

type Manager interface {
	HistoryCommitmentAtHeight(ctx context.Context, height uint64) (common.Hash, error)
	SubscribeStateEvents(ctx context.Context) <-chan *StateAdvancedEvent
}

type Simulated struct {
	currentHeight *atomic.Uint64
	maxHeight     uint64
	lock          sync.RWMutex
	leaves        []common.Hash
	stateTree     util.MerkleExpansion
	l2BlockTimes  time.Duration
	feed          *protocol.EventFeed[*StateAdvancedEvent]
}

type StateAdvancedEvent struct {
	HistoryCommitment *util.HistoryCommitment
}

type Opt func(*Simulated)

func WithL2BlockTimes(d time.Duration) Opt {
	return func(s *Simulated) {
		s.l2BlockTimes = d
	}
}

func NewSimulatedManager(ctx context.Context, maxHeight uint64, leaves []common.Hash, opts ...Opt) *Simulated {
	s := &Simulated{
		maxHeight:     maxHeight,
		currentHeight: &atomic.Uint64{},
		leaves:        leaves,
		stateTree:     util.ExpansionFromLeaves(leaves),
		l2BlockTimes:  time.Second,
		feed:          protocol.NewEventFeed[*StateAdvancedEvent](ctx),
	}
	for _, o := range opts {
		o(s)
	}
	return s
}

// HistoryCommitmentAtHeight --
// TODO: Match up with the existing state manager methods to rewind state, for example, for
// easier integration into the Nitro codebase.
func (s *Simulated) HistoryCommitmentAtHeight(_ context.Context, height uint64) (common.Hash, error) {
	s.lock.RLock()
	if height >= uint64(len(s.leaves)) {
		return [32]byte{}, fmt.Errorf("height %d exceeds available states %d", height, len(s.leaves))
	}
	treeAtHeight := util.ExpansionFromLeaves(s.leaves[:height])
	s.lock.RUnlock()
	h := &util.HistoryCommitment{
		Height: height,
		Merkle: treeAtHeight.Root(),
	}
	return h.Hash(), nil
}

func (s *Simulated) SubscribeStateEvents(ctx context.Context) <-chan *StateAdvancedEvent {
	return s.feed.Subscribe(ctx)
}

func (s *Simulated) AdvanceL2Chain(ctx context.Context) {
	tick := time.NewTicker(s.l2BlockTimes)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			s.currentHeight.Add(1)
			height := s.currentHeight.Load()
			s.lock.RLock()
			treeAtHeight := util.ExpansionFromLeaves(s.leaves[:height])
			s.lock.RUnlock()
			s.feed.Append(&StateAdvancedEvent{
				HistoryCommitment: &util.HistoryCommitment{
					Height: height,
					Merkle: treeAtHeight.Root(),
				},
			})
		case <-ctx.Done():
			return
		}
	}
}
