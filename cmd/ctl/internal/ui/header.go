package ui

import (
	"fmt"
	"time"

	"github.com/OffchainLabs/challenge-protocol-v2/cmd/ctl/internal/data"
	"github.com/rivo/tview"
)

func WithHeader(p page) page {
	return func() (title string, content tview.Primitive) {
		title, content = p()
		return title, tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(header(), 8, 0, false).
			AddItem(content, 0, 1, true)
	}
}

// Header consists of the server information
func header() tview.Primitive {
	left := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().SetText("Server: localhost:1234"), 1, 0, false).
		AddItem(dynamicBlockNumber(), 1, 0, false).
		AddItem(tview.NewBox(), 1, 0, false).
		AddItem(tview.NewBox(), 1, 0, false).
		AddItem(tview.NewBox(), 1, 0, false).
		AddItem(tview.NewTextView().SetText("Note: Using mock data"), 1, 0, false)

	left.SetBorderPadding(0, 0, 1, 0)

	flex := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(left, 0, 1, false).
		AddItem(tview.NewTextView().SetText(logoSmall), 9, 0, false)

	flex.SetBorder(true).
		SetTitle("Challenge V2")

	return flex
}

func dynamicBlockNumber() tview.Primitive {
	txt := tview.NewTextView()

	bn := data.CurrentBlockNumber()
	format := "Current block number: %d"

	txt.SetText(fmt.Sprintf(format, bn))

	go func() {
		for {
			time.Sleep(1 * time.Second)
			bn++
			app.QueueUpdateDraw(func() {
				txt.SetText(fmt.Sprintf(format, bn))
			})
		}
	}()

	return txt
}
