package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rwcoding/mcss/internal"
	"log"
	"net/http"
	"strings"
)

func main() {
	internal.InitCmd()
	internal.InitScript()

	if internal.Options.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.Default()
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

		var html []byte
		var err error
		for _, v := range internal.GetViewPath() {
			file := v + path
			if v[:1] == "@" {
				file = v[1:] + path
			}
			html, err = internal.ParseFile(file, nil)
			if html != nil {
				break
			}
		}

		//file := internal.GetViewPath() + path
		//html, err := internal.ParseFile(file, nil)

		if err != nil {
			log.Println(err)
			context.String(http.StatusNotFound, "")
		} else {
			context.Data(http.StatusOK, "text/html; charset=utf-8", html)
		}
	})

	fmt.Println("running ······ ")
	if err := app.Run(internal.Options.Addr); err != nil {
		log.Fatal(err)
	}
}
