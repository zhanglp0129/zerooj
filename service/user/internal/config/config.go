package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	common2 "zerooj/common"
)

type Config struct {
	zrpc.RpcServerConf
	Database    common2.DatabaseConfig
	RedisClient common2.RedisConf
}
