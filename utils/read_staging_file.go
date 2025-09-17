package utils

import (
	"encoding/csv"
	"github.com/joaberch/got/internal/model"
	"log"
	"os"
)

// ReadStagingFile reads a two-column CSV at the given path and returns a model.Tree representing the staging area.
// Each CSV record is expected to have at least two fields: the object hash at index 0 and the file path at index 1.
// For each record a model.TreeEntry is appended with Name set to the path, Hash set to the hash, Mode "file", and Type "blob".
// If the file cannot be opened or the CSV cannot be read, the function calls log.Fatal and terminates the program.
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
