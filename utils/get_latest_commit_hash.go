package utils

import (
	"path/filepath"
	"strings"
)

// GetLatestCommitHash reads the file ".got/head" and returns its contents as a trimmed string.
// If the file cannot be read, an error is returned.
// If the file is empty, an empty string is returned with no error.
func GetLatestCommitHash() (string, error) {
	headPath := filepath.Join(".got", "head")
	data, err := GetFileContent(headPath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}
