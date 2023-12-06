package linkline

type LinkLine uint16

const (
	Empty = 0b0_0000_0000_0000

	// This cell contains a horizontal line that connects to a parent.
	HorizontalParent = 0b0_0000_0000_0001

	// This cell contains a horizontal line that connects to an ancestor.
	HorizontalAncestor = 0b0_0000_0000_0010

	// The descendent of this cell is connected to the parent.
	VerticalParent = 0b0_0000_0000_0100

	// The descendent of this cell is connected to an ancestor.
	VerticalAncestor = 0b0_0000_0000_1000

	// The parent of this cell is linked in this link row and the child
	// is to the left.
	LeftForkParent = 0b0_0000_0001_0000

	// The ancestor of this cell is linked in this link row and the child
	// is to the left.
	LeftForkAncestor = 0b0_0000_0010_0000

	// The parent of this cell is linked in this link row and the child
	// is to the right.
	RightForkParent = 0b0_0000_0100_0000

	// The ancestor of this cell is linked in this link row and the child
	// is to the right.
	RightForkAncestor = 0b0_0000_1000_0000

	// The child of this cell is linked to parents on the left.
	LeftMergeParent = 0b0_0001_0000_0000

	// The child of this cell is linked to ancestors on the left.
	LeftMergeAncestor = 0b0_0010_0000_0000

	// The child of this cell is linked to parents on the right.
	RightMergeParent = 0b0_0100_0000_0000

	// The child of this cell is linked to ancestors on the right.
	RightMergeAncestor = 0b0_1000_0000_0000

	// The target node of this link line is the child of this column.
	// This disambiguates between the node that is connected in this link
	// line, and other nodes that are also connected vertically.
	Child = 0b1_0000_0000_0000

	Horizontal     = HorizontalParent | HorizontalAncestor
	Vertical       = VerticalParent | VerticalAncestor
	LeftFork       = LeftForkParent | LeftForkAncestor
	RightFork      = RightForkParent | RightForkAncestor
	LeftMerge      = LeftMergeParent | LeftMergeAncestor
	RightMerge     = RightMergeParent | RightMergeAncestor
	AnyMerge       = LeftMerge | RightMerge
	AnyFork        = LeftFork | RightFork
	AnyForkOrMerge = AnyMerge | AnyFork
)

// Intersects returns true if any of the bits overlap in a or b.
func (ll LinkLine) Intersects(b LinkLine) bool {
	return ll&b != 0b0
}
