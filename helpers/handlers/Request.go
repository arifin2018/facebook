package handlers

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/arifin2018/facebook/config"
	"github.com/arifin2018/facebook/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())

func RequestUploadFile(f *fiber.Ctx, user *models.User) (error, string) {
	file, err := f.FormFile("image")
	if err != nil {
		log.Println(err.Error())
		return errors.New(err.Error()), ""
	}

	if file.Size > (1 << 20) { // 2MB
		return errors.New("file size exceeds 2MB limit"), ""
	}
	destination := fmt.Sprintf("./storage/files/%s-%s", time.Now().Format("20060102150405"), file.Filename)
	if err := f.SaveFile(file, destination); err != nil {
		return err, ""
	}
	return nil, destination
}

type RequestMultipleUploadFile interface {
	RequestMultipleUploadFile(f *fiber.Ctx) (error, interface{})
}

func SetupDefaultConfigJwt() error {
	timeexp, err := strconv.Atoi(os.Getenv("TIMEEXP"))
	if err != nil {
		return err
	}
	config.Timeexp = timeexp
	*config.DefaultConfigJwt = config.ConfigJwt{
		Exp: time.Now().Add(time.Hour * time.Duration(config.Timeexp)),
	}
	return nil
}
