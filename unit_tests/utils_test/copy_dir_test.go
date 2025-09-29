package utils

import (
	"github.com/joaberch/got/utils"
	"os"
	"path/filepath"
	"testing"
)

func TestCopyDir_Success(t *testing.T) {
	src := t.TempDir()
	dst := filepath.Join(src, "dest")

	subDir := filepath.Join(src, "subDir")
	err := os.MkdirAll(subDir, 0777)
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile(filepath.Join(src, "file1.txt"), []byte("test"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile(filepath.Join(subDir, "file2.txt"), []byte("test2"), 0644)

	err = utils.CopyDir(src, dst)
	if err != nil {
		t.Fatal(err)
	}

	check := []string{
		filepath.Join(dst, "file1.txt"),
		filepath.Join(dst, "subDir", "file2.txt"),
	}
	for _, path := range check {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("%s does not exist", path)
		}
	}
}

func TestCopyDir_NoSource(t *testing.T) {
	err := utils.CopyDir(filepath.Join("testdata", "no_source"), t.TempDir())
	if err == nil {
		t.Errorf("no error copying directory %s, one was expected", t.TempDir())
	}
}

func TestCopyDir_DestIsFile(t *testing.T) {
	src := t.TempDir()
	dst := filepath.Join(src, "file1.txt")

	err := os.WriteFile(dst, []byte("test"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	err = utils.CopyDir(src, dst)
	if err == nil {
		t.Errorf("no error copying file %s, one was expected", dst)
	}
}
