package utils

import (
	"os"
)

// AddToHead clears the .got/head file and writes hash as the current HEAD commit hash.
// It first empties the head file (via ClearFile) and then writes the provided hash (written as-is; no trailing newline).
// Any error encountered while clearing, opening, writing, or closing the file causes the program to exit via log.Fatal.
func AddToHead(headPath string, hash string) error {
	err := ClearFile(headPath)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(headPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			err = closeErr
		}
	}()

	_, err = file.WriteString(hash)
	if err != nil {
		return err
	}
	return nil
}
