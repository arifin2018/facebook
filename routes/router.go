package routes

import (
	"github.com/arifin2018/facebook/controllers"
	"github.com/gofiber/fiber/v2"
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
}
