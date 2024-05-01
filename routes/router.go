package routes

import (
	"github.com/arifin2018/facebook/controllers"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/post", controllers.PostIndex)
	app.Post("/post", controllers.PostCreate)

	app.Get("/user", controllers.UserIndex)
	app.Get("/user/:id", controllers.UserFindById)
	app.Post("/user", controllers.UserCreate)
	app.Put("/user/:id", controllers.UserUpdate)
	app.Delete("/user/:id", controllers.UserDelete)
}
