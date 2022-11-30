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
	HistoryCommitmentUpTo(ctx context.Context, height uint64) (util.HistoryCommitment, error)
	PrefixProof(ctx context.Context, from, to uint64) ([]common.Hash, error)
	HasHistoryCommitment(ctx context.Context, commitment util.HistoryCommitment) bool
	LatestHistoryCommitment(ctx context.Context) (util.HistoryCommitment, error)
}

type Simulated struct {
	stateRoots []common.Hash
}

func New(stateRoots []common.Hash) *Simulated {
	if len(stateRoots) == 0 {
		panic("must have state roots")
	}
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
		Height:    uint64(len(s.stateRoots)) - 1,
		StateRoot: s.stateRoots[len(s.stateRoots)-1],
	}, nil
}

func (s *Simulated) HistoryCommitmentUpTo(ctx context.Context, height uint64) (util.HistoryCommitment, error) {
	exp := util.ExpansionFromLeaves(s.stateRoots[:height])
	return util.HistoryCommitment{
		Height: height,
		Merkle: exp.Root(),
	}, nil
}

func (s *Simulated) PrefixProof(ctx context.Context, lo, hi uint64) ([]common.Hash, error) {
	exp := util.ExpansionFromLeaves(s.stateRoots[:lo])
	return util.GeneratePrefixProof(
		lo,
		exp,
		s.stateRoots[lo:hi],
	), nil
}

func (s *Simulated) HasHistoryCommitment(ctx context.Context, commitment util.HistoryCommitment) bool {
	if commitment.Height >= uint64(len(s.stateRoots)) {
		panic("commitment height out of range")
	}
	merkle := util.ExpansionFromLeaves(s.stateRoots[:commitment.Height]).Root()
	return merkle == commitment.Merkle
}

func (s *Simulated) LatestHistoryCommitment(ctx context.Context) (util.HistoryCommitment, error) {
	return util.HistoryCommitment{
		Height: uint64(len(s.stateRoots)) - 1,
		Merkle: util.ExpansionFromLeaves(s.stateRoots).Root(),
	}, nil
}
