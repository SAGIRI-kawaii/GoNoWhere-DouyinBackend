package svc

import (
	videos "mini-douyin/service/video/model"
	"mini-douyin/service/video/rpc/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	VideoModel  videos.VideosModel
	FavorModel  videos.FavoritesModel
	FollowModel videos.FollowsModel
	UserModel   videos.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
