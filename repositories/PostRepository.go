package repositories

import (
	"github.com/arifin2018/facebook/models"
	"github.com/gofiber/fiber/v2"
)

type PostClass struct {
	Post models.Post
}

func FindPost() {

}

func (post *PostClass) CreatePost(f *fiber.Ctx) {
	result := DB.Create(&post.Post)
	if result.Error != nil {
		panic(result.Error.Error())
	}
}
