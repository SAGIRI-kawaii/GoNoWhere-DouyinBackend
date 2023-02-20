package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	videos "mini-douyin/service/video/model"
	"mini-douyin/service/video/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	VideoModel videos.VideosModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		VideoModel: videos.NewVideosModel(conn),
	}
}
