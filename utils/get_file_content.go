package utils

import (
	"log"
	"os"
)

// GetFileContent reads and returns the contents of the file at the given path.
// If the file cannot be read the function logs the error and calls log.Fatal,
// causing the program to exit.
func GetFileContent(path string) []byte {
	contents, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return contents
}
