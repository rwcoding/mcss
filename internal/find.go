package internal

import (
	"errors"
	"github.com/rwcoding/mcss/config"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var components sync.Map

func FindComponent(name string, fromFile string) (string, error) {
	v, ok := components.Load(name)
	if ok {
		file := v.(string)
		if FileExists(file) {
			return file, nil
		}
	}

	var pathList []string
	for _, v := range config.Options.Component {
		pathList = append(pathList, config.Options.Root+string(os.PathSeparator)+v)
	}
	root, _ := filepath.Abs(filepath.Dir(fromFile))
	pathList = append(pathList, root)

	file := ""
	if len(config.Options.Component) > 0 {
		for _, v := range pathList {
			_ = filepath.WalkDir(v, func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					log.Println("warning:", err)
					return err
				}
				if !d.IsDir() {
					if name+".html" == d.Name() {
						file = path
						return err
					}
				}
				return err
			})
			if file != "" {
				return file, nil
			}
		}
	}
	return "", errors.New("can not find component " + name)
}

func FileExists(file string) bool {
	if _, err := os.Stat(file); err != nil {
		return false
	}
	return true
}
