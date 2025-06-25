package english

var (
	langPack map[string]string
)

func init() {
	langPack = map[string]string{
		"This PC":       "This PC",
		"Devices":       "Devices",
		"HomeDir":       "Home",
		"CurrentDir":    "Current Directory",
		"Modified":      "Modified",
		"Size":          "Size",
		"Access Denied": "Access Denied",
		"Favorites":     "Favorites",
		"Cancel":        "Cancel",
		"Accept":        "Accept",
		"Back":          "Back",
		"Next":          "Next",
		"Finish":        "Finish",
		"Include":       "Include <%s>?",
	}
}

func GetString(id string) string {
	return langPack[id]
}
