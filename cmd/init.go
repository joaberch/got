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
		fullpath := filepath.Join(pwd + "/.got/" + name)

		if fileType == "Folder" {
			err = os.MkdirAll(fullpath, 0755)
			if err != nil {
				log.Fatal(err)
			}
		} else if fileType == "File" {
			_, err = os.Create(fullpath)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
