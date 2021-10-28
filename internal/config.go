package internal

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func init() {
	InitConsole(os.Args)
	parse()
}

var Options struct {
	Debug     bool                   `toml:"debug"`
	Addr      string                 `toml:"addr"`
	View      []string               `toml:"view"`
	Component []string               `toml:"component"`
	Mcss      map[string]interface{} `toml:"mcss"`
	Script    map[string]interface{} `toml:"script"`
	Iset      map[string]interface{} `toml:"iset"`
	VoidTag   []string               `toml:"void_tag"`
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
			if f, err = os.Open(dir + string(os.PathSeparator) + "mcss.local.toml"); err != nil {
				if f, err = os.Open(dir + string(os.PathSeparator) + "mcss.toml"); err != nil {
					if dir, err = filepath.Abs(filepath.Dir(os.Args[0])); err == nil {
						f, err = os.Open(dir + string(os.PathSeparator) + "mcss.toml")
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

	_, err = toml.Decode(string(content), &Options)
	if err != nil {
		log.Fatal("failed to parse config file, ", err.Error())
	}

	Options.Root = dir
}

func GetViewPath() []string {
	ret := []string{}
	for _, v := range Options.View {
		if v[:1] == "@" {
			ret = append(ret, "@"+FormatPath(Options.Root+"/"+v[1:]))
		} else {
			ret = append(ret, FormatPath(Options.Root+"/"+v))
		}
	}
	return ret
}

func GetTmpPath() string {
	if Options.TmpPath == "" {
		return FormatPath(Options.Root + "/tmp")
	} else {
		return FormatPath(Options.Root + "/" + Options.TmpPath)
	}
}
