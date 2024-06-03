package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"time"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/OffchainLabs/bold/containers/option"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

var weiToGwei = big.NewInt(1000000000)

type gasPaymentRequest struct {
	confirmedAssertionInChallenge common.Hash
	essentialChallengeBranches    [][]*edge
}

type serviceFeePaymentRequest struct {
	claimedItemTyp claimType
	claimedItem    common.Hash
	edgeCreationTx common.Hash
}

type pendingGasPayment struct {
	PayableAddresses map[common.Address]*payableGasInfo `json:"payable_addresses"`
}

type payableGasInfo struct {
	EdgeCreationTotalCostGwei *big.Int `json:"creation_total_cost_gwei"`
	ConfirmationTotalCostGwei *big.Int `json:"confirmation_total_cost_gwei"`
}

type pendingServiceFeePayment struct {
	NumBlocksPayable             uint64         `json:"num_blocks_payable"`
	RecommendedYieldGweiPerBlock string         `json:"recommended_yield_gwei_per_block"`
	Staker                       common.Address `json:"staker"`
	ForStakeType                 claimType      `json:"for_stake_type"`
	StakedItemHash               common.Hash    `json:"staked_item_hash"`
	StakedItemChallengeLevel     string         `json:"staked_item_challenge_level"`
	StakedItemCreationTxHash     common.Hash    `json:"staked_item_creation_tx_hash"`
	StakedItemConfirmationTxHash common.Hash    `json:"staked_item_confirmation_tx_hash"`
}

// Will convert a payment request into a pending payment.
// TODO: Allow reimbursement tool to execute pending payments by itself
// through a signing wallet or keystore.
//
// Once a pending payment is created, it can notify the user via their preferred
// means and will output its data as a JSON file for processing.
func (s *service) processServiceFeePaymentRequest(ctx context.Context, request *serviceFeePaymentRequest) {
	// If the item is an assertion, its stake is locked up between its creation time
	// and min(firstChildCreationTime, confirmationTime).
	// Otherwise, if it is an edge, its stake is locked up from creation until confirmation.
	var numberOfBlocksToPay uint64
	var creationTxHash, confirmationTxHash common.Hash
	signer := types.NewCancunSigner(s.chainId)
	challengeLevel := option.None[protocol.ChallengeLevel]()
	if request.claimedItemTyp == assertionTyp {
		assertion, err := s.rollupBindings.GetAssertion(&bind.CallOpts{}, request.claimedItem)
		if err != nil {
			panic(err)
		}
		creationInfo, err := s.readAssertionCreationInfo(ctx, protocol.AssertionHash{Hash: request.claimedItem})
		if err != nil {
			panic(err)
		}
		confirmedAtBlock, confirmTxHash, err := s.fetchAssertionConfirmationBlock(ctx, creationInfo.CreationBlock, request.claimedItem)
		if err != nil {
			panic(err)
		}
		start := assertion.CreatedAtBlock
		end := assertion.FirstChildBlock
		if confirmedAtBlock < end || end == 0 {
			end = confirmedAtBlock
		}
		if start > end {
			panic("Invalid block range")
		}
		numberOfBlocksToPay = end - start
		creationTxHash = creationInfo.TransactionHash
		confirmationTxHash = confirmTxHash
	} else {
		eg, err := s.chalManager.GetEdge(&bind.CallOpts{}, request.claimedItem)
		if err != nil {
			panic(err)
		}
		numberOfBlocksToPay = eg.ConfirmedAtBlock - eg.CreatedAtBlock
		creationTxHash = request.edgeCreationTx
		_, confirmTx, err := s.fetchEdgeConfirmationBlockHeader(ctx, eg.ConfirmedAtBlock, &edge{ChallengeEdge: &eg, txHash: creationTxHash, id: request.claimedItem})
		if err != nil {
			panic(err)
		}
		challengeLevel = option.Some(protocol.ChallengeLevel(eg.Level))
		confirmationTxHash = confirmTx.Hash()
	}
	tx, _, err := s.client.TransactionByHash(ctx, creationTxHash)
	if err != nil {
		panic(err)
	}
	staker, err := types.Sender(signer, tx)
	if err != nil {
		panic(err)
	}
	pending := &pendingServiceFeePayment{
		NumBlocksPayable:             numberOfBlocksToPay,
		RecommendedYieldGweiPerBlock: s.recommendedL1YieldPerBlock.String(),
		Staker:                       staker,
		StakedItemCreationTxHash:     creationTxHash,
		StakedItemConfirmationTxHash: confirmationTxHash,
		StakedItemHash:               request.claimedItem,
		ForStakeType:                 request.claimedItemTyp,
	}
	if challengeLevel.IsSome() {
		pending.StakedItemChallengeLevel = challengeLevel.Unwrap().String()
	}
	fname := filepath.Join(s.pendingPaymentsOutputDir, fmt.Sprintf("pending_srv_fee_payment_%d_%s_%#x", time.Now().Unix(), request.claimedItemTyp, request.claimedItem))
	f, err := os.Create(fname)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			panic(err)
		}
	}()
	if err = json.NewEncoder(f).Encode(pending); err != nil {
		panic(err)
	}
	fmt.Println("Wrote pending service fee payment request", fname)
}

