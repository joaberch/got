package cmd

import (
	"bytes"
	"fmt"
	"github.com/joaberch/got/utils"
	"os"
)

func Diff() error {
	//head -> contains latest commit hash
	headHash, err := utils.GetLatestCommitHash()
	if err != nil {
		return fmt.Errorf("failed to get latest commit hash: %v", err)
	}

	//Commit -> contains tree hash
	commit, err := utils.GetCommitFromHash(headHash)
	if err != nil {
		return fmt.Errorf("failed to get commit: %v", err)
	}

	//Tree -> contains hash of blob(s)
	tree, err := utils.GetTreeFromCommit(commit)

	//Compare each file in the tree with its current version
	for _, entry := range tree.Entries {
		committedBlob, err := utils.GetBlobFromHash(entry.Hash)
		committedData := committedBlob.Content
		if err != nil {
			fmt.Printf("error reading blob file %s: %s", entry.Name, err)
			continue
		}

		currentData, err := os.ReadFile(entry.Name)
		if err != nil {
			fmt.Printf("error reading blob file %s: %s", entry.Name, err)
			continue
		}

		if !bytes.Equal(currentData, committedData) {
			fmt.Printf("Diff for %s:\n", entry.Name)
			utils.ShowLineDiff(string(committedData), string(currentData))
		}
	}
	return nil
}
