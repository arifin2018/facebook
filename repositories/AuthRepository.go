package repositories

import (
	"github.com/arifin2018/facebook/models"
	"github.com/gofiber/fiber/v2"
)

func (userClass *UserClass) FindByEmail(f *fiber.Ctx, user *models.User) {
	result := DB.Where("email = ?", userClass.User.Email).First(user)
	if result.Error != nil {
		panic(result.Error.Error())
	}
}