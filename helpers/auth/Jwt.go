package auth

import (
	"github.com/arifin2018/facebook/config"
	"github.com/arifin2018/facebook/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

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
