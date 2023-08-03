package api_test

import (
	"context"
	"errors"

	"github.com/OffchainLabs/bold/api"
	protocol "github.com/OffchainLabs/bold/chain-abstraction"
)

var _ = api.EdgesProvider(&FakeEdgesProvider{})
var _ = api.AssertionsProvider(&FakeAssertionProvider{})

type FakeEdgesProvider struct {
	Edges []protocol.SpecEdge
}

func (f *FakeEdgesProvider) GetEdges() []protocol.SpecEdge {
	return f.Edges
}

type FakeAssertionProvider struct {
}

func (f *FakeAssertionProvider) ReadAssertionCreationInfo(ctx context.Context, ah protocol.AssertionHash) (*protocol.AssertionCreatedInfo, error) {
	return nil, errors.New("not implemented")
}

func (f *FakeAssertionProvider) LatestCreatedAssertionHashes(ctx context.Context) ([]protocol.AssertionHash, error) {
	return nil, errors.New("not implemented")
}
