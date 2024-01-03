// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE

package backend

import (
	"context"
	"testing"
	"time"

	retry "github.com/OffchainLabs/bold/runtime"
	"github.com/stretchr/testify/require"
)

func TestLocalAnvilLoadAccounts(t *testing.T) {
	a, err := NewAnvilLocal(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if err := a.loadAccounts(); err != nil {
		t.Fatal(err)
	}
	if len(a.accounts) == 0 {
		t.Error("No accounts generated")
	}
}

func TestLocalAnvilStarts(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	a, err := NewAnvilLocal(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if err := a.Start(ctx); err != nil {
		t.Fatal(err)
	}
	_, err = retry.UntilSucceeds(ctx, func() (bool, error) {
		if _, err := a.DeployRollup(ctx); err != nil {
			return false, err
		}
		return true, nil
	})
	require.NoError(t, err)

	// There should be at least 100 blocks
	bn, err2 := a.Client().HeaderByNumber(ctx, nil)
	if err2 != nil {
		t.Fatal(err2)
	}
	if bn.Number.Uint64() < 100 {
		t.Errorf("Expected at least 100 blocks at start, but got only %d", bn)
	}
}
