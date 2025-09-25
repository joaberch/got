package utils

import (
	"github.com/joaberch/got/internal/model"
	"github.com/joaberch/got/utils"
	"os"
	"path/filepath"
	"testing"
)

func TestGetCommitFromHash_Success(t *testing.T) {
	hash := "abc123"
	commitDir := filepath.Join(".got", "objects", "commits")
	commitPath := filepath.Join(commitDir, hash)
	content := []byte(`{"message": "test","author":"Tester"}`)

	err := os.MkdirAll(commitDir, 0755)
	if err != nil {
		t.Errorf("os.MkdirAll(%q, 0755) failed", commitDir)
	}
	defer func() {
		err = os.RemoveAll(".got")
		if err != nil {
			t.Errorf("os.RemoveAll(%q) failed", commitDir)
		}
	}()

	err = os.WriteFile(commitPath, content, 0644)
	if err != nil {
		t.Errorf("os.WriteFile(%q, content, 0644) failed", commitPath)
	}

	commit, err := utils.GetCommitFromHash(hash)
	if err != nil {
		t.Errorf("utils.GetCommitFromHash(%s) failed : %s", hash, commit.Message)
	}
	if commit.Message != "test" {
		t.Errorf("utils.GetCommitFromHash(%s) failed", hash)
	}
	if commit.Author != "Tester" {
		t.Errorf("utils.GetCommitFromHash(%s) failed", hash)
	}
}

func TestGetCommitFromHash_NoFile(t *testing.T) {
	hash := "abc123"
	commit, err := utils.GetCommitFromHash(hash)
	if err == nil {
		t.Errorf("utils.GetCommitFromHash(%s) failed", hash)
	}
	if commit != (model.Commit{}) {
		t.Errorf("utils.GetCommitFromHash(%s) failed", hash)
	}
}

func TestGetCommitFromHash_InvalidCommit(t *testing.T) {
	hash := "abc123"
	commitDir := filepath.Join(".got", "objects", "commits")
	commitPath := filepath.Join(commitDir, hash)
	content := []byte("invalid")

	err := os.MkdirAll(commitDir, 0755)
	if err != nil {
		t.Errorf("os.MkdirAll(%q, 0755) failed", commitDir)
	}
	defer func() {
		err = os.RemoveAll(".got")
		if err != nil {
			t.Errorf("os.RemoveAll(%q) failed", commitDir)
		}
	}()

	err = os.WriteFile(commitPath, content, 0644)
	if err != nil {
		t.Errorf("os.WriteFile(%q, content, 0644) failed", commitPath)
	}
	commit, err := utils.GetCommitFromHash(hash)
	if err == nil {
		t.Errorf("utils.GetCommitFromHash(%s) failed", hash)
	}
	if commit.Message != "" {
		t.Errorf("utils.GetCommitFromHash(%s) failed", hash)
	}
}
