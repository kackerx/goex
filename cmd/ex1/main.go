package main

import (
	"fmt"

	"goex1/cmd/ex1/wire"
	"goex1/internal/conf"
	"goex1/internal/data/cache"
	"goex1/pkg/validator"
)

func main() {
	validator.InitTrans()

	cfg := conf.NewConfig()
	cache.InitCache(cfg)

	g, clearUp, err := wire.NewWire(cfg)
	defer clearUp()
	if err != nil {
		panic(err)
	}

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	if err = g.Run(addr); err != nil {
		panic(err)
	}
}
