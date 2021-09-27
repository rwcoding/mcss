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
	IsClose    bool
}

func (n *Node) Add(node *Node) {
	n.Children = append(n.Children, node)
}

func (n *Node) String(fromFile string) string {
	if n.Type == TextType || n.Type == DocType || n.Type == NullType {
		return n.Content
	}

	var head []string
	var tail []string
	var innerHead []string
	var innerTail []string

	for k, _ := range n.Attributes {
		if k[:1] == "@" {
			if text, ok := Options.Iset[k[1:]]; ok && text != "" {
				ParseIset(k, text.(string), &n.Attributes, &head, &tail, &innerHead, &innerTail)
			}
		}
	}
	ReverseStringSlice(head)
	ReverseStringSlice(innerHead)

	if IsComponent(n.Tag) {

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
				return strings.Join(head, "") + string(b) + strings.Join(tail, "")
			}
		}
	}
	s := strings.Builder{}

	s.WriteString(strings.Join(head, ""))

	s.WriteString("<")
	s.WriteString(n.Tag)
	for k, v := range n.Attributes {
		s.WriteString(" ")
		s.WriteString(k)
		s.WriteString("=\"")
		s.WriteString(v)
		s.WriteString("\"")
	}
	s.WriteString(">")
	s.WriteString(strings.Join(innerHead, ""))
	if !n.NeedClose {
		s.WriteString(strings.Join(innerTail, ""))
		s.WriteString(strings.Join(tail, ""))
		return s.String()
	}

	if len(n.Children) > 0 {
		for _, v := range n.Children {
			s.WriteString(v.String(fromFile))
		}
	}

	s.WriteString(strings.Join(innerTail, ""))
	s.WriteString("</")
	s.WriteString(n.Tag)
	s.WriteString(">")

	s.WriteString(strings.Join(tail, ""))

	return s.String()
}

func nodesToString(nodes []*Node, file string) []byte {
	s := strings.Builder{}
	for _, v := range nodes {
		s.WriteString(v.String(file))
	}
	return []byte(s.String())
}
