package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	common "zerooj/common/init"
)

type Config struct {
	zrpc.RpcServerConf
	Database    common.DatabaseConfig
	RedisClient common.RedisConf
}
