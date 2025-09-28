package utils

import (
	"encoding/json"
	"github.com/joaberch/got/internal/model"
)

// DeserializeBlob parses JSON-encoded data into a model.Blob.
// The input `data` should contain JSON representing a Blob; the function
// returns the parsed Blob and any error produced by json.Unmarshal.
func DeserializeBlob(data []byte) (model.Blob, error) {
	var blob model.Blob
	err := json.Unmarshal(data, &blob)
	return blob, err
}
