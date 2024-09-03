package config

import (
	"oracleX-backend/pkg/zero-contrib/jwtx"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/gateway"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	Gateway gateway.GatewayConf

	MysqlDB struct {
		DataSource string `json:"DataSource"`
	} `json:"MysqlDB"`

	JWT      []*jwtx.Config  `json:"JWT"`
	Cache    cache.CacheConf `json:"Cache"`
	BizCache redis.RedisConf `json:"BizCache"`

	RedisDB struct {
		Addr   string `json:"Addr"`
		DB     int    `json:"DB"`
		Passwd string `json:"Passwd"`
	} `json:"RedisDB"`

	TelegramBotConfig []TelegramBotConfig `json:"TelegramBotConfig"`
}

type TelegramBotConfig struct {
	Name   string
	Secret string
	Expire int64
}
