package german

var (
	langPack map[string]string
)

func init() {
	langPack = map[string]string{
		"This PC":  "Dieser PC",
		"Devices":  "Geräte",
		"HomeDir":  "Hauptverzeichnis",
		"Modified": "Bearbeitet am",
		"Size":     "Größe",
		"Cancel":   "Abbrechen",
		"Accept":   "Akzeptieren",
		"Back":     "Zurück",
		"Next":     "Weiter",
		"Finish":   "Fertig",
		"Include":  "<%s> hinzufügen?",
	}
}

func GetString(id string) string {
	return langPack[id]
}
