package utils

import (
	"Got/internal/model"
	"encoding/json"
)

// DeserializeTree unmarshals JSON-encoded data into a model.Tree and returns a pointer to it.
// The input data must contain a JSON representation of a model.Tree. If unmarshalling fails,
// the returned error describes the failure and the returned *model.Tree will be the zero-value instance.
func DeserializeTree(data []byte) (*model.Tree, error) {
	var tree model.Tree
	err := json.Unmarshal(data, &tree)
	return &tree, err
}
