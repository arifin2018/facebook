package requestValidator

import (
	"errors"

	"github.com/arifin2018/facebook/config"
	"github.com/arifin2018/facebook/helpers/handlers"
	"github.com/arifin2018/facebook/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func HandlersValidateStructRole(f *fiber.Ctx, permission *models.Role) error {
	if err := handlers.Validate.Struct(permission); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, validationError := range validationErrors {
			return errors.New(validationError.Error())
		}
	}
	return nil
}

func InitProviderValidatorRequestRole() {
	validatorRoleUniqueDataColumn(&models.Role{})
}

func validatorRoleUniqueDataColumn(role *models.Role) {
	handlers.Validate.RegisterValidation("role-unique-data-column", func(fl validator.FieldLevel) bool {
		var count int64
		field := fl.Field().String()
		fieldName := fl.Param()
		config.DB.Model(role).Where(fieldName+" = ?", field).Count(&count)
		return count == 0
	})
}
