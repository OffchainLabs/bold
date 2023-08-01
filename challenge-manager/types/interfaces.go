// Package types includes types and interfaces specific to the challenge manager instance.
package types

import (
	"context"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/rpc"
)

// ChallengeManager defines an offchain, challenge manager, which will be
// an active participant in interacting with the on-chain contracts.
type ChallengeManager interface {
	ChallengeCreator
	ChallengeReader
}

// ChallengeCreator defines a struct which can initiate a challenge on an assertion hash
// by creating a level zero, block challenge edge onchain.
type ChallengeCreator interface {
	ChallengeAssertion(ctx context.Context, id protocol.AssertionHash) error
}

// ChallengeReader defines a struct which can read the challenge of a challenge manager.
type ChallengeReader interface {
	Mode() Mode
	MaxDelaySeconds() int
}

// RPCClient defines a subset of the rpc.Client struct as needed by the challenge manager.
type RPCClient interface {
	BatchCallContext(ctx context.Context, b []rpc.BatchElem) error
}

// ChallengeManagerBackend defines a subset of the ethclient struct as needed by the challenge manager.
type ChallengeManagerBackend interface {
	bind.ContractBackend
	Client() RPCClient
}
