package services

import (
	"github.com/arifin2018/facebook/models"
	"github.com/arifin2018/facebook/repositories"
	"github.com/gofiber/fiber/v2"
)

var PermissionRepository = &repositories.Permissions{
	Permission: &models.Permission{},
}

func IndexPermission(f *fiber.Ctx, permission *[]models.Permission) error {
	if err := PermissionRepository.IndexPermission(f, permission); err != nil {
		return err
	}
	return nil
}

func CreatePermission(f *fiber.Ctx, permission *models.Permission) error {
	PermissionRepository.Permission = permission
	if err := PermissionRepository.CreatePermission(f); err != nil {
		return err
	}
	return nil
}
