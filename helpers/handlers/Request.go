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

func RequestMultipleUploadFile(f *fiber.Ctx,postImages *[]models.PostImages) error {
	form, err := f.MultipartForm()
	if err != nil { 
		return err
	}
	
	for formFieldName, fileHeaders := range form.File {
		for _, file := range fileHeaders {
			destination := fmt.Sprintf("./storage/files/%s-%s-%s",formFieldName, time.Now().Format("20060102150405"), file.Filename)
			if err := f.SaveFile(file, destination); err != nil {
				return err
			}
			postImage := models.PostImages{}
			postImage.PostId = 1
			postImage.Caption = "awd"
			postImage.Url = destination
			*postImages = append(*postImages, postImage)
		}
	}
	return nil
}