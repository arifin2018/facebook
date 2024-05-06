package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// mysql://root:password@tcp(127.0.0.1:3306)/restful_api_facebook?charset=utf8mb4
func Mysql() *gorm.DB {
	dsn := "root:password@tcp(127.0.0.1:3306)/restful_api_facebook?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	return db
}

var DB = Mysql()
