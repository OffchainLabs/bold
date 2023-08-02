// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE

// Package solimpl includes an easy-to-use abstraction
// around the challenge protocol contracts using their Go
// bindings and exposes minimal details of Ethereum's internals.
package solimpl

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/OffchainLabs/bold/solgen/go/bridgegen"
	"github.com/OffchainLabs/bold/solgen/go/rollupgen"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

var (
	ErrNotFound         = errors.New("item not found on-chain")
	ErrAlreadyExists    = errors.New("item already exists on-chain")
	ErrPrevDoesNotExist = errors.New("assertion predecessor does not exist")
	ErrTooLate          = errors.New("too late to create assertion sibling")
)

var assertionCreatedId common.Hash

func init() {
	rollupAbi, err := rollupgen.RollupCoreMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	assertionCreatedEvent, ok := rollupAbi.Events["AssertionCreated"]
	if !ok {
		panic("RollupCore ABI missing AssertionCreated event")
	}
	assertionCreatedId = assertionCreatedEvent.ID
}

// ChainBackend to interact with the underlying blockchain.
type ChainBackend interface {
	bind.ContractBackend
	ReceiptFetcher
}

// ChainCommitter defines a type of chain backend that supports
// committing changes via a direct method, such as a simulated backend
// for testing purposes.
type ChainCommitter interface {
	Commit() common.Hash
}

// ReceiptFetcher defines the ability to retrieve transactions receipts from the chain.
type ReceiptFetcher interface {
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
}

// AssertionChain is a wrapper around solgen bindings
// that implements the protocol interface.
type AssertionChain struct {
	backend    ChainBackend
	rollup     *rollupgen.RollupCore
	userLogic  *rollupgen.RollupUserLogic
	txOpts     *bind.TransactOpts
	rollupAddr common.Address
}

// NewAssertionChain instantiates an assertion chain
// instance from a chain backend and provided options.
func NewAssertionChain(
	_ context.Context,
	rollupAddr common.Address,
	txOpts *bind.TransactOpts,
	backend ChainBackend,
) (*AssertionChain, error) {
	chain := &AssertionChain{
		backend:    backend,
		txOpts:     txOpts,
		rollupAddr: rollupAddr,
	}
	coreBinding, err := rollupgen.NewRollupCore(
		rollupAddr, chain.backend,
	)
	if err != nil {
		return nil, err
	}
	assertionChainBinding, err := rollupgen.NewRollupUserLogic(
		rollupAddr, chain.backend,
	)
	if err != nil {
		return nil, err
	}
	chain.rollup = coreBinding
	chain.userLogic = assertionChainBinding
	return chain, nil
}

func (a *AssertionChain) GetAssertion(ctx context.Context, assertionHash protocol.AssertionHash) (protocol.Assertion, error) {
	var b [32]byte
	copy(b[:], assertionHash.Bytes())
	res, err := a.userLogic.GetAssertion(&bind.CallOpts{Context: ctx}, b)
	if err != nil {
		return nil, err
	}
	if res.Status == uint8(protocol.NoAssertion) {
		return nil, errors.Wrapf(
			ErrNotFound,
			"assertion with id %#x",
			assertionHash,
		)
	}
	return &Assertion{
		id:    assertionHash,
		chain: a,
	}, nil
}

