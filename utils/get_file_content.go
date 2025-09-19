package utils

import (
	"fmt"
	"os"
)

// GetFileContent reads and returns the contents of the file at the given path.
// If the file cannot be read the function logs the error and calls log.Fatal,
// includes the file path.
func GetFileContent(path string) ([]byte, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file at %s: %w", path, err)
	}

	return contents, nil
}
