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

func ReverseStringSlice(input []string) {
	if len(input) == 0 {
		return
	}
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
}
