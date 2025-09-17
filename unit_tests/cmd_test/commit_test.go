package cmd

import (
	"github.com/joaberch/got/cmd"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCommit_Success(t *testing.T) {
	//Setup directory
	tmpDir := t.TempDir()
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

	//Create .got
	err = os.Mkdir(".got", os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to create .got directory: %v", err)
	}
	defer func() {
		err = os.RemoveAll(".got")
		if err != nil {
			t.Error()
		}
	}()

	//Create .got/objects/blobs
	err = os.MkdirAll(".got/objects/blobs", os.ModePerm)
	//Create .got/objects/commits
	err = os.MkdirAll(".got/objects/commits", os.ModePerm)
	//Create .got/objects/trees
	err = os.MkdirAll(".got/objects/trees", os.ModePerm)

	//Create file.txt
	file, err := os.Create("file.txt")
	if err != nil {
		t.Fatalf("Failed to create file.txt: %v", err)
	}
	_, err = file.Write([]byte("Hello World"))
	if err != nil {
		t.Fatalf("Failed to write file.txt: %v", err)
	}
	err = file.Close()
	if err != nil {
		t.Fatalf("Failed to close file.txt: %v", err)
	}

	//Create staging file
	stagingPath := filepath.Join(tmpDir, ".got", "staging.csv")
	err = os.WriteFile(stagingPath, []byte("file.txt,abc123\n"), os.ModePerm) //Format : file_path,blob_name
	if err != nil {
		t.Fatalf("Failed to write staging file: %v", err)
	}
	//Create commit file
	commitsPath := filepath.Join(tmpDir, ".got", "commits.csv")
	err = os.WriteFile(commitsPath, []byte(""), os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to write commit file: %v", err)
	}
	//Create head file
	headPath := filepath.Join(tmpDir, ".got", "head")
	err = os.WriteFile(headPath, []byte(""), os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to write head file: %v", err)
	}

	//Act
	cmd.Commit("My message")

	//Check empty staging file
	content, err := os.ReadFile(stagingPath)
	if err != nil {
		t.Fatalf("Failed to read staging file: %v", err)
	}
	if len(content) != 0 {
		t.Fatalf("Expected empty staging file")
	}

	//Check commits file has 1 entry
	content, err = os.ReadFile(commitsPath)
	lines := string(content)
	if !strings.Contains(lines, "My message") {
		t.Fatalf("Expected 'My message' but got '%v'", lines)
	}
	if err != nil {
		t.Fatalf("Failed to read commit file: %v", err)
	}
	if len(content) < 1 {
		t.Fatalf("Expected multiple entry on commit file %v", len(content))
	}

	//Check head file contains commit hash
	content, err = os.ReadFile(headPath)
	if err != nil {
		t.Fatalf("Failed to read head file: %v", err)
	}
	if len(content) < 1 {
		t.Fatalf("Expected multiple entry on head file %v", len(content))
	}

	//Check commit objects
	objectsCommitsPath := filepath.Join(tmpDir, ".got", "objects", "commits")
	entries, err := os.ReadDir(objectsCommitsPath)
	if err != nil {
		t.Fatalf("Failed to read objects commits directory: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("Expected one entry on objects directory %v", len(entries))
	}

	//Check tree objects
	objectsTreePath := filepath.Join(tmpDir, ".got", "objects", "trees")
	entries, err = os.ReadDir(objectsTreePath)
	if err != nil {
		t.Fatalf("Failed to read objects tree directory: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("Expected one entry on objects directory %v", len(entries))
	}

	//Check blob objects
	objectsBlobsPath := filepath.Join(tmpDir, ".got", "objects", "blobs")
	entries, err = os.ReadDir(objectsBlobsPath)
	if err != nil {
		t.Fatalf("Failed to read objects blobs directory: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("Expected one entry on objects directory %v", len(entries))
	}
}

func TestCommit_NoFolder(t *testing.T) {
	tmpDir := t.TempDir()
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

	err = cmd.Commit("My message")
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}
}
