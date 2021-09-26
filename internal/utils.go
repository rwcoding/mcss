package internal

import (
	"os"
	"strings"
)

func FileExists(file string) bool {
	if _, err := os.Stat(file); err != nil {
		return false
	}
	return true
}

func FormatPath(path string) string {
	return strings.ReplaceAll(path, "/", string(os.PathSeparator))
}
