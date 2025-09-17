package utils_test

import (
	"github.com/joaberch/got/utils"
	"os"
	"testing"
)

func TestClearFile_Success(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "TestClearFile_Success")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = os.Remove(tmpFile.Name())
		if err != nil {
			t.Fatal(err)
		}
	}()

	_, err = tmpFile.WriteString("content")
	if err != nil {
		t.Fatal(err)
	}
	err = tmpFile.Close()
	if err != nil {
		t.Fatal(err)
	}

	err = utils.ClearFile(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	if len(content) != 0 {
		t.Fatalf("expected empty, got %s", content)
	}
}
