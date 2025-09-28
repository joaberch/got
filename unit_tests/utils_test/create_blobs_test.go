package utils_test

import (
	"github.com/joaberch/got/internal/model"
	"github.com/joaberch/got/utils"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateBlobs_ReadError(t *testing.T) {
	base := t.TempDir()
	tree := model.Tree{
		Entries: []model.TreeEntry{
			{
				Name: filepath.Join(base, "nonexistent", "file.txt"),
				Hash: "blob123",
			},
		},
	}

	err := utils.CreateBlobs(tree)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestCreateBlobs_Success(t *testing.T) {
	tmpDir := t.TempDir()

	filePath := filepath.Join(tmpDir, "file.txt")
	content := []byte("this is a test file")
	err := os.WriteFile(filePath, content, 0666)
	if err != nil {
		t.Fatal(err)
	}

	blobsDir := filepath.Join(tmpDir, ".got", "objects", "blobs")
	err = os.MkdirAll(blobsDir, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	tree := model.Tree{
		Entries: []model.TreeEntry{
			{
				Name: filePath,
				Hash: "blob123",
			},
		},
	}
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	err = os.Chdir(tmpDir)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = os.Chdir(oldWd)
		if err != nil {
			t.Fatal(err)
		}
	}()

	err = utils.CreateBlobs(tree)
	if err != nil {
		t.Fatal(err)
	}

	blobPath := filepath.Join(".got", "objects", "blobs", "blob123")
	blobContent, err := os.ReadFile(blobPath)
	if err != nil {
		t.Fatal(err)
	}

	if string(blobContent) != string(content) {
		t.Errorf("blob content does not match")
	}
}
