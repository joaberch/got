package cmd

import (
	"Got/internal/model"
	"Got/utils"
)

// Add stages the file at the given path by creating a blob from its contents
// and recording the blob's hash in the repository staging area.
//
// The path is the filesystem path of the file to stage; the function reads the
// file contents, constructs a model.Blob, generates its hash, and writes an
// entry (path and hash) into the staging area. This function does not return
// an error.
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
