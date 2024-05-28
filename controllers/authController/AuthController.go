package authcontroller

import (
	"fmt"
	"log"
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
	checkLogin, findUser, err := services.Login(f, user)
	if checkLogin {
		token, err := auth.JWTClaims(f, findUser)
		if err != nil {
			log.Println(err.Error())
			return f.Status(fiber.StatusForbidden).JSON(map[string]interface{}{
				"data":  nil,
				"token": nil,
				"exp":   nil,
				"error": err.Error(),
			})
		}
		if err := services.FindUserById(f, findUser); err != nil {
			log.Println(err.Error())
			return f.Status(fiber.StatusForbidden).JSON(map[string]interface{}{
				"data":  findUser,
				"token": nil,
				"exp":   nil,
				"error": err.Error(),
			})
		}
		keyRedisToken := fmt.Sprintf("token_userid_%v", findUser.Id)
		if err := config.Redis.Set(f.Context(), keyRedisToken, token, time.Duration(config.Timeexp)*time.Hour).Err(); err != nil {
			log.Println(err)
			return f.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{
				"data":  findUser,
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
		log.Println(err.Error())
		return f.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{
			"data":  findUser,
			"token": nil,
			"exp":   nil,
			"error": err.Error(),
		})
	}
}
