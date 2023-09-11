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

// Batch index for an Arbitrum L2 state.
type Batch uint64

// Height for a BOLD history commitment.
type Height uint64

// OpcodeIndex within an Arbitrator machine for an L2 message.
type OpcodeIndex uint64

// StepSize is the number of opcode increments used for stepping through
// machines for BOLD challenges.
type StepSize uint64

// ConfigSnapshot for an assertion on Arbitrum.
type ConfigSnapshot struct {
	RequiredStake           *big.Int
	ChallengeManagerAddress common.Address
	ConfirmPeriodBlocks     uint64
	WasmModuleRoot          [32]byte
	InboxMaxCount           *big.Int
}

// Provider defines an L2 state backend that can provide history commitments, execution
// states, prefix proofs, and more for the BOLD protocol.
type Provider interface {
	ExecutionProvider
	GeneralHistoryCommitter
	GeneralPrefixProver
	OneStepProofProvider
	HistoryChecker
}

type ExecutionProvider interface {
	// Produces the latest state to assert to L1 from the local state manager's perspective.
	ExecutionStateAtMessageNumber(ctx context.Context, messageNumber uint64) (*protocol.ExecutionState, error)
	// If the state manager locally has this execution state, returns its message count and no error.
	// Returns ErrChainCatchingUp if catching up to chain.
	// Returns ErrNoExecutionState if the state manager does not have this execution state.
	ExecutionStateMsgCount(ctx context.Context, state *protocol.ExecutionState) (uint64, error)
}

type GeneralHistoryCommitter interface {
	HistoryCommitment(
		ctx context.Context,
		wasmModuleRoot common.Hash,
		batch Batch,
		startHeights []Height,
		upToHeight option.Option[Height],
	) (commitments.History, error)
}

type GeneralPrefixProver interface {
	PrefixProof(
		ctx context.Context,
		wasmModuleRoot common.Hash,
		batch Batch,
		startHeights []Height,
		fromMessageNumber Height,
		upToHeight option.Option[Height],
	) ([]byte, error)
}

type OneStepProofProvider interface {
	OneStepProofData(
		ctx context.Context,
		wasmModuleRoot common.Hash,
		postState rollupgen.ExecutionState,
		startHeights []Height,
		upToHeight option.Option[Height],
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
		heights []Height,
		history History,
	) (bool, error)
}
