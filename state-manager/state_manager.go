package statemanager

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/OffchainLabs/new-rollup-exploration/util"
	"github.com/ethereum/go-ethereum/common"
)

type Manager interface {
	HistoryCommitmentAtHeight(ctx context.Context, height uint64) common.Hash
}

type Simulated struct {
	height    *atomic.Uint64
	lock      sync.RWMutex
	stateTree util.MerkleExpansion
}

func NewSimulatedManager() *Simulated {
	return &Simulated{
		height:    &atomic.Uint64{},
		stateTree: util.NewEmptyMerkleExpansion(),
	}
}

func (s *Simulated) Start() {}

func (s *Simulated) HistoryCommitmentAtHeight(_ context.Context, height uint64) common.Hash {
	s.lock.RLock()
	commit := util.HistoryCommitment{
		Height: height,
		Merkle: s.stateTree.Root(),
	}
	s.lock.RUnlock()
	return commit.Hash()
}
