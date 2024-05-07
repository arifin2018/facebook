package handlers

import (
	"fmt"
	"time"

	"github.com/arifin2018/facebook/models"
	"github.com/gofiber/fiber/v2"
)

func RequestUploadFile(f *fiber.Ctx, user *models.User) error {
	file, err := f.FormFile("image")
	if err != nil {
		// Handle error
		panic(err)
	}
	destination := fmt.Sprintf("./storage/files/%s-%s", time.Now().Format("20060102150405"), file.Filename)
	if err := f.SaveFile(file, destination); err != nil {
		return err
	}
	user.Image = destination
	return nil
}
