package services

import (
	"strconv"

	"github.com/arifin2018/facebook/models"
	"github.com/arifin2018/facebook/repositories"
	"github.com/gofiber/fiber/v2"
)

var userClass = repositories.UserClass{
	User: models.User{},
}

func GetUser(f *fiber.Ctx, users *[]models.User) {
	userClass.FindUser(f, users)
}

func FindUserById(f *fiber.Ctx, user *models.User) {
	// userClass.FindUser(f, users)
	userClass.User = *user
	userClass.FindUserById(f)
	*user = userClass.User
}

func CreateUser(f *fiber.Ctx, user *models.User) models.User {
	userClass.User = models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Image:    user.Image,
	}

	userData := userClass.CreateUser(f)
	user.Id = userData.Id
	return userData
}

func UpdateUser(f *fiber.Ctx, user *models.User) {
	userClass.UpdateUser(f, user)
}

func DeleteUser(f *fiber.Ctx, id string) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	userClass.DeleteUser(f, intId)
}
