package cmd

import (
	"Got/internal/model"
	"Got/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func Commit(message string) {
	stagingPath := filepath.Join(".got", "staging.csv")
	commitsPath := filepath.Join(".got", "commits.csv")

	tree := utils.ReadStagingFile(stagingPath)
	treeHash := tree.GenerateHash()

	commit := model.Commit{
		TreeHash:   treeHash,
		ParentHash: "TODO",
		Author:     "TODO - none for MVP",
		Message:    message,
		Timestamp:  time.Now().Unix(),
	}

	treeSerialized := tree.Serialize()
	fmt.Print("test", treeHash)
	err := os.WriteFile(".got/objects/"+treeHash, treeSerialized, 0644)
	if err != nil {
		log.Fatal(err)
	}

	commitSerialized, err := commit.Serialize()
	if err != nil {
		log.Fatal(err)
	}
	commitHash := commit.Hash() //TODO - redundant in serialization

	err = utils.WriteObject(commitHash, commitSerialized)
	if err != nil {
		log.Fatal(err)
	}

	utils.AddEntryInCommitsFile(commitsPath, commitHash, commit)

	err = utils.ClearFile(stagingPath)
	if err != nil {
		log.Fatal(err)
	}
}
