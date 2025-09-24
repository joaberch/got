package utils

import (
	"bufio"
	"fmt"
	"github.com/joaberch/got/internal/model"
	"os"
	"path/filepath"
)

// UpdateIndexedPaths appends the distinct file from the tree in the 'indexed_paths.csv'
func UpdateIndexedPaths(tree model.Tree) error {
	indexPath := filepath.Join(".got", "indexed_paths.csv")

	//Get the already indexed
	existing := make(map[string]bool)
	if _, err := os.Stat(indexPath); err == nil {
		file, err := os.Open(indexPath)
		if err != nil {
			return fmt.Errorf("error opening indexed path: %s", err)
		}
		defer func() {
			closeErr := file.Close()
			if closeErr != nil {
				err = closeErr
			}
		}()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			existing[scanner.Text()] = true
		}
	}

	file, err := os.OpenFile(indexPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("error opening indexed path: %s", err)
	}
	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			err = closeErr
		}
	}()

	writer := bufio.NewWriter(file)
	for _, entry := range tree.Entries {
		if !existing[entry.Name] {
			_, err = writer.WriteString(entry.Name + "\n")
			if err != nil {
				return fmt.Errorf("error writing to indexed path: %s", err)
			}
		}
	}
	return writer.Flush()
}
