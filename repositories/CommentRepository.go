package repositories

import (
	"github.com/arifin2018/facebook/config"
	"github.com/arifin2018/facebook/models"
	"github.com/gofiber/fiber/v2"
)

type Comment struct {
	Comment *models.Comment
}

func (comment Comment) CreateComment(f *fiber.Ctx) error {
	result := config.DB.Create(&comment.Comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
