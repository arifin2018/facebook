package repositories

import (
	"github.com/arifin2018/facebook/config"
	"github.com/arifin2018/facebook/models"
	"github.com/gofiber/fiber/v2"
)

type Permissions struct {
	Permission *models.Permission
}

func (Permission Permissions) IndexPermission(f *fiber.Ctx, permissions *[]models.Permission) error {
	if err := config.DB.Find(permissions); err != nil {
		return err.Error
	}
	return nil
}

func (Permission Permissions) CreatePermission(f *fiber.Ctx) error {
	if err := config.DB.Create(Permission.Permission); err != nil {
		return err.Error
	}
	return nil
}
