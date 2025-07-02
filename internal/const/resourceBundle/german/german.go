package german

var (
	langPack map[string]string
)

func init() {
	langPack = map[string]string{
		"This PC":              "Dieser PC",
		"Devices":              "Geräte",
		"HomeDir":              "Hauptverzeichnis",
		"CurrentDir":           "Aktuelles Verzeichnis",
		"Modified":             "Bearbeitet am",
		"Size":                 "Größe",
		"Access Denied":        "Zugriff verweigert",
		"Favorites":            "Favoriten",
		"Cancel":               "Abbrechen",
		"Accept":               "Akzeptieren",
		"Back":                 "Zurück",
		"Next":                 "Weiter",
		"Finish":               "Fertig",
		"Include":              "<%s> hinzufügen?",
		"JpackageNotInstalled": "jpackage wurde nicht korrekt in die PATH-Variable eingebunden.",
		"SkipjpackageCheck":    "jpackage Check übersprungen",
		"SuccessfulExit":       "jpackage wurde mit deinen Optionen ausgeführt",
	}
}

func GetString(id string) string {
	return langPack[id]
}
