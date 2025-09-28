package utils

import (
	"github.com/joaberch/got/internal/model"
	"github.com/joaberch/got/utils"
	"os"
	"path/filepath"
	"testing"
)

func TestGetTreeFromCommit_Success(t *testing.T) {
	treeHash := "abc123"
	treePath := filepath.Join(".got", "objects", "trees", treeHash)
	content := []byte("{\"Entries\":[{\"Name\":\"go.mod\",\"Mode\":\"file\",\"Type\":\"blob\",\"Hash\":\"28e5edb9f5fd9dc7120bc6a423af595175f68260\"}]}")

	err := os.MkdirAll(filepath.Dir(treePath), os.ModePerm)
	if err != nil {
		t.Fatalf("error creating directory: %v", err)
	}
	defer func() {
		err = os.RemoveAll(".got")
		if err != nil {
			t.Fatalf("error removing directory: %v", err)
		}
	}()
	err = os.WriteFile(treePath, content, os.ModePerm)
	if err != nil {
		t.Fatalf("error writing file: %v", err)
	}

	commit := model.Commit{TreeHash: treeHash}
	tree, err := utils.GetTreeFromCommit(commit)
	if err != nil {
		t.Fatalf("error getting tree: %v", err)
	}
	if len(tree.Entries) != 1 {
		t.Fatalf("expected tree entry 'file.txt', got %v", tree.Entries)
	}
}

func TestGetTreeFromCommit_NoFile(t *testing.T) {
	commit := model.Commit{TreeHash: "abc123"}
	tree, err := utils.GetTreeFromCommit(commit)
	if err == nil {
		t.Fatalf("expected error getting tree")
	}
	if len(tree.Entries) != 0 {
		t.Fatalf("expected tree '{}', got %v", tree)
	}
}

func TestGetTreeFromCommit_InvalidTree(t *testing.T) {
	treeHash := "abc123"
	treePath := filepath.Join(".got", "objects", "trees", treeHash)
	content := []byte("{>")

	err := os.MkdirAll(filepath.Dir(treePath), os.ModePerm)
	if err != nil {
		t.Fatalf("error creating directory: %v", err)
	}
	err = os.WriteFile(treePath, content, os.ModePerm)
	if err != nil {
		t.Fatalf("error writing file: %v", err)
	}
	defer func() {
		err = os.RemoveAll(".got")
		if err != nil {
			t.Fatalf("error removing directory: %v", err)
		}
	}()

	commit := model.Commit{TreeHash: treeHash}
	tree, err := utils.GetTreeFromCommit(commit)
	if err == nil {
		t.Fatalf("expected error getting tree: %v", err)
	}
	if len(tree.Entries) != 0 {
		t.Fatalf("expected tree entry 'file.txt', got %v", tree.Entries)
	}
}
