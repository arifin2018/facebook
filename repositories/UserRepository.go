package repositories

import (
	"github.com/arifin2018/facebook/config"
	"github.com/arifin2018/facebook/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB = config.Mysql()

type UserClass struct {
	User models.User
}

func (user *UserClass) FindUser(f *fiber.Ctx, users *[]models.User) *gorm.DB {
	result := DB.Find(&users)
	if result.Error != nil {
		panic(result.Error.Error())
	}
	return result
}

func (user *UserClass) FindUserById(f *fiber.Ctx) *gorm.DB {
	result := DB.First(&user.User, user.User.Id)
	if result.Error != nil {
		panic(result.Error.Error())
	}
	return result
}

func (user *UserClass) CreateUser(f *fiber.Ctx) models.User {
	err := DB.Create(&user.User).Error
	if err != nil {
		panic(err.Error())
	}
	return user.User
}

func (user *UserClass) UpdateUser(f *fiber.Ctx, userUpdate *models.User) models.User {
	DB.Save(&userUpdate)
	return user.User
}

func (user *UserClass) DeleteUser(f *fiber.Ctx, id int) error {
	err := DB.Delete(&user.User, id).Error
	if err != nil {
		return err
	}
	return nil
}
