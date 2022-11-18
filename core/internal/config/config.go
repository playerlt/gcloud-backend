package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	Mysql struct {
		DataSource string
	}
	Redis struct {
		Addr string
	}
	FileudRPC      zrpc.RpcClientConf
	UsersRPC       zrpc.RpcClientConf
	ShareRPC       zrpc.RpcClientConf
	PostsRPC       zrpc.RpcClientConf
	PrivateRepoRPC zrpc.RpcClientConf
	PublicRepoRPC  zrpc.RpcClientConf
}
