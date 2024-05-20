package repositories

import (
	"github.com/arifin2018/facebook/config"
	"github.com/arifin2018/facebook/models"
	"github.com/gofiber/fiber/v2"
)

type PostClass struct {
	Post models.Post
}

func (post *PostClass) FindPost(f *fiber.Ctx, posts *[]models.Post) error {
	result := config.DB.Find(posts)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (post *PostClass) FindByIDPost(f *fiber.Ctx) error {
	result := config.DB.First(&post.Post, post.Post.ID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (post *PostClass) CreatePost(f *fiber.Ctx) error {
	result := config.DB.Create(&post.Post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (post *PostClass) UpdatePost(f *fiber.Ctx) error {
	result := config.DB.Model(post.Post).UpdateColumns(&post.Post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (post *PostClass) DeletePost(f *fiber.Ctx) error {
	result := config.DB.Delete(&post.Post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
