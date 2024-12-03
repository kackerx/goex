//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"goex1/internal/appservice"
	"goex1/internal/conf"
	"goex1/internal/data"
	"goex1/internal/domain/service"
	"goex1/internal/handler"
	"goex1/internal/server"
)

var repositorySet = wire.NewSet(
	data.NewRedis,
	data.NewDb,
	data.NewData,
	data.NewUserRepo,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var appServiceSet = wire.NewSet(
	appservice.NewAppService,
	appservice.NewUserAppService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

var serverSet = wire.NewSet(server.NewHTTPServer)

var initCacheSet = wire.NewSet(initCache)

func NewWire(cfg *conf.Conf) (*gin.Engine, func(), error) {
	panic(wire.Build(
		repositorySet,
		serviceSet,
		appServiceSet,
		handlerSet,
		serverSet,
		initCacheSet,
	))
}
