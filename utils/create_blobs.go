package utils

import (
	"Got/internal/model"
	"log"
	"os"
)

func CreateBlobs(tree model.Tree) {
	for _, entry := range tree.Entries {
		content, err := os.ReadFile(entry.Name)
		if err != nil {
			log.Fatal(err)
		}

		blobHash := entry.Hash

		err = WriteObject("blobs", blobHash, content)
		if err != nil {
			log.Fatal(err)
		}
	}
}
