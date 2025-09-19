package cmd_test

import (
	"bytes"
	"github.com/joaberch/got/cmd"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestLog_Success(t *testing.T) {
	tempDir := t.TempDir()
	gotDir := filepath.Join(tempDir, ".got")
	err := os.Mkdir(gotDir, 0777)
	if err != nil {
		t.Fatalf("os.Mkdir(%q) failed: %v", gotDir, err)
	}

	commitsPath := filepath.Join(gotDir, "commits.csv")
	//hash, treeHash, Author, Message, Timestamp
	content := `abc123,tree1,Tester,Initial commit,3953895789
def456,tree2,Tester,second commit,738275835`
	err = os.WriteFile(commitsPath, []byte(content), 0777)
	if err != nil {
		t.Fatalf("os.WriteFile(%q) failed: %v", commitsPath, err)
	}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("os.Getwd() failed: %v", err)
	}
	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatalf("os.Chdir() failed: %v", err)
	}
	defer func() {
		err = os.Chdir(oldWd)
		if err != nil {
			t.Fatalf("os.Chdir() failed: %v", err)
		}
	}()

	err = cmd.Log()
	if err != nil {
		t.Fatalf("cmd.Log() failed: %v", err)
	}

	err = w.Close()
	if err != nil {
		t.Fatalf("w.Close() failed: %v", err)
	}
	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	if err != nil {
		t.Fatalf("io.Copy() failed: %v", err)
	}
	os.Stdout = oldStdout

	output := buf.String()

	if !contains(output, "Commit : abc123") || !contains(output, "Commit : def456") {
		t.Fatalf("output does not contain expected commit")
	}
}

func contains(s, substr string) bool {
	return bytes.Contains([]byte(s), []byte(substr))
}
