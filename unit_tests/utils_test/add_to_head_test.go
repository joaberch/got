package utils_test

import (
	"github.com/joaberch/got/utils"
	"os"
	"path/filepath"
	"testing"
)

func TestAddToHead_FileOpenError(t *testing.T) {
	invalidPath := filepath.Join("invalid/path", "head")

	err := utils.AddToHead(invalidPath, "hash123")
	if err == nil {
		t.Error("AddToHead should return an error")
	}
}

func TestAddToHead_Success(t *testing.T) {
	tmpDir := t.TempDir()
	headPath := filepath.Join(tmpDir, "head")

	err := utils.AddToHead(headPath, "hash123")
	if err != nil {
		t.Fatal(err)
	}

	content, err := os.ReadFile(headPath)
	if err != nil {
		t.Fatal(err)
	}

	if string(content) != "hash123" {
		t.Fatalf("expected 'hash123', got '%s'", string(content))
	}
}
