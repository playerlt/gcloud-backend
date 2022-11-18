package models

import (
	"gcloud/core/define"
	"log"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
	初始化数据库
*/
func Init(dataSource string) *gorm.DB {
	engine, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	// engine, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		log.Printf("Xorm New Engine Error:%v", err)
		return nil
	}
	//dbAutoMigrate(engine)
	return engine
}

/*
	初始化redis
*/
func InitRedis(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: define.RedisPassword, // no password set
		DB:       0,                    // use default DB
	})
}

// 自动迁移表结构
func dbAutoMigrate(DB *gorm.DB) {
	_ = DB.AutoMigrate(
		&GongdeBasic{},
		&PostsBasic{},
		&PostsCommentBasic{},
		&PostsFeedback{},
		&PublicRepository{},
		&RepositoryPool{},
		&ShareBasic{},
		&UserBasic{},
		&UserRepository{},
	)
}
