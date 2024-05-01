package controllers

import (
	"fmt"
	"time"

	"github.com/arifin2018/facebook/models"
	"github.com/arifin2018/facebook/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func PostIndex(f *fiber.Ctx) error {
	user := new(models.User)
	f.BodyParser(user)
	services.GetPost(f)
	return nil
}

func PostCreate(f *fiber.Ctx) error {
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
	services.CreatePost(f, user)
	return f.Status(fiber.StatusBadRequest).JSON(user)
}
