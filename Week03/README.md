学习笔记

### 原则
#### 1. 把是否使用gorouting权限交给调用者显示调用
#### 2. 管控住gorouting生命周期

### sync
#### sync.Once 多次调用只执行一次


### errgroup
为什么不在协程里拼装数据, 原因是可能会引起并发冲突


计算密集型，context超时处理基本不处理, 很难打断
 

---- 
网络请求超时控制怎么处理?

动态计算deadLine时间, 如果小于配置的时间, 超时时间就是dealLine, 否则为配置的时间


注意context生命周期



event(data<-, done) -> back goroutine
close back gorutine   -> stop event
得知后台goroutine什么时候退出  前台event什么时候退出


