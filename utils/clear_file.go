package utils

import "os"

// ClearFile clears the file given
func ClearFile(path string) error {
	return os.WriteFile(path, []byte(""), 0644)
}
