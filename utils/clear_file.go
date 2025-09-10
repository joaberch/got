package utils

import "os"

func ClearFile(path string) error {
	return os.WriteFile(path, []byte(""), 0644)
}
