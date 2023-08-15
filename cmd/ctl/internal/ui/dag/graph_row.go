package dag

import "github.com/OffchainLabs/bold/cmd/ctl/internal/ui/dag/linkline"

// GraphRow represents an item as a row on the graph to be rendered.
type GraphRow[N Item] struct {
	// The node for this row.
	Node N

	// The glyph for this row.
	Glyph string

	// The message for this row.
	Message string

	// True if this row is for a merge commit.
	// TODO: There are no merge commits?
	Merge bool

	// The node columns for this row.
	NodeLine []NodeLine

	// The link columns for this row, if a link row is necessary.
	LinkLine []linkline.LinkLine

	// The location of any terminators, if necessary. Other columns should be
	// filled in with pad lines.
	TermLine []bool

	// The pad columns for this row.
	PadLines []PadLine
}
