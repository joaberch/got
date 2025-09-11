package utils

import (
	"log"
	"os"
	"path/filepath"
)

// CreateFilePath creates the mandatory file and folder for got to work after the init
func CreateFilePath(fullPath string, fileType string) {
	const dirPerm = 0755

	if fileType == "Folder" {
		err := os.MkdirAll(fullPath, dirPerm)
		if err != nil {
			log.Fatal(err)
		}
	} else if fileType == "File" {
		//Create the parent folder if it doesn't exist
		parentDir := filepath.Dir(fullPath)
		err := os.MkdirAll(parentDir, dirPerm)
		if err != nil {
			log.Fatal(err)
		}

		_, err = os.Create(fullPath)
		if err != nil {
			log.Fatal(err)
		}
	}
}
