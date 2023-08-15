package dag

import (
	"strings"

	"github.com/OffchainLabs/bold/cmd/ctl/internal/ui/dag/linkline"
)

type Renderer[N Item] interface {
	// Returns the witdth of the graph line, possibly including another node.
	Width(newNode *N, newParents []*Ancestor[N]) uint64

	// Reserve a column for the given node.
	Reserve(node *N)

	// Render the next row.
	NextRow(node *N, parents []*Ancestor[N], glyph, message string) GraphRow[N]
}

type Columns[N Item] []Column[N]

// Find the index of the node in the columns. Returns -1 if not found.
func (c Columns[N]) Find(node N) int {
	for i, cc := range c {
		if cc.Type != ColumnTypeEmpty && cc.node.ID() == node.ID() {
			return i
		}
	}
	return -1
}

// FindEmpty returns the arg if the column is empty, or -1 if not.
// TODO: This should be IsEmptyAt(index int) bool
func (c Columns[N]) FindEmpty(index int) int {
	if index > len(c) || index < 0 {
		return -1
	}
	if c[index].Type == ColumnTypeEmpty {
		return index
	}
	return -1
}

// FirstEmpty returns the index of the first empty column, or -1 if none.
func (c Columns[N]) FirstEmpty() int {
	for i, cc := range c {
		if cc.Type == ColumnTypeEmpty {
			return i
		}
	}
	return -1
}

// NewEmpty returns the index of a new empty column, appended to the end.
func (c *Columns[N]) NewEmpty() int {
	*c = append(*c, Column[N]{Type: ColumnTypeEmpty})
	return len(*c) - 1
}

// ResetColumns resets the columns list and returns the new
func ResetColumns[N Item](c Columns[N]) Columns[N] {
	// Filter columns in place to remove empty columns after reset.
	n := 0
	for _, val := range c {
		val.Reset()
		if val.Type != ColumnTypeEmpty {
			c[n] = val
			n++
		}
	}
	return c[:n] // truncate list
}

func (c Columns[N]) NumEmpty() int {
	n := 0
	for _, cc := range c {
		if cc.Type == ColumnTypeEmpty {
			n++
		}
	}
	return n
}

func (c Columns[N]) AsNodeLines() []NodeLine {
	lines := make([]NodeLine, len(c))
	for i, cc := range c {
		lines[i] = cc.ToNodeLine()
	}
	return lines
}

func (c Columns[N]) AsLinkLines() []linkline.LinkLine {
	lines := make([]linkline.LinkLine, len(c))
	for i, cc := range c {
		lines[i] = cc.ToLinkLine()
	}
	return lines
}

func (c Columns[N]) AsPadLines() []PadLine {
	lines := make([]PadLine, len(c))
	for i, cc := range c {
		lines[i] = cc.ToPadLine()
	}
	return lines
}

// Swap two columns.
func (c Columns[N]) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

type GraphRowRenderer[N Item] struct {
	columns Columns[N]
}

func (g *GraphRowRenderer[N]) Width(node *N, parents []Ancestor[N]) int {
	width := len(g.columns)
	empty := g.columns.NumEmpty()
	if node != nil {
		// If the node is not already allocated, and there is no space
		// for the node, then adding the new node would create a new column.
		if g.columns.Find(*node) == -1 {
			if empty == 0 {
				width++
			} else {
				empty--
			}
		}
	}
	unallocatedParents := 0
	for _, p := range parents {
		if g.columns.Find(p.node) == -1 {
			unallocatedParents++
		}
	}
	unallocatedParents -= empty
	if unallocatedParents < 0 {
		unallocatedParents = 0
	}
	if unallocatedParents > 0 {
		width += unallocatedParents - 1
	}

	return width
}

// Reserve a column for a node using the first empty column or allocating a new one.`s`
func (g *GraphRowRenderer[N]) Reserve(node N) {
	if g.columns.Find(node) == -1 {
		col := Column[N]{
			Type: ColumnTypeReserved,
			node: node,
		}
		// Insert into first empty or append.
		if i := g.columns.FirstEmpty(); i > -1 {
			g.columns[i] = col
		} else {
			g.columns = append(g.columns, col)
		}
	}
}

