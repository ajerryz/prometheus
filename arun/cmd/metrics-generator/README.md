# 

# 一个暴露指标统计案例
## 一、整体流程
```text
Go 应用
  ↓
Prometheus Client（埋点）
  ↓
/metrics HTTP 接口
  ↓
Prometheus scrape
  ↓
PromQL / Grafana
```

## 二、引入 Prometheus Go Client
依赖
```shell
go get github.com/prometheus/client_golang/prometheus
go get github.com/prometheus/client_golang/prometheus/promhttp
```

## 三、最小可运行示例(建议先跑)
