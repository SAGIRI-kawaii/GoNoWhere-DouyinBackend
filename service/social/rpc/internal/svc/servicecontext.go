package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mini-douyin/service/social/model/follows"
	"mini-douyin/service/social/rpc/internal/config"
	"mini-douyin/service/user/model"
)

type ServiceContext struct {
	Config      config.Config
	FollowModel follows.FollowsModel
	UserModel   model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		FollowModel: follows.NewFollowsModel(conn, c.CacheRedis),
		UserModel:   model.NewUsersModel(conn, c.CacheRedis),
	}
}
