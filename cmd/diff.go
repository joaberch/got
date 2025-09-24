package cmd

import (
	"bytes"
	"fmt"
	"github.com/joaberch/got/utils"
	"os"
)

// Diff compares the current working-tree files against the latest commit and prints per-file line-based
// differences for any files that have changed.
//
// It resolves the latest commit, walks the commit tree, and for each tree entry compares the committed blob
// content to the current file on disk. Per-entry read errors are printed and that entry is skipped.
//
// Returns an error only if resolving the latest commit hash or the commit object fails; otherwise it returns nil.
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
	if err != nil {
		return fmt.Errorf("failed to get tree: %v", err)
	}

	stagedEntries, err := utils.GetStagedEntries()
	if err != nil {
		return err
	}

	commitedFiles := make(map[string]string)
	for _, entry := range tree.Entries {
		commitedFiles[entry.Name] = entry.Hash
	}

	//Detect the modified file
	for name, hash := range commitedFiles {
		commitedBlob, err := utils.GetBlobFromHash(hash)
		if err != nil {
			fmt.Printf("failed to get commited blob: %v", err)
			continue
		}

		currentData, err := os.ReadFile(name)
		if err != nil {
			fmt.Printf("failed to read file: %v", err)
			continue
		}

		if !bytes.Equal(currentData, commitedBlob.Content) {
			fmt.Printf("Modified file: %s\n", name)
			utils.ShowLineDiff(string(commitedBlob.Content), string(currentData))
		}
	}

	//Detect the added file
	for _, stagedPath := range stagedEntries {
		if _, exists := commitedFiles[stagedPath]; !exists {
			fmt.Printf("New file staged: %s\n", stagedPath)
			currentData, err := os.ReadFile(stagedPath)
			if err != nil {
				fmt.Printf("failed to read file: %v", err)
			}
			utils.ShowLineDiff("", string(currentData))
		}
	}

	//Detect the deleted file
	for name, hash := range commitedFiles {
		_, err = os.Stat(name)
		if os.IsNotExist(err) {
			fmt.Printf("Deleted file: %s\n", name)

			commitedBlob, err := utils.GetBlobFromHash(hash)
			if err != nil {
				fmt.Printf("failed to get commited blob: %v", err)
				continue
			}
			utils.ShowLineDiff(string(commitedBlob.Content), "")
		}
	}
	return nil
}
