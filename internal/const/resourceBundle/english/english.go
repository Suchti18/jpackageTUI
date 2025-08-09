package english

var (
	langPack map[string]string
)

func init() {
	langPack = map[string]string{
		"This PC":              "This PC",
		"Devices":              "Devices",
		"HomeDir":              "Home",
		"CurrentDir":           "Current Directory",
		"Modified":             "Modified",
		"Size":                 "Size",
		"Access Denied":        "Access Denied",
		"Favorites":            "Favorites",
		"Cancel":               "Cancel",
		"Accept":               "Accept",
		"Back":                 "Back",
		"Next":                 "Next",
		"Finish":               "Finish",
		"Include":              "Include <%s>?",
		"JpackageNotInstalled": "jpackage is not installed properly. Make sure to include it in the PATH variable. Use --force to skip this check (Checkout the README for more information)",
		"SkipjpackageCheck":    "Skipping jpackage check",
		"SuccessfulExit":       "Command finished successfully",
		"PressKeyToContinue":   "Press any key to continue",
	}
}

func GetString(id string) string {
	return langPack[id]
}
