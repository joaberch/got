package utils

import (
	"log"
	"os"
	"path/filepath"
)

// AddToHead changes the latest commit hash in the head file
func AddToHead(hash string) {
	headPath := filepath.Join(".got", "head")
	err := ClearFile(headPath)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(headPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	_, err = file.WriteString(hash)
	if err != nil {
		log.Fatal(err)
	}
}
