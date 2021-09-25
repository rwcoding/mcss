package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	assert.Equal(t, Options.Addr, ":8080")
	fmt.Println("root:", Options.Root)
	fmt.Println("root:", Options.Mcss)
}
