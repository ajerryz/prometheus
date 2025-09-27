package main

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

const metricsPath = "/metrics"

var (
	// 请求总数，Counter 类型 累加
	httpRequestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gin_http_request_total",
			Help: "Total number of HTTP requests received",
		},
		[]string{"method", "uri", "status_code"})

	//请求延迟 Histogram 类型，
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "gin_http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: []float64{0.01, 0.05, 0.5, 2, 5, 10},
		},
		[]string{"method", "uri"})

	// 响应大小 Summary 类型, 适合直接统计 平均值、分位值(客户端计算)
	httpResponseSize = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "gin_http_response_size_bytes",
			Help: "Size of HTTP requests",
		},
		[]string{"method", "uri"},
	)
)

// 注册
func init() {
	prometheus.MustRegister(
		httpRequestTotal,
		httpRequestDuration,
		httpResponseSize,
	)
}

func Prometheus() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 忽略 /metrics
		if c.Request.URL.Path == metricsPath {
			c.Next()
			return
		}

		// 1.开始记录请求时间
		startTime := time.Now()
		c.Next()

		// 2.请求完成 更新指标
		duration := time.Since(startTime).Seconds() // 计算延迟
		statusCode := c.Writer.Status()             // 获取响应状态码
		method := c.Request.Method                  // 获取请求方法
		uri := c.FullPath()                         // 获取请求Path
		responseSize := float64(c.Writer.Size())    // 获取响应大小，字节
		// 3.更新指标
		httpRequestTotal.WithLabelValues(method, uri, strconv.Itoa(statusCode)).Inc()
		httpRequestDuration.WithLabelValues(method, uri).Observe(duration)
		httpResponseSize.WithLabelValues(method, uri).Observe(responseSize)
	}
}
