package dao

import "github.com/google/wire"

//将 dao interface 绑定到 dao 实现上
var Provider = wire.NewSet(New, wire.Bind(new(Dao), new(*dao)))

type Dao interface {
	GetUserInfo() string
}

type dao struct {
}

func New() *dao {
	return &dao{}
}

func (d *dao) GetUserInfo() string {
	return "hello dao"
}