func (a *AssertionChain) LatestConfirmed(ctx context.Context) (protocol.Assertion, error) {
	res, err := a.rollup.LatestConfirmed(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	return a.GetAssertion(ctx, protocol.AssertionHash{Hash: res})
}

// CreateAssertion makes an on-chain claim given a previous assertion hash, execution state,
// and a commitment to a post-state.
func (a *AssertionChain) CreateAssertion(
	ctx context.Context,
	parentAssertionCreationInfo *protocol.AssertionCreatedInfo,
	postState *protocol.ExecutionState,
) (protocol.Assertion, error) {
	if !parentAssertionCreationInfo.InboxMaxCount.IsUint64() {
		return nil, errors.New("prev assertion creation info inbox max count not a uint64")
	}
	newOpts := copyTxOpts(a.txOpts)
	if postState.GlobalState.Batch == 0 {
		return nil, errors.New("assertion post state cannot have a batch count of 0, as only genesis can")
	}
	bridgeAddr, err := a.userLogic.Bridge(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve bridge address for user rollup logic contract")
	}
	bridge, err := bridgegen.NewIBridgeCaller(bridgeAddr, a.backend)
	if err != nil {
		return nil, errors.Wrapf(err, "could not initialize bridge at address %#x", bridgeAddr)
	}
	inboxBatchAcc, err := bridge.SequencerInboxAccs(
		&bind.CallOpts{Context: ctx},
		new(big.Int).SetUint64(postState.GlobalState.Batch-1),
	)
	if err != nil {
		return nil, errors.Wrapf(err, "could not get sequencer inbox accummulator at batch %d", postState.GlobalState.Batch-1)
	}

	computedHash, err := a.userLogic.RollupUserLogicCaller.ComputeAssertionHash(
		&bind.CallOpts{Context: ctx},
		parentAssertionCreationInfo.AssertionHash,
		postState.AsSolidityStruct(),
		inboxBatchAcc,
	)
	if err != nil {
		return nil, errors.Wrap(err, "could not compute assertion hash")
	}
	existingAssertion, err := a.GetAssertion(ctx, protocol.AssertionHash{Hash: computedHash})
	switch {
	case err == nil:
		return existingAssertion, nil
	case !errors.Is(err, ErrNotFound):
		return nil, errors.Wrapf(err, "could not fetch assertion with computed hash %#x", computedHash)
	default:
	}

	receipt, err := transact(ctx, a.backend, func() (*types.Transaction, error) {
		return a.userLogic.NewStakeOnNewAssertion(
			newOpts,
			parentAssertionCreationInfo.RequiredStake,
			rollupgen.AssertionInputs{
				BeforeStateData: rollupgen.BeforeStateData{
					PrevPrevAssertionHash: parentAssertionCreationInfo.ParentAssertionHash,
					SequencerBatchAcc:     parentAssertionCreationInfo.AfterInboxBatchAcc,
					ConfigData: rollupgen.ConfigData{
						RequiredStake:       parentAssertionCreationInfo.RequiredStake,
						ChallengeManager:    parentAssertionCreationInfo.ChallengeManager,
						ConfirmPeriodBlocks: parentAssertionCreationInfo.ConfirmPeriodBlocks,
						WasmModuleRoot:      parentAssertionCreationInfo.WasmModuleRoot,
						NextInboxPosition:   parentAssertionCreationInfo.InboxMaxCount.Uint64(),
					},
				},
				BeforeState: parentAssertionCreationInfo.AfterState,
				AfterState:  postState.AsSolidityStruct(),
			},
			computedHash,
		)
	})
	if createErr := handleCreateAssertionError(err, postState.GlobalState.BlockHash); createErr != nil {
		return nil, fmt.Errorf("failed to create assertion: %w", createErr)
	}
	if len(receipt.Logs) == 0 {
		return nil, errors.New("no logs observed from assertion creation")
	}
	var assertionCreated *rollupgen.RollupCoreAssertionCreated
	var found bool
	for _, log := range receipt.Logs {
		creationEvent, err := a.rollup.ParseAssertionCreated(*log)
		if err == nil {
			assertionCreated = creationEvent
			found = true
			break
		}
	}
	if !found {
		return nil, errors.New("could not find assertion created event in logs")
	}
	return a.GetAssertion(ctx, protocol.AssertionHash{Hash: assertionCreated.AssertionHash})
}

func (a *AssertionChain) GenesisAssertionHash(ctx context.Context) (common.Hash, error) {
	return a.userLogic.GenesisAssertionHash(&bind.CallOpts{Context: ctx})
}

// ConfirmAssertionByChallengeWinner attempts to confirm an assertion onchain
// if there is a winning, level zero, block challenge edge that claims it.
func (a *AssertionChain) ConfirmAssertionByChallengeWinner(
	ctx context.Context,
	assertionHash protocol.AssertionHash,
	winningEdgeId protocol.EdgeId,
) error {
	var b [32]byte
	copy(b[:], assertionHash.Bytes())
	node, err := a.userLogic.GetAssertion(&bind.CallOpts{Context: ctx}, b)
	if err != nil {
		return err
	}
	if node.Status == uint8(protocol.AssertionConfirmed) {
		return nil
	}
	creationInfo, err := a.ReadAssertionCreationInfo(ctx, assertionHash)
	if err != nil {
		return err
	}
	// If the assertion is genesis, return nil.
	if creationInfo.ParentAssertionHash == [32]byte{} {
		return nil
	}
	prevCreationInfo, err := a.ReadAssertionCreationInfo(ctx, protocol.AssertionHash{Hash: creationInfo.ParentAssertionHash})
	if err != nil {
		return err
	}
	latestConfirmed, err := a.LatestConfirmed(ctx)
	if err != nil {
		return err
	}
	if creationInfo.ParentAssertionHash != latestConfirmed.Id().Hash {
		return fmt.Errorf(
			"parent id %#x is not the latest confirmed assertion %#x",
			creationInfo.ParentAssertionHash,
			latestConfirmed.Id(),
		)
	}
	if !prevCreationInfo.InboxMaxCount.IsUint64() {
		return errors.New("assertion prev creation info inbox max count was not a uint64")
	}
	receipt, err := transact(ctx, a.backend, func() (*types.Transaction, error) {
		return a.userLogic.RollupUserLogicTransactor.ConfirmAssertion(
			copyTxOpts(a.txOpts),
			b,
			creationInfo.ParentAssertionHash,
			creationInfo.AfterState,
			winningEdgeId.Hash,
			rollupgen.ConfigData{
				WasmModuleRoot:      prevCreationInfo.WasmModuleRoot,
				ConfirmPeriodBlocks: prevCreationInfo.ConfirmPeriodBlocks,
				RequiredStake:       prevCreationInfo.RequiredStake,
				ChallengeManager:    prevCreationInfo.ChallengeManager,
				NextInboxPosition:   prevCreationInfo.InboxMaxCount.Uint64(),
			},
			creationInfo.AfterInboxBatchAcc,
		)
	})
	if err != nil {
		return err
	}
	if len(receipt.Logs) == 0 {
		return errors.New("no logs observed from assertion confirmation")
	}
	return nil
}

// SpecChallengeManager creates a new spec challenge manager
func (a *AssertionChain) SpecChallengeManager(ctx context.Context) (protocol.SpecChallengeManager, error) {
	challengeManagerAddr, err := a.userLogic.RollupUserLogicCaller.ChallengeManager(
		&bind.CallOpts{Context: ctx},
	)
	if err != nil {
		return nil, err
	}
	return NewSpecChallengeManager(
		ctx,
		challengeManagerAddr,
		a,
		a.backend,
		a.txOpts,
	)
}

// AssertionUnrivaledBlocks gets the number of blocks an assertion was unrivaled. That is, it looks up the
// assertion's parent, and from that parent, computes second_child_creation_block - first_child_creation_block.
// If an assertion is a second child, this function will return 0.
func (a *AssertionChain) AssertionUnrivaledBlocks(ctx context.Context, assertionHash protocol.AssertionHash) (uint64, error) {
	var b [32]byte
	copy(b[:], assertionHash.Bytes())
	wantNode, err := a.rollup.GetAssertion(&bind.CallOpts{Context: ctx}, b)
	if err != nil {
		return 0, err
	}
	if wantNode.Status == uint8(protocol.NoAssertion) {
		return 0, errors.Wrapf(
			ErrNotFound,
			"assertion with id %#x",
			assertionHash,
		)
	}
	// If the assertion requested is not the first child, it was never unrivaled.
	if !wantNode.IsFirstChild {
		return 0, nil
	}
	assertion := &Assertion{
		id:    assertionHash,
		chain: a,
	}
	prevId, err := assertion.PrevId(ctx)
	if err != nil {
		return 0, err
	}
	copy(b[:], prevId.Bytes())
	prevNode, err := a.rollup.GetAssertion(&bind.CallOpts{Context: ctx}, b)
	if err != nil {
		return 0, err
	}
	if prevNode.Status == uint8(protocol.NoAssertion) {
		return 0, errors.Wrapf(
			ErrNotFound,
			"assertion with id %#x",
			assertionHash,
		)
	}
	// If there is no second child, we simply return the number of blocks
	// since the assertion was created and its parent.
	if prevNode.SecondChildBlock == 0 {
		latestHeader, err := a.backend.HeaderByNumber(ctx, nil)
		if err != nil {
			return 0, err
		}
		if !latestHeader.Number.IsUint64() {
			return 0, errors.New("latest header number is not a uint64")
		}
		num := latestHeader.Number.Uint64()

		// Should never happen.
		if wantNode.CreatedAtBlock > num {
			return 0, fmt.Errorf(
				"assertion creation block %d > latest block number %d for assertion hash %#x",
				wantNode.CreatedAtBlock,
				num,
				assertionHash,
			)
		}
		return num - wantNode.CreatedAtBlock, nil
	}
	// Should never happen.
	if prevNode.FirstChildBlock > prevNode.SecondChildBlock {
		return 0, fmt.Errorf(
			"first child creation block %d > second child creation block %d for assertion hash %#x",
			prevNode.FirstChildBlock,
			prevNode.SecondChildBlock,
			prevId,
		)
	}
	return prevNode.SecondChildBlock - prevNode.FirstChildBlock, nil
}

func (a *AssertionChain) TopLevelAssertion(ctx context.Context, edgeId protocol.EdgeId) (protocol.AssertionHash, error) {
	cm, err := a.SpecChallengeManager(ctx)
	if err != nil {
		return protocol.AssertionHash{}, err
	}
	edgeOpt, err := cm.GetEdge(ctx, edgeId)
	if err != nil {
		return protocol.AssertionHash{}, err
	}
	if edgeOpt.IsNone() {
		return protocol.AssertionHash{}, errors.New("edge was nil")
	}
	return edgeOpt.Unwrap().AssertionHash(ctx)
}

func (a *AssertionChain) TopLevelClaimHeights(ctx context.Context, edgeId protocol.EdgeId) (protocol.OriginHeights, error) {
	cm, err := a.SpecChallengeManager(ctx)
	if err != nil {
		return protocol.OriginHeights{}, err
	}
	edgeOpt, err := cm.GetEdge(ctx, edgeId)
	if err != nil {
		return protocol.OriginHeights{}, err
	}
	if edgeOpt.IsNone() {
		return protocol.OriginHeights{}, errors.New("edge was nil")
	}
	edge := edgeOpt.Unwrap()
	return edge.TopLevelClaimHeight(ctx)
}

// LatestCreatedAssertion retrieves the latest assertion from the rollup contract by reading the
// latest confirmed assertion and then querying the contract log events for all assertions created
// since that block and returning the most recent one.
func (a *AssertionChain) LatestCreatedAssertion(ctx context.Context) (protocol.Assertion, error) {
	latestConfirmed, err := a.LatestConfirmed(ctx)
	if err != nil {
		return nil, err
	}
	createdAtBlock, err := latestConfirmed.CreatedAtBlock()
	if err != nil {
		return nil, err
	}
	var query = ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(createdAtBlock),
		ToBlock:   nil, // Latest block.
		Addresses: []common.Address{a.rollupAddr},
		Topics:    [][]common.Hash{{assertionCreatedId}},
	}
	logs, err := a.backend.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}

	// The logs are likely sorted by blockNumber, index, but we find the latest one, just in case,
	// while ignoring any removed logs from a reorged event.
	var latestBlockNumber uint64
	var latestLogIndex uint
	var latestLog *types.Log
	for _, log := range logs {
		l := log
		if l.Removed {
			continue
		}
		if l.BlockNumber > latestBlockNumber ||
			(l.BlockNumber == latestBlockNumber && l.Index >= latestLogIndex) {
			latestBlockNumber = l.BlockNumber
			latestLogIndex = l.Index
			latestLog = &l
		}
	}

	if latestLog == nil {
		return nil, errors.New("no assertion creation events found")
	}

	creationEvent, err := a.rollup.ParseAssertionCreated(*latestLog)
	if err != nil {
		return nil, err
	}
	return a.GetAssertion(ctx, protocol.AssertionHash{Hash: creationEvent.AssertionHash})
}

