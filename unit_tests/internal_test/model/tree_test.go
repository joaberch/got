package model_test

import (
	"github.com/joaberch/got/internal/model"
	"testing"
)

func TestTree_GenerateHashAndSerialize(t *testing.T) {
	tree := model.Tree{
		Entries: []model.TreeEntry{
			{Name: "file.txt", Hash: "abc123"},
			{Name: "image.png", Hash: "def456"},
		},
	}

	hash, err := tree.GenerateHash()
	if err != nil {
		t.Fatal(err)
	}
	if len(hash) != 40 {
		t.Fatalf("hash length should be 40 got %d", len(hash))
	}

	data, err := tree.Serialize()
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Fatalf("data is empty")
	}
}
