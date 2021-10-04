package internal

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var contentType map[string]string = map[string]string{
	"html":  "text/html; charset=utf-8",
	"css":   "text/css",
	"js":    "application/x-javascript",
	"png":   "image/png",
	"jpg":   "image/jpeg",
	"gif":   "image/gif",
	"ico":   "image/x-icon",
	"svg":   "image/svg+xml",
	"bmp":   "image/bmp",
	"webp":  "image/webp",
	"jpeg":  "image/jpeg",
	"eot":   "application/octet-stream",
	"ttf":   "application/octet-stream",
	"woff":  "application/octet-stream",
	"woff2": "application/octet-stream",
}

func StaticHandler(context *gin.Context) {
	path := context.Request.URL.Path

	ext := path[strings.LastIndex(path, ".")+1:]
	ct, ok := contentType[ext]
	if !ok {
		ct = "text/html; charset=utf-8"
	}

	excluded := strings.Split(context.Query("excluded"), ",")
	for _, t := range []string{"css", "js", "less"} {
		if strings.Contains(path, "*."+t) {
			dir := Options.Root + string(os.PathSeparator) + strings.ReplaceAll(path, "*."+t, "")
			context.Data(http.StatusOK, ct, scan(dir, t, excluded))
			return
		}
	}

	file := Options.Root + string(os.PathSeparator) + path
	b, err := ioutil.ReadFile(file)
	if err != nil {
		context.Data(http.StatusNotFound, ct, []byte{})
		return
	}

	context.Data(http.StatusOK, ct, b)
}

func scan(dir, fileType string, excluded []string) []byte {
	var content bytes.Buffer
	_ = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Println("warning:", err)
			return err
		}
		if !d.IsDir() {
			name := d.Name()
			if name[len(name)-len(fileType):] != fileType {
				return err
			}
			for _, v := range excluded {
				if v == name {
					return err
				}
			}
			c, err := ioutil.ReadFile(path)
			if err != nil {
				log.Println("warning:", err)
				return err
			}
			if fileType == "js" {
				content.WriteString("\n; // " + path + " \n")
			}
			content.Write(c)
		}
		return err
	})
	return content.Bytes()
}
