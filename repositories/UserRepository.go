package repositories

import (
	"github.com/arifin2018/facebook/config"
	"github.com/arifin2018/facebook/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserClass struct {
	User models.User
}

func (user *UserClass) FindUser(f *fiber.Ctx, users *[]models.User) *gorm.DB {
	result := config.DB.Model(&users).Preload("UserRole").Preload("UserRole.RolePermisssion").Find(&users)
	if result.Error != nil {
		panic(result.Error.Error())
	}
	return result
}

func (user *UserClass) FindUserById(f *fiber.Ctx) *gorm.DB {
	result := config.DB.Model(&user.User).Preload("UserRole").Preload("UserRole.RolePermisssion").First(&user.User, user.User.Id)
	if result.Error != nil {
		panic(result.Error.Error())
	}
	return result
}

func (user *UserClass) CreateUser(f *fiber.Ctx) error {
	tx := config.DB.Begin()

	if err := tx.Select("Name", "Email", "Password", "Image").Create(&user.User).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (user *UserClass) UpdateUser(f *fiber.Ctx, userUpdate *models.User) models.User {
	config.DB.Save(&userUpdate)
	return user.User
}

func (user *UserClass) DeleteUser(f *fiber.Ctx, id int) error {
	err := config.DB.Delete(&user.User, id).Error
	if err != nil {
		return err
	}
	return nil
}
