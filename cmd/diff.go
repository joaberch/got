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
// Returns an error only if resolving the latest commit hash, or the commit object fails; otherwise it returns nil.
func Diff(verbose bool, args []string) error {
	var targetFile string
	if len(args) > 1 {
		targetFile = args[1]

	}
	//head -> contains the latest commit hash
	headHash, err := utils.GetLatestCommitHash()
	if err != nil {
		return fmt.Errorf("could not get latest commit hash: %w", err)
	}
	if headHash == "" {
		//No commit done
		stagedEntries, err := utils.GetStagedEntries()
		if err != nil {
			return fmt.Errorf("unable to get staged entries: %w", err)
		}

		for _, entry := range stagedEntries {
			if targetFile != "" && entry != targetFile {
				continue
			}
			fmt.Printf("New file staged: %s\n", entry)
			currentData, err := utils.GetFileContent(entry)
			if err != nil {
				if verbose {
					fmt.Printf("unable to get file content: %v", err)
				}
				continue
			}
			utils.ShowLineDiff("", string(currentData))
		}
		return nil
	}

	//Commit -> contains tree hash
	commit, err := utils.GetCommitFromHash(headHash)
	if err != nil {
		return fmt.Errorf("failed to get commit: %w", err)
	}

	//Tree -> contains hash of blob(s)
	tree, err := utils.GetTreeFromCommit(commit)
	if err != nil {
		return fmt.Errorf("failed to get tree: %w", err)
	}

	stagedEntries, err := utils.GetStagedEntries()
	if err != nil {
		return err
	}

	indexedMap, err := utils.GetIndexedPathsWithHash()
	if err != nil {
		return fmt.Errorf("failed to get indexed paths: %w", err)
	}

	committedFiles := make(map[string]string)
	for _, entry := range tree.Entries {
		committedFiles[entry.Name] = entry.Hash
	}

	for path, lastBlobHash := range indexedMap {
		if targetFile != "" && path != targetFile {
			continue
		}
		currentData, err := utils.GetFileContent(path)
		if os.IsNotExist(err) {
			fmt.Printf("Deleted file: %s\n", path)
			committedBlob, err := utils.GetBlobFromHash(lastBlobHash)
			if err == nil {
				utils.ShowLineDiff(string(committedBlob.Content), "")
			}
			continue
		} else if err != nil {
			if verbose {
				fmt.Printf("Failed to get committed blob: %v", err)
			}
			continue
		}

		committedBlob, err := utils.GetBlobFromHash(lastBlobHash)
		if err != nil {
			if verbose {
				fmt.Printf("Failed to get committed blob: %v", err)
			}
			continue
		}

		if !bytes.Equal(committedBlob.Content, currentData) {
			fmt.Printf("Modified file: %s\n", path)
			utils.ShowLineDiff(string(committedBlob.Content), string(currentData))
		}
	}

	//Detect the added file
	for _, stagedPath := range stagedEntries {
		if _, exists := committedFiles[stagedPath]; !exists {
			if targetFile != "" && stagedPath != targetFile {
				continue
			}
			fmt.Printf("New file staged: %s\n", stagedPath)
			currentData, err := utils.GetFileContent(stagedPath)
			if err != nil {
				if verbose {
					fmt.Printf("failed to read file: %v", err)
				}
				continue
			}
			utils.ShowLineDiff("", string(currentData))
		}
	}

	return nil
}
