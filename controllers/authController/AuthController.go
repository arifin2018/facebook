package authcontroller

import (
	"time"

	"github.com/arifin2018/facebook/config"
	"github.com/arifin2018/facebook/helpers/auth"
	"github.com/arifin2018/facebook/models"
	"github.com/arifin2018/facebook/services"
	"github.com/gofiber/fiber/v2"
)

func LoginController(f *fiber.Ctx) error {
	user := new(models.User)
	f.BodyParser(user)
	checkLogin, err, findUser := services.Login(f, user)
	if checkLogin {
		token, err := auth.JWTClaims(f, findUser)
		if err != nil {
			return f.Status(fiber.StatusForbidden).JSON(map[string]interface{}{
				"data":  nil,
				"token": nil,
				"exp":   nil,
				"error": err.Error(),
			})
		}
		return f.Status(fiber.StatusAccepted).JSON(map[string]interface{}{
			"data":  findUser,
			"token": token,
			"exp":   time.Now().Add(time.Hour * time.Duration(config.Timeexp)).Format(time.DateTime),
			"error": nil,
		})
	} else {
		return f.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{
			"data":  findUser,
			"token": nil,
			"exp":   nil,
			"error": err.Error(),
		})
	}
}
