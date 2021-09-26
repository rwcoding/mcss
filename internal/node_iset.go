package internal

import (
	"log"
	"strings"
)

const (
	ISET_ATTR      = "as"
	ISET_ATTR_DATA = "ds"
	ISET_TEMPLATE  = "ts"
	ISET_OUT       = "ot"
	ISET_IN        = "in"
)

func ParseIset(name, text string, attr *map[string]string, head, tail, innerHead, innerTail *[]string) {
	value := (*attr)[name]
	for _, v := range strings.Split(strings.TrimSpace(text), "||") {
		desc := strings.Split(strings.TrimSpace(v), "|")
		size := len(desc)
		if size == 0 {
			continue
		}
		cmd := strings.TrimSpace(desc[0])
		p1 := ""
		p2 := ""

		if size < 2 {
			continue
		}
		var params []string
		if cmd == ISET_ATTR || cmd == ISET_ATTR_DATA {
			params = strings.Split(strings.TrimSpace(desc[1]), ":")
		} else {
			params = desc[1:]
		}

		if len(params) >= 1 {
			p1 = strings.TrimSpace(params[0])
		}
		if len(params) == 2 {
			p2 = strings.TrimSpace(params[1])
		}

		if cmd == ISET_ATTR || cmd == ISET_ATTR_DATA {
			if p1 == "" {
				continue
			}
			if cmd == ISET_ATTR_DATA {
				p1 = "data-" + p1
			}
			if cmd == ISET_ATTR && p1 == "class" {
				if tmp, ok := (*attr)["class"]; ok {
					p2 = tmp + " " + p2
				}
			}
			if p2 == "" {
				(*attr)[p1] = value
			} else {
				//(*attr)[p1] = strings.ReplaceAll(strings.ReplaceAll(p2, "@v", value), "\"", "\\\"")
				(*attr)[p1] = strings.ReplaceAll(p2, "@v", value)
			}
		}
		if cmd == ISET_TEMPLATE {
			if p1 == "" || p2 == "" {
				log.Println("warning: iset-tp must bu have head and tail")
				continue
			}
			*head = append(*head, strings.ReplaceAll(p1, "@v", value))
			*tail = append(*tail, strings.ReplaceAll(p2, "@v", value))
		}
		if cmd == ISET_OUT {
			if p1 == "" && p2 == "" {
				log.Println("warning: iset-ht must bu have head or tail")
				continue
			}
			*head = append(*head, strings.ReplaceAll(p1, "@v", value))
			*tail = append(*tail, strings.ReplaceAll(p2, "@v", value))
		}
		if cmd == ISET_IN {
			if p1 == "" && p2 == "" {
				log.Println("warning: iset-in must bu have head or tail")
				continue
			}
			*innerHead = append(*innerHead, strings.ReplaceAll(p1, "@v", value))
			*innerTail = append(*innerTail, strings.ReplaceAll(p2, "@v", value))
		}
		delete(*attr, name)
	}
}
