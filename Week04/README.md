### 学习笔记

#### 1. 编码小技巧

以下代码在顶部, 编译期间检查 orderRepo 必须实现了 biz.OrderResp 这个interface
```
var _ biz.OrderResp = (*orderRepo)(nil)
```

#### 2. DI依赖注入管理
引入wire包管理依赖
```
github.com/google/wire

执行命令,生成可执行文件
go build  /Users/wenba/go/pkg/mod/github.com/google/wire@v0.4.0/cmd/wire/main.go

编写wire文件
// +build wireinject 编译时忽略此文件

编译依赖管理-生成 gen 文件
wire di/*


```

参考代码: https://github.com/DrmagicE/wire-examples/blob/master/example/di/wire.go



#### 案例代码
https://github.com/go-kratos/service-layout

https://github.com/bilibili/kratos-demo