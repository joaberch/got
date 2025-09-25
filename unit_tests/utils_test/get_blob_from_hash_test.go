package utils

import (
	"github.com/joaberch/got/utils"
	"os"
	"path/filepath"
	"testing"
)

func TestGetBlobFromHash_Success(t *testing.T) {
	hash := "abc123"
	blobDir := filepath.Join(".got", "objects", "blobs")
	blobPath := filepath.Join(blobDir, hash)

	expectedContent := "This is a test blob"
	err := os.MkdirAll(blobDir, os.ModePerm)
	if err != nil {
		t.Errorf("error creating directory %s: %s", blobDir, err)
	}
	defer func() {
		errRemove := os.RemoveAll(".got")
		if errRemove != nil {
			err = errRemove
		}
	}()

	err = os.WriteFile(blobPath, []byte(expectedContent), os.ModePerm)
	if err != nil {
		t.Errorf("error creating file %s: %s", blobPath, err)
	}

	blob, err := utils.GetBlobFromHash(hash)
	if err != nil {
		t.Errorf("error getting blob from hash %s: %s", hash, err)
	}
	if !(expectedContent == string(blob.Content)) {
		t.Errorf("error getting blob from hash %s: %s", hash, err)
	}
	err = os.Remove(blobPath)
	if err != nil {
		return
	}
}

func TestGetBlobFromHash_NoFile(t *testing.T) {
	hash := "abc123"
	blob, err := utils.GetBlobFromHash(hash)
	if err == nil {
		t.Errorf("error getting blob from hash %s: %s", hash, err)
	}
	if blob.Content != nil {
		t.Errorf("error getting blob from hash %s: %s", hash, err)
	}
}
