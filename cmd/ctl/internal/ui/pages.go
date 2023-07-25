package ui

import (
	"fmt"
	"strconv"

	"github.com/rivo/tview"
)

type page func() (title string, content tview.Primitive)

var app *tview.Application

func StartPages() {
	pages := []page{
		CoverPage,
		WithHeader(EdgeTreePage),
		WithHeader(EdgeTablePage),
	}

	p := tview.NewPages()

	// The bottom row has some info on where we are.
	info := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetHighlightedFunc(func(added, removed, remaining []string) {
			p.SwitchToPage(added[0])
		})

	for index, pg := range pages {
		title, primitive := pg()
		p.AddPage(strconv.Itoa(index), primitive, true, index == 0)
		fmt.Fprintf(info, `%d ["%d"][darkcyan]%s[white][""]  `, index+1, index, title)
	}

	info.Highlight("0")

	// Create the main layout.
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(p, 0, 1, true).
		AddItem(info, 1, 1, false)

	// Shortcuts to navigate the slides.
	app = tview.NewApplication()

	if err := app.SetRoot(layout, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
