package internal

import "strings"

func IsComponent(tag string) bool {
	return strings.Contains(tag, "-")
}
