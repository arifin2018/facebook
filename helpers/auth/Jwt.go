package auth

import (
	"github.com/arifin2018/facebook/config"
	"github.com/arifin2018/facebook/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var MeData = models.Me{
	User: models.User{},
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
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}

func ValidationUserValidLogin(c *fiber.Ctx) error {
	if err := c.Locals("user"); err != nil {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		email := claims["email"].(string)
		err := config.DB.Where("email = ?", email).First(&MeData.User)
		if err.Error != nil {
			panic(err.Error.Error())
		}
		return c.Next()
	}
	return c.SendString("failed token invalid")
}
