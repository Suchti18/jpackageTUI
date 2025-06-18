package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/nils/jpackageTUI/internal/option"
	"github.com/rivo/tview"
)

type Primitive interface {
	focus()
	blur()
	tview.Primitive
}

const (
	none     = "default"
	appimage = "app-image"
	exe      = "exe"
	msi      = "msi"
	rpm      = "rpm"
	deb      = "deb"
	pkg      = "pkg"
	dmg      = "dmg"
)

var (
	UI          *Ui
	typeOptions = []string{none, appimage, exe, msi, rpm, deb, pkg, dmg}
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
	genericOptionUI := NewOptionsUI([]*option.Option{
		option.NewOption(
			"Type",
			"Type of the result",
			"--type",
			option.CrossPlatform,
			true,
			true,
			typeOptions),
		option.NewOption(
			"App Version",
			"Version of the application and/or package",
			"--app-version",
			option.CrossPlatform,
			true,
			true,
			[]string{}),
	})

	grid := tview.NewGrid().
		SetRows(0).
		SetColumns(0).
		AddItem(genericOptionUI.GetPrimitive(), 0, 0, 1, 1, 0, 0, true)

	grid.SetBorder(true)

	ui.pages = tview.NewPages().
		AddAndSwitchToPage("main", grid, true)

	ui.app.SetRoot(ui.pages, true)
	ui.app.SetFocus(genericOptionUI.GetPrimitive())

	ui.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			ui.app.Stop()
		}

		return event
	})

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
