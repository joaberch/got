package model

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
)

// Tree stores all the data of each file
type Tree struct {
	Entries []TreeEntry
	Hash    string
}

// GenerateHash using data
func (tree Tree) GenerateHash() (string, error) {
	data, err := json.Marshal(tree.Entries)
	if err != nil {
		return tree.Hash, err
	}

	sum := sha1.Sum(data)
	tree.Hash = fmt.Sprintf("%x", sum)
	return tree.Hash, nil
}

// Serialize the data in json
func (tree Tree) Serialize() ([]byte, error) {
	data, err := json.Marshal(tree)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}
