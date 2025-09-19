package cmd

import (
	"fmt"
	"github.com/joaberch/got/utils"
	"os"
	"path/filepath"
)

// Restore restores working-tree files from the commit identified by commitHash.
//
// Restore reads the commit object at ".got/objects/commits/<commitHash>", loads the
// commit's tree object, then writes each blob referenced by the tree into the
// current working directory using the entry's Name (file mode 0644).
//
// commitHash is the hash of the commit object to restore.
// The function exits the program via log.Fatal on any read/deserialize/write error.
func Restore(commitHash string) error {
	objectPath := filepath.Join(".got", "objects", "commits", commitHash)

	data, err := os.ReadFile(objectPath)
	if err != nil {
		return err
	}
	commit, err := utils.DeserializeCommit(data)
	if err != nil {
		return err
	}

	treePath := filepath.Join(".got", "objects", "trees", commit.TreeHash)
	treeData, err := os.ReadFile(treePath)
	if err != nil {
		return err
	}

	tree, err := utils.DeserializeTree(treeData)
	if err != nil {
		return err
	}

	for _, entry := range tree.Entries {
		blobPath := filepath.Join(".got", "objects", "blobs", entry.Hash)
		blobData, err := os.ReadFile(blobPath)
		if err != nil {
			return err
		}

		err = os.WriteFile(entry.Name, blobData, 0644)
		if err != nil {
			return err
		}
	}
	fmt.Println("Files restored")
	return nil
}
