package utils

import (
	"fmt"
	"github.com/joaberch/got/internal/model"
	"os"
	"path/filepath"
)

func GetBlobFromHash(hash string) (model.Blob, error) {
	blobPath := filepath.Join(".got", "objects", "blobs", hash)
	data, err := os.ReadFile(blobPath)
	if err != nil {
		return model.Blob{}, fmt.Errorf("error reading blob file %s: %s", blobPath, err)
	}
	return model.Blob{Content: data}, nil
}
