package readers

import (
	"bytes"
	"os"
)

func reachedEOF(b *bytes.Reader) bool {
	t := make([]byte, 1)
	if v, _ := b.Read(t); v != 0 {
		return false
	}
	return true
}

func fileExists(path string) bool {
	// if fi, err := os.Stat(path); err == nil {
	// return fi.Size(), true
	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}
