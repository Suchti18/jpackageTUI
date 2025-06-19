package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/nils/jpackageTUI/internal/option"
	"github.com/rivo/tview"
	"strconv"
)

type Primitive interface {
	focus()
	blur()
	tview.Primitive
}

const (
	appImage = "app-image"
	exe      = "exe"
	msi      = "msi"
	rpm      = "rpm"
	deb      = "deb"
	pkg      = "pkg"
	dmg      = "dmg"
)

var (
	UI *Ui

	typeOptions = []string{appImage, exe, msi, rpm, deb, pkg, dmg}
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
	genericOptionUI.addNextButton()

	linkOptionUI := NewOptionsUI([]*option.Option{
		option.NewOption(
			"Add modules",
			"A comma (\",\") separated list of modules to add",
			"--add-modules",
			option.CrossPlatform,
			false,
			true,
			[]string{}),
		option.NewOption(
			"Module path",
			"Each path is either a directory of modules or the path to a modular jar, and is absolute or relative to the current directory.",
			"--module-path",
			option.CrossPlatform,
			false,
			true,
			[]string{}),
	})
	linkOptionUI.addFinishButton()

	ui.pages = tview.NewPages().
		AddAndSwitchToPage("1", genericOptionUI.GetPrimitive(), true).
		AddPage("2", linkOptionUI.GetPrimitive(), true, false)

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

func nextPage() {
	name, _ := UI.pages.GetFrontPage()
	s, _ := strconv.Atoi(name)
	UI.pages.SwitchToPage(strconv.Itoa(s + 1))
}

func finish() {

}
