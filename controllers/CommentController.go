package controllers

import (
	"strconv"

	"github.com/arifin2018/facebook/helpers/auth"
	"github.com/arifin2018/facebook/helpers/handlers"
	"github.com/arifin2018/facebook/models"
	"github.com/arifin2018/facebook/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CommentIndex(f *fiber.Ctx) error {
	post_id := f.Params("post_id")	
	comments := new([]models.Comment)
	postIdInt,err := strconv.Atoi(post_id)
	if err != nil {
		panic(err)
	}
	services.IndexCommentByPostId(f,postIdInt,comments)
	return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
		"data":comments,
	})
}

func CommentCreate(f *fiber.Ctx) error {
	comment := new(models.Comment)
	if err := f.BodyParser(comment); err != nil {
		panic(err.Error())
	}
	if err := handlers.Validate.Struct(comment); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, validationError := range validationErrors {
			return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
				"data": validationError.Error(),
			})
		}
	}
	userid, _ := strconv.Atoi(auth.MeData.User.Id)
	comment.UserId = uint(userid)
	if err := services.CreateComment(f, comment); err != nil {
		return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
		"data": comment,
	})
}

func CommentDelete(f *fiber.Ctx) error {
	return nil
}
