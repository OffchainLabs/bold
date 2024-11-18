package util

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	protocol "github.com/offchainlabs/bold/chain-abstraction"
	"math/big"
)

var (
	_ protocol.ChainBackend = &BackendWrapper{}
)

type ethClient = ethclient.Client
type BackendWrapper struct {
	*ethClient
}

func NewBackendWrapper(client *ethclient.Client) *BackendWrapper {
	return &BackendWrapper{client}
}

func (b BackendWrapper) HeaderByNumberIsUint64(ctx context.Context, number *big.Int) (*types.Header, error) {
	header, err := b.ethClient.HeaderByNumber(ctx, number)
	if err != nil {
		return nil, err
	}
	if !header.Number.IsUint64() {
		return nil, errors.New("block number is not uint64")
	}
	return header, nil
}
