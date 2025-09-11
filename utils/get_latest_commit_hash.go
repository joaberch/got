package utils

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

// GetLatestCommitHash returns the hash of the latest commit
func GetLatestCommitHash() string {
	headPath := filepath.Join(".got", "head")
	file, err := os.Open(headPath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var content = ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		content += line + "\n"
	}
	return content
}
