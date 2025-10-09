# prometheus yaml配置文件
- Prometheus 通过命令行参数和配置文件进行配置。
- 命令行参数配置的是不可改变的系统参数（例如存储位置、磁盘和内存中保存的数据量等），而配置文件则定义了与抓取作业及其实例相关的所有内容，以及需要加载的规则文件。
- 要查看所有可用的命令行标志，请运行`./prometheus -h`。


Prometheus 可以在`运行时重新加载其配置`。如果新配置格式不正确，则更改将无法应用。
- 通过SIGHUP向 Prometheus 进程发送 
- 或向端点发送 HTTP POST 请求/-/reload（--web.enable-lifecycle启用该标志时）可以触发配置重新加载。这也会重新加载所有已配置的规则文件。



# 配置文件
要指定要加载哪个配置文件，请使用`--config.file`标志。


## 配置
全局配置指定在所有其他配置上下文中有效的参数。它们也可作为其他配置部分的默认值。
```yaml
global:
  scrape_interval: 10s # 默认情况下抓取目标的频率。default=1m
  scrape_timeout: 10s  # 抓去目标的请求超时时间，该值不能大于抓取间隔。default=10s
  #......
```
### `scrape_config`


## 记录规则

## 告警规则

