package cmd

import (
	"crypto/md5"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Status displays the status of the staging area
func Status() {
	stagingPath := filepath.Join(".got", "staging.csv")
	file, err := os.Open(stagingPath)
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, record := range records {
		oldHash := record[0]
		path := record[1]

		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("🔴 Removed : %s \n", path)
		}

		newHash := fmt.Sprintf("%x", md5.Sum(content))
		if oldHash == newHash {
			continue //Nothing changed
		} else {
			fmt.Printf("🟠 Modified : %s \n", path)
		}
	}
}
