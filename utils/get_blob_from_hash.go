package utils

import (
	"fmt"
	"github.com/joaberch/got/internal/model"
	"os"
	"path/filepath"
)

// GetBlobFromHash reads the blob file for the given hash from ".got/objects/blobs" and returns it as a model.Blob.
// The hash is treated as the blob filename; on success the Blob's Content contains the file bytes.
// Returns a non-nil error if the blob file cannot be read.
func GetBlobFromHash(hash string) (model.Blob, error) {
	blobPath := filepath.Join(".got", "objects", "blobs", hash)
	data, err := os.ReadFile(blobPath)
	if err != nil {
		return model.Blob{}, fmt.Errorf("error reading blob file %s: %s", blobPath, err)
	}
	return model.Blob{Content: data}, nil
}
