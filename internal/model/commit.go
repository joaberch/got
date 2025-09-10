package model

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"log"
)

type Commit struct {
	TreeHash   string
	ParentHash string
	Author     string
	Message    string
	Timestamp  int64
}

// Serialize the commit object in json
func (commit *Commit) Serialize() ([]byte, error) {
	return json.Marshal(commit)
}

func (commit *Commit) Hash() string {
	data, err := commit.Serialize()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x", sha1.Sum(data))
}
