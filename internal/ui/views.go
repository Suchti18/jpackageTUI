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
			false,
			option.Text),
		option.NewOption(
			"App Version",
			"Version of the application and/or package",
			"--app-version",
			option.CrossPlatform,
			true,
			true,
			[]string{},
			false,
			option.Text),
		option.NewOption(
			"Copyright",
			"Copyright for the application",
			"--copyright",
			option.CrossPlatform,
			true,
			true,
			[]string{},
			false,
			option.Text),
		option.NewOption(
			"Vendor",
			"Vendor of the application",
			"--vendor",
			option.CrossPlatform,
			true,
			true,
			[]string{},
			false,
			option.Text),
		option.NewOption(
			"Description",
			"Description of the application",
			"--description",
			option.CrossPlatform,
			true,
			true,
			[]string{},
			false,
			option.Text),
		option.NewOption(
			"Icon",
			"Path of the icon of the application package",
			"--icon",
			option.CrossPlatform,
			true,
			true,
			[]string{},
			false,
			option.File),
		option.NewOption(
			"Type",
			"The type of package to create\n\nValid values are: {\"app-image\", \"exe\", \"msi\", \"rpm\", \"deb\", \"pkg\", \"dmg\"}\n\nIf this option is not specified a platform dependent default type will be created.\n",
			"--type",
			option.CrossPlatform,
			true,
			true,
			typeOptions,
			false,
			option.Text),
		option.NewOption(
			"Temp",
			"Path of a new or empty directory used to create temporary files\n\nIf specified, the temp dir will not be removed upon the task completion and must be removed manually.\n\nIf not specified, a temporary directory will be created and removed upon the task completion.\n",
			"--temp",
			option.CrossPlatform,
			true,
			true,
			[]string{},
			false,
			option.Folder),
		option.NewOption(
			"Destination",
			"Path where generated output file is placed\n\nDefaults to the current working directory.",
			"--dest",
			option.CrossPlatform,
			true,
			true,
			[]string{},
			false,
			option.Folder),
	}, "Generic Options")
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
			false,
			option.Text),
		option.NewOption(
			"Module path",
			"Each path is either a directory of modules or the path to a modular jar, and is absolute or relative to the current directory.",
			"--module-path",
			option.CrossPlatform,
			false,
			true,
			[]string{},
			false,
			option.Folder),
		option.NewOption(
			"--win-console",
			"Creates a console launcher for the application, should be specified for application which requires console interactions",
			"--win-console",
			option.Win,
			true,
			true,
			[]string{},
			true,
			option.Text),
	}, "Options for creating the runtime image")
	linkOptionUI.addBackButton()
	linkOptionUI.addFinishButton()
}
