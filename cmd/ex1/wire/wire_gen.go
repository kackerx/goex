// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"goex1/internal/appservice"
	"goex1/internal/conf"
	"goex1/internal/data"
	"goex1/internal/data/cache"
	"goex1/internal/domain/service"
	"goex1/internal/handler"
	"goex1/internal/server"
)

// Injectors from wire.go:

func NewWire(cfg *conf.Conf) (*gin.Engine, func(), error) {
	handlerHandler := handler.NewHandler()
	appService := appservice.NewAppService()
	serviceService := service.NewService()
	db := data.NewDb(cfg)
	client := data.NewRedis(cfg)
	dataData := data.NewData(db, client)
	userRepo := data.NewUserRepo(dataData)
	userService := service.NewUserService(serviceService, userRepo)
	userAppService := appservice.NewUserAppService(appService, userService)
	userHandler := handler.NewUserHandler(handlerHandler, userAppService)
	engine := server.NewHTTPServer(cfg, userHandler)
	return engine, func() {
	}, nil
}

// wire.go:

var repositorySet = wire.NewSet(data.NewRedis, data.NewDb, data.NewData, data.NewUserRepo)

var serviceSet = wire.NewSet(service.NewService, service.NewUserService)

var appServiceSet = wire.NewSet(appservice.NewAppService, appservice.NewUserAppService)

var handlerSet = wire.NewSet(handler.NewHandler, handler.NewUserHandler)

var serverSet = wire.NewSet(server.NewHTTPServer)

var initCacheSet = wire.NewSet(initCache)

func initCache(cfg *conf.Conf) {
	cache.InitCache(cfg)
}
