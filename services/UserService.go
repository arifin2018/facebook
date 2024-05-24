package services

import (
	"log"
	"os"

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
	userClass.User = *user
	userClass.FindUserById(f)
	*user = userClass.User
}

func CreateUser(f *fiber.Ctx, user *models.User) error {
	userClass.User = models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Image:    user.Image,
		Roleid:   user.Roleid,
	}

	err := userClass.CreateUser(f)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(f *fiber.Ctx, user *models.User) {
	userClass.UpdateUser(f, user)
}

func DeleteUser(f *fiber.Ctx, id int) {
	userClass.User.Id = uint(id)
	userClass.FindUserById(f)
	e := os.Remove(userClass.User.Image)
	if e != nil {
		log.Fatal(e)
	}
	userClass.DeleteUser(f, id)
}
