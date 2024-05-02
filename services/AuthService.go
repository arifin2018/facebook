package services

import (
	"errors"

	"github.com/arifin2018/facebook/helpers/auth"
	"github.com/arifin2018/facebook/models"
	"github.com/gofiber/fiber/v2"
)

func Login(f *fiber.Ctx, user *models.User) (bool, error, *models.User) {
	userClass.User = *user
	user = new(models.User)
	err := userClass.FindByEmail(f, user)
	if err != nil {
		return false, err, nil
	}
	userFind := user
	checkPassword := auth.CheckPasswordHash(userClass.User.Password, user.Password)
	return checkPassword, errors.New("Password doesn't match"), userFind
}
