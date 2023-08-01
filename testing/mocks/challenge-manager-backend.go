package mocks

import (
	"context"
	"errors"

	"github.com/OffchainLabs/bold/challenge-manager/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/rpc"
)

type mockBackend struct {
	*backends.SimulatedBackend
}

func (mb *mockBackend) Client() types.RPCClient {
	return mb
}

func (mb *mockBackend) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	return errors.New("mockBackend.BatchCallContext not implemented")
}

func SimulatedBackendToChallangeManagerBackend(sb *backends.SimulatedBackend) types.ChallengeManagerBackend {
	return &mockBackend{
		sb,
	}
}
