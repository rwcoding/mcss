package internal

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	text := "<!doctype html><html><body><style>h1{height:1000px;}</style><h1>Hello World</h1></body></html>"
	nodes, err := Parse([]byte(text))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("length: ", len(nodes))

	for _, node := range nodes {
		fmt.Println("type ---", node.Type, " --- tag --", node.Tag, "---con---", node.Content)
	}

	fmt.Println(nodesToString(nodes, ""))

}
