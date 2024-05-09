package services

import (
	"encoding/json"

	"github.com/arifin2018/facebook/helpers/handlers"
	"github.com/arifin2018/facebook/models"
	"github.com/arifin2018/facebook/repositories"
	"github.com/gofiber/fiber/v2"
)

var RequestMultipleUploadFile handlers.RequestMultipleUploadFile
var PostImagesRepository = repositories.PostImages{
	PostImages: &[]models.PostImages{},
}

func IndexPotsImages(f *fiber.Ctx, id int, postImage *[]models.PostImages) {
	PostImagesRepository.PostImages = postImage
	PostImagesRepository.IndexPostImages(f, id)
}

func CreatePostImage(f *fiber.Ctx, postImage *models.PostImages) *[]models.PostImages {
	postImages := new([]models.PostImages)
	RequestMultipleUploadFile = &models.PostImages{
		PostId: postImage.PostId,
	}
	err, data := RequestMultipleUploadFile.RequestMultipleUploadFile(f)
	if err != nil {
		panic(err.Error())
	}

	dataConvertJson, _ := json.Marshal(data)
	err = json.Unmarshal([]byte(dataConvertJson), postImages)
	if err != nil {
		panic(err)
	}
	PostImagesRepository.PostImages = postImages
	PostImagesRepository.CreatePostImages(f)
	return postImages
}
