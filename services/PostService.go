package services

import (
	"github.com/arifin2018/facebook/models"
	"github.com/arifin2018/facebook/repositories"
	"github.com/gofiber/fiber/v2"
)

func GetPost(f *fiber.Ctx) {
	userClass := repositories.UserClass{
		User: models.User{
			Name:     "arifin2018",
			Email:    "arifin2018",
			Password: "arifin2018",
		},
	}
	userClass.FindPostById(f)
}

func CreatePost(f *fiber.Ctx, user *models.User) models.User {
	userClass := repositories.UserClass{
		User: models.User{
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
			Image:    user.Image,
		},
	}
	userData := userClass.CreatePost(f)
	user.Id = userData.Id
	return userData
}
