### 学习笔记

### 小技巧

以下代码在顶部, 编译期间检查 orderRepo 必须实现了 biz.OrderResp 这个interface
```
var _ biz.OrderResp = (*orderRepo)(nil)
```


#### 案例代码
https://github.com/go-kratos/service-layout