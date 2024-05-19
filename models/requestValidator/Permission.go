package requestValidator

import (
	"errors"

	"github.com/arifin2018/facebook/config"
	"github.com/arifin2018/facebook/helpers/handlers"
	"github.com/arifin2018/facebook/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func HandlersValidateStructPermission(f *fiber.Ctx, permission *models.Permission) error {
	if err := handlers.Validate.Struct(permission); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, validationError := range validationErrors {
			return errors.New(validationError.Error())
		}
	}
	return nil
}

func InitProviderValidatorRequestPermission() {
	validatorPermissionUniqueDataColumn(&models.Permission{})
}

func validatorPermissionUniqueDataColumn(post *models.Permission) {
	handlers.Validate.RegisterValidation("permission-unique-data-column", func(fl validator.FieldLevel) bool {
		var count int64
		field := fl.Field().String()
		fieldName := fl.Param()
		config.DB.Model(post).Where(fieldName+" = ?", field).Count(&count)
		return count == 0
	})
}
