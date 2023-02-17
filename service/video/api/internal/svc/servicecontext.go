package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"mini-douyin/service/video/api/internal/config"
	"mini-douyin/service/video/rpc/videoservice"
)

type ServiceContext struct {
	Config config.Config

	VideoRpc videoservice.Video
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		VideoRpc: videoservice.NewVideoService(zrpc.MustNewClient(c.VideoRpc)),
	}
}
