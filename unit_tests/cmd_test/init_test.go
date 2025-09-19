package cmd_test

import (
	"github.com/joaberch/got/cmd"
	"github.com/joaberch/got/internal/model"
	"os"
	"path/filepath"
	"testing"
)

func TestInit_Success(t *testing.T) {
	tempDir := t.TempDir()
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

	//FileList
	err = cmd.Init()
	if err != nil {
		t.Fatal(err)
	}

	//Check .got
	gotPath := filepath.Join(tempDir, ".got")
	if _, err = os.Stat(gotPath); os.IsNotExist(err) {
		t.Fatal(".got does not exist")
	}

	//Check each file in filelist has been created
	for name, fileType := range model.FilesList {
		fullPath := filepath.Join(gotPath, name)
		info, err := os.Stat(fullPath)
		if err != nil {
			t.Fatal(err)
		}
		if info.IsDir() && fileType == "file" {
			t.Errorf("%s should be a directory", name)
		}
		if fileType == "dir" && !info.IsDir() {
			t.Errorf("%s should be a directory", name)
		}
	}
}
