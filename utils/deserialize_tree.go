package utils

import (
	"encoding/json"
	"github.com/joaberch/got/internal/model"
)

// DeserializeTree decodes JSON-encoded data into a model.Tree.
// If decoding succeeds it returns the populated Tree and a nil error.
// On failure it returns the zero-value Tree and the decoding error.
// The input must be JSON representing a model.Tree.
func DeserializeTree(data []byte) (model.Tree, error) {
	var tree model.Tree
	err := json.Unmarshal(data, &tree)
	return tree, err
}
