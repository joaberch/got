package utils

import (
	"Got/internal/model"
)

// CreateBlob constructs a Git-style blob object from the contents given
func CreateBlob(contents []byte) model.Blob {
	blob := model.Blob{
		Content: contents,
	}
	blob.GenerateHash()
	return blob
}
