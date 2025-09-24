package utils

import (
	"bufio"
	"fmt"
	"github.com/joaberch/got/internal/model"
	"os"
	"path/filepath"
	"strings"
)

// UpdateIndexedPaths appends the distinct file from the tree in the 'indexed_paths.csv'
func UpdateIndexedPaths(tree model.Tree) error {
	indexPath := filepath.Join(".got", "indexed_paths.csv")

	//Get the already indexed
	existing := make(map[string]string)
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
			parts := strings.Split(scanner.Text(), ",")
			if len(parts) == 2 {
				existing[parts[0]] = parts[1]
			}
		}
	}

	for _, entry := range tree.Entries {
		existing[entry.Name] = entry.Hash
	}

	file, err := os.OpenFile(indexPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
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
	for name, hash := range existing {
		_, err := writer.WriteString(fmt.Sprintf("%s,%s\n", name, hash))
		if err != nil {
			return fmt.Errorf("error writing to indexed path: %s", err)
		}
	}
	return writer.Flush()
}
