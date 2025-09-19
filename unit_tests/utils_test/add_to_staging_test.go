package utils_test

import (
	"encoding/csv"
	"github.com/joaberch/got/utils"
	"os"
	"path/filepath"
	"testing"
)

func TestAddToStaging_FileOpenError(t *testing.T) {
	tmpDir := t.TempDir()

	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	err = os.Chdir(tmpDir)
	if err != nil {
		return
	}
	defer func() {
		err = os.Chdir(oldWd)
		if err != nil {
			t.Fatal(err)
		}
	}()

	err = utils.AddToStaging("undefined/path/to/file.txt", "hash123")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestAddToStaging_Success(t *testing.T) {
	tmpDir := t.TempDir()
	gotDir := filepath.Join(tmpDir, ".got")
	err := os.Mkdir(gotDir, 0777)
	if err != nil {
		t.Fatal(err)
	}

	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	err = os.Chdir(tmpDir)
	if err != nil {
		return
	}
	defer func() {
		err = os.Chdir(oldWd)
		if err != nil {
			t.Fatal(err)
		}
	}()

	err = utils.AddToStaging("path/to/file.txt", "hash123")
	if err != nil {
		t.Fatal(err)
	}

	stagingPath := filepath.Join(gotDir, "staging.csv")
	file, err := os.Open(stagingPath)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(records) != 1 {
		t.Fatalf("expected 1 record, got %d", len(records))
	}

	expected := []string{"path/to/file.txt", "hash123"}
	for i, val := range expected {
		if records[0][i] != val {
			t.Fatalf("expected %s, got %s", val, records[0][i])
		}
	}
}
