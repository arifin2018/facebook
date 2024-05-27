package middlewares

import (
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func JWTMiddlewareCheckAuth() fiber.Handler {
	var codeJwt = os.Getenv("CODE_JWT")
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(codeJwt)},
	})
}
