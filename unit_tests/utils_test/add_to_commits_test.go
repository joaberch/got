package utils_test

import (
	"encoding/csv"
	"fmt"
	"github.com/joaberch/got/internal/model"
	"github.com/joaberch/got/utils"
	"os"
	"testing"
	"time"
)

func TestAddToCommits_FileOpenError(t *testing.T) {
	commit := model.Commit{
		TreeHash:  "abcd",
		Author:    "tester",
		Message:   "test message",
		Timestamp: time.Now().Unix(),
	}

	err := utils.AddToCommits("/not/a/valid/path/file.csv", "hash123", commit)
	if err == nil {
		t.Error("Expected an error")
	}
}

func TestAddToCommits_Success(t *testing.T) {
	//Do
	tmpFile, err := os.CreateTemp("", "commits_*.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = os.Remove(tmpFile.Name())
		if err != nil {
			t.Fatal(err)
		}
	}()
	defer func() {
		err = tmpFile.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()

	commit := model.Commit{
		TreeHash:   "test123",
		Author:     "tester",
		ParentHash: "",
		Message:    "Initial commit",
		Timestamp:  time.Now().Unix(),
	}

	err = utils.AddToCommits(tmpFile.Name(), "hash123", commit)
	if err != nil {
		t.Fatal(err)
	}

	file, err := os.Open(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()

	//Check
	tmpFile.Seek(0, 0)
	reader := csv.NewReader(tmpFile)
	records, err := reader.ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(records) != 1 {
		t.Fatalf("expected 1 record, got %d", len(records))
	}

	expected := []string{
		"hash123",
		commit.TreeHash,
		commit.Author,
		commit.Message,
		fmt.Sprintf("%d", commit.Timestamp),
	}

	for i, val := range expected {
		if records[0][i] != val {
			t.Fatalf("expected %s, got %s", val, records[0][i])
		}
	}
}
