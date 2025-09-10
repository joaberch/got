package cmd

import (
	"Got/internal/model"
	"Got/utils"
)

func Add(path string) {
	//Read the file
	contents := utils.StreamFile(path)

	blob := model.Blob{
		Content: contents,
	}

	//Get file hash
	blob.GenerateHash()

	//Add (the relative path, hash, (perm)) to staging.csv
	utils.AddEntryInStagingFile(path, blob.Hash)
}
