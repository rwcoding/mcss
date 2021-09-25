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
	_ = filepath.WalkDir(Options.View, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Println("warning:", err)
			return err
		}
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".html") && !strings.Contains(d.Name(), "-") {
			r, err := ParseFile(path, nil)
			if err != nil {
				log.Println("warning:", err)
			}
			k := path[len(Options.View):]
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

	file := Options.Root + string(os.PathSeparator) + Options.TmpPath + string(os.PathSeparator) + "mcss.html.json"
	return ioutil.WriteFile(file, s, 0777)
}
