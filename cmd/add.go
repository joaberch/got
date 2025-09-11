package cmd

import (
	"Got/internal/model"
	"Got/utils"
)

// Add a file in the staging area
func Add(path string) {
	//Read the file
	contents := utils.GetFileContent(path)

	blob := model.Blob{
		Content: contents,
	}

	//Get file hash
	blob.GenerateHash()

	//Add (the relative path, hash, (perm)) to staging.csv
	utils.AddToStaging(path, blob.Hash)
}
