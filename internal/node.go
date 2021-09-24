package internal

import (
	"github.com/rwcoding/mcss/config"
	"github.com/rwcoding/mcss/hds"
	"log"
	"os"
	"strings"
)

func nodesToString(nodes []*hds.Node) []byte {
	s := strings.Builder{}
	for _,v := range nodes {
		s.WriteString(nodeToString(v))
	}
	return []byte(s.String())
}

func nodeToString(n *hds.Node) string {
	if n.Type == hds.TextType || n.Type == hds.DocType || n.Type == hds.NullType {
		return n.Content
	}
	if strings.Contains(n.Tag, "-") {
		file := config.Options.View + string(os.PathSeparator) + n.Tag + ".html"
		b, err := ParseFile(file, n.Attributes)
		if err != nil {
			log.Println("warning: "+err.Error())
		} else {
			return string(b)
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
		for _,v := range n.Children {
			s.WriteString(nodeToString(v))
		}
	}
	s.WriteString("</")
	s.WriteString(n.Tag)
	s.WriteString(">")
	return s.String()
}
