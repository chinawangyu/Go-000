### 学习笔记
#### 1. 服务
服务关注单一业务, 

#### 2. 为什么服务间使用protobuf、thrift,而不使用json
json过于灵活, 只能靠接口文档约束对方传值, 容易出现问题

而protobuf定义结构固定, 能很好的约束双方

#### 3. 注册中心接入后为什么还需要健康检查

注册与服务 ---> 服务发现(服务与注册中心通讯)

服务与服务 ---> 健康检查(服务与服务通讯)

#### 4. 健康检查作用
平滑发布
兜底

#### 5. 客户端发现/服务端发现
服务 -> 服务

服务 -> LB(负载均衡) -> 服务



#### 7. 多租户

![avatar](多租户图.jpg)

上图中虚线右半部分为基础服务(稳定版本), 右侧为我们的dev新开发分支. 
请求入口为服务A, header中加入 dev1 标签, 在判断当中如果包含 dev1 标签检测是否有在运行服务,
如果存在则转入左侧请求, 如果为发现则为右侧基础版本。 dev1 标签贯穿整个请求流程生命周期.
 


适用场景: 多测评环境、灰度发布、压测


#### 8. 野协程panic问题代码
./code/panic.go


### 图解回收垃圾
https://zhuanlan.zhihu.com/p/297177002?utm_source=wechat_session&utm_medium=social&utm_oi=26711194337280&utm_campaign=shareopn

### 注册中心 (bilibili)
https://github.com/bilibili/discovery


### 问题列表
https://shimo.im/docs/x8dxHkQRcdCHX8j3/read

https://shimo.im/docs/WxJp66WCtjVwKDK3


### github小知识
```
ERROR: Permission to chinawangyu/Go-000.git denied to deploy key
fatal: 无法读取远程仓库。

请确认您有正确的访问权限并且仓库存在。
```

ssh-key 重新生成密钥文件, 配置到github中未生效, 需要额外执行
```
ssh-add ~/.ssh/id_rsa_github
```
