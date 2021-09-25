package internal

import (
	"log"
	"strings"
)

func InitScript() {
	boot, ok := Options.Script["_boot"]
	if !ok {
		return
	}

	for _, v := range strings.Split(boot.(string), " ") {
		name := strings.TrimSpace(v)
		if name == "" {
			continue
		}
		script, ok := Options.Script[name]
		if !ok || script == "" {
			log.Println("warning: fail to find script " + name)
			continue
		}
		go (func(script string) {
			if err := RunCmd(script); err != nil {
				log.Println("error:", err)
			}
		})(script.(string))
	}
}
