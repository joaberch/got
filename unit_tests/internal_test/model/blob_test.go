package model_test

import (
	"github.com/joaberch/got/internal/model"
	"testing"
)

func TestBlob_GenerateHash(t *testing.T) {
	content := []byte("Hello World")
	//expectedHash1 := sha1.Sum(content)
	expectedHash2 := "0a4d55a8d778e5022fab701977c5d840bbc486d0"

	blob := model.Blob{Content: content}
	hash := blob.GenerateHash()

	if hash != expectedHash2 {
		t.Errorf("error generating hash %s", hash)
	}
	if blob.Hash != expectedHash2 {
		t.Errorf("error generating hash %s", blob.Hash)
	}
}
