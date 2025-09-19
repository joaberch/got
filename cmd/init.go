package cmd

import (
	"fmt"
	"github.com/joaberch/got/internal/model"
	"github.com/joaberch/got/utils"
	"os"
	"path/filepath"
)

// Init creates a ".got" directory in the current working directory and populates it
// with the mandatory files and folders defined in model.FilesList using
// utils.CreateFilePath. It returns an error if the working directory cannot be
// determined, if the ".got" directory already exists, or if creating any required
// file or directory fails.
func Init() error {
	pwd, err := os.Getwd() //get current folder
	if err != nil {
		return fmt.Errorf("error getting working directory: %s", err)
	}
	gotPath := filepath.Join(pwd, ".got")

	if _, err = os.Stat(gotPath); !os.IsNotExist(err) { //Check if already exist
		return fmt.Errorf("this directory already exists: %s: %s", gotPath, err)
	}

	for name, fileType := range model.FilesList { //Create mandatory files/folder
		fullPath := filepath.Join(gotPath, name)
		err = utils.CreateFilePath(fullPath, fileType)
		if err != nil {
			return fmt.Errorf("error creating file %s: %s", name, err)
		}
	}
	return nil
}
