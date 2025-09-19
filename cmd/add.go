package cmd

import (
	"errors"
	"fmt"
	"github.com/joaberch/got/internal/model"
	"github.com/joaberch/got/utils"
	"strings"
)

// Add stages the file at the given path by creating a blob from its contents,
// computing its hash, and recording the path→hash pair in the staging area.
// 
// The function returns an error if the path contains ".got" (the tool ignores
// its own metadata files), if the file contents cannot be read, or if writing
// the entry to the staging area fails.
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
