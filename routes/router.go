package routes

import (
	"github.com/arifin2018/facebook/controllers"
	"github.com/arifin2018/facebook/helpers/auth"
	"github.com/arifin2018/facebook/helpers/middlewares"
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

	api := app.Group("api")
	api.Use("/", middlewares.JWTMiddlewareCheckAuth())

	// Restricted Routes
	api.Use(auth.ValidationUserValidLogin)

	post := api.Group("post")
	post.Get("/", controllers.PostIndex)
	post.Get("/:id", controllers.PostFindById)
	post.Post("/", controllers.PostCreate)
	post.Put("/:id", controllers.PostUpdate)
	post.Delete("/:id", controllers.PostDelete)
}
