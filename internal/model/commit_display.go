package model

import (
	"fmt"
	"time"
)

type CommitDisplay struct {
	Hash      string
	Author    string
	Message   string
	Timestamp int64
}

func (commit CommitDisplay) Display() {
	fmt.Printf("Commit : %s\n", commit.Hash)
	fmt.Printf("Author : %s\n", commit.Author)
	fmt.Printf("Message : %s\n", commit.Message)
	readableTime := time.Unix(commit.Timestamp, 0).Format("2006-01-02 15:04:05")
	fmt.Printf("Date     : %s\n\n", readableTime)
}
