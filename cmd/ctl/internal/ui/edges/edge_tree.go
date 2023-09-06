package edges

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/OffchainLabs/bold/api"
	"github.com/OffchainLabs/bold/cmd/ctl/internal/data"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const edgeTreeDefaultText = `Select an edge

Use the arrow keys to navigate the tree.
Press Enter to select an edge.

The key format is as follows:
ASSERTIONHASH-EDGETYPE-STARTCOMMITHEIGHT-STARTCOMMITROOT-ENDCOMMITHEIGHT-ENDCOMMITROOT

The tree is sorted by CreatedAtBlock descending. 
i.e. the most recent edges are at the top.
`

func EdgeTreePage() (title string, content tview.Primitive) {
	edges, err := data.LoadEdgesFromDisk()
	if err != nil {
		panic(err)
	}
	m := EdgesToMap(edges)

	edgeView := tview.NewTextView().
		SetWrap(true).
		SetText(edgeTreeDefaultText)
	edgeView.SetBorderPadding(1, 1, 2, 0)
	setEdgeViewContents := func(e *api.Edge) {
		b, err := json.MarshalIndent(e, "", "  ")
		if err != nil {
			edgeView.SetText(fmt.Sprintf("ERROR: Failed to marshal edge: %v", err))
		} else {
			edgeView.SetText(string(b))
		}
	}

	tree := makeTreeView(m, setEdgeViewContents)

	tree.GetRoot().SetSelectedFunc(func() {
		edgeView.SetText(edgeTreeDefaultText)
	})

	mainView := tview.NewFlex().
		AddItem(tree, 0, 1, true).
		AddItem(edgeView, 0, 1, false)

	footer := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().SetText("TODO: Filter(s), sort, refresh data, etc"), 1, 0, false)

	footer.SetBorder(true)

	return "Edge Tree", tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(mainView, 0, 1, true).
		AddItem(footer, 6, 0, true)
}

func makeTreeView(m map[common.Hash]*api.Edge, setEdgeView func(*api.Edge)) *tview.TreeView {
	root := tview.NewTreeNode("Edges")

	keys := make([]common.Hash, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]].CreatedAtBlock > m[keys[j]].CreatedAtBlock
	})

	for _, k := range keys {
		v := m[k]
		if v.IsRootChallenge() {
			addChildren(root, v, m, setEdgeView)
		}
	}

	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)
	tree.SetBorder(true)
	return tree
}

func addChildren(node *tview.TreeNode, edge *api.Edge, m map[common.Hash]*api.Edge, selectedCB func(*api.Edge)) {
	me := tview.NewTreeNode(edgeToTreeNodeName(edge)).
		SetSelectedFunc(func() {
			selectedCB(edge)
		})
	if edge.HasRival {
		me.SetColor(tcell.ColorRed)
	}

	node.AddChild(me)

	if uc, ok := m[edge.UpperChildID]; ok {
		addChildren(me, uc, m, selectedCB)
	}
	if lc, ok := m[edge.LowerChildID]; ok {
		addChildren(me, lc, m, selectedCB)
	}
}

func EdgesToMap(edges []*api.Edge) map[common.Hash]*api.Edge {
	m := make(map[common.Hash]*api.Edge, len(edges))
	for _, e := range edges {
		m[e.ID] = e
	}
	return m
}

func edgeToTreeNodeName(e *api.Edge) string {
	// ASSERTIONHASH-EDGETYPE-STARTCOMMITHEIGHT-STARTCOMMITROOT-ENDCOMMITHEIGHT-ENDCOMMITROOT

	trunc := func(b []byte) []byte {
		return b[:4]
	}

	return fmt.Sprintf(
		"%#x-%s-%d-%#x-%d-%#x",
		trunc(e.AssertionHash[:]),
		e.Type,
		e.StartCommitment.Height,
		trunc(e.StartCommitment.Hash[:]),
		e.EndCommitment.Height,
		trunc(e.EndCommitment.Hash[:]),
	)
}
