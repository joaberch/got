package utils

import (
	"Got/internal/model"
	"encoding/json"
)

// DeserializeCommit and returns it
func DeserializeCommit(data []byte) (*model.Commit, error) {
	var commit model.Commit
	err := json.Unmarshal(data, &commit)
	return &commit, err
}
