package cmd

import (
	"Got/internal/model"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func Commit() {
	tree := model.Tree{}
	stagingPath := filepath.Join(".got", "staging.csv")
	commitsPath := filepath.Join(".got", "commits.csv")

	file, err := os.Open(stagingPath)
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, record := range records {
		hash := record[0]
		path := record[1]

		tree.Entries = append(tree.Entries, model.TreeEntry{
			Name: path,
			Hash: hash,
			Mode: "file",
			Type: "blob",
		})
	}

	commit := model.Commit{
		TreeHash:   tree.GenerateHash(),
		ParentHash: "TODO",
		Author:     "TODO - none for MVP",
		Message:    "TODO - none for MVP",
		Timestamp:  time.Now().Unix(),
	}

	commitSerialized, err := commit.Serialize()
	if err != nil {
		log.Fatal(err)
	}
	commitHash := commit.Hash() //TODO - redundant in serialization
	objectPath := filepath.Join(".got", "objects", commitHash)
	err = os.WriteFile(objectPath, commitSerialized, 0644)
	if err != nil {
		log.Fatal(err)
	}

	file, err = os.OpenFile(commitsPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{
		commitHash,
		commit.TreeHash,
		commit.Author,
		commit.Message,
		fmt.Sprintf("%d", commit.Timestamp),
	})
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(stagingPath, []byte{}, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
