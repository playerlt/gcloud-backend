package models

import (
	"time"

	"gorm.io/gorm"
)

type ShareBasic struct {
	Id                     int    `gorm:"primarykey"`
	Identity               string `gorm:"type:varchar(36)"`
	UserIdentity           string `gorm:"type:varchar(36)"`
	UserRepositoryIdentity string `gorm:"type:varchar(36)"`
	RepositoryIdentity     string `gorm:"type:varchar(36)"`
	ExpiredTime            int
	ClickNum               int
	Desc                   string
	CreatedAt              time.Time      `gorm:"created"`
	UpdatedAt              time.Time      `gorm:"updated"`
	DeletedAt              gorm.DeletedAt `gorm:"deleted"`
}

func (table ShareBasic) TableName() string {
	return "share_basic"
}
