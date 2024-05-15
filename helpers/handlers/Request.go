package handlers

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/arifin2018/facebook/config"
	"github.com/arifin2018/facebook/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())

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

type RequestMultipleUploadFile interface {
	RequestMultipleUploadFile(f *fiber.Ctx) (error, interface{})
}

func SetupDefaultConfigJwt() {
	timeexp, err := strconv.Atoi(os.Getenv("TIMEEXP"))
	if err != nil {
		panic(err.Error())
	}
	config.Timeexp = timeexp
	*config.DefaultConfigJwt = config.ConfigJwt{
		Exp: time.Now().Add(time.Hour * time.Duration(config.Timeexp)),
	}
}
