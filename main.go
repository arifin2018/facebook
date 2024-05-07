package main

import (
	"github.com/arifin2018/facebook/config"
	"github.com/arifin2018/facebook/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.Logger(app)
	routes.Router(app)
	app.Listen(":3000")
}
