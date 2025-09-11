package cmd

import (
	"Got/internal/model"
	"Got/utils"
	"log"
	"os"
	"path/filepath"
)

// Init initializes a .got directory in the current working directory.
// 
// If the current working directory cannot be determined or a .got directory
// already exists, the function logs a fatal error and exits the process.
// Otherwise it creates the mandatory files and folders defined in
// model.FilesList under the newly created .got directory using utils.CreateFilePath.
func Init() {
	pwd, err := os.Getwd() //get current folder
	if err != nil {
		log.Fatal(err)
	}
	gotPath := filepath.Join(pwd, ".got")

	if _, err = os.Stat(gotPath); !os.IsNotExist(err) { //Check if already exist
		log.Fatal("This directory already exists")
		return
	}

	for name, fileType := range model.FilesList { //Create mandatory files/folder
		fullPath := filepath.Join(gotPath, name)
		utils.CreateFilePath(fullPath, fileType)
	}
}
