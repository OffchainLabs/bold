// Package l2stateprovider defines a dependency which provides L2 states and proofs
// needed for the challenge manager to interact with Arbitrum chains' rollup and challenge
// contracts.
//
// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE
package l2stateprovider

import (
	"context"
	"errors"
	"math/big"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/OffchainLabs/bold/containers/option"
	"github.com/OffchainLabs/bold/solgen/go/rollupgen"
	commitments "github.com/OffchainLabs/bold/state-commitments/history"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrNoExecutionState = errors.New("chain does not have execution state")
)

type ConfigSnapshot struct {
	RequiredStake           *big.Int
	ChallengeManagerAddress common.Address
	ConfirmPeriodBlocks     uint64
	WasmModuleRoot          [32]byte
	InboxMaxCount           *big.Int
}

type Provider interface {
	ExecutionProvider
	HistoryCommitmentProvider
	PrefixProver
	OneStepProofProvider
	HistoryChecker
}

func WithAllStateRoots() ClaimHeight {
	return ClaimHeight{
		From: 0,
		To:   option.None[uint64](),
	}
}

type MessageNumberRange struct {
	From uint64
	To   uint64
}

type ClaimHeight struct {
	From uint64
	To   option.Option[uint64]
}

type HistoryCommitmentProvider interface {
	HistoryCommitment(
		ctx context.Context,
		wasmModuleRoot common.Hash,
		messageNumberRange MessageNumberRange,
		claimHeights ...ClaimHeight,
	) (commitments.History, error)
}

type ExecutionProvider interface {
	// Produces the latest state to assert to L1 from the local state manager's perspective.
	ExecutionStateAtMessageNumber(ctx context.Context, messageNumber uint64) (*protocol.ExecutionState, error)
	// If the state manager locally has this execution state, returns its message count and no error.
	// Returns ErrChainCatchingUp if catching up to chain.
	// Returns ErrNoExecutionState if the state manager does not have this execution state.
	ExecutionStateMsgCount(ctx context.Context, state *protocol.ExecutionState) (uint64, error)
}

type PrefixProver interface {
	PrefixProof(
		ctx context.Context,
		wasmModuleRoot common.Hash,
		messageNumberRange MessageNumberRange,
		claimHeights ...ClaimHeight,
	) ([]byte, error)
}

type OneStepProofProvider interface {
	OneStepProofData(
		ctx context.Context,
		wasmModuleRoot common.Hash,
		postState rollupgen.ExecutionState,
		messageNumber,
		bigStep,
		smallStep uint64,
	) (data *protocol.OneStepData, startLeafInclusionProof, endLeafInclusionProof []common.Hash, err error)
}

type History struct {
	Height     uint64
	MerkleRoot common.Hash
}

type HistoryChecker interface {
	AgreesWithHistoryCommitment(
		ctx context.Context,
		wasmModuleRoot common.Hash,
		assertionInboxMaxCount uint64,
		parentAssertionAfterStateBatch uint64,
		challengeLevel protocol.ChallengeLevel,
		history History,
		claimHeights ...ClaimHeight,
	) (bool, error)
}
