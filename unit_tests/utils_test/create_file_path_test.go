package utils_test

import (
	"github.com/joaberch/got/utils"
	"os"
	"path"
	"testing"
)

func TestCreateFilePath_File(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := path.Join(tmpDir, "file")

	err := utils.CreateFilePath(filePath, "File")
	if err != nil {
		t.Fatal(err)
	}

	info, err := os.Stat(filePath)
	if err != nil {
		t.Fatal(err)
	}
	if info.IsDir() {
		t.Fatal("expected file")
	}
}

func TestCreateFilePath_Folder(t *testing.T) {
	tmpDir := t.TempDir()
	folderPath := path.Join(tmpDir, "nested", "folder")

	err := utils.CreateFilePath(folderPath, "Folder")
	if err != nil {
		t.Fatal(err)
	}

	info, err := os.Stat(folderPath)
	if err != nil {
		t.Fatal(err)
	}
	if !info.IsDir() {
		t.Fatal("expected folder")
	}
}
