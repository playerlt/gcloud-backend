package models

import (
	"time"

	"gorm.io/gorm"
)

type PostsCommentBasic struct {
	Id            int    `gorm:"primarykey"`
	Identity      string `gorm:"type:varchar(36)"`
	UserIdentity  string `gorm:"type:varchar(36)"`
	PostsIdentity string `gorm:"type:varchar(36)"`
	ReplyIdentity string `gorm:"type:varchar(36)"`
	ReplyName     string
	Content       string
	Mention       string
	Like          int
	Dislike       int
	Read          int
	CreatedAt     time.Time      `gorm:"created"`
	UpdatedAt     time.Time      `gorm:"updated"`
	DeletedAt     gorm.DeletedAt `gorm:"deleted"`
}

func (table PostsCommentBasic) TableName() string {
	return "posts_comment_basic"
}
