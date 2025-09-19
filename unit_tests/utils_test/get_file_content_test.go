package utils_test

import (
	"github.com/joaberch/got/utils"
	"os"
	"testing"
)

func TestGetFileContent_NoFile(t *testing.T) {
	content, err := utils.GetFileContent("file/not/found.txt")
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
	if content != nil {
		t.Fatalf("Expected nil, got %s", string(content))
	}
}

func TestGetFileContent_Success(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "TestGetFileContent_Success")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err := os.Remove(tmpFile.Name())
		if err != nil {
			t.Fatal(err)
		}
	}()

	expectedContent := []byte("This is the content")
	_, err = tmpFile.Write(expectedContent)
	if err != nil {
		t.Fatal(err)
	}
	err = tmpFile.Close()
	if err != nil {
		return
	}

	content, err := utils.GetFileContent(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if string(content) != string(expectedContent) {
		t.Fatalf("Expected '%s', got '%s'", string(expectedContent), string(content))
	}
}
