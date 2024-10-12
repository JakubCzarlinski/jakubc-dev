package main

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof"
	"project/gen/App"
	"project/src/flags"
	"project/src/logging"

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

	router.GET("/", HomePage)

	if flags.UseHttps {
		logging.Log(logging.Green("Listening on https://localhost:3000"))
		router.RunTLS(":3000", "./server.crt", "./server.key")
	} else {
		logging.Log(logging.Green("Listening on http://localhost:3000"))
		router.Run(":3000")
	}
}

func HomePage(ginContext *gin.Context) {
	props := App.AppProps{}
	headContents := make(map[string]struct{})
	component := App.Home(&props, headContents)
	component.Render(context.Background(), ginContext.Writer)
}
