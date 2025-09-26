package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/joaberch/got/utils"
	"os"
	"path/filepath"
)

// Push the objects to the local path
func Push() error {
	data, err := os.ReadFile(".got/.gotconfig")
	if err != nil {
		return fmt.Errorf("error reading .gotconfig: %w", err)
	}

	var config map[string]string
	if err := json.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("error unmarshaling .gotconfig: %w", err)
	}

	remotePath, ok := config["local"]
	if !ok {
		return fmt.Errorf("no 'local' key in .gotconfig")
	}

	if _, err = os.Stat(remotePath); os.IsNotExist(err) {
		return fmt.Errorf("remote path does not exist: %s", remotePath)
	}

	err = utils.CopyDir(".got/objects", filepath.Join(remotePath, "objects"))
	if err != nil {
		return fmt.Errorf("error copying remote objects: %w", err)
	}

	files := []string{"commits.csv", "head"}
	for _, file := range files {
		src := filepath.Join(".got", file)
		dst := filepath.Join(remotePath, file)

		err = utils.CopyFile(src, dst)
		if err != nil {
			return fmt.Errorf("error copying remote objects: %w", err)
		}
	}

	fmt.Println("Pushed remote objects to:", remotePath)
	return nil
}
