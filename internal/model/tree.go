package model

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"log"
)

type Tree struct {
	Entries []TreeEntry
	Hash    string
}

func (tree Tree) GenerateHash() string {
	data, err := json.Marshal(tree.Entries)
	if err != nil {
		log.Fatal(err)
	}

	sum := sha1.Sum(data)
	tree.Hash = fmt.Sprintf("%x", sum)
	return tree.Hash
}
