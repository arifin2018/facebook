package controllers

import (
	"time"

	"github.com/arifin2018/facebook/helpers/auth"
	"github.com/arifin2018/facebook/models"
	"github.com/arifin2018/facebook/services"
	"github.com/gofiber/fiber/v2"
)

func LoginController(f *fiber.Ctx) error {
	user := new(models.User)
	f.BodyParser(user)
	checkLogin := services.Login(f, user)
	if checkLogin {
		token, err := auth.JWTClaims(f, user)
		if err != nil {
			return f.Status(fiber.StatusForbidden).JSON(map[string]interface{}{
				"data": err.Error(),
			})
		}
		return f.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{
			"data":  user,
			"token": token,
			"exp":   time.Now().Add(time.Hour * 72).Unix(),
		})
	} else {
		return f.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{
			"data":  user,
			"token": nil,
			"exp":   nil,
		})
	}
}
