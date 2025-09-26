package cmd

import (
	"fmt"
	"github.com/joaberch/got/internal/model"
	"github.com/joaberch/got/utils"
	"path/filepath"
	"strings"
	"time"
)

// Commit creates a new commit from the current staging state (.got/staging.csv) and updates the repository.
//
// It reads the staging file, generates a tree and its blobs, writes the tree and commit objects to the object store,
// appends the commit to .got/commits.csv, updates .got/head with the new commit hash, and clears the staging file.
//
// The message parameter is used as the commit message.
// Returns an error if any step (reading staging, hashing/serializing, writing objects, updating commits/head, or clearing staging) fails.
func Commit(message []string) error {
	stagingPath := filepath.Join(".got", "staging.csv")
	commitsPath := filepath.Join(".got", "commits.csv")

	tree, err := utils.ReadStagingFile(stagingPath)
	if err != nil {
		return fmt.Errorf("error reading staging file: %w", err)
	}
	if len(tree.Entries) == 0 {
		return fmt.Errorf("cannot commit: staging area is empty")
	}
	treeHash, err := tree.GenerateHash()
	if err != nil {
		return fmt.Errorf("error generating tree hash: %w", err)
	}

	err = utils.CreateBlobs(tree) //.got/objects/blobs
	if err != nil {
		return fmt.Errorf("error creating blobs: %w", err)
	}

	latestCommitHash, err := utils.GetLatestCommitHash()
	if err != nil {
		return fmt.Errorf("error getting latest commit hash: %w", err)
	}

	commit := model.Commit{
		TreeHash:   treeHash,
		ParentHash: latestCommitHash,
		Author:     "TODO - none for MVP",
		Message:    strings.Join(message, " "),
		Timestamp:  time.Now().Unix(),
	}

	treeSerialized, err := tree.Serialize()
	if err != nil {
		return fmt.Errorf("error serializing tree: %w", err)
	}
	err = utils.WriteObject("trees", treeHash, treeSerialized) //.got/objects/trees
	if err != nil {
		return fmt.Errorf("error writing trees: %w", err)
	}

	commitSerialized, err := commit.Serialize()
	if err != nil {
		return fmt.Errorf("error serializing commit: %w", err)
	}
	commitHash := commit.Hash(commitSerialized)

	err = utils.WriteObject("commits", commitHash, commitSerialized)
	if err != nil {
		return fmt.Errorf("error writing commits: %w", err)
	}

	err = utils.AddToCommits(commitsPath, commitHash, commit)
	if err != nil {
		return fmt.Errorf("error adding to commits: %w", err)
	}
	headPath := filepath.Join(".got", "head")
	err = utils.AddToHead(headPath, commitHash)
	if err != nil {
		return fmt.Errorf("error adding to head: %w", err)
	}

	err = utils.ClearFile(stagingPath)
	if err != nil {
		return fmt.Errorf("error clearing staging file: %w", err)
	}

	err = utils.UpdateIndexedPaths(tree)
	if err != nil {
		return fmt.Errorf("error updating indexed paths: %w", err)
	}

	return nil
}
