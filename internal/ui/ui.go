package ui

import (
	"github.com/rivo/tview"
)

type Primitive interface {
	focus()
	blur()
	tview.Primitive
}

var (
	UI *Ui
)

type Ui struct {
	app          *tview.Application
	pages        *tview.Pages
	current      int
	primitives   []Primitive
	primitiveLen int
	updater      chan func()
}

func New() *Ui {
	ui := &Ui{
		app: tview.NewApplication(),
	}

	ui.updater = make(chan func(), 100)

	UI = ui

	return ui
}

func (ui *Ui) Start() error {
	NewOptionsUI()

	grid := tview.NewGrid().
		SetRows(0).
		SetColumns(0).
		AddItem(OptionUI, 0, 0, 1, 1, 0, 0, true)

	ui.pages = tview.NewPages().
		AddAndSwitchToPage("main", grid, true)

	ui.app.SetRoot(ui.pages, true)
	ui.app.SetFocus(OptionUI)

	go func() {
		for f := range UI.updater {
			go ui.app.QueueUpdateDraw(f)
		}
	}()

	if err := ui.app.Run(); err != nil {
		ui.app.Stop()
		return err
	}

	return nil
}
