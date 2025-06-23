//go:build darwin
// +build darwin

package mounts

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetDrivesLetters() ([]string, error) {
	volumesPath := "/Volumes"

	// Überprüfen, ob das Verzeichnis existiert
	if _, err := os.Stat(volumesPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("Verzeichnis %s existiert nicht", volumesPath)
	}

	// Liste der Laufwerke abrufen
	files, err := os.ReadDir(volumesPath)
	if err != nil {
		return nil, fmt.Errorf("Fehler beim Lesen des Verzeichnisses %s: %w", volumesPath, err)
	}

	var drives []string
	for _, file := range files {
		if file.IsDir() {
			drives = append(drives, filepath.Join(volumesPath, file.Name()))
		}
	}

	return drives, nil
}
