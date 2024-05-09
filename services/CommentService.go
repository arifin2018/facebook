package services

import (
	"github.com/arifin2018/facebook/models"
	"github.com/arifin2018/facebook/repositories"
	"github.com/gofiber/fiber/v2"
)

var CommentRepository = repositories.Comment{
	Comment: &models.Comment{},
}

func IndexComment(f *fiber.Ctx, id int, postImage *[]models.PostImages) {}

func CreateComment(f *fiber.Ctx, comment *models.Comment) error {
	CommentRepository.Comment = comment
	if err := CommentRepository.CreateComment(f); err != nil {
		return err
	}
	return nil
}
func DeleteComment(f *fiber.Ctx, id int, postImage *[]models.PostImages) {}
