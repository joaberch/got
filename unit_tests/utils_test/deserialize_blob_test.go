package utils

import (
	"encoding/json"
	"github.com/joaberch/got/internal/model"
	"github.com/joaberch/got/utils"
	"testing"
)

func TestDeserializeBlob_Success(t *testing.T) {
	original := model.Blob{
		Content: []byte("This is a test"),
	}
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("json.Marshal(original) failed: %v", err)
	}

	blob, err := utils.DeserializeBlob(data)
	if err != nil {
		t.Fatalf("DeserializeBlob(original) failed: %v", err)
	}
	if string(blob.Content) != "This is a test" {
		t.Fatalf("Expected 'This is a test', got '%s'", string(blob.Content))
	}
}

func TestDeserializeBlob_InvalidJSON(t *testing.T) {
	data := []byte(`{invalid json}`)
	_, err := utils.DeserializeBlob(data)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestDeserializeBlob_EmptyJSON(t *testing.T) {
	data := []byte(``)
	_, err := utils.DeserializeBlob(data)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}
