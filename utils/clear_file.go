package utils

import "os"

// ClearFile truncates (or creates) the named file and leaves it empty.
// It writes zero bytes to the file using mode 0644 and returns any error from the underlying write.
func ClearFile(path string) error {
	return os.WriteFile(path, []byte(""), 0644)
}
