package main

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/OffchainLabs/bold/solgen/go/rollupgen"
	"github.com/OffchainLabs/bold/util"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

func (s *service) scanChainEvents(ctx context.Context) {
	latestConfirmed, err := s.rollupBindings.LatestConfirmed(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	latestConfirmedAssertion, err := s.readAssertionCreationInfo(
		ctx,
		protocol.AssertionHash{Hash: latestConfirmed},
	)
	if err != nil {
		panic(err)
	}

	// Gather all challenged assertions.
	challengedAssertions := make([]*protocol.AssertionCreatedInfo, 0)
	_ = challengedAssertions

	fromBlock := latestConfirmedAssertion.CreationBlock
	ticker := time.NewTicker(s.chainScanInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			latestBlock, err := s.client.HeaderByNumber(ctx, util.GetSafeBlockNumber())
			if err != nil {
				log.Error("Could not get header by number", "err", err)
				continue
			}
			if !latestBlock.Number.IsUint64() {
				log.Error("Latest block number was not a uint64")
				continue
			}
			toBlock := latestBlock.Number.Uint64()
			if fromBlock == toBlock {
				continue
			}
			filterOpts := &bind.FilterOpts{
				Start:   fromBlock,
				End:     &toBlock,
				Context: ctx,
			}
			if err = s.processAllAssertionsInRange(ctx, s.rollupBindings, filterOpts); err != nil {
				log.Error("Could not process assertions in range", "err", err)
				continue
			}
			if err != nil {
				log.Error("Could not check for assertion added", "err", err)
				continue
			}

			// Assertion confirmation.
			it, err := s.rollupBindings.FilterAssertionConfirmed(filterOpts, nil)
			if err != nil {
				panic(err)
			}
			defer func() {
				if err = it.Close(); err != nil {
					log.Error("Could not close filter iterator", "err", err)
				}
			}()
			for it.Next() {
				if it.Error() != nil {
					panic(err)
				}
				s.confirmedAssertionsObserved <- it.Event.AssertionHash
			}

			// Scan for edge confirmations.
			it2, err := s.chalManager.FilterEdgeConfirmedByTime(filterOpts, nil, nil)
			if err != nil {
				panic(err)
			}
			defer func() {
				if err = it2.Close(); err != nil {
					log.Error("Could not close filter iterator", "err", err)
				}
			}()
			for it2.Next() {
				if it2.Error() != nil {
					panic(err)
				}
				fmt.Println("Edge confirmed by time")
				eg, err := s.chalManager.GetEdge(&bind.CallOpts{}, it2.Event.EdgeId)
				if err != nil {
					panic(err)
				}
				s.confirmedEdgesObserved <- &edge{
					id:            it2.Event.EdgeId,
					txHash:        it2.Event.Raw.TxHash,
					ChallengeEdge: &eg,
				}
			}

			it3, err := s.chalManager.FilterEdgeConfirmedByOneStepProof(filterOpts, nil, nil)
			if err != nil {
				panic(err)
			}
			defer func() {
				if err = it3.Close(); err != nil {
					log.Error("Could not close filter iterator", "err", err)
				}
			}()
			for it3.Next() {
				if it3.Error() != nil {
					panic(err)
				}
				fmt.Println("Edge confirmed by osp")
				eg, err := s.chalManager.GetEdge(&bind.CallOpts{}, it3.Event.EdgeId)
				if err != nil {
					panic(err)
				}
				s.confirmedEdgesObserved <- &edge{
					id:            it3.Event.EdgeId,
					txHash:        it3.Event.Raw.TxHash,
					ChallengeEdge: &eg,
				}
			}

			fromBlock = toBlock
		case <-ctx.Done():
			return
		}
	}
}

// This function will scan for all assertion creation events to determine which
// ones are canonical and which ones must be challenged.
func (s *service) processAllAssertionsInRange(
	ctx context.Context,
	filterer *rollupgen.RollupUserLogic,
	filterOpts *bind.FilterOpts,
) error {
	it, err := filterer.FilterAssertionCreated(filterOpts, nil, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err = it.Close(); err != nil {
			log.Error("Could not close filter iterator", "err", err)
		}
	}()
	for it.Next() {
		if it.Error() != nil {
			return it.Error()
		}
		hash := it.Event.AssertionHash
		fmt.Println("Got the creation event", common.Hash(hash).Hex())
		creationInfo, err := s.readAssertionCreationInfo(
			ctx, protocol.AssertionHash{Hash: hash},
		)
		if err != nil {
			return err
		}
		parentHash := creationInfo.ParentAssertionHash
		if parentHash == (common.Hash{}) {
			genesis, err := s.rollupBindings.GenesisAssertionHash(&bind.CallOpts{})
			if err != nil {
				return err
			}
			parentHash = genesis
		}
		parentAssertion, err := s.rollupBindings.GetAssertion(&bind.CallOpts{}, parentHash)
		if err != nil {
			return err
		}
		// If the parent has a second child, that means this assertion is either part of a challenge
		// or will be part of a challenge.
		if parentAssertion.SecondChildBlock > 0 {
			s.claimedChallengeAssertions <- creationInfo
		}
	}
	return nil
}

func (s *service) readAssertionCreationInfo(
	ctx context.Context,
	id protocol.AssertionHash,
) (*protocol.AssertionCreatedInfo, error) {
	var creationBlock uint64
	var topics [][]common.Hash
	if id == (protocol.AssertionHash{}) {
		rollupDeploymentBlock, err := s.rollupBindings.RollupDeploymentBlock(util.GetSafeCallOpts(&bind.CallOpts{Context: ctx}))
		if err != nil {
			return nil, err
		}
		if !rollupDeploymentBlock.IsUint64() {
			return nil, errors.New("rollup deployment block is not a uint64")
		}
		creationBlock = rollupDeploymentBlock.Uint64()
		topics = [][]common.Hash{{assertionCreatedId}}
	} else {
		var b [32]byte
		copy(b[:], id.Bytes())
		node, err := s.rollupBindings.GetAssertion(util.GetSafeCallOpts(&bind.CallOpts{Context: ctx}), b)
		if err != nil {
			return nil, err
		}
		creationBlock = node.CreatedAtBlock
		topics = [][]common.Hash{{assertionCreatedId}, {id.Hash}}
	}
	var query = ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(creationBlock),
		ToBlock:   new(big.Int).SetUint64(creationBlock),
		Addresses: []common.Address{s.rollupAddr},
		Topics:    topics,
	}
	logs, err := s.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}
	if len(logs) != 1 {
		return nil, errors.New("expected one log")
	}
	ethLog := logs[0]
	parsedLog, err := s.rollupBindings.ParseAssertionCreated(ethLog)
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
		TransactionHash:     ethLog.TxHash,
		CreationBlock:       ethLog.BlockNumber,
	}, nil
}
