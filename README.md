# XTimer

- 定时任务是一个与业务相关性不大，且常见的任务场景，可以抽离出的通用服务，本项目使用kratos框架实现一个简单的定时任务服务。

- 实现思路参考了[小徐先生](https://github.com/xiaoxuxiansheng/xtimer)，感谢他在go相关技术上的分享。

## Hello Xtimer

XTimer主要分为Migrator、Scheduler、Trigger、Executor模块。整体思路如下：

- XTimer提供了接口去创建、激活、删除、查询定时器，后续由Migrator模块根据定时器的Cron表达式会批量生成具体的任务
- Scheduler模块，按秒轮询，分配协程调度Trigger模块跟踪处理每个分片中任务
- Trigger模块，按秒轮询，拉取到期的任务，分配协程调度Executor模块执行任务
- Executor模块，执行具体任务，

## Todo

- 接入prometheus监控任务执行情况
- 支持根据任务数量动态分桶
- 结合本地缓存，实现三级存储结构，进一步优化性能
