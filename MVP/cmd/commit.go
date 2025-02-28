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

func GenerateCommitID(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func CommitChange(message string) {
	//Check message
	if message == "" {
		fmt.Errorf("commit message cannot be empty")
	}

	//Open the staging file
	file, err := os.Open(stagingPath)
	if err != nil {
		fmt.Errorf("error while opening %v : %v", stagingPath, err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Errorf("error while reading %v : %v", stagingPath, err)
	}

	//Generate commit object
	var files []CommitFile
	for _, record := range records {
		files = append(files, CommitFile{
			Path:     record[0],
			Checksum: record[1],
			Status:   record[2],
		})
		TrackFile(record[0]) //add the new file to the tracking
	}

	// create the commit object
	commit := Commit{
		CommitId: GenerateCommitID(time.Now().Format(time.RFC3339)),
		Date:     time.Now().Format(time.RFC3339),
		Message:  message,
		Files:    files,
	}

	//Write the commit in a file
	commitFilePath := fmt.Sprintf(".got/objects/%s.json", commit.CommitId)
	commitFile, err := os.Create(commitFilePath)
	if err != nil {
		fmt.Errorf("error while creating commit file : %v", err)
	}
	defer func(commitFile *os.File) {
		err := commitFile.Close()
		if err != nil {

		}
	}(commitFile)

	if err = json.NewEncoder(commitFile).Encode(commit); err != nil {
		fmt.Errorf("error while encoding commit : %v", err)
	}

	//Clean staging area
	if err = os.Truncate(stagingPath, 0); err != nil {
		fmt.Errorf("error while clearing %v : %v", stagingPath, err)
	}

	fmt.Printf("Commit %s created successfully !", commit.CommitId)
}

func TrackFile(path string) {
	//Check tracking file exist
	if _, err := os.Stat(trackingFilePath); os.IsNotExist(err) {
		fmt.Errorf("tracking file does not exist, initialize got first")
	}

	//Open file in append mode
	file, err := os.OpenFile(trackingFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Errorf("error opening %v: %v", trackingFilePath, err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	//Load what's already tracked to prevent duplicate
	existingFiles := make(map[string]bool)

	trackReader, err := os.Open(trackingFilePath)
	if err != nil {
		fmt.Errorf("error opening %v: %v", trackingFilePath, err)
	}
	defer func(trackReader *os.File) {
		err := trackReader.Close()
		if err != nil {

		}
	}(trackReader)

	csvReader := csv.NewReader(trackReader)
	existingRecords, err := csvReader.ReadAll()
	if err == nil {
		for _, record := range existingRecords {
			if len(record) > 0 {
				existingFiles[record[0]] = true
			}
		}
	}

	//Check duplicate
	if _, exists := existingFiles[path]; exists {
		//Already tracked
		return
	}

	//Add the new path in the tracking file
	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	if err := csvWriter.Write([]string{path}); err != nil {
		fmt.Errorf("error writing to %v: %v", trackingFilePath, err)
	}
}
