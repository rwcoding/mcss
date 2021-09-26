package internal

import (
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"
)

// 组件名称-路径映射缓存
var components sync.Map

func FindComponent(name string, fromFile string) (string, error) {
	v, ok := components.Load(name)
	if ok {
		file := v.(string)
		if FileExists(file) {
			return file, nil
		} else {
			components.Delete(name)
		}
	}

	var pathList []string
	for _, v := range Options.Component {
		pathList = append(pathList, Options.Root+string(os.PathSeparator)+v)
	}
	root, _ := filepath.Abs(filepath.Dir(fromFile))
	pathList = append(pathList, root)

	file := ""
	if len(Options.Component) > 0 {
		for _, v := range pathList {
			_ = filepath.WalkDir(v, func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					log.Println("warning:", err)
					return err
				}
				if !d.IsDir() {
					if name+".html" == d.Name() {
						file = path
						components.Store(name, path)
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
