package utils

import (
	"Got/internal/model"
	"encoding/csv"
	"log"
	"os"
)

// ReadStagingFile returns the staging area as a tree
func ReadStagingFile(path string) model.Tree {
	tree := model.Tree{}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, record := range records {
		hash := record[0]
		recordPath := record[1]

		tree.Entries = append(tree.Entries, model.TreeEntry{
			Name: recordPath,
			Hash: hash,
			Mode: "file",
			Type: "blob",
		})
	}
	return tree
}
