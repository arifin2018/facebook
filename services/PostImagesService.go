package services

import (
	"fmt"

	"github.com/arifin2018/facebook/helpers/handlers"
	"github.com/arifin2018/facebook/models"
	"github.com/gofiber/fiber/v2"
)
var RequestMultipleUploadFile handlers.RequestMultipleUploadFile

func CreatePostImage(f *fiber.Ctx,postImage *models.PostImages)  {
	RequestMultipleUploadFile = &models.PostImages{
		PostId: postImage.PostId,
	}
	err,data := RequestMultipleUploadFile.RequestMultipleUploadFile(f)
	fmt.Println(err)
	fmt.Println(data)
	// handlers.RequestMultipleUploadFile(f,postImages)
}