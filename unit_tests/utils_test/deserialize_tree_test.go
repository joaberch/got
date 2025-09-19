package utils_test

import (
	"github.com/joaberch/got/utils"
	"testing"
)

func TestDeserializeTree_EmptyJSON(t *testing.T) {
	jsonData := []byte(`{}`)

	tree, err := utils.DeserializeTree(jsonData)
	if err != nil {
		t.Fatal(err)
	}

	if len(tree.Entries) != 0 {
		t.Fatalf("Expected 0 entries, got %d", len(tree.Entries))
	}
}

func TestDeserializeTree_InvalidJson(t *testing.T) {
	jsonData := []byte(`<hello>hello</hello>`)

	_, err := utils.DeserializeTree(jsonData)
	if err == nil {
		t.Fatal(err)
	}
}

func TestDeserializeTree_IncompatibleJSON(t *testing.T) {
	jsonData := []byte(`{ "Name": "file.txt", "Hash": "abc123" }`)

	tree, err := utils.DeserializeTree(jsonData)
	if err != nil {
		t.Fatal(err)
	}
	if len(tree.Entries) != 0 {
		t.Fatal("expected empty but got", tree)
	}
}

func TestDeserializeTree_Success(t *testing.T) {
	jsonData := []byte(`
{
	"Entries": [
		{ "Name": "file.txt", "Hash": "abc123" },
		{ "Name": "data.csv", "Hash": "def456" }
	]
}`)

	tree, err := utils.DeserializeTree(jsonData)
	if err != nil {
		t.Fatal(err)
	}

	if len(tree.Entries) != 2 {
		t.Fatalf("Expected 2 entries, got %d", len(tree.Entries))
	}

	if tree.Entries[0].Name != "file.txt" || tree.Entries[0].Hash != "abc123" ||
		tree.Entries[1].Hash != "def456" || tree.Entries[1].Name != "data.csv" {
		t.Fatalf("Wrong entry")
	}
}
