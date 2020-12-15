// +build wireinject   此处含义是运行时防止 build 此文件
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/google/wire"
	"project/internal/http/dao"
	"project/internal/http/service"
)

func NewService() (*service.Service, error) {
	wire.Build(service.New, dao.Provider)
	//wire.Build(service.New, dao.NewTestDao())
	return nil, nil
}

func NewTestService() (*service.Service, error) {
	wire.Build(service.New, dao.TestProvider)
	//wire.Build(service.New, dao.NewTestDao())
	return nil, nil
}
