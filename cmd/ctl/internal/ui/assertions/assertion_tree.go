package assertions

import (
	"github.com/OffchainLabs/bold/api"
	"github.com/OffchainLabs/bold/cmd/ctl/internal/data"
	"github.com/OffchainLabs/bold/cmd/ctl/internal/ui/dag"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rivo/tview"
)

const defaultText = `Select an assertion

Use the arrow keys to navigate the tree.
Press Enter to select an assertion.

The tree is sorted by CreationBlock descending. 
i.e. the most recent assertions are at the top.

TODO: Make the assertions tree interactive.
`

// AssertionTreePage builds and renders the assertion tree page.
// TODO: Assertions are better represented as a DAG. Consider implementing a DAG rendering view like renderdag in Sapling.
// TODO: See: https://github.com/facebook/sapling/tree/64374229f4fb6d613c9558d90b2f669c76753fe8/eden/scm/lib/renderdag
func AssertionTreePage() (title string, content tview.Primitive) {
	assertions, err := data.LoadAssertionsFromDisk()
	if err != nil {
		panic(err)
	}

	d := dag.DAG[*AssertionItem]{}
	d.Init()

	for _, a := range assertions {
		d.AddNode(dag.Node[*AssertionItem]{
			Item:     &AssertionItem{a},
			ID:       a.AssertionHash,
			ParentID: a.ParentAssertionHash,
		})
	}

	renderedDAG := d.RenderString()

	tree := tview.NewTextView().
		SetWrap(true).
		SetText(renderedDAG)
	tree.SetBorder(true)

	details := tview.NewTextView().
		SetWrap(true).
		SetText(defaultText)
	details.SetBorderPadding(1, 1, 2, 0)

	mainView := tview.NewFlex().
		AddItem(tree, 0, 1, true).
		AddItem(details, 0, 1, false)

	footer := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().SetText("TODO: Filter(s), sort, refresh data, interactivity, etc"), 1, 0, false)

	footer.SetBorder(true)

	return "Assertion Tree", tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(mainView, 0, 1, true).
		AddItem(footer, 6, 0, true)
}

var _ = dag.Item(&AssertionItem{})

type AssertionItem struct {
	*api.Assertion
}

func (a *AssertionItem) DisplayID() string {
	return a.ID().Hex()
}

func (a *AssertionItem) Description() string {
	return "" // TODO: Provide description.
}

func (a *AssertionItem) Timestamp() int64 {
	return int64(a.Assertion.CreationBlock)
}

func (a *AssertionItem) ID() common.Hash {
	if a == nil {
		return common.Hash{}
	}
	return a.AssertionHash
}
