package services

import (
	"github.com/arifin2018/facebook/helpers/auth"
	"github.com/arifin2018/facebook/models"
	"github.com/gofiber/fiber/v2"
)

func Login(f *fiber.Ctx, user *models.User) bool {
	userClass.User = *user
	user = new(models.User)
	userClass.FindByEmail(f, user)
	checkPassword := auth.CheckPasswordHash(userClass.User.Password, user.Password)
	return checkPassword
}
