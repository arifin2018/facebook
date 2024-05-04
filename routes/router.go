package routes

import (
	"fmt"

	"github.com/arifin2018/facebook/controllers"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Router(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/auth/login", controllers.LoginController)

	user := app.Group("user")
	user.Get("/", controllers.UserIndex)
	user.Get("/:id", controllers.UserFindById)
	user.Post("/", controllers.UserCreate)
	user.Put("/:id", controllers.UserUpdate)
	user.Delete("/:id", controllers.UserDelete)

	app.Get("/post", controllers.PostIndex)
	app.Post("/post", controllers.PostCreate)

	app.Use("/api", func(c *fiber.Ctx) error {
		jwtware.New(jwtware.Config{
			SigningKey: jwtware.SigningKey{Key: []byte("secret")},
		})
		return c.Next()
	})

	// Restricted Routes
	app.Get("/restricted", func(c *fiber.Ctx) error {
		if err := c.Locals("user"); err != nil {
			user := c.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			fmt.Println(claims)
			email := claims["email"].(string)
			c.Next()
			return c.SendString("Welcome " + email)
		}
		return c.SendString("failed")
	})
}
