package models

import (
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	Id        int    `gorm:"primarykey"`
	Identity  string `gorm:"type:varchar(36)"`
	Name      string
	Password  string
	Email     string
	Avatar    string
	Capacity  int
	CreatedAt time.Time      `gorm:"created"`
	UpdatedAt time.Time      `gorm:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"deleted"`
}

func (table UserBasic) TableName() string {
	return "user_basic"
}
