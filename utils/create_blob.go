package utils

import (
	"crypto/sha1"
	"fmt"
)

// CreateBlob constructs a Git-style blob object from the contents given
func CreateBlob(contents []byte) ([]byte, string) {
	header := fmt.Sprintf("blob %d\x00", len(contents))

	fullBlob := append([]byte(header), contents...)

	hash := sha1.Sum(fullBlob)
	hashString := fmt.Sprintf("%x", hash)

	return fullBlob, hashString
}
