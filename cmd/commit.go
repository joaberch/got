package cmd

import (
	"fmt"
	"github.com/joaberch/got/internal/model"
	"github.com/joaberch/got/utils"
	"path/filepath"
	"time"
)

// Commit creates a new commit from the current staging state.
//
// It builds the tree and blob objects from .got/staging.csv, writes the tree and commit
// objects into the object store, appends the new commit to .got/commits.csv, updates HEAD,
// and clears the staging file.
//
// The message parameter is used as the commit message. On write/serialization failures the
// Commit creates a new commit from the current staging state (.got/staging.csv) and updates the repository.
//
// It reads the staging file, generates a tree and its blobs, writes the tree and commit objects to the object store,
// appends the commit to .got/commits.csv, updates .got/head with the new commit hash, and clears the staging file.
//
// The message parameter is used as the commit message.
// Returns an error if any step (reading staging, hashing/serializing, writing objects, updating commits/head, or clearing staging) fails.
func Commit(message string) error {
	stagingPath := filepath.Join(".got", "staging.csv")
	commitsPath := filepath.Join(".got", "commits.csv")

	tree, err := utils.ReadStagingFile(stagingPath)
	if err != nil {
		return fmt.Errorf("error reading staging file: %s", err)
	}
	treeHash, err := tree.GenerateHash()
	if err != nil {
		return fmt.Errorf("error generating tree hash: %s", err)
	}

	err = utils.CreateBlobs(tree) //.got/objects/blobs
	if err != nil {
		return fmt.Errorf("error creating blobs: %s", err)
	}

	latestCommitHash, err := utils.GetLatestCommitHash()
	if err != nil {
		return fmt.Errorf("error getting latest commit hash: %s", err)
	}

	commit := model.Commit{
		TreeHash:   treeHash,
		ParentHash: latestCommitHash,
		Author:     "TODO - none for MVP",
		Message:    message,
		Timestamp:  time.Now().Unix(),
	}

	treeSerialized, err := tree.Serialize()
	if err != nil {
		return fmt.Errorf("error serializing tree: %s", err)
	}
	err = utils.WriteObject("trees", treeHash, treeSerialized) //.got/objects/trees
	if err != nil {
		return fmt.Errorf("error writing trees: %s", err)
	}

	commitSerialized, err := commit.Serialize()
	if err != nil {
		return fmt.Errorf("error serializing commit: %s", err)
	}
	commitHash := commit.Hash(commitSerialized)

	err = utils.WriteObject("commits", commitHash, commitSerialized)
	if err != nil {
		return fmt.Errorf("error writing commits: %s", err)
	}

	err = utils.AddToCommits(commitsPath, commitHash, commit)
	if err != nil {
		return fmt.Errorf("error adding to commits: %s", err)
	}
	headPath := filepath.Join(".got", "head")
	err = utils.AddToHead(headPath, commitHash)
	if err != nil {
		return fmt.Errorf("error adding to head: %s", err)
	}

	err = utils.ClearFile(stagingPath)
	if err != nil {
		return fmt.Errorf("error clearing staging file: %s", err)
	}
	return nil
}
