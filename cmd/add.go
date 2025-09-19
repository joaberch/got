package cmd

import (
	"errors"
	"fmt"
	"github.com/joaberch/got/internal/model"
	"github.com/joaberch/got/utils"
	"strings"
)

// Add stages the file at the given path by creating a blob from its contents
// and recording the blob's hash in the repository staging area.
//
// The path is the filesystem path of the file to stage; the function reads the
// file contents, constructs a model.Blob, generates its hash, and writes an
// entry (path and hash) into the staging area. This function does not return
// an error.
func Add(path string) error {
	if strings.Contains(path, ".got") {
		return errors.New("path contains '.got', got doesn't process itself")
	}

	//Read the file
	contents, err := utils.GetFileContent(path)
	if err != nil {
		return fmt.Errorf("error getting file contents: %s", err)
	}

	blob := model.Blob{
		Content: contents,
	}

	//Get file hash
	blob.GenerateHash()

	//Add (the relative path, hash, (perm)) to staging.csv
	err = utils.AddToStaging(path, blob.Hash)
	if err != nil {
		return fmt.Errorf("error adding to staging file: %s", err)
	}
	return nil
}
