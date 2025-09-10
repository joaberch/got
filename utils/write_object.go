package utils

import (
	"os"
	"path/filepath"
)

func WriteObject(hash string, content []byte) error {
	objectPath := filepath.Join(".got", "objects", hash)
	return os.WriteFile(objectPath, content, 0644)
}
