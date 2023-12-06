package dag

import (
	"github.com/OffchainLabs/bold/cmd/ctl/internal/ui/dag/linkline"
	"github.com/ethereum/go-ethereum/common"
)

var (
	emptyHash = common.Hash{}
)

type AncestorType int

const (
	AncestorTypeAncestor AncestorType = iota
	AncestorTypeParent
	AncestorTypeAnon
)

type Ancestor[N Item] struct {
	Type AncestorType
	node N
}

// IsDirect returns true when the ancestor is a direct parent of the node or an anonymous node.
func (a *Ancestor[N]) IsDirect() bool {
	return a.Type != AncestorTypeAncestor
}

// ToColumn cases an ancestor to a column.
func (a *Ancestor[N]) ToColumn() *Column[N] {
	switch a.Type {
	case AncestorTypeAncestor:
		return &Column[N]{
			Type: ColumnTypeAncestor,
			node: a.node,
		}
	case AncestorTypeParent:
		return &Column[N]{
			Type: ColumnTypeParent,
			node: a.node,
		}
	case AncestorTypeAnon:
		return nil
	}
	return nil
}

// ID is the ID of the node.
func (a *Ancestor[N]) ID() common.Hash {
	if a.Type == AncestorTypeAnon {
		return emptyHash
	}
	return a.node.ID()
}

// ToLinkLink returns the LinkLine for the ancestor.
func (a *Ancestor[N]) ToLinkLine(direct, indirect linkline.LinkLine) linkline.LinkLine {
	if a.IsDirect() {
		return direct
	}
	return indirect
}

type AncestorColumnBounds struct {
	target      int
	minAncestor int
	minParent   int
	maxAncestor int
	maxParent   int
}

func NewAncestorColumnBounds[N Item](columns map[int]Ancestor[N], target int) *AncestorColumnBounds {
	if len(columns) == 0 {
		return nil
	}
	acb := &AncestorColumnBounds{
		target:      target,
		minAncestor: target,
		minParent:   target,
		maxAncestor: target,
		maxParent:   target,
	}

	for k, v := range columns {
		if k < acb.minAncestor {
			acb.minAncestor = k
		}
		if k > acb.maxAncestor {
			acb.maxAncestor = k
		}
		if v.IsDirect() {
			if k < acb.minParent {
				acb.minParent = k
			}
			if k > acb.maxParent {
				acb.maxParent = k
			}
		}
	}
	return acb
}

func (acb *AncestorColumnBounds) HorizontalLine(index int) linkline.LinkLine {
	if index == acb.target {
		return linkline.Empty
	} else if index > acb.minParent && index < acb.maxParent {
		return linkline.HorizontalParent
	} else if index > acb.minAncestor && index < acb.maxAncestor {
		return linkline.HorizontalAncestor
	}
	return linkline.Empty
}

func (acb *AncestorColumnBounds) Range() (from, to int) {
	if acb == nil {
		return 0, 0
	}
	if acb.minAncestor < acb.maxAncestor {
		from = acb.minAncestor + 1
		to = acb.maxAncestor
	}
	return
}
