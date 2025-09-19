package utils

import (
	"fmt"
	"github.com/joaberch/got/internal/model"
	"os"
	"path/filepath"
)

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
