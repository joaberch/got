package utils

import (
	"Got/internal/model"
	"encoding/json"
)

// DeserializeTree and returns it
func DeserializeTree(data []byte) (*model.Tree, error) {
	var tree model.Tree
	err := json.Unmarshal(data, &tree)
	return &tree, err
}
