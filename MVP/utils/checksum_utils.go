package utils

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func GetChecksum(path string) (string, error) {
	//Open the file
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	//SHA1 object
	h := sha1.New()
	//Copy the file content in the Hash object
	if _, err := io.Copy(h, file); err != nil {
		return "", err
	}

	//Return the hash as a string
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
