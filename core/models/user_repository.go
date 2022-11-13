package models

import (
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	Id                 int    `gorm:"primarykey"`
	Identity           string `gorm:"type:varchar(36)"`
	UserIdentity       string `gorm:"type:varchar(36)"`
	ParentId           int64
	RepositoryIdentity string
	Ext                string
	Name               string
	CreatedAt          time.Time      `gorm:"created"`
	UpdatedAt          time.Time      `gorm:"updated"`
	DeletedAt          gorm.DeletedAt `gorm:"deleted"`
}

func (table UserRepository) TableName() string {
	return "user_repository"
}
