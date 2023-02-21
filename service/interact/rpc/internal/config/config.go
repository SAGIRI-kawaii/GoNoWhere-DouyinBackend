package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource string
	}
	CacheRedis cache.ClusterConf
	Salt       string //?????????????
	UserRpc    zrpc.RpcClientConf
	VideoRpc   zrpc.RpcClientConf
}
