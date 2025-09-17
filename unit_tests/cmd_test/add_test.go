package cmd_test

import (
	"crypto/sha1"
	"encoding/csv"
	"encoding/hex"
	"github.com/joaberch/got/cmd"
	"os"
	"path/filepath"
	"testing"
)

func TestAdd_Success(t *testing.T) {
	tmpDir := t.TempDir()

	//Create file to add
	filePath := filepath.Join(tmpDir, "file.txt")
	content := []byte("Hello World")
	err := os.WriteFile(filePath, content, 0666)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	//Create .got folder
	gotDir := filepath.Join(tmpDir, ".got")
	err = os.MkdirAll(gotDir, 0777)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	//Go to temp folder
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	err = os.Chdir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to change working directory: %v", err)
	}
	defer func() {
		err = os.Chdir(oldWd)
		if err != nil {
			t.Fatalf("Failed to change working directory: %v", err)
		}
	}()

	//Act
	err = cmd.Add("file.txt")
	if err != nil {
		t.Fatalf("Failed to add test file: %v", err)
	}

	//Read staging
	stagingPath := filepath.Join(".got", "staging.csv")
	file, err := os.Open(stagingPath)
	if err != nil {
		t.Fatalf("Failed to open staging file: %v", err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			t.Fatalf("Failed to close staging file: %v", err)
		}
	}()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("Failed to read staging file: %v", err)
	}

	if len(records) != 1 {
		t.Fatalf("Expected 1 record, got %d", len(records))
	}

	hash := sha1.New()
	hash.Write(content)
	expectedHash := hex.EncodeToString(hash.Sum(nil))
	if records[0][0] != expectedHash {
		t.Fatalf("Expected 'Hello World', got '%s'", records[0][0])
	}

	if records[0][1] != "file.txt" {
		t.Fatalf("Expected 'file.txt', got '%s'", records[0][1])
	}
}

func TestAdd_NoFile(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "file.txt")

	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	err = os.Chdir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to change working directory: %v", err)
	}
	defer func() {
		err = os.Chdir(oldWd)
		if err != nil {
			t.Fatalf("Failed to change working directory: %v", err)
		}
	}()

	//Should stop test
	err = cmd.Add(filePath)
	if err == nil {
		t.Fatalf("Expected an error, got nil")
	}
}
