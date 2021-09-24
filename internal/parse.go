package internal

import (
	"github.com/flosch/pongo2/v4"
	"github.com/rwcoding/mcss/hds"
	"strings"
)

func ParseFile(file string, data map[string]string) ([]byte, error)  {
	tpl, err := pongo2.FromFile(file)
	if err != nil {
		return nil, err
	}

	params := pongo2.Context{}
	for k,v := range data {
		if v[:1] == "[" {
			r := []string{}
			for _,vv := range strings.Split(v[1:len(v)-1], ",") {
				r = append(r, strings.TrimSpace(vv))
			}
			params[k] = r
		} else if v[:1] == "{" {
			r := map[string]string{}
			for _,vv := range strings.Split(v[1:len(v)-1], ",") {
				or := strings.Split(strings.TrimSpace(vv), ":")
				if len(or) == 2  {
					r[strings.TrimSpace(or[0])] = strings.TrimSpace(or[1])
				}
			}
			params[k] = r
		} else {
			params[k] = v
		}
	}

	out, err := tpl.ExecuteBytes(params)
	if err != nil {
		return nil, err
	}
	nodes, err := hds.Parse(out)
	if err != nil {
		return nil, err
	}
	return nodesToString(nodes), nil
}