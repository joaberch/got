package cmd

import (
	"Got/internal/model"
	"log"
	"os"
	"path/filepath"
)

func Init() {
	pwd, err := os.Getwd() //get current folder
	if err != nil {
		log.Fatal(err)
	}

	if _, err = os.Stat(pwd + "/.got"); !os.IsNotExist(err) { //Check if already exist
		log.Fatal("This directory already exists")
		return
	}

	for name, fileType := range model.FilesList { //Create mandatory files/folder
		fullPath := filepath.Join(pwd + "/.got/" + name)

		if fileType == "Folder" {
			err = os.MkdirAll(fullPath, 0755)
			if err != nil {
				log.Fatal(err)
			}
		} else if fileType == "File" {
			//Create the parent folder if it doesn't exist
			parentDir := filepath.Dir(fullPath)
			err = os.MkdirAll(parentDir, 0755)
			if err != nil {
				log.Fatal(err)
			}
			
			_, err = os.Create(fullPath)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
