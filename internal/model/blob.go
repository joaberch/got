package model

import (
	"crypto/sha1"
	"encoding/hex"
)

// Blob is an object to store the file content as blob and name the file as his hash (SHA-1)
type Blob struct {
	Hash    string
	Content []byte
}

// GenerateHash generates the hash of the blob using his content
func (blob *Blob) GenerateHash() string {
	sum := sha1.Sum(blob.Content)
	blob.Hash = hex.EncodeToString(sum[:])
	return blob.Hash
}
