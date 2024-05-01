package controllers

import (
	"fmt"
	"time"

	"github.com/arifin2018/facebook/models"
	"github.com/arifin2018/facebook/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserIndex(f *fiber.Ctx) error {
	user := new(models.User)
	f.BodyParser(user)
	services.GetUser(f)
	return nil
}

func UserCreate(f *fiber.Ctx) error {
	user := new(models.User)
	f.BodyParser(user)
	file, _ := f.FormFile("image")
	destination := fmt.Sprintf("./storage/files/%s-%s", time.Now().Format("20060102150405"), file.Filename)
	if err := f.SaveFile(file, destination); err != nil {
		return err
	}
	user.Image = destination

	if err := validate.Struct(user); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, validationError := range validationErrors {
			return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
				"data": validationError.Error(),
			})
		}
	}
	services.CreateUser(f, user)
	return f.Status(fiber.StatusBadRequest).JSON(user)
}
