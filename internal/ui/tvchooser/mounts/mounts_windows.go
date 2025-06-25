package mounts

import (
	"fmt"
	"golang.org/x/sys/windows"
)

func GetDrivesLetters() ([]string, error) {
	drivesBitmask, err := windows.GetLogicalDrives()
	if err != nil {
		return nil, fmt.Errorf("error while retrieving Drives <%w>", err)
	}

	var drives []string
	for i := 0; i < 26; i++ { // A-Z
		if drivesBitmask&(1<<i) != 0 {
			drive := fmt.Sprintf("%c", 'A'+i)
			drives = append(drives, drive)
		}
	}

	return drives, nil
}
