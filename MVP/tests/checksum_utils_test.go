package tests

import (
	"MVP/utils"
	"crypto/sha1"
	"os"
	"testing"
)

func TestGetChecksum_SameHash(t *testing.T) {
	path := "./file.txt"
	file, err := os.Create(path)
	file.WriteString("data")
	defer os.Remove(path)
	defer file.Close()
	os.WriteFile(path, []byte("data"), 0644)
	checksum, err := utils.GetChecksum(path)
	if err != nil {
		t.Errorf("Error while getting checksum : %v", err)
	}

	if checksum == "" {
		t.Errorf("Checksum should not be empty")
	}

	expectedChecksum, err := utils.GetChecksum(path)
	if err != nil {
		t.Errorf("Error while getting checksum : %v", err)
	}

	if checksum != expectedChecksum {
		t.Errorf("Checksum should be %v, but got %v, mismatch", expectedChecksum, checksum)
	}
}

func TestGetChecksum_DifferentHash(t *testing.T) {
	path := "./file.txt"
	file, err := os.Create(path)
	if err != nil {
		t.Errorf("Error while creating file : %v", err)
	}

	file.WriteString("data")
	defer os.Remove(path)
	defer file.Close()
	os.WriteFile(path, []byte("data"), 0644)
	checksum, err := utils.GetChecksum(path)

	os.WriteFile(path, []byte("AnotherData"), 0644)
	checksum2, err := utils.GetChecksum(path)

	if checksum == checksum2 {
		t.Errorf("Checksum should not be same")
	}
}

func TestGetChecksum_FileNotFound(t *testing.T) {
	path := "./file.txt"
	checksum, err := utils.GetChecksum(path)
	if err == nil {
		t.Errorf("Error should be thrown")
	}

	if checksum != "" {
		t.Errorf("Checksum should be empty")
	}
}

func TestGetChecksum_Length(t *testing.T) {
	path := "./file.txt"
	file, err := os.Create(path)
	file.WriteString("data")
	defer os.Remove(path)
	defer file.Close()
	os.WriteFile(path, []byte("data"), 0644)
	checksum, err := utils.GetChecksum(path)
	if err != nil {
		t.Errorf("Error while getting checksum : %v", err)
	}

	if len(checksum) != sha1.Size*2 {
		t.Errorf("Checksum length should be %v, but got %v", sha1.Size*2, len(checksum))
	}
}
