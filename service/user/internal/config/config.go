package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zerooj/common"
)

type Config struct {
	zrpc.RpcServerConf
	Gorm common.GormConfig
}
