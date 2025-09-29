package utils

import (
	"github.com/joaberch/got/utils"
	"os"
	"path/filepath"
	"testing"
)

func TestGetIndexedPathsWithHash_Success(t *testing.T) {
	tempDir := t.TempDir()
	gotDir := filepath.Join(tempDir, ".got")

	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = os.Chdir(oldWd)
		if err != nil {
			t.Fatal(err)
		}
	}()

	err = os.MkdirAll(gotDir, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	indexPath := filepath.Join(gotDir, "indexed_paths.csv")
	content := "file.txt,abc123\nimage.png,def456\n"
	err = os.WriteFile(indexPath, []byte(content), os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	paths, err := utils.GetIndexedPathsWithHash()
	if err != nil {
		t.Fatal(err)
	}
	if len(paths) != 2 || paths["file.txt"] != "abc123" || paths["image.png"] != "def456" {
		t.Fatalf("expected paths %v, got %v", []string{"file.txt", "image.png"}, paths)
	}
}

func TestGetIndexedPathsWithHash_NoFile(t *testing.T) {
	paths, err := utils.GetIndexedPathsWithHash()
	if err == nil {
		t.Fatalf("expected error, got paths %v", paths)
	}
	if len(paths) != 0 {
		t.Fatalf("expected no paths, got %v", paths)
	}
}

func TestGetIndexedPathsWithHash_InvalidContent(t *testing.T) {
	tempDir := t.TempDir()
	gotDir := filepath.Join(tempDir, ".got")

	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = os.Chdir(oldWd)
		if err != nil {
			t.Fatal(err)
		}
	}()

	err = os.MkdirAll(gotDir, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	indexPath := filepath.Join(gotDir, "indexed_paths.csv")
	content := "file.txt-abc123\nimage.png,def456\ninvalid\n"
	err = os.WriteFile(indexPath, []byte(content), os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	paths, err := utils.GetIndexedPathsWithHash()
	if err != nil {
		t.Fatal(err)
	}
	if len(paths) != 1 || paths["image.png"] != "def456" { //Process valid value
		t.Fatalf("expected paths %v, got %v", []string{"file.txt", "image.png"}, paths)
	}
}
