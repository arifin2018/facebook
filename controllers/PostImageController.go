package controllers

import (
	"strconv"
	"strings"

	"github.com/arifin2018/facebook/helpers/handlers"
	"github.com/arifin2018/facebook/models"
	"github.com/arifin2018/facebook/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func PostImageIndex(f *fiber.Ctx) error {
	postImages := new([]models.PostImages)
	id, _ := strconv.Atoi(f.Params("post_id"))
	services.IndexPotsImages(f, id, postImages)
	return f.Status(fiber.StatusAccepted).JSON(map[string]interface{}{
		"data": postImages,
	})
}

func PostImageCreate(f *fiber.Ctx) error {
	postImage := new(models.PostImages)
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

	PostImages := services.CreatePostImage(f, postImage)
	return f.Status(fiber.StatusAccepted).JSON(map[string]interface{}{
		"data": PostImages,
	})
}

func PostImageDelete(f *fiber.Ctx) error {
	id := f.Query("id")
	dataId := strings.Split(id, ",")
	services.DeletePostImage(f, dataId)
	return nil
}
