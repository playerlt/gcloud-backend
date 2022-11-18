package svc

import (
	"gcloud/core/internal/config"
	"gcloud/core/internal/middleware"
	"gcloud/core/models"
	"gcloud/fileud/rpc/fileudclient"
	"gcloud/private_repo/rpc/privaterepo"
	"gcloud/public_repo/rpc/publicrepo"
	"gcloud/share/rpc/shareclient"
	"gcloud/users/rpc/usersclient"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config         config.Config   // 配置 (core-api.yaml)
	Engine         *gorm.DB        // orm
	RDB            *redis.Client   // Redis
	Auth           rest.Middleware // auth
	FileudRPC      fileudclient.Fileud
	UsersRPC       usersclient.Users
	ShareRPC       shareclient.Share
	PrivateRepoRPC privaterepo.PrivateRepo
	PublicRepoRPC  publicrepo.PublicRepo
}

// 上下文
func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		Engine:         models.Init(c.Mysql.DataSource),
		RDB:            models.InitRedis(c.Redis.Addr),
		Auth:           middleware.NewAuthMiddleware().Handle,
		FileudRPC:      fileudclient.NewFileud(zrpc.MustNewClient(c.FileudRPC)),
		UsersRPC:       usersclient.NewUsers(zrpc.MustNewClient(c.UsersRPC)),
		ShareRPC:       shareclient.NewShare(zrpc.MustNewClient(c.ShareRPC)),
		PrivateRepoRPC: privaterepo.NewPrivateRepo(zrpc.MustNewClient(c.PrivateRepoRPC)),
		PublicRepoRPC:  publicrepo.NewPublicRepo(zrpc.MustNewClient(c.PublicRepoRPC)),
	}
}
