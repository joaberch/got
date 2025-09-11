package utils

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
)

// AddToStaging adds an entry in the staging.csv file
func AddToStaging(path string, hash string) {
	stagingPath := filepath.Join(".got", "staging.csv")

	file, err := os.OpenFile(stagingPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{hash, path})
	if err != nil {
		log.Fatal(err)
	}
}
