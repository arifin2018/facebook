package controllers

import (
	"github.com/arifin2018/facebook/models"
	"github.com/arifin2018/facebook/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func PostIndex(f *fiber.Ctx) error {
	return nil
}

func PostCreate(f *fiber.Ctx) error {
	post := new(models.Post)
	f.BodyParser(post)
	services.CreatePost(f, post)
	return nil
}
