package utils

import (
	"crypto/sha1"
)

func GetChecksum(path string) string {
	h := sha1.New()
	h.Write([]byte(path))
	return string(h.Sum(nil))
}
