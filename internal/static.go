package internal

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
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

	file := Options.Root + string(os.PathSeparator) + path
	b, err := ioutil.ReadFile(file)
	if err != nil {
		context.Data(http.StatusOK, ct, []byte{})
		return
	}

	context.Data(http.StatusOK, ct, b)
}
