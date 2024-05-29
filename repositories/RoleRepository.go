package repositories

import (
	"log"

	"github.com/arifin2018/facebook/config"
	"github.com/arifin2018/facebook/models"
	"github.com/gofiber/fiber/v2"
)

type Role struct {
	Role *models.Role
}

func (Role Role) IndexRole(f *fiber.Ctx, role *[]models.Role) error {
	if err := config.DB.Find(role); err != nil {
		log.Println(err.Error)
		return err.Error
	}
	return nil
}

func (Role Role) CreateRole(f *fiber.Ctx) error {
	if err := config.DB.Create(Role.Role); err != nil {
		log.Println(err.Error)
		return err.Error
	}
	return nil
}

func (Role Role) UpdateRole(f *fiber.Ctx) error {
	if err := config.DB.Model(Role.Role).Omit("Created_at").Updates(Role.Role); err != nil {
		log.Println(err.Error)
		return err.Error
	}
	return nil
}

func (Role Role) DeleteRole(f *fiber.Ctx) error {
	if err := config.DB.Delete(Role.Role); err != nil {
		return err.Error
	}
	return nil
}
