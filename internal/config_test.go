package internal

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"regexp"
	"strings"
	"testing"
)

func TestConfig(t *testing.T) {
	assert.Equal(t, Options.Addr, ":8080")
	fmt.Println("root:", Options.Root)
	fmt.Println("root:", Options.Mcss)
}

func TestString(t *testing.T) {
	str := " a b  c   d  "
	reg, _ := regexp.Compile("\\s{2,}")
	s := reg.ReplaceAllString(strings.TrimSpace(str), " ")
	arr := strings.Split(s, " ")
	fmt.Println(len(arr))
	fmt.Println(arr)
}
