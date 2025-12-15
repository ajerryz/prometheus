# Prometheus Rules
`prometheus`支持两种类型的规则，可以配置并定期执行：`Record Rule`,`Alter Rule`。
要在`prometheus`中添加规则，需要创建一个规则语句的文件，并让`prometheus`通过配置`rule_files`中的字段引用这些规则文件。

规则文件可以在运行时通过发送`sighup`到`prometheus`进程进行重新加载。只有规则文件格式正确时，更改才会生效。

# Recording Rules

## 语法规则检查
要在不启用`prometheus server`的情况下快速检查规则文件语法是否正确，可以使用Prometheus自带的`promtool`命令行工具：`promtool check rules /path/to/example.rules.yaml`

## Recording Rules
- 记录规则(Recording Rules)允许您预先计算常用或计算量大的表达式，并将其结果保存为一组新的时间序列。
- 查询预先计算的结果通常比每次需要时执行原始表达式要快得多。这对于每次刷新时都需要重复查询相同表达式的仪表板尤其有用。
- 记录规则和警报规则存在于规则组中。组内的规则按固定间隔顺序运行，且评估时间相同。记录规则的名称必须是 有效的指标名称。警报规则的名称必须是 有效的标签值。


规则文件的语法是:
```yaml
groups:
  - name: group-example-1   # 规则组名称，文件中必须唯一
    interval: 10s           # 计算组中record的频率,默认global.evaluation_interval
    limit: 0                # 限制警报规则和记录规则可以产生的警报系列的数量。0 表示没有限制。
    query_offset: 0
    labels:
      name: demo
    rules:
      - record: new_metric  # 输出到新时间序列的指标名称，必须是有效的指标名称
        expr: ''            # PromQL表达式，每个评估周期都会在当前时间点进行评估，并将其记录为一组新的时间序列，指标名称为record字段指定
        labels:
          label1: labelValue1
          label2: labelValue2
      - alert: ''           # 告警的名称，必须是有效的标签值
        expr: ''            # 待评估的PromQL表达式，每个评估周期都会在当前时间点进行评估
        for: 0s   # 
        keep_firing_for: 0s
        labels:
          label1: labelValue1
        annotations:
          labelName: tmpl_string
```


# Alerting Rules
