package ui

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const logo = `
░█████╗░██╗░░██╗░█████╗░██╗░░░░░██╗░░░░░███████╗███╗░░██╗░██████╗░███████╗  ██╗░░░██╗██████╗░
██╔══██╗██║░░██║██╔══██╗██║░░░░░██║░░░░░██╔════╝████╗░██║██╔════╝░██╔════╝  ██║░░░██║╚════██╗
██║░░╚═╝███████║███████║██║░░░░░██║░░░░░█████╗░░██╔██╗██║██║░░██╗░█████╗░░  ╚██╗░██╔╝░░███╔═╝
██║░░██╗██╔══██║██╔══██║██║░░░░░██║░░░░░██╔══╝░░██║╚████║██║░░╚██╗██╔══╝░░  ░╚████╔╝░██╔══╝░░
╚█████╔╝██║░░██║██║░░██║███████╗███████╗███████╗██║░╚███║╚██████╔╝███████╗  ░░╚██╔╝░░███████╗
░╚════╝░╚═╝░░╚═╝╚═╝░░╚═╝╚══════╝╚══════╝╚══════╝╚═╝░░╚══╝░╚═════╝░╚══════╝  ░░░╚═╝░░░╚══════╝`

const logoSmall = `██╗░░░██╗██████╗░
██║░░░██║╚════██╗
╚██╗░██╔╝░░███╔═╝
░╚████╔╝░██╔══╝░░
░░╚██╔╝░░███████╗
░░░╚═╝░░░╚══════╝`

const (
	subtitle   = `Challenge / Edge browser for Arbitrum Protocol Challenge V2`
	navigation = `ctrl-C: Exit`
)

func CoverPage() (title string, content tview.Primitive) {
	// What's the size of the logo?
	lines := strings.Split(logo, "\n")
	logoWidth := 94
	logoHeight := len(lines)
	logoBox := tview.NewTextView().
		SetTextColor(ThemeLogoColor)

	fmt.Fprint(logoBox, logo)
	// Create a frame for the subtitle and navigation infos.
	frame := tview.NewFrame(tview.NewBox()).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText(subtitle, true, tview.AlignCenter, ThemeSubtitleColor).
		AddText("", true, tview.AlignCenter, ThemeSubtitleColor).
		AddText(navigation, true, tview.AlignCenter, tcell.ColorDarkMagenta)

	// Create a Flex layout that centers the logo and subtitle.
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewBox(), 0, 7, false).
		AddItem(tview.NewFlex().
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(logoBox, logoWidth, 1, false).
			AddItem(tview.NewBox(), 0, 1, false), logoHeight, 1, true).
		AddItem(frame, 0, 10, false)

	return "Start", flex
}
