package api

import (
	"context"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
)

type EdgesProvider interface {
	GetEdges() []protocol.SpecEdge
}

// TODO: Query logs from AssertionCreated events
type AssertionsProvider interface {
	ReadAssertionCreationInfo(context.Context, protocol.AssertionHash) (*protocol.AssertionCreatedInfo, error)
}
