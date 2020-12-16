// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"task/http/internal/dao"
	"task/http/internal/service"
)

// Injectors from wire.go:

func NewUserService() *service.Service {
	userDao := dao.NewUser()
	serviceService := service.NewUser(userDao)
	return serviceService
}
