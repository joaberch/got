package utils

import (
	"fmt"
	"github.com/joaberch/got/internal/model"
	"os"
	"path/filepath"
)

// GetCommitFromHash returns the Commit object stored at .got/objects/commits/{hash}.
// It reads the commit file for the provided hash and deserializes its contents.
// Returns an error if the file cannot be read or if deserialization fails.
func GetCommitFromHash(hash string) (model.Commit, error) {
	commitPath := filepath.Join(".got", "objects", "commits", hash)
	data, err := os.ReadFile(commitPath)
	if err != nil {
		return model.Commit{}, fmt.Errorf("error reading commit file a commit needs to be done first to see the differences %s: %s", commitPath, err)
	}
	commit, err := DeserializeCommit(data)
	if err != nil {
		return model.Commit{}, fmt.Errorf("error parsing commit file %s: %s", commitPath, err)
	}
	return commit, nil
}
