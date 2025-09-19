package utils_test

import (
	"github.com/joaberch/got/utils"
	"testing"
)

func TestDeserializeCommit_InvalidJSON(t *testing.T) {
	jsonData := []byte(`{<?xml version="1.0" encoding="UTF-8" ?>
 <root>
     <TreeHash>abc123</TreeHash>
     <Author>Tester</Author>
     <Message>Hello World!</Message>
     <Timestamp>1694956800</Timestamp>
 </root>
}`)

	_, err := utils.DeserializeCommit(jsonData)
	if err == nil {
		t.Fatal(err)
	}
}

func TestDeserializeCommit_Success(t *testing.T) {
	jsonData := []byte(`{
		"TreeHash": "abc123",
		"Author": "Tester",
		"Message": "Hello World!",
		"Timestamp": 1694956800}`)

	commit, err := utils.DeserializeCommit(jsonData)
	if err != nil {
		t.Fatal(err)
	}

	if commit.TreeHash != "abc123" || commit.Author != "Tester" || commit.Message != "Hello World!" || commit.Timestamp != 1694956800 {
		t.Fatal("Unexpected result in", commit)
	}
}

func TestDeserializeCommit_IncompleteJSON(t *testing.T) {
	jsonData := []byte(`{"Author": "Tester", "Message": "Hello World!"}`)
	_, err := utils.DeserializeCommit(jsonData)
	if err != nil {
		t.Fatal(err)
	}
}
