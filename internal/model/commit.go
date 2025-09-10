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

// DeserializeCommit has been completely written by goland IDE, thx
func DeserializeCommit(data []byte) (*Commit, error) {
	var commit Commit
	err := json.Unmarshal(data, &commit)
	return &commit, err
}

func (commit *Commit) Hash() string {
	data, err := commit.Serialize()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x", sha1.Sum(data))
}
