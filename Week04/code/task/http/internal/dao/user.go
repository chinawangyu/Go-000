package dao

import "github.com/google/wire"

type UserDao interface {
	GetUsers() string
}


var Provider = wire.NewSet(NewUser, wire.Bind(new(UserDao), new(*userDao)))

type userDao struct {
}

func NewUser() *userDao {
	return &userDao{}
}

func (d *userDao) GetUsers() string {
	return "hello user list"
}
