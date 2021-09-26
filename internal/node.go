package internal

import (
	"log"
	"strings"
)

type NodeType uint32

const (
	NullType NodeType = iota
	TextType
	TagType
	DocType
)

func (nt NodeType) String() string {
	if nt == TextType {
		return "text"
	} else if nt == TagType {
		return "tag"
	} else if nt == DocType {
		return "doc"
	}
	return "null"
}

type Node struct {
	Type       NodeType
	Tag        string
	NeedClose  bool
	Content    string
	Attributes map[string]string
	Children   []*Node
}

func (n *Node) Add(node *Node) {
	n.Children = append(n.Children, node)
}

func (n *Node) String(fromFile string) string {
	if n.Type == TextType || n.Type == DocType || n.Type == NullType {
		return n.Content
	}
	if strings.Contains(n.Tag, "-") {

		if len(n.Children) > 0 {
			n.Attributes["content"] = string(nodesToString(n.Children, fromFile))
		}

		file, err := FindComponent(n.Tag, fromFile)
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

	hasIf, hasIfOk := n.Attributes["@if"]
	if hasIfOk && hasIf != "" {
		s.WriteString(strings.ReplaceAll(templateIfStart, "--", hasIf))
		delete(n.Attributes, "@if")
	}

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
			s.WriteString(v.String(fromFile))
		}
	}
	s.WriteString("</")
	s.WriteString(n.Tag)
	s.WriteString(">")

	if hasIfOk && hasIf != "" {
		s.WriteString(templateIfEnd)
	}
	return s.String()
}

func nodesToString(nodes []*Node, file string) []byte {
	s := strings.Builder{}
	for _, v := range nodes {
		s.WriteString(v.String(file))
	}
	return []byte(s.String())
}
