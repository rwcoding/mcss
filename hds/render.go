package hds

import "strings"

func ToString(nodes []*Node) string {
	s := strings.Builder{}
	for _,v := range nodes{
		s.WriteString(v.String())
	}
	return s.String()
}