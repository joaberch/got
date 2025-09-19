package utils

import (
	"fmt"
	"github.com/joaberch/got/internal/model"
	"os"
	"path/filepath"
)

func GetCommitFromHash(hash string) (model.Commit, error) {
	commitPath := filepath.Join(".got", "objects", "commits", hash)
	data, err := os.ReadFile(commitPath)
	if err != nil {
		return model.Commit{}, fmt.Errorf("error reading commit file %s: %s", commitPath, err)
	}
	commit, err := DeserializeCommit(data)
	if err != nil {
		return model.Commit{}, fmt.Errorf("error parsing commit file %s: %s", commitPath, err)
	}
	return commit, nil
}
