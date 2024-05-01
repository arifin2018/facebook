package routes

import (
	"github.com/arifin2018/facebook/controllers"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/Post", controllers.PostIndex)
	app.Post("/Post", controllers.PostCreate)
}
