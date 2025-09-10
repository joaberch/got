package utils

import (
	"Got/internal/model"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func AddEntryInCommitsFile(commitsPath string, commitHash string, commit model.Commit) {
	file, err := os.OpenFile(commitsPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = file.Close()
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
}
