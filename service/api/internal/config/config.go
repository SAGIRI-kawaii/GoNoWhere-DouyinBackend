package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	InteractRpc zrpc.RpcClientConf
	MessageRpc  zrpc.RpcClientConf
	SocialRpc   zrpc.RpcClientConf
	UserRpc     zrpc.RpcClientConf
	VideoRpc    zrpc.RpcClientConf
}
