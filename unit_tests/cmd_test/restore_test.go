package cmd_test

import (
	"github.com/joaberch/got/cmd"
	"github.com/joaberch/got/internal/model"
	"os"
	"path/filepath"
	"testing"
)

func TestRestore_Success(t *testing.T) {
	//Setup
	tempDir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = os.Chdir(oldWd)
		if err != nil {
			t.Fatal(err)
		}
	}()
	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatal(err)
	}
	base := filepath.Join(tempDir, ".got", "objects")
	err = os.MkdirAll(filepath.Join(base, "commits"), 0755)
	if err != nil {
		t.Errorf("os.MkdirAll(%s) error %v", base, err)
	}
	err = os.MkdirAll(filepath.Join(base, "trees"), 0755)
	if err != nil {
		t.Errorf("os.MkdirAll(%s) error %v", base, err)
	}
	err = os.MkdirAll(filepath.Join(base, "blobs"), 0755)
	if err != nil {
		t.Errorf("os.MkdirAll(%s) error %v", base, err)
	}

	//Create blob
	blobContent := []byte("blob content")
	blobHash := "blob123"
	blobPath := filepath.Join(base, "blobs", blobHash)
	err = os.WriteFile(blobPath, blobContent, 0644)
	if err != nil {
		t.Errorf("os.WriteFile error: %v", err)
	}

	//Tree with blob in entry
	tree := model.Tree{
		Entries: []model.TreeEntry{
			{Name: "tree1", Hash: blobHash},
		},
	}
	treeData, err := tree.Serialize()
	if err != nil {
		t.Errorf("tree.Serialize() error: %v", err)
	}
	treeHash := "tree123"
	treePath := filepath.Join(base, "trees", treeHash)
	err = os.WriteFile(treePath, treeData, 0644)
	if err != nil {
		t.Errorf("os.WriteFile error: %v", err)
	}

	//Commit with tree
	commit := model.Commit{
		TreeHash: treeHash,
		Message:  "test message",
	}
	commitData, err := commit.Serialize()
	if err != nil {
		t.Errorf("commit.Serialize() error: %v", err)
	}
	commitHash := "commit789"
	commitPath := filepath.Join(base, "commits", commitHash)
	err = os.WriteFile(commitPath, commitData, 0644)
	if err != nil {
		t.Errorf("os.WriteFile error: %v", err)
	}

	//Restore commit
	err = cmd.Restore(commitHash)
	if err != nil {
		t.Errorf("cmd.Restore() error: %v", err)
	}

	restoredPath := filepath.Join(tempDir, "tree1")
	data, err := os.ReadFile(restoredPath)
	if err != nil {
		t.Errorf("os.ReadFile error: %v", err)
	}
	if string(data) != string(blobContent) {
		t.Errorf("restored.txt content not equal to blob content")
	}
}

func TestRestore_BlobMissing(t *testing.T) {
	//Setup
	tempDir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = os.Chdir(oldWd)
		if err != nil {
			t.Fatal(err)
		}
	}()
	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatal(err)
	}
	//Folder setup
	base := filepath.Join(tempDir, ".got", "objects")
	err = os.MkdirAll(filepath.Join(base, "commits"), 0755)
	if err != nil {
		t.Errorf("os.MkdirAll(%s) error %v", base, err)
	}
	err = os.MkdirAll(filepath.Join(base, "trees"), 0755)
	if err != nil {
		t.Errorf("os.MkdirAll(%s) error %v", base, err)
	}
	err = os.MkdirAll(filepath.Join(base, "blobs"), 0755)
	if err != nil {
		t.Errorf("os.MkdirAll(%s) error %v", base, err)
	}

	//Tree with no blob
	tree := model.Tree{
		Entries: []model.TreeEntry{
			{Name: "missing", Hash: "blob123"},
		},
	}
	treeData, err := tree.Serialize()
	if err != nil {
		t.Errorf("tree.Serialize() error: %v", err)
	}
	treeHash := "tree123"
	treePath := filepath.Join(base, "trees", treeHash)
	err = os.WriteFile(treePath, treeData, 0644)
	if err != nil {
		t.Errorf("os.WriteFile error: %v", err)
	}

	//Commit to tree
	commit := model.Commit{
		TreeHash: treeHash,
		Message:  "test message",
	}
	commitData, err := commit.Serialize()
	if err != nil {
		t.Errorf("commit.Serialize() error: %v", err)
	}
	commitHash := "commit789"
	commitPath := filepath.Join(base, "commits", commitHash)
	err = os.WriteFile(commitPath, commitData, 0644)
	if err != nil {
		t.Errorf("os.WriteFile error: %v", err)
	}

	//Act
	err = cmd.Restore(commitHash)
	if err == nil {
		t.Fatalf("Expected error but got none")
	}
}
