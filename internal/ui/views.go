package ui

import (
	"github.com/nils/jpackageTUI/internal/const/args"
	"github.com/nils/jpackageTUI/internal/option"
	"runtime"
)

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
	genericOptionUI    *OptionUI
	linkOptionUI       *OptionUI
	macOptionUI        *OptionUI
	linuxOptionUI      *OptionUI
	windowsOptionUI    *OptionUI
	appPackageOptionUI *OptionUI

	typeOptions = []string{appImage, exe, msi, rpm, deb, pkg, dmg}
)

func LoadAll() {
	loadGeneric()
	UI.optionUIs = append(UI.optionUIs, genericOptionUI)

	// Add additional options if the add argument is set
	if args.HasArg(args.AllArg) {
		loadCreatingAppPackageOptions()
		UI.optionUIs = append(UI.optionUIs, appPackageOptionUI)

		if runtime.GOOS == "windows" {
			loadWindowsOptions()
			UI.optionUIs = append(UI.optionUIs, windowsOptionUI)
		} else if runtime.GOOS == "darwin" {
			loadMacOptions()
			UI.optionUIs = append(UI.optionUIs, macOptionUI)
		} else if runtime.GOOS == "linux" {
			loadLinuxOptions()
			UI.optionUIs = append(UI.optionUIs, linuxOptionUI)
		}
	}

	loadLinks()
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
			"Modules",
			"A comma (\",\") separated list of modules to add\n\nThis module list, along with the main module (if specified) will be passed to jlink as the --add-module argument. If not specified, either just the main module (if --module is specified), or the default set of modules (if --main-jar is specified) are used.",
			"--add-modules",
			option.CrossPlatform,
			true,
			true,
			[]string{},
			false,
			option.Text),
		option.NewOption(
			"jlink options",
			"A space separated list of options to pass to jlink\n\nIf not specified, defaults to \"--strip-native-commands --strip-debug --no-man-pages --no-header-files\"",
			"--jlink-options",
			option.CrossPlatform,
			true,
			true,
			[]string{},
			false,
			option.Text),
		option.NewOption(
			"Runtime image",
			"Path of the predefined runtime image that will be copied into the application image (absolute path or relative to the current directory)\n\nIf --runtime-image is not specified, jpackage will run jlink to create the runtime image using options specified by --jlink-options.",
			"--runtime-image",
			option.CrossPlatform,
			true,
			true,
			[]string{},
			false,
			option.Folder),
		option.NewOption(
			"Module",
			"The main module (and optionally main class) of the application",
			"--module",
			option.CrossPlatform,
			true,
			true,
			[]string{},
			false,
			option.Text),
		option.NewOption(
			"Module path",
			"Each path is either a directory of modules or the path to a modular jar, and is absolute or relative to the current directory.",
			"--module-path",
			option.CrossPlatform,
			true,
			true,
			[]string{},
			false,
			option.Folder),
		option.NewOption(
			"Input files",
			"Path of the input directory that contains the files to be packaged\n\nAll files in the input directory will be packaged into the application image.",
			"--input",
			option.CrossPlatform,
			true,
			true,
			[]string{},
			false,
			option.Folder),
		option.NewOption(
			"Main class arguments",
			"Command line arguments to pass to the main class if no command line arguments are given to the launcher",
			"--arguments",
			option.CrossPlatform,
			true,
			true,
			[]string{},
			false,
			option.Text),
		option.NewOption(
			"JVM options",
			"Options to pass to the Java runtime",
			"--java-options",
			option.CrossPlatform,
			true,
			true,
			[]string{},
			false,
			option.Text),
		option.NewOption(
			"Add launcher",
			"Name of launcher, and a path to a Properties file that contains a list of key, value pairs (absolute path or relative to the current directory)\n\nThe keys \"module\", \"main-jar\", \"main-class\", \"arguments\", \"java-options\", \"app-version\", \"icon\", \"linux-app-category\", \"linux-app-release\", and \"win-console\" can be used.\n\nThese options are added to, or used to overwrite, the original command line options to build an additional alternative launcher. The main application launcher will be built from the command line options. Additional alternative launchers can be built using this option, and this option can be used multiple times to build multiple additional launchers.",
			"--add-launcher",
			option.CrossPlatform,
			true,
			true,
			[]string{},
			false,
			option.Text),
		option.NewOption(
			"Main class",
			"Qualified name of the application main class to execute\n\nThis option can only be used if --main-jar is specified.",
			"--main-class",
			option.CrossPlatform,
			true,
			true,
			[]string{},
			false,
			option.Text),
		option.NewOption(
			"Main jar",
			"The main JAR of the application; containing the main class (specified as a path relative to the input path)\n\nEither --module or --main-jar option can be specified but not both.",
			"--main-jar",
			option.CrossPlatform,
			true,
			true,
			[]string{},
			false,
			option.File),
		option.NewOption(
			"--win-console",
			"Creates a console launcher for the application, should be specified for application which requires console interactions (available only when running on Windows)",
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
	macOptionUI = NewOptionsUI([]*option.Option{
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
	macOptionUI.addBackButton()
	macOptionUI.addNextButton()
}

func loadCreatingAppPackageOptions() {
	appPackageOptionUI = NewOptionsUI([]*option.Option{
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
	appPackageOptionUI.addBackButton()
	appPackageOptionUI.addNextButton()
}

func loadWindowsOptions() {
	windowsOptionUI = NewOptionsUI([]*option.Option{
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
		option.NewOption(
			"Help url",
			"URL where user can obtain further information or technical support",
			"--win-help-url",
			option.Win,
			true,
			true,
			[]string{},
			false,
			option.Text),
		option.NewOption(
			"Start Menu shortcut",
			"Request to add a Start Menu shortcut for this application",
			"--win-menu",
			option.Win,
			true,
			true,
			[]string{},
			true,
			option.Text),
		option.NewOption(
			"Start Menu group",
			"Start Menu group this application is placed in",
			"--win-menu-group",
			option.Win,
			true,
			true,
			[]string{},
			false,
			option.Text),
		option.NewOption(
			"Per user install",
			"Request to perform an install on a per-user basis",
			"--win-per-user-install",
			option.Win,
			true,
			true,
			[]string{},
			true,
			option.Text),
		option.NewOption(
			"Desktop shortcut",
			"Request to create a desktop shortcut for this application",
			"--win-shortcut",
			option.Win,
			true,
			true,
			[]string{},
			true,
			option.Text),
		option.NewOption(
			"Shortcut prompt",
			"Adds a dialog to enable the user to choose if shortcuts will be created by installer",
			"--win-shortcut-prompt",
			option.Win,
			true,
			true,
			[]string{},
			true,
			option.Text),
		option.NewOption(
			"Update URL",
			"URL of available application update information",
			"--win-update-url",
			option.Win,
			true,
			true,
			[]string{},
			false,
			option.Text),
		option.NewOption(
			"Upgrade UUID",
			"UUID associated with upgrades for this package",
			"--win-upgrade-uuid",
			option.Win,
			true,
			true,
			[]string{},
			false,
			option.Text),
	}, "Windows platform options")
	windowsOptionUI.addBackButton()
	windowsOptionUI.addNextButton()
}

func loadLinuxOptions() {
	linuxOptionUI = NewOptionsUI([]*option.Option{
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
		option.NewOption(
			"Linux maintainer",
			"Maintainer's email address for .deb bundle",
			"--linux-deb-maintainer",
			option.Linux,
			true,
			true,
			[]string{},
			false,
			option.Text),
		option.NewOption(
			"Menu group",
			"Menu group this application is placed in",
			"--linux-menu-group",
			option.Linux,
			true,
			true,
			[]string{},
			false,
			option.Text),
		option.NewOption(
			"Required packages",
			"Required packages or capabilities for the application",
			"--linux-package-deps",
			option.Linux,
			true,
			true,
			[]string{},
			true,
			option.Text),
		option.NewOption(
			"License type",
			"Type of the license (\"License: <value>\" of the RPM .spec)",
			"--linux-rpm-license-type",
			option.Linux,
			true,
			true,
			[]string{},
			false,
			option.Text),
		option.NewOption(
			"Release value",
			"Release value of the RPM <name>.spec file or Debian revision value of the DEB control file",
			"--linux-app-release",
			option.Linux,
			true,
			true,
			[]string{},
			false,
			option.Text),
		option.NewOption(
			"Group",
			"Group value of the RPM <name>.spec file or Section value of DEB control file",
			"--linux-app-category",
			option.Linux,
			true,
			true,
			[]string{},
			false,
			option.Text),
		option.NewOption(
			"Create a shortcut",
			"Creates a shortcut for the application.",
			"--linux-shortcut",
			option.Linux,
			true,
			true,
			[]string{},
			true,
			option.Text),
	}, "Linux platform options")
	linuxOptionUI.addBackButton()
	linuxOptionUI.addNextButton()
}
