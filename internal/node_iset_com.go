package internal

import (
	"log"
	"strings"
)

const (
	ISET_DT_MAP          = "kv"
	ISET_DT_OBJECT_ARRAY = "oa"
	ISET_DT_STRING_ARRAY = "sa"
)

// ParseIsetCom 解析内置指令集 @com
func ParseIsetCom(attr *map[string]string) map[string]interface{} {
	v, ok := (*attr)["@com"]
	if !ok {
		return nil
	}
	delete(*attr, "@com")
	result := map[string]interface{}{}
	rules := strings.Split(strings.TrimSpace(v), "||")
	for _, v := range rules {
		rule := strings.Split(strings.TrimSpace(v), "|")
		if len(rule) != 3 {
			log.Println("warning: rule expected " + v)
			continue
		}
		dataType := strings.TrimSpace(rule[0])
		if dataType != ISET_DT_OBJECT_ARRAY && dataType != ISET_DT_MAP && dataType != ISET_DT_STRING_ARRAY {
			log.Println("warning: type error for rule " + v)
			continue
		}
		name := strings.TrimSpace(rule[1])
		filed := strings.Split(strings.TrimSpace(rule[2]), ",")
		if name == "" || len(filed) == 0 {
			log.Println("warning: rule expected (name or field) " + v)
			continue
		}

		var dataMap = map[string]string{}
		var dataStringArray []string
		var dataObjectArray []map[string]string

		for _, key := range filed {
			d, ok := (*attr)[strings.TrimSpace(key)]
			if !ok {
				continue
			}
			if dataType == ISET_DT_MAP {
				p := strings.Split(strings.TrimSpace(d), ":")
				if len(p) != 2 {
					continue
				}
				dataMap[strings.TrimSpace(p[0])] = strings.TrimSpace(p[1])
			}
			if dataType == ISET_DT_STRING_ARRAY {
				dataStringArray = append(dataStringArray, strings.TrimSpace(d))
			}
			if dataType == ISET_DT_OBJECT_ARRAY {
				p := strings.Split(d, ",")
				tmp := map[string]string{}
				for _, vv := range p {
					kvs := strings.Split(vv, ":")
					if len(kvs) != 2 {
						continue
					}
					tmp[strings.TrimSpace(kvs[0])] = strings.TrimSpace(kvs[1])
				}
				dataObjectArray = append(dataObjectArray, tmp)
			}
		}

		if dataType == ISET_DT_MAP {
			result[name] = dataMap
		}

		if dataType == ISET_DT_STRING_ARRAY {
			result[name] = dataStringArray
		}

		if dataType == ISET_DT_OBJECT_ARRAY {
			result[name] = dataObjectArray
		}
	}
	return result
}
