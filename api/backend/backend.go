// Package backend handles the business logic for API data fetching
// for BOLD challenge information. It is meant to be fairly abstract and
// well-tested.
package backend

import (
	"context"

	"github.com/OffchainLabs/bold/api"
	"github.com/OffchainLabs/bold/api/db"
	protocol "github.com/OffchainLabs/bold/chain-abstraction"
	watcher "github.com/OffchainLabs/bold/challenge-manager/chain-watcher"
)

type Backend struct {
	db               db.ReadOnlyDatabase
	chainDataFetcher protocol.AssertionChain
	chainWatcher     *watcher.Watcher
}

func NewBackend(
	db db.ReadOnlyDatabase,
	chainDataFetcher protocol.AssertionChain,
	chainWatcher *watcher.Watcher,
) *Backend {
	return &Backend{
		db:               db,
		chainDataFetcher: chainDataFetcher,
		chainWatcher:     chainWatcher,
	}
}

func (b *Backend) GetAssertions(ctx context.Context, opts ...db.AssertionOption) ([]*api.JsonAssertion, error) {
	assertions, err := b.db.GetAssertions(opts...)
	if err != nil {
		return nil, err
	}
	// TODO: Fetch updated data about assertion statuses from the chain
	// and populate those fields in the response.
	return assertions, nil
}

func (b *Backend) GetEdges(ctx context.Context, opts ...db.EdgeOption) ([]*api.JsonEdge, error) {
	assertions, err := b.db.GetEdges(opts...)
	if err != nil {
		return nil, err
	}
	// TODO: Fetch updated data about edge statuses from the chain
	// and populate those fields in the response.
	return assertions, nil
}

func (b *Backend) GetMiniStakes(ctx context.Context, assertionHash protocol.AssertionHash, opts ...db.EdgeOption) ([]*api.JsonEdge, error) {
	return nil, nil
}

func (b *Backend) LatestConfirmedAssertion(ctx context.Context) (*api.JsonAssertion, error) {
	assertionDetails, err := b.chainDataFetcher.LatestConfirmed(ctx)
	if err != nil {
		return nil, err
	}
	_ = assertionDetails
	return &api.JsonAssertion{}, nil
}
