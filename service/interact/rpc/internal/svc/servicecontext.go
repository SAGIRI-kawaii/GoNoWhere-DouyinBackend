package svc

import (
	"mini-douyin/service/interact/model/comments"
	"mini-douyin/service/interact/model/favorites"

	"mini-douyin/service/interact/model/users"
	"mini-douyin/service/interact/model/videos"

	"mini-douyin/service/interact/rpc/internal/config"
	"mini-douyin/service/user/rpc/userclient"
	"mini-douyin/service/video/rpc/videoservice"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	CommentModel  comments.CommentsModel
	FavoriteModel favorites.FavoritesModel
	VideoModel    videos.VideosModel
	UserModel     users.UsersModel
	UserRpc       userclient.User
	VideoRpc      videoservice.VideoService
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		CommentModel:  comments.NewCommentsModel(conn, c.CacheRedis),
		FavoriteModel: favorites.NewFavoritesModel(conn, c.CacheRedis),
		VideoModel:    videos.NewVideosModel(conn, c.CacheRedis),
		UserModel:     users.NewUsersModel(conn, c.CacheRedis),
		UserRpc:       userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		VideoRpc:      videoservice.NewVideoService(zrpc.MustNewClient(c.VideoRpc)),
	}
}
