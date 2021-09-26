package internal

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CmdBuild() error {
	data := map[string]string{}
	viewPath := GetViewPath()
	_ = filepath.WalkDir(viewPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Println("warning:", err)
			return err
		}
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".html") && !IsComponent(d.Name()) {
			r, err := ParseFile(path, nil)
			if err != nil {
				log.Println("warning:", err)
			}
			k := path[len(viewPath):]
			k = strings.ReplaceAll(k, ".html", "")
			k = strings.ReplaceAll(k, string(os.PathSeparator), "")
			data[k] = string(r)
		}
		return err
	})
	s, err := json.Marshal(data)
	if err != nil {
		return err
	}

	file := GetTmpPath() + string(os.PathSeparator) + "mcss.html.json"
	return ioutil.WriteFile(file, s, 0777)
}
