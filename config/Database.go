package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Context = context.Background()
var DB = Mysql()
var Redis = redis.NewClient(&redis.Options{
	Addr:     "35.240.145.253:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

// mysql://root:password@tcp(127.0.0.1:3306)/restful_api_facebook?charset=utf8mb4
// migrate -database mysql://root:password@tcp(127.0.0.1:3306)/restful_api_facebook?charset=utf8mb4 -path db/migrations up
// migrate -database mysql://root:password@tcp(127.0.0.1:3306)/restful_api_facebook?charset=utf8mb4 -path db/migrations down
// migrate -database mysql://root:password@tcp(127.0.0.1:3306)/restful_api_facebook?charset=utf8mb4 -path db/migrations down

// migrate -path db/migrations -database mysql://root:password@tcp(127.0.0.1:3306)/restful_api_facebook?charset=utf8mb4 force 3
func Mysql() *gorm.DB {
	file, err := os.OpenFile("./storage/logs/database/"+time.Now().Format("01-02-2006")+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	newLogger := logger.New(
		log.New(file, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	dsn := "root:Arifin123!@tcp(35.240.145.253:3306)/restful_api_facebook?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	return db
}
