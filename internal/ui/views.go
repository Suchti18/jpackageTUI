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
			"Name",
			"Name of the application and/or package",
			"--name",
			option.CrossPlatform,
			true,
			true,
			[]string{},
			false),
		option.NewOption(
			"App Version",
			"Version of the application and/or package",
			"--app-version",
			option.CrossPlatform,
			true,
			true,
			[]string{},
			false),
		option.NewOption(
			"Copyright",
			"Copyright for the application",
			"--copyright",
			option.CrossPlatform,
			true,
			true,
			[]string{},
			false),
		option.NewOption(
			"Description",
			"Description of the application",
			"--description",
			option.CrossPlatform,
			true,
			true,
			[]string{},
			false),
		option.NewOption(
			"Type",
			"Type of the result",
			"--type",
			option.CrossPlatform,
			true,
			true,
			typeOptions,
			false),
	})
	genericOptionUI.addNextButton()
}

func loadLinks() {
	linkOptionUI = NewOptionsUI([]*option.Option{
		option.NewOption(
			"Module",
			"The main module (and optionally main class) of the application",
			"--module",
			option.CrossPlatform,
			false,
			true,
			[]string{},
			false),
		option.NewOption(
			"Module path",
			"Each path is either a directory of modules or the path to a modular jar, and is absolute or relative to the current directory.",
			"--module-path",
			option.CrossPlatform,
			false,
			true,
			[]string{},
			false),
		option.NewOption(
			"--win-console",
			"Creates a console launcher for the application, should be specified for application which requires console interactions",
			"--win-console",
			option.Win,
			true,
			true,
			[]string{},
			true),
	})
	linkOptionUI.addBackButton()
	linkOptionUI.addFinishButton()
}
