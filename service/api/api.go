package main

import (
	"fmt"

	"mini-douyin/service/api/internal/config"
	"mini-douyin/service/api/internal/handler"
	"mini-douyin/service/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

//var configFile = flag.String("f", "etc/api.yaml", "the config file")

func main() {
	//flag.Parse()

	var c config.Config
	conf.MustLoad("common/config/rpc/api.yaml", &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
