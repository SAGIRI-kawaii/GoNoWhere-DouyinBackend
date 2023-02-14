package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mini-douyin/service/social/model/follows"
	"mini-douyin/service/social/rpc/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	FollowModel follows.FollowsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		FollowModel: follows.NewFollowsModel(conn, c.CacheRedis),
	}
}
