package utils

import (
	"encoding/csv"
	"fmt"
	"github.com/joaberch/got/internal/model"
	"os"
)

// ReadStagingFile reads a two-column CSV at the given path and returns a model.Tree representing the staging area.
// Each CSV record is expected to have at least two fields: the object hash at index 0 and the file path at index 1.
// For each record a model.TreeEntry is appended with Name set to the path, Hash set to the hash, Mode "file", and Type "blob".
// If the file cannot be opened or the CSV cannot be read, the function calls log.Fatal and terminates the program.
//
// record[0] = path ro real file to include in the commit
// ReadStagingFile reads a two-column CSV from path and builds a model.Tree.
//
// The CSV is expected to have the file path in column 0 and the blob/hash in column 1.
// Each row with at least two fields produces a TreeEntry with Name set to the file path,
// Hash set to the blob name, Mode set to "file", and Type set to "blob". Rows with
// fewer than two fields are skipped.
//
// Returns a non-nil error if opening or reading the file fails. Any error encountered
// when closing the file is recorded in the deferred close but is not propagated to the caller.

func ReadStagingFile(path string) (model.Tree, error) {
	tree := model.Tree{}

	file, err := os.Open(path)
	if err != nil {
		return tree, fmt.Errorf("failed to open file: %w", err)
	}
	defer func(file *os.File) {
		errClose := file.Close()
		if errClose != nil {
			err = fmt.Errorf("failed to close file: %w", errClose)
		}
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return tree, fmt.Errorf("failed to read file: %w", err)
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
