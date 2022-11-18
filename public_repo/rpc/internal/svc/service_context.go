package svc

import (
	"gcloud/core/models"
	"gcloud/public_repo/rpc/internal/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *gorm.DB      // orm
	RDB    *redis.Client // Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.Init(c.Mysql.DataSource),
		RDB:    models.InitRedis(c.Redis.Addr),
	}
}
