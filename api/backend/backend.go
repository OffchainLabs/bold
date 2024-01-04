// Package backend handles the business logic for API data fetching
// for BOLD challenge information. It is meant to be fairly abstract and
// well-tested.
package backend

import (
	"context"

	protocol "github.com/OffchainLabs/bold/chain-abstraction"
)

type Backend struct{}

type Opt func(b *Backend)

func (b *Backend) AllAssertions(
	ctx context.Context,
) ([]protocol.SpecEdge, error) {
	return nil, nil
}

func (b *Backend) FetchAssertion(ctx context.Context, assertionHash protocol.AssertionHash) ([]protocol.SpecEdge, error) {
	return nil, nil
}

func WithLimit(n uint64) Opt {
	return func(b *Backend) {}
}

func WithOffset(n uint64) Opt {
	return func(b *Backend) {}
}

func WithStartHistory(n uint64) Opt {
	return func(b *Backend) {}
}

func WithEndHistory(n uint64) Opt {
	return func(b *Backend) {}
}

func (b *Backend) AllEdges(ctx context.Context, opts ...Opt) ([]protocol.SpecEdge, error) {
	return nil, nil
}

func (b *Backend) FetchEdge(
	ctx context.Context, challengedAssertionHash protocol.AssertionHash,
) ([]protocol.SpecEdge, error) {
	return nil, nil
}

func (b *Backend) AllMiniStakes(ctx context.Context, opts ...Opt) ([]protocol.SpecEdge, error) {
	return nil, nil
}
