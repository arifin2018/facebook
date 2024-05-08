package models

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PostImages struct {
	PostId 	uint `gorm:"column:Post_id" validate:"required" json:"postId"`
	Url 	string `gorm:"column:Url" validate:"required" json:"url"`
	gorm.Model
}
func (postImage *PostImages) RequestMultipleUploadFile(f *fiber.Ctx) (error,interface{}) {
	form, err := f.MultipartForm()
	if err != nil { 
		return err, nil
	}
	postId := postImage.PostId
	postImages := []PostImages{}
	for formFieldName, fileHeaders := range form.File {
		for _, file := range fileHeaders {
			destination := fmt.Sprintf("./storage/files/%s-%s-%s",formFieldName, time.Now().Format("20060102150405"), file.Filename)
			if err := f.SaveFile(file, destination); err != nil {
				return err,nil
			}
			postImage := PostImages{}
			postImage.PostId = postId
			postImage.Url = destination
			postImages = append(postImages, postImage)
		}
	}
	return nil,postImages
}

