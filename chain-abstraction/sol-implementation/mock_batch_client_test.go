package solimpl_test

import (
	"context"

	solimpl "github.com/OffchainLabs/bold/chain-abstraction/sol-implementation"
	"github.com/ethereum/go-ethereum/rpc"
)

var _ = solimpl.BatchClient(&MockBatchClient{})

type MockResult struct {
	Result interface{}
	Error  error
}

type MockBatchClient struct {
	Results []MockResult
	Error   error

	// Inputs represents the inputs that were passed to BatchCallContext for the most recent
	// BatchCallContext call.
	Inputs []rpc.BatchElem
}

func (m *MockBatchClient) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	if len(m.Results) != len(b) {
		panic("MockBatchClient: BatchCallContext: len(m.Results) != len(b)")
	}
	for i := 0; i < len(b); i++ {
		b[i].Result = m.Results[i].Result
		b[i].Error = m.Results[i].Error
	}
	m.Inputs = b
	return m.Error
}
