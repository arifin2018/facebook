package controllers

import (
	"github.com/arifin2018/facebook/helpers"
	"github.com/arifin2018/facebook/helpers/handlers"
	"github.com/arifin2018/facebook/models"
	"github.com/arifin2018/facebook/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserIndex(f *fiber.Ctx) error {
	users := []models.User{}
	services.GetUser(f, &users)
	return f.Status(fiber.StatusAccepted).JSON(map[string]interface{}{
		"data": users,
	})
}

func UserFindById(f *fiber.Ctx) error {
	params := f.Params("id")
	user := new(models.User)
	user.Id = params
	services.FindUserById(f, user)
	return f.Status(fiber.StatusAccepted).JSON(map[string]interface{}{
		"data": user,
	})
}

func UserCreate(f *fiber.Ctx) error {
	user := new(models.User)
	f.BodyParser(user)
	password := user.Password
	user.Password, _ = helpers.HashPassword(password)
	handlers.RequestUploadFile(f, user)

	if err := validate.Struct(user); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, validationError := range validationErrors {
			return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
				"data": validationError.Error(),
			})
		}
	}
	services.CreateUser(f, user)
	return f.Status(fiber.StatusAccepted).JSON(map[string]interface{}{
		"data": user,
	})
}

func UserUpdate(f *fiber.Ctx) error {
	params := f.Params("id")
	user := new(models.User)
	user.Id = params
	services.FindUserById(f, user)
	f.BodyParser(user)

	if user.Password != "" {
		password := user.Password
		user.Password, _ = helpers.HashPassword(password)
	}

	services.UpdateUser(f, user)

	return f.Status(fiber.StatusAccepted).JSON(map[string]interface{}{
		"data": user,
	})
}

func UserDelete(f *fiber.Ctx) error {
	params := f.Params("id")
	services.DeleteUser(f, params)
	return f.Status(fiber.StatusAccepted).JSON(map[string]interface{}{
		"data": models.User{},
	})
}
