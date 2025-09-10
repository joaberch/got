package utils

import (
	"os"
	"path/filepath"
)

func WriteObject(path string, hash string, content []byte) error {
	objectPath := filepath.Join(".got", "objects", path, hash)
	if _, err := os.Stat(objectPath); os.IsNotExist(err) { //Handle duplicate
		return os.WriteFile(objectPath, content, 0644)
	}
	return nil //Duplicate can happen and shouldn't crash the app
}
