package config

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Logger(app *fiber.App) {
	file, err := os.OpenFile("./storage/logs/"+time.Now().Format("01-02-2006")+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	app.Use(logger.New(logger.Config{
		Output: file,
		Format: "body : ${body} | queryParams : ${queryParams} | reqHeaders : ${reqHeaders} | ${time} | ${status} | ${latency} | ${ip} | ${method} | url : ${url} | path : ${path} | route : ${route} | ${error}\n",
	}))
	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
