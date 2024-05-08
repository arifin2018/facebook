package controllers

import (
	"github.com/arifin2018/facebook/helpers/handlers"
	"github.com/arifin2018/facebook/models"
	"github.com/arifin2018/facebook/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func PostImageIndex(f *fiber.Ctx) error {
	return nil
}

func PostImageCreate(f *fiber.Ctx) error {
	postImage := new(models.PostImages) 
	// postImages := []models.PostImages{}
	f.BodyParser(postImage)
	if err := f.BodyParser(postImage); err != nil {
		panic(err.Error())
	}
	if err := handlers.Validate.Struct(postImage); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, validationError := range validationErrors {
			return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
				"data": validationError.Error(),
			})
		}
	}
	
	services.CreatePostImage(f, postImage)
	return nil
}