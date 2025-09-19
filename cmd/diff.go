package cmd

import (
	"bytes"
	"fmt"
	"github.com/joaberch/got/utils"
	"os"
	"path/filepath"
	"strings"
)

func Diff() error {
	//head -> contains latest commit hash
	headPath := filepath.Join(".got", "head")
	headHashBytes, err := os.ReadFile(headPath)
	if err != nil {
		return fmt.Errorf("error reading head file: %s", err)
	}
	headHash := strings.TrimSpace(string(headHashBytes))

	//Commit -> contains tree hash
	commitPath := filepath.Join(".got", "objects", "commits", headHash)
	commitData, err := os.ReadFile(commitPath)
	if err != nil {
		return fmt.Errorf("error reading commit file: %s", err)
	}
	commit, err := utils.DeserializeCommit(commitData)
	if err != nil {
		return fmt.Errorf("error deserializing commit: %s", err)
	}

	//Tree -> contains hash of blob(s)
	treePath := filepath.Join(".got", "objects", "trees", commit.TreeHash)
	treeData, err := os.ReadFile(treePath)
	if err != nil {
		return fmt.Errorf("error reading tree file: %s", err)
	}
	tree, err := utils.DeserializeTree(treeData)
	if err != nil {
		return fmt.Errorf("error deserializing trees: %s", err)
	}

	//Compare each file in the tree with its current version
	for _, entry := range tree.Entries {
		blobPath := filepath.Join(".got", "objects", "blobs", entry.Hash)
		committedData, err := os.ReadFile(blobPath)
		if err != nil {
			fmt.Printf("error reading blob file %s: %s", blobPath, err)
			continue
		}

		currentData, err := os.ReadFile(entry.Name)
		if err != nil {
			fmt.Printf("error reading blob file %s: %s", blobPath, err)
			continue
		}

		if !bytes.Equal(currentData, committedData) {
			fmt.Printf("Diff for %s:\n", entry.Name)
			utils.ShowLineDiff(string(committedData), string(currentData))
		}
	}
	return nil
}
