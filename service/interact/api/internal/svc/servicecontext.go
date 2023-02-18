package svc

import (
	"mini-douyin/service/interact/api/internal/config"
	"mini-douyin/service/interact/rpc/interactclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	InteractRpc interactclient.Interact
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		InteractRpc: interactclient.NewInteract(zrpc.MustNewClient(c.InteractRpc)),
	}
}
