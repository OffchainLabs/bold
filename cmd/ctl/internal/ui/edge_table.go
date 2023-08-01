package ui

import (
	"fmt"

	"github.com/OffchainLabs/bold/api"
	"github.com/OffchainLabs/bold/cmd/ctl/internal/data"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func EdgeTablePage() (title string, content tview.Primitive) {
	edges, err := data.LoadEdgesFromDisk()
	if err != nil {
		panic(err)
	}

	table := tview.NewTable().SetBorders(true)
	table.SetContent(&EdgesTableContent{Edges: edges})

	mainView := table

	footer := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().SetText("TODO: Filter(s), sort, refresh data, etc"), 1, 0, false)

	footer.SetBorder(true)

	return "Edge Table", tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(mainView, 0, 1, true).
		AddItem(footer, 6, 0, true)
}

var _ = tview.TableContent(&EdgesTableContent{})

type EdgesTableContent struct {
	tview.TableContentReadOnly // Inherit no-op functions for any mutating functions.

	// Edges represent the fetched data
	Edges []*api.Edge

	// Columns represents the ordered list of columns
	Columns []string

	Selected string
}

func (e *EdgesTableContent) GetCell(row, column int) *tview.TableCell {
	if row > len(e.Edges) || column >= len(e.ColumnNames()) || row < 0 || column < 0 {
		return tview.NewTableCell("MISSING!")
	}

	if row == 0 {
		return tview.NewTableCell(e.ColumnNames()[column])
	}

	var str string

	ee := e.Edges[row-1]

	switch e.ColumnNames()[column] {
	case ColumnID:
		str = ee.ID.Hex()
	case ColumnType:
		str = ee.Type
	case ColumnStartCommitmentHeight:
		if ee.StartCommitment == nil {
			str = "MISSING: StartCommitment"
		} else {
			str = fmt.Sprintf("%d", ee.StartCommitment.Height)
		}
	case ColumnStartCommitmentHash:
		if ee.StartCommitment == nil {
			str = "MISSING: StartCommitment"
		} else {
			str = ee.StartCommitment.Hash.Hex()
		}
	case ColumnEndCommitmentHeight:
		if ee.EndCommitment == nil {
			str = "MISSING: EndCommitment"
		} else {
			str = fmt.Sprintf("%d", ee.EndCommitment.Height)
		}
	case ColumnEndCommitmentHash:
		if ee.EndCommitment == nil {
			str = "MISSING: EndCommitment"
		} else {
			str = ee.EndCommitment.Hash.Hex()
		}
	case ColumnCreatedAtBlock:
		str = fmt.Sprintf("%d", ee.CreatedAtBlock)
	case ColumnMutualID:
		str = ee.MutualID.Hex()
	case ColumnOriginID:
		str = ee.ClaimID.Hex()
	case ColumnClaimID:
		str = ee.ClaimID.Hex()
	case ColumnHasChildren:
		str = fmt.Sprintf("%t", ee.HasChildren)
	case ColumnLowerChildID:
		str = ee.LowerChildID.Hex()
	case ColumnUpperChildID:
		str = ee.UpperChildID.Hex()
	case ColumnMiniStaker:
		str = ee.MiniStaker.Hex()
	case ColumnAssertionHash:
		str = ee.AssertionHash.Hex()
	case ColumnTimeUnrivaled:
		// TODO: What unit is time unrivaled?
		str = fmt.Sprintf("%d", ee.TimeUnrivaled)
	case ColumnHasRival:
		str = fmt.Sprintf("%t", ee.HasRival)
	case ColumnStatus:
		str = ee.Status
	case ColumnHasLengthOneRival:
		str = fmt.Sprintf("%t", ee.HasLengthOneRival)
	case ColumnTopLevelClaimBlockChallengeOriginHeight:
		if ee.TopLevelClaimHeight == nil {
			str = "MISSING: TopLevelClaimHeight"
		} else {
			str = fmt.Sprintf("%d", ee.TopLevelClaimHeight.BlockChallengeOriginHeight)
		}
	case ColumnTopLevelClaimBigStepChallengeOriginHeight:
		if ee.TopLevelClaimHeight == nil {
			str = "MISSING: TopLevelClaimHeight"
		} else {
			str = fmt.Sprintf("%d", ee.TopLevelClaimHeight.BigStepChallengeOriginHeight)
		}
	default:
		str = "ERROR: Unhandled field name"
	}

	cell := tview.NewTableCell(str).SetSelectable(true).SetClickedFunc(func() bool {
		e.Selected = str
		return true
	})

	if e.Selected != "" && str == e.Selected {
		cell.SetBackgroundColor(tcell.ColorGrey)
	}

	return cell
}

func (e *EdgesTableContent) GetRowCount() int {
	return len(e.Edges) + 1 // Add one for the column headers
}

func (e *EdgesTableContent) GetColumnCount() int {
	return len(e.ColumnNames())
}

var (
	ColumnID                                        = "ID"
	ColumnType                                      = "Type"
	ColumnStartCommitmentHeight                     = "StartCommitment.Height"
	ColumnStartCommitmentHash                       = "StartCommitment.Hash"
	ColumnEndCommitmentHeight                       = "EndCommitment.Height"
	ColumnEndCommitmentHash                         = "EndCommitment.Hash"
	ColumnCreatedAtBlock                            = "CreatedAtBlock"
	ColumnMutualID                                  = "MutualID"
	ColumnOriginID                                  = "OriginID"
	ColumnClaimID                                   = "ClaimID"
	ColumnHasChildren                               = "HasChildren"
	ColumnLowerChildID                              = "LowerChildID"
	ColumnUpperChildID                              = "UpperChildID"
	ColumnMiniStaker                                = "MiniStaker"
	ColumnAssertionHash                             = "AssertionHash"
	ColumnTimeUnrivaled                             = "TimeUnrivaled"
	ColumnHasRival                                  = "HasRival"
	ColumnStatus                                    = "Status"
	ColumnHasLengthOneRival                         = "HasLengthOneRival"
	ColumnTopLevelClaimBlockChallengeOriginHeight   = "TopLevelClaimHeight.BlockChallengeOriginHeight"
	ColumnTopLevelClaimBigStepChallengeOriginHeight = "TopLevelClaimHeight.BigStepChallengeOriginHeight"
)

var DefaultColumnOrder = []string{
	ColumnID,
	ColumnType,
	ColumnStartCommitmentHeight,
	ColumnStartCommitmentHash,
	ColumnEndCommitmentHeight,
	ColumnEndCommitmentHash,
	ColumnCreatedAtBlock,
	ColumnMutualID,
	ColumnOriginID,
	ColumnClaimID,
	ColumnHasChildren,
	ColumnLowerChildID,
	ColumnUpperChildID,
	ColumnMiniStaker,
	ColumnAssertionHash,
	ColumnTimeUnrivaled,
	ColumnHasRival,
	ColumnStatus,
	ColumnHasLengthOneRival,
	ColumnTopLevelClaimBlockChallengeOriginHeight,
	ColumnTopLevelClaimBigStepChallengeOriginHeight,
}

func (e *EdgesTableContent) ColumnNames() []string {
	if len(e.Columns) == 0 {
		return DefaultColumnOrder
	}
	return e.Columns
}
