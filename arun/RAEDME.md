# arun

# 目录介绍
- `prometheus`: 跑源码构建的`prometheus server`
- `prometool`: 跑源码构建的`prometool` 可能会没有需要构建
- `cmd/`: 自定义的一些app,用于用程序实现指标暴露
- `docs`: 文档目录
- `grafana`: grafana目录


# prometheus
- prometheus: [官方文档](https://prometheus.io/docs/prometheus/latest/getting_started/)

# grafana
- grafana oss: https://grafana.com/grafana/download?edition=oss&pg=oss-graf&platform=linux&plcmt=hero-btn-1

```shell
# 标准Linux
wget https://dl.grafana.com/grafana/release/12.2.0/grafana_12.2.0_17949786146_linux_amd64.tar.gz

# Macos
curl -O https://dl.grafana.com/grafana/release/12.2.0/grafana_12.2.0_17949786146_darwin_amd64.tar.gz
```