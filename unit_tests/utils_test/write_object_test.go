package utils_test

import (
	"github.com/joaberch/got/utils"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteObject_Success(t *testing.T) {
	tmpDir := t.TempDir()
	objectDir := filepath.Join(tmpDir, ".got", "objects", "blobs")
	err := os.MkdirAll(objectDir, 0755)
	if err != nil {
		t.Fatalf("Error creating object dir: %v", err)
	}

	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting current working directory: %v", err)
	}
	err = os.Chdir(tmpDir)
	if err != nil {
		t.Fatalf("Error changing working directory: %v", err)
	}
	defer func() {
		err = os.Chdir(oldWd)
		if err != nil {
			t.Fatalf("Error changing working directory: %v", err)
		}
	}()

	content := []byte("Hello World")
	err = utils.WriteObject("blobs", "hash123", content)
	if err != nil {
		t.Fatalf("Error writing object: %v", err)
	}

	objectPath := filepath.Join(tmpDir, ".got", "objects", "blobs", "hash123")
	data, err := os.ReadFile(objectPath)
	if err != nil {
		t.Fatalf("Error reading object: %v", err)
	}

	if string(data) != string(content) {
		t.Fatalf("Expected %s but found %s", string(content), string(data))
	}
}

func TestWriteObject_Duplicate(t *testing.T) {
	tmpDir := t.TempDir()
	objectDir := filepath.Join(tmpDir, ".got", "objects", "blobs")
	err := os.MkdirAll(objectDir, 0755)
	if err != nil {
		t.Fatalf("Error creating object dir: %v", err)
	}

	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting current working directory: %v", err)
	}
	err = os.Chdir(tmpDir)
	if err != nil {
		t.Fatalf("Error changing working directory: %v", err)
	}
	defer func() {
		err = os.Chdir(oldWd)
		if err != nil {
			t.Fatalf("Error changing working directory: %v", err)
		}
	}()

	initialContent := []byte("Hello World")
	err = utils.WriteObject("blobs", "hash123", initialContent)
	if err != nil {
		t.Fatalf("Error writing object: %v", err)
	}

	newContent := []byte("New Content")
	err = utils.WriteObject("blobs", "hash123", newContent)
	if err != nil {
		t.Fatalf("Error writing object: %v", err)
	}

	objectPath := filepath.Join(tmpDir, ".got", "objects", "blobs", "hash123")
	data, err := os.ReadFile(objectPath)
	if err != nil {
		t.Fatalf("Error reading object: %v", err)
	}

	if string(data) != string(initialContent) {
		t.Fatalf("Expected %s but found %s", string(initialContent), string(newContent))
	}
}
