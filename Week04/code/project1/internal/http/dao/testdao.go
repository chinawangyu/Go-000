package dao

import "github.com/google/wire"

// 将 dao interface 绑定到 testdao实现上
var TestProvider = wire.NewSet(NewTestDao, wire.Bind(new(Dao), new(*testdao)))

type testdao struct {
}

func NewTestDao() *testdao {
	return &testdao{}
}

func (d *testdao) GetUserInfo() string {
	return "hello test dao"
}
