package utils

import (
	"log"
	"os"
	"path/filepath"
)

// AddToHead clears the .got/head file and writes hash as the current HEAD commit hash.
// It first empties the head file (via ClearFile) and then writes the provided hash (written as-is; no trailing newline).
// Any error encountered while clearing, opening, writing, or closing the file causes the program to exit via log.Fatal.
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
