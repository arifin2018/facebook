package main

import (
	"github.com/arifin2018/facebook/config"
	appFolder "github.com/arifin2018/facebook/config/app"
	"github.com/arifin2018/facebook/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.SetupEnv()
	appFolder.InitProvider()
	appFolder.InitProviderRequestValidator()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "http://127.0.0.1:5500,http://127.0.0.1:3000,https://rtc.lenna.ai",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	config.Logger(app)
	routes.Router(app)
	app.Listen(":3000")
}
