package util

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"

	protocol "github.com/offchainlabs/bold/chain-abstraction"
)

type EthereumReader interface {
	ethereum.BlockNumberReader
	ethereum.ChainReader
	ethereum.ChainStateReader
	ethereum.ContractCaller
	ethereum.GasEstimator
	ethereum.GasPricer
	ethereum.GasPricer1559
	ethereum.FeeHistoryReader
	ethereum.LogFilterer
	ethereum.PendingStateReader
	ethereum.PendingContractCaller
	ethereum.ChainIDReader
	ethereum.TransactionReader
	CallContractAtHash(ctx context.Context, msg ethereum.CallMsg, blockHash common.Hash) ([]byte, error)
	TransactionSender(ctx context.Context, tx *types.Transaction, block common.Hash, index uint) (common.Address, error)
	Close()
	Client() rpc.ClientInterface
}

type EthereumWriter interface {
	EthereumReader
	ethereum.TransactionSender
	bind.ContractBackend
}

type EthereumReadWriter interface {
	EthereumReader
	EthereumWriter
}

var (
	_ protocol.ChainBackend = &BackendWrapper{
		desiredBlockNum: rpc.LatestBlockNumber,
	}
)

type BackendWrapper struct {
	EthereumReadWriter
	desiredBlockNum rpc.BlockNumber
}

func NewBackendWrapper(client EthereumReadWriter, desiredBlockNum rpc.BlockNumber) *BackendWrapper {
	return &BackendWrapper{client, desiredBlockNum}
}

func (b BackendWrapper) HeaderU64(ctx context.Context) (uint64, error) {
	header, err := b.HeaderByNumber(ctx, big.NewInt(int64(b.desiredBlockNum)))
	if err != nil {
		return 0, err
	}
	if !header.Number.IsUint64() {
		return 0, errors.New("block number is not uint64")
	}
	return header.Number.Uint64(), nil
}
