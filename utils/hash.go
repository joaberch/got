package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func HashContents(contents []byte) string {
	hash := sha1.New()
	hash.Write(contents)
	return hex.EncodeToString(hash.Sum(nil))
}
