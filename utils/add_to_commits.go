package utils

import (
	"encoding/csv"
	"fmt"
	"github.com/joaberch/got/internal/model"
	"log"
	"os"
)

// AddToCommits appends a commit record to the CSV file at the given path.
// 
// It writes a single CSV row containing: commitHash, commit.TreeHash, commit.Author,
// commit.Message, and commit.Timestamp (as a decimal string). commitsPath is the
// filesystem path to the CSV file; commitHash is the commit's hash, and commit
// supplies the remaining fields.
//
// On any filesystem or write error the function calls log.Fatal and terminates the process.
func AddToCommits(commitsPath string, commitHash string, commit model.Commit) {
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
