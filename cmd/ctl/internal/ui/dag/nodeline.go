package dag

type NodeLine uint

const (
	// Blank
	NodeLineBlank NodeLine = iota
	//Vertical line indicating an ancestor.
	NodeLineAncestor
	// Vertical line indicating a parent.
	NodeLineParent
	// The node for this row.
	NodeLineNode
)
