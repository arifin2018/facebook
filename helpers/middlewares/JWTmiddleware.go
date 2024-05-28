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
		ErrorHandler: func(f *fiber.Ctx, err error) error {
			return f.Status(fiber.StatusAccepted).JSON(map[string]interface{}{
				"messages": "Invalid or expired token",
			})
		},
	})
}
