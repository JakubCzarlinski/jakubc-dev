package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"project/gen/App"
	"project/src/assets"
	"project/src/flags"
	"project/src/pages/head"
	"time"

	"github.com/JakubCzarlinski/go-logging"
	"github.com/gin-gonic/gin"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	gin.SetMode(gin.ReleaseMode)

	var router *gin.Engine
	if flags.UseGinDefault {
		router = gin.Default()
	} else {
		router = gin.New()
	}
	router.UseH2C = true

	err := assets.HostStaticFiles(router)
	if err != nil {
		panic(err)
	}

	router.GET("/", HomePage)

	if flags.UseLiveReload {
		router.GET("/sse", func(ginContext *gin.Context) {
			ginContext.Header("Content-Type", "text/event-stream")
			ginContext.Header("Cache-Control", "no-cache")
			ginContext.Header("Connection", "keep-alive")
			ginContext.Writer.Flush()

			for {
				message := "data:" + "\n\n"
				message += "retry: 300\n\n"
				_, err := ginContext.Writer.Write([]byte(message))
				if err != nil {
					logging.Info(logging.Blue("Reloading..."))
					return
				}
				ginContext.Writer.Flush()
				time.Sleep(5 * time.Second)
			}
		})
	}

	if flags.UseHttps {
		logging.Info(logging.Green("Listening on https://localhost:443"))
		router.RunTLS(":443", "./server.crt", "./server.key")
	} else {
		logging.Info(logging.Green("Listening on http://localhost:80"))
		router.Run(":80")
	}
}

func HomePage(ginContext *gin.Context) {
	props := App.AppProps{}
	headContents := make(map[string]struct{})
	component := App.Home(&props, headContents)
	head.DefaultPageRender(component, headContents, ginContext.Writer, false)
}
