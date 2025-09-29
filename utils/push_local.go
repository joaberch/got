package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// PushLocal push the commits data and the objects to the local host
func PushLocal(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("remote path does not exist: %s", path)
	}

	err := CopyDir(".got/objects", filepath.Join(path, "objects"))
	if err != nil {
		return fmt.Errorf("error copying remote objects: %w", err)
	}

	files := []string{"commits.csv", "head"}
	for _, file := range files {
		src := filepath.Join(".got", file)
		dst := filepath.Join(path, file)

		err = CopyFile(src, dst)
		if err != nil {
			return fmt.Errorf("error copying remote objects: %w", err)
		}
	}

	fmt.Println("Pushed remote objects to:", path)
	return nil
}
