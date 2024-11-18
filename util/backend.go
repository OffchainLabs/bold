package util

import (
	"context"
	"errors"
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

func (b BackendWrapper) HeaderNumberUint64(ctx context.Context, number *big.Int) (uint64, error) {
	header, err := b.ethClient.HeaderByNumber(ctx, number)
	if err != nil {
		return 0, err
	}
	if !header.Number.IsUint64() {
		return 0, errors.New("block number is not uint64")
	}
	return header.Number.Uint64(), nil
}
