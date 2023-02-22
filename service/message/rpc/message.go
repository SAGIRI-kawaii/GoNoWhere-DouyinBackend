package main

import (
	"fmt"

	"mini-douyin/service/message/rpc/internal/config"
	"mini-douyin/service/message/rpc/internal/server"
	"mini-douyin/service/message/rpc/internal/svc"
	"mini-douyin/service/message/rpc/message"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//var configFile = flag.String("f", "etc/message.yaml", "the config file")

func main() {
	//flag.Parse()

	var c config.Config
	conf.MustLoad("common/config/rpc/message.yaml", &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		message.RegisterDouyinRelationServiceServer(grpcServer, server.NewDouyinRelationServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
