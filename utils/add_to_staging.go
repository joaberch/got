package utils

import (
	"encoding/csv"
	"os"
	"path/filepath"
)

// AddToStaging appends a CSV record to the .got/staging.csv file containing the provided
// hash and path (written in that order). The file is created if it does not already exist.
func AddToStaging(path string, hash string) error {
	stagingPath := filepath.Join(".got", "staging.csv")

	file, err := os.OpenFile(stagingPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func() {
		errClose := file.Close()
		if errClose != nil {
			err = errClose
		}
	}()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{path, hash})
	if err != nil {
		return err
	}
	return nil
}
