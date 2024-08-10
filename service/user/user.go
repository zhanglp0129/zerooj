package main

import (
	"flag"
	"fmt"

	"zerooj/service/user/internal/config"
	baseinfoServer "zerooj/service/user/internal/server/baseinfo"
	followServer "zerooj/service/user/internal/server/follow"
	otherServer "zerooj/service/user/internal/server/other"
	profileServer "zerooj/service/user/internal/server/profile"
	registerloginServer "zerooj/service/user/internal/server/registerlogin"
	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterRegisterLoginServer(grpcServer, registerloginServer.NewRegisterLoginServer(ctx))
		user.RegisterBaseInfoServer(grpcServer, baseinfoServer.NewBaseInfoServer(ctx))
		user.RegisterProfileServer(grpcServer, profileServer.NewProfileServer(ctx))
		user.RegisterFollowServer(grpcServer, followServer.NewFollowServer(ctx))
		user.RegisterOtherServer(grpcServer, otherServer.NewOtherServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
