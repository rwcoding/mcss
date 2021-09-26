package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseIset(t *testing.T) {
	attr := map[string]string{
		"@bind": "age",
		"id":    "ele",
	}
	text := "ap|class:bind || dp|bind || tp|{{ if user }}|{{ endif }} || ht|<start>|<end>"
	name := "@bind"

	var head []string
	var tail []string
	var innerHead []string
	var innerTail []string

	ParseIset(name, text, &attr, &head, &tail, &innerHead, &innerTail)

	ReverseStringSlice(head)

	assert.Equal(t, attr["data-bind"], "age")
	assert.Equal(t, attr["class"], "bind")
	assert.Equal(t, head[0], "<start>")
	assert.Equal(t, tail[1], "<end>")
}
