package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"project/gen/App"
	"project/src/assets"
	"project/src/flags"
	"project/src/pages/head"

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

	if flags.UseHttps {
		logging.Info(logging.Green("Listening on https://localhost:3000"))
		router.RunTLS(":3000", "./server.crt", "./server.key")
	} else {
		logging.Info(logging.Green("Listening on http://localhost:3000"))
		router.Run(":3000")
	}
}

func HomePage(ginContext *gin.Context) {
	props := App.AppProps{}
	headContents := make(map[string]struct{})
	component := App.Home(&props, headContents)
	head.DefaultPageRender(component, headContents, ginContext.Writer, false)
}
