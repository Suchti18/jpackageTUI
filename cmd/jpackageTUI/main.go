package main

import (
	"github.com/rivo/tview"
	"time"
)

var (
	app  *tview.Application
	flex *tview.Flex
)

func addNewButton() {
	text := tview.NewTextView().SetText("Hello, world")
	flex.AddItem(text, 0, 1, false)
}

func every5Sec() {
	tick := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-tick.C:
			addNewButton()
		}
	}
}

func main() {
	app = tview.NewApplication()
	app.EnableMouse(true)

	flex = tview.NewFlex()
	flex.SetBorder(true)
	flex.SetTitle("jpackageTUI")
	flex.SetDirection(tview.FlexRow)

	button := tview.NewButton("Click me")
	button.SetFocusFunc(func() {
		addNewButton()
	})

	flex.AddItem(tview.NewTextView().SetText("Hello, world"), 0, 1, false)
	flex.AddItem(button, 0, 1, false)

	//go every5Sec()

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