func (g *GraphRowRenderer[N]) NextRow(node N, parents []Ancestor[N], glyph, message string) GraphRow[N] {
	// Find a column for this node.
	column := g.columns.Find(node)
	if column == -1 {
		column = g.columns.FirstEmpty()
	}
	if column == -1 {
		column = g.columns.NewEmpty()
	}
	g.columns[column].Type = ColumnTypeEmpty

	// This row is for a merge if there are multiple parents.
	merge := len(parents) > 1

	// Build the initial node line.
	nodeLines := g.columns.AsNodeLines()
	nodeLines[column] = NodeLineNode

	// Build the initial link line.
	linkLines := g.columns.AsLinkLines()
	needLinkLine := false

	// Build the initial term line.
	termLines := make([]bool, len(g.columns))
	needTermLine := false

	// Build the initial pad line.
	padLines := g.columns.AsPadLines()

	// Assign each parent to a column.
	parentColumns := make(map[int]Ancestor[N])
	for _, p := range parents {
		// Check if the parent already has a column.
		if idx := g.columns.Find(p.node); idx != -1 {
			g.columns[idx].Merge(p.ToColumn())
			parentColumns[idx] = p
			continue
		}

		// Assign the parent to an empty column, preferring the column
		// the current node is going in, to maintain linearity.
		if idx := g.columns.FindEmpty(column); idx != -1 {
			g.columns[idx].Merge(p.ToColumn())
			parentColumns[idx] = p
			continue
		}

		// There are no empty columns left. Make a new column.
		parentColumns[len(g.columns)] = p
		nodeLines = append(nodeLines, NodeLineBlank)
		padLines = append(padLines, PadLineBlank)
		linkLines = append(linkLines, linkline.Empty)
		termLines = append(termLines, false)
		g.columns = append(g.columns, *p.ToColumn())
	}

	// Mark parent columns with anonymous parents as terminating.
	for i, p := range parentColumns {
		if p.ID() == emptyHash {
			termLines[i] = true
			needTermLine = true
		}
	}

	// Check if we can move the parent to the current column.
	if len(parents) == 1 {
		for parentColumn, parent := range parentColumns {
			if parentColumn > column {
				g.columns.Swap(column, parentColumn)
				delete(parentColumns, parentColumn)
				parentColumns[column] = parent

				// Generate a line from this column to the old parent column.
				// We need to continue with the style that was being used for the
				// parent column.
				wasDirect := linkLines[parentColumn]&linkline.VerticalParent == linkline.VerticalParent
				if wasDirect {
					linkLines[column] |= linkline.RightForkParent
				} else {
					linkLines[column] |= linkline.RightForkAncestor
				}
				for i := 0; i < parentColumn; i++ {
					if wasDirect {
						linkLines[i] |= linkline.HorizontalParent
					} else {
						linkLines[i] |= linkline.HorizontalAncestor
					}
				}
				if wasDirect {
					linkLines[parentColumn] = linkline.LeftMergeParent
				} else {
					linkLines[parentColumn] = linkline.LeftMergeAncestor
				}
				needLinkLine = true
				// The pad line for the old parent column is now blank.
				padLines[parentColumn] = PadLineBlank
			}
		}
	}

	// Connect the node column to all the parent columns.
	acb := NewAncestorColumnBounds[N](parentColumns, column)
	for i, to := acb.Range(); i < to; i++ {
		linkLines[i] |= acb.HorizontalLine(i)
		needLinkLine = true
	}

	// If there is a parent or ancestor to the right of the node column,
	// the node merges from the right.
	if acb.maxParent > column {
		linkLines[column] |= linkline.RightMergeParent
		needLinkLine = true
	} else if acb.maxAncestor > column {
		linkLines[column] |= linkline.RightMergeAncestor
		needLinkLine = true
	}

	// If there is as parent or ancestor to the left of the node column,
	// the node forks from the left.
	if acb.minParent < column {
		linkLines[column] |= linkline.LeftMergeParent
		needLinkLine = true
	} else if acb.minAncestor < column {
		linkLines[column] |= linkline.LeftMergeAncestor
		needLinkLine = true
	}

	// Each parent or ancestor forks toward the node column.
	for i, p := range parentColumns {
		padLines[i] = g.columns[i].ToPadLine()
		if i < column {
			linkLines[i] |= p.ToLinkLine(linkline.RightForkParent, linkline.RightForkAncestor)
		} else if i == column {
			linkLines[i] |= linkline.Child | p.ToLinkLine(linkline.VerticalParent, linkline.VerticalAncestor)
		} else {
			linkLines[i] |= p.ToLinkLine(linkline.LeftForkParent, linkline.LeftForkAncestor)
		}
	}

	// Now that we have assigned all of the columns, reset their state.
	g.columns = ResetColumns[N](g.columns)

	// Filter out the link line or term line if they are not needed.
	if !needLinkLine {
		linkLines = nil
	}
	if !needTermLine {
		termLines = nil
	}

	return GraphRow[N]{
		Node:     node,
		Glyph:    glyph,
		Message:  message,
		Merge:    merge,
		NodeLine: nodeLines,
		LinkLine: linkLines,
		TermLine: termLines,
		PadLines: padLines,
	}
}

