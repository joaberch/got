package utils

import (
	"fmt"
	"github.com/joaberch/got/internal/model"
	"os"
)

// CreateBlobs reads the file at entry.Name and writes its contents to the "blobs" object store using entry.Hash as the object name.
// It returns a wrapped error if any file read or object write fails; on success it returns nil.
func CreateBlobs(tree model.Tree) error {
	for _, entry := range tree.Entries {
		//entry.Name = path to real file
		//entry.Hash = blob name to create (file name)
		content, err := os.ReadFile(entry.Name)
		if err != nil {
			return fmt.Errorf("failed to read file at %s: %w", entry.Name, err)
		}

		err = WriteObject("blobs", entry.Hash, content)
		if err != nil {
			return fmt.Errorf("failed to write to blobs file at %s: %w", entry.Name, err)
		}
	}
	return nil
}
