package dag_test

import (
	"testing"

	"github.com/OffchainLabs/bold/cmd/ctl/internal/ui/dag"
	"github.com/ethereum/go-ethereum/common"
)

type testItem struct {
	identifier string
	desc       string
	timestamp  int64
}

func (t *testItem) DisplayID() string {
	return t.identifier
}

func (t *testItem) Description() string {
	return t.desc
}

func (t *testItem) Timestamp() int64 {
	return t.timestamp
}

func TestDagRenderString_Scenario1(t *testing.T) {

	want := `
◉  E
│  
│ ◉  D
~ │ 
  │ ◉  C
  │ │
  ◉ │  B
  ├─╯  
  ◉  A
  │  
  ~`

	nodes := []dag.Node[*testItem]{
		{
			Item: &testItem{
				identifier: "A",
				timestamp:  0,
			},
			ID:       common.HexToHash("0xA"),
			ParentID: common.HexToHash("0x0"),
		},
		{
			Item: &testItem{
				identifier: "B",
				timestamp:  1,
			},
			ID:       common.HexToHash("0xB"),
			ParentID: common.HexToHash("0xA"),
		},
		{
			Item: &testItem{
				identifier: "C",
				timestamp:  2,
			},
			ID:       common.HexToHash("0xC"),
			ParentID: common.HexToHash("0xA"),
		},
		{
			Item: &testItem{
				identifier: "D",
				timestamp:  3,
			},
			ID:       common.HexToHash("0xD"),
			ParentID: common.HexToHash("0xB"),
		},
		{
			Item: &testItem{
				identifier: "E",
				timestamp:  4,
			},
			ID:       common.HexToHash("0xE"),
			ParentID: common.HexToHash("0x0"),
		},
	}

	d := dag.DAG[*testItem]{}

	for _, n := range nodes {
		d.AddNode(n)
	}

	if got := d.RenderString(); got != want {
		t.Errorf("dag.RenderString() = %v, want %v", got, want)
	}
}
