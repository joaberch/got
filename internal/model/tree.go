package model

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"log"
)

// Tree stores all the data of each file
type Tree struct {
	Entries []TreeEntry
	Hash    string
}

// GenerateHash using data
func (tree Tree) GenerateHash() string {
	data, err := json.Marshal(tree.Entries)
	if err != nil {
		log.Fatal(err)
	}

	sum := sha1.Sum(data)
	tree.Hash = fmt.Sprintf("%x", sum)
	return tree.Hash
}

// Serialize the data in json
func (tree Tree) Serialize() []byte {
	data, err := json.Marshal(tree)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
