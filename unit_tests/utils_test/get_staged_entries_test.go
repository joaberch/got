package utils

import (
	"github.com/joaberch/got/utils"
	"os"
	"path/filepath"
	"testing"
)

func TestGetStagedEntries_Success(t *testing.T) {
	stagingPath := filepath.Join(".got", "staging.csv")
	content := []byte("file.txt,abc123\nfile.csv,def456\n")

	err := os.MkdirAll(filepath.Dir(stagingPath), 0755)
	if err != nil {
		t.Errorf("os.MkdirAll(%q) failed", stagingPath)
	}
	defer func() {
		err = os.RemoveAll(".got")
		if err != nil {
			t.Errorf("os.RemoveAll(%q) failed", stagingPath)
		}
	}()

	err = os.WriteFile(stagingPath, content, 0644)
	if err != nil {
		t.Errorf("os.WriteFile(%q) failed", stagingPath)
	}

	entries, err := utils.GetStagedEntries()
	if err != nil {
		t.Errorf("utils.GetStagedEntries failed: %v", err)
	}

	expected := []string{"file.txt", "file.csv"}

	if len(entries) != len(expected) {
		t.Errorf("expected %d entries, got %d", len(expected), len(entries))
	}
	for index, entry := range entries {
		if entry != expected[index] {
			t.Errorf("expected entry %q, got %q", expected[index], entry)
		}
	}
}

func TestGetStagedEntries_NoFile(t *testing.T) {
	err := os.RemoveAll(".got")
	if err != nil {
		t.Errorf("os.RemoveAll(%q) failed", filepath.Join(".got", "staging.csv"))
	}

	entries, err := utils.GetStagedEntries()
	if err == nil {
		t.Errorf("utils.GetStagedEntries failed: %v", err)
	}
	if entries != nil {
		t.Errorf("utils.GetStagedEntries failed: %v", err)
	}
}

func TestGetStagedEntries_Empty(t *testing.T) {
	stagingPath := filepath.Join(".got", "staging.csv")
	content := []byte("")

	err := os.MkdirAll(filepath.Dir(stagingPath), 0755)
	if err != nil {
		t.Errorf("os.MkdirAll(%q) failed", stagingPath)
	}
	defer func() {
		err = os.RemoveAll(".got")
		if err != nil {
			t.Errorf("os.RemoveAll(%q) failed", stagingPath)
		}
	}()

	err = os.WriteFile(stagingPath, content, 0644)
	if err != nil {
		t.Errorf("os.WriteFile(%q) failed", stagingPath)
	}

	entries, err := utils.GetStagedEntries()
	if err != nil {
		t.Errorf("utils.GetStagedEntries failed: %v", err)
	}

	expected := []string{}
	if len(entries) != len(expected) {
		t.Errorf("expected %d entries, got %d", len(expected), len(entries))
	}
	for index, entry := range entries {
		if entry != expected[index] {
			t.Errorf("expected entry %q, got %q", expected[index], entry)
		}
	}
}
