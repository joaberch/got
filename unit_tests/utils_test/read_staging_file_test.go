package utils_test

import (
	"encoding/csv"
	"github.com/joaberch/got/utils"
	"os"
	"testing"
)

func TestReadStagingFile_EmptyFile(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "staging.csv")
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

	tree, err := utils.ReadStagingFile(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if len(tree.Entries) != 0 {
		t.Fatalf("Expected 0 entries but found %d", len(tree.Entries))
	}
}

func TestReadStagingFile_NoFile(t *testing.T) {
	_, err := utils.ReadStagingFile("path/to/file.csv")
	if err == nil {
		t.Error("Expected an error")
	}
}

func TestReadStagingFile_Success(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "staging.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = os.Remove(tmpFile.Name())
		if err != nil {
			t.Fatal(err)
		}
	}()

	writer := csv.NewWriter(tmpFile)
	err = writer.WriteAll([][]string{
		{"hash123", "path/to/file.txt"},
		{"hash456", "path/to/data.csv"},
	})
	if err != nil {
		t.Fatal(err)
	}
	writer.Flush()

	err = tmpFile.Close()
	if err != nil {
		t.Fatal(err)
	}
	tree, err := utils.ReadStagingFile(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if len(tree.Entries) != 2 {
		t.Fatalf("Expected 2 entries, got %d", len(tree.Entries))
	}

	if tree.Entries[0].Hash != "hash123" || tree.Entries[1].Hash != "hash456" ||
		tree.Entries[0].Name != "path/to/file.txt" || tree.Entries[1].Name != "path/to/data.csv" {
		t.Fatalf("Expected hash \"hash123\", got %s", tree.Entries[0].Hash)
	}
}
