package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetStagedEntries fetch the relative path of the files added to the staging area
func GetStagedEntries() ([]string, error) {
	stagingPath := filepath.Join(".got", "staging.csv")
	file, err := os.Open(stagingPath)
	if err != nil {
		return nil, fmt.Errorf("error opening staging file: %v", err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Printf("error closing staging file: %v", err)
		}
	}()

	var entries []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) >= 1 {
			entries = append(entries, parts[0])
		}
	}

	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading staging file: %v", err)
	}

	return entries, nil
}
