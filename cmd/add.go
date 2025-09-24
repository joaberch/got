package cmd

import (
	"errors"
	"fmt"
	"github.com/joaberch/got/internal/model"
	"github.com/joaberch/got/utils"
	"os"
	"path/filepath"
	"strings"
)

// Add stages the file at the given path by creating a blob from its contents,
// computing its hash, and recording the path→hash pair in the staging area.
//
// The function returns an error if the path contains ".got" (the tool ignores
// its own metadata files), if the file contents cannot be read, or if writing
// the entry to the staging area fails.
func Add(path string) error { //TODO - auto-stage from file already added
	if strings.Contains(path, ".got") {
		return errors.New("path contains '.got', got doesn't process itself")
	}

	f, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("os.Stat(%s): %v", path, err)
	}

	if f.IsDir() { //If dir, recursive
		entries, err := os.ReadDir(path)
		if err != nil {
			return fmt.Errorf("os.ReadDir(%s): %v", path, err)
		}

		for _, entry := range entries {
			entryPath := filepath.Join(path, entry.Name())
			if strings.Contains(entryPath, ".got") {
				continue
			}

			err = Add(entryPath)
			if err != nil {
				return err
			}
		}
	} else {
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
	}
	return nil
}
