package routes

import (
	"github.com/arifin2018/facebook/controllers"
	"github.com/arifin2018/facebook/helpers/auth"
	"github.com/arifin2018/facebook/helpers/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	// app.Post("testLogin", func(c *fiber.Ctx) error {
	// 	user := c.FormValue("user")
	// 	pass := c.FormValue("pass")

	// 	// Throws Unauthorized error
	// 	if user != "john" || pass != "doe" {
	// 		return c.SendStatus(fiber.StatusUnauthorized)
	// 	}

	// 	// Create the Claims
	// 	claims := jwt.MapClaims{
	// 		"name":  "John Doe",
	// 		"admin": true,
	// 		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	// 	}

	// 	// Create token
	// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 	// Generate encoded token and send it as response.
	// 	t, err := token.SignedString([]byte("secret"))
	// 	if err != nil {
	// 		return c.SendStatus(fiber.StatusInternalServerError)
	// 	}

	// 	return c.JSON(fiber.Map{"token": t})
	// })
	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	// }))
	// app.Get("/testrestricted", func(c *fiber.Ctx) error {
	// 	user := c.Locals("user").(*jwt.Token)
	// 	claims := user.Claims.(jwt.MapClaims)
	// 	name := claims["name"].(string)
	// 	return c.SendString("Welcome " + name)
	// })

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
	post.Post("/", controllers.PostCreate)
}
