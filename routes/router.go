package routes

import (
	"fmt"
	"os"

	"github.com/arifin2018/facebook/controllers"
	authcontroller "github.com/arifin2018/facebook/controllers/authController"
	"github.com/arifin2018/facebook/helpers/auth"
	"github.com/arifin2018/facebook/helpers/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Static("/storage/files", "./storage/files")
	app.Static("/storage/swagger", "./storage/swagger")

	app.Get("/json", func(c *fiber.Ctx) error {
		data2 := c.Query("files")
		// Baca file JSON
		filesName := fmt.Sprintf("./storage/swagger/%v", data2)
		data, err := os.ReadFile(filesName)
		if err != nil {
			return err
		}

		// Kirim JSON sebagai respons
		return c.Status(fiber.StatusOK).Send(data)
	})

	app.Post("/auth/login", authcontroller.LoginController)

	user := app.Group("user")
	user.Get("/", controllers.UserIndex)
	user.Get("/:id", controllers.UserFindById)
	user.Post("/", controllers.UserCreate)
	user.Put("/:id", controllers.UserUpdate)
	user.Delete("/:id", controllers.UserDelete)

	api := app.Group("api")
	api.Use("/", middlewares.JWTMiddlewareCheckAuth())
	// Restricted Routes
	api.Use("/", auth.ValidationUserValidLogin)

	post := api.Group("post")
	post.Get("/", middlewares.UserRole("admin,staff"), controllers.PostIndex)
	// post.Get("/", controllers.PostIndex)
	post.Get("/:id", controllers.PostFindById)
	post.Post("/", middlewares.UserRole("admin"), controllers.PostCreate)
	post.Put("/:id", controllers.PostUpdate)
	post.Delete("/:id", controllers.PostDelete)

	postImage := api.Group("postImage")
	postImage.Get("/:post_id", controllers.PostImageIndex)
	postImage.Post("/", controllers.PostImageCreate)
	postImage.Delete("/", controllers.PostImageDelete)

	comment := api.Group("comment")
	comment.Get("/:post_id", controllers.CommentIndex)
	comment.Post("/", controllers.CommentCreate)
	comment.Delete("/", controllers.CommentDelete)

	permissions := api.Group("permission")
	permissions.Get("/", authcontroller.PermissionIndex)
	permissions.Post("/", authcontroller.PermissionCreate)
	permissions.Put("/:id", authcontroller.PermissionUpdate)
	permissions.Delete("/:id", authcontroller.PermissionDelete)

	roles := api.Group("role")
	roles.Get("/", authcontroller.RoleIndex)
	roles.Post("/", authcontroller.RoleCreate)
	roles.Put("/:id", authcontroller.RoleUpdate)
	roles.Delete("/:id", authcontroller.RoleDelete)
}
