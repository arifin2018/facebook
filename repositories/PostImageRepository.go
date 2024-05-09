package repositories

import (
	"github.com/arifin2018/facebook/config"
	"github.com/arifin2018/facebook/models"
	"github.com/gofiber/fiber/v2"
)

type PostImages struct {
	PostImages *[]models.PostImages
}

func (PostImages *PostImages) IndexPostImages(f *fiber.Ctx, id int) {
	config.DB.Where("Post_id = ?", id).Find(&PostImages.PostImages)
}

func (PostImages *PostImages) CreatePostImages(f *fiber.Ctx) {
	result := config.DB.Create(PostImages.PostImages)
	if result.Error != nil {
		panic(result.Error.Error())
	}
}
