package ui

import "github.com/rivo/tview"

var (
	OptionUI *tview.Form
)

func NewOptionsUI() {
	OptionUI = tview.NewForm()

	options := []string{"default", "app-image", "exe", "msi", "rpm", "deb", "pkg", "dmg"}

	OptionUI.AddDropDown("Type", options, 0, func(option string, optionIndex int) {

	})
}
