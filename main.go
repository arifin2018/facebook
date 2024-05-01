package main

import (
	"github.com/arifin2018/facebook/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.Router(app)
	app.Listen(":3000")
}
