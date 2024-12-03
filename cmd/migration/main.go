package main

import (
	"goex1/internal/conf"
	"goex1/internal/data"
	"goex1/internal/server"
)

func main() {
	cfg := conf.NewConfig()
	db := data.NewDb(cfg)
	migrate := server.NewMigrate(db)
	migrate.Start()
}
