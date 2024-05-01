package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func PostIndex(f *fiber.Ctx) error {
	return nil
}

func PostCreate(f *fiber.Ctx) error {
	return nil
}
