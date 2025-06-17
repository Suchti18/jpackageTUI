package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow)

	button := tview.NewButton("Add Text").SetSelectedFunc(func() {
		text := tview.NewTextView().SetText("Hello, world")
		flex.AddItem(text, 0, 1, false)
		app.ForceDraw()
	}).SetLabelColor(tcell.ColorWhite).SetLabelColorActivated(tcell.ColorRed)

	flex.AddItem(tview.NewTextView().SetText("Hello, world"), 0, 1, false)
	flex.AddItem(button, 0, 1, false)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
