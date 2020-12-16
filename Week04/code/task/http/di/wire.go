// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/google/wire"
	"task/http/internal/dao"
	"task/http/internal/service"
)

func NewUserService() (*service.Service) {
	panic(wire.Build(service.NewUser, dao.Provider))
}
