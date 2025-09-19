package utils

import (
	"encoding/json"
	"github.com/joaberch/got/internal/model"
)

// DeserializeTree unmarshals JSON-encoded data into a model.Tree and returns a pointer to it.
// The input data must contain a JSON representation of a model.Tree. If unmarshalling fails,
// DeserializeTree decodes JSON-encoded data into a model.Tree.
// If decoding succeeds it returns the populated Tree and a nil error.
// On failure it returns the zero-value Tree and the decoding error.
// The input must be JSON representing a model.Tree.
func DeserializeTree(data []byte) (model.Tree, error) {
	var tree model.Tree
	err := json.Unmarshal(data, &tree)
	return tree, err
}
