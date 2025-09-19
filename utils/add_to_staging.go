package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

// AddToStaging appends a CSV record to the .got/staging.csv file containing the provided
// hash and path (written in that order). The file is created if it does not already exist.
func AddToStaging(path string, hash string) error {
	stagingPath := filepath.Join(".got", "staging.csv")

	file, err := os.OpenFile(stagingPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open staging file at %s: %w", stagingPath, err)
	}
	defer func() {
		errClose := file.Close()
		if errClose != nil {
			err = fmt.Errorf("failed to close staging file at %s: %w", stagingPath, errClose)
		}
	}()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{path, hash})
	if err != nil {
		return fmt.Errorf("failed to write to staging file at %s: %w", stagingPath, err)
	}
	return nil
}
