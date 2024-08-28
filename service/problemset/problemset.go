package main

import (
	"flag"
	"fmt"

	"zerooj/service/problemset/internal/config"
	exampleServer "zerooj/service/problemset/internal/server/example"
	hintServer "zerooj/service/problemset/internal/server/hint"
	judgedataServer "zerooj/service/problemset/internal/server/judgedata"
	languageServer "zerooj/service/problemset/internal/server/language"
	problemServer "zerooj/service/problemset/internal/server/problem"
	submitServer "zerooj/service/problemset/internal/server/submit"
	tagServer "zerooj/service/problemset/internal/server/tag"
	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/problemset.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		problemset.RegisterProblemServer(grpcServer, problemServer.NewProblemServer(ctx))
		problemset.RegisterTagServer(grpcServer, tagServer.NewTagServer(ctx))
		problemset.RegisterExampleServer(grpcServer, exampleServer.NewExampleServer(ctx))
		problemset.RegisterHintServer(grpcServer, hintServer.NewHintServer(ctx))
		problemset.RegisterJudgeDataServer(grpcServer, judgedataServer.NewJudgeDataServer(ctx))
		problemset.RegisterLanguageServer(grpcServer, languageServer.NewLanguageServer(ctx))
		problemset.RegisterSubmitServer(grpcServer, submitServer.NewSubmitServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
