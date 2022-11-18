package svc

import (
	"gcloud/core/models"
	"gcloud/fileud/rpc/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *gorm.DB // orm

}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.Init(c.Mysql.DataSource),
	}
}
