package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mini-douyin/service/message/model/friends"

	"mini-douyin/service/social/model/follows"
	"mini-douyin/service/social/rpc/internal/config"
	"mini-douyin/service/user/model"
)

type ServiceContext struct {
	Config      config.Config
	FollowModel follows.FollowsModel
	UserModel   model.UsersModel
	FriendModel friends.FriendsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		FollowModel: follows.NewFollowsModel(conn, c.CacheRedis),
		UserModel:   model.NewUsersModel(conn, c.CacheRedis),
		FriendModel: friends.NewFriendsModel(conn, c.CacheRedis),
	}
}
