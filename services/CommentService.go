package services

import (
	"strconv"
	"strings"

	"github.com/arifin2018/facebook/models"
	"github.com/arifin2018/facebook/repositories"
	"github.com/gofiber/fiber/v2"
)

var CommentRepository = repositories.Comment{
	Comment: &models.Comment{},
}

func IndexCommentByPostId(f *fiber.Ctx, id int, comments *[]models.Comment) error {
	CommentRepository.Comment.ID = uint(id)
	if err := CommentRepository.IndexCommentByPostId(f, id, comments); err != nil {
		return err
	}
	return nil
}

func CreateComment(f *fiber.Ctx, comment *models.Comment) error {
	CommentRepository.Comment = comment
	if err := CommentRepository.CreateComment(f); err != nil {
		return err
	}
	return nil
}
func DeleteComment(f *fiber.Ctx, id string) error {
	dataId := strings.Split(id, ",")
	for _, v := range dataId {
		id, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		CommentRepository.Comment.ID = uint(id)
		if result := CommentRepository.DeleteComment(f); result != nil {
			return result
		}
	}
	return nil
}
