# Prometheus TSDB
Prometheus = 采集(抓取) + 存储 + 查询

其中TSDB负责:
- 存储 时序数据
- 高效追加写
- 高效范围查
- 不支持 update/delete (只能append / compaction)

核心特性:
高基数时间序列、写多读多、强时间维度。

# TSDB 与 传统数据库的差异
| 维度 | TSDB | MySQL/OLTP |
|---  |--- | ---|
| 写入 |Append-only | insert/update |
|主键 | (metric + labels + time) | 任意 |
|索引 | Label 倒排索引 | B+Tree|
|查询 | 时间范围扫描 | 精确匹配 |
|压缩 | 极致 | 一般 |
|删除 | 按时间批量 | 任意

TSDB：写时顺序、读时范围 的数据库

# Prometheus数据模型(TSDB基础)
1.时间序列(Series)
```text
Series = 唯一的标签集合
```
例如:`cpu_usage{host="node1", core="0"}

2.样本(Sample)
```text
(timestamp,value)
```
示例:`(169000000,0.73)`

3.数据块(Chunk)

TSDB不会一个 sample 一个存,而是:
```text
Series
 ├── Chunk1 (120 samples)
 ├── Chunk2
 ├── Chunk3
```
👉:Chunk是压缩、IO、查询的最小单位

# TSDB典型存储分层(通用架构)
```markdown
        ┌────────────┐
        │   Memory   │  ← 最新数据
        └─────┬──────┘
              │
        ┌─────▼──────┐
        │    WAL     │  ← 防丢数据
        └─────┬──────┘
              │
        ┌─────▼──────┐
        │   Blocks   │  ← 持久化
        └─────┬──────┘
              │
        ┌─────▼──────┐
        │ Compaction │
        └────────────┘
```

## WAL(Write Ahead Log)
WAL本质:`先顺序写日志，再更新内存`
- 优点：
  - 顺序IO,速度快
  - 崩溃可恢复
  - 不依赖随机写
- 缺点：
  - 占磁盘
  - 重启需要replay

TSDB的WAL写什么?
写入的不是SQL,而是:
```text
SeriesRef -> timestamp -> value
```
并包含:
- 新series创建
- label信息
- sample数据

## 内存结构(Head Block)
内存中存:
1. Series 元数据
2. Chunk (正在写)
3. 倒排索引(label -> series)

特点:
- 写入全部发生在内存
- 查询优先命中内存
- 定期flush成磁盘block

## Block(持久化文件)
Block的意义:
- 时间范围固定(如:2h)
- 只读
- 查询友好
- 可压缩、可合并

Block的典型组成:
```text
block/
├── meta.json     # 时间范围、统计信息
├── index         # label → series
├── chunks/       # 真正的时间序列数据
```

为什么TSDB必须block化？
- 文件不可变 -> 易cache
- 可 mmap
- 合并简单
- 删除简单(按 block)

## 索引设计(TSDB的灵魂)
1.为什么不用B+Tree?
- 查询不是`WHERE id=xxx`
- 而是:
```text
WHERE labelA="A" AND labelB!="y" AND time BETWEEN t1 AND t2
```

2.倒排索引(Inverted Index)
```text
label_name
 └── label_value
     └── [series_id1, series_id2, ...]
```
查询流程：
1. 找label对应的 series 集合
2. 多条件做集合运算(AND/OR)
3. 再按时间扫 chunk

3.高基数为什么是TSDB的杀手？
- 倒排索引爆炸
- 内存、磁盘、CPU全炸


## 压缩算法(TSDB为什么省空间)
1. 时间戳压缩(Delta-of-Delta)
2. 数值压缩(XOR)

效果:
- 原始16 bytes/sample
- 压缩后 ~1~2 bytes/sample

## Compaction(合并压缩)


# TSDB的常见实例对比
|产品 | 特点|
|--- | --- |
|Prometheus TSDB | 单机、简单 |
| VictoriaMerics | 高压缩、高性能 |
| Cortex/Mimir | 分布式、对象存储 | 
| InfluxDB | SQL 风格 |
| OpenTSDB | 基于 HBase |