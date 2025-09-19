package utils

import (
	"os"
	"path/filepath"
	"strings"
)

// GetLatestCommitHash reads the repository head file at ".got/head" and returns its contents
// with surrounding whitespace trimmed.
// It returns the trimmed commit hash and any error encountered while reading the file.
// If the file is empty or contains only whitespace, an empty string is returned.
func GetLatestCommitHash() (string, error) {
	headPath := filepath.Join(".got", "head")
	data, err := os.ReadFile(headPath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}
