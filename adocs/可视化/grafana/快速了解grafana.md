# 一、Grafana 是什么?
Grafana 是一个“统一查询 + 可视化 + 告警” 的观测平台(Visualization & Observability UI)

它本身**不存数据**，只做三件事：
1. 向数据源发送查询
2. 把结果渲染成图表
3. 基于查询结果做告警

# 二、Grafana架构与核心组件
```text
Browser
  ↓
Grafana Server
  ↓
DataSource Plugin
  ↓
Prometheus / Elasticsearch / Loki / MySQL
```
1️⃣ Grafana Server
   - Web UI
   - Dashboard 管理
   - 变量解析
   - 告警调度
   - 用户与权限
⚠️：Grafana 本身是**无状态的**(状态在DB)
1️

2️⃣数据源(Data Source)
Grafana支持多种数据源,(💡：数据源插件 = 查询翻译器)

| 类型 | 常见用途 |
|--- | --- |
|Prometheus | 指标 |
|Elasticsearch | 日志 |
|Loki | 日志(轻量) | 
| MySQL/PostreSQL | 业务数据 |
| InfuxDB | 时序 |

三、Dashboard(仪表盘) 核心概念
Dashboard由什么组成?
```text
Dashboard
 ├── Rows
 │    └── Panels
 │         └── Query + Visualization
 └── Variables
```

Panel(面板)是核心:

一个Panel包含:
1. Query(PromQL /ES DSL / SQL)
2. Time Range
3. Transform
4. Visualization
5. Threshold / Unit / Legend

# 四、Grafana查询是如何工作的(关键)
**1️⃣查询流程(以Prometheus为例)**
```text
Panel
 ↓
PromQL
 ↓
Grafana 注入时间范围 & step
 ↓
HTTP 请求 Prometheus /api/v1/query_range
 ↓
返回时序数据
 ↓
渲染图表
```
💡：Grafana不是执行引擎，只是拼请求

**2️⃣Grafana自动加入了什么?**

时间范围: `[$_range]`
步长:`step = $_interval`

# 五、时间范围 & Interval(重点)
1️⃣Time Range
- Dashboard 级
- Panel 可 override

2️⃣Interval(自动采样)

Grafana会根据: `time range /max data points` 计算出`$_interval`


# 六、变量(Grafana的灵魂)
**1️⃣为什么要变量？**
- 多服务
- 多环境
- 多实例
- 一个Dashboard 复用

**2️⃣常见变量类型**
- Query: label_values
- Custom: dev,prod
- Constant: 固定值
- Datasource: 切换数据源

**3️⃣Prometheus变量示例**
```PromQL
label_values(http_requests_total, service)
```
在查询中使用:
```PromQL
sum(rate(http_requests_total{service="$service"}[5m]))
```
⚠️：`$_all`是一个坑点(正则)



# Grafana VS Kibana
| | Grafana | Kibana  |
|--|---|---------|
|数据源| 多 | ES only |
| 指标| 🌟🌟🌟🌟🌟 |    🌟🌟     |
|日志| 🌟🌟🌟  | 🌟🌟🌟🌟🌟        |
|告警| 强 | 强       |
|统一观测| 🌟🌟🌟🌟🌟| 🌟🌟|
