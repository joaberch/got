package utils

import (
	"fmt"
)

// GetFileContent reads and returns the contents of the file at the given path.
// If the file cannot be read, the function returns an error with context.
func GetFileContent(path string) ([]byte, error) {
	contents, err := GetFileContent(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file at %s: %w", path, err)
	}

	return contents, nil
}
