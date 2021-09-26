package internal

import (
	"log"
	"strings"
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
		if cmd == "ap" || cmd == "dp" {
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

		if cmd == "ap" || cmd == "dp" {
			if p1 == "" {
				continue
			}
			if cmd == "dp" {
				p1 = "data-" + p1
			}
			if p2 == "" {
				(*attr)[p1] = value
			} else {
				//(*attr)[p1] = strings.ReplaceAll(strings.ReplaceAll(p2, "@v", value), "\"", "\\\"")
				(*attr)[p1] = strings.ReplaceAll(p2, "@v", value)
			}
		}
		if cmd == "tp" {
			if p1 == "" || p2 == "" {
				log.Println("warning: iset-tp must bu have head and tail")
				continue
			}
			*head = append(*head, strings.ReplaceAll(p1, "@v", value))
			*tail = append(*tail, strings.ReplaceAll(p2, "@v", value))
		}
		if cmd == "ht" {
			if p1 == "" && p2 == "" {
				log.Println("warning: iset-ht must bu have head or tail")
				continue
			}
			*head = append(*head, strings.ReplaceAll(p1, "@v", value))
			*tail = append(*tail, strings.ReplaceAll(p2, "@v", value))
		}
		if cmd == "in" {
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
