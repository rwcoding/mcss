package internal

import (
	"github.com/flosch/pongo2/v4"
	"strings"
)

func ParseFile(file string, data map[string]string) ([]byte, error) {
	tpl, err := pongo2.FromFile(file)
	if err != nil {
		return nil, err
	}

	params := pongo2.Context{
		"mcss": Options.Mcss,
	}

	for np, nd := range ParseIsetCom(&data) {
		if np == "" {
			continue
		}
		params[np] = nd
	}

	for k, v := range data {
		fv := strings.TrimSpace(v)
		if fv != "" && fv[:1] == "[" {
			r := []string{}
			for _, vv := range strings.Split(fv[1:len(fv)-1], ",") {
				r = append(r, strings.TrimSpace(vv))
			}
			params[k] = r
		} else if fv != "" && fv[:1] == "{" {
			r := map[string]string{}
			for _, vv := range strings.Split(fv[1:len(fv)-1], ",") {
				or := strings.Split(strings.TrimSpace(vv), ":")
				if len(or) == 2 {
					r[strings.TrimSpace(or[0])] = strings.TrimSpace(or[1])
				}
			}
			params[k] = r
		} else {
			params[k] = fv
		}
	}

	out, err := tpl.ExecuteBytes(params)
	if err != nil {
		return nil, err
	}

	nodes, err := Parse(out)
	if err != nil {
		return nil, err
	}
	return nodesToString(nodes, file), nil
}
