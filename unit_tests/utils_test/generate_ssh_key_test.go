package utils

import (
	"github.com/joaberch/got/utils"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateSSHKey_Success(t *testing.T) {
	tempDir := t.TempDir()
	keyPath := filepath.Join(tempDir, "id_rsa")

	err := utils.GenerateSSHKey(keyPath)
	if err != nil {
		t.Fatal(err)
	}

	if _, err = os.Stat(keyPath); err != nil {
		t.Fatal(err)
	}
	if _, err = os.Stat(keyPath + ".pub"); err != nil {
		t.Fatal(err)
	}
}

func TestGenerateSSHKey_AlreadyExist(t *testing.T) {
	tempDir := t.TempDir()
	keyPath := filepath.Join(tempDir, "id_rsa")

	err := os.WriteFile(keyPath, []byte("test"), 0644) //Simulate an already existing key
	if err != nil {
		t.Fatal(err)
	}
	err = utils.GenerateSSHKey(keyPath)
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "already exists") {
		t.Fatal(err)
	}
}
