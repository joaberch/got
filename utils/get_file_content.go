package utils

import (
	"log"
	"os"
)

// GetFileContent returns the content of the file given
func GetFileContent(path string) []byte {
	contents, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return contents
}
