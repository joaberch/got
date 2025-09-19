package utils

import (
	"fmt"
	"github.com/joaberch/got/internal/model"
	"os"
	"path/filepath"
)

// GetTreeFromCommit reads and parses the tree object referenced by the given commit.
// 
// GetTreeFromCommit looks up the tree file at ".got/objects/trees/<commit.TreeHash>", reads
// its contents and deserializes them with DeserializeTree. It returns the parsed model.Tree on
// success. If the file cannot be read or the data cannot be parsed, it returns an empty
// model.Tree and a non-nil error describing the failure.
func GetTreeFromCommit(commit model.Commit) (model.Tree, error) {
	treePath := filepath.Join(".got", "objects", "trees", commit.TreeHash)
	data, err := os.ReadFile(treePath)
	if err != nil {
		return model.Tree{}, fmt.Errorf("error reading tree file %s: %s", treePath, err)
	}
	tree, err := DeserializeTree(data)
	if err != nil {
		return model.Tree{}, fmt.Errorf("error parsing tree file %s: %s", treePath, err)
	}
	return tree, nil
}
