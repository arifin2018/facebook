package controllers

import (
	"log"
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
	postIdInt, err := strconv.Atoi(post_id)
	if err != nil {
		log.Println(err.Error())
		return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	services.IndexCommentByPostId(f, postIdInt, comments)
	return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
		"data": comments,
	})
}

func CommentCreate(f *fiber.Ctx) error {
	comment := new(models.Comment)
	if err := f.BodyParser(comment); err != nil {
		log.Println(err.Error())
		return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	if err := handlers.Validate.Struct(comment); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, validationError := range validationErrors {
			return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
				"data": validationError.Error(),
			})
		}
	}
	comment.UserId = uint(auth.MeData.User.Id)
	if err := services.CreateComment(f, comment); err != nil {
		log.Println(err.Error())
		return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
		"data": comment,
	})
}

func CommentDelete(f *fiber.Ctx) error {
	id := f.Query("id")
	if err := services.DeleteComment(f, id); err != nil {
		log.Println(err.Error())
		return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
		"data": models.Comment{},
	})
}
