package utils

import (
	"fmt"
	"github.com/joaberch/got/internal/model"
	"os"
)

// CreateBlobs reads each file referenced by tree.Entries and writes its content as a blob object named by the entry's Hash into the objects/blobs store.
// If reading or writing any entry fails, the function logs the error and exits the program.
// CreateBlobs reads each entry in the tree and stores its file contents as a blob.
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
