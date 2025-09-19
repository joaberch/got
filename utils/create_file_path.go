package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// CreateFilePath creates a directory or file at the given path.
//
// If fileType is "Folder", it creates the directory tree rooted at fullPath.
// If fileType is "File", it ensures the parent directory exists and creates an empty file at fullPath.
// For any other fileType value the function does nothing.
//
// CreateFilePath creates filesystem entries at fullPath according to fileType.
// 
// If fileType == "Folder", it creates the directory tree at fullPath (mkdir -p semantics).
// If fileType == "File", it ensures the parent directory exists and creates an empty file at fullPath.
// For any other fileType value the function does nothing.
// 
// On failure the function returns a non-nil error describing the filesystem error (directory creation,
// file creation, or file close failure). On success it returns nil.
func CreateFilePath(fullPath string, fileType string) error {
	const dirPerm = 0755

	if fileType == "Folder" {
		err := os.MkdirAll(fullPath, dirPerm)
		if err != nil {
			return fmt.Errorf("failed to create directory at %s: %w", fullPath, err)
		}
	} else if fileType == "File" {
		//Create the parent folder if it doesn't exist
		parentDir := filepath.Dir(fullPath)
		err := os.MkdirAll(parentDir, dirPerm)
		if err != nil {
			return fmt.Errorf("failed to create directory at %s: %w", parentDir, err)
		}

		file, err := os.Create(fullPath)
		if err != nil {
			return fmt.Errorf("failed to create file at %s: %w", fullPath, err)
		}
		defer func() {
			closeErr := file.Close()
			if err != nil {
				err = fmt.Errorf("failed to close file at %s: %w", fullPath, closeErr)
			}
		}()
	}
	return nil
}
