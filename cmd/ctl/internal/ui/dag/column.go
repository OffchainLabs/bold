package dag

import "github.com/OffchainLabs/bold/cmd/ctl/internal/ui/dag/linkline"

type ColumnType int

const (
	ColumnTypeEmpty ColumnType = iota
	ColumnTypeBlocked
	ColumnTypeReserved
	ColumnTypeAncestor
	ColumnTypeParent
)

type Column[N Item] struct {
	Type ColumnType
	node N
}

// Matches returns true if the column is not empty and the node matches the
// node assigned to this column.
func (c *Column[N]) Matches(n N) bool {
	if c.Type == ColumnTypeEmpty || c.Type == ColumnTypeBlocked {
		return false
	}
	return c.node.ID() == n.ID()
}

// Variant returns the column type.
func (c *Column[N]) Variant() uint {
	return uint(c.Type)
}

// Merge merges the column with another column.
func (c *Column[N]) Merge(other *Column[N]) {
	if other == nil {
		return
	}
	if other.Variant() > c.Variant() {
		c.Type = other.Type
		c.node = other.node
	}
}

// Reset clears a blocked column.
func (c *Column[N]) Reset() {
	if c.Type == ColumnTypeBlocked {
		c.Type = ColumnTypeEmpty
	}
}

func (c *Column[N]) ToLinkLine() linkline.LinkLine {
	switch c.Type {
	case ColumnTypeAncestor:
		return linkline.VerticalAncestor
	case ColumnTypeParent:
		return linkline.VerticalParent
	default:
		return linkline.Empty
	}
}

func (c *Column[N]) ToNodeLine() NodeLine {
	switch c.Type {
	case ColumnTypeAncestor:
		return NodeLineAncestor
	case ColumnTypeParent:
		return NodeLineParent
	}
	return NodeLineBlank
}

func (c *Column[N]) ToPadLine() PadLine {
	switch c.Type {
	case ColumnTypeAncestor:
		return PadLineAncestor
	case ColumnTypeParent:
		return PadLineParent
	}
	return PadLineBlank
}
