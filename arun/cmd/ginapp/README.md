# 下载依赖
```shell
go get github.com/gin-gonic/gin

go get github.com/prometheus/client_golang/prometheus
go get github.com/prometheus/client_golang/prometheus/promhttp
go get github.com/gin-contrib/prometheus
```

# 实现方案
通过 gin 中的 middleware 方式 实现通用请求的 常用指标.详细见`metrics.go`


# 实验
1. 启动ginapp
```shell
go run .
```
2. 使用脚本发送请求
```shell
cd tools
python -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
python main.py
```