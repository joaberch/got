package model

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
)

// Commit is an object to store the data of the commit
type Commit struct {
	TreeHash   string
	ParentHash string
	Author     string
	Message    string
	Timestamp  int64
}

// Serialize the commit object in JSON
func (commit *Commit) Serialize() ([]byte, error) {
	return json.Marshal(commit)
}

func (commit *Commit) Hash(data []byte) string {
	return fmt.Sprintf("%x", sha1.Sum(data))
}
