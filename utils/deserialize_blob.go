package utils

import (
	"encoding/json"
	"github.com/joaberch/got/internal/model"
)

func DeserializeBlob(data []byte) (model.Blob, error) {
	var blob model.Blob
	err := json.Unmarshal(data, &blob)
	return blob, err
}
