package internal

import (
	"github.com/rwcoding/mcss/hds"
	"log"
	"strings"
)

func nodesToString(nodes []*hds.Node, file string) []byte {
	s := strings.Builder{}
	for _, v := range nodes {
		s.WriteString(nodeToString(v, file))
	}
	return []byte(s.String())
}

//todo 元素快捷指令
func nodeToString(n *hds.Node, file string) string {
	if n.Type == hds.TextType || n.Type == hds.DocType || n.Type == hds.NullType {
		return n.Content
	}
	if strings.Contains(n.Tag, "-") {

		if len(n.Children) > 0 {
			n.Attributes["content"] = string(nodesToString(n.Children, file))
		}

		file, err := FindComponent(n.Tag, file)
		if err != nil {
			log.Println("warning: " + err.Error())
		} else {
			b, err := ParseFile(file, n.Attributes)
			if err != nil {
				log.Println("warning: " + err.Error())
				if len(n.Children) > 0 {
					delete(n.Attributes, "content")
				}
			} else {
				return string(b)
			}
		}
	}
	s := strings.Builder{}
	s.WriteString("<")
	s.WriteString(n.Tag)
	for k, v := range n.Attributes {
		s.WriteString(" ")
		s.WriteString(k)
		s.WriteString("=")
		s.WriteString("\"")
		s.WriteString(v)
		s.WriteString("\"")
	}
	s.WriteString(">")
	if !n.NeedClose {
		return s.String()
	}
	if len(n.Children) > 0 {
		for _, v := range n.Children {
			s.WriteString(nodeToString(v, file))
		}
	}
	s.WriteString("</")
	s.WriteString(n.Tag)
	s.WriteString(">")
	return s.String()
}
