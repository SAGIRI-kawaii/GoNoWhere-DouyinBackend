package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mini-douyin/service/message/model"
	"mini-douyin/service/message/rpc/internal/config"
	"mini-douyin/service/user/model"
)

type ServiceContext struct {
	Config config.Config

	MessageModel follows.MessagesModel
	UserModel    model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		MessageModel: follows.NewMessagesModel(conn, c.CacheRedis),
		UserModel:    model.NewUsersModel(conn, c.CacheRedis),
	}
}
