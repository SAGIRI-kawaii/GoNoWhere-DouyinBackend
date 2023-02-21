package svc

import (
	"mini-douyin/service/api/internal/config"
	"mini-douyin/service/interact/rpc/interactclient"
	"mini-douyin/service/message/rpc/douyinrelationservice"
	"mini-douyin/service/social/rpc/followclient"
	"mini-douyin/service/user/rpc/userclient"
	"mini-douyin/service/video/rpc/videoservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	InteractRpc interactclient.Interact
	MessageRpc  douyinrelationservice.DouyinRelationService
	SocialRpc   followclient.Follow
	UserRpc     userclient.User
	VideoRpc    videoservice.VideoService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		InteractRpc: interactclient.NewInteract(zrpc.MustNewClient(c.InteractRpc)),
		MessageRpc:  douyinrelationservice.NewDouyinRelationService(zrpc.MustNewClient(c.MessageRpc)),
		SocialRpc:   followclient.NewFollow(zrpc.MustNewClient(c.SocialRpc)),
		UserRpc:     userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		VideoRpc:    videoservice.NewVideoService(zrpc.MustNewClient(c.VideoRpc)),
	}
}
