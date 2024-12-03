package data

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"goex1/internal/conf"
)

type Data struct {
	db    *gorm.DB
	cache *redis.Client
}

func NewData(db *gorm.DB, cache *redis.Client) *Data {
	return &Data{db: db, cache: cache}
}

func NewDb(cfg *conf.Conf) *gorm.DB {
	master := cfg.Data.Master
	db, err := gorm.Open(mysql.Open(master.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDb, _ := db.DB()
	sqlDb.SetMaxOpenConns(master.MaxOpen)
	sqlDb.SetMaxIdleConns(master.MaxIdle)
	sqlDb.SetConnMaxLifetime(time.Minute * time.Duration(master.MaxLifeTime))

	return db
}

func NewRedis(cfg *conf.Conf) *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr:     cfg.Data.Redis.Addr,
		Password: cfg.Data.Redis.Password,
		DB:       cfg.Data.Redis.Db,
		PoolSize: cfg.Data.Redis.PoolSize,
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if _, err := cli.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	return cli
}
