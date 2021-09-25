package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCmdBuild(t *testing.T) {
	err := CmdBuild()
	assert.Equal(t, err, nil)
}