// ReadAssertionCreationInfo for an assertion sequence number by looking up its creation
// event from the rollup contracts.
func (a *AssertionChain) ReadAssertionCreationInfo(
	ctx context.Context, id protocol.AssertionHash,
) (*protocol.AssertionCreatedInfo, error) {
	var creationBlock uint64
	var topics [][]common.Hash
	if id == (protocol.AssertionHash{}) {
		rollupDeploymentBlock, err := a.rollup.RollupDeploymentBlock(&bind.CallOpts{Context: ctx})
		if err != nil {
			return nil, err
		}
		if !rollupDeploymentBlock.IsUint64() {
			return nil, errors.New("rollup deployment block was not a uint64")
		}
		creationBlock = rollupDeploymentBlock.Uint64()
		topics = [][]common.Hash{{assertionCreatedId}}
	} else {
		var b [32]byte
		copy(b[:], id.Bytes())
		node, err := a.rollup.GetAssertion(&bind.CallOpts{Context: ctx}, b)
		if err != nil {
			return nil, err
		}
		creationBlock = node.CreatedAtBlock
		topics = [][]common.Hash{{assertionCreatedId}, {id.Hash}}
	}
	var query = ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(creationBlock),
		ToBlock:   new(big.Int).SetUint64(creationBlock),
		Addresses: []common.Address{a.rollupAddr},
		Topics:    topics,
	}
	logs, err := a.backend.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}
	if len(logs) == 0 {
		return nil, errors.New("no assertion creation logs found")
	}
	if len(logs) > 1 {
		return nil, errors.New("found multiple instances of requested node")
	}
	ethLog := logs[0]
	parsedLog, err := a.rollup.ParseAssertionCreated(ethLog)
	if err != nil {
		return nil, err
	}
	afterState := parsedLog.Assertion.AfterState
	return &protocol.AssertionCreatedInfo{
		ConfirmPeriodBlocks: parsedLog.ConfirmPeriodBlocks,
		RequiredStake:       parsedLog.RequiredStake,
		ParentAssertionHash: parsedLog.ParentAssertionHash,
		BeforeState:         parsedLog.Assertion.BeforeState,
		AfterState:          afterState,
		InboxMaxCount:       parsedLog.InboxMaxCount,
		AfterInboxBatchAcc:  parsedLog.AfterInboxBatchAcc,
		AssertionHash:       parsedLog.AssertionHash,
		WasmModuleRoot:      parsedLog.WasmModuleRoot,
		ChallengeManager:    parsedLog.ChallengeManager,
	}, nil
}

