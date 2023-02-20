package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mini-douyin/service/message/model/messages"
	"mini-douyin/service/message/model/users"
	"mini-douyin/service/message/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	MessageModel messages.MessagesModel
	UserModel    users.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		MessageModel: messages.NewMessagesModel(conn, c.CacheRedis),
		UserModel:    users.NewUsersModel(conn, c.CacheRedis),
	}
}
