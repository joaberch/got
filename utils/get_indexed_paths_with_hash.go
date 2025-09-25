package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetIndexedPathsWithHash returns a map of each file processed with their latest hash
func GetIndexedPathsWithHash() (map[string]string, error) {
	indexPath := filepath.Join(".got", "indexed_paths.csv")
	file, err := os.Open(indexPath)
	if err != nil {
		return nil, fmt.Errorf("error opening indexed path: %s", err)
	}
	defer func() {
		errClose := file.Close()
		if errClose != nil {
			err = errClose
		}
	}()

	paths := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) == 2 {
			paths[parts[0]] = parts[1]
		}
	}
	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading indexed paths: %s", err)
	}
	return paths, nil
}
