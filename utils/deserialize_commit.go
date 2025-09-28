package utils

import (
	"encoding/json"
	"github.com/joaberch/got/internal/model"
)

// DeserializeCommit deserializes JSON-encoded data into a model.Commit.
// 
// The input `data` must contain a JSON representation of model.Commit. The
// function returns the decoded Commit value and any error from json.Unmarshal.
// No additional validation is performed; on error the returned Commit is the
// zero value for model.Commit.
func DeserializeCommit(data []byte) (model.Commit, error) {
	var commit model.Commit
	err := json.Unmarshal(data, &commit)
	return commit, err
}
