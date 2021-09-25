package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rwcoding/mcss/internal"
	"log"
	"net/http"
	"strings"
)

func main() {
	internal.InitCmd()
	internal.InitScript()

	app := gin.Default()
	if internal.Options.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	app.NoRoute(func(context *gin.Context) {
		path := context.Request.URL.Path
		if path[len(path)-1:] == "/" {
			path += "index.html"
		}
		if !strings.Contains(path, ".html") {
			if strings.Contains(path, ".") {
				internal.StaticHandler(context)
				return
			}
			path += ".html"
		}
		file := internal.Options.View + path

		html, err := internal.ParseFile(file, nil)

		if err != nil {
			log.Println(err)
			context.String(http.StatusNotFound, "")
		} else {
			context.Data(http.StatusOK, "text/html; charset=utf-8", html)
		}
	})
	if err := app.Run(internal.Options.Addr); err != nil {
		log.Fatal(err)
	}
}
