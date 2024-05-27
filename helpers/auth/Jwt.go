package auth

import (
	"os"

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
			return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
				"data": errorMessage,
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
