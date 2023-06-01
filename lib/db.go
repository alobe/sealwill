package lib

import (
	"github.com/alobe/seawill/config"
	"github.com/alobe/seawill/model"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Rds *redis.Client

func InitDB() {
	cfg := config.Get()
	var err error
	DB, err = gorm.Open(mysql.Open(cfg.Mysql.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 注册gorm模型
	err = DB.AutoMigrate(&model.User{})

	if err != nil {
		panic(err)
	}

	opt, err := redis.ParseURL(cfg.Redis.Url)
	if err != nil {
		panic(err)
	}

	Rds = redis.NewClient(opt)
}

func CloseDB() {
	// DB.Close()
	Rds.Close()
}
