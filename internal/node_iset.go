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

func ParseIset(name string, rules []interface{}, attr *map[string]string, head, tail, innerHead, innerTail *[]string) {
	value := (*attr)[name]

	if _, ok := rules[0].(string); ok {
		rules = append([]interface{}{}, rules)
	}

	for _, v := range rules {
		desc := v.([]interface{})
		if _, ok := desc[0].(string); !ok {
			continue
		}
		if len(desc) < 2 {
			continue
		}
		cmd := strings.TrimSpace(desc[0].(string))

		for kk, vv := range desc {
			if kk == 0 {
				continue
			}
			if cmd == ISET_ATTR || cmd == ISET_ATTR_DATA {
				if item, ok := vv.(map[string]string); ok {
					for _k, _v := range item {
						if cmd == ISET_ATTR_DATA {
							_k = "data-" + _k
						}
						if cmd == ISET_ATTR && _k == "class" {
							if tmp, ok := (*attr)["class"]; ok {
								_v = tmp + " " + _v
							}
						}
						(*attr)[_k] = strings.ReplaceAll(_v, "@v", value)
					}
				}
			}

			if cmd == ISET_OUT || cmd == ISET_TEMPLATE {
				item, ok := vv.(string)
				if !ok {
					continue
				}
				if kk == 1 {
					*head = append(*head, strings.ReplaceAll(item, "@v", value))
				}
				if kk == 2 {
					*tail = append(*tail, strings.ReplaceAll(item, "@v", value))
				}
			}

			if cmd == ISET_IN {
				item, ok := vv.(string)
				if !ok {
					continue
				}
				if kk == 1 {
					*innerHead = append(*innerHead, strings.ReplaceAll(item, "@v", value))
				}
				if kk == 2 {
					*innerTail = append(*innerTail, strings.ReplaceAll(item, "@v", value))
				}
			}
		}

		delete(*attr, name)
	}
}

func ParseIsetV1(name, text string, attr *map[string]string, head, tail, innerHead, innerTail *[]string) {
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
