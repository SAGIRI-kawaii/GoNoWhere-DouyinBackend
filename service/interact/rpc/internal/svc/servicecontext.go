package svc

import (
	"mini-douyin/service/interact/model"
	"mini-douyin/service/interact/rpc/internal/config"
	"mini-douyin/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	CommentModel  model.CommentsModel
	FavoriteModel model.FavoritesModel
	UserModel     model.UsersModel
	UserRpc       userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		CommentModel:  model.NewCommentsModel(conn, c.CacheRedis),
		FavoriteModel: model.NewFavoritesModel(conn, c.CacheRedis),
		UserRpc:       userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
