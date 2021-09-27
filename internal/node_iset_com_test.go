package internal

import (
	"fmt"
	"testing"
)

func TestParseIsetCom(t *testing.T) {

	attr := map[string]string{
		"@com": "sa|users|s1,s2,s3 || oa|objs|o1,o2,o3 || kv|kvs|k1,k2,k3",
		"s1":   "s-1",
		"s2":   "s-2",
		"s3":   "s-3",

		"o1": "id:1,name:lucy",
		"o2": "id:2,name:lili",
		"o3": "id:3,name:mimi",

		"k1": "age:20",
		"k2": "addr:new york",
		"k3": "email:1@e.com",
	}

	r := ParseIsetCom(&attr)
	fmt.Println(r["kvs"])
	fmt.Println(r["users"])
	fmt.Println(r["objs"])
}
