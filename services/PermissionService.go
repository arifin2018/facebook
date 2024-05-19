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

func UpdatePermission(f *fiber.Ctx, id int, permission *models.Permission) error {
	permission.ID = uint(id)
	PermissionRepository.Permission = permission
	if err := PermissionRepository.UpdatePermission(f); err != nil {
		return err
	}
	return nil
}

func DeletePermission(f *fiber.Ctx, id int) error {
	PermissionRepository.Permission.ID = uint(id)
	if err := PermissionRepository.DeletePermission(f); err != nil {
		return err
	}
	return nil
}