func NewAsciiRenderer[N Item]() *AsciiRenderer[N] {
	return &AsciiRenderer[N]{
		renderer:     &GraphRowRenderer[N]{},
		minRowHeight: 2,
	}
}

type AsciiRenderer[N Item] struct {
	renderer     *GraphRowRenderer[N]
	minRowHeight int
	extraPadLine string
}

func (a *AsciiRenderer[N]) Width(node *N, parents []Ancestor[N]) int {
	return a.renderer.Width(node, parents)*2 + 1
}

func (a *AsciiRenderer[N]) Reserve(node N) {
	a.renderer.Reserve(node)
}

func (a *AsciiRenderer[N]) NextRow(node N, parents []Ancestor[N], glyph, message string) string {
	line := a.renderer.NextRow(node, parents, glyph, message)
	var out string
	// Split message into lines and pad to the minimum number of lines.
	messageLines := strings.Split(line.Message, "\n")
	for i := len(messageLines); i < a.minRowHeight; i++ {
		messageLines = append(messageLines, "")
	}
	needExtraPadLine := false
	// Render the previous extra pad line.
	if len(a.extraPadLine) > 0 {
		out += a.extraPadLine + "\n"
		a.extraPadLine = ""
	}

	// Render the nodeline.
	nodeLine := ""
	for _, nl := range line.NodeLine {
		switch nl {
		case NodeLineNode:
			nodeLine += line.Glyph
		case NodeLineParent:
			nodeLine += glyphs[GlyphParent]
		case NodeLineAncestor:
			nodeLine += glyphs[GlyphAncestor]
		case NodeLineBlank:
			nodeLine += glyphs[GlyphSpace]
		}
	}

	// Pop a message line and write it.
	if len(messageLines) > 0 {
		ml := messageLines[0]
		messageLines = messageLines[1:]
		nodeLine += " " + ml
	}

	out += strings.TrimSuffix(nodeLine, " ") + "\n"

	// Render the link line
	if len(line.LinkLine) > 0 {
		lline := ""
		for _, ll := range line.LinkLine {
			if ll.Intersects(linkline.Horizontal) {
				if ll.Intersects(linkline.Child) {
					lline += glyphs[GlyphJoinBoth]
				} else if ll.Intersects(linkline.AnyFork) && ll.Intersects(linkline.AnyMerge) {
					lline += glyphs[GlyphJoinBoth]
				} else if ll.Intersects(linkline.AnyFork) && ll.Intersects(linkline.VerticalParent) && !line.Merge {
					lline += glyphs[GlyphJoinBoth]
				} else if ll.Intersects(linkline.AnyFork) {
					lline += glyphs[GlyphForkBoth]
				} else if ll.Intersects(linkline.AnyMerge) {
					lline += glyphs[GlyphMergeBoth]
				} else {
					lline += glyphs[GlyphHorizontal]
				}
			} else if ll.Intersects(linkline.VerticalParent) && !line.Merge {
				left := ll.Intersects(linkline.LeftMerge | linkline.LeftFork)
				right := ll.Intersects(linkline.RightMerge | linkline.RightFork)
				switch {
				case left && right:
					lline += glyphs[GlyphJoinBoth]
				case left && !right:
					lline += glyphs[GlyphJoinLeft]
				case !left && right:
					lline += glyphs[GlyphJoinRight]
				default:
					lline += glyphs[GlyphParent]
				}
			} else if ll.Intersects(linkline.VerticalParent|linkline.VerticalAncestor) && !ll.Intersects(linkline.LeftFork|linkline.RightFork) {
				left, right := ll.Intersects(linkline.LeftMerge), ll.Intersects(linkline.RightMerge)
				switch {
				case left && right:
					lline += glyphs[GlyphJoinBoth]
				case left && !right:
					lline += glyphs[GlyphJoinLeft]
				case !left && right:
					lline += glyphs[GlyphJoinRight]
				default:
					if ll.Intersects(linkline.VerticalAncestor) {
						lline += glyphs[GlyphAncestor]
					} else {
						lline += glyphs[GlyphParent]
					}
				}
			} else if ll.Intersects(linkline.LeftFork) && ll.Intersects(linkline.LeftMerge|linkline.Child) {
				lline += glyphs[GlyphJoinLeft]
			} else if ll.Intersects(linkline.RightFork) && ll.Intersects(linkline.RightMerge|linkline.Child) {
				lline += glyphs[GlyphJoinRight]
			} else if ll.Intersects(linkline.LeftMerge) && ll.Intersects(linkline.RightMerge) {
				lline += glyphs[GlyphMergeBoth]
			} else if ll.Intersects(linkline.LeftFork) && ll.Intersects(linkline.RightFork) {
				lline += glyphs[GlyphForkBoth]
			} else if ll.Intersects(linkline.LeftFork) {
				lline += glyphs[GlyphForkLeft]
			} else if ll.Intersects(linkline.LeftMerge) {
				lline += glyphs[GlyphMergeLeft]
			} else if ll.Intersects(linkline.RightFork) {
				lline += glyphs[GlyphForkRight]
			} else if ll.Intersects(linkline.RightMerge) {
				lline += glyphs[GlyphMergeRight]
			} else {
				lline += glyphs[GlyphSpace]
			}
		}
		if len(messageLines) > 0 {
			if len(messageLines[0]) > 0 {
				lline += " " + messageLines[0]
			}
			messageLines = messageLines[1:]
		}
		out += strings.TrimSuffix(lline, " ") + "\n"
	}

	// Render the term line
	if len(line.TermLine) > 0 {
		termStrings := [2]string{glyphs[GlyphParent], glyphs[GlyphTermination]}
		for _, ts := range termStrings {
			tline := ""
			for i, tl := range line.TermLine {
				if tl {
					tline += ts
				} else {
					switch line.PadLines[i] {
					case PadLineParent:
						tline += "| "
					case PadLineAncestor:
						tline += ". "
					case PadLineBlank:
						tline += glyphs[GlyphSpace]
					}
				}
			}
			if len(messageLines) > 0 {
				if len(messageLines[0]) > 0 {
					tline += " " + messageLines[0]
				}
				messageLines = messageLines[1:]
			}

			out += strings.TrimSuffix(tline, " ") + "\n"
		}
		needExtraPadLine = true
	}

	basePadLine := ""
	for _, pl := range line.PadLines {
		switch pl {
		case PadLineParent:
			basePadLine += "│ "
		case PadLineAncestor:
			basePadLine += ". "
		case PadLineBlank:
			basePadLine += glyphs[GlyphSpace]
		}
	}
	// Render any pad lines
	for _, ml := range messageLines {
		pl := basePadLine
		pl += " " + ml
		out += strings.TrimSuffix(pl, " ") + "\n"
		needExtraPadLine = false
	}

	if needExtraPadLine {
		a.extraPadLine = basePadLine
	}

	return out
}
