package utils

import (
	"os"
	"path/filepath"
)

// WriteObject writes content to ".got/objects/{path}/{hash}" for the object type given by path
// (e.g., "blobs", "trees", "commits"). The file is created only if it does not already exist;
// if the target file exists the function returns nil. Only errors from writing the file are
// returned. Note: parent directories are not created by this function and must exist beforehand.
func WriteObject(path string, hash string, content []byte) error {
	objectPath := filepath.Join(".got", "objects", path, hash)
	if _, err := os.Stat(objectPath); os.IsNotExist(err) { //Handle duplicate
		return os.WriteFile(objectPath, content, 0644)
	}
	return nil //Duplicate can happen and shouldn't crash the app
}
