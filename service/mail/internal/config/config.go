package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zerooj/common"
)

type Config struct {
	zrpc.RpcServerConf
	SMTP        SMTPConfig
	RedisClient common.RedisConf
}

type SMTPConfig struct {
	Host     string
	Port     uint16
	SSL      bool
	Sender   string
	Username string
	Password string
}
