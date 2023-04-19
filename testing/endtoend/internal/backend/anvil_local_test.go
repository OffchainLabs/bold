package backend

import (
	"context"
	"testing"
	"time"
)

func TestLocalAnvilLoadAccounts(t *testing.T) {
	a, err := NewAnvilLocal(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if err := a.loadAccounts(); err != nil {
		t.Fatal(err)
	}
	if a.alice == nil {
		t.Error("Alice is nil")
	}
	if a.bob == nil {
		t.Error("Bob is nil")
	}
	if a.deployer == nil {
		t.Error("Deployer is nil")
	}
}

func TestLocalAnvilStarts(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancel()
	a, err := NewAnvilLocal(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if err := a.Start(); err != nil {
		t.Fatal(err)
	}
	if err := a.Stop(); err != nil {
		t.Fatal(err)
	}
}
