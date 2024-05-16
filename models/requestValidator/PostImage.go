package requestValidator

import (
	"fmt"
	"os"
	"strconv"

	"github.com/arifin2018/facebook/helpers/handlers"
	"github.com/arifin2018/facebook/models"
	"github.com/go-playground/validator/v10"
)

func InitProviderValidatorRequestPostImage() {
	validatorPostImageFileSize(&models.PostImages{})
}

func validatorPostImageFileSize(postImage *models.PostImages) {
	handlers.Validate.RegisterValidation("file-size-post-image", func(fl validator.FieldLevel) bool {
		maxSize, err := strconv.Atoi(fl.Param())
		if err != nil {
			fmt.Println(err)
			return false
		}


		filePath, ok := fl.Field().Interface().(string)
		if !ok {
			fmt.Println(ok)
			return false
		}
		// Membaca informasi file
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			return false
		}
		return fileInfo.Size() <= int64(maxSize)
	})
}