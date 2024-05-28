package auth

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/arifin2018/facebook/config"
	"github.com/arifin2018/facebook/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var MeData = models.Me{
	User: &models.User{}, // Menggunakan pointer ke User
}

func ValidationUserValidLogin(c *fiber.Ctx) error {
	if err := c.Locals("user"); err != nil {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		email := claims["email"].(string)
		MeData.User = nil
		if err := config.DB.Where("email = ?", email).First(&MeData.User).Error; err != nil {
			var errorMessage string = err.Error()
			log.Println(errorMessage)
			return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
				"data": errorMessage,
			})
		}
		tokenKeyUserId := fmt.Sprintf("token_userid_%v", MeData.User.Id)
		getAuthorization := c.Get("Authorization")
		tokenAuthorization := strings.Split(getAuthorization, " ")
		val, err := config.Redis.Get(c.Context(), tokenKeyUserId).Result()
		if err != nil || len(tokenAuthorization) != 2 {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
				"data": err,
			})
		}
		if val != tokenAuthorization[1] {
			return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
				"messages": "Invalid or expired token",
			})
		}
		return c.Next()
	}
	return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
		"data": "failed token invalid",
	})
}

func JWTClaims(f *fiber.Ctx, user *models.User) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"email": user.Email,
		"time":  time.Now(),
		"exp":   config.DefaultConfigJwt.Exp.Unix(),
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	var codeJwt = os.Getenv("CODE_JWT")
	t, err := token.SignedString([]byte(codeJwt))
	if err != nil {
		return "", err
	}

	return t, nil
}
