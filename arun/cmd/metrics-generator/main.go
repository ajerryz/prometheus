package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	quickMetricsDemo()
}

func quickMetricsDemo() {
	// 一个简单的Counter

	// 1.创建一个计数器类型指标
	httpRequestTotal := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "ns", // metrics 的规范: namespace_subsystem_name
		Subsystem: "sys",
		Name:      "http_requests_total",
		Help:      "The total number of HTTP requests made.",
		ConstLabels: prometheus.Labels{
			"method": "GET",
			"status": "200",
		},
	})

	// 注册指标
	prometheus.MustRegister(httpRequestTotal)

	// 业务接口
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		httpRequestTotal.Inc()
		_, _ = fmt.Fprintf(w, "Hello, ok")
	})

	// (可选),mock 请求,直接操作客户端埋点
	go func() {
		for {
			httpRequestTotal.Add(1000)
			time.Sleep(time.Duration(5) * time.Second)
		}
	}()

	// 暴露 /metrics
	http.Handle("/metrics", promhttp.Handler())

	// 启动业务本身服务

	if err := http.ListenAndServe(":8788", nil); err != nil {
		panic(err)
	}
}
