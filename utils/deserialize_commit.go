package utils

import (
	"encoding/json"
	"github.com/joaberch/got/internal/model"
)

// DeserializeCommit deserializes JSON-encoded data into a model.Commit and returns
// a pointer to the resulting Commit along with any error produced by json.Unmarshal.
// The function does no additional validation of the decoded value.
func DeserializeCommit(data []byte) (*model.Commit, error) {
	var commit model.Commit
	err := json.Unmarshal(data, &commit)
	return &commit, err
}