func (s *service) processGasPaymentRequest(ctx context.Context, request *gasPaymentRequest) {
	pending := &pendingGasPayment{
		PayableAddresses: make(map[common.Address]*payableGasInfo),
	}
	for _, branch := range request.essentialChallengeBranches {
		for _, eg := range branch {
			tx, _, err := s.client.TransactionByHash(ctx, eg.txHash)
			if err != nil {
				panic(err)
			}
			header, err := s.client.HeaderByNumber(ctx, big.NewInt(int64(eg.CreatedAtBlock)))
			if err != nil {
				panic(err)
			}
			creationCostGwei, confirmationCostGwei, err := s.determinePayableGasForEdge(ctx, eg, tx, header)
			if err != nil {
				panic(err)
			}
			signer := types.NewCancunSigner(s.chainId)
			sender, err := signer.Sender(tx)
			if err != nil {
				panic(err)
			}
			if _, ok := pending.PayableAddresses[sender]; !ok {
				pending.PayableAddresses[sender] = &payableGasInfo{
					EdgeCreationTotalCostGwei: new(big.Int),
					ConfirmationTotalCostGwei: new(big.Int),
				}
			}
			existingCreation := pending.PayableAddresses[sender].EdgeCreationTotalCostGwei
			existingConf := pending.PayableAddresses[sender].ConfirmationTotalCostGwei
			pending.PayableAddresses[sender].EdgeCreationTotalCostGwei = new(big.Int).Add(existingCreation, creationCostGwei)
			pending.PayableAddresses[sender].ConfirmationTotalCostGwei = new(big.Int).Add(existingConf, confirmationCostGwei)
		}
	}
	fname := filepath.Join(s.pendingPaymentsOutputDir, fmt.Sprintf("pending_gas_payment_%d_confirmed_claimed_assertion_%#x", time.Now().Unix(), request.confirmedAssertionInChallenge))
	f, err := os.Create(fname)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			panic(err)
		}
	}()
	if err = json.NewEncoder(f).Encode(pending); err != nil {
		panic(err)
	}
	fmt.Println("Wrote pending service fee payment request", fname)
}

// TODO: Is the basefee in gwei?
func (s *service) determinePayableGasForEdge(
	ctx context.Context, eg *edge, tx *types.Transaction, header *types.Header,
) (*big.Int, *big.Int, error) {
	// If the edge is a refinement move, we compute its creation cost
	// using the refinement predetermined gas costs.
	var determineConfirmationCost func() uint64
	var determineCreationCost func() uint64
	if eg.ClaimId != ([32]byte{}) {
		determineCreationCost = func() uint64 {
			switch eg.Level {
			case protocol.NewBlockChallengeLevel().Uint8():
				return blockRefinementCreateGasCost
			case protocol.NewBlockChallengeLevel().Uint8() + 1:
				return bigStepRefinementCreateGasCost
			case protocol.NewBlockChallengeLevel().Uint8() + 2:
				return smallStepRefinementCreateGasCost
			default:
				return 1 // TODO: Error instead.
			}
		}
		// Only edges with claim ids are confirmed by time, so we determine
		// how much that confirmation transaction cost.
		determineConfirmationCost = func() uint64 {
			return confirmEdgeByTimeGasCost
		}
	} else {
		// Otherwise, the edge was a bisection edge.
		determineCreationCost = func() uint64 {
			switch eg.Level {
			case protocol.NewBlockChallengeLevel().Uint8():
				return blockBisectGasCost
			case protocol.NewBlockChallengeLevel().Uint8() + 1:
				return bigStepBisectGasCost
			case protocol.NewBlockChallengeLevel().Uint8() + 2:
				return smallStepBisectGasCost
			default:
				return 1 // TODO: Error instead.
			}
		}
		// We check if the edge was one-step-proven, and if so, we check
		// how much that one step proof cost.
		determineConfirmationCost = func() uint64 {
			if isOneStepProven(eg) {
				return confirmByOneStepProofGasCost
			}
			return confirmEdgeByTimeGasCost
		}
	}
	// TODO: Use the min of this and the actual cost onchain.
	creationGas := new(big.Int).Mul(header.BaseFee, big.NewInt(int64(determineCreationCost())))
	creationReceipt, err := s.client.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return nil, nil, err
	}
	fmt.Printf("Creation block base fee %d\n", header.BaseFee.Uint64())
	fmt.Printf("Creation gas estimated %d vs. actual used %d\n", creationGas.Uint64(), new(big.Int).Mul(creationReceipt.EffectiveGasPrice, big.NewInt(int64(creationReceipt.GasUsed))))

	confirmationGasGwei := big.NewInt(0)
	if protocol.EdgeStatus(eg.Status) == protocol.EdgeConfirmed {
		confirmationBlockHeader, confirmationTx, err := s.fetchEdgeConfirmationBlockHeader(ctx, eg.ConfirmedAtBlock, eg)
		if err != nil {
			return nil, nil, err
		}
		confirmationReceipt, err := s.client.TransactionReceipt(ctx, confirmationTx.Hash())
		if err != nil {
			return nil, nil, err
		}
		confirmationGas := new(big.Int).Mul(confirmationBlockHeader.BaseFee, big.NewInt(int64(determineConfirmationCost())))
		fmt.Printf("Confirmation block base fee %d\n", confirmationBlockHeader.BaseFee.Uint64())
		fmt.Printf("Confirmation gas estimated %d vs. actual used %d\n", confirmationGas.Uint64(), new(big.Int).Mul(confirmationReceipt.EffectiveGasPrice, big.NewInt(int64(confirmationReceipt.GasUsed))))
		confirmationGasGwei = new(big.Int).Div(confirmationGas, weiToGwei)
	}
	creationGasGwei := new(big.Int).Div(creationGas, weiToGwei)
	return creationGasGwei, confirmationGasGwei, nil
}

