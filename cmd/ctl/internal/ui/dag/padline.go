package dag

type PadLine uint

const (
	PadLineBlank PadLine = iota

	// Vertical line indicating an ancestor.
	PadLineAncestor

	// Vertical line indiciating a parent.
	PadLineParent
)
