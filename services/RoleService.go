package services

import (
	"github.com/arifin2018/facebook/models"
	"github.com/arifin2018/facebook/repositories"
	"github.com/gofiber/fiber/v2"
)

var RoleRepository = &repositories.Role{
	Role: &models.Role{},
}

func IndexRole(f *fiber.Ctx, role *[]models.Role) error {
	if err := RoleRepository.IndexRole(f, role); err != nil {
		return err
	}
	return nil
}

func CreateRole(f *fiber.Ctx, role *models.Role) error {
	RoleRepository.Role = role
	if err := RoleRepository.CreateRole(f); err != nil {
		return err
	}
	return nil
}

func UpdateRole(f *fiber.Ctx, id int, role *models.Role) error {
	role.ID = uint(id)
	RoleRepository.Role = role
	if err := RoleRepository.UpdateRole(f); err != nil {
		return err
	}
	return nil
}

func DeleteRole(f *fiber.Ctx, id int) error {
	RoleRepository.Role.ID = uint(id)
	if err := RoleRepository.DeleteRole(f); err != nil {
		return err
	}
	return nil
}
