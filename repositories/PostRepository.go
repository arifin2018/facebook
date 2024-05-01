package repositories

import (
	"github.com/arifin2018/facebook/config"
	"github.com/arifin2018/facebook/models"
	"github.com/gofiber/fiber/v2"
)

var DB = config.Mysql()

type UserClass struct {
	User models.User
}

func (user *UserClass) FindPostById(f *fiber.Ctx) {

}

func (user *UserClass) CreatePost(f *fiber.Ctx) models.User {
	err := DB.Create(&user.User).Error
	if err != nil {
		panic(err.Error())
	}
	return user.User
}
