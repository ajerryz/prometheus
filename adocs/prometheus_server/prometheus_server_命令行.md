# Prometheus 命令行

# flags
prometheus 的 `flags`:
```text
flag                                    描述                                  默认

-h,--help                               显示帮助信息
--version                               显示版本
--config.file                           prometheus配置文件路径                    prometheus.yml
--config.auto-reload-interval           配置文件检测修改时间间隔                    30s
--web.listen-address ...                监听UI,API和遥测数据的地址。可以重复         0.0.0.0:9090
--auto-gomaxprocs                       自动设置GOMAXPROCS以匹配Linux容器CPU配额   true
--auto-gomemlimit                       自动设置GOMEMLIMIT以匹配Linux容器或系统内存限制 true
--auto-gomemlimit.ratio                 预留的GOMEMLIMIT内存与检测到最大容器内存的比例   0.9
--web.config.file                       
--web.read-timeout                      读取请求超时并关闭空闲连接之前的最大持续时间   5m
--web.max-connections                   所有监听器的同时最大连接数                  512
--web.max-notifications-subscribers     限制同时接收实时的订阅者数量上限。           16
                                        如果达到上限，新的订阅请求将被拒绝，直到
                                        现有连接关闭。
--web.external-rul                      Prometheus 外部可访问的 URL（例如，如果 
                                        Prometheus 通过反向代理提供服务）。用于生
                                        成指向 Prometheus 本身的相对和绝对链接。
                                        如果 URL 包含路径部分，它将用作 Prometheus 服务的
                                        所有 HTTP 端点的前缀。如果省略，则会自动生成相关
                                        的 URL 组件。
--web.route-prefix                      Web 端点内部路由的前缀。
                                        默认为 --web.external-url 的路径。
--web.user-assets                       静态资产目录的路径，可在 /user 处获取。
--web.enable-lifecycle                  通过 HTTP 请求启用关机和重新加载。           false
--web.enable-admin-api                  为管理员控制操作启用API端点                 false
--web.enable-remote-write-receiver      启用API端点接受远程写入请求                 false
......
```
[官方详细文件](https://prometheus.io/docs/prometheus/latest/command-line/prometheus/)