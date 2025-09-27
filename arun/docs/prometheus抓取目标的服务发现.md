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
// TODO


### 基于三方注册中心 consul

### 基于三方注册中心 etcd

