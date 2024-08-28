package common

import "github.com/zeromicro/go-zero/zrpc"

type ClientConfig struct {
	User       zrpc.RpcClientConf `json:",optional"`
	Mail       zrpc.RpcClientConf `json:",optional"`
	Problemset zrpc.RpcClientConf `json:",optional"`
}
