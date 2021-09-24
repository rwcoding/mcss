package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rwcoding/mcss/config"
	_ "github.com/rwcoding/mcss/config"
	"github.com/rwcoding/mcss/internal"
	"log"
	"net/http"
	"strings"
)

func main() {
	app := gin.Default()
	app.NoRoute(func(context *gin.Context) {
		path := context.Request.URL.Path
		if path[len(path)-1:] == "/" {
			path += "index.html"
		}
		if !strings.Contains(path, ".html") {
			path += ".html"
		}
		file := config.Options.View + path

		html, err := internal.ParseFile(file, nil)

		if err != nil {
			log.Println(err)
			context.String(http.StatusNotFound, "")
		} else {
			context.Data(http.StatusOK, "text/html; charset=utf-8", html)
		}
	})
	if err := app.Run(config.Options.Addr); err != nil {
		log.Fatal(err)
	}
}