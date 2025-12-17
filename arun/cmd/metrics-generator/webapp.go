package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func bootWebApp() {
	r := gin.New()
	r.Use(gin.Recovery())
	//r.Use(gin.Logger())

	// 注册
	prometheus.MustRegister(
		httpRequestsTotal,
		httpDuration,
		inflightRequests,
	)

	r.Use(PrometheusMiddleware())
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	RouteRegister(r)

	if err := r.Run(":8788"); err != nil {
		panic(err)
	}
}
