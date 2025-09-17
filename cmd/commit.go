package cmd

import (
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
// function calls log.Fatal and terminates the program.
func Commit(message string) error {
	stagingPath := filepath.Join(".got", "staging.csv")
	commitsPath := filepath.Join(".got", "commits.csv")

	tree, err := utils.ReadStagingFile(stagingPath)
	if err != nil {
		return err
	}
	treeHash := tree.GenerateHash()

	err = utils.CreateBlobs(tree) //.got/objects/blobs
	if err != nil {
		return err
	}

	latestCommitHash, err := utils.GetLatestCommitHash()
	if err != nil {
		return err
	}

	commit := model.Commit{
		TreeHash:   treeHash,
		ParentHash: latestCommitHash,
		Author:     "TODO - none for MVP",
		Message:    message,
		Timestamp:  time.Now().Unix(),
	}

	treeSerialized := tree.Serialize()
	err = utils.WriteObject("trees", treeHash, treeSerialized) //.got/objects/trees
	if err != nil {
		return err
	}

	commitSerialized, err := commit.Serialize()
	if err != nil {
		return err
	}
	commitHash := commit.Hash(commitSerialized)

	err = utils.WriteObject("commits", commitHash, commitSerialized)
	if err != nil {
		return err
	}

	err = utils.AddToCommits(commitsPath, commitHash, commit)
	if err != nil {
		return err
	}
	headPath := filepath.Join(".got", "head")
	err = utils.AddToHead(headPath, commitHash)
	if err != nil {
		return err
	}

	err = utils.ClearFile(stagingPath)
	if err != nil {
		return err
	}
	return nil
}
