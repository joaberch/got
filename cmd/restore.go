package cmd

import (
	"fmt"
	"github.com/joaberch/got/utils"
	"os"
	"path/filepath"
)

// Restore restores working-tree files from the commit identified by commitHash.
// 
// It reads the commit object at ".got/objects/commits/<commitHash>", deserializes it
// to obtain the root tree hash, reads and deserializes the tree object at
// ".got/objects/trees/<treeHash>", then writes each blob found in
// ".got/objects/blobs/<blobHash>" to the working directory using the entry's Name
// with file mode 0644.
//
// commitHash is the hash of the commit object to restore.
//
// Returns an error if any read, deserialization, or write operation fails. On
// success the function prints "Files restored" and returns nil.
func Restore(commitHash string) error {
	objectPath := filepath.Join(".got", "objects", "commits", commitHash)

	data, err := os.ReadFile(objectPath)
	if err != nil {
		return fmt.Errorf("error reading file %s: %s", objectPath, err)
	}
	commit, err := utils.DeserializeCommit(data)
	if err != nil {
		return fmt.Errorf("error deserializing commit: %s", err)
	}

	treePath := filepath.Join(".got", "objects", "trees", commit.TreeHash)
	treeData, err := os.ReadFile(treePath)
	if err != nil {
		return fmt.Errorf("error reading tree file %s: %s", treePath, err)
	}

	tree, err := utils.DeserializeTree(treeData)
	if err != nil {
		return fmt.Errorf("error deserializing tree %s: %s", treePath, err)
	}

	for _, entry := range tree.Entries {
		blobPath := filepath.Join(".got", "objects", "blobs", entry.Hash)
		blobData, err := os.ReadFile(blobPath)
		if err != nil {
			return fmt.Errorf("error reading blob file %s: %s", blobPath, err)
		}

		err = os.WriteFile(entry.Name, blobData, 0644)
		if err != nil {
			return fmt.Errorf("error writing blob file %s: %s", blobPath, err)
		}
	}
	fmt.Println("Files restored")
	return nil
}
