package cmd

import (
	"crypto/sha1"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type CommitFile struct {
	Path     string
	Checksum string
	Status   string
}

type Commit struct {
	CommitId string       `json:"commitId"`
	Date     string       `json:"date"`
	Message  string       `json:"message"`
	Files    []CommitFile `json:"files"`
}

func GenerateCommitID(data string, message string, files []CommitFile) string {
	h := sha1.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func CommitChange(message string) error {
	//Check message
	if message == "" {
		return fmt.Errorf("commit message cannot be empty")
	}

	//Open the staging file
	file, err := os.Open(stagingPath)
	if err != nil {
		return fmt.Errorf("error while opening %v : %v", stagingPath, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("error while reading %v : %v", stagingPath, err)
	}

	//Generate commit object
	var files []CommitFile
	for _, record := range records {
		files = append(files, CommitFile{
			Path:     record[0],
			Checksum: record[1],
			Status:   record[2],
		})
	}

	commit := Commit{
		CommitId: GenerateCommitID(time.Now().Format(time.RFC3339), message, files),
		Date:     time.Now().Format(time.RFC3339),
		Message:  message,
		Files:    files,
	}

	//Write the commit in a file
	commitFilePath := fmt.Sprintf(".got/objects/%s.json", commit.CommitId)
	commitFile, err := os.Create(commitFilePath)
	if err != nil {
		return fmt.Errorf("error while creating commit file : %v", err)
	}
	defer commitFile.Close()

	if err = json.NewEncoder(commitFile).Encode(commit); err != nil {
		return fmt.Errorf("error while encoding commit : %v", err)
	}

	//Clean staging area
	if err = os.Truncate(stagingPath, 0); err != nil {
		return fmt.Errorf("error while clearing %v : %v", stagingPath, err)
	}

	fmt.Printf("Commit %s created successfully !", commit.CommitId)
	return nil
}
