package utils

import (
	"github.com/joaberch/got/utils"
	"os"
	"path/filepath"
	"testing"
)

func TestCopyFile_Success(t *testing.T) {
	tempDir := t.TempDir()
	src := filepath.Join(tempDir, "source.txt")
	dst := filepath.Join(tempDir, "dest.txt")

	content := []byte("Hello World")
	err := os.WriteFile(src, content, 0644)
	if err != nil {
		t.Errorf("error writing file %s: %v", src, err)
	}

	err = utils.CopyFile(src, dst)
	if err != nil {
		t.Errorf("error copying file %s: %v", src, err)
	}

	data, err := os.ReadFile(dst)
	if err != nil {
		t.Errorf("error reading file %s: %v", dst, err)
	}
	if string(data) != string(content) {
		t.Errorf("error copying file %s: %v", dst, string(data))
	}
}

func TestCopyFile_NoFile(t *testing.T) {
	tempDir := t.TempDir()
	src := filepath.Join(tempDir, "source.txt")
	dst := filepath.Join(tempDir, "dest.txt")

	err := utils.CopyFile(src, dst)
	if err == nil {
		t.Errorf("error copying file %s: %v", src, err)
	}
}

func TestCopyFile_NoDst(t *testing.T) {
	tempDir := t.TempDir()
	src := filepath.Join(tempDir, "source.txt")
	dst := filepath.Join(tempDir, "missing_folder", "dest.txt")

	err := os.WriteFile(src, []byte("Hello World"), 0644)
	if err != nil {
		t.Errorf("error writing file %s: %v", src, err)
	}

	err = utils.CopyFile(src, dst)
	if err == nil {
		t.Errorf("no error copying file %s, one was expected", src)
	}
}
