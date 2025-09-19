package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetLatestCommitHash reads the file ".got/head" and returns its entire contents as a string.
// Each line in the file is appended with a trailing newline in the returned value.
// If the file cannot be opened or closed, the function logs the error and exits the process via log.Fatal.
// If the file is empty, an empty string is returned.
func GetLatestCommitHash() (string, error) {
	headPath := filepath.Join(".got", "head")
	file, err := os.Open(headPath)
	if err != nil {
		return "", fmt.Errorf("failed to open head file at %s: %w", headPath, err)
	}
	defer func() {
		errClose := file.Close()
		if errClose != nil {
			err = fmt.Errorf("failed to close head file at %s: %w", headPath, errClose)
		}
	}()

	var content = ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		content += line + ""
	}
	return strings.TrimSpace(content), nil
}
