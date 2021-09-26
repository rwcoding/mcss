package internal

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	InitConsole(os.Args)
	parse()
}

var Options struct {
	Debug     bool                   `yaml:"debug"`
	Addr      string                 `yaml:"addr"`
	View      string                 `yaml:"view"`
	Component []string               `yaml:"component"`
	Mcss      map[string]interface{} `yaml:"mcss"`
	Script    map[string]interface{} `yaml:"script"`
	Template  map[string]interface{} `yaml:"template"`
	VoidTag   []string               `yaml:"void_tag"`
	Root      string
	TmpPath   string `yaml:"tmp_path"`
}

func parse() {
	configFile := ""

	if conf := GetFlag("conf"); conf != "" {
		configFile = conf
	}

	var content []byte
	var f http.File
	var err error
	var dir string

	if configFile != "" {
		f, err = os.Open(configFile)
		dir, _ = filepath.Abs(filepath.Dir(configFile))
	} else {
		if dir, err = os.Getwd(); err == nil {
			if f, err = os.Open(dir + string(os.PathSeparator) + "mcss.local.yaml"); err != nil {
				if f, err = os.Open(dir + string(os.PathSeparator) + "mcss.yaml"); err != nil {
					if dir, err = filepath.Abs(filepath.Dir(os.Args[0])); err == nil {
						f, err = os.Open(dir + string(os.PathSeparator) + "mcss.yaml")
					}
				}
			}
		}
	}
	if err != nil {
		log.Println("warning:failed to find config file, use sample config")
		content = []byte(sample)
	} else {
		defer f.Close()
		content, err = ioutil.ReadAll(f)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = yaml.Unmarshal(content, &Options)
	if err != nil {
		log.Fatal("failed to parse config file, ", err.Error())
	}

	Options.Root = dir
	Options.View = strings.ReplaceAll(dir+"/"+Options.View, "/", string(os.PathSeparator))
	if Options.TmpPath == "" {
		Options.TmpPath = "tmp"
	}

	if s, ok := Options.Template["if_start"]; ok && s != "" {
		templateIfStart = s.(string)
	}

	if s, ok := Options.Template["if_end"]; ok && s != "" {
		templateIfEnd = s.(string)
	}
}
