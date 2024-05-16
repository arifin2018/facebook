package requestValidator

import (
	"errors"

	"github.com/arifin2018/facebook/config"
	"github.com/arifin2018/facebook/helpers/handlers"
	"github.com/arifin2018/facebook/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HandlersValidateStructPost(f *fiber.Ctx, post *models.Post) error {
	if err := handlers.Validate.Struct(post); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, validationError := range validationErrors {
			return errors.New(validationError.Error())
		}
	}
	return nil
}

func InitProviderValidatorRequestPost() {
	validatorPostHasOneContentUserid(&models.Post{})
}

func validatorPostHasOneContentUserid(post *models.Post) {
	handlers.Validate.RegisterValidation("check-userid-same-content-post", func(fl validator.FieldLevel) bool {
		content := fl.Field().String()
		fieldUserId, _, _, _ := fl.GetStructFieldOKAdvanced2(fl.Parent(), "UserId")
		if result := config.DB.Where("content = ? and user_id = ?", content, fieldUserId.Uint()).First(&post); result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				return true
			}
		}
		return false
	})
}
