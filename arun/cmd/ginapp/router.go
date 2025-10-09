package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var engine *gin.Engine

func boot() {
	engine = gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())
	engine.Use(Prometheus()) // 将自定义指标采集的 middleware 进行注册

	// 集成 prometheus
	engine.GET(metricsPath, gin.WrapH(promhttp.Handler()))

	// router
	{
		helloController := NewHelloController()
		group := engine.Group("/hello")
		group.GET("", helloController.Hello)
		group.GET("/:name", helloController.HelloPath)
		group.GET("/400", helloController.Error400)
		group.GET("/500", helloController.Error500)
	}

	host := flag.String("host", "localhost", "The hostname of the server")
	port := flag.Int("port", 8080, "The port number of the server")

	flag.Parse()

	hostAndPort := fmt.Sprintf("%s:%d", *host, *port)

	fmt.Printf("metrics:http://%v/metrics\n", hostAndPort)
	if err := engine.Run(hostAndPort); err != nil {
		panic(err)
	}
}
