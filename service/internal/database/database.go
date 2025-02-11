package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/dig"
	"gorm.io/gorm"
	"meme_coin_api/service/internal/config"
)

type MemeCoinOut struct {
	dig.Out

	MySQLMemeCoin *gorm.DB      `name:"meme_coin"`
	RedisMemeCoin *redis.Client `name:"meme_coin"`
}

const (
	mysqlMemeCoin = "meme_coin"
	redisMemeCoin = "meme_coin"
)

func NewMemeCoin(ctx context.Context, dbms config.DatabaseManagementSystem) MemeCoinOut {
	return MemeCoinOut{
		MySQLMemeCoin: newMySQL(mysqlMemeCoin, dbms.MariaDBSystems[mysqlMemeCoin]),
		RedisMemeCoin: newRedis(ctx, redisMemeCoin, dbms.RedisSystems[redisMemeCoin]),
	}
}
