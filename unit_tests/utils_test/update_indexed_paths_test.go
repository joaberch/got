package utils

import (
	"github.com/joaberch/got/internal/model"
	"github.com/joaberch/got/utils"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestUpdateIndexedPaths_Success(t *testing.T) {
	indexPath := filepath.Join(".got", "indexed_paths.csv")
	err := os.MkdirAll(filepath.Dir(indexPath), os.ModePerm)
	if err != nil {
		t.Fatalf("error creating directory: %v", err)
	}
	defer func() {
		err = os.RemoveAll(".got")
		if err != nil {
			t.Fatalf("error removing directory: %v", err)
		}
	}()

	tree := model.Tree{
		Entries: []model.TreeEntry{
			{Name: "file.txt", Hash: "abc123"},
			{Name: "file.csv", Hash: "def456"},
		},
	}

	err = utils.UpdateIndexedPaths(tree)
	if err != nil {
		t.Fatalf("error updating indexed paths: %v", err)
	}

	data, err := os.ReadFile(indexPath)
	if err != nil {
		t.Fatalf("error reading file: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	expected := map[string]string{
		"file.txt": "abc123",
		"file.csv": "def456",
	}
	if len(lines) != len(expected) {
		t.Fatalf("expected %v lines, got %v", len(expected), len(lines))
	}

	for _, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			t.Errorf("expected 2 columns, got %v", len(parts))
			continue
		}
		name, hash := parts[0], parts[1]
		if expected[name] != hash {
			t.Errorf("expected %v, got %v", expected[name], hash)
		}
	}
}

func TestUpdateIndexedPaths_Append(t *testing.T) {
	indexPath := filepath.Join(".got", "indexed_paths.csv")
	err := os.MkdirAll(filepath.Dir(indexPath), os.ModePerm)
	if err != nil {
		t.Fatalf("error creating directory: %v", err)
	}
	defer func() {
		err = os.RemoveAll(".got")
		if err != nil {
			t.Fatalf("error removing directory: %v", err)
		}
	}()

	initial := []byte("file1.txt,abc123\n")
	err = os.WriteFile(indexPath, initial, os.ModePerm)
	if err != nil {
		t.Fatalf("error creating file: %v", err)
	}

	tree := model.Tree{
		Entries: []model.TreeEntry{
			{Name: "file3.txt", Hash: "ghi789"},
			{Name: "file2.csv", Hash: "def456"},
		},
	}

	err = utils.UpdateIndexedPaths(tree)
	if err != nil {
		t.Fatalf("error updating indexed paths: %v", err)
	}

	data, err := os.ReadFile(indexPath)
	if err != nil {
		t.Fatalf("error reading file: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	expected := map[string]string{
		"file1.txt": "abc123",
		"file3.txt": "ghi789",
		"file2.csv": "def456",
	}

	if len(lines) != len(expected) {
		t.Fatalf("expected %v lines, got %v", len(expected), len(lines))
	}
	for _, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 2 || expected[parts[0]] != parts[1] {
			t.Errorf("expected 2 columns, got %v", len(parts))
		}
	}
}

func TestUpdateIndexedPaths_ExistingEntry(t *testing.T) {
	indexPath := filepath.Join(".got", "indexed_paths.csv")
	err := os.MkdirAll(filepath.Dir(indexPath), os.ModePerm)
	if err != nil {
		t.Fatalf("error creating directory: %v", err)
	}
	defer func() {
		err = os.RemoveAll(".got")
		if err != nil {
			t.Fatalf("error removing directory: %v", err)
		}
	}()

	initial := []byte("file1.txt,abc123\n")
	err = os.WriteFile(indexPath, initial, os.ModePerm)
	if err != nil {
		t.Fatalf("error creating file: %v", err)
	}
	tree := model.Tree{
		Entries: []model.TreeEntry{
			{Name: "file1.txt", Hash: "def456"},
		},
	}

	err = utils.UpdateIndexedPaths(tree)
	if err != nil {
		t.Fatalf("error updating indexed paths: %v", err)
	}

	data, err := os.ReadFile(indexPath)
	if err != nil {
		t.Fatalf("error reading file: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	expected := map[string]string{
		"file1.txt": "def456",
	}

	if len(lines) != len(expected) {
		t.Fatalf("expected %v lines, got %v", len(expected), len(lines))
	}
	for _, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 2 || expected[parts[0]] != parts[1] {
			t.Errorf("expected 2 columns, got %v", len(parts))
		}
	}
}
