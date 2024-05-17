package repositories

import (
	"github.com/arifin2018/facebook/config"
	"github.com/arifin2018/facebook/models"
	"github.com/gofiber/fiber/v2"
)

type Comment struct {
	Comment *models.Comment
}

func (comment Comment) IndexCommentByPostId(f *fiber.Ctx, postId int, comments *[]models.Comment) error {
	result := config.DB.Find(&comments, "Post_id = ?", comment.Comment.ID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (comment Comment) CreateComment(f *fiber.Ctx) error {
	result := config.DB.Create(&comment.Comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (comment Comment) DeleteComment(f *fiber.Ctx) error {
	if result := config.DB.Delete(&comment.Comment); result.Error != nil {
		return result.Error
	}
	return nil
}
