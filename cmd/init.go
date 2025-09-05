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

	err = os.Mkdir(pwd+"/.got", 0755) //Create folder
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range model.FilesList {
		_, err = os.Create(filepath.Join(pwd, file))
		if err != nil {
			log.Fatal(err)
		}
	}
}
