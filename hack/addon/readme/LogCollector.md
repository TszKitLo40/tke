# LogCollector说明

## 组件介绍

提供的集群内日志采集功能，支持将集群内服务或集群节点特定路径文件的日志发送至 Kafka、Elasticsearch等消费端，支持采集容器标准输出日志，容器内文件日志以及主机内文件日志。更提供事件持久化、审计等功能，实时记录集群事件及操作日志记录，帮助运维人员存储和分析集群内部资源生命周期、资源调度、异常告警等情况。

日志收集功能需要为每个集群手动开启。日志收集功能开启后，日志收集 Agent 会在集群内以 Daemonset 的形式运行。用户可以通过日志收集规则配置日志的采集源和消费端，日志收集 Agent 会从用户配置的采集源进行日志收集，并将日志内容发送至用户指定的消费端。

需要注意的是，使用日志收集功能需要您确认 Kubernetes 集群内节点能够访问日志消费端。

### LogCollector使用场景

日志收集功能适用于需要对 Kubernetes 集群内服务日志进行存储和分析的用户。用户可以通过配置日志收集规则进行集群内日志的收集并将收集到的日志发送至 Kafka 的指定 Topic 或 日志服务 CLS 的指定日志主题以供用户的其它基础设施进行消费。

### 部署在集群内kubernetes对象

在集群内部署LogCollector Add-on , 将在集群内部署以下kubernetes对象

| kubernetes对象名称 | 类型 | 默认占用资源 | 所属Namespaces |
| ----------------- | --- | ---------- | ------------- |
| log-collector |DaemonSet |每节点0.3核CPU, 250MB内存|kube-system|