package hds

import "strings"

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
	Type NodeType
	Tag string
	NeedClose bool
	Content string
	Attributes map[string]string
	Children []*Node
}

func (n *Node) Add(node *Node) {
	n.Children = append(n.Children, node)
}

func (n *Node) String() string {
	if n.Type == TextType || n.Type == DocType || n.Type == NullType {
		return n.Content
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
			s.WriteString(v.String())
		}
	}
	s.WriteString("</")
	s.WriteString(n.Tag)
	s.WriteString(">")
	return s.String()
}


