# Dashboard
# 一、Dashboard是什么
`Dashboard = 多个Panel + 统一时间范围 + 变量 + 权限`

Dashboard解决的问题:
- 快速定位问题
- 建立系统健康模型
- 支撑告警和容量评估


Dashboard层级结构:
```text
Organization
 └── Folder
      └── Dashboard
           ├── Variables
           ├── Rows
           │    └── Panels
           └── Time Range
```
- Folder = 权限边界
- Dashboard = 业务视角
- Panel = 技术视角


# 二、Dashboard核心组成
**1️⃣Time Range (时间范围)**
- Dashboard 级默认
- Panel 可 override

💡：时间范围影响 PromQL 的 step 和 采样密度

**2️⃣ Variables (Dashboard灵魂)**
常见变量设计:
- env : 环境
- service : 服务名
- instance : 实例
- datasource : 数据源

变量示例:
```text
label_values(http_requests_total{env="$env"},service)

使用:
sum(rate(http_requests_total{service="$service"}[5m]))
```

**3️⃣ Panels (真正的主角)**

一个 Panel = Query + Visualization + Options

## 2.1、Panel 详解
1️⃣：查询

Prometheus Panel
```text
Grafana 会发送: /api/v1/query_range
此时自动注入:
- start/end
- step = $_interval

示例:
sum(rate(http_requests_total[5m])) by (service)
```
Elasticsearch Panel
```text
Grafana实际构造: POST index/_search
此时自动转换:
- 自动加时间 filter
- UI查询 -> DSL
```

2️⃣ Visualization(可视化)

常用类型与用途:
- Time series: 趋势
- Stat: 单值
- Gauge: 阈值
- Table: 明细
- Logs: 日志

3️⃣Filed/Unit/Leggen

- Unit: ms / % / req/s
- Legen: {{service}}
- Threshold: 颜色分段



## 2,2 Panel 高级能力
1️⃣ Transform(强但容易被忽略)

常见 Transform:
- Reduce : 多点 -> 单值
- Group By : 聚合
- Join : 多查询合并
- Add field : 派生字段

💡：有些逻辑放 Grafana 层比 PromQL 更好

2️⃣Overrides(字段级控制)

- 单位
- 小数位
- 颜色
- 阈值

# 三、Dashboard设计方法论(生产经验)
1️⃣自上而下设计：
- 整体健康
- 关键指标
- 异常定位
- 明细日志


2️⃣ 黄金四指标(必备)
以Prometheus PromQL为例
- QPS : rate
- 错误率: error / total
- 延迟: p95 / p99
- 饱和度: CPU/内存

3️⃣每个Dashboard 建议结构
```text
Row 1：全局状态（Stat）
Row 2：流量 & 错误
Row 3：延迟
Row 4：资源
Row 5：日志入口
```