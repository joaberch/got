package utils

import (
	"fmt"
	"os"
)

// AddToHead clears the .got/head file and writes hash as the current HEAD commit hash.
// It first empties the head file (via ClearFile) and then writes the provided hash (written as-is; no trailing newline).
// AddToHead clears the file at headPath and writes hash as the new HEAD value.
// 
// AddToHead removes any existing content in headPath, opens (or creates) the file
// with mode 0644, and writes the provided hash exactly as given (no trailing
// newline). It returns a wrapped error if clearing the file, opening it, or
// writing the hash fails. The function attempts to close the file when done;
// any close error is observed but may not be propagated to the caller.
func AddToHead(headPath string, hash string) error {
	err := ClearFile(headPath)
	if err != nil {
		return fmt.Errorf("failed to clear file %s: %w", headPath, err)
	}

	file, err := os.OpenFile(headPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", headPath, err)
	}
	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			err = fmt.Errorf("failed to close file %s: %w", headPath, closeErr)
		}
	}()

	_, err = file.WriteString(hash)
	if err != nil {
		return fmt.Errorf("failed to write to commits file at %s: %w", headPath, err)
	}
	return nil
}
