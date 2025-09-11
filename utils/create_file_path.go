package utils

import (
	"log"
	"os"
	"path/filepath"
)

// CreateFilePath creates a directory or file at the given path.
// 
// If fileType is "Folder", it creates the directory tree rooted at fullPath.
// If fileType is "File", it ensures the parent directory exists and creates an empty file at fullPath.
// For any other fileType value the function does nothing.
// 
// Any filesystem error is fatal: the function calls log.Fatal on failure, terminating the program.
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
