package utils

import (
	"github.com/joaberch/got/cmd"
	"github.com/joaberch/got/internal/model"
	"github.com/joaberch/got/utils"
	"os"
	"path/filepath"
	"testing"
)

func TestPushLocalTest_Success(t *testing.T) {
	tempDir := t.TempDir()
	err := os.MkdirAll(filepath.Join(tempDir, "output"), os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("os.Getwd() failed: %v", err)
	}
	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatalf("os.Chdir() failed: %v", err)
	}
	defer func() {
		err := os.Chdir(oldWd)
		if err != nil {
			t.Fatalf("os.Chdir() failed: %v", err)
		}
	}()

	err = cmd.Init()
	if err != nil {
		t.Fatalf("cmd.Init() failed: %v", err)
	}

	content := []byte("Hello World")
	err = os.WriteFile("file.txt", content, 0644)
	if err != nil {
		t.Fatalf("os.WriteFile() failed: %v", err)
	}

	err = cmd.Add("file.txt")
	if err != nil {
		t.Fatalf("cmd.Add() failed: %v", err)
	}

	err = cmd.Commit([]string{"initial commit"})
	if err != nil {
		t.Fatalf("cmd.Commit() failed: %v", err)
	}

	err = cmd.SetRemote(model.Local, []string{"set-remote", "local", "output"})
	if err != nil {
		t.Fatalf("cmd.SetRemote() failed: %v", err)
	}

	err = cmd.Push()
	if err != nil {
		t.Fatalf("cmd.Push() failed: %v", err)
	}

	commitHash, err := utils.GetLatestCommitHash()
	if err != nil {
		t.Fatalf("utils.GetLatestCommitHash() failed: %v", err)
	}
	commit, err := utils.GetCommitFromHash(commitHash)
	if err != nil {
		t.Fatalf("utils.GetCommitFromHash() failed: %v", err)
	}
	tree, err := utils.GetTreeFromCommit(commit)
	if err != nil {
		t.Fatalf("utils.GetTreeFromCommit() failed: %v", err)
	}
	var blobHash []string
	for _, entry := range tree.Entries {
		blobHash = append(blobHash, entry.Hash)
	}
	if len(blobHash) != 1 {
		t.Fatalf("len(blobHash) != 1")
	}
	if blobHash[0] == "" {
		t.Fatalf("blobHash[0] is empty")
	}
	blobOutputPath := filepath.Join("output", "objects", "blobs", blobHash[0])

	check := []string{
		blobOutputPath,
		"output/commits.csv",
		"output/head",
	}

	for _, path := range check {
		if _, err = os.Stat(path); os.IsNotExist(err) {
			t.Fatalf("os.Stat() failed: %v", err)
		}
	}
	data, err := os.ReadFile(blobOutputPath)
	if err != nil {
		t.Fatalf("os.ReadFile() failed: %v", err)
	}
	if string(data) != string(content) {
		t.Fatalf("os.ReadFile() failed: %v", err)
	}
}
