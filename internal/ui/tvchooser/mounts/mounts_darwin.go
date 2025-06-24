package mounts

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetDrivesLetters() ([]string, error) {
	volumesPath := "/Volumes"

	// Check if the folder exist
	if _, err := os.Stat(volumesPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("<%s> does not exist", volumesPath)
	}

	// Reading all connected drives
	files, err := os.ReadDir(volumesPath)
	if err != nil {
		return nil, fmt.Errorf("error while reading folder <%s>: <%w>", volumesPath, err)
	}

	var drives []string
	for _, file := range files {
		if file.IsDir() {
			drives = append(drives, filepath.Join(volumesPath, file.Name()))
		}
	}

	return drives, nil
}
