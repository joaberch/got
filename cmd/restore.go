package cmd

import (
	"Got/utils"
	"fmt"
	"log"
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
func Restore(commitHash string) {
	objectPath := filepath.Join(".got", "objects", "commits", commitHash)

	data, err := os.ReadFile(objectPath)
	if err != nil {
		log.Fatal(err)
	}

	commit, err := utils.DeserializeCommit(data)
	if err != nil {
		log.Fatal(err)
	}

	treePath := filepath.Join(".got", "objects", "trees", commit.TreeHash)
	treeData, err := os.ReadFile(treePath)
	if err != nil {
		log.Fatal(err)
	}

	tree, err := utils.DeserializeTree(treeData)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range tree.Entries {
		blobPath := filepath.Join(".got", "objects", "blobs", entry.Hash)
		blobData, err := os.ReadFile(blobPath)
		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile(entry.Name, blobData, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Files restored")
}
