package utils

import (
	"encoding/csv"
	"fmt"
	"github.com/joaberch/got/internal/model"
	"os"
)

// AddToCommits appends a single commit record as a CSV row to the file at commitsPath.
// 
// The written row fields (in order) are: commitHash, commit.TreeHash, commit.Author,
// commit.Message, and commit.Timestamp formatted as a decimal string.
// 
// It returns a non-nil error if the file cannot be opened or if writing the CSV row fails.
// Errors from closing the file are captured internally but are not propagated to the caller.
func AddToCommits(commitsPath string, commitHash string, commit model.Commit) error {
	file, err := os.OpenFile(commitsPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open commits file at %s: %w", commitsPath, err)
	}
	defer func() {
		errClose := file.Close()
		if errClose != nil {
			err = fmt.Errorf("failed to  close commits file at %s: %w", commitsPath, errClose)
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
		return fmt.Errorf("failed to write to commits file at %s: %w", commitsPath, err)
	}
	return nil
}
