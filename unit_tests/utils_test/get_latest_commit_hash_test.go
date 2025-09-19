package utils_test

import (
	"github.com/joaberch/got/utils"
	"os"
	"path/filepath"
	"testing"
)

func TestGetLatestCommitHash_NoFile(t *testing.T) {
	tmpDir := t.TempDir()

	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Could not get working directory: %s", err)
	}
	defer func() {
		err = os.Chdir(oldWd)
		if err != nil {
			t.Fatalf("Could not chdir to %s: %s", oldWd, err)
		}
	}()
	err = os.Chdir(tmpDir)
	if err != nil {
		t.Fatalf("Could not chdir to %s: %s", tmpDir, err)
	}

	_, err = utils.GetLatestCommitHash()
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestGetLatestCommitHash_EmptyFile(t *testing.T) {
	tmpDir := os.TempDir()
	gotDir := filepath.Join(tmpDir, ".got")
	err := os.MkdirAll(gotDir, os.ModePerm)
	if err != nil {
		t.Fatalf("Could not create directory %s: %s", gotDir, err)
	}
	defer func() {
		err = os.RemoveAll(gotDir)
		if err != nil {
			t.Fatalf("Could not remove directory %s: %s", gotDir, err)
		}
	}()

	headPath := filepath.Join(gotDir, "head")
	err = os.WriteFile(headPath, []byte(""), 0644)
	if err != nil {
		t.Fatalf("Could not create file %s: %s", headPath, err)
	}

	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Could not get working directory: %s", err)
	}
	defer func() {
		err = os.Chdir(oldWd)
		if err != nil {
			t.Fatalf("Could not chdir to %s: %s", oldWd, err)
		}
	}()
	err = os.Chdir(tmpDir)
	if err != nil {
		t.Fatalf("Could not chdir to %s: %s", tmpDir, err)
	}

	hash, err := utils.GetLatestCommitHash()
	if err != nil {
		t.Fatalf("Could not get latest commit hash: %s", err)
	}

	if hash != "" {
		t.Fatalf("Could not get latest commit hash")
	}
}

func TestGetLatestCommitHash_Success(t *testing.T) {
	tmpDir := os.TempDir()
	gotDir := filepath.Join(tmpDir, ".got")
	err := os.MkdirAll(gotDir, os.ModePerm)
	if err != nil {
		t.Fatalf("Could not create directory %s: %s", gotDir, err)
	}
	defer func() {
		err = os.RemoveAll(gotDir)
		if err != nil {
			t.Fatalf("Could not remove directory %s: %s", gotDir, err)
		}
	}()

	headPath := filepath.Join(gotDir, "head")
	expected := "abc123"
	err = os.WriteFile(headPath, []byte(expected), 0644)
	if err != nil {
		t.Fatalf("Could not create file %s: %s", headPath, err)
	}

	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Could not get working directory: %s", err)
	}
	defer func() {
		err = os.Chdir(oldWd)
		if err != nil {
			t.Fatalf("Could not chdir to %s: %s", oldWd, err)
		}
	}()
	err = os.Chdir(tmpDir)
	if err != nil {
		t.Fatalf("Could not chdir to %s: %s", tmpDir, err)
	}

	hash, err := utils.GetLatestCommitHash()
	if err != nil {
		t.Fatalf("Could not get latest commit hash: %s", err)
	}

	if hash != expected {
		t.Fatalf("Expected %q, got %q", expected, hash)
	}
}
