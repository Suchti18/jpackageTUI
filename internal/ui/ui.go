package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/nils/jpackageTUI/internal/option"
	"github.com/rivo/tview"
	"strconv"
)

var (
	UI *Ui
)

type Ui struct {
	app       *tview.Application
	pages     *tview.Pages
	current   int
	optionUIs []*OptionUI
}

func New() *Ui {
	ui := &Ui{
		app: tview.NewApplication(),
	}

	UI = ui

	return ui
}

func (ui *Ui) Start() error {
	LoadAll()

	ui.pages = tview.NewPages()

	for index, optionUI := range ui.optionUIs {
		ui.pages.AddPage(strconv.Itoa(index), optionUI.GetPrimitive(), true, false)
	}
	ui.current = 0
	switchPage(ui.current)

	ui.app.SetRoot(ui.pages, true)
	ui.app.EnableMouse(true)

	ui.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			UI.app.Stop()
		}

		return event
	})

	if err := ui.app.Run(); err != nil {
		UI.app.Stop()
		return err
	}

	return nil
}

func previousPage() {
	UI.current--
	switchPage(UI.current)
}

func nextPage() {
	UI.current++
	switchPage(UI.current)
}

func switchPage(name int) {
	pageName := strconv.Itoa(name)

	if UI.pages.HasPage(pageName) {
		UI.pages.SwitchToPage(pageName)
	}
}

func finish() {
	for _, optionUI := range UI.optionUIs {
		option.AddMapToRepo(optionUI.getData())
	}

	UI.app.Stop()
}
