package main

import (
	"fmt"

	"mini-douyin/service/interact/rpc/interact"
	"mini-douyin/service/interact/rpc/internal/config"
	"mini-douyin/service/interact/rpc/internal/server"
	"mini-douyin/service/interact/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//var configFile = flag.String("f", "etc/interact.yaml", "the config file")

func main() {
	//flag.Parse()

	var c config.Config
	conf.MustLoad("common/config/rpc/interact.yaml", &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		interact.RegisterInteractServer(grpcServer, server.NewInteractServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
