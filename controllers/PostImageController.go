package controllers

import (
	"log"
	"strconv"
	"strings"

	"github.com/arifin2018/facebook/models"
	"github.com/arifin2018/facebook/services"
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
	if err := f.BodyParser(postImage); err != nil {
		log.Println(err.Error())
		return f.Status(fiber.StatusAccepted).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}

	PostImages, err := services.CreatePostImage(f, postImage)
	if err != nil {
		log.Println(err.Error())
		return f.Status(fiber.StatusAccepted).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	return f.Status(fiber.StatusAccepted).JSON(map[string]interface{}{
		"data": PostImages,
	})
}

func PostImageDelete(f *fiber.Ctx) error {
	id := f.Query("id")
	dataId := strings.Split(id, ",")
	services.DeletePostImage(f, dataId)
	return f.Status(fiber.StatusAccepted).JSON(map[string]interface{}{
		"data": models.PostImages{},
	})
}
