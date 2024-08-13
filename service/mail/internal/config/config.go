package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	SMTP        SMTPConfig
	RedisClient redis.RedisConf
}

type SMTPConfig struct {
	Host     string
	Port     uint16
	SSL      bool
	Sender   string
	Username string
	Password string
}
