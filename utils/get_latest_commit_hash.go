package utils

import (
	"bufio"
	"os"
	"path/filepath"
)

// GetLatestCommitHash reads the file ".got/head" and returns its entire contents as a string.
// Each line in the file is appended with a trailing newline in the returned value.
// If the file cannot be opened or closed, the function logs the error and exits the process via log.Fatal.
// If the file is empty, an empty string is returned.
func GetLatestCommitHash() (string, error) {
	headPath := filepath.Join(".got", "head")
	file, err := os.Open(headPath)
	if err != nil {
		return "", err
	}
	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			err = closeErr
		}
	}()

	var content = ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		content += line + "\n"
	}
	return content, nil
}
