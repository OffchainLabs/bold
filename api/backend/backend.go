// Package backend handles the business logic for API data fetching
// for BOLD challenge information. It is meant to be fairly abstract and
// well-tested.
package backend

import (
	"context"

	"github.com/OffchainLabs/bold/api/db"
	protocol "github.com/OffchainLabs/bold/chain-abstraction"
)

type OnchainDataReader interface {
	Assertions()
}

type Database interface {
	Assertions()
}

type Backend struct {
	db db.ReadOnlyDatabase
}

type AssertionOpt func(b *Backend)

// - limit: the max number of items in the response
// - offset: the offset index in the DB
// - inbox_max_count: assertions that have a specified value for InboxMaxCount
// - from_block_number: items that were created since a specific block number. Defaults to latest confirmed assertion
// - to_block_number: caps the response to assertions up to and including a block number
func WithAssertionLimit(n uint64) AssertionOpt {
	return func(b *Backend) {}
}

func WithAssertionOffset(n uint64) AssertionOpt {
	return func(b *Backend) {}
}

func (b *Backend) AllAssertions(
	ctx context.Context,
	opts ...AssertionOpt,
) ([]protocol.SpecEdge, error) {
	return nil, nil
}

func (b *Backend) FetchAssertion(ctx context.Context, assertionHash protocol.AssertionHash) ([]protocol.SpecEdge, error) {
	return nil, nil
}

type EdgeOpt func(b *Backend)

func WithEdgeLimit(n uint64) EdgeOpt {
	return func(b *Backend) {}
}

func WithEdgeOffset(n uint64) EdgeOpt {
	return func(b *Backend) {}
}

func WithStartHistory(n uint64) EdgeOpt {
	return func(b *Backend) {}
}

func WithEndHistory(n uint64) EdgeOpt {
	return func(b *Backend) {}
}

func (b *Backend) AllEdges(ctx context.Context, opts ...EdgeOpt) ([]protocol.SpecEdge, error) {
	return nil, nil
}

func (b *Backend) FetchEdge(
	ctx context.Context, challengedAssertionHash protocol.AssertionHash,
) ([]protocol.SpecEdge, error) {
	return nil, nil
}

func (b *Backend) AllMiniStakes(ctx context.Context, opts ...EdgeOpt) ([]protocol.SpecEdge, error) {
	return nil, nil
}

func (b *Backend) FetchChallenge(ctx context.Context) error {
	return nil
}
