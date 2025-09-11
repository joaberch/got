package utils

import (
	"os"
	"path/filepath"
)

// WriteObject creates an object file of the type given as 'path' (blobs/trees/commits) with the file name 'hash' and the content 'content'
func WriteObject(path string, hash string, content []byte) error {
	objectPath := filepath.Join(".got", "objects", path, hash)
	if _, err := os.Stat(objectPath); os.IsNotExist(err) { //Handle duplicate
		return os.WriteFile(objectPath, content, 0644)
	}
	return nil //Duplicate can happen and shouldn't crash the app
}