func handleCreateAssertionError(err error, blockHash common.Hash) error {
	if err == nil {
		return nil
	}
	errS := err.Error()
	switch {
	case strings.Contains(errS, "Assertion already exists"):
		return errors.Wrapf(
			ErrAlreadyExists,
			"commit block hash %#x",
			blockHash,
		)
	case strings.Contains(errS, "Assertion does not exist"):
		return ErrPrevDoesNotExist
	case strings.Contains(errS, "Too late to create sibling"):
		return ErrTooLate
	default:
		return err
	}
}

// Runs a callback function meant to write to a chain backend, and if the
// chain backend supports committing directly, we call the commit function before
// returning. This function additionally waits for the transaction to complete and returns
// an optional transaction receipt. It returns an error if the
// transaction had a failed status on-chain, or if the execution of the callback
// failed directly.
func transact(ctx context.Context, backend ChainBackend, fn func() (*types.Transaction, error)) (*types.Receipt, error) {
	tx, err := fn()
	if err != nil {
		return nil, err
	}
	if commiter, ok := backend.(ChainCommitter); ok {
		commiter.Commit()
	}
	receipt, err := bind.WaitMined(ctx, backend, tx)
	if err != nil {
		return nil, err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		callMsg := ethereum.CallMsg{
			From:       common.Address{},
			To:         tx.To(),
			Gas:        0,
			GasPrice:   nil,
			Value:      tx.Value(),
			Data:       tx.Data(),
			AccessList: tx.AccessList(),
		}
		if _, err := backend.CallContract(ctx, callMsg, nil); err != nil {
			return nil, errors.Wrap(err, "failed transaction")
		}
	}
	return receipt, nil
}

// copyTxOpts creates a deep copy of the given transaction options.
func copyTxOpts(opts *bind.TransactOpts) *bind.TransactOpts {
	copied := &bind.TransactOpts{
		From:     opts.From,
		Context:  opts.Context,
		NoSend:   opts.NoSend,
		Signer:   opts.Signer,
		GasLimit: opts.GasLimit,
	}

	if opts.Nonce != nil {
		copied.Nonce = new(big.Int).Set(opts.Nonce)
	}
	if opts.Value != nil {
		copied.Value = new(big.Int).Set(opts.Value)
	}
	if opts.GasPrice != nil {
		copied.GasPrice = new(big.Int).Set(opts.GasPrice)
	}
	if opts.GasFeeCap != nil {
		copied.GasFeeCap = new(big.Int).Set(opts.GasFeeCap)
	}
	if opts.GasTipCap != nil {
		copied.GasTipCap = new(big.Int).Set(opts.GasTipCap)
	}
	return copied
}
