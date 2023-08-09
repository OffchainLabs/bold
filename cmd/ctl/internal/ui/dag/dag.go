package dag

import "github.com/ethereum/go-ethereum/common"

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
	GlyphCount
	GlyphCircle
)

var glyphs = []string{"  ", "──", "│ ", "╷ ", "╯ ", "╰─", "┴─", "╮ ", "╭─", "┬─", "┤ ", "├─", "┼─", "~ ", "◉ "}

type Item interface {
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
// The graph will be drawn with the following rules:
// - Each column will start with a parent.
// - The left most column will be the most recent parent.
// - Each row will be two lines
//   - The first line will be a dot with the node ID and timestamp
//   - The second line will be additional node metadata and/or description. Essentially, it needs to be a condensed text representation of the node.
//   - Each row line pair will be a single table row, which is selectable for more details.
//
// - A column is added for each diverging parent.
// - There are no "merges" to consider, so a column will never merge to the left.
type DAG[T Item] struct {
	Nodes map[common.Hash]Node[T]
}

func (d *DAG[T]) AddNode(node Node[T]) {

}

func (d *DAG[T]) RenderString() string {
	// Determine the order of IDs.

	// For each node, find the parent and add it to the column.
	return ""
}
