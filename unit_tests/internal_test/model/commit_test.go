package model_test

import (
	"github.com/joaberch/got/internal/model"
	"testing"
	"time"
)

func TestCommit_SerializeAndHash(t *testing.T) {
	commit := model.Commit{
		TreeHash:   "abc123",
		ParentHash: "def456",
		Author:     "tester",
		Message:    "test message",
		Timestamp:  time.Now().Unix(),
	}

	data, err := commit.Serialize()
	if err != nil {
		t.Error(err)
	}
	if len(data) == 0 {
		t.Error("data is empty")
	}

	hash := commit.Hash(data)
	if len(hash) != 40 {
		t.Error("hash length is wrong")
	}
}
