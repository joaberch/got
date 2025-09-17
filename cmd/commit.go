package cmd

import (
	"github.com/joaberch/got/internal/model"
	"github.com/joaberch/got/utils"
	"log"
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
func Commit(message string) {
	stagingPath := filepath.Join(".got", "staging.csv")
	commitsPath := filepath.Join(".got", "commits.csv")

	tree, err := utils.ReadStagingFile(stagingPath)
	if err != nil {
		log.Fatal(err)
	}
	treeHash, err := tree.GenerateHash()
	if err != nil {
		log.Fatal(err)
	}

	err = utils.CreateBlobs(tree) //.got/objects/blobs
	if err != nil {
		log.Fatal(err)
	}

	parentHash, err := utils.GetLatestCommitHash()
	commit := model.Commit{
		TreeHash:   treeHash,
		ParentHash: parentHash,
		Author:     "TODO - none for MVP",
		Message:    message,
		Timestamp:  time.Now().Unix(),
	}

	treeSerialized, err := tree.Serialize()
	if err != nil {
		log.Fatal(err)
	}
	err = utils.WriteObject("trees", treeHash, treeSerialized) //.got/objects/trees
	if err != nil {
		log.Fatal(err)
	}

	commitSerialized, err := commit.Serialize()
	if err != nil {
		log.Fatal(err)
	}
	commitHash := commit.Hash(commitSerialized)

	err = utils.WriteObject("commits", commitHash, commitSerialized)
	if err != nil {
		log.Fatal(err)
	}

	err = utils.AddToCommits(commitsPath, commitHash, commit)
	if err != nil {
		log.Fatal(err)
	}
	err = utils.AddToHead(commitHash)
	if err != nil {
		log.Fatal(err)
	}

	err = utils.ClearFile(stagingPath)
	if err != nil {
		log.Fatal(err)
	}
}
