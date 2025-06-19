package ui

import "github.com/nils/jpackageTUI/internal/option"

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
	genericOptionUI *OptionUI
	linkOptionUI    *OptionUI

	typeOptions = []string{appImage, exe, msi, rpm, deb, pkg, dmg}
)

func LoadAll() {
	loadGeneric()
	loadLinks()

	UI.optionUIs = append(UI.optionUIs, genericOptionUI)
	UI.optionUIs = append(UI.optionUIs, linkOptionUI)
}

func loadGeneric() {
	genericOptionUI = NewOptionsUI([]*option.Option{
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
}

func loadLinks() {
	linkOptionUI = NewOptionsUI([]*option.Option{
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
	linkOptionUI.addBackButton()
	linkOptionUI.addFinishButton()
}
