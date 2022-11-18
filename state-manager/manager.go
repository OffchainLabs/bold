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
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("prefix", "state-manager")

// Manager defines a struct that can provide local state data and historical
// Merkle commitments to L2 state for the validator.
type Manager interface {
	HasStateCommitment(ctx context.Context, commitment protocol.StateCommitment) bool
	StateCommitmentAtHeight(ctx context.Context, height uint64) (protocol.StateCommitment, error)
	LatestStateCommitment(ctx context.Context) (protocol.StateCommitment, error)
	SubscribeStateEvents(ctx context.Context, ch chan<- *L2StateEvent)
}

type L2StateEvent struct{}

type Simulated struct {
	currentHeight   *atomic.Uint64
	maxHeight       uint64
	lock            sync.RWMutex
	leaves          []common.Hash
	knownStateRoots map[common.Hash]bool
	stateTree       util.MerkleExpansion
	l2BlockTimes    time.Duration
	feed            *protocol.EventFeed[*L2StateEvent]
}

type Opt func(*Simulated)

func WithL2BlockTimes(d time.Duration) Opt {
	return func(s *Simulated) {
		s.l2BlockTimes = d
	}
}

func NewSimulatedManager(ctx context.Context, maxHeight uint64, leaves []common.Hash, opts ...Opt) *Simulated {
	s := &Simulated{
		maxHeight:       maxHeight,
		currentHeight:   &atomic.Uint64{},
		leaves:          leaves,
		stateTree:       util.ExpansionFromLeaves(leaves[:1]),
		l2BlockTimes:    time.Second,
		feed:            protocol.NewEventFeed[*L2StateEvent](ctx),
		knownStateRoots: make(map[common.Hash]bool),
	}
	for _, o := range opts {
		o(s)
	}
	return s
}

func (s *Simulated) SubscribeStateEvents(ctx context.Context, ch chan<- *L2StateEvent) {
	s.feed.Subscribe(ctx, ch)
}

func (s *Simulated) HasStateCommitment(ctx context.Context, commit protocol.StateCommitment) bool {
	// TODO: State commit is not the same as history commit! They are treated as the same for now
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.knownStateRoots[commit.StateRoot]
}

func (s *Simulated) LatestStateCommitment(_ context.Context) (protocol.StateCommitment, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return protocol.StateCommitment{
		Height:    s.currentHeight.Load(),
		StateRoot: s.stateTree.Root(),
	}, nil
}

func (s *Simulated) StateCommitmentAtHeight(_ context.Context, height uint64) (protocol.StateCommitment, error) {
	s.lock.RLock()
	if height >= uint64(len(s.leaves)) {
		return protocol.StateCommitment{
			Height: 0, StateRoot: common.Hash{},
		}, fmt.Errorf("height %d exceeds available states %d", height, len(s.leaves))
	}
	treeAtHeight := util.ExpansionFromLeaves(s.leaves[:height+1])
	s.lock.RUnlock()
	return protocol.StateCommitment{
		Height:    height,
		StateRoot: treeAtHeight.Root(),
	}, nil
}

func (s *Simulated) AdvanceL2Chain(ctx context.Context) {
	tick := time.NewTicker(s.l2BlockTimes)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			s.currentHeight.Add(1)
			height := s.currentHeight.Load()
			s.lock.Lock()
			s.stateTree = s.stateTree.AppendLeaf(s.leaves[height])
			s.feed.Append(&L2StateEvent{})
			s.knownStateRoots[s.stateTree.Root()] = true
			log.WithFields(logrus.Fields{
				"newHeight": height,
				"merkle":    util.FormatHash(s.stateTree.Root()),
			}).Debug("Advancing L2 chain state")
			s.lock.Unlock()
		case <-ctx.Done():
			return
		}
	}
}
