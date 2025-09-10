package model

import (
	"crypto/sha1"
	"encoding/hex"
)

type Blob struct {
	Hash    string
	Content []byte
}

func (blob *Blob) GenerateHash() string {
	sum := sha1.Sum(blob.Content)
	blob.Hash = hex.EncodeToString(sum[:])
	return blob.Hash
}
