package cmd_test

import (
	"bytes"
	"github.com/joaberch/got/cmd"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDiff_AddFile(t *testing.T) {
	tmpDir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	err = os.Chdir(tmpDir)
	if err != nil {
		t.Fatalf("failed to change to temp dir: %v", err)
	}
	defer func() {
		err = os.Chdir(oldWd)
		if err != nil {
			t.Fatalf("failed to change to original dir: %v", err)
		}
	}()

	err = cmd.Init()
	if err != nil {
		t.Fatalf("failed to init command: %v", err)
	}

	// Create file.txt with initial content
	err = os.WriteFile("file.txt", []byte("Hello World1"), 0644)
	if err != nil {
		t.Fatalf("Failed to write file.txt: %v", err)
	}

	err = cmd.Add(filepath.Join(tmpDir, "file.txt"))
	if err != nil {
		t.Fatalf("failed to add file: %v", err)
	}

	err = cmd.Commit([]string{"initial"})
	if err != nil {
		t.Fatalf("failed to commit: %v", err)
	}

	err = os.WriteFile("newfile.txt", []byte("Hello World1"), 0644)
	if err != nil {
		t.Fatalf("Failed to write file.txt: %v", err)
	}
	err = cmd.Add(filepath.Join(tmpDir, "newfile.txt"))
	if err != nil {
		t.Fatalf("failed to add file: %v", err)
	}

	err = cmd.Diff(false)
	if err != nil {
		t.Fatalf("failed to diff: %v", err)
	}
}

func TestDiff_ChangeFileSuccess(t *testing.T) {
	tmpDir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	err = os.Chdir(tmpDir)
	if err != nil {
		t.Fatalf("failed to change to temp dir: %v", err)
	}
	defer func() {
		err = os.Chdir(oldWd)
		if err != nil {
			t.Fatalf("failed to change to original dir: %v", err)
		}
	}()

	// Create .got structure
	err = os.MkdirAll(".got/objects/blobs", os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to create blobs dir: %v", err)
	}
	err = os.MkdirAll(".got/objects/commits", os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to create commits dir: %v", err)
	}
	err = os.MkdirAll(".got/objects/trees", os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to create trees dir: %v", err)
	}

	indexedPath := filepath.Join(".got", "indexed_paths.csv")
	err = os.WriteFile(indexedPath, []byte("file.txt,abc123\n"), 0644)
	if err != nil {
		t.Fatalf("Failed to write indexed paths file: %v", err)
	}

	// Create file.txt with initial content
	err = os.WriteFile("file.txt", []byte("Hello World"), 0644)
	if err != nil {
		t.Fatalf("Failed to write file.txt: %v", err)
	}

	// Simulate blob object
	blobHash := "abc123"
	blobPath := filepath.Join(".got", "objects", "blobs", blobHash)
	err = os.WriteFile(blobPath, []byte("Hello World"), 0644)
	if err != nil {
		t.Fatalf("Failed to write blob: %v", err)
	}

	// Simulate tree object
	treeHash := "tree123"
	treePath := filepath.Join(".got", "objects", "trees", treeHash)
	treeContent := `{"Entries": [{"Name": "file.txt", "Hash": "abc123"}]}`
	err = os.WriteFile(treePath, []byte(treeContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write tree: %v", err)
	}

	// Simulate commit object
	commitHash := "commit123"
	commitPath := filepath.Join(".got", "objects", "commits", commitHash)
	commitContent := `{"TreeHash": "tree123", "Message": "Initial commit"}`

	err = os.WriteFile(commitPath, []byte(commitContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write commit: %v", err)
	}

	// Write head file
	headPath := filepath.Join(".got", "head")
	err = os.WriteFile(headPath, []byte(commitHash), 0644)
	if err != nil {
		t.Fatalf("Failed to write head file: %v", err)
	}

	// Simulate index file
	indexPath := filepath.Join(".got", "index.csv")
	err = os.WriteFile(indexPath, []byte("file.txt,abc123\n"), 0644)
	if err != nil {
		t.Fatalf("Failed to write index file: %v", err)
	}

	// Simulate staging file
	stagingPath := filepath.Join(".got", "staging.csv")
	err = os.WriteFile(stagingPath, []byte("file.txt,abc123\n"), 0644)
	if err != nil {
		t.Fatalf("Failed to write staging file: %v", err)
	}

	// Modify file.txt to simulate a diff
	err = os.WriteFile("file.txt", []byte("Hello Universe"), 0644)
	if err != nil {
		t.Fatalf("Failed to modify file.txt: %v", err)
	}

	// Capture stdout
	var buf bytes.Buffer
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Act
	err = cmd.Diff(false)
	if err != nil {
		t.Fatalf("Diff failed: %v", err)
	}

	// Restore stdout
	err = w.Close()
	if err != nil {
		t.Fatalf("Diff failed: %v", err)
	}
	os.Stdout = stdout
	_, err = buf.ReadFrom(r)
	if err != nil {
		t.Fatalf("Diff failed: %v", err)
	}
	output := buf.String()

	// Assert
	if !strings.Contains(output, "Modified file: file.txt") {
		t.Errorf("Expected diff output for modified file, got: %s", output)
	}
}
