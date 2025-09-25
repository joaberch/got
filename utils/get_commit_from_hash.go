package utils

import (
	"fmt"
	"github.com/joaberch/got/internal/model"
	"path/filepath"
)

// GetCommitFromHash returns the Commit object stored at .got/objects/commits/{hash}.
// It reads the commit file for the provided hash and deserializes its contents.
// Returns an error if the file cannot be read or if deserialization fails.
func GetCommitFromHash(hash string) (model.Commit, error) {
	commitPath := filepath.Join(".got", "objects", "commits", hash)
	data, err := GetFileContent(commitPath)
	if err != nil {
		return model.Commit{}, fmt.Errorf("error reading commit %s: %w", commitPath, err)
	}
	commit, err := DeserializeCommit(data)
	if err != nil {
		return model.Commit{}, fmt.Errorf("error parsing commit file %s: %w", commitPath, err)
	}
	return commit, nil
}
