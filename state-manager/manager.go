package statemanager

import (
	"context"

	"github.com/OffchainLabs/new-rollup-exploration/protocol"
	"github.com/OffchainLabs/new-rollup-exploration/util"
	"github.com/ethereum/go-ethereum/common"
)

// Manager defines a struct that can provide local state data and historical
// Merkle commitments to L2 state for the validator.
type Manager interface {
	HasStateCommitment(ctx context.Context, commitment protocol.StateCommitment) bool
	StateCommitmentAtHeight(ctx context.Context, height uint64) (protocol.StateCommitment, error)
	LatestStateCommitment(ctx context.Context) (protocol.StateCommitment, error)
	HasHistoryCommitment(ctx context.Context, commitment util.HistoryCommitment) bool
	LatestHistoryCommitment(ctx context.Context) (util.HistoryCommitment, error)
	SubscribeStateEvents(ctx context.Context, ch chan<- *L2StateEvent)
}

type L2StateEvent struct{}

type Simulated struct {
	stateRoots []common.Hash
}

func New(stateRoots []common.Hash) *Simulated {
	return &Simulated{stateRoots}
}

func (s *Simulated) HasStateCommitment(ctx context.Context, commitment protocol.StateCommitment) bool {
	if commitment.Height >= uint64(len(s.stateRoots)) {
		panic("commitment height out of range")
	}
	return s.stateRoots[commitment.Height] == commitment.StateRoot
}

func (s *Simulated) StateCommitmentAtHeight(ctx context.Context, height uint64) (protocol.StateCommitment, error) {
	if height >= uint64(len(s.stateRoots)) {
		panic("commitment height out of range")
	}
	return protocol.StateCommitment{
		Height:    height,
		StateRoot: s.stateRoots[height],
	}, nil
}

func (s *Simulated) LatestStateCommitment(ctx context.Context) (protocol.StateCommitment, error) {
	return protocol.StateCommitment{
		Height:    uint64(len(s.stateRoots)),
		StateRoot: s.stateRoots[len(s.stateRoots)-1],
	}, nil
}

func (s *Simulated) HasHistoryCommitment(ctx context.Context, commitment util.HistoryCommitment) bool {
	if commitment.Height >= uint64(len(s.stateRoots)) {
		panic("commitment height out of range")
	}
	return false
}

func (s *Simulated) LatestHistoryCommitment(ctx context.Context) (util.HistoryCommitment, error) {
	panic("not implemented") // TODO: Implement
}

func (s *Simulated) SubscribeStateEvents(ctx context.Context, ch chan<- *L2StateEvent) {
	panic("not implemented") // TODO: Implement
}
