package authcontroller

import (
	"fmt"

	"github.com/arifin2018/facebook/models"
	"github.com/arifin2018/facebook/models/requestValidator"
	"github.com/arifin2018/facebook/services"
	"github.com/gofiber/fiber/v2"
)

func PermissionIndex(f *fiber.Ctx) error {
	permissions := new([]models.Permission)
	if err := services.IndexPermission(f, permissions); err != nil {
		return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	fmt.Println(permissions)
	return nil
}

func PermissionCreate(f *fiber.Ctx) error {
	permission := new(models.Permission)
	if err := f.BodyParser(permission); err != nil {
		return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	if err := requestValidator.HandlersValidateStructPermission(f, permission); err != nil {
		return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	if err := services.CreatePermission(f, permission); err != nil {
		return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"data": err.Error(),
		})
	}
	return f.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
		"data": permission,
	})
}

func PermissionUpdate(f *fiber.Ctx) error {
	return nil
}

func PermissionDelete(f *fiber.Ctx) error {
	return nil
}
