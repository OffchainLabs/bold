// Package dag provides a directed acyclic graph and rendering utilities.
// This package is largely based on github.com/facebook/sapling's dag and renderdag crates
// (MIT License).
package dag

import (
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

const (
	GlyphSpace = iota
	GlyphHorizontal
	GlyphParent
	GlyphAncestor
	GlyphMergeLeft
	GlyphMergeRight
	GlyphMergeBoth
	GlyphForkLeft
	GlyphForkRight
	GlyphForkBoth
	GlyphJoinLeft
	GlyphJoinRight
	GlyphJoinBoth
	GlyphTermination
	GlyphCircle
)

var glyphs = []string{
	"  ", // GlyphSpace
	"──", // GlyphHorizontal
	"│ ", // GlyphParent
	"╷ ", // GlyphAncestor
	"╯ ", // GlyphMergeLeft
	"╰─", // GlyphMergeRight
	"┴─", // GlyphMergeBoth
	"╮ ", // GlyphForkLeft
	"╭─", // GlyphForkRight
	"┬─", // GlyphForkBoth
	"┤ ", // GlyphJoinLeft
	"├─", // GlyphJoinRight
	"┼─", // GlyphJoinBoth
	"~ ", // GlyphTermination
	"◉ ", // GlyphCircle
}

type Item interface {
	ID() common.Hash
	DisplayID() string
	Description() string
	Timestamp() int64
}

type Node[T Item] struct {
	Item T

	// The node ID is a unique identifier for the node.
	ID       common.Hash
	ParentID common.Hash
}

// DAG is a directed acyclic graph
// The graph is represented as a set of nodes and a set of edges, where each edge is a pair of nodes.
type DAG[T Item] struct {
	nodes   map[common.Hash]Node[T]
	heads   map[common.Hash]bool
	parents map[common.Hash]bool
}

// Init initializes the DAG.
func (d *DAG[T]) Init() {
	d.nodes = make(map[common.Hash]Node[T])
	d.heads = make(map[common.Hash]bool)
	d.parents = make(map[common.Hash]bool)
}

// AddNode adds a node to the DAG.
func (d *DAG[T]) AddNode(node Node[T]) {
	// Add node to storage.
	d.nodes[node.ID] = node
	// If no other node in the map references this one, then it's a head.
	delete(d.heads, node.ParentID)
	d.heads[node.ID] = true
	// Add parent to storage.
	d.parents[node.ParentID] = true
}

func (d *DAG[T]) RenderString() (s string) {
	r := NewAsciiRenderer[T]()

	// Ensure nodes are sorted be timestamp where the most recent one is at index 0.
	nodes := make([]Node[T], 0, len(d.nodes))
	for _, node := range d.nodes {
		n := node
		nodes = append(nodes, n)
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Item.Timestamp() > nodes[j].Item.Timestamp()
	})

	for _, node := range nodes {

		typ := AncestorTypeParent
		parent := &Ancestor[T]{
			node: d.nodes[node.ParentID].Item,
			Type: typ,
		}

		msg := strings.Join([]string{
			node.Item.DisplayID(),
			node.Item.Description(),
		}, "\n")

		row := r.NextRow(node.Item, []Ancestor[T]{*parent}, glyphs[GlyphCircle], msg)
		s += row
	}

	return
}
