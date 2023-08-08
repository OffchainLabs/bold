package assertions

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/OffchainLabs/bold/api"
	"github.com/OffchainLabs/bold/cmd/ctl/internal/data"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rivo/tview"
)

const defaultText = `Select an assertion

Use the arrow keys to navigate the tree.
Press Enter to select an assertion.

The tree is sorted by CreationBlock descending. 
i.e. the most recent assertions are at the top.
`

// AssertionTreePage builds and renders the assertion tree page.
// TODO: Assertions are better represented as a DAG. Consider implementing a DAG rendering view like renderdag in Sapling.
// TODO: See: https://github.com/facebook/sapling/tree/64374229f4fb6d613c9558d90b2f669c76753fe8/eden/scm/lib/renderdag
func AssertionTreePage() (title string, content tview.Primitive) {
	assertions, err := data.LoadAssertionsFromDisk()
	if err != nil {
		panic(err)
	}

	m := AssertionsToMap(assertions)
	edgeView := tview.NewTextView().
		SetWrap(true).
		SetText(defaultText)
	edgeView.SetBorderPadding(1, 1, 2, 0)
	setEdgeViewContents := func(e *api.Assertion) {
		b, err := json.MarshalIndent(e, "", "  ")
		if err != nil {
			edgeView.SetText(fmt.Sprintf("ERROR: Failed to marshal edge: %v", err))
		} else {
			edgeView.SetText(fmt.Sprintf("%s", b))
		}
	}

	tree := makeTreeView(m, setEdgeViewContents)

	tree.GetRoot().SetSelectedFunc(func() {
		edgeView.SetText(defaultText)
	})

	mainView := tview.NewFlex().
		AddItem(tree, 0, 1, true).
		AddItem(edgeView, 0, 1, false)

	footer := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().SetText("TODO: Filter(s), sort, refresh data, etc"), 1, 0, false)

	footer.SetBorder(true)

	return "Assertion Tree", tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(mainView, 0, 1, true).
		AddItem(footer, 6, 0, true)
}

func makeTreeView(m map[common.Hash]*api.Assertion, setEdgeView func(*api.Assertion)) *tview.TreeView {
	root := tview.NewTreeNode("Assertions")

	keys := make([]common.Hash, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]].CreationBlock > m[keys[j]].CreationBlock
	})

	for _, k := range keys {
		v := m[k]
		addChildren(root, v, m, setEdgeView)
	}

	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)
	tree.SetBorder(true)
	return tree
}

func addChildren(node *tview.TreeNode, assertion *api.Assertion, m map[common.Hash]*api.Assertion, selectedCB func(*api.Assertion)) {
	me := tview.NewTreeNode(assertionToTreeNodeName(assertion)).
		SetSelectedFunc(func() {
			selectedCB(assertion)
		})

	node.AddChild(me)
}

func AssertionsToMap(assertions []*api.Assertion) map[common.Hash]*api.Assertion {
	m := make(map[common.Hash]*api.Assertion, len(assertions))
	for _, a := range assertions {
		m[a.AssertionHash] = a
	}
	return m
}

func assertionToTreeNodeName(e *api.Assertion) string {
	// TODO: Define name format.

	trunc := func(b []byte) []byte {
		return b[:4]
	}

	return fmt.Sprintf(
		"%#x",
		trunc(e.AssertionHash[:]),
	)
}
