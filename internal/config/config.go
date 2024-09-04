package config

import (
	"oracleX-backend/pkg/zero-contrib/jwtx"

	"github.com/zeromicro/go-zero/gateway"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	Gateway gateway.GatewayConf

	JWT []*jwtx.Config `json:"JWT"`

	TelegramBotConfig []TelegramBotConfig `json:"TelegramBotConfig"`

	MysqlDB struct {
		DataSource string `json:"DataSource"`
	} `json:"MysqlDB"`

	RedisDB struct {
		Addr   string `json:"Addr"`
		DB     int    `json:"DB"`
		Passwd string `json:"Passwd"`
	} `json:"RedisDB"`
}

type TelegramBotConfig struct {
	Name   string
	Secret string
	Expire int64
}
