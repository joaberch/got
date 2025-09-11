package cmd

import (
	"Got/internal/model"
	"Got/utils"
	"log"
	"os"
	"path/filepath"
)

// Init the got folder with the mandatory files and folder
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
