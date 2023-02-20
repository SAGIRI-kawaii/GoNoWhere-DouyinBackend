package svc

import (
	"mini-douyin/service/user/model/logins"
	"mini-douyin/service/user/model/users"

	"mini-douyin/service/user/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	UserModel  users.UsersModel
	LoginModel logins.LoginsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		UserModel:  users.NewUsersModel(conn, c.CacheRedis),
		LoginModel: logins.NewLoginsModel(conn, c.CacheRedis),
	}
}
