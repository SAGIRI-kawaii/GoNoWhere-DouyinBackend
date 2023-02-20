package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		FavorModel:  videos.NewFavoritesModel(conn, c.CacheRedis),
		FollowModel: videos.NewFollowsModel(conn, c.CacheRedis),
		UserModel:   videos.NewUsersModel(conn, c.CacheRedis),
		VideoModel:  videos.NewVideosModel(conn, c.CacheRedis),
	}
}