func (s *service) fetchEdgeConfirmationBlockHeader(
	ctx context.Context, confirmedAtBlock uint64, eg *edge,
) (*types.Header, *types.Transaction, error) {
	var blockNum uint64
	var txHash common.Hash
	found := false
	if isOneStepProven(eg) {
		it, err := s.chalManager.FilterEdgeConfirmedByOneStepProof(&bind.FilterOpts{
			Start:   confirmedAtBlock,
			End:     &confirmedAtBlock,
			Context: ctx,
		}, [][32]byte{eg.id}, nil)
		defer func() {
			if err = it.Close(); err != nil {
				log.Error("Could not close filter iterator", "err", err)
			}
		}()
		for it.Next() {
			if it.Error() != nil {
				panic(err)
			}
			if it.Event.EdgeId == eg.id {
				blockNum = it.Event.Raw.BlockNumber
				txHash = it.Event.Raw.TxHash
				found = true
				break
			}
		}
	} else {
		it, err := s.chalManager.FilterEdgeConfirmedByTime(&bind.FilterOpts{
			Start:   confirmedAtBlock,
			End:     &confirmedAtBlock,
			Context: ctx,
		}, [][32]byte{eg.id}, nil)
		defer func() {
			if err = it.Close(); err != nil {
				log.Error("Could not close filter iterator", "err", err)
			}
		}()
		for it.Next() {
			if it.Error() != nil {
				panic(err)
			}
			if it.Event.EdgeId == eg.id {
				blockNum = it.Event.Raw.BlockNumber
				txHash = it.Event.Raw.TxHash
				found = true
				break
			}
		}
	}
	if !found {
		return nil, nil, errors.New("no edge confirmation tx found")
	}
	header, err := s.client.HeaderByNumber(ctx, big.NewInt(int64(blockNum)))
	if err != nil {
		return nil, nil, err
	}
	tx, _, err := s.client.TransactionByHash(ctx, txHash)
	if err != nil {
		return nil, nil, err
	}
	return header, tx, nil
}

func (s *service) fetchAssertionConfirmationBlock(ctx context.Context, creationBlock uint64, assertionHash common.Hash) (uint64, common.Hash, error) {
	it, err := s.rollupBindings.FilterAssertionConfirmed(&bind.FilterOpts{
		Start:   creationBlock,
		End:     nil,
		Context: ctx,
	}, [][32]byte{assertionHash})
	defer func() {
		if err = it.Close(); err != nil {
			log.Error("Could not close filter iterator", "err", err)
		}
	}()
	for it.Next() {
		if it.Error() != nil {
			panic(err)
		}
		if it.Event.AssertionHash == assertionHash {
			return it.Event.Raw.BlockNumber, it.Event.Raw.TxHash, nil
		}
	}
	return 0, common.Hash{}, fmt.Errorf("assertion %#x not yet confirmed", assertionHash)
}

func isOneStepProven(eg *edge) bool {
	isSmallStep := eg.Level == protocol.NewBlockChallengeLevel().Uint8()+2
	end := eg.EndHeight.Uint64()
	start := eg.StartHeight.Uint64()
	return isSmallStep && end-start == 1
}
