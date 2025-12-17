package main

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// 请求数量 Counter
	httpRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "ns",
		Subsystem: "ginapp",
		Name:      "http_requests_total",
		Help:      "total number of http requests",
	}, []string{"method", "url", "status"})
	// 请求耗时 Histogram
	httpDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "ns",
			Subsystem: "ginapp",
			Name:      "http_request_duration_seconds",
			Help:      "HTTP request duration",
			Buckets:   prometheus.DefBuckets,
		}, []string{"method", "url", "status"})
	// 当前(总)并发请求
	inflightRequests = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "ns",
		Subsystem: "ginapp",
		Name:      "http_inflight_requests",
		Help:      "Number of inflight requests",
	})
)

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		inflightRequests.Inc() // 总并发数+1
		defer inflightRequests.Dec()

		c.Next()

		latency := time.Since(start).Seconds()
		status := c.Writer.Status()
		path := c.FullPath()
		if path == "" {
			path = "unknown"
		}
		httpRequestsTotal.WithLabelValues(
			c.Request.Method,
			path,
			strconv.Itoa(status),
		).Inc()

		httpDuration.WithLabelValues(
			c.Request.Method,
			path,
			strconv.Itoa(status),
		).Observe(latency)
	}
}
