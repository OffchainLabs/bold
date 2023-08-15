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

func (t *testItem) ID() common.Hash {
	if t == nil {
		return common.Hash{}
	}
	return common.BytesToHash([]byte(t.identifier))
}

func TestDagRenderString_ABC(t *testing.T) {

	want := `
◉  A
│ 
◉  B
│ 
◉  C
│
~
`[1:]

	nodes := []dag.Node[*testItem]{
		{
			Item: &testItem{
				identifier: "A",
				timestamp:  3,
			},
			ID:       common.HexToHash("0xA"),
			ParentID: common.HexToHash("0xB"),
		},
		{
			Item: &testItem{
				identifier: "B",
				timestamp:  2,
			},
			ID:       common.HexToHash("0xB"),
			ParentID: common.HexToHash("0xC"),
		},
		{
			Item: &testItem{
				identifier: "C",
				timestamp:  1,
			},
			ID:       common.HexToHash("0xC"),
			ParentID: common.HexToHash("0x0"),
		},
	}

	d := dag.DAG[*testItem]{}

	d.Init()

	for _, n := range nodes {
		d.AddNode(n)
	}

	if got := d.RenderString(); got != want {
		t.Errorf("dag.RenderString() = %v, want %v", got, want)
	}
}

func TestDagRenderString_MultiForkScenario(t *testing.T) {

	want := `
◉  G
│ 
│ ◉  E
│ │ 
│ │ ◉  D
│ ├─╯
◉ │  C
│ │ 
│ ◉  B
├─╯
◉  A
│
~
`[1:]

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
			ParentID: common.HexToHash("0xB"),
		},
		{
			Item: &testItem{
				identifier: "F",
				timestamp:  5,
			},
			ID:       common.HexToHash("0xF"),
			ParentID: common.HexToHash("0xB"),
		},
		{
			Item: &testItem{
				identifier: "G",
				timestamp:  6,
			},
			ID:       common.HexToHash("0xF"),
			ParentID: common.HexToHash("0xC"),
		},
	}

	d := dag.DAG[*testItem]{}

	d.Init()

	for _, n := range nodes {
		d.AddNode(n)
	}

	if got := d.RenderString(); got != want {
		t.Errorf("dag.RenderString() = \n%v, want \n%v", got, want)
	}
}
