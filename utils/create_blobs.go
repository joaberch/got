package utils

import (
	"github.com/joaberch/got/internal/model"
	"os"
)

// CreateBlobs reads each file referenced by tree.Entries and writes its content as a blob object named by the entry's Hash into the objects/blobs store.
// If reading or writing any entry fails, the function logs the error and exits the program.
// Each entry is expected to provide Name (filesystem path) and Hash (blob identifier).
func CreateBlobs(tree model.Tree) error {
	for _, entry := range tree.Entries {
		content, err := os.ReadFile(entry.Name)
		if err != nil {
			return err
		}

		blobHash := entry.Hash

		err = WriteObject("blobs", blobHash, content)
		if err != nil {
			return err
		}
	}
	return nil
}
