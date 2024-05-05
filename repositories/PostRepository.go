package repositories

import (
	"github.com/arifin2018/facebook/config"
	"github.com/arifin2018/facebook/models"
	"github.com/gofiber/fiber/v2"
)

type PostClass struct {
	Post models.Post
}

func (post *PostClass) FindPost(f *fiber.Ctx, posts *[]models.Post) {
	result := config.DB.Find(posts)
	if result.Error != nil {
		panic(result.Error.Error())
	}
}

func (post *PostClass) CreatePost(f *fiber.Ctx) {
	result := config.DB.Create(&post.Post)
	if result.Error != nil {
		panic(result.Error.Error())
	}
}
