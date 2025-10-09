# 服务发现
Prometheus 本身的配置（prometheus.yml）是静态的，修改后必须重启或发送 SIGHUP 信号才能生效。为了实现目标（Target）的动态配置，我们需要利用 Prometheus 提供的服务发现（Service Discovery）机制。


## 核心思想
动态配置的核心思想是：不要在 prometheus.yml 中硬编码目标的 IP 和端口，而是让 Prometheus 定期从一个外部的 “数据源”（如 Kubernetes API、Consul、DNS 等）去发现这些目标。


## 服务发现的方案
- 基于文件的动态发现
- 基于DNS的动态发现
- 基于Kubernetes的服务发现
- 使用第三方注册中心
  - consul
  - etcd
  - ZooKeeper


### 基于文件的动态发现
工作原理:
1. 配置文件：你在 prometheus.yml 中配置一个 file_sd_configs 块，指定 Prometheus 要监控的文件或目录。
2. 目标文件：这些被监控的文件（通常是 .yml 或 .json 文件）包含了目标列表及其标签。
3. 监控与重载：Prometheus 会定期（默认每 5 分钟）检查这些文件的修改时间（mtime）。如果文件有更新，Prometheus 会自动加载新的目标配置，无需重启。这个过程是无缝的，不会中断正在进行的监控。

实验：
1. 编写`file_sd_configs`模块，例如:[prometheus_dynamicfile.yml](../../prometheus_dynamicfile.yml)。看其中的`file_sd_job`配置。
2. 编写targets,例如:[ginapp_targets.json](../../targets/ginapp_targets.json),在该文件中可动态编写抓取目标
3. 启动prometheus,`./prometheus --config.file=prometheus_dynamaicfile.yml`
4. 启动`ginapp`,`go run . --port 8080`
5. 启动`ginapp`,`go run . --port 9999`
6. 再次修改`ginapp_targets.json`动态将端口为9999的配置加入进来。



### 基于三方注册中心 etcd
工作原理:
1. 服务注册：你的应用程序（或其编排工具，如 Nomad）在启动时，会将自己的信息（如 IP 地址、端口、服务名、环境标签等）以 JSON 格式存入 etcd 的一个特定路径下，例如 /services/my-app/192.168.1.10:8080。
2. Prometheus 配置：你在 prometheus.yml 中配置一个 etcd_sd_configs 块，告诉 Prometheus 去哪个 etcd 集群、以及在哪个键前缀下查找目标。
3. 发现与监控：Prometheus 会连接到 etcd 集群，并持续监听（watch）指定前缀下的键变化。
   - 当一个新服务注册时（新增一个键），Prometheus 会立即收到通知，并自动将这个新目标添加到其监控列表中。
   - 当一个服务下线时（删除一个键），Prometheus 同样会收到通知，并停止对该目标的抓取。
   - 当一个服务信息更新时（修改一个键的值），Prometheus 会更新对应的目标信息。

这个过程是完全自动化和实时的，无需人工干预或重启 Prometheus。


？？高版本貌似移除了对 etcd 的支持


### 基于三方注册中心 consul
//TODO


