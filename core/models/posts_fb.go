package models

import (
	"time"

	"gorm.io/gorm"
)

type PostsFeedback struct {
	Id            int    `gorm:"primarykey"`
	Identity      string `gorm:"type:varchar(36)"`
	UserIdentity  string `gorm:"type:varchar(36)"`
	PostsIdentity string `gorm:"type:varchar(36)"`
	Type          string
	Count         int
	Read          int
	CreatedAt     time.Time      `gorm:"created"`
	UpdatedAt     time.Time      `gorm:"updated"`
	DeletedAt     gorm.DeletedAt `gorm:"deleted"`
}

func (table PostsFeedback) TableName() string {
	return "posts_fb"
}
