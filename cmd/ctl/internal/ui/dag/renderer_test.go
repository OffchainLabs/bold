package dag_test

import (
	"testing"

	"github.com/OffchainLabs/bold/cmd/ctl/internal/ui/dag"
)

func TestColumns_NewEmpty(t *testing.T) {
	cols := dag.Columns[*testItem]{}
	if i := cols.NewEmpty(); i != 0 {
		t.Fatalf("expected 0, got %d", i)
	}
	if len(cols) != 1 {
		t.Fatalf("expected 1, got %d", len(cols))
	}
}

func TestColumns_Reset(t *testing.T) {
	c := dag.Columns[*testItem]{}
	c.NewEmpty()
	if len(c) != 1 {
		t.Fatalf("expected 1, got %d", len(c))
	}
	c = dag.ResetColumns[*testItem](c)
	if len(c) != 0 {
		t.Fatalf("expected 0, got %d", len(c))
	}
}
