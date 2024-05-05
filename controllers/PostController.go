package controllers

import (
	response "github.com/arifin2018/facebook/helpers/handlers/Response"
	"github.com/arifin2018/facebook/models"
	"github.com/arifin2018/facebook/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func PostIndex(f *fiber.Ctx) error {
	var post response.PaginateInterface = models.Post{}
	posts := []models.Post{}
	services.GetPost(f, &posts)
	data := response.PaginationData{
		Count:           new(int64),
		TotalPages:      new(int64),
		CountTotalPages: new(float64),
		Model:           &posts,
	}
	post.DataPagination(f, &data)
	return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
		"data":              posts,
		"count_total_pages": data.CountTotalPages,
	})
}

func PostCreate(f *fiber.Ctx) error {
	post := new(models.Post)
	f.BodyParser(post)
	services.CreatePost(f, post)
	return nil
}
