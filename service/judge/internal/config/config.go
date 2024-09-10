package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Run RunConfig
}

type RunConfig struct {
	WorkingDirectory string
	SrcFilename      string
	CompileCommand   string `json:",optional"`
	RunCommand       string
}
