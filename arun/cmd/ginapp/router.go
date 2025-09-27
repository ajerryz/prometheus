package main

import (
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

	fmt.Printf("metrics:http://localhost:8080/metrics\n")
	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}
}
