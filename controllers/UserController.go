package controllers

import (
	"log"
	"os"
	"strconv"

	"github.com/arifin2018/facebook/helpers/auth"
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
	paramsIdInt, _ := strconv.Atoi(params)
	user.Id = uint(paramsIdInt)
	services.FindUserById(f, user)
	return f.Status(fiber.StatusAccepted).JSON(map[string]interface{}{
		"data": user,
	})
}

func UserCreate(f *fiber.Ctx) error {
	user := new(models.User)
	f.BodyParser(user)
	password := user.Password
	user.Password, _ = auth.HashPassword(password)
	err, destinationFile := handlers.RequestUploadFile(f, user)
	if err != nil {
		log.Println(err.Error())
		return f.Status(fiber.StatusAccepted).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	user.Image = destinationFile
	if err := handlers.Validate.Struct(user); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, validationError := range validationErrors {
			return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
				"data": validationError.Error(),
			})
		}
	}

	if err := services.CreateUser(f, user); err != nil {
		if err := os.Remove(user.Image); err != nil {
			log.Println(err.Error())
			return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
				"data": err.Error(),
			})
		}
		log.Println(err.Error())
		return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	if err := services.FindUserById(f, user); err != nil {
		log.Println(err.Error())
		return f.Status(fiber.StatusForbidden).JSON(map[string]interface{}{
			"data":  user,
			"token": nil,
			"exp":   nil,
			"error": err.Error(),
		})
	}
	return f.Status(fiber.StatusAccepted).JSON(map[string]interface{}{
		"data": user,
	})
}

func UserUpdate(f *fiber.Ctx) error {
	params := f.Params("id")
	user := new(models.User)
	paramsIdInt, _ := strconv.Atoi(params)
	user.Id = uint(paramsIdInt)
	services.FindUserById(f, user)
	f.BodyParser(user)

	if user.Password != "" {
		password := user.Password
		user.Password, _ = auth.HashPassword(password)
	}

	services.UpdateUser(f, user)

	return f.Status(fiber.StatusAccepted).JSON(map[string]interface{}{
		"data": user,
	})
}

func UserDelete(f *fiber.Ctx) error {
	params := f.Params("id")
	paramsIdInt, _ := strconv.Atoi(params)
	services.DeleteUser(f, paramsIdInt)
	return f.Status(fiber.StatusAccepted).JSON(map[string]interface{}{
		"data": models.User{},
	})
}
