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

func loadMacOptions() {
	linkOptionUI = NewOptionsUI([]*option.Option{
		option.NewOption(
			"Mac package identifier",
			"An identifier that uniquely identifies the application for macOS\n\nDefaults to the the main class name.\n\nMay only use alphanumeric (A-Z,a-z,0-9), hyphen (-), and period (.) characters.\n",
			"--mac-package-identifier",
			option.Mac,
			true,
			true,
			[]string{},
			false,
			option.Text),
		option.NewOption(
			"Mac package name",
			"Name of the application as it appears in the Menu Bar\n\nThis can be different from the application name.\n\nThis name must be less than 16 characters long and be suitable for displaying in the menu bar and the application Info window. Defaults to the application name.\n",
			"--mac-package-name",
			option.Mac,
			true,
			true,
			[]string{},
			false,
			option.Text),
		option.NewOption(
			"Mac package signing prefix",
			"When signing the application package, this value is prefixed to all components that need to be signed that don't have an existing package identifier.",
			"--mac-package-signing-prefix",
			option.Mac,
			true,
			true,
			[]string{},
			false,
			option.Text),
		option.NewOption(
			"Mac sign",
			"Request that the bundle be signed.",
			"--mac-sign",
			option.Mac,
			true,
			true,
			[]string{},
			true,
			option.Text),
		option.NewOption(
			"Mac signing keychain",
			"Name of the keychain to search for the signing identity\n\nIf not specified, the standard keychains are used.\n",
			"--mac-signing-keychain",
			option.Mac,
			true,
			true,
			[]string{},
			false,
			option.Text),
		option.NewOption(
			"Mac signing key user name",
			"Team or user name portion in Apple signing identities",
			"--mac-signing-key-user-name",
			option.Mac,
			true,
			true,
			[]string{},
			false,
			option.Text),
		option.NewOption(
			"Mac App store app",
			"Indicates that the jpackage output is intended for the Mac App Store.",
			"--mac-app-store",
			option.Mac,
			true,
			true,
			[]string{},
			true,
			option.Text),
		option.NewOption(
			"Mac entitlements",
			"Path to file containing entitlements to use when signing executables and libraries in the bundle",
			"--mac-entitlements",
			option.Mac,
			true,
			true,
			[]string{},
			false,
			option.File),
		option.NewOption(
			"Mac app category",
			"String used to construct LSApplicationCategoryType in application plist\n\nThe default value is \"utilities\".\n",
			"--mac-app-category",
			option.Mac,
			true,
			true,
			[]string{},
			false,
			option.Text),
	}, "macOS platform options")
}

func loadCreatingAppPackageOptions() {
	linkOptionUI = NewOptionsUI([]*option.Option{
		option.NewOption(
			"About URL",
			"URL of the application's home page",
			"--about-url",
			option.CrossPlatform,
			true,
			true,
			[]string{},
			false,
			option.Text),
		option.NewOption(
			"App image location",
			"Location of the predefined application image that is used to build an installable package\n\nSee create-app-image mode options to create the application image.",
			"--app-image",
			option.CrossPlatform,
			true,
			true,
			[]string{},
			false,
			option.Folder),
		option.NewOption(
			"File associations properties",
			"Path to a Properties file that contains list of key, value pairs (absolute path or relative to the current directory)\n\nThe keys \"extension\", \"mime-type\", \"icon\", and \"description\" can be used to describe the association.",
			"--file-associations",
			option.CrossPlatform,
			true,
			false,
			[]string{},
			false,
			option.File),
		option.NewOption(
			"Install directory",
			"Absolute path of the installation directory of the application (on macos or linux), or relative sub-path of the installation directory such as \"Program Files\" or \"AppData\" (on Windows)",
			"--install-dir",
			option.CrossPlatform,
			true,
			false,
			[]string{},
			false,
			option.Folder),
		option.NewOption(
			"License file",
			"Path to the license file",
			"--license-file",
			option.CrossPlatform,
			true,
			false,
			[]string{},
			false,
			option.File),
		option.NewOption(
			"Override resources",
			"Path to override jpackage resources\n\nIcons, template files, and other resources of jpackage can be over-ridden by adding replacement resources to this directory.",
			"--resource-dir",
			option.CrossPlatform,
			true,
			false,
			[]string{},
			false,
			option.Folder),
		option.NewOption(
			"Predefined runtime image",
			"Path of the predefined runtime image to install\n\nOption is required when creating a runtime installer.",
			"--runtime-image",
			option.CrossPlatform,
			true,
			false,
			[]string{},
			false,
			option.Folder),
	}, "Options for creating the application package")
}

func loadWindowsOptions() {
	linkOptionUI = NewOptionsUI([]*option.Option{
		option.NewOption(
			"Add directory dialog",
			"Adds a dialog to enable the user to choose a directory in which the product is installed.",
			"--win-dir-chooser",
			option.Win,
			true,
			true,
			[]string{},
			true,
			option.Text),
	}, "Windows platform options")
}

func loadLinuxOptions() {
	linkOptionUI = NewOptionsUI([]*option.Option{
		option.NewOption(
			"Linux package name",
			"Name for Linux package\n\nDefaults to the application name.",
			"--linux-package-name",
			option.Linux,
			true,
			true,
			[]string{},
			false,
			option.Text),
	}, "Linux platform options")
}
