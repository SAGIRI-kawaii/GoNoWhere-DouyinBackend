package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mini-douyin/service/message/model/friends"
	"mini-douyin/service/message/model/messages"
	"mini-douyin/service/user/model/users"

	"mini-douyin/service/social/model/follows"
	"mini-douyin/service/social/rpc/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	FollowModel  follows.FollowsModel
	UserModel    users.UsersModel
	FriendModel  friends.FriendsModel
	MessageModel messages.MessagesModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		FollowModel:  follows.NewFollowsModel(conn, c.CacheRedis),
		UserModel:    users.NewUsersModel(conn, c.CacheRedis),
		FriendModel:  friends.NewFriendsModel(conn, c.CacheRedis),
		MessageModel: messages.NewMessagesModel(conn, c.CacheRedis),
	}
}
