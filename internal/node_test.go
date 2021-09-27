package internal

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	text := ""
	//text := "<!doctype html><html><body><style>h1{height:1000px;}</style><h1>Hello World</h1></body></html>"
	//text := "<div><div>abc</div></div>"
	text = "<div><label>abc</label><div><input></div><div></div></div>"
	//text = "<doc><label>abc</label><div><input></div><div></div></doc>"
	//text := "<div><p>abc</p></div>"
	nodes, err := Parse([]byte(text))
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, node := range nodes {
		fmt.Println("type ---", node.Type, " --- tag --", node.Tag, "---con---", node.Content)
	}

	fmt.Println(string(nodesToString(nodes, "")))

}
