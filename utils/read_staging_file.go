package utils

import (
	"encoding/csv"
	"github.com/joaberch/got/internal/model"
	"os"
)

// ReadStagingFile reads a two-column CSV at the given path and returns a model.Tree representing the staging area.
// Each CSV record is expected to have at least two fields: the object hash at index 0 and the file path at index 1.
// For each record a model.TreeEntry is appended with Name set to the path, Hash set to the hash, Mode "file", and Type "blob".
// If the file cannot be opened or the CSV cannot be read, the function calls log.Fatal and terminates the program.
//
// record[0] = path ro real file to include in the commit
// record[1] = blob name (hash) can be fictive

func ReadStagingFile(path string) (model.Tree, error) {
	tree := model.Tree{}

	file, err := os.Open(path)
	if err != nil {
		return tree, err
	}
	defer func(file *os.File) {
		errClose := file.Close()
		if errClose != nil {
			err = errClose
		}
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return tree, err
	}
	for _, record := range records {
		if len(record) < 2 {
			continue //Skip
		}
		filePath := record[0]
		blobName := record[1]

		tree.Entries = append(tree.Entries, model.TreeEntry{
			Name: filePath,
			Hash: blobName,
			Mode: "file",
			Type: "blob",
		})
	}
	return tree, nil
}
